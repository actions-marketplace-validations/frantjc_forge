package command

import (
	"os"
	"strconv"

	"github.com/docker/docker/client"
	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/contaminate"
	"github.com/frantjc/forge/internal/hostfs"
	"github.com/frantjc/forge/pkg/forgeactions"
	"github.com/frantjc/forge/pkg/githubactions"
	"github.com/frantjc/forge/pkg/ore"
	"github.com/frantjc/forge/pkg/runtime/docker"
	"github.com/spf13/cobra"
)

func NewUse() *cobra.Command {
	var (
		env, with map[string]string
		cmd       = &cobra.Command{
			Use:           "use",
			Short:         "Use a GitHub Action",
			Args:          cobra.ExactArgs(1),
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(cmd *cobra.Command, args []string) error {
				var (
					ctx = cmd.Context()
					wd  = WorkdirFrom(ctx)
				)

				globalContext, err := githubactions.NewGlobalContextFromPath(ctx, wd)
				if err != nil {
					return err
				}

				if verbosity, _ := strconv.Atoi(cmd.Flag("verbose").Value.String()); verbosity > 0 {
					globalContext.SecretsContext[githubactions.SecretActionsStepDebug] = githubactions.SecretDebugValue
				}

				for _, dir := range []string{hostfs.RunnerTmp, hostfs.RunnerToolcache} {
					if err = os.MkdirAll(dir, 0o755); err != nil {
						return err
					}
				}

				c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
				if err != nil {
					return err
				}

				_, err = forge.NewFoundry(docker.New(c)).Process(
					contaminate.WithMounts(ctx, []*forge.Mount{
						{
							Source:      wd,
							Destination: forgeactions.DefaultWorkspace,
						},
						{
							Source:      hostfs.RunnerTmp,
							Destination: forgeactions.DefaultRunnerTemp,
						},
						{
							Source:      hostfs.RunnerToolcache,
							Destination: forgeactions.DefaultRunnerToolCache,
						},
					}...),
					&ore.Action{
						Uses:          args[0],
						With:          with,
						Env:           env,
						GlobalContext: globalContext,
					},
					forge.StdDrains(),
				)
				return err
			},
		}
	)

	cmd.Flags().StringToStringVarP(&with, "env", "e", nil, "env values")
	cmd.Flags().StringToStringVarP(&with, "with", "w", nil, "with values")
	cmd.Flags().StringVar(&forgeactions.Node12ImageReference, "node12-image", forgeactions.DefaultNode12ImageReference, "node12 image")
	cmd.Flags().StringVar(&forgeactions.Node16ImageReference, "node16-image", forgeactions.DefaultNode16ImageReference, "node16 image")

	return cmd
}
