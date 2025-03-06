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

type ModifyCmd struct {
	Input string `arg:"" help:"source file to truncate"`
	Start uint32 `short:"f" optional:"" help:"start bytes to modify" default:"0"`
	Count uint   `short:"c" optional:"" help:"how many bytes to modify" default:"1"`
	Data  string `short:"D" optional:"" help:"data to write to file, encode in hex, no 0x prefix" default:"00"`
	Size  uint   `short:"s" optional:"" help:"size of data" default:"1"`
}

var client struct {
	Copy     CopyCmd     `cmd:"" aliases:"c" help:"copy part of a file"`
	Truncate TruncateCmd `cmd:"" aliases:"t" help:"truncate a file"`
	Modify   ModifyCmd   `cmd:"" aliases:"m" help:"modify a file"`
	Verbose  bool        `short:"v" help:"verbose" default:"false"`
	Debug    bool        `short:"d" help:"debug" default:"false"`
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
	if client.Truncate.Input != "" {
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

func bytesModify() {
	err := fsissors.BytesModify(client.Modify.Input, client.Modify.Start, client.Modify.Count, client.Modify.Size, client.Modify.Data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to modify bytes: %s\n", err.Error())
	}
}
func main() {
	ctx := kong.Parse(&client)
	fsissors.Verbose = client.Verbose
	fsissors.Debug = client.Debug
	switch ctx.Command() {
	case "copy <source> <target>":
		copyFile()
	case "truncate <input> <size>":
		truncateFile()
	case "modify <input>":
		bytesModify()
	default:
		_, _ = fmt.Fprintf(os.Stderr, "unknown command: %s\n", ctx.Command())
		_ = ctx.PrintUsage(true)
	}
}
