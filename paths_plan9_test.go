// +build plan9

package xdg_test

import (
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
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
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(home, "lib", ".cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			expected: "/tmp",
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			expected: filepath.Join(home, "lib", "state"),
			actual:   &xdg.StateHome,
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
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    "$home/lib/data",
			expected: filepath.Join(home, "lib", "data"),
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
			value:    "$home/lib/config",
			expected: filepath.Join(home, "lib", "config"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			value:    "/lib/config",
			expected: []string{"/lib/config"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    "$home/lib",
			expected: filepath.Join(home, "lib"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			value:    "$home/lib/runtime",
			expected: filepath.Join(home, "lib", "runtime"),
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			value:    "$home/lib",
			expected: filepath.Join(home, "lib"),
			actual:   &xdg.StateHome,
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
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			value:    "$home/lib/desktop",
			expected: filepath.Join(home, "lib", "desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			value:    "$home/lib/downloads",
			expected: filepath.Join(home, "lib", "downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			value:    "$home/lib/documents",
			expected: filepath.Join(home, "lib", "documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			value:    "$home/lib/music",
			expected: filepath.Join(home, "lib", "music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			value:    "$home/lib/pictures",
			expected: filepath.Join(home, "lib", "pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			value:    "$home/lib/videos",
			expected: filepath.Join(home, "lib", "videos"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			value:    "$home/lib/templates",
			expected: filepath.Join(home, "lib", "templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			value:    "$home/lib/public",
			expected: filepath.Join(home, "lib", "public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}
