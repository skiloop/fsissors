package fsissors

import (
	"errors"
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
	if fileName == "" || fileOut == "" {
		return errors.New("文件名不能为空")
	}
	if bufSize > 1024*1024*10 { // 限制缓冲区最大为 10MB
		bufSize = 1024 * 1024 * 10
	}
	if bufSize == 0 {
		bufSize = 1024
	}
	if whence != io.SeekStart && whence != io.SeekCurrent && whence != io.SeekEnd {
		whence = io.SeekStart
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
		return errors.New("position exceed file size")
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
		_ = out.Sync()
		_ = out.Close()
	}(out)
	return Copy(fin, out, size)
}

// Copy copies data from the reader to the writer using a buffer of the specified size.
// If size is greater than 0, it copies exactly that many bytes.
// If size is 0 or less, it copies until EOF.
func Copy(reader io.Reader, writer io.Writer, size int64) (err error) {
	if size > 0 {
		// Copy exactly 'size' bytes from reader to writer
		_, err = io.CopyN(writer, reader, size)
	} else {
		// Copy until EOF
		_, err = io.Copy(writer, reader)
	}
	return err
}

// MemCopyFile /*
//func MemCopyFile(in string, from int64, size int64, out string, offset int64) error {
//	return nil
//}
