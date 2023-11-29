package main

import (
	"context"
	"path"

	"github.com/gmodx/prometheus-data-generator/config"
	"github.com/gmodx/prometheus-data-generator/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Name       = "Prometheus data generator"
	binaryName = "prometheus_data_generator"
)

var (
	configFile string
)

func main() {
	rootCmd := RootCmd()
	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func getCfg() *config.Config {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read config file: %v", err)
	}

	var config config.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Failed to unmarshal config: %v", err)
	}

	for i := range config.TemplateWithoutUnixConfigs {
		if config.TemplateWithoutUnixConfigs[i].TemplateValuePath == "" {
			config.TemplateWithoutUnixConfigs[i].TemplateValuePath = config.GlobalConfig.TemplateValuePath
		}
		if config.TemplateWithoutUnixConfigs[i].Days == 0 {
			config.TemplateWithoutUnixConfigs[i].Days = config.GlobalConfig.Days
		}
		if config.TemplateWithoutUnixConfigs[i].EndTimeUnix == 0 {
			config.TemplateWithoutUnixConfigs[i].EndTimeUnix = config.GlobalConfig.EndTimeUnix
		}
		if config.TemplateWithoutUnixConfigs[i].ResolutionSeconds == 0 {
			config.TemplateWithoutUnixConfigs[i].ResolutionSeconds = config.GlobalConfig.ResolutionSeconds
		}
	}

	return &config
}

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   binaryName,
		Short: binaryName,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := getCfg()
			ctx := context.Background()
			blockHours := 2

			for _, templateConfig := range cfg.TemplateWithoutUnixConfigs {
				tplPath := path.Join(cfg.TemplateDir, templateConfig.Name+".template")
				tplValuePath := path.Join(cfg.TemplateValueDir, templateConfig.TemplateValuePath)

				GenerateSamples_WithoutUnix(ctx, templateConfig.Name, tplPath, tplValuePath, cfg.OutputDir, templateConfig.ResolutionSeconds, templateConfig.StartTime(), templateConfig.EndTime(), blockHours)
			}

			for _, templateConfig := range cfg.TemplateWithUnixConfigs {
				tplPath := path.Join(cfg.TemplateDir, templateConfig.Name+".template")
				tplValuePath := path.Join(cfg.TemplateValueDir, templateConfig.TemplateValuePath)

				GenerateSamples_WithUnix(ctx, templateConfig.Name, tplPath, tplValuePath, cfg.OutputDir, blockHours)
			}
		},
	}

	cmd.Flags().StringVar(&configFile, "config.file", "config.yaml", "configuration file path")
	return cmd
}
