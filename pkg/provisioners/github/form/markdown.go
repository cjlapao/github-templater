package form

type MarkdownItem struct {
	ItemType   GithubFormType         `json:"type" yaml:"type"`
	Attributes map[string]interface{} `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

func NewMarkdownItem() MarkdownItem {
	return MarkdownItem{
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

func (i MarkdownItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":       i.ItemType,
		"attributes": i.Attributes,
	}
}
