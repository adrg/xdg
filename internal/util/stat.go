//go:build !windows
// +build !windows

package util

import "os"

// PathExists returns true if the specified path exists.
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
