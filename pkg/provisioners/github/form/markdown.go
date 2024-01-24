package form

import (
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type MarkdownItem struct {
	ID         string                 `json:"-" yaml:"-"`
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func NewMarkdownItem(id string) MarkdownItem {
	return MarkdownItem{
		ID:         id,
		ItemType:   GithubFormTypeMarkdown,
		Attributes: map[string]interface{}{},
	}
}

func (i MarkdownItem) Type() GithubFormType {
	return i.ItemType
}

func (i MarkdownItem) Value(v string) {
	i.Attributes["value"] = v
}

func (i MarkdownItem) IsRequired(v bool) {
}

func (i MarkdownItem) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if value, ok := i.Attributes["value"].(string); ok {
		if value == "" {
			diag.AddError(fmt.Errorf("markdown %v value is required", i.ID))
		}
	} else {
		if i.Attributes["label"] == nil {
			diag.AddError(fmt.Errorf("markdown %v label is required", i.ID))
		} else {
			diag.AddError(fmt.Errorf("markdown %v label is not a string", i.ID))
		}
	}

	return diag
}
