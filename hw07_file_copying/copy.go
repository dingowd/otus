package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Bar ...
type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Открываем исходный файл
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer srcFile.Close()
	// Создаем новый файл
	dstFile, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer dstFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	if offset > srcInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	buf := make([]byte, 1)
	var bytesToCopy, n int64
	if limit == 0 || limit > srcInfo.Size()-offset {
		bytesToCopy = srcInfo.Size() - offset
	} else {
		bytesToCopy = limit + 1
	}
	if _, err := srcFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	var bar Bar
	bar.NewOption(0, bytesToCopy)

	// Работает
	for {
		if n == bytesToCopy {
			break
		}
		if _, err := srcFile.Read(buf); err != nil {
			panic(err)
		}
		if _, err := dstFile.Write(buf); err != nil {
			panic(err)
		}
		n++
		bar.Play(n)
	}
	return nil
}
