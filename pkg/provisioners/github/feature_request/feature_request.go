package feature_request

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
	"github.com/cjlapao/github-templater/pkg/provisioners/github/form"
	github_template "github.com/cjlapao/github-templater/pkg/provisioners/github/template"
)

var potentialFileNames = []string{
	"feature_request.yaml",
	"feature_request.yml",
	"feature_request.md",
}

type FeatureRequestProcessor struct {
	id                   string
	name                 string
	diagnostics          *diagnostics.Diagnostics
	cfg                  *config.Config
	ctx                  *context.ProvisionerContext
	config               interfaces.ProvisionerConfig
	featureRequestConfig *github_template.IssueFormTemplate
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *FeatureRequestProcessor {
	result := &FeatureRequestProcessor{
		id:     "github_feature_request_processor",
		name:   "Feature Request Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p *FeatureRequestProcessor) SetConfig(config *github_template.IssueFormTemplate) {
	p.featureRequestConfig = config
}

func (p *FeatureRequestProcessor) Name() string {
	return p.name
}

func (p *FeatureRequestProcessor) ID() string {
	return p.id
}

func (p *FeatureRequestProcessor) Process() diagnostics.Diagnostics {
	issueTemplateFolderPath := path.Join(p.config.WorkingDirectory, constants.GithubFolder, constants.IssueTemplateFolder)
	if !helper.FileExists(issueTemplateFolderPath) {
		err := os.Mkdir(issueTemplateFolderPath, 0o755)
		if err != nil {
			p.diagnostics.AddError(err)
			return *p.diagnostics
		}
	}

	if err := p.checkIfFileExists(issueTemplateFolderPath); err != nil {
		return *p.diagnostics
	}

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

	filePath := path.Join(issueTemplateFolderPath, "feature_request.yml")
	tmpl, err := template.New("default").Funcs(helpers.FuncMap).Parse(github_template.DefaultFormTemplate)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	// Execute the template
	var output bytes.Buffer
	err = tmpl.Execute(&output, p.featureRequestConfig)
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

func (p *FeatureRequestProcessor) checkIfFileExists(folder string) error {
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
		override := p.cfg.RequestBoolFromUser(constants.FeatureRequestFileExistsEnvVar, fmt.Sprintf("A feature request file \"%v\" already exists, do you want to override it? [y/n]", foundFile), false)
		if !override {
			msg := fmt.Sprintf("A feature request file \"%v\" already exists, ignoring provisioner on request by user", foundFile)
			p.ctx.LogWarn(msg)
			p.diagnostics.AddWarning(msg)
			return errors.New("feature request file already exists")
		}
		err := os.Remove(path.Join(folder, "feature_request.yml"))
		if err != nil {
			p.diagnostics.AddError(err)
			return err
		}
	}

	return nil
}

func (p *FeatureRequestProcessor) generateDefault() *github_template.IssueFormTemplate {
	defaultConfig := github_template.NewIssueTemplate("Feature Request")
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
