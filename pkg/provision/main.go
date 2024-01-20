package provision

import (
	"fmt"
	"sync"

	logger "github.com/cjlapao/common-go-logger"
	"github.com/cjlapao/github-templater/config"
	"github.com/cjlapao/github-templater/context"
	"github.com/cjlapao/github-templater/interfaces"
	"github.com/pkg/errors"
)

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

	return result
}

func (p *ProvisionerService) Provision() error {
	language := p.cfg.RequestFromUser("LANGUAGE", fmt.Sprintf("%v What language do you want to provision?", logger.IconBell))
	if language == "" {
		return errors.New("No language provided")
	}

	p.ctx.LogInfo("Provisioning for language %v", language)
	availableProvisioners := p.GetByLanguage(language)
	if len(availableProvisioners) == 0 {
		p.ctx.LogError("No provisioners found for language %v", language)
		return nil
	}

	wg := sync.WaitGroup{}
	for _, provisioner := range availableProvisioners {
		wg.Add(1)
		go func(provisioner interfaces.Provisioner) {
			base := provisioner.New(*p.ctx, p.config)
			err := base.Provision()
			if err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(provisioner)
	}

	wg.Wait()

	return nil
}

func (p *ProvisionerService) getConfig() interfaces.ProvisionerConfig {
	options := interfaces.ProvisionerConfig{}
	options.WorkingDirectory = p.cfg.GetKey("WORKING_DIRECTORY")
	if options.WorkingDirectory == "" {
		options.WorkingDirectory = "."
	}
	return options
}
