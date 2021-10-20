package xdg

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"
)

func initBaseDirs(home string) {
	roamingAppDataDir := os.Getenv("APPDATA")
	if roamingAppDataDir == "" {
		roamingAppDataDir := filepath.Join(home, "Roaming")
	}

	localAppDataDir := os.Getenv("LOCALAPPDATA")
	if localAppDataDir == "" {
		localAppDataDir = filepath.Join(home, "Local")
	}

	programDataDir := os.Getenv("PROGRAMDATA")
	if programDataDir == "" {
		if systemDrive := os.Getenv("SystemDrive"); systemDrive != "" {
			programDataDir = filepath.Join(systemDrive, "ProgramData")
		} else {
			programDataDir = home
		}
	}

	winDir := os.Getenv("windir")
	if winDir == "" {
		if winDir = os.Getenv("SystemRoot"); winDir == "" {
			winDir = home
		}
	}

	// Initialize base directories.
	baseDirs.dataHome = xdgPath(envDataHome, localAppDataDir)
	baseDirs.data = xdgPaths(envDataDirs, roamingAppDataDir, programDataDir)
	baseDirs.configHome = xdgPath(envConfigHome, localAppDataDir)
	baseDirs.config = xdgPaths(envConfigDirs, programDataDir)
	baseDirs.stateHome = xdgPath(envStateHome, localAppDataDir)
	baseDirs.cacheHome = xdgPath(envCacheHome, filepath.Join(localAppDataDir, "cache"))
	baseDirs.runtime = xdgPath(envRuntimeDir, localAppDataDir)

	// Initialize non-standard directories.
	baseDirs.applications = []string{
		filepath.Join(roamingAppDataDir, "Microsoft", "Windows", "Start Menu", "Programs"),
	}
	baseDirs.fonts = []string{
		filepath.Join(winDir, "Fonts"),
		filepath.Join(localAppDataDir, "Microsoft", "Windows", "Fonts"),
	}
}

func initUserDirs(home string) {
	publicDir := os.Getenv("PUBLIC")
	if publicDir == "" {
		publicDir = filepath.Join(home, "Public")
	}

	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Documents = xdgPath(envDocumentsDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envMusicDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envPicturesDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envVideosDir, filepath.Join(home, "Videos"))
	UserDirs.Templates = xdgPath(envTemplatesDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envPublicShareDir, publicDir)
}
