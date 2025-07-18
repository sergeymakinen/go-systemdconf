// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// JournalRemoteFile represents journal-remote.conf, journal-remote.conf.d — Configuration files for the service accepting remote journal uploads
// (see https://www.freedesktop.org/software/systemd/man/journal-remote.conf.html for details).
type JournalRemoteFile struct {
	systemdconf.File

	Remote JournalRemoteRemoteSection // parameters of systemd-journal-remote.service
}

// JournalRemoteRemoteSection represents parameters of systemd-journal-remote.service
// (see https://www.freedesktop.org/software/systemd/man/journal-remote.conf.html#Options for details).
type JournalRemoteRemoteSection struct {
	systemdconf.Section

	// Periodically sign the data in the journal using Forward Secure Sealing.
	//
	// Added in version 229.
	Seal systemdconf.Value

	// One of "host" or "none".
	//
	// Added in version 220.
	SplitMode systemdconf.Value

	// SSL key in PEM format.
	//
	// Added in version 220.
	ServerKeyFile systemdconf.Value

	// SSL certificate in PEM format.
	//
	// Added in version 220.
	ServerCertificateFile systemdconf.Value

	// SSL CA certificate.
	//
	// Added in version 220.
	TrustedCertificateFile systemdconf.Value

	// These are analogous to SystemMaxUse=, SystemKeepFree=, SystemMaxFileSize= and SystemMaxFiles= in journald.conf.
	//
	// MaxUse= controls how much disk space the systemd-journal-remote may use up at most. KeepFree= controls how much disk space
	// systemd-journal-remote shall leave free for other uses. systemd-journal-remote will respect both limits and use the
	// smaller of the two values.
	//
	// MaxFiles= controls how many individual journal files to keep at most. Note that only archived files are deleted to reduce
	// the number of files until this limit is reached; active files will stay around. This means that, in effect, there might still
	// be more journal files around in total than this limit after a vacuuming operation is complete.
	//
	// Added in version 253.
	MaxUse systemdconf.Value

	// These are analogous to SystemMaxUse=, SystemKeepFree=, SystemMaxFileSize= and SystemMaxFiles= in journald.conf.
	//
	// MaxUse= controls how much disk space the systemd-journal-remote may use up at most. KeepFree= controls how much disk space
	// systemd-journal-remote shall leave free for other uses. systemd-journal-remote will respect both limits and use the
	// smaller of the two values.
	//
	// MaxFiles= controls how many individual journal files to keep at most. Note that only archived files are deleted to reduce
	// the number of files until this limit is reached; active files will stay around. This means that, in effect, there might still
	// be more journal files around in total than this limit after a vacuuming operation is complete.
	//
	// Added in version 253.
	KeepFree systemdconf.Value

	// These are analogous to SystemMaxUse=, SystemKeepFree=, SystemMaxFileSize= and SystemMaxFiles= in journald.conf.
	//
	// MaxUse= controls how much disk space the systemd-journal-remote may use up at most. KeepFree= controls how much disk space
	// systemd-journal-remote shall leave free for other uses. systemd-journal-remote will respect both limits and use the
	// smaller of the two values.
	//
	// MaxFiles= controls how many individual journal files to keep at most. Note that only archived files are deleted to reduce
	// the number of files until this limit is reached; active files will stay around. This means that, in effect, there might still
	// be more journal files around in total than this limit after a vacuuming operation is complete.
	//
	// Added in version 253.
	MaxFileSize systemdconf.Value

	// These are analogous to SystemMaxUse=, SystemKeepFree=, SystemMaxFileSize= and SystemMaxFiles= in journald.conf.
	//
	// MaxUse= controls how much disk space the systemd-journal-remote may use up at most. KeepFree= controls how much disk space
	// systemd-journal-remote shall leave free for other uses. systemd-journal-remote will respect both limits and use the
	// smaller of the two values.
	//
	// MaxFiles= controls how many individual journal files to keep at most. Note that only archived files are deleted to reduce
	// the number of files until this limit is reached; active files will stay around. This means that, in effect, there might still
	// be more journal files around in total than this limit after a vacuuming operation is complete.
	//
	// Added in version 253.
	MaxFiles systemdconf.Value
}
