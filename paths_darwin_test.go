// +build darwin

package xdg_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
)

func TestDefaultBaseDirs(t *testing.T) {
	home := xdg.Home
	homeAppSupport := filepath.Join(home, "Library", "Application Support")
	rootAppSupport := "/Library/Application Support"

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			expected: homeAppSupport,
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			expected: []string{rootAppSupport},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			expected: homeAppSupport,
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name: "XDG_CONFIG_DIRS",
			expected: []string{
				filepath.Join(home, "Library", "Preferences"),
				rootAppSupport,
				"/Library/Preferences",
			},
			actual: &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(home, "Library", "Caches"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			expected: homeAppSupport,
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			expected: homeAppSupport,
			actual:   &xdg.StateHome,
		},
		&envSample{
			name: "XDG_APPLICATION_DIRS",
			expected: []string{
				"/Applications",
			},
			actual: &xdg.ApplicationDirs,
		},
		&envSample{
			name: "XDG_FONT_DIRS",
			expected: []string{
				filepath.Join(home, "Library/Fonts"),
				"/Library/Fonts",
				"/System/Library/Fonts",
				"/Network/Library/Fonts",
			},
			actual: &xdg.FontDirs,
		},
	)
}

func TestCustomBaseDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    "~/Library/data",
			expected: filepath.Join(home, "Library/data"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			value:    "~/Library/data:/Library/Application Support",
			expected: []string{filepath.Join(home, "Library/data"), "/Library/Application Support"},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			value:    "~/Library/config",
			expected: filepath.Join(home, "Library/config"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			value:    "~/Library/config:/Library/Preferences",
			expected: []string{filepath.Join(home, "Library/config"), "/Library/Preferences"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    "~/Library/cache",
			expected: filepath.Join(home, "Library/cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			value:    "~/Library/runtime",
			expected: filepath.Join(home, "Library/runtime"),
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			value:    "~/Library/state",
			expected: filepath.Join(home, "Library/state"),
			actual:   &xdg.StateHome,
		},
	)
}

func TestDefaultUserDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			expected: filepath.Join(home, "Desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			expected: filepath.Join(home, "Downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			expected: filepath.Join(home, "Documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			expected: filepath.Join(home, "Music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			expected: filepath.Join(home, "Pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			expected: filepath.Join(home, "Movies"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			expected: filepath.Join(home, "Templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			expected: filepath.Join(home, "Public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}

func TestCustomUserDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			value:    "$HOME/Library/Desktop",
			expected: filepath.Join(home, "Library/Desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			value:    "$HOME/Library/Downloads",
			expected: filepath.Join(home, "Library/Downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			value:    "$HOME/Library/Documents",
			expected: filepath.Join(home, "Library/Documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			value:    "$HOME/Library/Music",
			expected: filepath.Join(home, "Library/Music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			value:    "$HOME/Library/Pictures",
			expected: filepath.Join(home, "Library/Pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			value:    "$HOME/Library/Movies",
			expected: filepath.Join(home, "Library/Movies"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			value:    "$HOME/Library/Templates",
			expected: filepath.Join(home, "Library/Templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			value:    "$HOME/Library/Public",
			expected: filepath.Join(home, "Library/Public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}
