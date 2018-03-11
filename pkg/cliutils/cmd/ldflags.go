package cmd

import (
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/spf13/cobra"

	"github.com/izumin5210/clicontrib/cbuild"
)

func newLdflagsCommand(outWriter io.Writer, errWriter io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "ldflags NAME VERSION",
		Short:         "",
		Long:          "",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, args []string) error {
			cfg := cbuild.Config{}

			cfg.Name = args[0]
			cfg.Version = args[1]

			var (
				out []byte
				err error
			)

			cmd := exec.Command("git", "rev-parse", "HEAD")
			cmd.Stderr = errWriter
			out, err = cmd.Output()
			if err != nil {
				return err
			}
			cfg.GitCommit = string(out)

			cmd = exec.Command("git", "describe", "--exact-match", "--abbrev=0", "--tags")
			out, err = cmd.Output()
			if err == nil {
				cfg.GitTag = string(out)
			}

			cmd = exec.Command("git", "describe", "--abbrev=0", "--tags")
			out, err = cmd.Output()
			if err == nil {
				cfg.GitNearestTag = string(out)
			}

			cmd = exec.Command("git", "status", "--porcelain")
			cmd.Stderr = errWriter
			out, err = cmd.Output()
			if err != nil {
				return err
			}
			if len(out) == 0 {
				cfg.GitTreeState = cbuild.TreeStateClean
			} else {
				cfg.GitTreeState = cbuild.TreeStateDirty
			}

			cfg.BuildTime = time.Now().UTC()

			fmt.Fprintln(outWriter, cfg.Ldflags())

			return nil
		},
	}

	return cmd
}
