//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || nacl || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js,wasm nacl linux netbsd openbsd solaris

package pathutil

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// Exists returns true if the specified path exists.
func Exists(fs afero.Fs, path string) bool {
	_, err := fs.Stat(path)
	return err == nil || os.IsExist(err)
}

// ExpandHome substitutes `~` and `$HOME` at the start of the specified
// `path` using the provided `home` location.
func ExpandHome(path, home string) string {
	if path == "" || home == "" {
		return path
	}
	if path[0] == '~' {
		return filepath.Join(home, path[1:])
	}
	if strings.HasPrefix(path, "$HOME") {
		return filepath.Join(home, path[5:])
	}

	return path
}
