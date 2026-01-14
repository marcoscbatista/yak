// Package main is the entry point for the yak CLI.
package main

import (
	"log"
	"os"

	"marcoscbatista/yak/internal/cli"
	"marcoscbatista/yak/internal/clipboard"
	"marcoscbatista/yak/internal/reader"
)

func main() {
	cb, err := clipboard.NewClipboard()
	if err != nil {
		log.Fatal(err)
	}
	rd := reader.NewReader()
	cli := cli.NewCli(cb, rd)
	os.Exit(run(cli))
}
