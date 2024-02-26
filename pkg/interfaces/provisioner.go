package interfaces

import (
	provisioner_context "github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

type Provisioner interface {
	New(ctx provisioner_context.ProvisionerContext, config ProvisionerConfig) Provisioner
	Name() string
	ID() string
	Context() *provisioner_context.ProvisionerContext
	Run() diagnostics.Diagnostics
	Languages() []string
}
