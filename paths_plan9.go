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
	homeLibDir := filepath.Join(home, "lib")
	rootLibDir := "/lib"

	// Initialize standard directories.
	baseDirs.dataHome = xdgPath(envDataHome, homeLibDir)
	baseDirs.data = xdgPaths(envDataDirs, rootLibDir)
	baseDirs.configHome = xdgPath(envConfigHome, homeLibDir)
	baseDirs.config = xdgPaths(envConfigDirs, rootLibDir)
	baseDirs.stateHome = xdgPath(envStateHome, filepath.Join(homeLibDir, "state"))
	baseDirs.cacheHome = xdgPath(envCacheHome, filepath.Join(homeLibDir, "cache"))
	baseDirs.runtime = xdgPath(envRuntimeDir, "/tmp")

	// Initialize non-standard directories.
	baseDirs.applications = []string{
		filepath.Join(home, "bin"),
		"/bin",
	}

	baseDirs.fonts = []string{
		filepath.Join(homeLibDir, "font"),
		"/lib/font",
	}
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(userdirs.EnvDesktopDir, filepath.Join(home, "desktop"))
	UserDirs.Download = xdgPath(userdirs.EnvDownloadDir, filepath.Join(home, "downloads"))
	UserDirs.Documents = xdgPath(userdirs.EnvDocumentsDir, filepath.Join(home, "documents"))
	UserDirs.Music = xdgPath(userdirs.EnvMusicDir, filepath.Join(home, "music"))
	UserDirs.Pictures = xdgPath(userdirs.EnvPicturesDir, filepath.Join(home, "pictures"))
	UserDirs.Videos = xdgPath(userdirs.EnvVideosDir, filepath.Join(home, "videos"))
	UserDirs.Templates = xdgPath(userdirs.EnvTemplatesDir, filepath.Join(home, "templates"))
	UserDirs.PublicShare = xdgPath(userdirs.EnvPublicShareDir, filepath.Join(home, "public"))
}
