package github

import (
	"os"
	"path"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/bug_report"
	github_constants "github.com/cjlapao/github-templater/pkg/provisioners/github/constants"
)

type GitHubProvisioner struct {
	id          string
	name        string
	languages   []string
	ctx         *context.ProvisionerContext
	cfg         *config.Config
	diagnostics *diagnostics.Diagnostics
	config      interfaces.ProvisionerConfig
}

func New(ctx context.ProvisionerContext, providerConfig interfaces.ProvisionerConfig) *GitHubProvisioner {
	result := &GitHubProvisioner{
		id:        "1a3a0ae0-bb6b-4e69-8d14-6aa3e150ec3e",
		name:      "github",
		ctx:       &ctx,
		cfg:       config.Get(),
		languages: []string{"all"},
		config:    providerConfig,
	}

	result.ctx.WithValue(constants.ContextResourceName, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics
	return result
}

func (p GitHubProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p GitHubProvisioner) Name() string {
	return p.name
}

func (p GitHubProvisioner) ID() string {
	return p.id
}

func (p GitHubProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p GitHubProvisioner) Languages() []string {
	return p.languages
}

func (p GitHubProvisioner) Provision() diagnostics.Diagnostics {
	p.ctx.LogInfo("Provisioning GitHub")
	p.tryGenerateGithubFolder()
	bugReport := bug_report.New(p.ctx, p.config)
	if bugReportDiag := bugReport.Process(); bugReportDiag.HasErrors() {
		return bugReportDiag
	}

	p.ctx.LogInfo("GolangCI Linter provisioned")
	return *p.diagnostics
}

func (p GitHubProvisioner) tryGenerateGithubFolder() error {
	githubFolderPath := path.Join(p.config.WorkingDirectory, github_constants.GithubFolder)
	if !helper.FileExists(githubFolderPath) {
		err := os.Mkdir(githubFolderPath, 0o755)
		if err != nil {
			return err
		}
	}
	return nil
}
