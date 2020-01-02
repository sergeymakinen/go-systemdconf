// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// DeviceFile represents information about a device unit as exposed in the sysfs/udev(7) device tree
type DeviceFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Install InstallSection // Installation information for the unit
}
