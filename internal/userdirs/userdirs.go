package userdirs

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/adrg/xdg/internal/pathutil"
)

// ParseConfigFile parses the user directories config file at the specified
// location. The returned map contains pairs consisting of the user directory
// names and their paths.
func ParseConfigFile(name string) map[string]string {
	f, err := os.Open(name)
	if err != nil {
		return nil
	}
	defer f.Close()

	return ParseConfig(f)
}

// ParseConfig parses the user directories config file contained in the
// provided reader. The returned map contains pairs consisting of the user
// directory names and their paths.
func ParseConfig(r io.Reader) map[string]string {
	dirs := map[string]string{
		"XDG_DESKTOP_DIR":     "",
		"XDG_DOWNLOAD_DIR":    "",
		"XDG_DOCUMENTS_DIR":   "",
		"XDG_MUSIC_DIR":       "",
		"XDG_PICTURES_DIR":    "",
		"XDG_VIDEOS_DIR":      "",
		"XDG_TEMPLATES_DIR":   "",
		"XDG_PUBLICSHARE_DIR": "",
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		if !strings.HasPrefix(line, "XDG_") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			continue
		}

		// Parse key.
		key := strings.TrimSpace(parts[0])
		if _, ok := dirs[key]; !ok {
			continue
		}

		// Parse value.
		runes := []rune(strings.TrimSpace(parts[1]))

		lenRunes := len(runes)
		if lenRunes <= 2 || runes[0] != '"' {
			continue
		}

		for i := 1; i < lenRunes; i++ {
			if runes[i] == '"' {
				dirs[key] = pathutil.ExpandHome(string(runes[1:i]))
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil
	}

	return dirs
}
