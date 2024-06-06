package app

import (
	"io/fs"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/kinfinity/terraforge/pkg/cli"
	"github.com/kinfinity/terraforge/pkg/utils"

)

var AppTree = []string{
	"/bootstrap",
	"/environments",
	"/modules",
}

type AppOptions struct {
	Name     string
	RootDir  fs.FileInfo
	Executor *utils.Executor
}

type TerraforgeApp struct {
	Name     string
	RootDir  os.File
	Executor *utils.Executor
}

var (
	executor string
)

func NewTerraforgeApp(options *AppOptions) *TerraforgeApp {
	return &TerraforgeApp{
		Name:     options.Name,
		Executor: options.Executor,
	}
}

func CreateAppCommand(logger *log.Logger) *cobra.Command {

	createCommand := &cobra.Command{
		Use:   "create-app <name>",
		Short: "Create new Terraforge App",
		Long:  "Creates a new Terraforge App to manage multi-environment Infrastructure with supported Executors\n Executors: \n - terraform \n - opentofu",
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.Help()
				os.Exit(0)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			AppName := args[0]

			// Setup Options
			workingDir, err := cli.IsEmptyDir(AppName)
			if err != nil {
				// print help on fail
				return err
			}

			if created, err := cli.CreateAppStructure(AppName, AppTree); err != nil && !created {
				return err
			}

			options := &AppOptions{
				Name:     AppName,
				RootDir:  workingDir,
				Executor: utils.NewExecutor(executor),
			}

			// Create New Terraforge App
			tfApp := NewTerraforgeApp(options)
			logger.Info(tfApp.Name + "\n")
			logger.Info("Executor: " + tfApp.Executor.Name.String())
			logger.Info("where: " + tfApp.Executor.Path)

			return nil

		},
	}

	createCommand.
		Flags().
		StringVarP(
			&executor,
			"executor", "e",
			"terraform",
			"This flag is used to Setup the Environment Executor",
		)

	return createCommand
}

// After command has been validated check for app in location with Name or .
//
// Create with default executor & backend [ AWS| ... ]
// -- executor [ OpenTofu | Terraform ]
//
