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
	// Home contains the path of the user's home directory.
	Home string

	// DataHome defines the base directory relative to which user-specific
	// data files should be stored. This directory is defined by the
	// $XDG_DATA_HOME environment variable. If the variable is not set,
	// a default equal to $HOME/.local/share should be used.
	DataHome string

	// DataDirs defines the preference-ordered set of base directories to
	// search for data files in addition to the DataHome base directory.
	// This set of directories is defined by the $XDG_DATA_DIRS environment
	// variable. If the variable is not set, the default directories
	// to be used are /usr/local/share and /usr/share, in that order. The
	// DataHome directory is considered more important than any of the
	// directories defined by DataDirs. Therefore, user data files should be
	// written relative to the DataHome directory, if possible.
	DataDirs []string

	// ConfigHome defines the base directory relative to which user-specific
	// configuration files should be written. This directory is defined by
	// the $XDG_CONFIG_HOME environment variable. If the variable is not
	// not set, a default equal to $HOME/.config should be used.
	ConfigHome string

	// ConfigDirs defines the preference-ordered set of base directories to
	// search for configuration files in addition to the ConfigHome base
	// directory. This set of directories is defined by the $XDG_CONFIG_DIRS
	// environment variable. If the variable is not set, a default equal
	// to /etc/xdg should be used. The ConfigHome directory is considered
	// more important than any of the directories defined by ConfigDirs.
	// Therefore, user config files should be written relative to the
	// ConfigHome directory, if possible.
	ConfigDirs []string

	// StateHome defines the base directory relative to which user-specific
	// state files should be stored. This directory is defined by the
	// $XDG_STATE_HOME environment variable. If the variable is not set,
	// a default equal to ~/.local/state should be used.
	StateHome string

	// CacheHome defines the base directory relative to which user-specific
	// non-essential (cached) data should be written. This directory is
	// defined by the $XDG_CACHE_HOME environment variable. If the variable
	// is not set, a default equal to $HOME/.cache should be used.
	CacheHome string

	// RuntimeDir defines the base directory relative to which user-specific
	// non-essential runtime files and other file objects (such as sockets,
	// named pipes, etc.) should be stored. This directory is defined by the
	// $XDG_RUNTIME_DIR environment variable. If the variable is not set,
	// applications should fall back to a replacement directory with similar
	// capabilities. Applications should use this directory for communication
	// and synchronization purposes and should not place larger files in it,
	// since it might reside in runtime memory and cannot necessarily be
	// swapped out to disk.
	RuntimeDir string

	// UserDirs defines the locations of well known user directories.
	UserDirs UserDirectories

	// FontDirs defines the common locations where font files are stored.
	FontDirs []string

	// ApplicationDirs defines the common locations of applications.
	ApplicationDirs []string

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
	bd := BaseDirectories{
		fs: fs,
	}

	bd.Reload()

	return bd
}

// Reload refreshes base and user directories by reading the environment.
// Defaults are applied for XDG variables which are empty or not present
// in the environment.
func (bd *BaseDirectories) Reload() {
	// Initialize home directory.
	bd.Home = homeDir()

	// Initialize base and user directories.
	initDirs(bd, bd.Home)

	// Set standard directories.
	bd.DataHome = bd.dataHome
	bd.DataDirs = bd.data
	bd.ConfigHome = bd.configHome
	bd.ConfigDirs = bd.config
	bd.StateHome = bd.stateHome
	bd.CacheHome = bd.cacheHome
	bd.RuntimeDir = bd.runtime

	// Set non-standard directories.
	bd.FontDirs = bd.fonts
	bd.ApplicationDirs = bd.applications
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
