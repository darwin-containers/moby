//go:build !darwin && !windows

package appdefaults

const (
	Address         = "unix:///run/buildkit/buildkitd.sock"
	traceSocketPath = "/run/buildkit/otel-grpc.sock"
)
