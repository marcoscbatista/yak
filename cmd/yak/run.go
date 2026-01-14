package main

import (
	"fmt"
	"os"
	"path/filepath"

	"marcoscbatista/yak/internal/cli"
)

func run(cli *cli.CLI) int {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: yak <dir>")
		return 1
	}
	dir := os.Args[1]

	if !filepath.IsAbs(dir) {
		abs, err := filepath.Abs(dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		dir = abs
	}
	fmt.Println(dir)
	if err := cli.Run(dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
