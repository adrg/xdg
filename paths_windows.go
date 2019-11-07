package xdg

import (
	"os"
	"path/filepath"
)

func initBaseDirs(home string) {
	appDataDir := os.Getenv("APPDATA")
	if appDataDir == "" {
		appDataDir = filepath.Join(home, "AppData")
	}
	roamingAppDataDir := filepath.Join(appDataDir, "Roaming")

	localAppDataDir := os.Getenv("LOCALAPPDATA")
	if localAppDataDir == "" {
		localAppDataDir = filepath.Join(appDataDir, "Local")
	}

	programDataDir := os.Getenv("PROGRAMDATA")
	if programDataDir == "" {
		if systemDrive := os.Getenv("SystemDrive"); systemDrive != "" {
			programDataDir = filepath.Join(systemDrive, "ProgramData")
		} else {
			programDataDir = home
		}
	}

	baseDirs.DataHome = xdgPath(envDataHome, localAppDataDir)
	baseDirs.Data = xdgPaths(envDataDirs, roamingAppDataDir, programDataDir)
	baseDirs.ConfigHome = xdgPath(envConfigHome, localAppDataDir)
	baseDirs.Config = xdgPaths(envConfigDirs, programDataDir)
	baseDirs.CacheHome = xdgPath(envCacheHome, filepath.Join(localAppDataDir, "cache"))
	baseDirs.Runtime = xdgPath(envRuntimeDir, localAppDataDir)
}

func initUserDirs(home string) {
	publicDir := os.Getenv("PUBLIC")
	if publicDir == "" {
		publicDir = filepath.Join(home, "Public")
	}

	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Templates = xdgPath(envDownloadDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envDownloadDir, publicDir)
	UserDirs.Documents = xdgPath(envDownloadDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envDownloadDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envDownloadDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envDownloadDir, filepath.Join(home, "Videos"))
}
