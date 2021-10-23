//go:build windows
// +build windows

package util_test

import (
	"testing"

	"github.com/adrg/xdg/internal/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/windows"
)

func TestKnownFolderPath(t *testing.T) {
	expected := `C:\ProgramData`
	require.Equal(t, expected, util.KnownFolderPath(windows.FOLDERID_ProgramData, nil, nil))
	require.Equal(t, expected, util.KnownFolderPath(nil, []string{"ProgramData"}, nil))
	require.Equal(t, expected, util.KnownFolderPath(nil, nil, []string{expected}))
	require.Equal(t, "", util.KnownFolderPath(nil, nil, nil))
}
