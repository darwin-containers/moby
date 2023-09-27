package config // import "github.com/docker/docker/daemon/config"

import (
	"context"
	"github.com/containerd/containerd/v2/defaults"
	"github.com/containerd/log"
	"github.com/docker/docker/api/types/system"
)

const (
	minAPIVersion    string = "1.24"
	StockRuntimeName        = defaults.DefaultRuntime
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
	RemappedRoot string                    `json:"userns-remap,omitempty"`
}

func (conf *Config) GetExecRoot() string {
	return ""
}

func (conf *Config) GetInitPath() string {
	return ""
}

func (conf *Config) IsSwarmCompatible() error {
	return nil
}

func (conf *Config) ValidatePlatformConfig() error {
	if conf.MTU != 0 && conf.MTU != DefaultNetworkMtu {
		log.G(context.TODO()).Warn(`WARNING: MTU for the default network is not configurable on macOS, and this option will be ignored.`)
	}
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
