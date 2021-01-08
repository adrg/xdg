package xdg_test

import (
	"os"
	"testing"

	"github.com/adrg/xdg"
	"github.com/stretchr/testify/assert"
)

type envSample struct {
	name     string
	value    string
	expected interface{}
	actual   interface{}
}

func testDirs(t *testing.T, samples ...*envSample) {
	// Test home directory.
	if !assert.NotEmpty(t, xdg.Home) {
		t.FailNow()
	}
	t.Logf("Home: %s", xdg.Home)

	// Set environment variables.
	for _, sample := range samples {
		assert.NoError(t, os.Setenv(sample.name, sample.value))
	}
	xdg.Reload()

	// Test results.
	for _, sample := range samples {
		var actual interface{}
		switch v := sample.actual.(type) {
		case *string:
			actual = *v
		case *[]string:
			actual = *v
		}

		assert.Equal(t, sample.expected, actual)
		t.Logf("%s: %v", sample.name, actual)
	}
}

func TestBaseDirFuncs(t *testing.T) {
	type inputData struct {
		relPaths   []string
		pathFunc   func(string) (string, error)
		searchFunc func(string) (string, error)
	}

	inputs := []*inputData{
		{
			relPaths:   []string{"app.data", "appname/app.data"},
			pathFunc:   xdg.DataFile,
			searchFunc: xdg.SearchDataFile,
		},
		{
			relPaths:   []string{"app.yaml", "appname/app.yaml"},
			pathFunc:   xdg.ConfigFile,
			searchFunc: xdg.SearchConfigFile,
		},
		{
			relPaths:   []string{"app.cache", "appname/app.cache"},
			pathFunc:   xdg.CacheFile,
			searchFunc: xdg.SearchCacheFile,
		},
		{
			relPaths:   []string{"app.pid", "appname/app.pid"},
			pathFunc:   xdg.RuntimeFile,
			searchFunc: xdg.SearchRuntimeFile,
		},
		{
			relPaths:   []string{"app.state", "appname/app.state"},
			pathFunc:   xdg.StateFile,
			searchFunc: xdg.SearchStateFile,
		},
	}

	for _, input := range inputs {
		for _, relPath := range input.relPaths {
			// Get suitable path for input file.
			expFullPath, err := input.pathFunc(relPath)
			assert.NoError(t, err)

			// Create input file.
			f, err := os.Create(expFullPath)
			assert.NoError(t, err)
			assert.NoError(t, f.Close())

			// Search input file after creation.
			actFullPath, err := input.searchFunc(relPath)
			assert.NoError(t, err)
			assert.Equal(t, expFullPath, actFullPath)

			// Remove created file.
			assert.NoError(t, os.Remove(expFullPath))

			// Search input file after removal.
			_, err = input.searchFunc(relPath)
			assert.Error(t, err)

			// Check that the same path is returned.
			actFullPath, err = input.pathFunc(relPath)
			assert.NoError(t, err)
			assert.Equal(t, expFullPath, actFullPath)
		}
	}
}
