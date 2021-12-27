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

// Функция для прорисовки прогресс-бара.
func Bar(cur, total int64) {
	rate := ""
	progress := int64(float64(cur) / float64(total) * 100)
	for i := 0; i < int(progress/2); i++ {
		rate += "█"
	}
	fmt.Printf("\r[%-50s]%3d%%", rate, progress)
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
	// Получаем статистику по файлу-источнику
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	// Проверяем, что смещение не больше размера исходного файла
	if offset > srcInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	// Если смещение равно размеру исходного файла, то копируем 0 байт
	if offset == srcInfo.Size() {
		Bar(1, 1)
		return nil
	}
	// Для отображения прогресс-бара создаем буфер с размером 1
	buf := make([]byte, 1)
	// Вычисляем количество байт для копирования в соответствии с условиями задачи
	var bytesToCopy, n int64
	if limit == 0 || limit > srcInfo.Size()-offset {
		bytesToCopy = srcInfo.Size() - offset
	} else {
		bytesToCopy = limit
	}
	// Устанавливаем позицию, с которой нужно копировать
	if _, err := srcFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	// Копируем побайтово в заданный файл и выводим прогресс-бар
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
		Bar(n, bytesToCopy)
	}
	return nil
}
