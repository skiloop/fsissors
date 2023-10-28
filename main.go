package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/skiloop/fsissors/fsissors"
	"os"
)

type CopyCmd struct {
	Source     string `arg:"" help:"file to copy from"`
	Target     string `arg:"" help:"output file"`
	From       int64  `short:"f" optional:"" help:"offset to copy from" default:"0"`
	Size       int64  `short:"s" optional:"" help:"how many bytes to copy, 0 means to end of file" default:"0"`
	BufferSize uint   `short:"b" optional:"" help:"copy buffer size in bytes" default:"1024"`
}
type TruncateCmd struct {
	Input string `arg:"" help:"source file to truncate"`
	Size  int64  `arg:"" help:"size of output file. Nothing will be done when negative, equal to or  larger than origin file size"`
}

var client struct {
	Copy     CopyCmd     `cmd:"" aliases:"c" help:"copy part of a file"`
	Truncate TruncateCmd `cmd:"" aliases:"t" help:"truncate a file"`
}

func copyFile() {
	err := fsissors.FileCopy(client.Copy.Source, client.Copy.From,
		client.Copy.Target, 0, client.Copy.BufferSize, client.Copy.Size)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "file copy error: %s\n", err.Error())
		return
	}
}
func truncateFile() {
	if client.Truncate.Input == "" {
		accept := 'N'
		fmt.Printf("truncate %s to size %d (y/n)", client.Truncate.Input, client.Truncate.Size)
		_, _ = fmt.Scanf("%c", &accept)
		if accept != 'y' && accept != 'Y' {
			fmt.Println("nothing is done")
			return
		}
	}
	err := fsissors.FileTruncate(client.Truncate.Input, client.Truncate.Size)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to truncate file: %s", err.Error())
	}
}
func main() {
	ctx := kong.Parse(&client)
	switch ctx.Command() {
	case "copy <source> <target>":
		copyFile()
	case "truncate <input> <offset>":
		truncateFile()
	}
}
