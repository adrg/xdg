package xdg

import (
	"path/filepath"

	"github.com/adrg/xdg/internal/userdirs"
)

func initDirs(home string) {
	initBaseDirs(home)
	initUserDirs(home)
}

func initBaseDirs(home string) {
	homeAppSupport := filepath.Join(home, "Library", "Application Support")
	rootAppSupport := "/Library/Application Support"

	// Initialize standard directories.
	baseDirs.dataHome = xdgPath(envDataHome, homeAppSupport)
	baseDirs.data = xdgPaths(envDataDirs, rootAppSupport)
	baseDirs.configHome = xdgPath(envConfigHome, homeAppSupport)
	baseDirs.config = xdgPaths(envConfigDirs,
		filepath.Join(home, "Library", "Preferences"),
		rootAppSupport,
		"/Library/Preferences",
	)
	baseDirs.stateHome = xdgPath(envStateHome, homeAppSupport)
	baseDirs.cacheHome = xdgPath(envCacheHome, filepath.Join(home, "Library", "Caches"))
	baseDirs.runtime = xdgPath(envRuntimeDir, homeAppSupport)

	// Initialize non-standard directories.
	baseDirs.applications = []string{
		"/Applications",
	}

	baseDirs.fonts = []string{
		filepath.Join(home, "Library/Fonts"),
		"/Library/Fonts",
		"/System/Library/Fonts",
		"/Network/Library/Fonts",
	}
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(userdirs.EnvDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(userdirs.EnvDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Documents = xdgPath(userdirs.EnvDocumentsDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(userdirs.EnvMusicDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(userdirs.EnvPicturesDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(userdirs.EnvVideosDir, filepath.Join(home, "Movies"))
	UserDirs.Templates = xdgPath(userdirs.EnvTemplatesDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(userdirs.EnvPublicShareDir, filepath.Join(home, "Public"))
}
