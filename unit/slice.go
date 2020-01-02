// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// SliceFile represents information about a slice unit. A slice unit is a concept for hierarchically managing resources of a group of processes
type SliceFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Install InstallSection // Installation information for the unit
}
