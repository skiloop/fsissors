# fsissors

A file utility designed for editing large files.

## Installation

To install the `fsissors` command-line tool, run the following command:

```bash
go install github.com/skiloop/fsissors@latest
```

Make sure that your `$GOPATH/bin` (or `%USERPROFILE%\go\bin` on Windows) is added to your system's `PATH` environment variable so that the `fsissors` command can be executed globally.

## Features

- **Copy**: Copy a portion of a file to another file.
- **Truncate**: Truncate a file to a specified size.
- **Modify**: Modify specific bytes in a file.

## Usage

The `fsissors` tool provides the following commands:

### Copy a File

Copy a portion of a file to another file.

```bash
fsissors copy <source> <target> [--from=<offset>] [--size=<size>] [--buffer-size=<buffer-size>]
```

#### Options:
- `<source>`: The source file to copy from.
- `<target>`: The target file to copy to.
- `--from`: Offset to start copying from (default: `0`).
- `--size`: Number of bytes to copy (default: `0`, meaning copy to the end of the file).
- `--buffer-size`: Buffer size in bytes for copying (default: `1024`).

#### Example:
```bash
fsissors copy input.txt output.txt --from=100 --size=500 --buffer-size=2048
```

### Truncate a File

Truncate a file to a specified size.

```bash
fsissors truncate <input> <size>
```

#### Options:
- `<input>`: The file to truncate.
- `<size>`: The size to truncate the file to. If `<size>` is negative, the front part of the file will be removed.

#### Example:
```bash
fsissors truncate largefile.txt -1000
```

### Modify Bytes in a File

Modify specific bytes in a file.

```bash
fsissors modify <input> [--start=<offset>] [--count=<count>] [--data=<hex-data>] [--size=<size>]
```

#### Options:
- `<input>`: The file to modify.
- `--start`: The starting byte offset to modify (default: `0`).
- `--count`: The number of bytes to modify (default: `1`).
- `--data`: The data to write, encoded in hexadecimal (default: `00`).
- `--size`: The size of the data to write (default: `1`).

#### Example:
```bash
fsissors modify file.bin --start=10 --count=4 --data=FFEE --size=2
```

## Debugging and Verbose Mode

You can enable verbose or debug mode for more detailed output:

- `--verbose` or `-v`: Enable verbose output.
- `--debug` or `-d`: Enable debug mode.

#### Example:
```bash
fsissors copy input.txt output.txt --verbose
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
```