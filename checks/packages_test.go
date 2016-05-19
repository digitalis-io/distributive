package checks_test

import (
	"strings"
	"testing"
    "github.com/CiscoCloud/distributive/checks"

	gosssystem "github.com/aelsabbahy/goss/system"
)

func TestPacmanIgnore(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.PacmanIgnore{}, t)
	if gosssystem.DetectPackageManager() == "pacman" {
		testCheck([][]string{}, names, checks.PacmanIgnore{}, t)
	}
}

func TestInstalled(t *testing.T) {
	t.Parallel()
	// dpkg will fail on invalid package name
	validPackageNames := [][]string{}
	for _, name := range names {
		newName := strings.Replace(name[0], " ", "-", -1)
		validPackageNames = append(validPackageNames, []string{newName})
	}
	testParameters(validPackageNames, notLengthOne, checks.Installed{}, t)
	testCheck([][]string{}, names, checks.Installed{}, t)
}
