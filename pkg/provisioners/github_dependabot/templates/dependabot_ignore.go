package template

import (
	"errors"
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type DependabotIgnore struct {
	DependencyName string   `json:"dependency-name" yaml:"dependency-name"`
	UpdateTypes    []string `json:"update-types" yaml:"update-types"`
	Versions       []string `json:"versions" yaml:"versions"`
}

func NewDependabotIgnore() DependabotIgnore {
	return DependabotIgnore{
		UpdateTypes: []string{},
		Versions:    []string{},
	}
}

func (d DependabotIgnore) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()

	if d.DependencyName == "" {
		diag.AddError(errors.New("dependency-name is required"))
	}

	if len(d.UpdateTypes) == 0 {
		diag.AddError(errors.New("update-types is required"))
	} else {
		for _, t := range d.UpdateTypes {
			foundType := false
			for _, ut := range UpdateTypes {
				if t == ut {
					foundType = true
					break
				}
			}
			if !foundType {
				diag.AddError(fmt.Errorf("update-types must be one of the following: %v, found %v", UpdateTypes, t))
			}
		}
	}

	if len(d.Versions) == 0 {
		diag.AddError(errors.New("versions is required"))
	}

	return diag
}
