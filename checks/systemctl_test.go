package checks_test

import (
	"testing"
    "github.com/CiscoCloud/distributive/checks"
)

var activeServices = [][]string{
	{"dev-mqueue.mount"}, {"tmp.mount"}, {"dbus.service"},
}

func TestSystemctlLoaded(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.SystemctlLoaded{}, t)
	testCheck(activeServices, names, checks.SystemctlLoaded{}, t)
}

func TestSystemctlActive(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.SystemctlActive{}, t)
	testCheck(activeServices, names, checks.SystemctlLoaded{}, t)
}

func TestSystemctlSockPath(t *testing.T) {
	t.Parallel()
	invalidInputs := append(notLengthOne, names...)
	testParameters(fileParameters, invalidInputs, checks.SystemctlSockListening{}, t)
	testCheck([][]string{}, fileParameters, checks.SystemctlSockListening{}, t)
}

func TestSystemctlTimer(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.SystemctlTimer{}, t)
	testCheck([][]string{}, names, checks.SystemctlTimer{}, t)
}

func TestSystemctlTimerLoaded(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.SystemctlTimerLoaded{}, t)
	testCheck([][]string{}, names, checks.SystemctlTimerLoaded{}, t)
}

func TestSystemctlUnitFileStatus(t *testing.T) {
	t.Parallel()
	goodEggs := [][]string{
		{"dbus.service", "static"},
		{"polkit.service", "static"},
		{"systemd-initctl.service", "static"},
	}
	validInputs := appendParameter(names, "static")
	testParameters(validInputs, notLengthTwo, checks.SystemctlUnitFileStatus{}, t)
	testCheck(goodEggs, validInputs, checks.SystemctlUnitFileStatus{}, t)
}
