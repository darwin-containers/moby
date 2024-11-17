package daemon

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/containerd/log"
	"github.com/moby/moby/v2/daemon/config"
	"github.com/moby/moby/v2/daemon/container"
	"github.com/moby/moby/v2/daemon/libnetwork"
	"github.com/moby/moby/v2/daemon/libnetwork/drivers/bridge"
	"github.com/moby/moby/v2/daemon/network"
	"github.com/moby/sys/mount"
	"github.com/moby/sys/user"
	"github.com/opencontainers/selinux/go-selinux/label"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"golang.org/x/sys/unix"
)

func (daemon *Daemon) addLegacyLinks(
	ctx context.Context,
	cfg *config.Config,
	ctr *container.Container,
	epConfig *network.EndpointSettings,
	sb *libnetwork.Sandbox,
) error {
	ctx, span := otel.Tracer("").Start(ctx, "daemon.addLegacyLinks")
	defer span.End()

	if epConfig.EndpointID == "" {
		return nil
	}

	children := daemon.linkIndex.children(ctr)
	var parents map[string]*container.Container
	if !cfg.DisableBridge && ctr.HostConfig.NetworkMode.IsPrivate() {
		parents = daemon.linkIndex.parents(ctr)
	}
	if len(children) == 0 && len(parents) == 0 {
		return nil
	}
	for _, child := range children {
		if _, ok := child.NetworkSettings.Networks[network.DefaultNetwork]; !ok {
			return fmt.Errorf("cannot link to %s, as it does not belong to the default network", child.Name)
		}
	}

	var (
		childEndpoints []string
		cEndpointID    string
	)
	for linkAlias, child := range children {
		_, alias := path.Split(linkAlias)
		// allow access to the linked container via the alias, real name, and container hostname
		aliasList := alias + " " + child.Config.Hostname
		// only add the name if alias isn't equal to the name
		if alias != child.Name[1:] {
			aliasList = aliasList + " " + child.Name[1:]
		}
		defaultNW := child.NetworkSettings.Networks[network.DefaultNetwork]
		if defaultNW.IPAddress != "" {
			if err := sb.AddHostsEntry(ctx, aliasList, defaultNW.IPAddress); err != nil {
				return errors.Wrapf(err, "failed to add address to /etc/hosts for link to %s", child.Name)
			}
		}
		if defaultNW.GlobalIPv6Address != "" {
			if err := sb.AddHostsEntry(ctx, aliasList, defaultNW.GlobalIPv6Address); err != nil {
				return errors.Wrapf(err, "failed to add IPv6 address to /etc/hosts for link to %s", child.Name)
			}
		}
		cEndpointID = defaultNW.EndpointID
		if cEndpointID != "" {
			childEndpoints = append(childEndpoints, cEndpointID)
		}
	}

	var parentEndpoints []string
	for alias, parent := range parents {
		_, alias = path.Split(alias)
		// Update ctr's IP address in /etc/hosts files in containers with legacy-links to ctr.
		log.G(context.TODO()).Debugf("Update /etc/hosts of %s for alias %s with ip %s", parent.ID, alias, epConfig.IPAddress)
		if psb, _ := daemon.netController.GetSandbox(parent.ID); psb != nil {
			if err := psb.UpdateHostsEntry(alias, epConfig.IPAddress); err != nil {
				return errors.Wrapf(err, "failed to update /etc/hosts of %s for alias %s with IP %s",
					parent.ID, alias, epConfig.IPAddress)
			}
			if epConfig.GlobalIPv6Address != "" {
				if err := psb.UpdateHostsEntry(alias, epConfig.GlobalIPv6Address); err != nil {
					return errors.Wrapf(err, "failed to update /etc/hosts of %s for alias %s with IP %s",
						parent.ID, alias, epConfig.GlobalIPv6Address)
				}
			}
		}
		if cEndpointID != "" {
			parentEndpoints = append(parentEndpoints, cEndpointID)
		}
	}

	sb.UpdateLabels(bridge.LegacyContainerLinkOptions(parentEndpoints, childEndpoints))

	return nil
}

