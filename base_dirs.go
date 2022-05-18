package xdg

import (
	"github.com/adrg/xdg/internal/pathutil"
	"github.com/spf13/afero"
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
)

type BaseDirectories struct {
	fs afero.Fs

	dataHome   string
	data       []string
	configHome string
	config     []string
	stateHome  string
	cacheHome  string
	runtime    string

	// Non-standard directories.
	fonts        []string
	applications []string
}

func New(fs afero.Fs) BaseDirectories {
	return BaseDirectories{
		fs: fs,
	}
}

func (bd BaseDirectories) DataFile(relPath string) (string, error) {
	return pathutil.Create(bd.fs, relPath, append([]string{bd.dataHome}, bd.data...))
}

func (bd BaseDirectories) ConfigFile(relPath string) (string, error) {
	return pathutil.Create(bd.fs, relPath, append([]string{bd.configHome}, bd.config...))
}

func (bd BaseDirectories) StateFile(relPath string) (string, error) {
	return pathutil.Create(bd.fs, relPath, []string{bd.stateHome})
}

func (bd BaseDirectories) CacheFile(relPath string) (string, error) {
	return pathutil.Create(bd.fs, relPath, []string{bd.cacheHome})
}

func (bd BaseDirectories) RuntimeFile(relPath string) (string, error) {
	return pathutil.Create(bd.fs, relPath, []string{bd.runtime})
}

func (bd BaseDirectories) SearchDataFile(relPath string) (string, error) {
	return pathutil.Search(bd.fs, relPath, append([]string{bd.dataHome}, bd.data...))
}

func (bd BaseDirectories) SearchConfigFile(relPath string) (string, error) {
	return pathutil.Search(bd.fs, relPath, append([]string{bd.configHome}, bd.config...))
}

func (bd BaseDirectories) SearchStateFile(relPath string) (string, error) {
	return pathutil.Search(bd.fs, relPath, []string{bd.stateHome})
}

func (bd BaseDirectories) SearchCacheFile(relPath string) (string, error) {
	return pathutil.Search(bd.fs, relPath, []string{bd.cacheHome})
}

func (bd BaseDirectories) SearchRuntimeFile(relPath string) (string, error) {
	return pathutil.Search(bd.fs, relPath, []string{bd.runtime})
}
