package issuer_template_config

import (
	"errors"
	"os"
	"path"

	_ "embed"

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
	"config.yml",
	"config.yaml",
}

type IssueTemplateConfigProcessor struct {
	id             string
	name           string
	diagnostics    *diagnostics.Diagnostics
	cfg            *config.Config
	ctx            *context.ProvisionerContext
	config         interfaces.ProvisionerConfig
	configTemplate *template.IssueTemplateConfig
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *IssueTemplateConfigProcessor {
	result := &IssueTemplateConfigProcessor{
		id:     "github_issue_template_config_processor",
		name:   "Issue Template Config Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p *IssueTemplateConfigProcessor) Name() string {
	return p.name
}

func (p *IssueTemplateConfigProcessor) ID() string {
	return p.id
}

func (p *IssueTemplateConfigProcessor) AddContactLink(name string, url string, about string) diagnostics.Diagnostics {
	if p.configTemplate == nil {
		p.configTemplate = p.generateDefault()
	}
	configContactLink := template.NewIssueTemplateContactLink()
	configContactLink.Name = name
	configContactLink.URL = url
	configContactLink.About = about

	p.configTemplate.ContactLinks = append(p.configTemplate.ContactLinks, configContactLink)
	return *p.diagnostics
}

func (p *IssueTemplateConfigProcessor) Process() diagnostics.Diagnostics {
	issueTemplateFolderPath := path.Join(p.config.WorkingDirectory, constants.GithubFolder, constants.IssueTemplateFolder)
	if !helper.FileExists(issueTemplateFolderPath) {
		err := os.Mkdir(issueTemplateFolderPath, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
		}
	}

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Issue template", issueTemplateFolderPath, potentialFileNames, constants.IssueTemplateConfigFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

	if p.configTemplate == nil {
		p.configTemplate = p.generateDefault()
	}

	if p.configTemplate == nil {
		p.diagnostics.AddError(errors.New("issue template config is nil"))
		return *p.diagnostics
	}

	validateDiag := p.configTemplate.Validate()
	if validateDiag.HasErrors() {
		p.diagnostics.Append(validateDiag)
		return *p.diagnostics
	}

	filePath := path.Join(issueTemplateFolderPath, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		template.DefaultIssueTemplateConfig,
		p.configTemplate); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *IssueTemplateConfigProcessor) generateDefault() *template.IssueTemplateConfig {
	defaultConfig := template.NewIssueTemplateConfig()
	defaultConfig.BlankIssuesEnabled = false

	defaultConfig.ContactLinks = append(defaultConfig.ContactLinks, template.IssueTemplateContactLink{
		Name:  "GitHub",
		URL:   "https://example.com",
		About: "dsdsds",
	})
	return &defaultConfig
}