// createSecretsDir is used to create a dir suitable for storing container secrets.
// In practice this is using a tmpfs mount and is used for both "configs" and "secrets"
func (daemon *Daemon) createSecretsDir(ctr *container.Container) error {
	// retrieve possible remapped range start for root UID, GID
	uid, gid := daemon.idMapping.RootPair()
	dir, err := ctr.SecretMountPath()
	if err != nil {
		return errors.Wrap(err, "error getting container secrets dir")
	}

	// create tmpfs
	if err := user.MkdirAllAndChown(dir, 0o700, uid, gid); err != nil {
		return errors.Wrap(err, "error creating secret local mount path")
	}

	tmpfsOwnership := fmt.Sprintf("uid=%d,gid=%d", uid, gid)
	if err := mount.Mount("tmpfs", dir, "tmpfs", "nodev,nosuid,noexec,"+tmpfsOwnership); err != nil {
		return errors.Wrap(err, "unable to setup secret mount")
	}

	return nil
}

func (daemon *Daemon) remountSecretDir(ctr *container.Container) error {
	dir, err := ctr.SecretMountPath()
	if err != nil {
		return errors.Wrap(err, "error getting container secrets path")
	}
	if err := label.Relabel(dir, ctr.MountLabel, false); err != nil {
		log.G(context.TODO()).WithError(err).WithField("dir", dir).Warn("Error while attempting to set selinux label")
	}

	uid, gid := daemon.idMapping.RootPair()
	tmpfsOwnership := fmt.Sprintf("uid=%d,gid=%d", uid, gid)

	// remount secrets ro
	if err := mount.Mount("tmpfs", dir, "tmpfs", "remount,ro,"+tmpfsOwnership); err != nil {
		return errors.Wrap(err, "unable to remount dir as readonly")
	}

	return nil
}

func (daemon *Daemon) setupIPCDirs(ctr *container.Container) error {
	ipcMode := ctr.HostConfig.IpcMode

	switch {
	case ipcMode.IsContainer():
		ic, err := daemon.getIPCContainer(ipcMode.Container())
		if err != nil {
			return errors.Wrapf(err, "failed to join IPC namespace")
		}
		ctr.ShmPath = ic.ShmPath

	case ipcMode.IsHost():
		if _, err := os.Stat("/dev/shm"); err != nil {
			return errors.New("/dev/shm is not mounted, but must be for --ipc=host")
		}
		ctr.ShmPath = "/dev/shm"

	case ipcMode.IsPrivate(), ipcMode.IsNone():
		// c.ShmPath will/should not be used, so make it empty.
		// Container's /dev/shm mount comes from OCI spec.
		ctr.ShmPath = ""

	case ipcMode.IsEmpty():
		// A container was created by an older version of the daemon.
		// The default behavior used to be what is now called "shareable".
		fallthrough

	case ipcMode.IsShareable():
		uid, gid := daemon.idMapping.RootPair()
		if !ctr.HasMountFor("/dev/shm") {
			shmPath, err := ctr.ShmResourcePath()
			if err != nil {
				return err
			}

			if err := user.MkdirAllAndChown(shmPath, 0o700, uid, gid); err != nil {
				return err
			}

			shmproperty := "mode=1777,size=" + strconv.FormatInt(ctr.HostConfig.ShmSize, 10)
			if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, ctr.GetMountLabel())); err != nil {
				return fmt.Errorf("mounting shm tmpfs: %s", err)
			}

			if err := os.Chown(shmPath, uid, gid); err != nil {
				return err
			}
			ctr.ShmPath = shmPath
		}

	default:
		return fmt.Errorf("invalid IPC mode: %v", ipcMode)
	}

	return nil
}
