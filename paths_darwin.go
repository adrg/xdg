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
	baseDirs.Runtime = xdgPath(envDataHome, filepath.Join(home, "Library", "Application Support"))
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Templates = xdgPath(envDownloadDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envDownloadDir, filepath.Join(home, "Public"))
	UserDirs.Documents = xdgPath(envDownloadDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envDownloadDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envDownloadDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envDownloadDir, filepath.Join(home, "Movies"))
}
