package main

import (
	"flag"
	"fmt"
	"github.com/skiloop/fsissors/fsissors"
)

var (
	input    = flag.String("i", "", "input filename")
	output   = flag.String("o", "", "destination filename")
	position = flag.Int64("p", 0, "start position to copy")
	whence   = flag.Int("w", 0, "according to whence: 0 means relative to the origin of the file, 1 means "+
		"relative to the current offset, and 2 means relative to the end.")
	bufSize  = flag.Uint("b", 1024, "buffer size")
	truncate = flag.Bool("t", false, "truncate file")
)

func main() {
	flag.Parse()

	if *output != "" {
		err := fsissors.FileTailCopy(*input, *position, *output, *whence, *bufSize)
		if err != nil {
			fmt.Printf("file copy error: %s\n", err.Error())
		}
	}
	if *truncate {
		err := fsissors.FileTruncate(*input, *position)
		if err != nil {
			fmt.Printf("failed to truncate file: %s", err.Error())
		}
	}
}
}
