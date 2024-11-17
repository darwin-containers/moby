package config

type PlatformConfig struct {
}

func optionExecRoot(execRoot string) Option {
	return func(*Config) {}
}
