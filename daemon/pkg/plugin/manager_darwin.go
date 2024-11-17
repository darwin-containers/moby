package plugin

import (
	"fmt"

	"github.com/containerd/containerd/v2/core/mount"
	"golang.org/x/sys/unix"

	v2 "github.com/moby/moby/v2/daemon/pkg/plugin/v2"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func (pm *Manager) enable(p *v2.Plugin, c *controller, force bool) error {
	return fmt.Errorf("not implemented")
}

func (pm *Manager) initSpec(p *v2.Plugin) (*specs.Spec, error) {
	return nil, fmt.Errorf("not implemented")
}

func (pm *Manager) disable(p *v2.Plugin, c *controller) error {
	return fmt.Errorf("not implemented")
}

func (pm *Manager) restore(p *v2.Plugin, c *controller) error {
	return fmt.Errorf("not implemented")
}

// Shutdown plugins
func (pm *Manager) Shutdown() {
}

func recursiveUnmount(target string) error {
	return mount.UnmountRecursive(target, unix.MNT_FORCE)
}
