package form

import "strings"

type DropdownItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	ID         string                 `json:"id" yaml:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Validators map[string]interface{} `json:"validations,omitempty" yaml:"validations,omitempty"`
}

func NewDropdownItem(id string) DropdownItem {
	return DropdownItem{
		ItemType:   GithubFormTypeDropdown,
		ID:         id,
		Attributes: map[string]interface{}{},
		Validators: map[string]interface{}{},
	}
}

func (i DropdownItem) Type() GithubFormType {
	return i.ItemType
}

func (i DropdownItem) Label(v string) {
	i.Attributes["label"] = v
}

func (i DropdownItem) Description(v string) {
	i.Attributes["description"] = v
}

func (i DropdownItem) Multiple(v bool) {
	i.Attributes["multiple"] = v
}

func (i DropdownItem) AddOption(v string) {
	if ok := i.Attributes["options"]; ok != nil {
		if _, ok := i.Attributes["options"].([]string); !ok {
			i.Attributes["options"] = []string{}
		}
		options := i.Attributes["options"].([]string)
		for _, option := range options {
			if strings.EqualFold(option, v) {
				return
			}
		}
		i.Attributes["options"] = append(i.Attributes["options"].([]string), v)
	} else {
		i.Attributes["options"] = []string{v}
	}
}

func (i DropdownItem) Default(v int) {
	if ok := i.Attributes["options"]; ok == nil {
		return
	}
	options := i.Attributes["options"].([]string)
	if v < 0 || v >= len(options) {
		return
	}

	i.Attributes["default"] = v
}

func (i DropdownItem) IsRequired(v bool) {
	i.Validators["required"] = v
}

func (i DropdownItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":       i.ItemType,
		"attributes": i.Attributes,
	}
}
