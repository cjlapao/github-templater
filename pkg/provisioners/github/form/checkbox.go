package form

import "strings"

type CheckboxItemOption struct {
	Label    string `json:"label,omitempty" yaml:"label,omitempty"`
	Required bool   `json:"required,omitempty" yaml:"required,omitempty"`
}

type CheckboxItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	ID         string                 `json:"id" yaml:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Validators map[string]interface{} `json:"validations,omitempty" yaml:"validations,omitempty"`
}

func NewCheckboxItem(id string) CheckboxItem {
	return CheckboxItem{
		ItemType:   GithubFormTypeCheckbox,
		ID:         id,
		Attributes: map[string]interface{}{},
		Validators: map[string]interface{}{},
	}
}

func (i CheckboxItem) Type() GithubFormType {
	return i.ItemType
}

func (i CheckboxItem) Label(v string) {
	i.Attributes["label"] = v
}

func (i CheckboxItem) Description(v string) {
	i.Attributes["description"] = v
}

func (i CheckboxItem) AddOption(v string, checked bool) {
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

func (i CheckboxItem) IsRequired(v bool) {
}

func (i CheckboxItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":       i.ItemType,
		"attributes": i.Attributes,
	}
}
