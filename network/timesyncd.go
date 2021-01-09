// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf"

// TimesyncdFile represents timesyncd.conf, timesyncd.conf.d — Network Time Synchronization configuration files
// (see https://www.freedesktop.org/software/systemd/man/timesyncd.conf.html for details)
type TimesyncdFile struct {
	systemdconf.File

	Resolve TimesyncdResolveSection // [Resolve] section
}

// TimesyncdResolveSection represents [Resolve] section
// (see https://www.freedesktop.org/software/systemd/man/timesyncd.conf.html#Options for details)
type TimesyncdResolveSection struct {
	systemdconf.Section

	// A space-separated list of NTP server host names or IP addresses. During runtime this list is combined with any per-interface
	// NTP servers acquired from systemd-networkd.service. systemd-timesyncd will contact all configured system or per-interface
	// servers in turn until one is found that responds. When the empty string is assigned, the list of NTP servers is reset, and
	// all assignments prior to this one will have no effect. This setting defaults to an empty list.
	NTP systemdconf.Value

	// A space-separated list of NTP server host names or IP addresses to be used as the fallback NTP servers. Any per-interface
	// NTP servers obtained from systemd-networkd.service take precedence over this setting, as do any servers set via NTP=
	// above. This setting is hence only used if no other NTP server information is known. When the empty string is assigned, the
	// list of NTP servers is reset, and all assignments prior to this one will have no effect. If this option is not given, a compiled-in
	// list of NTP servers is used instead.
	FallbackNTP systemdconf.Value

	// Maximum acceptable root distance. Takes a time value (in seconds). Defaults to 5 seconds.
	RootDistanceMaxSec systemdconf.Value

	// The minimum and maximum poll intervals for NTP messages. Each setting takes a time value (in seconds). PollIntervalMinSec=
	// must not be smaller than 16 seconds. PollIntervalMaxSec= must be larger than PollIntervalMinSec=. PollIntervalMinSec=
	// defaults to 32 seconds, and PollIntervalMaxSec= defaults to 2048 seconds.
	PollIntervalMinSec systemdconf.Value

	// The minimum and maximum poll intervals for NTP messages. Each setting takes a time value (in seconds). PollIntervalMinSec=
	// must not be smaller than 16 seconds. PollIntervalMaxSec= must be larger than PollIntervalMinSec=. PollIntervalMinSec=
	// defaults to 32 seconds, and PollIntervalMaxSec= defaults to 2048 seconds.
	PollIntervalMaxSec systemdconf.Value

	// Specifies the delaying attempts to contact servers after network is online. Takes a time value (in seconds). Defaults
	// to 30 seconds and must not be smaller than 1 seconds.
	ConnectionRetrySec systemdconf.Value
}
