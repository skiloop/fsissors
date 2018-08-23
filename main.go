package main

import (
	"fmt"
	"github.com/skiloop/fsissors/fsissors"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	copyCmd        = kingpin.Command("copy", "copy a file").Alias("c")
	copyInput      = copyCmd.Arg("source", "file to copy from").Required().String()
	copyOutput     = copyCmd.Arg("target", "target file").Required().String()
	copyFrom       = copyCmd.Flag("from", "offset to copy").Short('f').Default("0").Int64()
	copySize       = copyCmd.Flag("size", "size to copy, 0 to copy to end of file").Short('s').Default("0").Int64()
	copyBufferSize = copyCmd.Flag("buffer", "copy buffer size in bytes").Short('b').Default("1024").Uint()

	truncate   = kingpin.Command("truncate", "truncate a file").Alias("t").Alias("trun")
	trunInput  = truncate.Arg("input", "file to truncate").Required().String()
	trunOutput = truncate.Arg("target", "file to truncate").String()
	trunFrom   = truncate.Flag("from", "offset to truncate").Required().Int64()
)

func copyFile() {
	err := fsissors.FileCopy(*copyInput, *copyFrom, *copyOutput, 0, *copyBufferSize, *copySize)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file copy error: %s\n", err.Error())
		return
	}
}
func truncateFile() {
	if *trunOutput == "" {
		accept := 'N'
		fmt.Fprintf(os.Stdout, "truncate %s to %d (y/N)", *trunInput, *trunFrom)
		fmt.Scanf("%c", &accept)
		if accept != 'y' && accept != 'Y' {
			return
		}
	}
	err := fsissors.FileTruncate(*trunInput, *trunFrom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to truncate file: %s", err.Error())
	}
}
func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	cmd := kingpin.Parse()
	switch cmd {
	case copyCmd.FullCommand():
		copyFile()
	case truncate.FullCommand():
		truncateFile()
	}
}
