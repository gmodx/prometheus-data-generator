package config

type Config struct {
	TemplateDir string `json:"template_dir"`
	Output      string `json:"output"`

	TemplateWithoutUnixConfigs []TemplateWithoutUnixConfig `json:"template_without_unix_configs"`
	TemplateWithUnixConfigs    []TemplateWithUnixConfig    `json:"template_with_unix_configs"`
}

type GlobalConfig struct {
	TemplateValuesPath string `json:"template_values_path"`
	Days               int    `json:"days"`
	ResolutionSeconds  int    `json:"resolution_seconds"`
	EndtimeUnix        int    `json:"endtime_unix"`
}

type TemplateWithoutUnixConfig struct {
	Name               string `json:"name"`
	TemplateValuesPath string `json:"template_values_path"`
	Days               int    `json:"days"`
	ResolutionSeconds  int    `json:"resolution_seconds"`
	EndtimeUnix        int    `json:"endtime_unix"`
}

type TemplateWithUnixConfig struct {
	Name               string `json:"name"`
	TemplateValuesPath string `json:"template_values_path"`
}
