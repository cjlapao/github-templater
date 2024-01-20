package makefile

import (
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/interfaces"
)

type MakeFileProvisioner struct {
	id          string
	name        string
	languages   []string
	ctx         *context.ProvisionerContext
	cfg         *config.Config
	diagnostics *diagnostics.Diagnostics
	config      interfaces.ProvisionerConfig
}

func New(ctx context.ProvisionerContext, providerConfig interfaces.ProvisionerConfig) *MakeFileProvisioner {
	result := &MakeFileProvisioner{
		id:        "31264948-5ec7-4d7f-abe3-fc7ec7f70344",
		name:      "makefile",
		ctx:       &ctx,
		cfg:       config.Get(),
		languages: []string{"all"},
		config:    providerConfig,
	}

	result.ctx.WithValue(constants.ContextResourceName, result.name)
	diagnostics := diagnostics.NewModuleDiagnostics(result.name)
	result.diagnostics = &diagnostics

	return result
}

func (p MakeFileProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p MakeFileProvisioner) Name() string {
	return p.name
}

func (p MakeFileProvisioner) ID() string {
	return p.id
}

func (p MakeFileProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p MakeFileProvisioner) Languages() []string {
	return p.languages
}

func (p MakeFileProvisioner) Provision() diagnostics.Diagnostics {
	p.ctx.LogInfo("Provisioning Makefile")
	// p.cfg.RequestFromUser("IAMGROOT", "Yep Everything is fine from Makefile")
	p.ctx.LogInfo("Makefile provisioned")
	return *p.diagnostics
}
