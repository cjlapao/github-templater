package form

import (
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type InputItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	ID         string                 `json:"id" yaml:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Validators map[string]interface{} `json:"validations,omitempty" yaml:"validations,omitempty"`
}

func NewInputItem(id string) InputItem {
	return InputItem{
		ItemType:   GithubFormTypeInput,
		ID:         id,
		Attributes: map[string]interface{}{},
		Validators: map[string]interface{}{},
	}
}

func (i InputItem) Type() GithubFormType {
	return i.ItemType
}

func (i InputItem) Label(v string) {
	i.Attributes["label"] = v
}

func (i InputItem) Description(v string) {
	i.Attributes["description"] = v
}

func (i InputItem) Placeholder(v string) {
	i.Attributes["placeholder"] = v
}

func (i InputItem) Value(v string) {
	i.Attributes["value"] = v
}

func (i InputItem) IsRequired(v bool) {
	i.Validators["required"] = v
}

func (i InputItem) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if label, ok := i.Attributes["label"].(string); ok {
		if label == "" {
			diag.AddError(fmt.Errorf("input %v label is required", i.ID))
		}
	} else {
		if i.Attributes["label"] == nil {
			diag.AddError(fmt.Errorf("input %v label is required", i.ID))
		} else {
			diag.AddError(fmt.Errorf("input %v label is not a string", i.ID))
		}
	}

	return diag
}
