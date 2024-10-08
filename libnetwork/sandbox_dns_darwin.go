package libnetwork

import (
	"context"
	"github.com/docker/docker/libnetwork/etchosts"
)

// Stub implementations for DNS related functions

func (sb *Sandbox) setupResolutionFiles(_ context.Context) error {
	return nil
}

func (sb *Sandbox) restoreHostsPath() {}

func (sb *Sandbox) restoreResolvConfPath() {}

func (sb *Sandbox) updateHostsFile(ctx context.Context, ifaceIPs []string) error {
	return nil
}

func (sb *Sandbox) deleteHostsEntries(recs []etchosts.Record) {}

func (sb *Sandbox) updateDNS(ipv6Enabled bool) error {
	return nil
}
