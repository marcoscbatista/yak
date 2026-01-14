# Yak

> A fast and simple CLI tool to copy file contents directly to your clipboard

**Yak** is a lightweight command-line utility written in Go that streamlines your workflow by copying file contents to the system clipboard with a single command. Inspired by Vim's "yank" command, Yak brings the same simplicity to your terminal.

## Installation

### Download Pre-compiled Binaries

Download the latest release for your platform from the [Releases](https://github.com/marcoscbatista/yak/releases) page.

Available packages:
- **Linux**: DEB, RPM, APK, Arch packages
- **macOS**: Universal binary
- **Windows**: Executable

### Install from Source
```bash
git clone https://github.com/marcoscbatista/yak.git
cd yak
go mod tidy
go build -o yak ./cmd/yak
```

Then move the binary to your PATH:
```bash
# Linux/macOS
sudo mv yak /usr/local/bin/

# Or add to your user bin
mv yak ~/.local/bin/
```

## Usage

Copy any file content to clipboard:
```bash
yak path/to/file.txt
```

### Examples
```bash
# Copy a configuration file
yak ~/.bashrc

# Copy source code
yak main.go

# Copy from relative paths
yak ./README.md

# Copy multi-line files
yak large-log-file.log
```

The content is now in your system clipboard, ready to paste anywhere!

## Use Cases

- Copy configuration files for documentation
- Transfer log contents for debugging
- Paste file contents into web forms or chat applications
- Streamline your development workflow

## Contributing

Contributions are welcome! Feel free to:

- Report bugs by opening an [issue](https://github.com/marcoscbatista/yak/issues)
- Suggest new features
- Submit pull requests

## License

Distributed under the **GPL-3.0 License**. See [LICENSE](LICENSE) for more information.

---
