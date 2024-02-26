package interfaces

type ProvisionerConfig struct {
	Language         string `json:"language" yaml:"language" TOML:"language"`
	WorkingDirectory string `json:"working_directory" yaml:"working_directory" TOML:"working-directory"`
}

func NewProvisionerOptions() ProvisionerConfig {
	return ProvisionerConfig{
		Language:         "any",
		WorkingDirectory: ".",
	}
}
