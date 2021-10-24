package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/adrg/xdg/internal/util"
	"github.com/stretchr/testify/require"
)

func TestPathExists(t *testing.T) {
	tempDir := os.TempDir()

	// Test regular file.
	pathFile := filepath.Join(tempDir, "regular")
	f, err := os.Create(pathFile)
	require.NoError(t, err)
	require.NoError(t, f.Close())
	require.True(t, util.PathExists(pathFile))

	// Test symlink.
	pathSymlink := filepath.Join(tempDir, "symlink")
	require.NoError(t, os.Symlink(pathFile, pathSymlink))
	require.True(t, util.PathExists(pathSymlink))

	// Test non-existent file.
	require.NoError(t, os.Remove(pathFile))
	require.False(t, util.PathExists(pathFile))
	require.False(t, util.PathExists(pathSymlink))
	require.NoError(t, os.Remove(pathSymlink))
	require.False(t, util.PathExists(pathSymlink))
}
