package template

import (
	"errors"
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type DependabotAllow struct {
	DependencyName string `json:"dependency-name" yaml:"dependency-name"`
	DependencyType string `json:"dependency-type" yaml:"dependency-type"`
}

func NewDependabotAllow() DependabotAllow {
	return DependabotAllow{}
}

func (d DependabotAllow) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()

	if d.DependencyName == "" {
		diag.AddError(errors.New("dependency-name is required"))
	}

	if d.DependencyType == "" {
		diag.AddError(errors.New("dependency-type is required"))
	}

	foundType := false
	for _, t := range DependencyTypes {
		if t == d.DependencyType {
			foundType = true
			break
		}
	}
	if !foundType {
		diag.AddError(fmt.Errorf("dependency-type must be one of the following: %v ", DependencyTypes))
	}

	return diag
}
