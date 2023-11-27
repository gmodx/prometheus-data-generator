package main

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	Name       = "Prometheus data generator"
	binaryName = "prometheus_data_generator"
)

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

func main() {
	rootCmd := RootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var (
	configFile string
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   binaryName,
		Short: binaryName,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.Flags().StringVar(&configFile, "config.file", "config.yaml", "configuration file path")
	return cmd
}
