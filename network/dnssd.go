// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf"

// DNSSDServiceSection represents discoverable network service announced in a local network with Multicast DNS broadcasts
type DNSSDServiceSection struct {
	systemdconf.Section

	// An instance name of the network service as defined in the section 4.1.1 of RFC 6763, e.g. "webserver".
	//
	// The option supports simple specifier expansion. The following expansions are understood:
	//
	// 	+-----------+----------------+--------------------------------+
	// 	| SPECIFIER |    MEANING     |            DETAILS             |
	// 	+-----------+----------------+--------------------------------+
	// 	| "%m"      | Machine ID     | The machine ID of the running  |
	// 	|           |                | system, formatted as string.   |
	// 	|           |                | See machine-id for more        |
	// 	|           |                | information.                   |
	// 	| "%b"      | Boot ID        | The boot ID of the running     |
	// 	|           |                | system, formatted as           |
	// 	|           |                | string. See random for more    |
	// 	|           |                | information.                   |
	// 	| "%H"      | Host name      | The hostname of the running    |
	// 	|           |                | system.                        |
	// 	| "%v"      | Kernel release | Identical to uname -r output.  |
	// 	+-----------+----------------+--------------------------------+
	Name systemdconf.Value

	// A type of the network service as defined in the section 4.1.2 of RFC 6763, e.g. "_http._tcp".
	Type systemdconf.Value

	// An IP port number of the network service.
	Port systemdconf.Value

	// A priority number set in SRV resource records corresponding to the network service.
	Priority systemdconf.Value

	// A weight number set in SRV resource records corresponding to the network service.
	Weight systemdconf.Value

	// A whitespace-separated list of arbitrary key/value pairs conveying additional information about the named service
	// in the corresponding TXT resource record, e.g. "path=/portal/index.html". Keys and values can contain C-style escape
	// sequences which get translated upon reading configuration files.
	//
	// This option together with TxtData= may be specified more than once, in which case multiple TXT resource records will be
	// created for the service. If the empty string is assigned to this option, the list is reset and all prior assignments will
	// have no effect.
	TxtText systemdconf.Value

	// A whitespace-separated list of arbitrary key/value pairs conveying additional information about the named service
	// in the corresponding TXT resource record where values are base64-encoded string representing any binary data, e.g. "data=YW55IGJpbmFyeSBkYXRhCg==".
	// Keys can contain C-style escape sequences which get translated upon reading configuration files.
	//
	// This option together with TxtText= may be specified more than once, in which case multiple TXT resource records will be
	// created for the service. If the empty string is assigned to this option, the list is reset and all prior assignments will
	// have no effect.
	TxtData systemdconf.Value
}

// DNSSDFile represents DNS-SD setup configuration is performed by systemd-resolved
type DNSSDFile struct {
	systemdconf.File

	Service DNSSDServiceSection // Discoverable network service announced in a local network with Multicast DNS broadcasts
}
