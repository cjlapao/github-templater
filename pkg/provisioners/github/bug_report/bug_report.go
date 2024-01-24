package bug_report

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
	"bug_report.yaml",
	"bug_report.yml",
	"bug_report.md",
}

type BugReportProcessor struct {
	id              string
	name            string
	diagnostics     *diagnostics.Diagnostics
	cfg             *config.Config
	ctx             *context.ProvisionerContext
	config          interfaces.ProvisionerConfig
	bugReportConfig *github_template.IssueFormTemplateConfig
}

func New(ctx *context.ProvisionerContext, provisionerConfig interfaces.ProvisionerConfig) *BugReportProcessor {
	result := &BugReportProcessor{
		id:     "github_bug_report_processor",
		name:   "Bug Report Processor",
		cfg:    config.Get(),
		ctx:    ctx,
		config: provisionerConfig,
	}

	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p *BugReportProcessor) Name() string {
	return p.name
}

func (p *BugReportProcessor) ID() string {
	return p.id
}

func (p *BugReportProcessor) SetConfig(config *github_template.IssueFormTemplateConfig) {
	p.bugReportConfig = config
}

func (p *BugReportProcessor) Process() diagnostics.Diagnostics {
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

	if p.bugReportConfig == nil {
		p.bugReportConfig = p.generateDefault()
	}

	if p.bugReportConfig == nil {
		p.diagnostics.AddError(errors.New("bug report config is nil"))
		return *p.diagnostics
	}

	validateDiag := p.bugReportConfig.Validate()
	if validateDiag.HasErrors() {
		p.diagnostics.Append(validateDiag)
		return *p.diagnostics
	}

	filePath := path.Join(issueTemplateFolderPath, "bug_report.yml")
	tmpl, err := template.New("default").Funcs(helpers.FuncMap).Parse(github_template.DefaultFormTemplate)
	if err != nil {
		p.diagnostics.AddError(err)
		return *p.diagnostics
	}

	// Execute the template
	var output bytes.Buffer
	err = tmpl.Execute(&output, p.bugReportConfig)
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

func (p *BugReportProcessor) checkIfFileExists(folder string) error {
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
		override := p.cfg.RequestBoolFromUser("GITHUB_BUG_REPORT_FILE_EXIST", fmt.Sprintf("A bug report file \"%v\" already exists, do you want to override it? [y/n]", foundFile), false)
		if !override {
			msg := fmt.Sprintf("A bug report file \"%v\" already exists, ignoring provisioner on request by user", foundFile)
			p.ctx.LogWarn(msg)
			p.diagnostics.AddWarning(msg)
			return errors.New("bug report file already exists")
		}
		err := os.Remove(path.Join(folder, "bug_report.yml"))
		if err != nil {
			p.diagnostics.AddError(err)
			return err
		}
	}

	return nil
}

func (p *BugReportProcessor) generateDefault() *github_template.IssueFormTemplateConfig {
	defaultConfig := github_template.NewIssueTemplateConfig("Bug Report")
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
