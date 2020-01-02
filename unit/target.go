// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// TargetFile represents information about a target unit of systemd, which is used for grouping units and as well-known synchronization points during start-up
type TargetFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Install InstallSection // Installation information for the unit
}
