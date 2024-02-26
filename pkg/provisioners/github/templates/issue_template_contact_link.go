package templates

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type IssueTemplateContactLink struct {
	Name  string `json:"name" yaml:"name"`
	URL   string `json:"url" yaml:"url"`
	About string `json:"about" yaml:"about"`
}

func NewIssueTemplateContactLink() IssueTemplateContactLink {
	return IssueTemplateContactLink{}
}

func (i IssueTemplateContactLink) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()

	if i.Name == "" {
		diag.AddError(errors.New("name is required"))
	}

	if i.URL == "" {
		diag.AddError(errors.New("url is required"))
	}
	if !strings.HasPrefix(i.URL, "https://") {
		diag.AddError(fmt.Errorf("url must start with https://"))
	}

	if i.About == "" {
		diag.AddError(errors.New("about is required"))
	}
	return diag
}
