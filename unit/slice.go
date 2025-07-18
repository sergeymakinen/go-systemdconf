// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// SliceFile represents systemd.slice — Slice unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.slice.html for details).
type SliceFile struct {
	systemdconf.File

	Unit    UnitSection    // generic information about the unit that is not dependent on the type of unit
	Install InstallSection // installation information for the unit
	Slice   SliceSection   // [Slice] section
}

// SliceSection represents [Slice] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.slice.html#Options for details).
type SliceSection struct {
	systemdconf.Section
	ResourceControlOptions
}
