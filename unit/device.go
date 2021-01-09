// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// DeviceFile represents systemd.device — Device unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.device.html for details)
type DeviceFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Install InstallSection // Installation information for the unit
}
