package checks_test

import (
    "testing"
    "github.com/CiscoCloud/distributive/checks"
)

func TestCommand(t *testing.T) {
	t.Parallel()
	validInputs := [][]string{
		{"sleep 0.00000001"}, {"echo this works"}, {"cd"}, {"mv --help"},
	}
	invalidInputs := notLengthOne
	goodEggs := validInputs
	badEggs := [][]string{
		{"sleep fail"}, {"cd /steppenwolf"}, {"mv /glass /bead-game"},
	}
	badEggs = append(badEggs, names...)
	testParameters(validInputs, invalidInputs, checks.Command{}, t)
	testCheck(goodEggs, badEggs, checks.Command{}, t)
}

func TestCommandOutputMatches(t *testing.T) {
	t.Parallel()
	validInputs := [][]string{
		{"echo siddhartha", "sid"}, {"cp --help", "cp"}, {"echo euler", "eu"},
	}
	invalidInputs := notLengthTwo
	goodEggs := validInputs
	badEggs := [][]string{
		{"echo siddhartha", "fail"},
		{"cp --help", "asdfalkjsdhldjfk"},
		{"echo haskell", "curry"},
	}
	testParameters(validInputs, invalidInputs, checks.CommandOutputMatches{}, t)
	testCheck(goodEggs, badEggs, checks.CommandOutputMatches{}, t)
}

func TestRunning(t *testing.T) {
	t.Parallel()
	validInputs := append(names, [][]string{
		{"proc"}, {"nginx"}, {"anything"}, {"worker"}, {"distributive"},
	}...)
	invalidInputs := notLengthOne
	goodEggs := [][]string{}
	badEggs := dirParameters
	testParameters(validInputs, invalidInputs, checks.Running{}, t)
	testCheck(goodEggs, badEggs, checks.Running{}, t)
}

func TestParseSensorsOutput(t *testing.T) {
	sensorsOut := `
	k8temp-pci-00c3
	Adapter: PCI adapter
	Core0 Temp:  +30.0°C
	Core0 Temp:  +30.0°C
	Core1 Temp:  +29.0°C
	Core1 Temp:  +36.0°C
	`
	t.Parallel()
	if len(checks.ParseSensorsOutput(sensorsOut)) < 1 {
		t.Errorf("parseSensorsOutput didn't parse correctly")
	}
}

func TestTemp(t *testing.T) {
	t.Parallel()
	validInputs := positiveInts[:len(positiveInts)-2] // only small ints
	invalidInputs := append(append(names, notInts...), notLengthOne...)
	testParameters(validInputs, invalidInputs, checks.Temp{}, t)
}

func TestModule(t *testing.T) {
	t.Parallel()
	validInputs := names
	invalidInputs := notLengthOne
	goodEggs := [][]string{}
	badEggs := names
	testParameters(validInputs, invalidInputs, checks.Module{}, t)
	testCheck(goodEggs, badEggs, checks.Module{}, t)
}

func TestKernelParameter(t *testing.T) {
	validInputs := names
	invalidInputs := notLengthOne
	goodEggs := [][]string{
		{"net.ipv4.conf.all.accept_local"},
		{"net.ipv4.conf.all.accept_redirects"},
		{"net.ipv4.conf.all.arp_accept"},
	}
	badEggs := names
	testParameters(validInputs, invalidInputs, checks.KernelParameter{}, t)
	testCheck(goodEggs, badEggs, checks.KernelParameter{}, t)
}

func TestPHPConfig(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(names, "dummy-value")
	invalidInputs := notLengthTwo
	testParameters(validInputs, invalidInputs, checks.PHPConfig{}, t)
}
