/*
Package xdg provides an implementation of the XDG Base Directory
Specification. The specification defines a set of standard paths for storing
application files including data and configuration files. For portability and
flexibility reasons, applications should use the XDG defined locations instead
of hardcoding paths.

	For more information about the XDG Base Directory Specification see:
	http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
*/
package xdg

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	// Home contains the path of the user's home directory. This directory
	// is defined by the environment variable $HOME.
	Home string

	// DataHome defines the base directory relative to which user-specific
	// data files should be stored. This directory is defined by the
	// environment variable $XDG_DATA_HOME. If this variable is not set,
	// a default equal to $HOME/.local/share should be used.
	DataHome string

	// ConfigHome defines the base directory relative to which user-specific
	// configuration files should be written. This directory is defined by
	// the environment variable $XDG_CONFIG_HOME. If this variable is not
	// not set, a default equal to $HOME/.config should be used.
	ConfigHome string

	// CacheHome defines the base directory relative to which user-specific
	// non-essential (cached) data should be written. This directory is
	// defined by the environment variable $XDG_CACHE_HOME. If this variable
	// is not set, a default equal to $HOME/.cache should be used.
	CacheHome string

	// RuntimeDir defines the base directory relative to which user-specific
	// non-essential runtime files and other file objects (such as sockets,
	// named pipes, ...) should be stored. This directory is defined by the
	// environment variable $XDG_RUNTIME_DIR. If this variable is not set,
	// applications should fall back to a replacement directory with similar
	// capabilities. Applications should use this directory for communication
	// and synchronization purposes and should not place larger files in it,
	// since it might reside in runtime memory and cannot necessarily be
	// swapped out to disk.
	RuntimeDir string

	// DataDirs defines the preference-ordered set of base directories to
	// search for data files in addition to the DataHome base directory.
	// This set of directories is defined by the environment variable
	// $XDG_DATA_DIRS. If this variable is not set, the default directories
	// to be used are /usr/local/share and /usr/share, in that order. The
	// DataHome directory is considered more important than any of the
	// directories defined by DataDirs. Therefore, user data files should be
	// written relative to the DataHome directory, if possible.
	DataDirs []string

	// ConfigDirs defines the preference-ordered set of base directories to
	// search for configuration files in addition to the ConfigHome base
	// directory. This set of directories is defined by the environment
	// variable $XDG_CONFIG_DIRS. If this variable is not set, a default
	// equal to /etc/xdg should be used. The ConfigHome directory is
	// considered more important than any of the directories defined by
	// ConfigDirs. Therefore, user config files should be written
	// relative to the ConfigHome directory, if possible.
	ConfigDirs []string
)

func init() {
	Home := os.Getenv("HOME")

	DataHome = xdgPath("DATA_HOME", path.Join(Home, ".local/share"))
	ConfigHome = xdgPath("CONFIG_HOME", path.Join(Home, ".config"))
	CacheHome = xdgPath("CACHE_HOME", path.Join(Home, ".cache"))

	RuntimeDir = xdgPath("RUNTIME_DIR",
		path.Join(os.TempDir(), fmt.Sprintf("%d", os.Getuid())))

	DataDirs = xdgPaths("DATA_DIRS", DataHome,
		[]string{"/usr/local/share", "/usr/share"})

	ConfigDirs = xdgPaths("CONFIG_DIRS", ConfigHome, []string{"/etc/xdg"})
}

// DataFile returns a suitable path containing the specified data file.
// Beside the name of the data file, the name parameter can contain
// directories in which the data file needs to be written (for example:
// appname/appdata.data). If the specified directories do not exist, they
// will be created. If the directories cannot be created, an error containing
// the tried paths is returned.
func DataFile(name string) (string, error) {
	return createPath(name, DataDirs)
}

// ConfigFile returns a suitable path containing the specified config file.
// Beside the name of the config file, the name parameter can contain
// directories in which the config file needs to be written (for example:
// appname/config.yaml). If the specified directories do not exist, they will
// be created. If the directories cannot be created, an error containing the
// tried paths is returned.
func ConfigFile(name string) (string, error) {
	return createPath(name, ConfigDirs)
}

