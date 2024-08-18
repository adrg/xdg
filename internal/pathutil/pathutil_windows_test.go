//go:build windows

package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
)

func TestUserHomeDir(t *testing.T) {
	home := pathutil.KnownFolder(windows.FOLDERID_Profile, nil, nil)
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}

	require.Equal(t, home, pathutil.UserHomeDir())
}

func TestKnownFolder(t *testing.T) {
	expected := `C:\ProgramData`
	require.Equal(t, expected, pathutil.KnownFolder(windows.FOLDERID_ProgramData, nil, nil))
	require.Equal(t, expected, pathutil.KnownFolder(nil, []string{"ProgramData"}, nil))
	require.Equal(t, expected, pathutil.KnownFolder(nil, nil, []string{expected}))
	require.Equal(t, "", pathutil.KnownFolder(nil, nil, nil))
}

func TestExpandHome(t *testing.T) {
	home := pathutil.UserHomeDir()

	require.Equal(t, "", pathutil.ExpandHome(""))
	require.Equal(t, home, pathutil.ExpandHome(home))
	require.Equal(t, home, pathutil.ExpandHome(`%USERPROFILE%`))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome(`%USERPROFILE%\appname`))
}

func TestUnique(t *testing.T) {
	home := pathutil.UserHomeDir()

	input := []string{
		"",
		home,
		filepath.Join(home, "foo"),
		"a",
		`%USERPROFILE%/foo`,
		`%USERPROFILE%\foo`,
		"a",
	}

	expected := []string{
		home,
		filepath.Join(home, "foo"),
	}

	require.EqualValues(t, expected, pathutil.Unique(input))
}

func TestFirst(t *testing.T) {
	home := pathutil.UserHomeDir()

	require.Equal(t, "", pathutil.First([]string{}))
	require.Equal(t, home, pathutil.First([]string{home}))
	require.Equal(t, home, pathutil.First([]string{"%USERPROFILE%"}))
	require.Equal(t, home, pathutil.First([]string{home, ""}))
	require.Equal(t, home, pathutil.First([]string{"", home}))
	require.Equal(t, home, pathutil.First([]string{"%USERPROFILE%", ""}))
	require.Equal(t, home, pathutil.First([]string{"", "%USERPROFILE%"}))
	require.Equal(t, `C:\Users\foo`, pathutil.First([]string{`C:\Users\foo`, `C:\Users\bar`}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{`%USERPROFILE%/foo`, `%USERPROFILE%/bar`}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{`%USERPROFILE%/foo`, `%USERPROFILE%/bar`}))
}
