package utils

import (
	"os/exec"
	"strings"
)

type ExecutorVal int

const (
	Terraform ExecutorVal = iota
	OpenTofu
)

// Get all executor values
func AllExecutors() []ExecutorVal {
	return []ExecutorVal{Terraform, OpenTofu}
}

// To String for Executor values
func (e ExecutorVal) String() string {
	switch e {
	case Terraform:
		return "Terraform"
	case OpenTofu:
		return "OpenTofu"
	default:
		return "Unknown"
	}
}

const (
	DefaultTerraformBin string = "/usr/bin/terraform"
	DefaultOpenTofuBin  string = ""
)

type Executor struct {
	Name ExecutorVal
	Path string
}

// Create and Validate Executor Instance
func NewExecutor(name string) *Executor {
	for _, v := range AllExecutors() {
		if strings.Contains(strings.ToLower(name), strings.ToLower(v.String())) {
			e := &Executor{
				Name: v,
				Path: "",
			}
			e.version()
			return e
		}
	}

	return nil
}

func (e *Executor) version() string {
	cmd := exec.Command(strings.ToLower(e.Name.String()), "--version")
	ExecutorVersion, err := cmd.Output()
	if err != nil {
		return err.Error()
	}

	return string(ExecutorVersion)
}
