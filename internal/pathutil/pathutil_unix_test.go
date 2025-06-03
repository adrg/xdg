//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris

package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adrg/xdg/internal/pathutil"
)

func TestUserHomeDir(t *testing.T) {
	home := os.Getenv("HOME")

	require.Equal(t, home, pathutil.UserHomeDir())
	require.NoError(t, os.Unsetenv("HOME"))
	require.Equal(t, "/", pathutil.UserHomeDir())
	require.NoError(t, os.Setenv("HOME", home))
}

func TestExpandHome(t *testing.T) {
	home := pathutil.UserHomeDir()

	require.Equal(t, "", pathutil.ExpandHome(""))
	require.Equal(t, home, pathutil.ExpandHome(home))
	require.Equal(t, home, pathutil.ExpandHome("~"))
	require.Equal(t, home, pathutil.ExpandHome("$HOME"))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("~/appname"))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("$HOME/appname"))
}

func TestUnique(t *testing.T) {
	home := pathutil.UserHomeDir()

	input := []string{
		"",
		home,
		filepath.Join(home, "foo"),
		"a",
		"~/foo",
		"$HOME/foo",
		"a",
		"~",
		"$HOME",
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
	require.Equal(t, home, pathutil.First([]string{"$HOME"}))
	require.Equal(t, home, pathutil.First([]string{"~"}))
	require.Equal(t, home, pathutil.First([]string{home, ""}))
	require.Equal(t, home, pathutil.First([]string{"", home}))
	require.Equal(t, home, pathutil.First([]string{"$HOME", ""}))
	require.Equal(t, home, pathutil.First([]string{"", "$HOME"}))
	require.Equal(t, home, pathutil.First([]string{"~", ""}))
	require.Equal(t, home, pathutil.First([]string{"", "~"}))
	require.Equal(t, "/home/test/foo", pathutil.First([]string{"/home/test/foo", "/home/test/bar"}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{"$HOME/foo", "$HOME/bar"}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{"~/foo", "~/bar"}))
}
