//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js,wasm nacl linux netbsd openbsd solaris

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
	require.Equal(t, home, util.ExpandHome("$HOME", home))
	require.Equal(t, filepath.Join(home, "appname"), util.ExpandHome("~/appname", home))
	require.Equal(t, filepath.Join(home, "appname"), util.ExpandHome("$HOME/appname", home))

	require.Equal(t, "", util.ExpandHome("", home))
	require.Equal(t, home, util.ExpandHome(home, ""))
	require.Equal(t, "", util.ExpandHome("", ""))

	require.Equal(t, home, util.ExpandHome(home, home))
	require.Equal(t, "/", util.ExpandHome("~", "/"))
	require.Equal(t, "/", util.ExpandHome("$HOME", "/"))
	require.Equal(t, "/usr/bin", util.ExpandHome("~/bin", "/usr"))
	require.Equal(t, "/usr/bin", util.ExpandHome("$HOME/bin", "/usr"))
}
