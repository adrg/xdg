package xdg

import (
	"path/filepath"
)

func initDirs(home string) {
	kf := initKnownFolders(home)
	initBaseDirs(home, kf)
	initUserDirs(home, kf)
}

func initBaseDirs(home string, kf *knownFolders) {
	// Initialize standard directories.
	baseDirs.dataHome = xdgPath(envDataHome, kf.localAppData)
	baseDirs.data = xdgPaths(envDataDirs, kf.roamingAppData, kf.programData)
	baseDirs.configHome = xdgPath(envConfigHome, kf.localAppData)
	baseDirs.config = xdgPaths(envConfigDirs, kf.programData)
	baseDirs.stateHome = xdgPath(envStateHome, kf.localAppData)
	baseDirs.cacheHome = xdgPath(envCacheHome, filepath.Join(kf.localAppData, "cache"))
	baseDirs.runtime = xdgPath(envRuntimeDir, kf.localAppData)

	// Initialize non-standard directories.
	baseDirs.applications = []string{
		kf.programs,
		kf.commonPrograms,
	}
	baseDirs.fonts = []string{
		kf.fonts,
		filepath.Join(kf.localAppData, "Microsoft", "Windows", "Fonts"),
	}
}

func initUserDirs(home string, kf *knownFolders) {
	UserDirs.Desktop = xdgPath(envDesktopDir, kf.desktop)
	UserDirs.Download = xdgPath(envDownloadDir, kf.downloads)
	UserDirs.Documents = xdgPath(envDocumentsDir, kf.documents)
	UserDirs.Music = xdgPath(envMusicDir, kf.music)
	UserDirs.Pictures = xdgPath(envPicturesDir, kf.pictures)
	UserDirs.Videos = xdgPath(envVideosDir, kf.videos)
	UserDirs.Templates = xdgPath(envTemplatesDir, kf.templates)
	UserDirs.PublicShare = xdgPath(envPublicShareDir, kf.public)
}
