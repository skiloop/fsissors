package fsissors

import (
	"fmt"
	"io"
	"os"
)

func FileCopy(fileName string, pos int64, fileOut string, whence int, bufSize uint, size int64) error {
	if bufSize == 0 {
		bufSize = 1024
	}
	if whence != 0 && whence != 1 && whence != 2 {
		whence = 0
	}
	fin, err := os.Open(fileName)
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
	if pos >= stat.Size() {
		return err
	}
	_, err = fin.Seek(pos, whence)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(fileOut, os.O_CREATE|os.O_APPEND|os.O_WRONLY, stat.Mode())
	if err != nil {

		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	return Copy(fin, out, bufSize, size)
}

func truncateFile(fin *os.File, pos int64) error {
	err := fin.Truncate(pos)
	if err != nil {
		return err
	}
	_, _ = fin.Seek(0, 0)
	_ = fin.Sync()
	return nil
}
func Copy(reader io.Reader, writer io.Writer, bufSize uint, size int64) (err error) {
	var n int
	var copySize int64
	copySize = 0
	buf := make([]byte, bufSize)
	for {
		n, err = reader.Read(buf)

		if err != nil && err != io.EOF {
			return err
		}
		if n > 0 {
			if size > 0 && copySize+int64(n) > size {
				n = int(size - copySize)
			}
			_, writeErr := writer.Write(buf[:n])
			if writeErr != nil {
				return writeErr
			}
			copySize += int64(n)
		}
		if err == io.EOF || (size > 0 && copySize >= size) {
			break
		}
	}
	return nil
}

// FileTruncate truncate file
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

// MemCopyFile /*
//func MemCopyFile(in string, from int64, size int64, out string, offset int64) error {
//	return nil
//}
