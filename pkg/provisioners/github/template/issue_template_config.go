package template

import (
	_ "embed"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

//go:embed issue_template_config.yml.tpl
var DefaultIssueTemplateConfig string

type IssueTemplateConfig struct {
	BlankIssuesEnabled bool                       `json:"blank_issues_enabled" yaml:"blank_issues_enabled"`
	ContactLinks       []IssueTemplateContactLink `json:"contact_links" yaml:"contact_links"`
}

func NewIssueTemplateConfig() IssueTemplateConfig {
	return IssueTemplateConfig{
		BlankIssuesEnabled: false,
		ContactLinks:       []IssueTemplateContactLink{},
	}
}

func (i IssueTemplateConfig) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()

	for _, contactLink := range i.ContactLinks {
		diag.Append(contactLink.Validate())
	}

	return diag
}
