package form

type TextAreaItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	ID         string                 `json:"id" yaml:"id"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Validators map[string]interface{} `json:"validators,omitempty" yaml:"validators,omitempty"`
}

func NewTextAreaItem(id string) TextAreaItem {
	return TextAreaItem{
		ItemType:   GithubFormTypeTextArea,
		ID:         id,
		Attributes: map[string]interface{}{},
		Validators: map[string]interface{}{},
	}
}

func (i TextAreaItem) Type() GithubFormType {
	return i.ItemType
}

func (i TextAreaItem) Label(v string) {
	i.Attributes["label"] = v
}

func (i TextAreaItem) Description(v string) {
	i.Attributes["description"] = v
}

func (i TextAreaItem) Placeholder(v string) {
	i.Attributes["placeholder"] = v
}

func (i TextAreaItem) Value(v string) {
	i.Attributes["value"] = v
}

func (i TextAreaItem) IsRequired(v bool) {
}

func (i TextAreaItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":       i.ItemType,
		"attributes": i.Attributes,
	}
}
