package golangcilinter

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	_ "embed"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/constants"
	"github.com/cjlapao/github-templater/context"
	"github.com/cjlapao/github-templater/interfaces"
)

//go:embed default.yml
var defaultTemplate string

type GolangCILinterProvisioner struct {
	id        string
	name      string
	languages []string
	test      string
	ctx       *context.ProvisionerContext
	config    interfaces.ProvisionerConfig
}

func New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) *GolangCILinterProvisioner {
	result := &GolangCILinterProvisioner{
		id:        "a762b8c7-9e43-45f8-aa55-04b9fc42e1b2",
		name:      "golangci_linter",
		ctx:       &ctx,
		languages: []string{"go", "golang"},
		config:    config,
	}

	result.ctx.WithValue(constants.ContextResourceName, result.name)
	result.test = fmt.Sprintf("Test date %v", time.Now().Format("2006-01-02 15:04:05.000000000 -0700 MST"))
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

func (p GolangCILinterProvisioner) Context() context.ProvisionerContext {
	return *p.ctx
}

func (p GolangCILinterProvisioner) Languages() []string {
	return p.languages
}

func (p GolangCILinterProvisioner) Provision() error {
	filePath := fmt.Sprintf("%s/.golangci.yml", p.config.WorkingDirectory)
	p.ctx.LogInfo("Provisioning GolangCI Linter")
	defaultConfig := Configuration{
		OlPrefixStyle: "one",
	}
	tmpl, err := template.New("default").Parse(defaultTemplate)
	if err != nil {
		return err
	}

	// Execute the template
	var output bytes.Buffer
	err = tmpl.Execute(&output, defaultConfig)
	if err != nil {
		return err
	}

	if helper.FileExists(filePath) {
		p.ctx.LogInfo("Found existing %v file, ignoring provisioner", filePath)
		return nil
	}

	err = helper.WriteToFile(output.String(), filePath)
	if err != nil {
		return err
	}

	p.ctx.LogInfo("GolangCI Linter provisioned")
	return nil
}
