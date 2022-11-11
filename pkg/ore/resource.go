package ore

import (
	"context"

	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/containerutil"
	"github.com/frantjc/forge/internal/contaminate"
	"github.com/frantjc/forge/pkg/concourse"
	"github.com/frantjc/forge/pkg/forgeconcourse"
	"github.com/frantjc/go-fn"
)

func (o *Resource) Liquify(ctx context.Context, containerRuntime forge.ContainerRuntime, drains *forge.Drains) (*forge.Metal, error) {
	image, err := forgeconcourse.PullImageForResourceType(ctx, containerRuntime, o.GetResourceType())
	if err != nil {
		return nil, err
	}

	containerConfig := forgeconcourse.ResourceToConfig(o.GetResource(), o.GetResourceType(), o.GetMethod())
	containerConfig.Mounts = contaminate.OverrideWithMountsFrom(ctx, containerConfig.GetMounts()...)

	container, err := containerutil.CreateSleepingContainer(ctx, containerRuntime, image, containerConfig)
	if err != nil {
		return nil, err
	}
	defer container.Stop(ctx)   //nolint:errcheck
	defer container.Remove(ctx) //nolint:errcheck

	exitCode, err := container.Exec(ctx, containerConfig, forgeconcourse.NewStreams(drains, &concourse.Input{
		Params: fn.Ternary(
			o.GetMethod() == forgeconcourse.MethodCheck,
			nil, o.GetParams(),
		),
		Source: o.GetResource().GetSource(),
		Version: fn.Ternary(
			o.GetMethod() == forgeconcourse.MethodPut,
			nil, o.GetVersion(),
		),
	}))
	if err != nil {
		return nil, err
	}

	return &forge.Metal{
		ExitCode: int64(exitCode),
	}, nil
}
