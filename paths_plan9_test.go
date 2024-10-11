//go:build plan9

package xdg_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/require"
)

func TestDefaultBaseDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			expected: filepath.Join(home, "lib"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			expected: []string{"/lib"},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			expected: filepath.Join(home, "lib"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			expected: []string{"/lib"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			expected: filepath.Join(home, "lib", "state"),
			actual:   &xdg.StateHome,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(home, "lib", "cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			expected: "/tmp",
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_BIN_HOME",
			expected: filepath.Join(home, "bin"),
			actual:   &xdg.BinHome,
		},
		&envSample{
			name: "XDG_APPLICATION_DIRS",
			expected: []string{
				filepath.Join(home, "bin"),
				"/bin",
			},
			actual: &xdg.ApplicationDirs,
		},
		&envSample{
			name: "XDG_FONT_DIRS",
			expected: []string{
				filepath.Join(home, "lib", "font"),
				"/lib/font",
			},
			actual: &xdg.FontDirs,
		},
	)
}

func TestCustomBaseDirs(t *testing.T) {
	homeLib := filepath.Join(xdg.Home, "lib")

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    filepath.Join(homeLib, "data"),
			expected: filepath.Join(homeLib, "data"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			value:    "/lib/data",
			expected: []string{"/lib/data"},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			value:    filepath.Join(homeLib, "config"),
			expected: filepath.Join(homeLib, "config"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			value:    "/lib/config",
			expected: []string{"/lib/config"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			value:    homeLib,
			expected: homeLib,
			actual:   &xdg.StateHome,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    homeLib,
			expected: homeLib,
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			value:    filepath.Join(homeLib, "runtime"),
			expected: filepath.Join(homeLib, "runtime"),
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_BIN_HOME",
			value:    filepath.Join(homeLib, "bin"),
			expected: filepath.Join(homeLib, "bin"),
			actual:   &xdg.BinHome,
		},
	)
}

func TestDefaultUserDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			expected: filepath.Join(home, "desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			expected: filepath.Join(home, "downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			expected: filepath.Join(home, "documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			expected: filepath.Join(home, "music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			expected: filepath.Join(home, "pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			expected: filepath.Join(home, "videos"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			expected: filepath.Join(home, "templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			expected: filepath.Join(home, "public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}

func TestCustomUserDirs(t *testing.T) {
	homeLib := filepath.Join(xdg.Home, "lib")

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			value:    filepath.Join(homeLib, "desktop"),
			expected: filepath.Join(homeLib, "desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			value:    filepath.Join(homeLib, "downloads"),
			expected: filepath.Join(homeLib, "downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			value:    filepath.Join(homeLib, "documents"),
			expected: filepath.Join(homeLib, "documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			value:    filepath.Join(homeLib, "music"),
			expected: filepath.Join(homeLib, "music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			value:    filepath.Join(homeLib, "pictures"),
			expected: filepath.Join(homeLib, "pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			value:    filepath.Join(homeLib, "videos"),
			expected: filepath.Join(homeLib, "videos"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			value:    filepath.Join(homeLib, "templates"),
			expected: filepath.Join(homeLib, "templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			value:    filepath.Join(homeLib, "public"),
			expected: filepath.Join(homeLib, "public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}

func TestHomeNotSet(t *testing.T) {
	envHomeVar := "home"
	envHomeVal := os.Getenv(envHomeVar)
	require.NoError(t, os.Unsetenv(envHomeVar))

	xdg.Reload()
	require.Equal(t, "/", xdg.Home)

	require.NoError(t, os.Setenv(envHomeVar, envHomeVal))
	xdg.Reload()
	require.Equal(t, envHomeVal, xdg.Home)
}
