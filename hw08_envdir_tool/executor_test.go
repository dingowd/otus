package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Test Run command", func(t *testing.T) {
		env, _ := ReadDir("testdata/env")
		expected := "bar"
		cmd := make([]string, 1)
		RunCmd(cmd, env)
		bar, _ := os.LookupEnv("BAR")
		require.Equal(t, expected, bar)
	})
}
