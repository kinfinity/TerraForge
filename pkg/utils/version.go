package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver/v3"

)

const Name = "TerraForge"

var (
	// Current Version based on git Tag
	version *semver.Version
)

func init() {
	// Get the version from git tags
	if version == nil {
		_version, err := getVersionFromGit()
		if err != nil {
			// log.Printf("Failed to get version from git tags")
			// don't panic - fallback on dev
			version, _ = semver.NewVersion("v0.0.0-dev")
		} else {
			version, _ = semver.NewVersion(_version)
		}
		// log.Printf("version: %s", version)
	}
}

// getVersionFromGit retrieves the version from the latest git tag
func getVersionFromGit() (string, error) {
	cmd := exec.Command("git", "tag", "--list='v*.*.*'")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	log.Printf("output: %s ", output)
	// Trim leading/trailing whitespaces and newline characters
	version := strings.TrimSpace(string(output[0]))

	return version, nil
}

func FullVersion() string {
	return fmt.Sprintf("%s Version %s", Name, version)
}

func Version() string {
	return version.String()
}
