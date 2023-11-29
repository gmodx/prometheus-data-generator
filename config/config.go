package config

import "time"

type Config struct {
	TemplateDir      string `mapstructure:"template_dir"`
	TemplateValueDir string `mapstructure:"template_value_dir"`
	OutputDir        string `mapstructure:"output_dir"`

	GlobalConfig GlobalConfig `mapstructure:"global_config"`

	TemplateWithoutUnixConfigs []TemplateWithoutUnixConfig `mapstructure:"template_without_unix_configs"`
	TemplateWithUnixConfigs    []TemplateWithUnixConfig    `mapstructure:"template_with_unix_configs"`
}

type GlobalConfig struct {
	TemplateValuePath string `mapstructure:"template_value_path"`
	Days              int    `mapstructure:"days"`
	ResolutionSeconds int    `mapstructure:"resolution_seconds"`
	EndTimeUnix       int64  `mapstructure:"endtime_unix"`
}

type TemplateWithoutUnixConfig struct {
	Name string `mapstructure:"name"`

	TemplateValuePath string `mapstructure:"template_value_path"`
	Days              int    `mapstructure:"days"`
	ResolutionSeconds int    `mapstructure:"resolution_seconds"`
	EndTimeUnix       int64  `mapstructure:"endtime_unix"`
}

type TemplateWithUnixConfig struct {
	Name string `mapstructure:"name"`

	TemplateValuePath string `mapstructure:"template_value_path"`
}

func (c TemplateWithoutUnixConfig) EndTime() time.Time {
	return time.Unix(c.EndTimeUnix, 0)
}

func (c TemplateWithoutUnixConfig) StartTime() time.Time {
	return c.EndTime().AddDate(0, 0, -c.Days)
}
