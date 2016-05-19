package checks_test

import (
	"testing"
    "github.com/CiscoCloud/distributive/checks"
)

var validHosts = [][]string{{"eff.org"}, {"mozilla.org"}, {"golang.org"}}

var invalidHosts = [][]string{
	{"serbapsidpuflnaskjdcasd.com"},
	{"eroiqewruqpwioepbpasdfb.net"},
	{"rjblrjbbrbbbnnzasdflbaj.org"},
}

var validURLs = prefixParameter(validHosts, "http://")
var invalidURLs = prefixParameter(invalidHosts, "http://")
var validHostsWithPort = suffixParameter(validHosts, ":80")
var invalidHostsWithPort = suffixParameter(invalidHosts, ":80")

var closedPorts = [][]string{
	{"49151"}, // reserved
	{"5310"},  // Outlaws (1997 video game)
	{"0"},     // reserved
	{"2302"},  // Halo: Combat Evolved multiplayer
}

func TestPort(t *testing.T) {
	t.Parallel()
	// only take smaller ones
	validInputs := positiveInts[:len(positiveInts)-2]
	invalidInputs := append(notInts, negativeInts...)
	testParameters(validInputs, invalidInputs, checks.Port{}, t)
	testCheck([][]string{}, closedPorts, checks.Port{}, t)
}

func TestPortTCP(t *testing.T) {
	t.Parallel()
	// only take smaller ones
	validInputs := positiveInts[:len(positiveInts)-2]
	invalidInputs := append(notInts, negativeInts...)
	testParameters(validInputs, invalidInputs, checks.PortTCP{}, t)
	testCheck([][]string{}, closedPorts, checks.PortTCP{}, t)
}

func TestPortUDP(t *testing.T) {
	t.Parallel()
	// only take smaller ones
	validInputs := positiveInts[:len(positiveInts)-2]
	invalidInputs := append(notInts, negativeInts...)
	testParameters(validInputs, invalidInputs, checks.PortUDP{}, t)
	testCheck([][]string{}, closedPorts, checks.PortUDP{}, t)
}

func TestInterfaceExists(t *testing.T) {
	t.Parallel()
	validInputs := names
	invalidInputs := notLengthOne
	goodEggs := [][]string{}
	badEggs := [][]string{}
	testParameters(validInputs, invalidInputs, checks.InterfaceExists{}, t)
	testCheck(goodEggs, badEggs, checks.InterfaceExists{}, t)
}

func TestUp(t *testing.T) {
	t.Parallel()
	validInputs := names
	invalidInputs := notLengthOne
	goodEggs := [][]string{}
	badEggs := [][]string{}
	testParameters(validInputs, invalidInputs, checks.Up{}, t)
	testCheck(goodEggs, badEggs, checks.Up{}, t)
}

func TestIP4(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(names, "0000:000:0000:000:0000:0000:000:0000")
	invalidInputs := notLengthTwo
	goodEggs := [][]string{}
	badEggs := validInputs
	testParameters(validInputs, invalidInputs, checks.IP4{}, t)
	testCheck(goodEggs, badEggs, checks.IP4{}, t)
}

func TestIP6(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(names, "0000:000:0000:000:0000:0000:000:0000")
	invalidInputs := notLengthTwo
	goodEggs := [][]string{}
	badEggs := validInputs
	testParameters(validInputs, invalidInputs, checks.IP6{}, t)
	testCheck(goodEggs, badEggs, checks.IP6{}, t)
}

func TestGatewayInterface(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(names, "0000:000:0000:000:0000:0000:000:0000")
	validInputs = append(validInputs, appendParameter(names, "192.168.0.1")...)
	testParameters(validInputs, notLengthTwo, checks.IP6{}, t)
	testCheck([][]string{}, validInputs, checks.IP6{}, t)
}

func TestHost(t *testing.T) {
	t.Parallel()
	validInputs := names
	invalidInputs := notLengthOne
	goodEggs := validHosts
	badEggs := invalidHosts
	testParameters(validInputs, invalidInputs, checks.Host{}, t)
	testCheck(goodEggs, badEggs, checks.Host{}, t)
}

func TestTCP(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.TCP{}, t)
	testCheck(validHostsWithPort, invalidHostsWithPort, checks.TCP{}, t)
}

func TestUDP(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.UDP{}, t)
	testCheck(validHostsWithPort, invalidHostsWithPort, checks.UDP{}, t)
}

func TestTCPTimeout(t *testing.T) {
	t.Parallel()
	goodEggs := appendParameter(validHostsWithPort, "5s")
	badEggs := appendParameter(validHostsWithPort, "1µs")
	validInputs := appendParameter(names, "5s")
	testParameters(validInputs, notLengthTwo, checks.TCPTimeout{}, t)
	testCheck(goodEggs, badEggs, checks.TCPTimeout{}, t)
}

func TestUDPTimeout(t *testing.T) {
	t.Parallel()
	goodEggs := appendParameter(validHostsWithPort, "5s")
	badEggs := appendParameter(validHostsWithPort, "1µs")
	validInputs := appendParameter(names, "5s")
	testParameters(validInputs, notLengthTwo, checks.UDPTimeout{}, t)
	testCheck(goodEggs, badEggs, checks.UDPTimeout{}, t)
}

func TestRoutingTableDestination(t *testing.T) {
	t.Parallel()
	// TODO get a list of valid IP addresses for these valid params
	invalidInputs := append(names, notLengthOne...)
	testParameters([][]string{}, invalidInputs, checks.RoutingTableDestination{}, t)
	//testCheck([][]string{}, [][]string{}, RoutingTableDestination{}, t)
}

func TestRoutingTableInterface(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.RoutingTableInterface{}, t)
	testCheck([][]string{}, names, checks.RoutingTableInterface{}, t)
}

func TestRoutingTableGateway(t *testing.T) {
	t.Parallel()
	testParameters(names, notLengthOne, checks.RoutingTableGateway{}, t)
	testCheck([][]string{}, names, checks.RoutingTableGateway{}, t)
}

func TestReponseMatches(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping tests that query remote servers in short mode")
	} else {
		validInputs := appendParameter(names, "match")
		invalidInputs := notLengthTwo
		goodEggs := appendParameter(validURLs, "html")
		badEggs := appendParameter(validURLs, "asfdjhow012u")
		testParameters(validInputs, invalidInputs, checks.ResponseMatches{}, t)
		testCheck(goodEggs, badEggs, checks.ResponseMatches{}, t)
	}
}

func TestReponseMatchesInsecure(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping tests that query remote servers in short mode")
	} else {
		validInputs := appendParameter(names, "match")
		invalidInputs := notLengthTwo
		goodEggs := appendParameter(validURLs, "html")
		badEggs := appendParameter(validURLs, "asfdjhow012u")
		testParameters(validInputs, invalidInputs, checks.ResponseMatchesInsecure{}, t)
		testCheck(goodEggs, badEggs, checks.ResponseMatchesInsecure{}, t)
	}
}
