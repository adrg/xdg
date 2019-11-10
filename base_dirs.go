package xdg

import "os"

// XDG Base Directory environment variables.
var (
	envDataHome   = "XDG_DATA_HOME"
	envDataDirs   = "XDG_DATA_DIRS"
	envConfigHome = "XDG_CONFIG_HOME"
	envConfigDirs = "XDG_CONFIG_DIRS"
	envCacheHome  = "XDG_CACHE_HOME"
	envRuntimeDir = "XDG_RUNTIME_DIR"
)

type baseDirectories struct {
	DataHome   string
	Data       []string
	ConfigHome string
	Config     []string
	CacheHome  string
	Runtime    string
}

func (bd baseDirectories) dataFile(relPath string) (string, error) {
	return createPath(relPath, append([]string{bd.DataHome}, bd.Data...))
}

func (bd baseDirectories) configFile(relPath string) (string, error) {
	return createPath(relPath, append([]string{bd.ConfigHome}, bd.Config...))
}

func (bd baseDirectories) cacheFile(relPath string) (string, error) {
	return createPath(relPath, []string{bd.CacheHome})
}

func (bd baseDirectories) runtimeFile(relPath string) (string, error) {
	fi, err := os.Lstat(bd.Runtime)
	if err != nil {
		if os.IsNotExist(err) {
			return createPath(relPath, []string{bd.Runtime})
		}
		return "", err
	}

	if fi.IsDir() {
		// The Runtime directory must be owned by the user.
		if err = os.Chown(bd.Runtime, os.Getuid(), os.Getgid()); err != nil {
			return "", err
		}
	} else {
		// For security reasons, the Runtime directory cannot be a symlink.
		if err = os.Remove(bd.Runtime); err != nil {
			return "", err
		}
	}

	return createPath(relPath, []string{bd.Runtime})
}

func (bd baseDirectories) searchDataFile(relPath string) (string, error) {
	return searchFile(relPath, append([]string{bd.ConfigHome}, bd.Data...))
}

func (bd baseDirectories) searchConfigFile(relPath string) (string, error) {
	return searchFile(relPath, append([]string{bd.ConfigHome}, bd.Config...))
}

func (bd baseDirectories) searchCacheFile(relPath string) (string, error) {
	return searchFile(relPath, []string{bd.CacheHome})
}

func (bd baseDirectories) searchRuntimeFile(relPath string) (string, error) {
	return searchFile(relPath, []string{bd.Runtime})
}
