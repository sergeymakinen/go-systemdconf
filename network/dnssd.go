// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf"

// DnssdFile represents systemd.dnssd — DNS-SD configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.dnssd.html for details)
type DnssdFile struct {
	systemdconf.File

	Service DnssdServiceSection // Discoverable network service announced in a local network with Multicast DNS broadcasts
}

// DnssdServiceSection represents discoverable network service announced in a local network with Multicast DNS broadcasts
// (see https://www.freedesktop.org/software/systemd/man/systemd.dnssd.html#%5BService%5D%20Section%20Options for details)
type DnssdServiceSection struct {
	systemdconf.Section

	// An instance name of the network service as defined in the section 4.1.1 of RFC 6763, e.g. "webserver".
	//
	// The option supports simple specifier expansion. The following expansions are understood:
	//
	// 	+-----------+-----------------------------+--------------------------------+
	// 	| SPECIFIER |           MEANING           |            DETAILS             |
	// 	+-----------+-----------------------------+--------------------------------+
	// 	| "%a"      | Architecture                | A short string identifying     |
	// 	|           |                             | the architecture of the        |
	// 	|           |                             | local system. A string such    |
	// 	|           |                             | as x86, x86-64 or arm64. See   |
	// 	|           |                             | the architectures defined      |
	// 	|           |                             | for ConditionArchitecture= in  |
	// 	|           |                             | systemd.unit for a full list.  |
	// 	| "%b"      | Boot ID                     | The boot ID of the running     |
	// 	|           |                             | system, formatted as           |
	// 	|           |                             | string. See random for more    |
	// 	|           |                             | information.                   |
	// 	| "%B"      | Operating system build ID   | The operating system build     |
	// 	|           |                             | identifier of the running      |
	// 	|           |                             | system, as read from           |
	// 	|           |                             | the BUILD_ID= field of         |
	// 	|           |                             | /etc/os-release. If not set,   |
	// 	|           |                             | resolves to an empty string.   |
	// 	|           |                             | See os-release for more        |
	// 	|           |                             | information.                   |
	// 	| "%H"      | Host name                   | The hostname of the running    |
	// 	|           |                             | system.                        |
	// 	| "%m"      | Machine ID                  | The machine ID of the running  |
	// 	|           |                             | system, formatted as string.   |
	// 	|           |                             | See machine-id for more        |
	// 	|           |                             | information.                   |
	// 	| "%o"      | Operating system ID         | The operating system           |
	// 	|           |                             | identifier of the running      |
	// 	|           |                             | system, as read from the ID=   |
	// 	|           |                             | field of /etc/os-release.      |
	// 	|           |                             | See os-release for more        |
	// 	|           |                             | information.                   |
	// 	| "%v"      | Kernel release              | Identical to uname -r output.  |
	// 	| "%w"      | Operating system version ID | The operating system           |
	// 	|           |                             | version identifier of the      |
	// 	|           |                             | running system, as read from   |
	// 	|           |                             | the VERSION_ID= field of       |
	// 	|           |                             | /etc/os-release. If not set,   |
	// 	|           |                             | resolves to an empty string.   |
	// 	|           |                             | See os-release for more        |
	// 	|           |                             | information.                   |
	// 	| "%W"      | Operating system variant ID | The operating system           |
	// 	|           |                             | variant identifier of the      |
	// 	|           |                             | running system, as read from   |
	// 	|           |                             | the VARIANT_ID= field of       |
	// 	|           |                             | /etc/os-release. If not set,   |
	// 	|           |                             | resolves to an empty string.   |
	// 	|           |                             | See os-release for more        |
	// 	|           |                             | information.                   |
	// 	| "%%"      | Single percent sign         | Use "%%" in place of "%" to    |
	// 	|           |                             | specify a single percent sign. |
	// 	+-----------+-----------------------------+--------------------------------+
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
