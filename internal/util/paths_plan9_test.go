//go:build plan9
// +build plan9

package util_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/util"
	"github.com/stretchr/testify/require"
)

func TestExpandHome(t *testing.T) {
	home := "/home/test"

	require.Equal(t, home, util.ExpandHome("~", home))
	require.Equal(t, home, util.ExpandHome("$home", home))
	require.Equal(t, filepath.Join(home, "appname"), util.ExpandHome("~/appname", home))
	require.Equal(t, filepath.Join(home, "appname"), util.ExpandHome("$home/appname", home))

	require.Equal(t, "", util.ExpandHome("", home))
	require.Equal(t, home, util.ExpandHome(home, ""))
	require.Equal(t, "", util.ExpandHome("", ""))

	require.Equal(t, home, util.ExpandHome(home, home))
	require.Equal(t, "/", util.ExpandHome("~", "/"))
	require.Equal(t, "/", util.ExpandHome("$home", "/"))
	require.Equal(t, "/usr/bin", util.ExpandHome("~/bin", "/usr"))
	require.Equal(t, "/usr/bin", util.ExpandHome("$home/bin", "/usr"))
}

func TestUniquePaths(t *testing.T) {
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

	require.EqualValues(t, expected, util.UniquePaths(input, "/home/test"))
}
