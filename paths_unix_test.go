// +build aix dragonfly freebsd linux nacl netbsd openbsd solaris

package xdg_test

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/adrg/xdg"
)

func TestDefaultBaseDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			expected: filepath.Join(home, ".local/share"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			expected: []string{"/usr/local/share", "/usr/share"},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			expected: filepath.Join(home, ".config"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			expected: []string{"/etc/xdg"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(home, ".cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			expected: filepath.Join(os.TempDir(), strconv.Itoa(os.Getuid())),
			actual:   &xdg.RuntimeDir,
		},
	)
}

func TestCustomBaseDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    "~/.data/home",
			expected: filepath.Join(home, ".data/home"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			value:    "~/.data/dirs:/usr/share",
			expected: []string{filepath.Join(home, ".data/dirs"), "/usr/share"},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			value:    "~/.config/home",
			expected: filepath.Join(home, ".config/home"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			value:    "~/.config/dirs:/etc/xdg",
			expected: []string{filepath.Join(home, ".config/dirs"), "/etc/xdg"},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    "~/.cache/home",
			expected: filepath.Join(home, ".cache/home"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			value:    filepath.Join(os.TempDir(), "runtime"),
			expected: filepath.Join(os.TempDir(), "runtime"),
			actual:   &xdg.RuntimeDir,
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
			expected: filepath.Join(home, "Videos"),
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
			value:    "$HOME/files/desktop",
			expected: filepath.Join(home, "files/desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			value:    "$HOME/files/downloads",
			expected: filepath.Join(home, "files/downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			value:    "$HOME/files/documents",
			expected: filepath.Join(home, "files/documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			value:    "$HOME/files/music",
			expected: filepath.Join(home, "files/music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			value:    "$HOME/files/pictures",
			expected: filepath.Join(home, "files/pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			value:    "$HOME/files/videos",
			expected: filepath.Join(home, "files/videos"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			value:    "$HOME/files/templates",
			expected: filepath.Join(home, "files/templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			value:    "$HOME/files/public",
			expected: filepath.Join(home, "files/public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}
