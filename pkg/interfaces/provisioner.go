package interfaces

import (
	provisioner_context "github.com/cjlapao/github-templater/pkg/context"
)

type Provisioner interface {
	New(ctx provisioner_context.ProvisionerContext, config ProvisionerConfig) Provisioner
	Name() string
	ID() string
	Context() provisioner_context.ProvisionerContext
	Provision() error
	Languages() []string
}
