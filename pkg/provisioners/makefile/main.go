package makefile

import (
	"errors"

	"github.com/cjlapao/common-go-dependency-tree/dependencytree"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/constants"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/interfaces"
)

type MakeFileProvisioner struct {
	id                 string
	name               string
	languages          []string
	ctx                *context.ProvisionerContext
	cfg                *config.Config
	diagnostics        *diagnostics.Diagnostics
	config             interfaces.ProvisionerConfig
	depTree            *dependencytree.DependencyTreeService[interfaces.Executor]
	dependencyTreeItem *dependencytree.DependencyTreeItem[interfaces.Executor]
}

func New(ctx context.ProvisionerContext, providerConfig interfaces.ProvisionerConfig) *MakeFileProvisioner {
	result := &MakeFileProvisioner{
		id:        constants.MakeFileProvisionerID,
		name:      "Makefile Provisioner",
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

func (p *MakeFileProvisioner) New(ctx context.ProvisionerContext, config interfaces.ProvisionerConfig) interfaces.Provisioner {
	item := New(ctx, config)
	return item
}

func (p *MakeFileProvisioner) Register() error {
	depTree := dependencytree.Get[interfaces.Executor](&MakeFileProvisioner{})
	if depTree == nil {
		p.ctx.LogError("Error getting dependency tree")
		return errors.New("Error getting dependency tree")
	}
	depItem, err := depTree.AddRootItem(p.id, p.name, p)
	if err != nil {
		p.ctx.LogError("Error adding %v to dependency tree, err: %v", p.name, err.Error())
		return err
	}
	p.depTree = depTree
	p.dependencyTreeItem = depItem

	return nil
}

func (p *MakeFileProvisioner) Name() string {
	return p.name
}

func (p *MakeFileProvisioner) ID() string {
	return p.id
}

func (p *MakeFileProvisioner) Context() *context.ProvisionerContext {
	return p.ctx
}

func (p *MakeFileProvisioner) Languages() []string {
	return p.languages
}

func (p *MakeFileProvisioner) Run() diagnostics.Diagnostics {
	p.ctx.LogInfo("Provisioning Makefile")
	// p.cfg.RequestFromUser("IAMGROOT", "Yep Everything is fine from Makefile")
	p.ctx.LogInfo("Makefile provisioned")
	return *p.diagnostics
}

func (p *MakeFileProvisioner) ReshuffleCallback() func() {
	return nil
}
