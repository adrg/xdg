package pathutil_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/pathutil"
	"github.com/stretchr/testify/require"
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
