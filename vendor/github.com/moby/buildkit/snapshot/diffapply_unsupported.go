//go:build !linux

package snapshot

import (
	"context"
	"runtime"

	"github.com/containerd/containerd/v2/leases"
	"github.com/containerd/containerd/v2/snapshots"
	"github.com/pkg/errors"
)

func (sn *mergeSnapshotter) diffApply(ctx context.Context, dest Mountable, diffs ...Diff) (_ snapshots.Usage, rerr error) {
	return snapshots.Usage{}, errors.New("diffApply not yet supported on " + runtime.GOOS)
}

func needsUserXAttr(ctx context.Context, sn Snapshotter, lm leases.Manager) (bool, error) {
	return false, errors.New("needs userxattr not supported on " + runtime.GOOS)
}
