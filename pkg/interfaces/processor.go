package interfaces

import "github.com/cjlapao/github-templater/pkg/diagnostics"

type Processor interface {
	ID() string
	Name() string
	Process() diagnostics.Diagnostics
}
