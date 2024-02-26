package issuer_template_config

import (
	"errors"
	"fmt"
	"os"
	"path"

	_ "embed"

	"github.com/cjlapao/common-go-dependency-tree/dependencytree"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	pkg_constants "github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/helpers"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/common"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/constants"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/templates"
)

var potentialFileNames = []string{
	"config.yml",
	"config.yaml",
}

type IssueTemplateConfigProcessor struct {
	id                 string
	name               string
	diagnostics        *diagnostics.Diagnostics
	cfg                *config.Config
	ctx                *context.ProvisionerContext
	config             interfaces.ProvisionerConfig
	configTemplate     *templates.IssueTemplateConfig
	depTree            *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *IssueTemplateConfigProcessor {
	result := &IssueTemplateConfigProcessor{
		id:     pkg_constants.GitHubIssueTemplateConfigProcessorID,
		name:   "Issue Template Config Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	result.ctx.WithValue(pkg_constants.ContextResourceName, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics

	return result
}

func (p *IssueTemplateConfigProcessor) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&IssueTemplateConfigProcessor{})
	if depTree == nil {
		errorMsg := "Error getting dependency tree"
		p.ctx.LogError(errorMsg)
		return errors.New(errorMsg)
	}
	treeItem, err := depTree.AddItem(p.id, p.name, pkg_constants.GitHubProvisionerID, p)
	if err != nil {
		errorMsg := fmt.Sprintf("Error adding %v to dependency tree, err: %v", p.name, err.Error())
		p.ctx.LogError(errorMsg)
		return errors.New(errorMsg)
	}

	p.dependencyTreeItem = treeItem
	p.depTree = depTree

	return nil
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
	configContactLink := templates.NewIssueTemplateContactLink()
	configContactLink.Name = name
	configContactLink.URL = url
	configContactLink.About = about

	p.configTemplate.ContactLinks = append(p.configTemplate.ContactLinks, configContactLink)
	return *p.diagnostics
}

func (p *IssueTemplateConfigProcessor) Run() diagnostics.Diagnostics {
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
		templates.DefaultIssueTemplateConfig,
		p.configTemplate); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *IssueTemplateConfigProcessor) generateDefault() *templates.IssueTemplateConfig {
	defaultConfig := templates.NewIssueTemplateConfig()
	defaultConfig.BlankIssuesEnabled = false

	defaultConfig.ContactLinks = append(defaultConfig.ContactLinks, templates.IssueTemplateContactLink{
		Name:  "GitHub",
		URL:   "https://example.com",
		About: "dsdsds",
	})
	return &defaultConfig
}

func (p IssueTemplateConfigProcessor) ReshuffleCallback() func() {
	return nil
}
