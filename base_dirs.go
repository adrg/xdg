package xdg

import (
	"os"

	"github.com/adrg/xdg/internal/pathutil"
)

// XDG Base Directory environment variables.
const (
	envDataHome   = "XDG_DATA_HOME"
	envDataDirs   = "XDG_DATA_DIRS"
	envConfigHome = "XDG_CONFIG_HOME"
	envConfigDirs = "XDG_CONFIG_DIRS"
	envStateHome  = "XDG_STATE_HOME"
	envCacheHome  = "XDG_CACHE_HOME"
	envRuntimeDir = "XDG_RUNTIME_DIR"

	// Non-standard.
	envBinHome = "XDG_BIN_HOME"
)

type baseDirectories struct {
	dataHome   string
	data       []string
	configHome string
	config     []string
	stateHome  string
	cacheHome  string
	runtime    string

	// Non-standard.
	binHome      string
	fonts        []string
	applications []string
}

func (bd baseDirectories) dataFile(relPath string) (string, error) {
	return pathutil.Create(relPath, append([]string{bd.dataHome}, bd.data...))
}

func (bd baseDirectories) configFile(relPath string) (string, error) {
	return pathutil.Create(relPath, append([]string{bd.configHome}, bd.config...))
}

func (bd baseDirectories) stateFile(relPath string) (string, error) {
	return pathutil.Create(relPath, []string{bd.stateHome})
}

func (bd baseDirectories) cacheFile(relPath string) (string, error) {
	return pathutil.Create(relPath, []string{bd.cacheHome})
}

func (bd baseDirectories) runtimeFile(relPath string) (string, error) {
	var paths []string
	for _, p := range pathutil.Unique([]string{bd.runtime, os.TempDir()}) {
		if pathutil.Exists(p) {
			paths = append(paths, p)
		}
	}
	return pathutil.Create(relPath, paths)
}

func (bd baseDirectories) searchDataFile(relPath string) (string, error) {
	return pathutil.Search(relPath, append([]string{bd.dataHome}, bd.data...))
}

func (bd baseDirectories) searchConfigFile(relPath string) (string, error) {
	return pathutil.Search(relPath, append([]string{bd.configHome}, bd.config...))
}

func (bd baseDirectories) searchStateFile(relPath string) (string, error) {
	return pathutil.Search(relPath, []string{bd.stateHome})
}

func (bd baseDirectories) searchCacheFile(relPath string) (string, error) {
	return pathutil.Search(relPath, []string{bd.cacheHome})
}

func (bd baseDirectories) searchRuntimeFile(relPath string) (string, error) {
	return pathutil.Search(relPath, pathutil.Unique([]string{bd.runtime, os.TempDir()}))
}
