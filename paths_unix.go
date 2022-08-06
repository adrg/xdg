//go:build aix || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris
// +build aix dragonfly freebsd js,wasm nacl linux netbsd openbsd solaris

package xdg

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/adrg/xdg/internal/pathutil"
)

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	return "/"
}

func initDirs(bd *BaseDirectories, home string) {
	initBaseDirs(bd, home)
	initUserDirs(bd, home)
}

func initBaseDirs(bd *BaseDirectories, home string) {
	// Initialize standard directories.
	bd.dataHome = xdgPath(envDataHome, filepath.Join(home, ".local", "share"))
	bd.data = xdgPaths(envDataDirs, "/usr/local/share", "/usr/share")
	bd.configHome = xdgPath(envConfigHome, filepath.Join(home, ".config"))
	bd.config = xdgPaths(envConfigDirs, "/etc/xdg")
	bd.stateHome = xdgPath(envStateHome, filepath.Join(home, ".local", "state"))
	bd.cacheHome = xdgPath(envCacheHome, filepath.Join(home, ".cache"))
	bd.runtime = xdgPath(envRuntimeDir, filepath.Join("/run/user", strconv.Itoa(os.Getuid())))

	// Initialize non-standard directories.
	appDirs := []string{
		filepath.Join(bd.dataHome, "applications"),
		filepath.Join(home, ".local/share/applications"),
		"/usr/local/share/applications",
		"/usr/share/applications",
	}

	fontDirs := []string{
		filepath.Join(bd.dataHome, "fonts"),
		filepath.Join(home, ".fonts"),
		filepath.Join(home, ".local/share/fonts"),
		"/usr/local/share/fonts",
		"/usr/share/fonts",
	}

	for _, dir := range bd.data {
		appDirs = append(appDirs, filepath.Join(dir, "applications"))
		fontDirs = append(fontDirs, filepath.Join(dir, "fonts"))
	}

	bd.applications = pathutil.Unique(appDirs, Home)
	bd.fonts = pathutil.Unique(fontDirs, Home)
}

func initUserDirs(bd *BaseDirectories, home string) {
	bd.UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	bd.UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	bd.UserDirs.Documents = xdgPath(envDocumentsDir, filepath.Join(home, "Documents"))
	bd.UserDirs.Music = xdgPath(envMusicDir, filepath.Join(home, "Music"))
	bd.UserDirs.Pictures = xdgPath(envPicturesDir, filepath.Join(home, "Pictures"))
	bd.UserDirs.Videos = xdgPath(envVideosDir, filepath.Join(home, "Videos"))
	bd.UserDirs.Templates = xdgPath(envTemplatesDir, filepath.Join(home, "Templates"))
	bd.UserDirs.PublicShare = xdgPath(envPublicShareDir, filepath.Join(home, "Public"))
}
