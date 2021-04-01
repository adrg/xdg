package xdg

import "os"

func exists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil || os.IsExist(err)
}