// CacheFile returns a suitable path containing the specified cache file.
// Beside the name of the cache file, the name parameter can contain
// directories in which the config file needs to be written (for example:
// appname/appname.cache). If the specified directories do not exist,
// they will be created. If the directories cannot be created, an error
// containing the tried paths is returned.
func CacheFile(name string) (string, error) {
	return createPath(name, []string{CacheHome})
}

// RuntimeFile returns a suitable path containing the specified runtime file.
// Beside the name of the runtime file, the name parameter can contain
// directories in which the runtime file needs to be written (for example:
// appname/app.pid). If the specified directories do not exist, they will
// be created. If the directories cannot be created, an error containing the
// tried paths is returned.
func RuntimeFile(name string) (string, error) {
	fi, err := os.Lstat(RuntimeDir)
	if err != nil {
		if os.IsNotExist(err) {
			return createPath(name, []string{RuntimeDir})
		}
		return "", err
	}

	if fi.IsDir() {
		// The Runtime directory must be owned by the user
		if err = os.Chown(RuntimeDir, os.Getuid(), os.Getgid()); err != nil {
			return "", err
		}
	} else {
		// For security reasons, the Runtime directory cannot be a symlink
		if err = os.Remove(RuntimeDir); err != nil {
			return "", err
		}
	}

	return createPath(name, []string{RuntimeDir})
}

// SearchDataFile searches for the data file specified by the name
// parameter in the DataHome and DataDirs directories. The name parameter
// can contain directories relative to the search paths (for example:
// appname/appdata.data). If the file cannot be found, an error specifying
// the searched paths is returned.
func SearchDataFile(name string) (string, error) {
	return searchFile(name, DataDirs)
}

// SearchConfigFile searches for the config file specified by the name
// parameter in the ConfigHome and ConfigDirs directories. The name parameter
// can contain directories relative to the search paths (for example:
// appname/config.yaml). If the file cannot be found, an error specifying the
// searched paths is returned.
func SearchConfigFile(name string) (string, error) {
	return searchFile(name, ConfigDirs)
}

// SearchCacheFile searches for the cache file specified by the name
// parameter in the CacheHome directory. The name parameter can contain
// directories relative to the search path (for example: appname/app.cache).
// If the file cannot be found, an error specifying the searched path is
// returned.
func SearchCacheFile(name string) (string, error) {
	return searchFile(name, []string{CacheHome})
}

// SearchRuntimeFile searches for the runtime file specified by the name
// parameter in the RuntimeHome directory. The name parameter can contain
// directories relative to the search path (for example: appname/app.pid).
// If the file cannot be found, an error specifying the searched path is
// returned.
func SearchRuntimeFile(name string) (string, error) {
	return searchFile(name, []string{RuntimeDir})
}

func xdgPath(name, defaultPath string) string {
	dir := os.Getenv(fmt.Sprintf("XDG_%s", name))
	if dir != "" && path.IsAbs(dir) {
		return dir
	}

	return defaultPath
}

func xdgPaths(name, homePath string, defs []string) []string {
	dirs := []string{}

	paths := strings.Split(os.Getenv(fmt.Sprintf("XDG_%s", name)), ":")
	for _, p := range paths {
		if p != "" && path.IsAbs(p) {
			dirs = append(dirs, p)
		}
	}

	if len(dirs) == 0 {
		dirs = append(dirs, defs...)
	}

	if homePath != "" {
		dirs = append([]string{homePath}, dirs...)
	}

	return dirs
}

func searchFile(name string, paths []string) (string, error) {
	for _, p := range paths {
		if filePath := path.Join(p, name); exists(filePath) {
			return filePath, nil
		}
	}

	msg := "Could not locate `%s` in the following paths: %s"
	return "", fmt.Errorf(msg, name, strings.Join(paths, ", "))
}

func createPath(name string, paths []string) (string, error) {
	triedPaths := []string{}

	for _, p := range paths {
		filePath := path.Join(p, name)
		dir, _ := path.Split(filePath)

		if exists(dir) {
			return filePath, nil
		}

		if err := os.MkdirAll(dir, os.ModeDir|0700); err == nil {
			return filePath, nil
		}

		triedPaths = append(triedPaths, dir)
	}

	msg := "Could not create any of the following paths: %s"
	return "", fmt.Errorf(msg, strings.Join(triedPaths, ", "))
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}
