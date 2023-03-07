package command

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/client"
	"github.com/frantjc/forge"
	"github.com/frantjc/forge/forgeconcourse"
	"github.com/frantjc/forge/internal/contaminate"
	"github.com/frantjc/forge/internal/hooks"
	"github.com/frantjc/forge/ore"
	"github.com/frantjc/forge/runtime/docker"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func newResource(method string, check bool) *cobra.Command {
	var (
		attach          bool
		conf, workdir   string
		version, params map[string]string
		cmd             = &cobra.Command{
			Use:           method,
			Short:         fmt.Sprintf("%s a Concourse resource", cases.Title(language.English).String(method)),
			Args:          cobra.ExactArgs(1),
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(cmd *cobra.Command, args []string) error {
				var (
					ctx    = cmd.Context()
					_      = forge.LoggerFrom(ctx)
					name   = args[0]
					config = &forgeconcourse.Config{}
					file   io.Reader
					err    error
					o      = &ore.Resource{
						Method:  method,
						Version: version,
						Params:  params,
					}
				)

				if cmd.Flag("conf").Changed {
					if file, err = os.Open(conf); err != nil {
						return err
					}
				} else {
					if file, err = os.Open(filepath.Join(workdir, conf)); err != nil {
						return err
					}
				}

				if err = yaml.NewDecoder(file).Decode(config); err != nil {
					return err
				}

				for _, r := range config.Resources {
					if r.Name == name {
						resource := r
						o.Resource = &resource
					}
				}
				if o.Resource == nil {
					return fmt.Errorf("resource not found: %s", name)
				}

				for _, t := range config.ResourceTypes {
					if t.Name == o.Resource.Type {
						resourceType := t
						o.ResourceType = &resourceType
					}
				}
				if o.ResourceType == nil {
					return fmt.Errorf("resource type not found: %s", o.Resource.Type)
				}

				c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
				if err != nil {
					return err
				}

				if attach {
					hooks.ContainerStarted.Listen(hookAttach(cmd))
				}

				return forge.NewFoundry(docker.New(c)).Process(
					contaminate.WithMounts(ctx, forge.Mount{
						Source:      workdir,
						Destination: filepath.Join(forgeconcourse.DefaultRootPath, o.Resource.Name),
					}),
					o,
					commandDrains(cmd),
				)
			},
		}
	)

	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	if !check {
		cmd.Flags().StringToStringVarP(&params, "params", "p", nil, "params for resource")
	}
	cmd.Flags().BoolVarP(&attach, "attach", "a", false, "attach to containers")
	cmd.Flags().StringToStringVarP(&version, "version", "v", nil, "version for resource")
	cmd.Flags().StringVarP(&conf, "conf", "c", "forge.yml", "config file for resource")
	_ = cmd.MarkFlagFilename("conf", "yaml", "yml", "json")
	cmd.Flags().StringVar(&workdir, "workdir", wd, "working directory for resource")
	_ = cmd.MarkFlagDirname("workdir")

	return cmd
}
