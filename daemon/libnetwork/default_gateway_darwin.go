package libnetwork

import (
	"runtime"

	"github.com/moby/moby/v2/daemon/libnetwork/types"
)

const libnGWNetwork = ""

func getPlatformOption() EndpointOption {
	return nil
}

func (c *Controller) createGWNetwork() (*Network, error) {
	return nil, types.NotImplementedErrorf("createGWNetwork is not implemented on " + runtime.GOOS)
}
