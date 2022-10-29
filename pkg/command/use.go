package command

import (
	"os"

	"github.com/docker/docker/client"
	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/contaminate"
	hfs "github.com/frantjc/forge/internal/hostfs"
	fa "github.com/frantjc/forge/pkg/forgeactions"
	"github.com/frantjc/forge/pkg/github/actions"
	"github.com/frantjc/forge/pkg/ore"
	"github.com/frantjc/forge/pkg/runtime/container/docker"
	"github.com/spf13/cobra"
)

func NewUse() *cobra.Command {
	var (
		env  = map[string]string{}
		with = map[string]string{}
		cmd  = &cobra.Command{
			Use:  "use",
			Args: cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				ctx := cmd.Context()

				wd, err := os.Getwd()
				if err != nil {
					cmd.PrintErrln(err)
					return
				}

				globalContext, err := actions.NewGlobalContextFromPath(ctx, wd)
				if err != nil {
					cmd.PrintErrln(err)
					return
				}

				c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
				if err != nil {
					cmd.PrintErrln(err)
					return
				}

				for _, dir := range []string{hfs.RunnerTmp, hfs.RunnerToolcache} {
					if err = os.MkdirAll(dir, 0755); err != nil {
						cmd.PrintErrln(err)
						return
					}
				}

				if _, err = forge.NewFoundry(docker.New(c)).Process(
					contaminate.WithMounts(ctx, []*forge.Mount{
						{
							Source:      wd,
							Destination: fa.DefaultWorkspace,
						},
						{
							Source:      hfs.RunnerTmp,
							Destination: fa.DefaultRunnerTemp,
						},
						{
							Source:      hfs.RunnerToolcache,
							Destination: fa.DefaultRunnerToolCache,
						},
					}...),
					&ore.Action{
						Uses:          args[0],
						With:          with,
						Env:           env,
						GlobalContext: globalContext,
					},
					forge.StdDrains(),
				); err != nil {
					cmd.PrintErrln(err)
				}
			},
		}
	)

	cmd.Flags().StringToStringVarP(&with, "env", "e", make(map[string]string), "env values")
	cmd.Flags().StringToStringVarP(&with, "with", "w", make(map[string]string), "with values")

	return cmd
}
