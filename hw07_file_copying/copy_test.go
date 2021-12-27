package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const to = "output.txt"
	t.Run("1000 bytes copy", func(t *testing.T) {
		from := "./testdata/input2.txt"
		defer os.Remove(to)
		expected := 1000
		Copy(from, to, 0, 1000)
		dstFile, _ := os.Open(to)
		dstInfo, _ := dstFile.Stat()
		dstFile.Close()
		require.Equal(t, expected, int(dstInfo.Size()))
	})
	t.Run("Big offset", func(t *testing.T) {
		from := "./testdata/input2.txt"
		srcFile, _ := os.Open(from)
		srcInfo, _ := srcFile.Stat()
		offset := srcInfo.Size() + 10
		srcFile.Close()
		defer os.Remove(to)
		expected := ErrOffsetExceedsFileSize
		require.Equal(t, expected, Copy(from, to, offset, 1000))
	})
	t.Run("File not exist", func(t *testing.T) {
		from := "./testdata/input3.txt"
		defer os.Remove(to)
		expected := ErrUnsupportedFile
		require.Equal(t, expected, Copy(from, to, 0, 1000))
	})
}
