package cmd

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/kinfinity/terraforge/pkg/cli"
	"github.com/kinfinity/terraforge/pkg/cmd/app"
	"github.com/kinfinity/terraforge/pkg/utils"
	"github.com/kinfinity/terraforge/pkg/utils/templates"
)

type TerraforgeOptions struct {
	Arguments []string
}

var (
	// Terraforge Version
	version string = utils.Version()
)

func NewDefaultTerraforgeCommand(logger *log.Logger) *cobra.Command {
	return NewDefaultTerraforgeCommandWithArgs(TerraforgeOptions{
		Arguments: os.Args,
	},
		logger,
	)
}

func NewDefaultTerraforgeCommandWithArgs(options TerraforgeOptions, logger *log.Logger) *cobra.Command {
	return NewTerraforgeCommand(logger)
}

func NewTerraforgeCommand(logger *log.Logger) *cobra.Command {
	// Create Root Command and add other commands downstream

	root := &cobra.Command{
		Use:   strings.ToLower(utils.Name),
		Short: "Terraform/OpenTofu environment and module management",
		Long:  "Terraforge is an IaC environment management bootstrapper and tool for OpenTofu and Terraform",
		// Setup Version flag
		Version: version,
	}

	var (
		configFileName string = "/terraforge.yaml"
		configFilePath string
	)
	configDir := cli.SetupConfig(log.New(), configFileName)

	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configDir)

	// Set up a persistent flag for the configFilePath
	root.Flags().StringVarP(&configFilePath, "config", "c", "", "config file (default is ./config.yaml)")
	// // Bind the viper configuration to the configFilePath flag1
	viper.BindPFlag("config", root.Flags().Lookup("config"))
	// Check config via flag
	if root.Flags().Changed("config") {
		// Configure Viper from --config
		viper.AddConfigPath(configFilePath) // --config overwrites  config in the .terraforge directory
		logger.Info("Overwrite Config File =", configFilePath+configFileName)
		if _, err := os.Stat(configFilePath + configFileName); os.IsNotExist(err) {
			logger.Error("New Config File Not Found: ", err)
		}
	}

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("fatal error config file:", err.Error())
	}
	logger.Info("Log File =", viper.GetString("terraforge.log-file"))

	commandGroup := &templates.CommandGroups{
		{
			Message: "Terraforge App Commands",
			Commands: []*cobra.Command{
				app.CreateAppCommand(logger),
			},
		},
	}
	commandGroup.Add(root)

	return root
}
