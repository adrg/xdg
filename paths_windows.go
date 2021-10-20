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

	UserDirs.Desktop = xdgPath(envDesktopDir,
		kfPath(windows.FOLDERID_Desktop, filepath.Join(home, "Desktop")))
	UserDirs.Download = xdgPath(envDownloadDir,
		kfPath(windows.FOLDERID_Downloads, filepath.Join(home, "Downloads")))
	UserDirs.Documents = xdgPath(envDocumentsDir,
		kfPath(windows.FOLDERID_Documents, filepath.Join(home, "Documents")))
	UserDirs.Music = xdgPath(envMusicDir,
		kfPath(windows.FOLDERID_Music, filepath.Join(home, "Music")))
	UserDirs.Pictures = xdgPath(envPicturesDir,
		kfPath(windows.FOLDERID_Pictures, filepath.Join(home, "Pictures")))
	UserDirs.Videos = xdgPath(envVideosDir,
		kfPath(windows.FOLDERID_Videos, filepath.Join(home, "Videos")))
	UserDirs.Templates = xdgPath(envTemplatesDir,
		kfPath(windows.FOLDERID_Templates, filepath.Join(home, "Templates")))
	UserDirs.PublicShare = xdgPath(envPublicShareDir,
		kfPath(windows.FOLDERID_Public, publicDir))
}

func kfPath(folderID *KNOWNFOLDERID, defaultPath string) string {
	knownPath, _ := windows.KnownFolderPath(folderID, KF_FLAG_DEFAULT)
	if knownPath = strings.TrimSpace(knownPath); knownPath == "" {
		return defaultPath
	}

	return knownPath
}
