/*
Copyright © 2024 kegbewatt@aortem.io

	TERRAFORGE - Easing Terraform environment bootstrapping and development
*/
package main

import (
	"github.com/kinfinity/terraforge/pkg/cli"
	"github.com/kinfinity/terraforge/pkg/cmd"
)

func main() {

	terraforgeCli := cli.NewCli()

	command := cmd.NewDefaultTerraforgeCommand(terraforgeCli.Logger())

	if err := terraforgeCli.Run(command); err != nil {
		terraforgeCli.Logger().Fatalf("%s", err)
	}

}
