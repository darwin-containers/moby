//go:build unix && !linux

package platform

func possibleCPUs() []int {
	// not implemented
	return nil
}
