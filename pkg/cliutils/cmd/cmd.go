package cmd

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/izumin5210/clicontrib"
	"github.com/izumin5210/clicontrib/cbuild"
)

// NewClicontribCommand creates a new command object.
func NewClicontribCommand(
	cwd string,
	inReader io.Reader,
	outWriter io.Writer,
	errWriter io.Writer,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:           cbuild.Default.Name,
		Short:         "",
		Long:          "",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	clicontrib.HandleLogFlags(cmd)
	cmd.AddCommand(
		newLdflagsCommand(outWriter, errWriter),
	)

	return cmd
}
