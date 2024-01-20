package bug_report

import "github.com/cjlapao/github-templater/pkg/provisioners/github/form"

type BugReportConfig struct {
	Name        string          `json:"name" yaml:"name"`
	Description string          `json:"description" yaml:"description"`
	Labels      []string        `json:"labels" yaml:"labels"`
	Assignees   []string        `json:"assignees" yaml:"assignees"`
	Body        form.GithubForm `json:"body" yaml:"body"`
}

func NewBugReportConfig(name string) *BugReportConfig {
	return &BugReportConfig{
		Name:        name,
		Description: "",
		Labels:      []string{},
		Assignees:   []string{},
		Body:        *form.NewGithubForm(),
	}
}
