// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// MountFile represents systemd.mount — Mount unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.mount.html for details).
type MountFile struct {
	systemdconf.File

	Unit    UnitSection    // generic information about the unit that is not dependent on the type of unit
	Mount   MountSection   // information about the file system mount points it supervises
	Install InstallSection // installation information for the unit
}

// MountSection represents information about the file system mount points it supervises
// (see https://www.freedesktop.org/software/systemd/man/systemd.mount.html#Options for details).
type MountSection struct {
	systemdconf.Section
	ExecOptions
	KillOptions
	ResourceControlOptions

	// Takes an absolute path or a fstab-style identifier of a device node, file or other resource to mount. See mount for details.
	// If this refers to a device node, a dependency on the respective device unit is automatically created. (See systemd.device
	// for more information.) This option is mandatory. Note that the usual specifier expansion is applied to this setting, literal
	// percent characters should hence be written as "%%". If this mount is a bind mount and the specified path does not exist yet
	// it is created as directory.
	What systemdconf.Value

	// Takes an absolute path of a file or directory for the mount point; in particular, the destination cannot be a symbolic link.
	// If the mount point does not exist at the time of mounting, it is created as either a directory or a file. The former is the usual
	// case; the latter is done only if this mount is a bind mount and the source (What=) is not a directory. This string must be reflected
	// in the unit filename. (See above.) This option is mandatory.
	Where systemdconf.Value

	// Takes a string for the file system type. See mount for details. This setting is optional.
	//
	// If the type is "overlay", and "upperdir=" or "workdir=" are specified as options and the directories do not exist, they
	// will be created.
	Type systemdconf.Value

	// Mount options to use when mounting. This takes a comma-separated list of options. This setting is optional. Note that the
	// usual specifier expansion is applied to this setting, literal percent characters should hence be written as "%%".
	Options systemdconf.Value

	// Takes a boolean argument. If true, parsing of the options specified in Options= is relaxed, and unknown mount options are
	// tolerated. This corresponds with mount's -s switch. Defaults to off.
	//
	// Added in version 215.
	SloppyOptions systemdconf.Value

	// Takes a boolean argument. If true, detach the filesystem from the filesystem hierarchy at time of the unmount operation,
	// and clean up all references to the filesystem as soon as they are not busy anymore. This corresponds with umount's -l switch.
	// Defaults to off.
	//
	// Added in version 232.
	LazyUnmount systemdconf.Value

	// Takes a boolean argument. If false, a mount point that shall be mounted read-write but cannot be mounted so is retried to
	// be mounted read-only. If true the operation will fail immediately after the read-write mount attempt did not succeed.
	// This corresponds with mount's -w switch. Defaults to off.
	//
	// Added in version 246.
	ReadWriteOnly systemdconf.Value

	// Takes a boolean argument. If true, force an unmount (in case of an unreachable NFS system). This corresponds with umount's
	// -f switch. Defaults to off.
	//
	// Added in version 232.
	ForceUnmount systemdconf.Value

	// Directories of mount points (and any parent directories) are automatically created if needed. This option specifies
	// the file system access mode used when creating these directories. Takes an access mode in octal notation. Defaults to 0755.
	DirectoryMode systemdconf.Value

	// Configures the time to wait for the mount command to finish. If a command does not exit within the configured time, the mount
	// will be considered failed and be shut down again. All commands still running will be terminated forcibly via SIGTERM, and
	// after another delay of this time with SIGKILL. (See KillMode= in systemd.kill.) Takes a unit-less value in seconds, or
	// a time span value such as "5min 20s". Pass 0 to disable the timeout logic. The default value is set from DefaultTimeoutStartSec=
	// option in systemd-system.conf.
	TimeoutSec systemdconf.Value
}
