// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// JournalUploadFile represents journal-upload.conf, journal-upload.conf.d — Configuration files for the journal upload service
// (see https://www.freedesktop.org/software/systemd/man/journal-upload.conf.html for details)
type JournalUploadFile struct {
	systemdconf.File

	Upload JournalUploadUploadSection // Parameters of systemd-journal-upload.service
}

// JournalUploadUploadSection represents parameters of systemd-journal-upload.service
// (see https://www.freedesktop.org/software/systemd/man/journal-upload.conf.html#Options for details)
type JournalUploadUploadSection struct {
	systemdconf.Section

	// The URL to upload the journal entries to. See the description of --url= option in systemd-journal-upload for the description
	// of possible values. There is no default value, so either this option or the command-line option must be always present to
	// make an upload.
	URL systemdconf.Value

	// SSL key in PEM format.
	ServerKeyFile systemdconf.Value

	// SSL CA certificate in PEM format.
	ServerCertificateFile systemdconf.Value

	// SSL CA certificate.
	TrustedCertificateFile systemdconf.Value
}
