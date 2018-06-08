package main

import (
	"flag"
	"github.com/skiloop/fsissors/fsissors"
)

var (
	input    = flag.String("i", "", "input filename")
	output   = flag.String("o", "", "destination filename")
	position = flag.Int64("p", 0, "start position to copy")
	whence   = flag.Int("w", 0, "according to whence: 0 means relative to the origin of the file, 1 means")
	bufSize  = flag.Uint("b", 1024, "buffer size")
)

func main() {
	flag.Parse()
	fsissors.FileTailCopy(*input, *position, *output, *whence, *bufSize)
}
