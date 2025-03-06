package fsissors

import (
	"encoding/hex"
	"fmt"
	"os"
)

// BytesModify modifies the specified part of the file to the specified value
func BytesModify(filename string, start uint32, count, size uint, hexData string) error {
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return fmt.Errorf("failed to decode hex data: %v", err)
	}

	if uint(len(data)) != size {
		return fmt.Errorf("data size %d does not match specified size %d", len(data), size)
	}

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(int64(start), 0)
	if err != nil {
		return err
	}

	for i := uint(0); i < count; i++ {
		_, err = file.Write(data)
		if err != nil {
			return err
		}
	}

	return file.Sync()
}
