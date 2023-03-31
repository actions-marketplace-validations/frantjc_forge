package forgeactions

import (
	"archive/tar"
	"context"
	"errors"
	"io"
	"strings"

	"github.com/frantjc/forge"
	"github.com/frantjc/forge/envconv"
	"github.com/frantjc/forge/githubactions"
	"golang.org/x/exp/maps"
)

func SetGlobalContextFromEnvFiles(ctx context.Context, globalContext *githubactions.GlobalContext, step string, container forge.Container) error {
	return DefaultMapping.SetGlobalContextFromEnvFiles(ctx, globalContext, step, container)
}

func (m *Mapping) SetGlobalContextFromEnvFiles(ctx context.Context, globalContext *githubactions.GlobalContext, step string, container forge.Container) error {
	_ = forge.LoggerFrom(ctx)
	globalContext = m.ConfigureGlobalContext(globalContext)

	rc, err := container.CopyFrom(ctx, m.GitHubPath)
	if err != nil {
		return err
	}
	defer rc.Close()

	r := tar.NewReader(rc)
	for {
		header, err := r.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		//nolint:gocritic
		switch header.Typeflag {
		case tar.TypeReg:

			switch {
			case strings.HasSuffix(m.GitHubOutputPath, header.Name):
				outputs, err := envconv.MapFromReader(io.LimitReader(r, header.Size))
				if err != nil {
					return err
				}

				if _, ok := globalContext.StepsContext[step]; !ok {
					globalContext.StepsContext[step] = &githubactions.StepContext{
						Outputs: outputs,
					}
				} else {
					maps.Copy(globalContext.StepsContext[step].Outputs, outputs)
				}
			case strings.HasSuffix(m.GitHubStatePath, header.Name):
				outputs, err := envconv.MapFromReader(io.LimitReader(r, header.Size))
				if err != nil {
					return err
				}

				for k, v := range outputs {
					globalContext.EnvContext["STATE_"+k] = v
				}
			}
		}
	}

	return rc.Close()
}
