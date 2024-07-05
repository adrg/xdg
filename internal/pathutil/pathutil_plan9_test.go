//go:build plan9

package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adrg/xdg/internal/pathutil"
)

func TestUserHomeDir(t *testing.T) {
	home := os.Getenv("home")
	defer os.Setenv("home", home)

	require.Equal(t, home, pathutil.UserHomeDir())

	os.Unsetenv("home")
	require.Equal(t, "/", pathutil.UserHomeDir())
}

func TestExpandHome(t *testing.T) {
	home := pathutil.UserHomeDir()

	require.Equal(t, "", pathutil.ExpandHome(""))
	require.Equal(t, home, pathutil.ExpandHome(home))
	require.Equal(t, home, pathutil.ExpandHome("~"))
	require.Equal(t, home, pathutil.ExpandHome("$home"))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("~/appname"))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("$home/appname"))
}

func TestUnique(t *testing.T) {
	home := pathutil.UserHomeDir()

	input := []string{
		"",
		home,
		filepath.Join(home, "foo"),
		"a",
		"~/foo",
		"$home/foo",
		"a",
		"~",
		"$home",
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
	require.Equal(t, home, pathutil.First([]string{"$home"}))
	require.Equal(t, home, pathutil.First([]string{"~"}))
	require.Equal(t, home, pathutil.First([]string{home, ""}))
	require.Equal(t, home, pathutil.First([]string{"", home}))
	require.Equal(t, home, pathutil.First([]string{"$home", ""}))
	require.Equal(t, home, pathutil.First([]string{"", "$home"}))
	require.Equal(t, home, pathutil.First([]string{"~", ""}))
	require.Equal(t, home, pathutil.First([]string{"", "~"}))
	require.Equal(t, "/home/test/foo", pathutil.First([]string{"/home/test/foo", "/home/test/bar"}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{"$home/foo", "$home/bar"}))
	require.Equal(t, filepath.Join(home, "foo"), pathutil.First([]string{"~/foo", "~/bar"}))
}
