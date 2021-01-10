// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// PstoreFile represents pstore.conf, pstore.conf.d — PStore configuration file
// (see https://www.freedesktop.org/software/systemd/man/pstore.conf.html for details)
type PstoreFile struct {
	systemdconf.File

	PStore PstorePStoreSection // Behavior of systemd-pstore, a tool for archiving the contents of the persistent storage filesystem, pstore
}

// PstorePStoreSection represents behavior of systemd-pstore, a tool for archiving the contents of the persistent storage filesystem, pstore
// (see https://www.freedesktop.org/software/systemd/man/pstore.conf.html#Options for details)
type PstorePStoreSection struct {
	systemdconf.Section

	// Controls where to archive (i.e. copy) files from the pstore filesystem. One of "none", "external", and "journal". When
	// "none", the tool exits without processing files in the pstore filesystem. When "external" (the default), files are archived
	// into /var/lib/systemd/pstore/, and logged into the journal. When "journal", pstore file contents are logged only in
	// the journal.
	Storage systemdconf.Value

	// Controls whether or not files are removed from pstore after processing. Takes a boolean value. When true, a pstore file
	// is removed from the pstore once it has been archived (either to disk or into the journal). When false, processing of pstore
	// files occurs normally, but the files remain in the pstore. The default is true in order to maintain the pstore in a nearly
	// empty state, so that the pstore has storage available for the next kernel error event.
	Unlink systemdconf.Value
}
