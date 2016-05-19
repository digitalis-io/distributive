package checks_test

import (
	"testing"
    "github.com/CiscoCloud/distributive/checks"
)

func TestZooKeeperRUOK(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping docker tests in short mode")
	} else {
		validInputs := [][]string{
			[]string{"2ms", "wikipedia.org:9814"},
			[]string{"1ms", "mozilla.org:9814"},
		}
		invalidInputs := [][]string{
			{"", "mozilla.net"},
			{"nottime", "wikipedia.org"},
		}
		// inputs that should lead to success
		goodEggs := [][]string{}
		// inputs that should lead to failure
		testParameters(validInputs, invalidInputs, checks.ZooKeeperRUOK{}, t)
		testCheck(goodEggs, validInputs, checks.ZooKeeperRUOK{}, t)
	}
}
