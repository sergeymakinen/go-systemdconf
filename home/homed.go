// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package home

import "github.com/sergeymakinen/go-systemdconf"

// HomedFile represents homed.conf, homed.conf.d — Home area/user account manager configuration files
// (see https://www.freedesktop.org/software/systemd/man/homed.conf.html for details)
type HomedFile struct {
	systemdconf.File

	Home HomedHomeSection // [Home] section
}

// HomedHomeSection represents [Home] section
// (see https://www.freedesktop.org/software/systemd/man/homed.conf.html#Options for details)
type HomedHomeSection struct {
	systemdconf.Section

	// The default storage to use for home areas. Takes one of "luks", "fscrypt", "directory", "subvolume", "cifs". For details
	// about these options, see homectl. If not configured or assigned the empty string, the default storage is automatically
	// determined: if not running in a container environment and /home/ is not itself encrypted, defaults to "luks". Otherwise
	// defaults to "subvolume" if /home/ is on a btrfs file system, and "directory" otherwise. Note that the storage selected
	// on the homectl command line always takes precedence.
	DefaultStorage systemdconf.Value

	// When using "luks" as storage (see above), selects the default file system to use inside the user's LUKS volume. Takes one
	// of "btrfs", "ext4" or "xfs". If not specified defaults to "btrfs". This setting has no effect if a different storage mechanism
	// is used. The file system type selected on the homectl command line always takes precedence.
	DefaultFileSystemType systemdconf.Value
}
