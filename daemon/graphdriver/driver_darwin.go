package graphdriver // import "github.com/docker/docker/daemon/graphdriver"

// List of drivers that should be used in an order
var priority = "vfs"

// GetFSMagic returns the filesystem id given the path.
func GetFSMagic(rootpath string) (FsMagic, error) {
	return FsMagicUnsupported, nil
}
