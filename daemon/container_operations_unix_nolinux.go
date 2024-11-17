//go:build unix && !linux

package daemon

import (
	"context"
	"errors"
	"runtime"

	"github.com/moby/moby/v2/daemon/config"
	"github.com/moby/moby/v2/daemon/container"
	"github.com/moby/moby/v2/daemon/libnetwork"
	"github.com/moby/moby/v2/daemon/network"
)

func (daemon *Daemon) addLegacyLinks(
	ctx context.Context,
	cfg *config.Config,
	ctr *container.Container,
	epConfig *network.EndpointSettings,
	sb *libnetwork.Sandbox,
) error {
	return nil
}

func (daemon *Daemon) createSecretsDir(ctr *container.Container) error {
	return errors.New("createSecretsDir is not supported on " + runtime.GOOS)
}

func (daemon *Daemon) remountSecretDir(ctr *container.Container) error {
	return errors.New("remountSecretDir is not supported on " + runtime.GOOS)
}

func (daemon *Daemon) setupIPCDirs(ctr *container.Container) error {
	return nil
}
