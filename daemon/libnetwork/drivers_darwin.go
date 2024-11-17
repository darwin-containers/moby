package libnetwork

import (
	"context"

	"github.com/moby/moby/v2/daemon/libnetwork/config"
	"github.com/moby/moby/v2/daemon/libnetwork/datastore"
	"github.com/moby/moby/v2/daemon/libnetwork/driverapi"
	"github.com/moby/moby/v2/daemon/libnetwork/drivers/host"
	"github.com/moby/moby/v2/daemon/libnetwork/drivers/null"
	"github.com/moby/moby/v2/daemon/libnetwork/drvregistry"
)

func registerNetworkDrivers(r driverapi.Registerer, cfg *config.Config, store *datastore.Store, pms *drvregistry.PortMappers) error {
	err := null.Register(r)
	if err != nil {
		return err
	}

	err = host.Register(r)
	if err != nil {
		return err
	}

	return nil
}

func registerPortMappers(ctx context.Context, r *drvregistry.PortMappers, cfg *config.Config) error {
	return nil
}
