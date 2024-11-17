package command

import (
	"github.com/moby/moby/v2/daemon/config"
	"github.com/moby/moby/v2/daemon/pkg/opts"
	"github.com/spf13/pflag"
)

func installConfigFlags(conf *config.Config, flags *pflag.FlagSet) error {
	installCommonConfigFlags(conf, flags)

	flags.Var(opts.NewNamedRuntimeOpt("runtimes", &conf.Runtimes, config.StockRuntimeName), "add-runtime", "Register an additional OCI compatible runtime")
	flags.StringVarP(&conf.SocketGroup, "group", "G", "docker", "Group for the unix socket")

	return nil
}
