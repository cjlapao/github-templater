package form

import (
	"fmt"
	"strings"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type CheckboxItemOption struct {
	Label    string `json:"label,omitempty" yaml:"label,omitempty"`
	Required bool   `json:"required,omitempty" yaml:"required,omitempty"`
}

type CheckboxesItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	ID         string                 `json:"id" yaml:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Validators map[string]interface{} `json:"validations,omitempty" yaml:"validations,omitempty"`
}

func NewCheckboxItems(id string) CheckboxesItem {
	return CheckboxesItem{
		ItemType:   GithubFormTypeCheckbox,
		ID:         id,
		Attributes: map[string]interface{}{},
		Validators: map[string]interface{}{},
	}
}

func (i CheckboxesItem) Type() GithubFormType {
	return i.ItemType
}

func (i CheckboxesItem) Label(v string) {
	i.Attributes["label"] = v
}

func (i CheckboxesItem) Description(v string) {
	i.Attributes["description"] = v
}

func (i CheckboxesItem) AddOption(v string, checked bool) {
	item := CheckboxItemOption{
		Label:    v,
		Required: checked,
	}
	if ok := i.Attributes["options"]; ok != nil {
		if _, ok := i.Attributes["options"].([]CheckboxItemOption); !ok {
			i.Attributes["options"] = []CheckboxItemOption{}
		}
		options := i.Attributes["options"].([]CheckboxItemOption)
		for _, option := range options {
			if strings.EqualFold(option.Label, v) {
				return
			}
		}
		i.Attributes["options"] = append(i.Attributes["options"].([]CheckboxItemOption), item)
	} else {
		i.Attributes["options"] = []CheckboxItemOption{item}
	}
}

func (i CheckboxesItem) IsRequired(v bool) {
}

func (i CheckboxesItem) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if label, ok := i.Attributes["label"].(string); ok {
		if label == "" {
			diag.AddError(fmt.Errorf("checkboxes %v label is required", i.ID))
		}
	} else {
		if i.Attributes["label"] == nil {
			diag.AddError(fmt.Errorf("checkboxes %v label is required", i.ID))
		} else {
			diag.AddError(fmt.Errorf("checkboxes %v label is not a string", i.ID))
		}
	}

	options := i.Attributes["options"].([]CheckboxItemOption)
	if len(options) == 0 {
		diag.AddError(fmt.Errorf("checkboxes %v options are required", i.ID))
	}
	for _, option := range options {
		if option.Label == "" {
			diag.AddError(fmt.Errorf("checkboxes %v option label is required", i.ID))
		}
	}

	return diag
}
