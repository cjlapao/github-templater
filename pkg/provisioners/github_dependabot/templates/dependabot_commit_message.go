package template

import (
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type DependabotCommitMessage struct {
	Prefix            string `json:"prefix" yaml:"prefix"`
	PrefixDevelopment string `json:"prefix-development" yaml:"prefix-development"`
	Include           string `json:"include" yaml:"include"`
}

func NewDependabotCommitMessage() DependabotCommitMessage {
	return DependabotCommitMessage{
		Include: "scope",
	}
}

func (d DependabotCommitMessage) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if d.Prefix != "" && len(d.Prefix) > 50 {
		diag.AddError(fmt.Errorf("prefix must be less than 50 characters"))
	}

	if d.PrefixDevelopment != "" && len(d.PrefixDevelopment) > 50 {
		diag.AddError(fmt.Errorf("prefix-development must be less than 50 characters"))
	}

	if d.Include != "" && d.Include != "scope" {
		diag.AddError(fmt.Errorf("include must be scope"))
	}

	return diag
}
