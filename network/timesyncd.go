// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// TimesyncdFile represents timesyncd.conf, timesyncd.conf.d — Network Time Synchronization configuration files
// (see https://www.freedesktop.org/software/systemd/man/timesyncd.conf.html for details).
type TimesyncdFile struct {
	systemdconf.File

	Time TimesyncdTimeSection // [Time] section
}

// TimesyncdTimeSection represents [Time] section
// (see https://www.freedesktop.org/software/systemd/man/timesyncd.conf.html#Options for details).
type TimesyncdTimeSection struct {
	systemdconf.Section

	// A space-separated list of NTP server host names or IP addresses. During runtime this list is combined with any per-interface
	// NTP servers acquired from systemd-networkd.service. systemd-timesyncd will contact all configured system or per-interface
	// servers in turn, until one responds. When the empty string is assigned, the list of NTP servers is reset, and all prior assignments
	// will have no effect. This setting defaults to an empty list.
	//
	// Added in version 216.
	NTP systemdconf.Value

	// A space-separated list of NTP server host names or IP addresses to be used as the fallback NTP servers. Any per-interface
	// NTP servers obtained from systemd-networkd.service take precedence over this setting, as do any servers set via NTP=
	// above. This setting is hence only relevant if no other NTP server information is known. When the empty string is assigned,
	// the list of NTP servers is reset, and all prior assignments will have no effect. If this option is not given, a compiled-in
	// list of NTP servers is used.
	//
	// Added in version 216.
	FallbackNTP systemdconf.Value

	// Maximum acceptable root distance, i.e. the maximum estimated time required for a packet to travel to the server we are connected
	// to from the server with the reference clock. If the current server does not satisfy this limit, systemd-timesyncd will
	// switch to a different server.
	//
	// Takes a time span value. The default unit is seconds, but other units may be specified, see systemd.time. Defaults to 5 seconds.
	//
	// Added in version 236.
	RootDistanceMaxSec systemdconf.Value

	// The minimum and maximum poll intervals for NTP messages. Polling starts at the minimum poll interval, and is adjusted within
	// the specified limits in response to received packets.
	//
	// Each setting takes a time span value. The default unit is seconds, but other units may be specified, see systemd.time. PollIntervalMinSec=
	// defaults to 32 seconds and must not be smaller than 16 seconds. PollIntervalMaxSec= defaults to 34 min 8 s (2048 seconds)
	// and must be larger than PollIntervalMinSec=.
	//
	// Added in version 236.
	PollIntervalMinSec systemdconf.Value

	// The minimum and maximum poll intervals for NTP messages. Polling starts at the minimum poll interval, and is adjusted within
	// the specified limits in response to received packets.
	//
	// Each setting takes a time span value. The default unit is seconds, but other units may be specified, see systemd.time. PollIntervalMinSec=
	// defaults to 32 seconds and must not be smaller than 16 seconds. PollIntervalMaxSec= defaults to 34 min 8 s (2048 seconds)
	// and must be larger than PollIntervalMinSec=.
	//
	// Added in version 236.
	PollIntervalMaxSec systemdconf.Value

	// Specifies the minimum delay before subsequent attempts to contact a new NTP server are made.
	//
	// Takes a time span value. The default unit is seconds, but other units may be specified, see systemd.time. Defaults to 30
	// seconds and must not be smaller than 1 second.
	//
	// Added in version 248.
	ConnectionRetrySec systemdconf.Value

	// The interval at which the current time is periodically saved to disk, in the absence of any recent synchronisation from
	// an NTP server. This is especially useful for offline systems with no local RTC, as it will guarantee that the system clock
	// remains roughly monotonic across reboots.
	//
	// Takes a time interval value. The default unit is seconds, but other units may be specified, see systemd.time. Defaults
	// to 60 seconds.
	//
	// Added in version 250.
	SaveIntervalSec systemdconf.Value
}
