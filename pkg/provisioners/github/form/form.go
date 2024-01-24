package form

import "github.com/cjlapao/github-templater/pkg/diagnostics"

type GithubFormType string

const (
	GithubFormTypeTextArea GithubFormType = "textarea"
	GithubFormTypeInput    GithubFormType = "input"
	GithubFormTypeDropdown GithubFormType = "dropdown"
	GithubFormTypeCheckbox GithubFormType = "checkboxes"
	GithubFormTypeMarkdown GithubFormType = "markdown"
)

type GithubForm struct {
	Items []GithubFormItem
}

func NewGithubForm() *GithubForm {
	return &GithubForm{
		Items: []GithubFormItem{},
	}
}

func (f *GithubForm) AddItem(item GithubFormItem) {
	f.Items = append(f.Items, item)
}

func (f *GithubForm) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	for _, item := range f.Items {
		diag.Append(item.Validate())
	}

	return diag
}

type GithubFormItem interface {
	Type() GithubFormType
	IsRequired(v bool)
	Validate() diagnostics.Diagnostics
}
