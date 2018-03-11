package main

import (
	"fmt"
	"os"

	"github.com/izumin5210/clicontrib/pkg/clicontrib/cmd"
)

var (
	inReader  = os.Stdin
	outWriter = os.Stdout
	errWriter = os.Stderr
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = cmd.NewClicontribCommand(
		cwd,
		inReader,
		outWriter,
		errWriter,
	).Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
