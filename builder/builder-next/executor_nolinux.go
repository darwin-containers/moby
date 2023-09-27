//go:build !linux

package buildkit

import (
	containerd "github.com/containerd/containerd/v2/client"
	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/libnetwork"
	"github.com/docker/docker/pkg/idtools"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/containerdexecutor"
	"github.com/moby/buildkit/executor/oci"
	"github.com/moby/buildkit/util/network/netproviders"
)

func newExecutor(root string, cgroupParent string, containerdClient *containerd.Client, net *libnetwork.Controller, dnsConfig *oci.DNSConfig, rootless bool, idmap idtools.IdentityMapping, apparmorProfile string) (executor.Executor, error) {
	nc := netproviders.Opt{
		Mode: "host",
	}
	np, _, err := netproviders.Providers(nc)
	if err != nil {
		return nil, err
	}

	opts := containerdexecutor.ExecutorOptions{
		Client:           containerdClient,
		Root:             root,
		CgroupParent:     cgroupParent,
		NetworkProviders: np,
		DNSConfig:        dnsConfig,
		ApparmorProfile:  apparmorProfile,
		Runtime: &containerdexecutor.RuntimeInfo{
			Name: containerdClient.Runtime(),
		},
	}

	return containerdexecutor.New(opts), nil
}

func getDNSConfig(config.DNSConfig) *oci.DNSConfig {
	return nil
}
