//go:build plan9

package pathutil_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/stretchr/testify/require"
)

func TestExpandHome(t *testing.T) {
	home := "/home/test"

	require.Equal(t, home, pathutil.ExpandHome("~", home))
	require.Equal(t, home, pathutil.ExpandHome("$home", home))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("~/appname", home))
	require.Equal(t, filepath.Join(home, "appname"), pathutil.ExpandHome("$home/appname", home))

	require.Equal(t, "", pathutil.ExpandHome("", home))
	require.Equal(t, home, pathutil.ExpandHome(home, ""))
	require.Equal(t, "", pathutil.ExpandHome("", ""))

	require.Equal(t, home, pathutil.ExpandHome(home, home))
	require.Equal(t, "/", pathutil.ExpandHome("~", "/"))
	require.Equal(t, "/", pathutil.ExpandHome("$home", "/"))
	require.Equal(t, "/usr/bin", pathutil.ExpandHome("~/bin", "/usr"))
	require.Equal(t, "/usr/bin", pathutil.ExpandHome("$home/bin", "/usr"))
}

func TestUnique(t *testing.T) {
	input := []string{
		"",
		"/home",
		"/home/test",
		"a",
		"~/appname",
		"$home/appname",
		"a",
		"/home",
	}

	expected := []string{
		"/home",
		"/home/test",
		"/home/test/appname",
	}

	require.EqualValues(t, expected, pathutil.Unique(input, "/home/test"))
}
