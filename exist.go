//go:build !windows
// +build !windows

package xdg

import "os"

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
