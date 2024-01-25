package bug_report

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
	"github.com/cjlapao/github-templater/pkg/provisioners/github/form"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/template"
)

var potentialFileNames = []string{
	"bug_report.yml",
	"bug_report.yaml",
	"bug_report.md",
}

type BugReportProcessor struct {
	id              string
	name            string
	diagnostics     *diagnostics.Diagnostics
	cfg             *config.Config
	ctx             *context.ProvisionerContext
	config          interfaces.ProvisionerConfig
	bugReportConfig *template.IssueFormTemplate
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

func (p *BugReportProcessor) SetConfig(config *template.IssueFormTemplate) {
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

	fileExistsDiags := common.CheckIfFileExists(p.ctx, "Bug Report", issueTemplateFolderPath, potentialFileNames, constants.BugReportFileExistsEnvVar)
	p.diagnostics.Append(fileExistsDiags)

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

	filePath := path.Join(issueTemplateFolderPath, potentialFileNames[0])
	if writeDiag := helpers.WriteTemplateToFile(filePath,
		template.DefaultFormTemplate,
		p.bugReportConfig); writeDiag.HasErrors() {
		p.diagnostics.Append(writeDiag)
		return *p.diagnostics
	}

	return *p.diagnostics
}

func (p *BugReportProcessor) generateDefault() *template.IssueFormTemplate {
	defaultConfig := template.NewIssueTemplate("Bug Report")
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
