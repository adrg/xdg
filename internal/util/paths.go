package util

import "path/filepath"

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
