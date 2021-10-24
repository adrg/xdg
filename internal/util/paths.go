package util

import (
	"path/filepath"
	"strings"
)

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
