// Package cli implements the command-line interface logic.
package cli

import (
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

func (c *CLI) Run(path string) error {
	content, err := c.reader.Read(path)
	if err != nil {
		return err
	}
	c.clipboard.Write(content)
	return nil
}
