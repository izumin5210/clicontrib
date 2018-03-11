package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"

	"github.com/izumin5210/clicontrib/cbuild"
	"github.com/izumin5210/clicontrib/pkg/cliutils"
)

func newLdflagsCommand(cfg *cliutils.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "ldflags NAME VERSION",
		Short:         "",
		Long:          "",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(_ *cobra.Command, args []string) error {
			buildCfg := cbuild.Config{}

			var (
				cmd *exec.Cmd
				out []byte
				err error
			)

			buildCfg.Name, err = cfg.Name()
			if err != nil {
				return err
			}

			buildCfg.Version, err = cfg.Version()
			if err != nil {
				return err
			}

			cmd = exec.Command("git", "rev-parse", "HEAD")
			cmd.Stderr = cfg.ErrWriter
			out, err = cmd.Output()
			if err != nil {
				return err
			}
			buildCfg.GitCommit = string(out)

			cmd = exec.Command("git", "describe", "--exact-match", "--abbrev=0", "--tags")
			out, err = cmd.Output()
			if err == nil {
				buildCfg.GitTag = string(out)
			}

			cmd = exec.Command("git", "describe", "--abbrev=0", "--tags")
			out, err = cmd.Output()
			if err == nil {
				buildCfg.GitNearestTag = string(out)
			}

			cmd = exec.Command("git", "status", "--porcelain")
			cmd.Stderr = cfg.ErrWriter
			out, err = cmd.Output()
			if err != nil {
				return err
			}
			if len(out) == 0 {
				buildCfg.GitTreeState = cbuild.TreeStateClean
			} else {
				buildCfg.GitTreeState = cbuild.TreeStateDirty
			}

			buildCfg.BuildTime = time.Now().UTC()

			fmt.Fprintln(cfg.OutWriter, buildCfg.Ldflags())

			return nil
		},
	}

	return cmd
}
