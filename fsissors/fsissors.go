package fsissors

import (
	"os"
	"io"
)

func FileTailCopy(fileName string, pos int64, fileOut string, whence int, bufSize uint) error {
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
	defer fin.Close()
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
	defer out.Close()
	return Copy(fin, out, bufSize)
}

func truncateFile(fin *os.File, pos int64) error {
	err := fin.Truncate(pos)
	if err != nil {
		return err
	}
	fin.Seek(0, 0)
	fin.Sync()
	return nil
}
func Copy(reader io.Reader, writer io.Writer, bufSize uint) (err error) {
	var n int
	buf := make([]byte, bufSize)

	for {
		n, err = reader.Read(buf)

		if err != nil && err != io.EOF {
			return err
		}
		if n > 0 {
			_, werr := writer.Write(buf[:n])
			if werr != nil {
				return werr
			}
		}
		if err == io.EOF {
			return nil
		}
	}
}

func FileTruncate(filename string, pos int64) error {
	fin, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer fin.Close()
	stat, err := fin.Stat()
	if err != nil {
		return err
	}
	if pos >= stat.Size() {
		return nil
	}
	return truncateFile(fin, pos)
}
