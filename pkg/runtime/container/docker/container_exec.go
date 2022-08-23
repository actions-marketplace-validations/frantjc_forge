package docker

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/frantjc/forge"
	"github.com/moby/term"
)

func (c *Container) Exec(ctx context.Context, config *forge.ContainerConfig, streams *forge.Streams) (int, error) {
	var (
		stdin          io.Reader
		stdout, stderr io.Writer
		tty            bool
		detachKeys     string
	)
	if streams != nil {
		stdin = streams.In
		stdout = streams.Out
		stderr = streams.Err
		tty = streams.Tty
		if tty {
			stderr = stdout
		}
		detachKeys = streams.DetachKeys
	}

	idr, err := c.Client.ContainerExecCreate(ctx, c.ID, types.ExecConfig{
		User:         config.User,
		Privileged:   config.Privileged,
		Env:          config.Env,
		WorkingDir:   config.Cwd,
		Cmd:          append(config.Entrypoint, config.Cmd...),
		Tty:          tty,
		DetachKeys:   detachKeys,
		AttachStdin:  stdin != nil,
		AttachStdout: stdout != nil,
		AttachStderr: stderr != nil,
	})
	if err != nil {
		return -1, err
	}

	hjr, err := c.Client.ContainerExecAttach(ctx, idr.ID, types.ExecStartCheck{
		Tty: tty,
	})
	if err != nil {
		return -1, err
	}
	defer hjr.Close()

	errC := make(chan error, 1)
	go func() {
		var err error
		if tty {
			_, err = io.Copy(stdout, hjr.Reader)
		} else {
			_, err = stdcopy.StdCopy(
				stdout,
				stderr,
				hjr.Reader,
			)
		}
		errC <- err
	}()

	if stdin != nil {
		if detachKeys != "" {
			detachKeysB, err := term.ToBytes(detachKeys)
			if err != nil {
				return -1, err
			}

			stdin = term.NewEscapeProxy(stdin, detachKeysB)
		}

		go func() {
			if _, err = io.Copy(hjr.Conn, stdin); err != nil {
				errC <- err
			}

			errC <- hjr.CloseWrite()
		}()
	}

	select {
	case err = <-errC:
		if _, ok := err.(term.EscapeError); ok {
			err = nil
		}
	case <-ctx.Done():
		err = ctx.Err()
	}

	cei, inspectErr := c.ContainerExecInspect(ctx, idr.ID)
	if inspectErr != nil {
		return -1, inspectErr
	}

	return cei.ExitCode, err
}
