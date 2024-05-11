/*
Copyright © 2024 kokou.egbewatt@gmail.com

	TERRAFORGE - Easing Terraform environment bootstrapping and development
*/
package main

import (
	"fmt"
	"os"

)

func main() {
	if len(os.Args) <= 1 {
		PrintUsage()
		os.Exit(0)
	}
}

func PrintUsage() {
	fmt.Println(
		`USAGE:
		
terraforge <command> <args>`)
}
