// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// DNSSDFile represents systemd.dnssd — DNS-SD configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.dnssd.html for details).
type DNSSDFile struct {
	systemdconf.File

	Service DNSSDServiceSection // discoverable network service announced in a local network with Multicast DNS broadcasts
}

// DNSSDServiceSection represents discoverable network service announced in a local network with Multicast DNS broadcasts
// (see https://www.freedesktop.org/software/systemd/man/systemd.dnssd.html#%5BService%5D%20Section%20Options for details).
type DNSSDServiceSection struct {
	systemdconf.Section

	// An instance name of the network service as defined in the section 4.1.1 of RFC 6763, e.g. "webserver".
	//
	// The option supports simple specifier expansion. The following expansions are understood:
	//
	//	+-----------+--------------------------------+--------------------------------+
	//	| SPECIFIER |            MEANING             |            DETAILS             |
	//	+-----------+--------------------------------+--------------------------------+
	//	| "%a"      | Architecture                   | A short string identifying     |
	//	|           |                                | the architecture of the        |
	//	|           |                                | local system. A string such    |
	//	|           |                                | as x86, x86-64 or arm64. See   |
	//	|           |                                | the architectures defined      |
	//	|           |                                | for ConditionArchitecture= in  |
	//	|           |                                | systemd.unit for a full list.  |
	//	| "%A"      | Operating system image version | The operating system image     |
	//	|           |                                | version identifier of the      |
	//	|           |                                | running system, as read from   |
	//	|           |                                | the IMAGE_VERSION= field of    |
	//	|           |                                | /etc/os-release. If not set,   |
	//	|           |                                | resolves to an empty string.   |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%b"      | Boot ID                        | The boot ID of the running     |
	//	|           |                                | system, formatted as           |
	//	|           |                                | string. See random for more    |
	//	|           |                                | information.                   |
	//	| "%B"      | Operating system build ID      | The operating system build     |
	//	|           |                                | identifier of the running      |
	//	|           |                                | system, as read from           |
	//	|           |                                | the BUILD_ID= field of         |
	//	|           |                                | /etc/os-release. If not set,   |
	//	|           |                                | resolves to an empty string.   |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%H"      | Host name                      | The hostname of the running    |
	//	|           |                                | system.                        |
	//	| "%m"      | Machine ID                     | The machine ID of the running  |
	//	|           |                                | system, formatted as string.   |
	//	|           |                                | See machine-id for more        |
	//	|           |                                | information.                   |
	//	| "%M"      | Operating system image         | The operating system image     |
	//	|           | identifier                     | identifier of the running      |
	//	|           |                                | system, as read from           |
	//	|           |                                | the IMAGE_ID= field of         |
	//	|           |                                | /etc/os-release. If not set,   |
	//	|           |                                | resolves to an empty string.   |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%o"      | Operating system ID            | The operating system           |
	//	|           |                                | identifier of the running      |
	//	|           |                                | system, as read from the ID=   |
	//	|           |                                | field of /etc/os-release.      |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%v"      | Kernel release                 | Identical to uname -r output.  |
	//	| "%w"      | Operating system version ID    | The operating system           |
	//	|           |                                | version identifier of the      |
	//	|           |                                | running system, as read from   |
	//	|           |                                | the VERSION_ID= field of       |
	//	|           |                                | /etc/os-release. If not set,   |
	//	|           |                                | resolves to an empty string.   |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%W"      | Operating system variant ID    | The operating system           |
	//	|           |                                | variant identifier of the      |
	//	|           |                                | running system, as read from   |
	//	|           |                                | the VARIANT_ID= field of       |
	//	|           |                                | /etc/os-release. If not set,   |
	//	|           |                                | resolves to an empty string.   |
	//	|           |                                | See os-release for more        |
	//	|           |                                | information.                   |
	//	| "%%"      | Single percent sign            | Use "%%" in place of "%" to    |
	//	|           |                                | specify a single percent sign. |
	//	+-----------+--------------------------------+--------------------------------+
	//
	// Added in version 236.
	Name systemdconf.Value

	// A type of the network service as defined in the section 4.1.2 of RFC 6763, e.g. "_http._tcp".
	//
	// Added in version 236.
	Type systemdconf.Value

	// A subtype of the network service as defined in the section 7.1 of RFC 6763, e.g. "_printer".
	//
	// Added in version 256.
	SubType systemdconf.Value

	// An IP port number of the network service.
	//
	// Added in version 236.
	Port systemdconf.Value

	// A priority number set in SRV resource records corresponding to the network service.
	//
	// Added in version 236.
	Priority systemdconf.Value

	// A weight number set in SRV resource records corresponding to the network service.
	//
	// Added in version 236.
	Weight systemdconf.Value

	// A whitespace-separated list of arbitrary key/value pairs conveying additional information about the named service
	// in the corresponding TXT resource record, e.g. "path=/portal/index.html". Keys and values can contain C-style escape
	// sequences which get translated upon reading configuration files.
	//
	// This option together with TxtData= may be specified more than once, in which case multiple TXT resource records will be
	// created for the service. If the empty string is assigned to this option, the list is reset and all prior assignments will
	// have no effect.
	//
	// Added in version 236.
	TxtText systemdconf.Value

	// A whitespace-separated list of arbitrary key/value pairs conveying additional information about the named service
	// in the corresponding TXT resource record where values are base64-encoded string representing any binary data, e.g. "data=YW55IGJpbmFyeSBkYXRhCg==".
	// Keys can contain C-style escape sequences which get translated upon reading configuration files.
	//
	// This option together with TxtText= may be specified more than once, in which case multiple TXT resource records will be
	// created for the service. If the empty string is assigned to this option, the list is reset and all prior assignments will
	// have no effect.
	//
	// Added in version 236.
	TxtData systemdconf.Value
}
