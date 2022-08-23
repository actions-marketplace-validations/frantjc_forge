package forge

import (
	"context"
	"fmt"
)

// Deposit is a cache for Metal
type Deposit interface {
	fmt.GoStringer

	Store(context.Context, Ore, *Metal) error
	Retrieve(context.Context, Ore) (*Metal, error)
}
