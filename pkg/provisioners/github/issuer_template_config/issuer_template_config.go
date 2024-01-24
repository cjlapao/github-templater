package issuer_template_config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"text/template"

	_ "embed"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/helpers"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/constants"
	github_template "github.com/cjlapao/github-templater/pkg/provisioners/github/template"
)

var potentialFileNames = []string{
	"config.yaml",
	"config.yml",
}

type IssueTemplateConfigProcessor struct {
	id             string
	name           string
	diagnostics    *diagnostics.Diagnostics
	cfg            *config.Config
	ctx            *context.ProvisionerContext
	config         interfaces.ProvisionerConfig
	configTemplate *github_template.IssueTemplateConfig
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
	configContactLink := github_template.NewIssueTemplateContactLink()
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

	if err := p.checkIfFileExists(issueTemplateFolderPath); err != nil {
		return *p.diagnostics
	}

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

	filePath := path.Join(issueTemplateFolderPath, "config.yml")
	tmpl, err := template.New("default").Funcs(helpers.FuncMap).Parse(github_template.DefaultIssueTemplateConfig)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	// Execute the template
	var output bytes.Buffer
	err = tmpl.Execute(&output, p.configTemplate)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	err = helper.WriteToFile(output.String(), filePath)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *IssueTemplateConfigProcessor) checkIfFileExists(folder string) error {
	fileExists := false
	foundFile := ""
	for _, fileName := range potentialFileNames {
		if helper.FileExists(path.Join(folder, fileName)) {
			fileExists = true
			foundFile = fileName
			break
		}
	}

	if fileExists {
		override := p.cfg.RequestBoolFromUser(constants.IssueTemplateConfigFileExistsEnvVar, fmt.Sprintf("A bug report file \"%v\" already exists, do you want to override it? [y/n]", foundFile), false)
		if !override {
			msg := fmt.Sprintf("A issue template config file \"%v\" already exists, ignoring provisioner on request by user", foundFile)
			p.ctx.LogWarn(msg)
			p.diagnostics.AddWarning(msg)
			return errors.New("issue template config file already exists")
		}
		err := os.Remove(path.Join(folder, "config.yml"))
		if err != nil {
			p.diagnostics.AddError(err)
			return err
		}
	}

	return nil
}

func (p *IssueTemplateConfigProcessor) generateDefault() *github_template.IssueTemplateConfig {
	defaultConfig := github_template.NewIssueTemplateConfig()
	defaultConfig.BlankIssuesEnabled = false

	defaultConfig.ContactLinks = append(defaultConfig.ContactLinks, github_template.IssueTemplateContactLink{
		Name:  "GitHub",
		URL:   "https://example.com",
		About: "dsdsds",
	})
	return &defaultConfig
}
