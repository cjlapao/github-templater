package stale

import (
	"errors"
	"os"
	"path"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/helpers"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/common"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/constants"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/template"
)

var potentialFileNames = []string{
	"stale.yml",
	"stale.yaml",
}

type StaleProcessor struct {
	id             string
	name           string
	diagnostics    *diagnostics.Diagnostics
	cfg            *config.Config
	ctx            *context.ProvisionerContext
	config         interfaces.ProvisionerConfig
	configTemplate *template.StaleConfig
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *StaleProcessor {
	result := &StaleProcessor{
		id:     "github_stale_processor",
		name:   "Stale Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p *StaleProcessor) Name() string {
	return p.name
}

func (p *StaleProcessor) ID() string {
	return p.id
}

func (p *StaleProcessor) Process() diagnostics.Diagnostics {
	if !p.cfg.RequestBoolFromUser(constants.StaleGenerateEnvVar, "Do you want to generate a stale.yml file? [y/n]", false) {
		return *p.diagnostics
	}

	staleFolderPath := path.Join(p.config.WorkingDirectory, constants.GithubFolder)
	if !helper.FileExists(staleFolderPath) {
		err := os.Mkdir(staleFolderPath, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
		}
	}

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Stale", staleFolderPath, potentialFileNames, constants.StaleFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

	if p.configTemplate == nil {
		config := template.NewStaleConfig()
		p.configTemplate = &config
	}

	if p.configTemplate == nil {
		p.diagnostics.AddError(errors.New("stale config is nil"))
		return *p.diagnostics
	}

	validateDiag := p.configTemplate.Validate()
	if validateDiag.HasErrors() {
		p.diagnostics.Append(validateDiag)
		return *p.diagnostics
	}

	filePath := path.Join(staleFolderPath, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		template.DefaultStaleTemplate,
		p.configTemplate.Generate()); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}
