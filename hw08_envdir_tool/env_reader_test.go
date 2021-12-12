package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Read value from BAR", func(t *testing.T) {
		expected := "bar"
		env, _ := ReadDir("testdata/env")
		require.Equal(t, expected, env["BAR"].Value)
	})
	t.Run("Read empty file", func(t *testing.T) {
		expected := false
		env, _ := ReadDir("testdata/env")
		require.Equal(t, expected, env["BAR"].NeedRemove)
	})
	t.Run("Read existing dir", func(t *testing.T) {
		_, err := ReadDir("testdata/env")
		require.Nil(t, err)
	})
}
