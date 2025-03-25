package fsissors

import (
	"fmt"
	"io"
	"os"
)

// FileTruncate truncates the file to the specified size.
// If the size is negative, it does nothing and returns nil.
// If the size is greater than or equal to the current file size, it does nothing and returns nil.
// Otherwise, it truncates the file to the specified size.
func FileTruncate(filename string, size int64) error {
	if size < 0 {
		if Verbose {
			fmt.Printf("nothing is done for size is negative: %d\n", size)
		}
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
	if size >= stat.Size() {
		if Verbose {
			fmt.Printf("input size %d >= %d\n", size, stat.Size())
		}
		return nil
	}
	fmt.Printf("truncate %s to size %d\n", filename, size)
	return truncateFile(fin, size)
}

func truncateFile(fin *os.File, pos int64) error {
	err := fin.Truncate(pos)
	if err != nil {
		return err
	}
	_, err = fin.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	return fin.Sync()
}
