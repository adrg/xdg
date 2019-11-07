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
	baseDirs.Runtime = xdgPath(envRuntimeDir, filepath.Join(os.TempDir(), strconv.Itoa(os.Getuid())))
}

func initUserDirs(home string) {
	UserDirs.Desktop = xdgPath(envDesktopDir, filepath.Join(home, "Desktop"))
	UserDirs.Download = xdgPath(envDownloadDir, filepath.Join(home, "Downloads"))
	UserDirs.Templates = xdgPath(envDownloadDir, filepath.Join(home, "Templates"))
	UserDirs.PublicShare = xdgPath(envDownloadDir, filepath.Join(home, "Public"))
	UserDirs.Documents = xdgPath(envDownloadDir, filepath.Join(home, "Documents"))
	UserDirs.Music = xdgPath(envDownloadDir, filepath.Join(home, "Music"))
	UserDirs.Pictures = xdgPath(envDownloadDir, filepath.Join(home, "Pictures"))
	UserDirs.Videos = xdgPath(envDownloadDir, filepath.Join(home, "Videos"))
}
