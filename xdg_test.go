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
