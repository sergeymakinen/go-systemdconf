// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// DeviceFile represents systemd.device — Device unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.device.html for details).
type DeviceFile struct {
	systemdconf.File

	Unit    UnitSection    // generic information about the unit that is not dependent on the type of unit
	Install InstallSection // installation information for the unit
}
