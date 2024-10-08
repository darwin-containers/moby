package containerd

import (
	"context"

	"github.com/containerd/containerd/v2/core/mount"
	"github.com/containerd/containerd/v2/core/snapshots"
)

const remapSuffix = "-remap"

func (i *ImageService) copyAndUnremapRootFS(ctx context.Context, dst, src []mount.Mount) error {
	return nil
}

func (i *ImageService) remapSnapshot(ctx context.Context, snapshotter snapshots.Snapshotter, id string, parentSnapshot string) error {
	return nil
}

func (i *ImageService) unremapRootFS(ctx context.Context, mounts []mount.Mount) error {
	return nil
}
