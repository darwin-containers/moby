package config

import (
	"context"
	"fmt"

	"github.com/containerd/containerd/v2/defaults"
	"github.com/containerd/log"
	"github.com/moby/moby/api/types/system"
)

const (
	StockRuntimeName = defaults.DefaultRuntime
)

type BridgeConfig struct {
	DefaultBridgeConfig
}

type DefaultBridgeConfig struct {
	commonBridgeConfig

	// MTU is not actually used on Windows, but the --mtu option has always
	// been there on Windows (but ignored).
	MTU int `json:"mtu,omitempty"`
}

type Config struct {
	CommonConfig

	CgroupParent string
	ResolvConf   string                    `json:"resolv-conf,omitempty"`
	Runtimes     map[string]system.Runtime `json:"runtimes,omitempty"`
}

func (conf *Config) GetExecRoot() string {
	return conf.ExecRoot
}

func (conf *Config) GetInitPath() string {
	return ""
}

func (conf *Config) IsSwarmCompatible() error {
	return nil
}

func (conf *Config) IsRootless() bool {
	return false
}

func (conf *Config) GetResolvConf() string {
	return conf.ResolvConf
}

func setPlatformDefaults(cfg *Config) error {
	cfg.Root = "/var/lib/docker"
	cfg.ExecRoot = "/var/run/docker"
	cfg.Pidfile = "/var/run/docker.pid"

	cfg.Runtimes = make(map[string]system.Runtime)
	return nil
}

// validatePlatformConfig checks if any platform-specific configuration settings are invalid.
func validatePlatformConfig(conf *Config) error {
	if conf.MTU != 0 && conf.MTU != DefaultNetworkMtu {
		log.G(context.TODO()).Warn(`WARNING: MTU for the default network is not configurable on Windows, and this option will be ignored.`)
	}
	return nil
}

// validatePlatformExecOpt validates if the given exec-opt and value are valid
// for the current platform.
func validatePlatformExecOpt(opt, value string) error {
	switch opt {
	case "isolation":
		// TODO(thaJeztah): add validation that's currently in Daemon.setDefaultIsolation()
		return nil
	case "native.cgroupdriver":
		return fmt.Errorf("option '%s' is only supported on linux", opt)
	default:
		return fmt.Errorf("unknown option: '%s'", opt)
	}
}
