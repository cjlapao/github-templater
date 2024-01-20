package interfaces

type ProvisionerConfig struct {
	WorkingDirectory string `json:"working_directory" yaml:"working_directory" TOML:"working-directory"`
}

func NewProvisionerOptions() ProvisionerConfig {
	return ProvisionerConfig{
		WorkingDirectory: ".",
	}
}
