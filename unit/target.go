// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// TargetFile represents systemd.target — Target unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.target.html for details).
type TargetFile struct {
	systemdconf.File

	Unit    UnitSection    // generic information about the unit that is not dependent on the type of unit
	Install InstallSection // installation information for the unit
}
