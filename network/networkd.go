// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf/v2"

// NetworkdFile represents networkd.conf, networkd.conf.d — Global Network configuration files
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html for details)
type NetworkdFile struct {
	systemdconf.File

	Network NetworkdNetworkSection // [Network] section
	DHCP    NetworkdDHCPSection    // DHCP Unique Identifier (DUID) value used by DHCP protocol
}

// NetworkdNetworkSection represents [Network] section
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BNetwork%5D%20Section%20Options for details)
type NetworkdNetworkSection struct {
	systemdconf.Section

	// Takes a boolean. If set to yes, then systemd-networkd measures the traffic of each interface, and networkctl status INTERFACE
	// shows the measured speed. Defaults to no.
	SpeedMeter systemdconf.Value

	// Specifies the time interval to calculate the traffic speed of each interface. If SpeedMeter=no, the value is ignored.
	// Defaults to 10sec.
	SpeedMeterIntervalSec systemdconf.Value

	// A boolean. When true, systemd-networkd will store any routes configured by other tools in its memory. When false, systemd-networkd
	// will not manage the foreign routes, thus they are kept even if KeepConfiguration= is false. Defaults to yes.
	ManageForeignRoutes systemdconf.Value
}

// NetworkdDHCPSection represents DHCP Unique Identifier (DUID) value used by DHCP protocol
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BDHCP%5D%20Section%20Options for details)
type NetworkdDHCPSection struct {
	systemdconf.Section

	// Specifies how the DUID should be generated. See RFC 3315 for a description of all the options.
	//
	// The following values are understood:
	//
	// 	vendor: If "DUIDType=vendor", then the DUID value will be generated using "43793" as the vendor identifier (systemd)
	// 	and hashed contents of machine-id. This is the default if DUIDType= is not specified.
	// 	uuid: If "DUIDType=uuid", and DUIDRawData= is not set, then the product UUID is used as a DUID value. If a system does not
	// 	have valid product UUID, then an application-specific machine-id is used as a DUID value. About the application-specific
	// 	machine ID, see sd_id128_get_machine_app_specific.
	// 	link-layer-time[:TIME], link-layer: If "link-layer-time" or "link-layer" is specified, then the MAC address of the
	// 	interface is used as a DUID value. The value "link-layer-time" can take additional time value after a colon, e.g. "link-layer-time:2018-01-23
	// 	12:34:56 UTC". The default time value is "2000-01-01 00:00:00 UTC".
	//
	// In all cases, DUIDRawData= can be used to override the actual DUID value that is used.
	DUIDType systemdconf.Value

	// Specifies the DHCP DUID value as a single newline-terminated, hexadecimal string, with each byte separated by ":". The
	// DUID that is sent is composed of the DUID type specified by DUIDType= and the value configured here.
	//
	// The DUID value specified here overrides the DUID that systemd-networkd.service generates from the machine ID. To configure
	// DUID per-network, see systemd.network. The configured DHCP DUID should conform to the specification in RFC 3315, RFC
	// 6355. To configure IAID, see systemd.network.
	//
	// 	DUIDType=vendor DUIDRawData=00:00:ab:11:f9:2a:c2:77:29:f9:5c:00
	//
	// This specifies a 14 byte DUID, with the type DUID-EN ("00:02"), enterprise number 43793 ("00:00:ab:11"), and identifier
	// value "f9:2a:c2:77:29:f9:5c:00".
	DUIDRawData systemdconf.Value
}
