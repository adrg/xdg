package xdg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg/internal/util"
)

func createPath(name string, paths []string) (string, error) {
	var searchedPaths []string
	for _, p := range paths {
		path := filepath.Join(p, name)
		dir := filepath.Dir(path)

		if util.PathExists(dir) {
			return path, nil
		}
		if err := os.MkdirAll(dir, os.ModeDir|0700); err == nil {
			return path, nil
		}

		searchedPaths = append(searchedPaths, dir)
	}

	return "", fmt.Errorf("could not create any of the following paths: %s",
		strings.Join(searchedPaths, ", "))
}

func searchFile(name string, paths []string) (string, error) {
	var searchedPaths []string
	for _, p := range paths {
		path := filepath.Join(p, name)
		if util.PathExists(path) {
			return path, nil
		}

		searchedPaths = append(searchedPaths, filepath.Dir(path))
	}

	return "", fmt.Errorf("could not locate `%s` in any of the following paths: %s",
		filepath.Base(name), strings.Join(searchedPaths, ", "))
}

func xdgPath(name, defaultPath string) string {
	dir := util.ExpandHome(os.Getenv(name), Home)
	if dir != "" && filepath.IsAbs(dir) {
		return dir
	}

	return defaultPath
}

func xdgPaths(name string, defaultPaths ...string) []string {
	dirs := uniquePaths(filepath.SplitList(os.Getenv(name)))
	if len(dirs) != 0 {
		return dirs
	}

	return uniquePaths(defaultPaths)
}

func uniquePaths(paths []string) []string {
	var (
		uniq     []string
		registry = map[string]struct{}{}
	)

	for _, p := range paths {
		dir := util.ExpandHome(p, Home)
		if dir != "" && filepath.IsAbs(dir) {
			if _, ok := registry[dir]; ok {
				continue
			}

			registry[dir] = struct{}{}
			uniq = append(uniq, dir)
		}
	}

	return uniq
}
