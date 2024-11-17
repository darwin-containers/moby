package libnetwork

import (
	"context"

	"github.com/moby/moby/v2/daemon/libnetwork/osl"
	"github.com/moby/moby/v2/daemon/libnetwork/types"
)

func releaseOSSboxResources(*osl.Namespace, *Endpoint) {}

func (sb *Sandbox) updateGateway(_, _ *Endpoint) error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func (sb *Sandbox) ExecFunc(func()) error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func (sb *Sandbox) releaseOSSbox() error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func (sb *Sandbox) restoreOslSandbox() error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

// NetnsPath is not implemented on Darwin (Sandbox.osSbox is always nil)
func (sb *Sandbox) NetnsPath() (path string, ok bool) {
	return "", false
}

func (sb *Sandbox) IPv6Enabled() (enabled, ok bool) {
	return false, true
}

func (sb *Sandbox) Statistics() (map[string]*types.InterfaceStatistics, error) {
	return nil, nil
}

func (sb *Sandbox) canPopulateNetworkResources() bool {
	return true
}

func (sb *Sandbox) populateNetworkResourcesOS(ctx context.Context, ep *Endpoint) error {
	return nil
}
