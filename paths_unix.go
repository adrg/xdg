// +build aix dragonfly freebsd linux nacl netbsd openbsd solaris

package xdg

import (
	"os"
	"path/filepath"
	"strconv"
)

func initBaseDirs(home string) {
	baseDirs.DataHome = xdgPath(envDataHome, filepath.Join(home, ".local", "share"))
	baseDirs.Data = xdgPaths(envDataDirs, "/usr/local/share", "/usr/share")
	baseDirs.ConfigHome = xdgPath(envConfigHome, filepath.Join(home, ".config"))
	baseDirs.Config = xdgPaths(envConfigDirs, "/etc/xdg")
	baseDirs.CacheHome = xdgPath(envCacheHome, filepath.Join(home, ".cache"))
	baseDirs.Runtime = xdgPath(envRuntimeDir, filepath.Join("/run/user", strconv.Itoa(os.Getuid())))
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Documents = xdgPath(envDocumentsDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envMusicDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envPicturesDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envVideosDir, filepath.Join(home, "Videos"))
	UserDirs.Templates = xdgPath(envTemplatesDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envPublicShareDir, filepath.Join(home, "Public"))
}
