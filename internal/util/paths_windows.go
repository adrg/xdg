package util

import (
	"os"

	"golang.org/x/sys/windows"
)

// KnownFolderPath returns the location of the folder with the specified ID.
// If that fails, the folder location is determined by reading the provided
// environment variables (the first non-empty value is returned).
// If that fails as well, the first non-empty fallback is returned.
// If all of the above fails, the function returns an empty string.
func KnownFolderPath(id *windows.KNOWNFOLDERID, envVars []string, fallbacks []string) string {
	if id != nil {
		flags := []uint32{windows.KF_FLAG_DEFAULT, windows.KF_FLAG_DEFAULT_PATH}
		for _, flag := range flags {
			if p, _ := windows.KnownFolderPath(id, flag|windows.KF_FLAG_DONT_VERIFY); p != "" {
				return p
			}
		}
	}

	for _, envVar := range envVars {
		if p := os.Getenv(envVar); p != "" {
			return p
		}
	}

	for _, fallback := range fallbacks {
		if fallback != "" {
			return fallback
		}
	}

	return ""
}
