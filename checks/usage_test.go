package checks_test

import (
	"github.com/CiscoCloud/distributive/chkutil"
    "github.com/CiscoCloud/distributive/checks"
	"testing"
)

var smallInts = [][]string{{"0"}, {"1"}, {"2"}}

var bigIntsUnder100 = [][]string{{"100"}, {"99"}, {"98"}}

var reallyBigInts = [][]string{
	{"999999999999999999"}, {"888888888888888888"}, {"777777777777777777"},
}

func TestMemoryUsage(t *testing.T) {
	t.Parallel()
	validInputs := append(smallInts, bigIntsUnder100...)
	invalidInputs := append(append(reallyBigInts, notInts...), negativeInts...)
	testParameters(validInputs, invalidInputs, checks.MemoryUsage{}, t)
	testCheck(bigIntsUnder100, smallInts, checks.MemoryUsage{}, t)
}

func TestSwapUsage(t *testing.T) {
	t.Parallel()
	validInputs := append(smallInts, bigIntsUnder100...)
	invalidInputs := append(append(notLengthOne, notInts...), negativeInts...)
	testParameters(validInputs, invalidInputs, checks.SwapUsage{}, t)
	testCheck(bigIntsUnder100, [][]string{}, checks.SwapUsage{}, t)
}

func testFreeMemoryOrSwap(t *testing.T, chk chkutil.Check) {
	bWinners := suffixParameter(smallInts, "b")
	kbWinners := suffixParameter(smallInts, "kb")
	mbWinners := suffixParameter(smallInts, "mb")
	mbLosers := suffixParameter(reallyBigInts, "mb")
	gbLosers := suffixParameter(reallyBigInts, "gb")
	tbLosers := suffixParameter(reallyBigInts, "tb")
	goodEggs := append(append(bWinners, kbWinners...), mbWinners...)
	badEggs := append(append(mbLosers, gbLosers...), tbLosers...)

	validInputs := append(goodEggs, badEggs...)
	invalidInputs := append(names, notInts...)

	testParameters(validInputs, invalidInputs, chk, t)
	testCheck(goodEggs, badEggs, chk, t)
}

func TestFreeMemory(t *testing.T) {
	t.Parallel()
	testFreeMemoryOrSwap(t, checks.FreeMemory{})
}

func TestFreeSwap(t *testing.T) {
	t.Parallel()
	testFreeMemoryOrSwap(t, checks.FreeSwap{})
}

// $1 - path, $2 maxpercent
func TestDiskUsage(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(dirParameters, "95")
	invalidInputs := append(notLengthTwo,
		[][]string{{"", ""}, {}, {"/", "garble"}}...,
	)
	goodEggs := [][]string{[]string{"/", "99"}}
	badEggs := [][]string{[]string{"/", "1"}}
	testParameters(validInputs, invalidInputs, checks.DiskUsage{}, t)
	testCheck(goodEggs, badEggs, checks.DiskUsage{}, t)
}

func TestInodeUsage(t *testing.T) {
	t.Parallel()
	// TODO: unknown which filesystems would be valid inputs, hence good/bad eggs
	validInputs := [][]string{}
	invalidInputs := append(notLengthTwo,
		[][]string{{"", ""}, {}, {"testfail", "garble"}, {"/dev/testfail"}}...,
	)
	testParameters(validInputs, invalidInputs, checks.DiskUsage{}, t)
}
