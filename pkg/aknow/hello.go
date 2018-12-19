package aknow

import "github.com/spf13/cobra"

func NewHelloCommand() *cobra.Command {
	var (
		flagVerbose bool
	)

	cmd := &cobra.Command{
		Use: "hello",
	}

	cmd.AddCommand(
		newInitHelloCommand(),
	)

	cmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose log")

	return cmd
}
