// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// JournalUploadFile represents journal-upload.conf, journal-upload.conf.d — Configuration files for the journal upload service
// (see https://www.freedesktop.org/software/systemd/man/journal-upload.conf.html for details).
type JournalUploadFile struct {
	systemdconf.File

	Upload JournalUploadUploadSection // parameters of systemd-journal-upload.service
}

// JournalUploadUploadSection represents parameters of systemd-journal-upload.service
// (see https://www.freedesktop.org/software/systemd/man/journal-upload.conf.html#Options for details).
type JournalUploadUploadSection struct {
	systemdconf.Section

	// The URL to upload the journal entries to. See the description of --url= option in systemd-journal-upload for the description
	// of possible values. There is no default value, so either this option or the command-line option must be always present to
	// make an upload.
	//
	// Added in version 232.
	URL systemdconf.Value

	// SSL key in PEM format.
	//
	// Added in version 232.
	ServerKeyFile systemdconf.Value

	// SSL CA certificate in PEM format.
	//
	// Added in version 232.
	ServerCertificateFile systemdconf.Value

	// SSL CA certificate.
	//
	// Added in version 232.
	TrustedCertificateFile systemdconf.Value

	// When network connectivity to the server is lost, this option configures the time to wait for the connectivity to get restored.
	// If the server is not reachable over the network for the configured time, systemd-journal-upload exits. Takes a value in
	// seconds (or in other time units if suffixed with "ms", "min", "h", etc). For details, see systemd.time.
	//
	// Added in version 249.
	NetworkTimeoutSec systemdconf.Value
}
