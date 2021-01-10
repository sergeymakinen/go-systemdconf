// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf/v2"

// AutomountFile represents systemd.automount — Automount unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.automount.html for details)
type AutomountFile struct {
	systemdconf.File

	Unit      UnitSection      // Generic information about the unit that is not dependent on the type of unit
	Automount AutomountSection // Information about the file system automount points it supervises
	Install   InstallSection   // Installation information for the unit
}

// AutomountSection represents information about the file system automount points it supervises
// (see https://www.freedesktop.org/software/systemd/man/systemd.automount.html#Options for details)
type AutomountSection struct {
	systemdconf.Section

	// Takes an absolute path of a directory of the automount point. If the automount point does not exist at time that the automount
	// point is installed, it is created. This string must be reflected in the unit filename. (See above.) This option is mandatory.
	Where systemdconf.Value

	// Directories of automount points (and any parent directories) are automatically created if needed. This option specifies
	// the file system access mode used when creating these directories. Takes an access mode in octal notation. Defaults to 0755.
	DirectoryMode systemdconf.Value

	// Configures an idle timeout. Once the mount has been idle for the specified time, systemd will attempt to unmount. Takes
	// a unit-less value in seconds, or a time span value such as "5min 20s". Pass 0 to disable the timeout logic. The timeout is disabled
	// by default.
	TimeoutIdleSec systemdconf.Value
}
