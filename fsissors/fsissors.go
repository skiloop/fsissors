package fsissors

import (
	"fmt"
	"io"
	"os"
)

// FileCopy 从源文件指定位置复制指定大小的数据到目标文件
// fileName: 源文件名
// pos: 起始位置
// fileOut: 目标文件名
// whence: 位置基准，可以是 io.SeekStart, io.SeekCurrent, io.SeekEnd
// bufSize: 缓冲区大小，如果为0则使用默认值1024
// size: 要复制的字节数，如果为0则复制到文件末尾
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
// Copy copies data from the reader to the writer using a buffer of the specified size.
// If size is greater than 0, it copies exactly that many bytes.
// If size is 0 or less, it copies until EOF.
func Copy(reader io.Reader, writer io.Writer, bufSize uint, size int64) (err error) {
	if size > 0 {
		// Copy exactly 'size' bytes from reader to writer
		_, err = io.CopyN(writer, reader, size)
	} else {
		// Copy until EOF
		_, err = io.Copy(writer, reader)
	}
	return err
}

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

// MemCopyFile /*
//func MemCopyFile(in string, from int64, size int64, out string, offset int64) error {
//	return nil
//}
