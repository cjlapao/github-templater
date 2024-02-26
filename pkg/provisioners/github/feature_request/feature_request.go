package feature_request

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
	"github.com/cjlapao/github-templater/pkg/provisioners/github/form"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/templates"
)

var potentialFileNames = []string{
	"feature_request.yml",
	"feature_request.yaml",
	"feature_request.md",
}

type FeatureRequestProcessor struct {
	id                   string
	name                 string
	diagnostics          *diagnostics.Diagnostics
	cfg                  *config.Config
	ctx                  *context.ProvisionerContext
	config               interfaces.ProvisionerConfig
	featureRequestConfig *templates.IssueFormTemplate
	depTree              *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem   *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *FeatureRequestProcessor {
	result := &FeatureRequestProcessor{
		id:     pkg_constants.GitHubFeatureRequestProcessorID,
		name:   "Feature Request Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	result.ctx.WithValue(pkg_constants.ContextResourceName, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics

	return result
}

func (p *FeatureRequestProcessor) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&FeatureRequestProcessor{})
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

	if err := treeItem.DependsOn(pkg_constants.GitHubIssueTemplateConfigProcessorID); err != nil {
		errorMsg := fmt.Sprintf("Error adding dependency to item %v in dependency tree, err: %v", p.name, err.Error())
		p.ctx.LogError(errorMsg)
		return errors.New(errorMsg)
	}

	p.dependencyTreeItem = treeItem
	p.depTree = depTree

	return nil
}

func (p *FeatureRequestProcessor) SetConfig(config *templates.IssueFormTemplate) {
	p.featureRequestConfig = config
}

func (p *FeatureRequestProcessor) Name() string {
	return p.name
}

func (p *FeatureRequestProcessor) ID() string {
	return p.id
}

func (p *FeatureRequestProcessor) Run() diagnostics.Diagnostics {
	issueTemplateFolderPath := path.Join(p.config.WorkingDirectory, constants.GithubFolder, constants.IssueTemplateFolder)
	if !helper.FileExists(issueTemplateFolderPath) {
		err := os.Mkdir(issueTemplateFolderPath, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
			return *p.diagnostics
		}
	}

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Feature Request", issueTemplateFolderPath, potentialFileNames, constants.FeatureRequestFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

	if p.featureRequestConfig == nil {
		p.featureRequestConfig = p.generateDefault()
	}

	if p.featureRequestConfig == nil {
		p.diagnostics.AddError(errors.New("feature request config is nil"))
		return *p.diagnostics
	}

	validateDiag := p.featureRequestConfig.Validate()
	if validateDiag.HasErrors() {
		p.diagnostics.Append(validateDiag)
		return *p.diagnostics
	}

	filePath := path.Join(issueTemplateFolderPath, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		templates.DefaultFormTemplate,
		p.featureRequestConfig); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *FeatureRequestProcessor) generateDefault() *templates.IssueFormTemplate {
	defaultConfig := templates.NewIssueTemplate("Feature Request")
	defaultConfig.Labels = append(defaultConfig.Labels, "triage", "feature request")
	defaultConfig.Description = "Suggest an idea for this project"
	featureDescriptionItem := form.NewTextAreaItem("feature_description")
	featureDescriptionItem.Label("Is your feature request related to a problem? Please describe.")
	featureDescriptionItem.IsRequired(true)
	featureDescriptionItem.Render("Markdown")
	featureDescriptionItem.Value("A clear and concise description of what the problem is. Ex. I'm always frustrated\nwhen [...]")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, featureDescriptionItem)
	describeFeatureItem := form.NewTextAreaItem("describe_feature")
	describeFeatureItem.Label("Describe the solution you'd like")
	describeFeatureItem.IsRequired(true)
	describeFeatureItem.Value("A clear and concise description of what you want to happen.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, describeFeatureItem)
	alternativesItem := form.NewTextAreaItem("describe_alternatives")
	alternativesItem.Label("Describe alternatives you've considered")
	alternativesItem.IsRequired(true)
	alternativesItem.Value("A clear and concise description of any alternative solutions or features you've considered.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, alternativesItem)
	additionalContextItem := form.NewTextAreaItem("additional_context")
	additionalContextItem.Label("Additional context")
	additionalContextItem.Value("Add any other context or screenshots about the feature request here.")
	defaultConfig.Body.Items = append(defaultConfig.Body.Items, additionalContextItem)

	return defaultConfig
}

func (p FeatureRequestProcessor) ReshuffleCallback() func() {
	return nil
}
