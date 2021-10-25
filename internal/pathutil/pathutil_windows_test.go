//go:build windows
// +build windows

package pathutil_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
)

func TestKnownFolderPath(t *testing.T) {
	expected := `C:\ProgramData`
	require.Equal(t, expected, pathutil.KnownFolderPath(windows.FOLDERID_ProgramData, nil, nil))
	require.Equal(t, expected, pathutil.KnownFolderPath(nil, []string{"ProgramData"}, nil))
	require.Equal(t, expected, pathutil.KnownFolderPath(nil, nil, []string{expected}))
	require.Equal(t, "", pathutil.KnownFolderPath(nil, nil, nil))
}

func TestExpandHome(t *testing.T) {
	home := `C:\Users\test`

	require.Equal(t, home, pathutil.ExpandHome(`%USERPROFILE%`, home))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome(`%USERPROFILE%\appname`, home))

	require.Equal(t, "", pathutil.ExpandHome("", home))
	require.Equal(t, home, pathutil.ExpandHome(home, ""))
	require.Equal(t, "", pathutil.ExpandHome("", ""))

	require.Equal(t, home, pathutil.ExpandHome(home, home))
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

	require.EqualValues(t, expected, pathutil.UniquePaths(input, `C:\Users\test`))
}
