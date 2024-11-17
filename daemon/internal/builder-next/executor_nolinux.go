//go:build !linux

package buildkit

import (
	ctd "github.com/containerd/containerd/v2/client"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/oci"
	"github.com/moby/buildkit/solver/llbsolver/cdidevices"
	"github.com/moby/moby/v2/daemon/libnetwork"
	"github.com/moby/sys/user"
)

func newExecutorGD(root, cgroupParent string, net *libnetwork.Controller, dnsConfig *oci.DNSConfig, rootless bool, idmap user.IdentityMapping, apparmorProfile string, cdiManager *cdidevices.Manager, containerdClient *ctd.Client) (executor.Executor, error) {
	return newExecutor(root, cgroupParent, net, dnsConfig, rootless, idmap, apparmorProfile, cdiManager, containerdClient)
}
