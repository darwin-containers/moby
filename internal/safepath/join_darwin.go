package safepath

import (
	"context"
	"github.com/pkg/errors"
	"runtime"
)

func Join(_ context.Context, path, subpath string) (*SafePath, error) {
	return nil, errors.New("safepath.Join is not supported on " + runtime.GOOS)
}
