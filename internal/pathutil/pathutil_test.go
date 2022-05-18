package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestExists(t *testing.T) {
	tempDir := os.TempDir()
	fs := afero.NewOsFs()

	// Test regular file.
	pathFile := filepath.Join(tempDir, "regular")
	f, err := os.Create(pathFile)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.True(t, pathutil.Exists(fs, pathFile))

	// Test symlink.
	pathSymlink := filepath.Join(tempDir, "symlink")
	require.NoError(t, os.Symlink(pathFile, pathSymlink))
	require.True(t, pathutil.Exists(fs, pathSymlink))

	// Test non-existent file.
	require.NoError(t, os.Remove(pathFile))
	require.False(t, pathutil.Exists(fs, pathFile))
	require.False(t, pathutil.Exists(fs, pathSymlink))
	require.NoError(t, os.Remove(pathSymlink))
	require.False(t, pathutil.Exists(fs, pathSymlink))
}

func TestCreate(t *testing.T) {
	tempDir := os.TempDir()
	fs := afero.NewOsFs()

	// Test path selection order.
	p, err := pathutil.Create(fs, "test", []string{tempDir, "\000a"})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	p, err = pathutil.Create(fs, "test", []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, filepath.Join(tempDir, "test"), p)

	// Test relative parent directories.
	expected := filepath.Join(tempDir, "appname", "config", "test")
	p, err = pathutil.Create(fs, filepath.Join("appname", "config", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	expected = filepath.Join(tempDir, "appname", "test")
	p, err = pathutil.Create(fs, filepath.Join("appname", "test"), []string{"\000a", tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)
	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))

	// Test invalid paths.
	_, err = pathutil.Create(fs, filepath.Join("appname", "test"), []string{"\000a"})
	require.Error(t, err)

	_, err = pathutil.Create(fs, "test", []string{filepath.Join(tempDir, "\000a")})
	require.Error(t, err)
}

func TestSearch(t *testing.T) {
	tempDir := os.TempDir()
	fs := afero.NewOsFs()

	// Test file not found.
	_, err := pathutil.Search(fs, "test", []string{tempDir, filepath.Join(tempDir, "appname")})
	require.Error(t, err)

	// Test file found.
	expected := filepath.Join(tempDir, "test")
	f, err := os.Create(expected)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	p, err := pathutil.Search(fs, "test", []string{tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	p, err = pathutil.Search(fs, "test", []string{filepath.Join(tempDir, "appname"), tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	require.NoError(t, os.Remove(expected))

	// Test relative parent directories.
	expected = filepath.Join(tempDir, "appname", "test")
	_, err = pathutil.Create(fs, filepath.Join("appname", "test"), []string{tempDir})
	require.NoError(t, err)
	f, err = os.Create(expected)
	require.NoError(t, err)
	require.NoError(t, f.Close())

	p, err = pathutil.Search(fs, filepath.Join("appname", "test"), []string{tempDir})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	p, err = pathutil.Search(fs, "test", []string{tempDir, filepath.Join(tempDir, "appname")})
	require.NoError(t, err)
	require.Equal(t, expected, p)

	require.NoError(t, os.RemoveAll(filepath.Dir(expected)))
}
