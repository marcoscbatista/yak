// Package cli implements the command-line interface logic.
package cli

import (
	"errors"
	"fmt"
	"path/filepath"

	"marcoscbatista/yak/internal/clipboard"
	"marcoscbatista/yak/internal/reader"
)

type CLI struct {
	clipboard clipboard.Clipboard
	reader    reader.Reader
}

func NewCli(cb clipboard.Clipboard, rd reader.Reader) *CLI {
	return &CLI{
		clipboard: cb,
		reader:    rd,
	}
}

func (c *CLI) Run(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: yak <dir>")
	}

	rawPath := args[1]

	path, err := filepath.Abs(rawPath)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	content, err := c.reader.Read(path)
	if err != nil {
		return err
	}

	return c.clipboard.Write(content)
}
