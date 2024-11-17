package safepath

import (
	"context"
	"runtime"

	"github.com/pkg/errors"
)

func Join(_ context.Context, path, subpath string) (*SafePath, error) {
	return nil, errors.New("safepath.Join is not supported on " + runtime.GOOS)
}
