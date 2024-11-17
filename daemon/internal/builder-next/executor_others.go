//go:build !linux && !windows

package buildkit

import (
	"context"

	ctd "github.com/containerd/containerd/v2/client"
	"github.com/containerd/log"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/containerdexecutor"
	"github.com/moby/buildkit/executor/oci"
	"github.com/moby/buildkit/solver/llbsolver/cdidevices"
	"github.com/moby/buildkit/util/network/netproviders"
	"github.com/moby/moby/v2/daemon/libnetwork"
	"github.com/moby/sys/user"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func newExecutor(root, cgroupParent string, net *libnetwork.Controller, dnsConfig *oci.DNSConfig, rootless bool, idmap user.IdentityMapping, apparmorProfile string, cdiManager *cdidevices.Manager, containerdClient *ctd.Client) (executor.Executor, error) {
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

func (iface *lnInterface) Set(s *specs.Spec) error {
	<-iface.ready
	if iface.err != nil {
		log.G(context.TODO()).WithError(iface.err).Error("failed to set networking spec")
		return iface.err
	}

	return nil
}
