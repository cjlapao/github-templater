package dependabot

import (
	"errors"
	"os"
	"path"

	"github.com/cjlapao/common-go-dependency-tree/dependencytree"
	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/helpers"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/common"
	github_constants "github.com/cjlapao/github-templater/pkg/provisioners/github/constants"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/form"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/templates"
)

var potentialFileNames = []string{
	"dependabot.yml",
	"dependabot.yaml",
}

type DependabotProvisioner struct {
	id                 string
	name               string
	diagnostics        *diagnostics.Diagnostics
	cfg                *config.Config
	ctx                *context.ProvisionerContext
	config             interfaces.ProvisionerConfig
	dependabotConfig   *templates.IssueFormTemplate
	languages          []string
	depTree            *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *DependabotProvisioner {
	result := &DependabotProvisioner{
		id:        constants.GitHubDependabotProvisionerID,
		name:      "GitHub Dependabot Processor",
		cfg:       config.Get(),
		ctx:       &ctx,
		config:    provisionerConfig,
		languages: []string{"go", "golang", "docker"},
	}

	result.ctx.WithValue(constants.GitHubProvisionerID, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p *DependabotProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p *DependabotProvisioner) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&DependabotProvisioner{})
	if depTree == nil {
		p.ctx.LogError("Error getting dependency tree")
		return errors.New("Error getting dependency tree")
	}
	depItem, err := depTree.AddItem(p.id, p.name, constants.GoLangCILinterProvisionerID, p)
	if err != nil {
		p.ctx.LogError("Error adding %v to dependency tree, err: %v", p.name, err.Error())
		return nil
	}
	p.depTree = depTree
	p.dependencyTreeItem = depItem

	return nil
}

func (p *DependabotProvisioner) Name() string {
	return p.name
}

func (p *DependabotProvisioner) ID() string {
	return p.id
}

func (p *DependabotProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p *DependabotProvisioner) Languages() []string {
	return p.languages
}

func (p *DependabotProvisioner) Run() diagnostics.Diagnostics {
	// Ask the user if they want to generate a dependabot config file
	if !p.cfg.RequestBoolFromUser(DependabotGenerateEnvVar, "Do you want to generate a dependabot config file?", false) {
		return *p.diagnostics
	}

	githubFolder := path.Join(p.config.WorkingDirectory, github_constants.GithubFolder)
	if !helper.FileExists(githubFolder) {
		err := os.Mkdir(githubFolder, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
		}
	}

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Dependabot", githubFolder, potentialFileNames, DependabotFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

	if p.dependabotConfig == nil {
		p.dependabotConfig = p.generateDefault()
	}

	if p.dependabotConfig == nil {
		p.diagnostics.AddError(errors.New("dependabot config is nil"))
		return *p.diagnostics
	}

	validateDiag := p.dependabotConfig.Validate()
	if validateDiag.HasErrors() {
		p.diagnostics.Append(validateDiag)
		return *p.diagnostics
	}

	filePath := path.Join(githubFolder, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		templates.DefaultFormTemplate,
		p.dependabotConfig); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *DependabotProvisioner) generateDefault() *templates.IssueFormTemplate {
	defaultConfig := templates.NewIssueTemplate("Bug Report")
	defaultConfig.Labels = append(defaultConfig.Labels, "bug")
	defaultConfig.Description = "Create a report to help us improve"
	bugDescriptionItem := form.NewTextAreaItem("bug_description")
	bugDescriptionItem.Label("Bug description")
	bugDescriptionItem.Render("Markdown")
	bugDescriptionItem.Value("**Please describe the bug**")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, bugDescriptionItem)
	reproduceStepsItem := form.NewTextAreaItem("reproduce_steps")
	reproduceStepsItem.Label("Steps to reproduce")
	reproduceStepsItem.Value("Steps to reproduce the behavior:\n1. Go to '...'\n2. Click on '....'\n3. Scroll down to '....'\n4. See error")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, reproduceStepsItem)
	versionItem := form.NewInputItem("version")
	versionItem.Label("Version")
	versionItem.Placeholder("e.g. v1.0.0")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, versionItem)
	expectedBehaviourItem := form.NewTextAreaItem("expected_behaviour")
	expectedBehaviourItem.Label("Expected behavior")
	expectedBehaviourItem.Value("A clear and concise description of what you expected to happen.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, expectedBehaviourItem)
	actualBehaviourItem := form.NewTextAreaItem("actual_behaviour")
	actualBehaviourItem.Label("Actual behavior")
	actualBehaviourItem.Value("A clear and concise description of what actually happened.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, actualBehaviourItem)
	screenshotsItem := form.NewTextAreaItem("screenshots")
	screenshotsItem.Label("Screenshots")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, screenshotsItem)
	additionalContextItem := form.NewTextAreaItem("additional_context")
	additionalContextItem.Label("Additional context")
	additionalContextItem.Value("Add any other context about the problem here.\n For example, if you are using a specific version of the language, or if you are using a specific operating system.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, additionalContextItem)

	validationItem := form.NewCheckboxItems("validation")
	validationItem.Label("Validation")
	validationItem.AddOption("Yes, I've included all of the above information (Version, settings, logs, etc.)", true)
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, validationItem)

	return defaultConfig
}

func (p *DependabotProvisioner) ReshuffleCallback() func() {
	return nil
}
