package form

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

func (i InputItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":       i.ItemType,
		"attributes": i.Attributes,
	}
}
