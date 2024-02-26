package golangci_linter

import (
	"bytes"
	"fmt"
	"text/template"

	_ "embed"

	"github.com/cjlapao/common-go-dependency-tree/dependencytree"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/interfaces"
)

//go:embed default.yml
var defaultTemplate string

type GolangCILinterProvisioner struct {
	id                 string
	name               string
	languages          []string
	ctx                *context.ProvisionerContext
	cfg                *config.Config
	diagnostics        *diagnostics.Diagnostics
	config             interfaces.ProvisionerConfig
	depTree            *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx context.ProvisionerContext, providerConfig interfaces.ProvisionerConfig) *GolangCILinterProvisioner {
	result := &GolangCILinterProvisioner{
		id:        constants.GoLangCILinterProvisionerID,
		name:      "GolangCI Linter Provisioner",
		ctx:       &ctx,
		cfg:       config.Get(),
		languages: []string{"go", "golang"},
		config:    providerConfig,
	}

	result.ctx.WithValue(constants.ContextResourceName, result.name)

	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics

	return result
}

func (p *GolangCILinterProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p *GolangCILinterProvisioner) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&GolangCILinterProvisioner{})
	if depTree == nil {
		p.ctx.LogError("Error getting dependency tree")
		return fmt.Errorf("Error getting dependency tree")
	}
	depItem, err := depTree.AddRootItem(p.id, p.name, p)
	if err != nil {
		p.ctx.LogError("Error adding %v to dependency tree, err: %v", p.name, err.Error())
		return err
	}
	p.depTree = depTree
	p.dependencyTreeItem = depItem
	return nil
}

func (p *GolangCILinterProvisioner) Name() string {
	return p.name
}

func (p *GolangCILinterProvisioner) ID() string {
	return p.id
}

func (p *GolangCILinterProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p *GolangCILinterProvisioner) Languages() []string {
	return p.languages
}

func (p *GolangCILinterProvisioner) Run() diagnostics.Diagnostics {
	filePath := fmt.Sprintf("%s/.golangci.yml", p.config.WorkingDirectory)
	p.ctx.LogInfo("Provisioning GolangCI Linter")
	defaultConfig := Configuration{
		OlPrefixStyle: "one",
	}
	// p.cfg.RequestFromUser("IAMGROOT", "Yep Everything is fine")
	tmpl, err := template.New("default").Parse(defaultTemplate)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	// Execute the template
	var output bytes.Buffer
	err = tmpl.Execute(&output, defaultConfig)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	if helper.FileExists(filePath) {
		p.ctx.LogInfo("Found existing %v file, ignoring provisioner", filePath)
		p.diagnostics.AddWarning(fmt.Sprintf("found existing %v file, ignoring provisioner", filePath))
		return *p.diagnostics
	}

	err = helper.WriteToFile(output.String(), filePath)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	p.ctx.LogInfo("GolangCI Linter provisioned")
	return *p.diagnostics
}

func (p *GolangCILinterProvisioner) ReshuffleCallback() func() {
	return nil
}
