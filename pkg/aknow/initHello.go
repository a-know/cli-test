package aknow

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newInitHelloCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "hello",
		Args: cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			fmt.Fprintln(os.Stdout, fmt.Sprintf("hello, %s!!", args[0]))
			return nil
		},
	}

	return cmd
}
