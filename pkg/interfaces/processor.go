package interfaces

import "github.com/cjlapao/github-templater/pkg/diagnostics"

type Processor interface {
	ID() string
	Name() string
	Run() diagnostics.Diagnostics
}
