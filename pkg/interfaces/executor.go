package interfaces

import (
	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type Executor interface {
	Register() error
	Run() diagnostics.Diagnostics
}
