package main

import (
	"fmt"
	"log"
	"os"

	"marcoscbatista/yak/internal/cli"
	"marcoscbatista/yak/internal/clipboard"
	"marcoscbatista/yak/internal/reader"
)

func run() int {
	cb, err := clipboard.NewClipboard()
	if err != nil {
		log.Fatal(err)
	}
	rd := reader.NewReader()
	app := cli.NewCli(cb, rd)

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}
