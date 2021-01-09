// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf"

// CoredumpFile represents coredump.conf, coredump.conf.d — Core dump storage configuration files
// (see https://www.freedesktop.org/software/systemd/man/coredump.conf.html for details)
type CoredumpFile struct {
	systemdconf.File

	Coredump CoredumpCoredumpSection // Behavior of systemd-coredump, a handler for core dumps invoked by the kernel
}

// CoredumpCoredumpSection represents behavior of systemd-coredump, a handler for core dumps invoked by the kernel
// (see https://www.freedesktop.org/software/systemd/man/coredump.conf.html#Options for details)
type CoredumpCoredumpSection struct {
	systemdconf.Section

	// Controls where to store cores. One of "none", "external", and "journal". When "none", the core dumps may be logged (including
	// the backtrace if possible), but not stored permanently. When "external" (the default), cores will be stored in /var/lib/systemd/coredump/.
	// When "journal", cores will be stored in the journal and rotated following normal journal rotation patterns.
	//
	// When cores are stored in the journal, they might be compressed following journal compression settings, see journald.conf.
	// When cores are stored externally, they will be compressed by default, see below.
	Storage systemdconf.Value

	// Controls compression for external storage. Takes a boolean argument, which defaults to "yes".
	Compress systemdconf.Value

	// The maximum size in bytes of a core which will be processed. Core dumps exceeding this size may be stored, but the backtrace
	// will not be generated.
	//
	// Setting Storage=none and ProcessSizeMax=0 disables all coredump handling except for a log entry.
	ProcessSizeMax systemdconf.Value

	// The maximum (uncompressed) size in bytes of a core to be saved.
	ExternalSizeMax systemdconf.Value

	// The maximum (uncompressed) size in bytes of a core to be saved.
	JournalSizeMax systemdconf.Value

	// Enforce limits on the disk space taken up by externally stored core dumps. MaxUse= makes sure that old core dumps are removed
	// as soon as the total disk space taken up by core dumps grows beyond this limit (defaults to 10% of the total disk size). KeepFree=
	// controls how much disk space to keep free at least (defaults to 15% of the total disk size). Note that the disk space used by
	// core dumps might temporarily exceed these limits while core dumps are processed. Note that old core dumps are also removed
	// based on time via systemd-tmpfiles. Set either value to 0 to turn off size-based clean-up.
	MaxUse systemdconf.Value

	// Enforce limits on the disk space taken up by externally stored core dumps. MaxUse= makes sure that old core dumps are removed
	// as soon as the total disk space taken up by core dumps grows beyond this limit (defaults to 10% of the total disk size). KeepFree=
	// controls how much disk space to keep free at least (defaults to 15% of the total disk size). Note that the disk space used by
	// core dumps might temporarily exceed these limits while core dumps are processed. Note that old core dumps are also removed
	// based on time via systemd-tmpfiles. Set either value to 0 to turn off size-based clean-up.
	KeepFree systemdconf.Value
}
