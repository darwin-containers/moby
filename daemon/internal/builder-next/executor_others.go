//go:build !linux && !windows

package buildkit

import (
	"context"

	"github.com/containerd/log"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/oci"
	"github.com/moby/buildkit/solver/llbsolver/cdidevices"
	"github.com/moby/moby/v2/daemon/libnetwork"
	"github.com/moby/sys/user"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func newExecutor(_, _ string, _ *libnetwork.Controller, _ *oci.DNSConfig, _ bool, _ user.IdentityMapping, _ string, _ *cdidevices.Manager, _, _ string) (executor.Executor, error) {
	return &stubExecutor{}, nil
}

func (iface *lnInterface) Set(s *specs.Spec) error {
	<-iface.ready
	if iface.err != nil {
		log.G(context.TODO()).WithError(iface.err).Error("failed to set networking spec")
		return iface.err
	}

	return nil
}
