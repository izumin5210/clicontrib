package ccmd

import (
	"time"

	"github.com/izumin5210/clicontrib/pkg/cbuild"
	"github.com/spf13/cobra"
)

// NewVersionCommand returns a new command object for printing version information.
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:           "version",
		Short:         "Print version information",
		Long:          "Print version information",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, _ []string) {
			cfg := cbuild.Default
			releaseType := "stable"
			if cfg.GitTag != cfg.GitNearestTag {
				releaseType = "canary"
			}
			fmtStr := "%s %s %s (%s %s)\n"
			if cfg.GitTreeState != cbuild.TreeStateClean {
				fmtStr = "%s %s %s (%s %s dirty)\n"
			}
			cmd.Printf(
				fmtStr,
				cfg.Name,
				cfg.Version,
				releaseType,
				cfg.BuildTime.Format(time.RFC3339),
				cfg.GitCommit[:7],
			)
		},
	}
}
