package template

import (
	"fmt"

	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type DependabotGroup struct {
	DependencyType  string   `json:"dependency-type,omitempty" yaml:"dependency-type,omitempty"`
	Patterns        []string `json:"patterns,omitempty" yaml:"patterns,omitempty"`
	ExcludePatterns []string `json:"exclude-patterns,omitempty" yaml:"exclude-patterns,omitempty"`
	UpdateTypes     []string `json:"update-types,omitempty" yaml:"update-types,omitempty"`
}

func NewDependabotGroup() DependabotGroup {
	return DependabotGroup{
		Patterns:        []string{},
		ExcludePatterns: []string{},
		UpdateTypes:     []string{},
	}
}

func (d DependabotGroup) Validate() diagnostics.Diagnostics {
	diag := diagnostics.New()
	if len(d.Patterns) == 0 {
		diag.AddError(fmt.Errorf("patterns is required"))
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

	if len(d.UpdateTypes) == 0 {
		diag.AddError(fmt.Errorf("update-types is required"))
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

	return diag
}
