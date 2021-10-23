package util

import (
	"os"
	"path/filepath"
)

// PathExists returns true if the specified path exists.
func PathExists(path string) bool {
	fi, err := os.Lstat(path)
	if fi != nil && fi.Mode()&os.ModeSymlink != 0 {
		_, err = filepath.EvalSymlinks(path)
	}

	return err == nil || os.IsExist(err)
}
