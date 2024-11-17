//go:build unix && !linux

package libnetwork

import (
	"context"
	"time"

	"github.com/moby/moby/v2/daemon/libnetwork/ipams/defaultipam"
)

type platformNetwork struct{} //nolint:nolintlint,unused // only populated on windows

// Stub implementations for DNS related functions

func (n *Network) startResolver() {
}

func addEpToResolver(
	ctx context.Context,
	netName, epName string,
	config *containerConfig,
	epIface *EndpointInterface,
	resolvers []*Resolver,
) error {
	return nil
}

func deleteEpFromResolver(epName string, epIface *EndpointInterface, resolvers []*Resolver) error {
	return nil
}

func defaultIpamForNetworkType(networkType string) string {
	return defaultipam.DriverName
}

func (n *Network) validatedAdvertiseAddrNMsgs() (*int, error) {
	return nil, nil
}

func (n *Network) validatedAdvertiseAddrInterval() (*time.Duration, error) {
	return nil, nil
}

func (n *Network) IsPruneable() bool {
	return false
}
