package golangci_linter

import (
	"bytes"
	"fmt"
	"text/template"

	_ "embed"

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
	id          string
	name        string
	languages   []string
	ctx         *context.ProvisionerContext
	cfg         *config.Config
	diagnostics *diagnostics.Diagnostics
	config      interfaces.ProvisionerConfig
}

func New(ctx context.ProvisionerContext, providerConfig interfaces.ProvisionerConfig) *GolangCILinterProvisioner {
	result := &GolangCILinterProvisioner{
		id:        "a762b8c7-9e43-45f8-aa55-04b9fc42e1b2",
		name:      "golangci_linter",
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

func (p GolangCILinterProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p GolangCILinterProvisioner) Name() string {
	return p.name
}

func (p GolangCILinterProvisioner) ID() string {
	return p.id
}

func (p GolangCILinterProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p GolangCILinterProvisioner) Languages() []string {
	return p.languages
}

func (p GolangCILinterProvisioner) Provision() diagnostics.Diagnostics {
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
