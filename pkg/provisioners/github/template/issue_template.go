package template

import (
	_ "embed"
	"errors"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/provisioners/github/form"
)

//go:embed issue_template.yml
var DefaultTemplate string

type IssueTemplateConfig struct {
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	Labels      []string        `json:"labels" yaml:"labels"`
	Assignees   []string        `json:"assignees" yaml:"assignees"`
	Projects    []string        `json:"projects" yaml:"projects"`
	Body        form.GithubForm `json:"body" yaml:"body"`
}

func NewIssueTemplateConfig(name string) *IssueTemplateConfig {
	return &IssueTemplateConfig{
		Name:        name,
		Description: "",
		Labels:      []string{},
		Assignees:   []string{},
		Projects:    []string{},
		Body:        *form.NewGithubForm(),
	}
}

func (i IssueTemplateConfig) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if i.Name == "" {
		diag.AddError(errors.New("name is required"))
	}

	if i.Description == "" {
		diag.AddError(errors.New("description is required"))
	}

	diag.Append(i.Body.Validate())

	return diag
}
