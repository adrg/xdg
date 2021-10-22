//go:build windows
// +build windows

package xdg_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/require"
)

func TestDefaultBaseDirs(t *testing.T) {
	home := xdg.Home
	systemDrive := `C:\`
	roamingAppData := filepath.Join(home, "AppData", "Roaming")
	localAppData := filepath.Join(home, "AppData", "Local")
	systemRoot := filepath.Join(systemDrive, "Windows")
	programData := filepath.Join(systemDrive, "ProgramData")

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			expected: localAppData,
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			expected: []string{roamingAppData, programData},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			expected: localAppData,
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			expected: []string{programData, roamingAppData},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			expected: localAppData,
			actual:   &xdg.StateHome,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			expected: filepath.Join(localAppData, "cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			expected: localAppData,
			actual:   &xdg.RuntimeDir,
		},
		&envSample{
			name: "XDG_APPLICATION_DIRS",
			expected: []string{
				filepath.Join(roamingAppData, "Microsoft", "Windows", "Start Menu", "Programs"),
				filepath.Join(programData, "Microsoft", "Windows", "Start Menu", "Programs"),
			},
			actual: &xdg.ApplicationDirs,
		},
		&envSample{
			name: "XDG_FONT_DIRS",
			expected: []string{
				filepath.Join(systemRoot, "Fonts"),
				filepath.Join(localAppData, "Microsoft", "Windows", "Fonts"),
			},
			actual: &xdg.FontDirs,
		},
	)
}

func TestCustomBaseDirs(t *testing.T) {
	home := xdg.Home
	appData := filepath.Join(home, "Appdata")
	localAppData := filepath.Join(appData, "Local")
	programData := filepath.Join(home, "ProgramData")

	require.NoError(t, os.Setenv("APPDATA", appData))
	require.NoError(t, os.Setenv("LOCALAPPDATA", localAppData))
	require.NoError(t, os.Setenv("PROGRAMDATA", programData))

	testDirs(t,
		&envSample{
			name:     "XDG_DATA_HOME",
			value:    filepath.Join(localAppData, "Data"),
			expected: filepath.Join(localAppData, "Data"),
			actual:   &xdg.DataHome,
		},
		&envSample{
			name:     "XDG_DATA_DIRS",
			value:    fmt.Sprintf("%s;%s", filepath.Join(localAppData, "Data"), filepath.Join(appData, "Data")),
			expected: []string{filepath.Join(localAppData, "Data"), filepath.Join(appData, "Data")},
			actual:   &xdg.DataDirs,
		},
		&envSample{
			name:     "XDG_CONFIG_HOME",
			value:    filepath.Join(localAppData, "Config"),
			expected: filepath.Join(localAppData, "Config"),
			actual:   &xdg.ConfigHome,
		},
		&envSample{
			name:     "XDG_CONFIG_DIRS",
			value:    fmt.Sprintf("%s;%s", filepath.Join(localAppData, "Config"), filepath.Join(appData, "Config")),
			expected: []string{filepath.Join(localAppData, "Config"), filepath.Join(appData, "Config")},
			actual:   &xdg.ConfigDirs,
		},
		&envSample{
			name:     "XDG_STATE_HOME",
			value:    filepath.Join(programData, "State"),
			expected: filepath.Join(programData, "State"),
			actual:   &xdg.StateHome,
		},
		&envSample{
			name:     "XDG_CACHE_HOME",
			value:    filepath.Join(programData, "Cache"),
			expected: filepath.Join(programData, "Cache"),
			actual:   &xdg.CacheHome,
		},
		&envSample{
			name:     "XDG_RUNTIME_DIR",
			value:    filepath.Join(programData, "Runtime"),
			expected: filepath.Join(programData, "Runtime"),
			actual:   &xdg.RuntimeDir,
		},
	)
}

func TestDefaultUserDirs(t *testing.T) {
	home := xdg.Home
	appData := filepath.Join(home, "AppData")
	roamingAppData := filepath.Join(appData, "Roaming")
	usersDir := `C:\Users`

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
			expected: filepath.Join(roamingAppData, "Microsoft", "Windows", "Templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			expected: filepath.Join(usersDir, "Public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}

func TestCustomUserDirs(t *testing.T) {
	home := xdg.Home

	testDirs(t,
		&envSample{
			name:     "XDG_DESKTOP_DIR",
			value:    filepath.Join(home, "Files/Desktop"),
			expected: filepath.Join(home, "Files/Desktop"),
			actual:   &xdg.UserDirs.Desktop,
		},
		&envSample{
			name:     "XDG_DOWNLOAD_DIR",
			value:    filepath.Join(home, "Files/Downloads"),
			expected: filepath.Join(home, "Files/Downloads"),
			actual:   &xdg.UserDirs.Download,
		},
		&envSample{
			name:     "XDG_DOCUMENTS_DIR",
			value:    filepath.Join(home, "Files/Documents"),
			expected: filepath.Join(home, "Files/Documents"),
			actual:   &xdg.UserDirs.Documents,
		},
		&envSample{
			name:     "XDG_MUSIC_DIR",
			value:    filepath.Join(home, "Files/Music"),
			expected: filepath.Join(home, "Files/Music"),
			actual:   &xdg.UserDirs.Music,
		},
		&envSample{
			name:     "XDG_PICTURES_DIR",
			value:    filepath.Join(home, "Files/Pictures"),
			expected: filepath.Join(home, "Files/Pictures"),
			actual:   &xdg.UserDirs.Pictures,
		},
		&envSample{
			name:     "XDG_VIDEOS_DIR",
			value:    filepath.Join(home, "Files/Videos"),
			expected: filepath.Join(home, "Files/Videos"),
			actual:   &xdg.UserDirs.Videos,
		},
		&envSample{
			name:     "XDG_TEMPLATES_DIR",
			value:    filepath.Join(home, "Files/Templates"),
			expected: filepath.Join(home, "Files/Templates"),
			actual:   &xdg.UserDirs.Templates,
		},
		&envSample{
			name:     "XDG_PUBLICSHARE_DIR",
			value:    filepath.Join(home, "Files/Public"),
			expected: filepath.Join(home, "Files/Public"),
			actual:   &xdg.UserDirs.PublicShare,
		},
	)
}
