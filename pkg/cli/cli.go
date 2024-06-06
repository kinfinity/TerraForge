package cli

import (
	"io"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

)

type TerraforgeCli struct {
	Name       string
	Version    string
	ConfigFile string
	IOStreams  struct {
		In  io.Reader
		Out io.Writer
		Err io.Writer
	}
	logger *log.Logger
}

func (t *TerraforgeCli) Run(cmd *cobra.Command) error {

	cmd.SetOutput(t.logger.Out)
	// Execute the Cobra command
	if err := cmd.Execute(); err != nil {
		t.logger.Panic("fatal error: ", err.Error())
		return err
	}
	return nil
}

func NewCli() *TerraforgeCli {
	/*
		Handle Root Configs
	*/
	var (
		cliName        string = "Terraforge"
		configFileName string = "/config.yaml"
		logFileName    string = "/.terraforge/terraforge.log"
	)

	homeDir, _ := os.UserHomeDir()
	logFilePath := path.Join(homeDir, logFileName)

	// Open log file
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Error opening default log file:", err)
	}

	logger := log.New()
	logger.Out = io.MultiWriter(
		os.Stdout,
		NewFileWriter(logFile),
	)
	// logger.SetFormatter(log.New().Formatter)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{})

	// Setup Config
	configDir := SetupConfig(logger, configFileName)
	logger.Info(configDir)

	return &TerraforgeCli{
		Name:    cliName,
		Version: "", // Pick from master version
		logger:  logger,
	}
}

// Return
func (t *TerraforgeCli) Logger() (logger *log.Logger) {
	return t.logger
}
