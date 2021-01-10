// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// JournalRemoteFile represents journal-remote.conf, journal-remote.conf.d — Configuration files for the service accepting remote journal uploads
// (see https://www.freedesktop.org/software/systemd/man/journal-remote.conf.html for details)
type JournalRemoteFile struct {
	systemdconf.File

	Remote JournalRemoteRemoteSection // Parameters of systemd-journal-remote.service
}

// JournalRemoteRemoteSection represents parameters of systemd-journal-remote.service
// (see https://www.freedesktop.org/software/systemd/man/journal-remote.conf.html#Options for details)
type JournalRemoteRemoteSection struct {
	systemdconf.Section

	// Periodically sign the data in the journal using Forward Secure Sealing.
	Seal systemdconf.Value

	// One of "host" or "none".
	SplitMode systemdconf.Value

	// SSL key in PEM format.
	ServerKeyFile systemdconf.Value

	// SSL certificate in PEM format.
	ServerCertificateFile systemdconf.Value

	// SSL CA certificate.
	TrustedCertificateFile systemdconf.Value
}
