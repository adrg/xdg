//go:build aix || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris

package xdg

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/adrg/xdg/internal/userdirs"
)

func initDirs(home string) {
	initBaseDirs(home)
	initUserDirs(home, baseDirs.configHome)
}

func initBaseDirs(home string) {
	// Initialize standard directories.
	baseDirs.dataHome = xdgPath(envDataHome, filepath.Join(home, ".local", "share"))
	baseDirs.data = xdgPaths(envDataDirs, "/usr/local/share", "/usr/share")
	baseDirs.configHome = xdgPath(envConfigHome, filepath.Join(home, ".config"))
	baseDirs.config = xdgPaths(envConfigDirs, "/etc/xdg")
	baseDirs.stateHome = xdgPath(envStateHome, filepath.Join(home, ".local", "state"))
	baseDirs.cacheHome = xdgPath(envCacheHome, filepath.Join(home, ".cache"))
	baseDirs.runtime = xdgPath(envRuntimeDir, filepath.Join("/run/user", strconv.Itoa(os.Getuid())))

	// Initialize non-standard directories.
	appDirs := []string{
		filepath.Join(baseDirs.dataHome, "applications"),
		filepath.Join(home, ".local/share/applications"),
		"/usr/local/share/applications",
		"/usr/share/applications",
	}

	fontDirs := []string{
		filepath.Join(baseDirs.dataHome, "fonts"),
		filepath.Join(home, ".fonts"),
		filepath.Join(home, ".local/share/fonts"),
		"/usr/local/share/fonts",
		"/usr/share/fonts",
	}

	for _, dir := range baseDirs.data {
		appDirs = append(appDirs, filepath.Join(dir, "applications"))
		fontDirs = append(fontDirs, filepath.Join(dir, "fonts"))
	}

	baseDirs.applications = pathutil.Unique(appDirs)
	baseDirs.fonts = pathutil.Unique(fontDirs)
}

func initUserDirs(home, configHome string) {
	userDirsMap := userdirs.ParseConfigFile(filepath.Join(configHome, "user-dirs.dirs"))

	UserDirs.Desktop = xdgPath(userdirs.EnvDesktopDir, userDirsMap[userdirs.EnvDesktopDir], filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(userdirs.EnvDownloadDir, userDirsMap[userdirs.EnvDownloadDir], filepath.Join(home, "Downloads"))
	UserDirs.Documents = xdgPath(userdirs.EnvDocumentsDir, userDirsMap[userdirs.EnvDocumentsDir], filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(userdirs.EnvMusicDir, userDirsMap[userdirs.EnvMusicDir], filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(userdirs.EnvPicturesDir, userDirsMap[userdirs.EnvPicturesDir], filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(userdirs.EnvVideosDir, userDirsMap[userdirs.EnvVideosDir], filepath.Join(home, "Videos"))
	UserDirs.Templates = xdgPath(userdirs.EnvTemplatesDir, userDirsMap[userdirs.EnvTemplatesDir], filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(userdirs.EnvPublicShareDir, userDirsMap[userdirs.EnvPublicShareDir], filepath.Join(home, "Public"))
}
