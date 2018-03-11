package cmd

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/izumin5210/clicontrib"
	"github.com/izumin5210/clicontrib/cbuild"
	"github.com/izumin5210/clicontrib/pkg/cliutils"
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

	cfg := cliutils.NewConfig(
		inReader,
		outWriter,
		errWriter,
	)

	var cfgFile string
	cobra.OnInitialize(func() { cfg.Init(cfgFile) })
	clicontrib.HandleLogFlags(cmd)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "./"+cbuild.Default.Name+".toml", "config file")

	cmd.AddCommand(
		newLdflagsCommand(cfg),
	)

	return cmd
}
