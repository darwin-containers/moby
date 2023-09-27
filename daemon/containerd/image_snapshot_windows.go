package containerd

import (
	"context"

	"github.com/containerd/containerd/v2/snapshots"
)

func (i *ImageService) remapSnapshot(ctx context.Context, snapshotter snapshots.Snapshotter, id string, parentSnapshot string) error {
	return nil
}
