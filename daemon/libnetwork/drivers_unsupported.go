//go:build !linux && !windows && !darwin

package libnetwork

func registerNetworkDrivers(cfg *config.Config, r driverapi.Registerer, store *datastore.Store, pms *drvregistry.PortMappers) error {
	return nil
}
