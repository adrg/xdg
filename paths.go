package xdg

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg/internal/util"
)

func xdgPath(name, defaultPath string) string {
	dir := util.ExpandHome(os.Getenv(name), Home)
	if dir != "" && filepath.IsAbs(dir) {
		return dir
	}

	return defaultPath
}

func xdgPaths(name string, defaultPaths ...string) []string {
	dirs := util.UniquePaths(filepath.SplitList(os.Getenv(name)), Home)
	if len(dirs) != 0 {
		return dirs
	}

	return util.UniquePaths(defaultPaths, Home)
}
