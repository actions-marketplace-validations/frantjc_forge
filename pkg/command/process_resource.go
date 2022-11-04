package command

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/docker/docker/client"
	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/contaminate"
	"github.com/frantjc/forge/pkg/forgeconcourse"
	"github.com/frantjc/forge/pkg/ore"
	"github.com/frantjc/forge/pkg/runtime/docker"
)

func processResource(ctx context.Context, method, name string, params, version map[string]string) error {
	var (
		_      = forge.LoggerFrom(ctx)
		config = &forgeconcourse.Config{}
		wd     = WorkdirFrom(ctx)
		err    error
	)

	for _, filename := range []string{"forge.json"} {
		if file, err := os.Open(filepath.Join(wd, filename)); err == nil {
			if err = json.NewDecoder(file).Decode(config); err == nil {
				break
			}
		}
	}
	if err != nil {
		return err
	}

	o := &ore.Resource{
		Method:  method,
		Version: version,
		Params:  params,
	}
	for _, r := range config.GetResources() {
		if r.GetName() == name {
			o.Resource = r
		}
	}
	if o.GetResource() == nil {
		return fmt.Errorf("resource not found: %s", name)
	}

	for _, t := range config.GetResourceTypes() {
		if t.GetName() == o.GetResource().GetType() {
			o.ResourceType = t
		}
	}
	if o.GetResourceType() == nil {
		return fmt.Errorf("resource type not found: %s", o.GetResource().GetType())
	}

	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	_, err = forge.NewFoundry(docker.New(c)).Process(
		contaminate.WithMounts(ctx, &forge.Mount{
			Source:      wd,
			Destination: filepath.Join(forgeconcourse.DefaultRootPath, o.GetResource().GetName()),
		}), o, forge.StdDrains(),
	)
	return err
}
