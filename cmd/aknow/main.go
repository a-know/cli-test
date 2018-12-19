package main

import (
	"fmt"
	"os"

	"github.com/a-know/cli-test/pkg/aknow"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cmd := aknow.NewHelloCommand()
	return errors.WithStack(cmd.Execute())
}
