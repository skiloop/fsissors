package fsissors

import (
	"fmt"
	"io"
	"os"
)

// FileTruncate truncates the file to the specified size.
// If the size is greater than or equal to the current file size, it does nothing and returns nil.
// If the size is negative and -size not equals to filesize, it does remove the beginning part of the file, -size is the removal size .
// Otherwise, it truncates the file to the specified size.
func FileTruncate(filename string, size int64) error {
	if size == 0 {
		fmt.Println("size is zero, nothing is done")
		return nil
	}
	fin, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func(fin *os.File) {
		_ = fin.Close()
	}(fin)
	stat, err := fin.Stat()
	if err != nil {
		return err
	}
	fileSize := stat.Size()
	if size > 0 {
		if size >= fileSize {
			if Verbose {
				fmt.Printf("input size %d >= %d\n", size, stat.Size())
			}
			return nil
		}
		fmt.Printf("truncate %s to size %d\n", filename, size)
		return truncateFile(fin, size)
	}

	size = -size
	if size >= fileSize {
		if Verbose {
			fmt.Printf("input size %d >= %d\n", size, stat.Size())
		}
		return nil
	}
	fmt.Printf("remove front part of file %s from begining to %d\n", filename, size)
	return removeFileFrontPart(fin, fileSize, size)
}

func truncateFile(fin *os.File, pos int64) error {
	err := fin.Truncate(pos)
	if err != nil {
		return err
	}
	return fin.Sync()
}

func removeFileFrontPart(file *os.File, fileSize, offset int64) error {

	// Move the remaining data to the beginning of the file
	bufSize := int64(1024 * 1024) // 1MB buffer
	buf := make([]byte, bufSize)
	var readPos, writePos int64 = offset, 0

	for readPos < fileSize {
		_, err := file.Seek(readPos, io.SeekStart)
		if err != nil {
			return err
		}
		n, err := file.Read(buf)
		if err != nil {
			return err
		}
		_, err = file.Seek(writePos, io.SeekStart)
		if err != nil {
			return err
		}
		_, err = file.Write(buf[:n])
		if err != nil {
			return err
		}
		readPos += int64(n)
		writePos += int64(n)
	}

	// Truncate the file to the new size
	return truncateFile(file, fileSize-offset)
}
