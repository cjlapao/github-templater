package form

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

type GithubFormItem interface {
	Type() GithubFormType
	IsRequired(v bool)
	ToMap() map[string]interface{}
}
