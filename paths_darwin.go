package xdg

import (
	"path/filepath"
)

func initBaseDirs(home string) {
	baseDirs.DataHome = xdgPath(envDataHome, filepath.Join(home, "Library", "Application Support"))
	baseDirs.Data = xdgPaths(envDataDirs, "/Library/Application Support")
	baseDirs.ConfigHome = xdgPath(envConfigHome, filepath.Join(home, "Library", "Preferences"))
	baseDirs.Config = xdgPaths(envConfigDirs)
	baseDirs.CacheHome = xdgPath(envCacheHome, filepath.Join(home, "Library", "Caches"))
	baseDirs.Runtime = xdgPath(envRuntimeDir, filepath.Join(home, "Library", "Application Support"))
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Documents = xdgPath(envDocumentsDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envMusicDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envPicturesDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envVideosDir, filepath.Join(home, "Movies"))
	UserDirs.Templates = xdgPath(envTemplatesDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envPublicShareDir, filepath.Join(home, "Public"))
}
