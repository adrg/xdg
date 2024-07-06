//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris

package userdirs_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/adrg/xdg/internal/userdirs"
	"github.com/stretchr/testify/require"
)

func TestParseConfigFile(t *testing.T) {
	// Test parsed values.
	f, err := os.CreateTemp("", "test_parse_config_file")
	require.NoError(t, err)

	var tmpFileRemoved bool
	defer func() {
		if !tmpFileRemoved {
			os.Remove(f.Name())
		}
	}()

	_, err = f.Write([]byte(`XDG_DOWNLOAD_DIR="/home/test/Downloads"`))
	require.NoError(t, err)

	err = f.Close()
	require.NoError(t, err)

	dirs := userdirs.ParseConfigFile(f.Name())
	require.NotNil(t, dirs)
	require.Equal(t, "/home/test/Downloads", dirs["XDG_DOWNLOAD_DIR"])

	// Test non-existent file.
	err = os.Remove(f.Name())
	require.NoError(t, err)
	tmpFileRemoved = true

	dirs = userdirs.ParseConfigFile(f.Name())
	require.NotNil(t, dirs)
}

func TestParseConfig(t *testing.T) {
	// Test parsed values.
	home := pathutil.UserHomeDir()

	dirs := userdirs.ParseConfig(strings.NewReader(`
		# This file is written by xdg-user-dirs-update
		# If you want to change or add directories, just edit the line you're
		# interested in. All local changes will be retained on the next run.
		# Format is XDG_xxx_DIR="$HOME/yyy", where yyy is a shell-escaped
		# homedir-relative path, or XDG_xxx_DIR="/yyy", where /yyy is an
		# absolute path. No other format is supported.
		#
		XDG_DESKTOP_DIR="$HOME/Desktop"
		XDG_DOWNLOAD_DIR="$HOME/Downloads"
		XDG_TEMPLATES_DIR="/home/test/Templates"
		XDG_PUBLICSHARE_DIR="~/Public"
		XDG_DOCUMENTS_DIR="$HOME/Documents"
		XDG_MUSIC_DIR="$HOME/Music" # Music user directory
		# XDG_PICTURES_DIR="$HOME/Pictures"
		XDG_VIDEOS_DIR=""

		NON_XDG_DIR="ignore"
		XDG_INVALID_DIR="ignore"
		XDG_DOWNLOAD_DIR
	`))

	require.NotNil(t, dirs)
	require.Equal(t, filepath.Join(home, "Desktop"), dirs["XDG_DESKTOP_DIR"])
	require.Equal(t, filepath.Join(home, "Downloads"), dirs["XDG_DOWNLOAD_DIR"])
	require.Equal(t, "/home/test/Templates", dirs["XDG_TEMPLATES_DIR"])
	require.Equal(t, filepath.Join(home, "Public"), dirs["XDG_PUBLICSHARE_DIR"])
	require.Equal(t, filepath.Join(home, "Documents"), dirs["XDG_DOCUMENTS_DIR"])
	require.Equal(t, filepath.Join(home, "Music"), dirs["XDG_MUSIC_DIR"])
	require.Equal(t, "", dirs["XDG_PICTURES_DIR"])
	require.Equal(t, "", dirs["XDG_VIDEOS_DIR"])

	// Test reader error.
	f, err := os.CreateTemp("", "test_parse_config")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	err = f.Close()
	require.NoError(t, err)

	dirs = userdirs.ParseConfig(f)
	require.NotNil(t, dirs)
}
