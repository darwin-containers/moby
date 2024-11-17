//go:build unix && !linux

package daemon

import "github.com/moby/moby/api/types/mount"

func (daemon *Daemon) validateBindDaemonRoot(m mount.Mount) (bool, error) {
	return false, nil
}
