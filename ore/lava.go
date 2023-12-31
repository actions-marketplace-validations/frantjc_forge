package ore

import (
	"bytes"
	"context"

	"github.com/frantjc/forge"
	"github.com/frantjc/forge/internal/contaminate"
)

// Lava is an Ore representing two Ores of which the
// stdout of the first is piped to the stdin of the second.
type Lava struct {
	From forge.Ore `json:"from,omitempty"`
	To   *Pure     `json:"to,omitempty"`
}

func (o *Lava) Liquify(ctx context.Context, containerRuntime forge.ContainerRuntime, drains *forge.Drains) (err error) {
	buf := new(bytes.Buffer)
	if err = o.From.Liquify(ctx, containerRuntime, &forge.Drains{
		Out: buf,
		Err: drains.Err,
		Tty: drains.Tty,
	}); err != nil {
		return
	}

	return o.To.Liquify(contaminate.WithInput(ctx, buf.Bytes()), containerRuntime, drains)
}

func (o *Lava) GetFrom() forge.Ore {
	return o.From
}

func (o *Lava) GetTo() forge.Ore {
	return o.To
}
