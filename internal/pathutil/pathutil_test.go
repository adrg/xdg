package pathutil_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/adrg/xdg/internal/pathutil"
)

func TestExists(t *testing.T) {
	tempDir := os.TempDir()

	// Test regular file.
	pathFile := filepath.Join(tempDir, "regular")
	f, err := os.Create(pathFile)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.True(t, pathutil.Exists(pathFile))

	// Test symlink.
	pathSymlink := filepath.Join(tempDir, "symlink")
	require.NoError(t, os.Symlink(pathFile, pathSymlink))
	require.True(t, pathutil.Exists(pathSymlink))

	// Test non-existent file.
	require.NoError(t, os.Remove(pathFile))
	require.False(t, pathutil.Exists(pathFile))
	require.False(t, pathutil.Exists(pathSymlink))
	require.NoError(t, os.Remove(pathSymlink))
	require.False(t, pathutil.Exists(pathSymlink))
}

func TestCreate(t *testing.T) {
	tempDir := os.TempDir()

	// Test path selection order.
	p, err := pathutil.Create("test", []string{tempDir, "\000a"})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	p, err = pathutil.Create("test", []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	// Test relative parent directories.
	expected := filepath.Join(tempDir, "appname", "config", "test")
	p, err = pathutil.Create(filepath.Join("appname", "config", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	expected = filepath.Join(tempDir, "appname", "test")
	p, err = pathutil.Create(filepath.Join("appname", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	// Test invalid paths.
	_, err = pathutil.Create(filepath.Join("appname", "test"), []string{"\000a"})
	require.Error(t, err)

	_, err = pathutil.Create("test", []string{filepath.Join(tempDir, "\000a")})
	require.Error(t, err)
}

func TestSearch(t *testing.T) {
	tempDir := os.TempDir()

	// Test file not found.
	_, err := pathutil.Search("test", []string{tempDir, filepath.Join(tempDir, "appname")})
	require.Error(t, err)

	// Test file found.
	expected := filepath.Join(tempDir, "test")
	f, err := os.Create(expected)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	p, err := pathutil.Search("test", []string{tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	p, err = pathutil.Search("test", []string{filepath.Join(tempDir, "appname"), tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	require.NoError(t, os.Remove(expected))

	// Test relative parent directories.
	expected = filepath.Join(tempDir, "appname", "test")
	_, err = pathutil.Create(filepath.Join("appname", "test"), []string{tempDir})
	require.NoError(t, err)
	f, err = os.Create(expected)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	p, err = pathutil.Search(filepath.Join("appname", "test"), []string{tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	p, err = pathutil.Search("test", []string{tempDir, filepath.Join(tempDir, "appname")})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))
}

func TestEnvPath(t *testing.T) {
	home := pathutil.UserHomeDir()
	val := filepath.Join(home, "test")

	require.NoError(t, os.Setenv("PATHUTIL_TEST_VAR", val))
	require.Equal(t, val, pathutil.EnvPath("PATHUTIL_TEST_VAR"))

	require.NoError(t, os.Setenv("PATHUTIL_TEST_VAR", ""))
	require.Equal(t, val, pathutil.EnvPath("PATHUTIL_TEST_VAR", val))
	require.Equal(t, val, pathutil.EnvPath("", val))

	require.NoError(t, os.Unsetenv("PATHUTIL_TEST_VAR"))
}

func TestEnvPathList(t *testing.T) {
	home := pathutil.UserHomeDir()
	pathList := []string{
		filepath.Join(home, "test1"),
		filepath.Join(home, "test2"),
		filepath.Join(home, "test3"),
	}
	val := strings.Join(pathList, string(os.PathListSeparator))

	require.NoError(t, os.Setenv("PATHUTIL_TEST_VAR", val))
	require.Equal(t, val, strings.Join(pathutil.EnvPathList("PATHUTIL_TEST_VAR"), string(os.PathListSeparator)))

	require.NoError(t, os.Setenv("PATHUTIL_TEST_VAR", ""))
	require.Equal(t, val, strings.Join(pathutil.EnvPathList("PATHUTIL_TEST_VAR", pathList...), string(os.PathListSeparator)))
	require.Equal(t, val, strings.Join(pathutil.EnvPathList("", pathList...), string(os.PathListSeparator)))

	require.NoError(t, os.Unsetenv("PATHUTIL_TEST_VAR"))
}

func TestCreateSkipsNonDirectoryParents(t *testing.T) {
	for _, name := range []string{"config.dat", filepath.Join("app", "config.dat")} {
		t.Run(name, func(t *testing.T) {
			first := filepath.Join(t.TempDir(), "first")
			second := t.TempDir()
			blocked := filepath.Dir(filepath.Join(first, name))
			require.NoError(t, os.MkdirAll(filepath.Dir(blocked), 0o700))
			require.NoError(t, os.WriteFile(blocked, []byte("keep me"), 0o600))

			p, err := pathutil.Create(name, []string{first, second})
			require.NoError(t, err)
			require.Equal(t, filepath.Join(second, name), p)
			require.NoError(t, os.WriteFile(p, []byte("config"), 0o600))
			contents, err := os.ReadFile(blocked)
			require.NoError(t, err)
			require.Equal(t, "keep me", string(contents))

			p, err = pathutil.Create(name, []string{first})
			require.Error(t, err)
			require.Empty(t, p)
			require.Contains(t, err.Error(), blocked)
		})
	}
}

func TestCreateWithSymlinkDirectory(t *testing.T) {
	dir := t.TempDir()
	link := filepath.Join(t.TempDir(), "link")
	require.NoError(t, os.Symlink(dir, link))

	p, err := pathutil.Create("config.dat", []string{link})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(link, "config.dat"), p)
	require.NoError(t, os.WriteFile(p, []byte("config"), 0o600))
	contents, err := os.ReadFile(filepath.Join(dir, "config.dat"))
	require.NoError(t, err)
	require.Equal(t, "config", string(contents))
}
