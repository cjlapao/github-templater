package provision

import (
	"fmt"

	"github.com/cjlapao/common-go-dependency-tree/dependencytree"
	logger "github.com/cjlapao/common-go-logger"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"github.com/cjlapao/github-templater/pkg/interfaces"
	"github.com/cjlapao/github-templater/pkg/provisioners/github"
	"github.com/pkg/errors"
)

var SupportedLanguages = []string{
	"go",
	"golang",
}

type ProvisionerService struct {
	cfg          *config.Config
	ctx          *context.ProvisionerContext
	provisioners []interfaces.Provisioner
	config       interfaces.ProvisionerConfig
}

func New() *ProvisionerService {
	result := &ProvisionerService{
		cfg:          config.Get(),
		ctx:          context.Get(),
		provisioners: []interfaces.Provisioner{},
		config:       interfaces.ProvisionerConfig{},
	}

	result.config = result.getConfig()
	return result
}

func (p *ProvisionerService) Add(provisioner interfaces.Provisioner) {
	newProvisioner := provisioner.New(*p.ctx, p.config)
	if executor, ok := newProvisioner.(interfaces.Executor); ok {
		executor.Register()
	}
	p.provisioners = append(p.provisioners, newProvisioner)
}

func (p *ProvisionerService) Get() []interfaces.Provisioner {
	return p.provisioners
}

func (p *ProvisionerService) GetByName(name string) interfaces.Provisioner {
	for _, provisioner := range p.provisioners {
		if provisioner.Name() == name {
			return provisioner
		}
	}

	return nil
}

func (p *ProvisionerService) GetByID(id string) interfaces.Provisioner {
	for _, provisioner := range p.provisioners {
		if provisioner.ID() == id {
			return provisioner
		}
	}

	return nil
}

func (p *ProvisionerService) GetByLanguage(language string) []interfaces.Provisioner {
	result := []interfaces.Provisioner{}

	for _, provisioner := range p.provisioners {
		for _, lang := range provisioner.Languages() {
			if lang == language || lang == "all" || language == "any" {
				result = append(result, provisioner)
			}
		}
	}

	dependencyTree := dependencytree.Get[interfaces.Executor](&github.GitHubProvisioner{})
	dependencyTree.SetDebug(config.Get().IsDebug())
	dependencyTree.SetVerbose(config.Get().IsDebug())

	_, err := dependencyTree.Build()
	if err != nil {
		p.ctx.LogError(err.Error())
	}

	return result
}

func (p *ProvisionerService) Provision() diagnostics.Diagnostics {
	diag := diagnostics.NewModuleDiagnostics("provisioner")
	language := p.cfg.RequestFromUser("LANGUAGE", fmt.Sprintf("%v What language do you want to provision?", logger.IconBell), "")
	if language == "" {
		diag.AddError(errors.New("No language provided"))
		return diag
	}
	if !p.isLanguageSupported(language) {
		msg := fmt.Sprintf("Language %v is not supported", language)
		diag.AddError(errors.New(msg))
		p.ctx.LogError(msg)
		return diag
	}

	p.ctx.LogInfo("Provisioning for language %v", language)
	p.config.Language = language
	availableProvisioners := p.GetByLanguage(language)
	if len(availableProvisioners) == 0 {
		p.ctx.LogError("No provisioners found for language %v", language)
		diag.AddInfo(fmt.Sprintf("No provisioners found for language %v", language))
		return diag
	}

	for _, provisioner := range availableProvisioners {
		base := provisioner.New(*p.ctx, p.config)
		if provisionerDiag := base.Run(); provisionerDiag.HasErrors() {
			diag.Append(provisionerDiag)
		}
	}

	return diag
}

func (p *ProvisionerService) getConfig() interfaces.ProvisionerConfig {
	options := interfaces.ProvisionerConfig{}
	options.WorkingDirectory = p.cfg.GetKey("WORKING_DIRECTORY")
	if options.WorkingDirectory == "" {
		options.WorkingDirectory = "."
	}
	return options
}

func (p *ProvisionerService) isLanguageSupported(language string) bool {
	for _, lang := range SupportedLanguages {
		if lang == language {
			return true
		}
	}

	return false
}
