package stale

import (
	"errors"
	"fmt"
	"os"
	"path"

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
	"stale.yml",
	"stale.yaml",
}

type StaleProcessor struct {
	id                 string
	name               string
	diagnostics        *diagnostics.Diagnostics
	cfg                *config.Config
	ctx                *context.ProvisionerContext
	config             interfaces.ProvisionerConfig
	configTemplate     *templates.StaleConfig
	depTree            *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *StaleProcessor {
	result := &StaleProcessor{
		id:     pkg_constants.GitHubStaleProcessorID,
		name:   "Stale Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	result.ctx.WithValue(pkg_constants.ContextResourceName, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics

	return result
}

func (p *StaleProcessor) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&StaleProcessor{})
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

func (p *StaleProcessor) Name() string {
	return p.name
}

func (p *StaleProcessor) ID() string {
	return p.id
}

func (p *StaleProcessor) Run() diagnostics.Diagnostics {
	if !p.cfg.RequestBoolFromUser(constants.StaleGenerateEnvVar, "Do you want to generate a stale.yml file? [y/n]", false) {
		return *p.diagnostics
	}

	githubFolder := path.Join(p.config.WorkingDirectory, constants.GithubFolder)
	if !helper.FileExists(githubFolder) {
		err := os.Mkdir(githubFolder, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
		}
	}

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Stale", githubFolder, potentialFileNames, constants.StaleFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

	if p.configTemplate == nil {
		config := templates.NewStaleConfig()
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

	filePath := path.Join(githubFolder, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		templates.DefaultStaleTemplate,
		p.configTemplate.Generate()); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p StaleProcessor) ReshuffleCallback() func() {
	return nil
}
