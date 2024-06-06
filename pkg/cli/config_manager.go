package cli

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

)

type TerraforgeConfig struct {
	LogLevel string `yaml:"log-level"`
	LogFile  string `yaml:"log-file"`
}

type Config struct {
	terraforgeConfig TerraforgeConfig `yaml:"terraforge_config"`
}

func InitializeConfig(logger *logrus.Logger, configDir string, configFilePath string) (err error) {
	// Create the directory if it doesn't exist
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		logger.Fatalln("Error creating directory:", err)
	}
	homeDir, _ := os.UserHomeDir()
	// create default config
	terraforgeConfig := TerraforgeConfig{
		LogLevel: "info",
		LogFile:  homeDir + "/.terraforge/terraforge.log",
	}
	config := Config{
		terraforgeConfig: terraforgeConfig,
	}

	// Marshal the configuration to YAML
	configYAML, err := yaml.Marshal(&config)
	if err != nil {
		logger.Fatalln("Error marshaling configuration to YAML:", err)
	}
	// Write the YAML to the config file
	err = os.WriteFile(configFilePath, configYAML, 0600)
	if err != nil {
		logger.Fatalln("Error writing configuration to file:", err)
	}

	// // create default log file
	err = os.WriteFile(terraforgeConfig.LogFile, []byte{}, 0666)
	if err != nil {
		logger.Fatalln("Error creating default log file:", err)
	}

	return nil
}
