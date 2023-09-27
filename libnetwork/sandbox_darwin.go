package libnetwork

import (
	"context"
	"github.com/docker/docker/libnetwork/osl"
	"github.com/docker/docker/libnetwork/types"
)

type containerConfigOS struct{} //nolint:nolintlint,unused // only populated on windows

func (sb *Sandbox) populateNetworkResources(context.Context, *Endpoint) error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func (sb *Sandbox) updateGateway(*Endpoint) error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func (sb *Sandbox) restoreOslSandbox() error {
	// not implemented on Darwin (Sandbox.osSbox is always nil)
	return nil
}

func releaseOSSboxResources(*osl.Namespace, *Endpoint) {}

func (sb *Sandbox) Statistics() (map[string]*types.InterfaceStatistics, error) {
	return nil, nil
}
