package checks_test

import (
	"testing"
    "github.com/CiscoCloud/distributive/checks"
)

var fileParameters = [][]string{
	{"/proc/net/tcp"},
	{"/bin/sh"},
	{"/proc/filesystems"},
	{"/proc/uptime"},
	{"/proc/cpuinfo"},
}

var dirParameters = [][]string{
	{"/dev"},
	{"/var"},
	{"/tmp"},
	{"/opt"},
	{"/usr"},
	{"/usr/bin"},
}

var symlinkParameters = [][]string{
	{"/bin/sh"},
}

var notPaths = append(notLengthOne,
	[]string{}, []string{`\{{[(`}, []string{"", "", ""}, []string{"fail"},
	[]string{""}, []string{"7"},
)

func TestFile(t *testing.T) {
	t.Parallel()
	validInputs := append(fileParameters, dirParameters...)
	validInputs = append(validInputs, symlinkParameters...)
	invalidInputs := notLengthOne
	goodEggs := fileParameters
	badEggs := append(dirParameters, symlinkParameters...)
	testParameters(validInputs, invalidInputs, checks.File{}, t)
	testCheck(goodEggs, badEggs, checks.File{}, t)
}

func TestDirectory(t *testing.T) {
	t.Parallel()
	validInputs := append(fileParameters, dirParameters...)
	validInputs = append(validInputs, symlinkParameters...)
	invalidInputs := notLengthOne
	goodEggs := dirParameters
	badEggs := append(fileParameters, symlinkParameters...)
	testParameters(validInputs, invalidInputs, checks.Directory{}, t)
	testCheck(goodEggs, badEggs, checks.Directory{}, t)
}

func TestSymlink(t *testing.T) {
	t.Parallel()
	validInputs := append(fileParameters, dirParameters...)
	validInputs = append(validInputs, symlinkParameters...)
	invalidInputs := notLengthOne
	goodEggs := symlinkParameters
	badEggs := append(dirParameters, fileParameters...)
	testParameters(validInputs, invalidInputs, checks.Symlink{}, t)
	testCheck(goodEggs, badEggs, checks.Symlink{}, t)
}

// $1 - algorithm, $2 - check against, $3 - path
func TestChecksum(t *testing.T) {
	t.Parallel()
	validInputs := [][]string{
		{"md5", "d41d8cd98f00b204e9800998ecf8427e", "/dev/null"},
		{"sha1", "da39a3ee5e6b4b0d3255bfef95601890afd80709", "/dev/null"},
		{"sha256", "chksum", "/proc/cpuinfo"},
		{"sha512", "chksum", "/proc/cpuinfo"},
	}
	// generate losers from all files - none of them have that checksum
	invalidInputs := [][]string{{}, {"", "", ""}, {"", ""}}
	invalidInputs = append(invalidInputs, names...)
	// TODO this fails when testing
	//goodEggs := [][]string{validInputs[0], validInputs[1]}
	//badEggs := [][]string{validInputs[2], validInputs[3]}
	testParameters(validInputs, invalidInputs, checks.Checksum{}, t)
	//testCheck(goodEggs, badEggs, Checksum{}, t)
}

func TestFileMatches(t *testing.T) {
	t.Parallel()
	validInputs := appendParameter(fileParameters, "")
	invalidInputs := append(names,
		[][]string{{"", ""}, {}, {"/notfile", "notmatch"}}...)
	invalidInputs = append(notLengthTwo, names...)
	goodEggs := validInputs
	badEggs := [][]string{
		{"/dev/null", "something"}, {"/proc/cpuinfo", "siddharthist"},
	}
	testParameters(validInputs, invalidInputs, checks.FileMatches{}, t)
	testCheck(goodEggs, badEggs, checks.FileMatches{}, t)
}

// $1 - path, $2 - givenMode (-rwxrwxrwx)
func TestPermissions(t *testing.T) {
	t.Parallel()
	valid1 := appendParameter(fileParameters, "----------")
	valid2 := appendParameter(dirParameters, "-rwxrwxrwx")
	valid3 := appendParameter(symlinkParameters, "-r--r--r--")
	validInputs := append(append(valid1, valid2...), valid3...)
	invalid1 := appendParameter(fileParameters, "nonsense")
	invalid2 := appendParameter(dirParameters, "-rrrwwwxxx")
	invalid3 := appendParameter(symlinkParameters, "")
	invalidInputs := append(append(invalid1, invalid2...), invalid3...)
	invalidInputs = append(invalidInputs, names...)
	goodEggs := [][]string{
		{"/dev/null", "-rw-rw-rw-"},
		{"/proc/", "-r-xr-xr-x"},
		{"/bin/", "-rwxr-xr-x"},
	}
	badEggs := appendParameter(fileParameters, "----------")
	testParameters(validInputs, invalidInputs, checks.Permissions{}, t)
	testCheck(goodEggs, badEggs, checks.Permissions{}, t)
}
