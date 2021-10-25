package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// UniquePaths eliminates the duplicate paths from the provided slice and
// returns the result. The items in the output slice are in the order in
// which they occur in the input slice. If a `home` location is provided,
// the paths are expanded using the `ExpandHome` function.
func UniquePaths(paths []string, home string) []string {
	var (
		uniq     []string
		registry = map[string]struct{}{}
	)

	for _, p := range paths {
		p := ExpandHome(p, home)
		if p != "" && filepath.IsAbs(p) {
			if _, ok := registry[p]; ok {
				continue
			}

			registry[p] = struct{}{}
			uniq = append(uniq, p)
		}
	}

	return uniq
}

// CreatePath returns a suitable location relative to which the file with the
// specified `name` can be written, based on the first writable path from the
// provided `paths` slice. The `name` parameter should contain the name of the
// file which is going to be written in the location returned by this function,
// but it can also contain a set of parent directories, which will be created
// relative to the selected parent path.
func CreatePath(name string, paths []string) (string, error) {
	var searchedPaths []string
	for _, p := range paths {
		p = filepath.Join(p, name)

		dir := filepath.Dir(p)
		if PathExists(dir) {
			return p, nil
		}
		if err := os.MkdirAll(dir, os.ModeDir|0700); err == nil {
			return p, nil
		}

		searchedPaths = append(searchedPaths, dir)
	}

	return "", fmt.Errorf("could not create any of the following paths: %s",
		strings.Join(searchedPaths, ", "))
}
