package ccmd

import (
	"fmt"

	"github.com/izumin5210/clicontrib/pkg/clog"
	"github.com/spf13/cobra"
)

// HandleLogFlags processes verbose and debug flags and initialize a logger.
func HandleLogFlags(cmd *cobra.Command) {
	var (
		debugEnabled, verboseEnabled bool
	)

	cmd.PersistentFlags().BoolVar(
		&debugEnabled,
		"debug",
		false,
		fmt.Sprintf("Debug level output"),
	)
	cmd.PersistentFlags().BoolVarP(
		&verboseEnabled,
		"verbose",
		"v",
		false,
		fmt.Sprintf("Verbose level output"),
	)
	cobra.OnInitialize(func() {
		if debugEnabled {
			clog.InitDebug()
		} else if verboseEnabled {
			clog.InitVerbose()
		}
	})
}
