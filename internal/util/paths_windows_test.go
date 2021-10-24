//go:build windows
// +build windows

package util_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
)

func TestKnownFolderPath(t *testing.T) {
	expected := `C:\ProgramData`
	require.Equal(t, expected, util.KnownFolderPath(windows.FOLDERID_ProgramData, nil, nil))
	require.Equal(t, expected, util.KnownFolderPath(nil, []string{"ProgramData"}, nil))
	require.Equal(t, expected, util.KnownFolderPath(nil, nil, []string{expected}))
	require.Equal(t, "", util.KnownFolderPath(nil, nil, nil))
}

func TestExpandHome(t *testing.T) {
	home := `C:\Users\test`

	require.Equal(t, home, util.ExpandHome(`%USERPROFILE%`, home))
	require.Equal(t, filepath.Join(home, "appname"), util.ExpandHome(`%USERPROFILE%\appname`, home))

	require.Equal(t, "", util.ExpandHome("", home))
	require.Equal(t, home, util.ExpandHome(home, ""))
	require.Equal(t, "", util.ExpandHome("", ""))

	require.Equal(t, home, util.ExpandHome(home, home))
}

func TestUniquePaths(t *testing.T) {
	input := []string{
		"",
		`C:\Users`,
		`C:\Users\test`,
		"a",
		`C:\Users\test\appname`,
		`%USERPROFILE%/appname`,
		"a",
		`C:\Users`,
	}

	expected := []string{
		`C:\Users`,
		`C:\Users\test`,
		`C:\Users\test\appname`,
	}

	require.EqualValues(t, expected, util.UniquePaths(input, `C:\Users\test`))
}
