// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// NetworkdFile represents networkd.conf, networkd.conf.d — Global Network configuration files
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html for details).
type NetworkdFile struct {
	systemdconf.File

	Network          NetworkdNetworkSection            // [Network] section
	IPv6AcceptRA     NetworkdIPv6AcceptRASection       // the default setting of the Neighbor Discovery
	IPv6AddressLabel []NetworkdIPv6AddressLabelSection // [IPv6AddressLabel] section
	DHCPv4           NetworkdDHCPv4Section             // the DHCP Unique Identifier (DUID) value used by DHCP protocol
	DHCPv6           NetworkdDHCPv6Section             // the DHCP Unique Identifier (DUID) value used by DHCPv6 protocol
	DHCPServer       NetworkdDHCPServerSection         // the default setting of the DHCP server
}

// NetworkdNetworkSection represents [Network] section
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BNetwork%5D%20Section%20Options for details).
type NetworkdNetworkSection struct {
	systemdconf.Section

	// Takes a boolean. If set to yes, then systemd-networkd measures the traffic of each interface, and networkctl status INTERFACE
	// shows the measured speed. Defaults to no.
	//
	// Added in version 244.
	SpeedMeter systemdconf.Value

	// Specifies the time interval to calculate the traffic speed of each interface. If SpeedMeter=no, the value is ignored.
	// Defaults to 10sec.
	//
	// Added in version 244.
	SpeedMeterIntervalSec systemdconf.Value

	// A boolean. When true, systemd-networkd will remove rules that are not configured in .network files (except for rules with
	// protocol "kernel"). When false, it will not remove any foreign rules, keeping them even if they are not configured in a .network
	// file. Defaults to yes.
	//
	// Added in version 249.
	ManageForeignRoutingPolicyRules systemdconf.Value

	// A boolean. When true, systemd-networkd will remove routes that are not configured in .network files (except for routes
	// with protocol "kernel", "dhcp" when KeepConfiguration= is true or "dhcp", and "static" when KeepConfiguration= is true
	// or "static"). When false, it will not remove any foreign routes, keeping them even if they are not configured in a .network
	// file. Defaults to yes.
	//
	// Added in version 246.
	ManageForeignRoutes systemdconf.Value

	// A boolean. When true, systemd-networkd will remove nexthops that are not configured in .network files (except for routes
	// with protocol "kernel"). When false, it will not remove any foreign nexthops, keeping them even if they are not configured
	// in a .network file. Defaults to yes.
	//
	// Added in version 256.
	ManageForeignNextHops systemdconf.Value

	// Defines the route table name. Takes a whitespace-separated list of the pairs of route table name and number. The route table
	// name and number in each pair are separated with a colon, i.e., "name:number". The route table name must not be "default",
	// "main", or "local", as these route table names are predefined with route table number 253, 254, and 255, respectively.
	// The route table number must be an integer in the range 1…4294967295, except for predefined numbers 253, 254, and 255. This
	// setting can be specified multiple times. If an empty string is specified, then the list specified earlier are cleared.
	// Defaults to unset.
	//
	// Added in version 248.
	RouteTable systemdconf.Value

	// Configures IPv4 packet forwarding for the system. Takes a boolean value. This controls the net.ipv4.conf.default.forwarding
	// and net.ipv4.conf.all.forwarding sysctl options. See IP Sysctl for more details about the sysctl options. Defaults
	// to unset and the sysctl options will not be changed.
	//
	// If an interface is configured with a .network file that enables IPMasquerade= for IPv4 (that is, "ipv4" or "both"), this
	// setting is implied unless explicitly specified. See IPMasquerade= in systemd.network for more details.
	//
	// Added in version 256.
	IPv4Forwarding systemdconf.Value

	// Configures IPv6 packet forwarding for the system. Takes a boolean value. This controls the net.ipv6.conf.default.forwarding
	// and net.ipv6.conf.all.forwarding sysctl options. See IP Sysctl for more details about the sysctl options. Defaults
	// to unset and the sysctl options will not be changed.
	//
	// If an interface is configured with a .network file that enables IPMasquerade= for IPv6 (that is, "ipv6" or "both"), this
	// setting is implied unless explicitly specified. See IPMasquerade= in systemd.network for more details.
	//
	// Added in version 256.
	IPv6Forwarding systemdconf.Value

	// Specifies the default value for per-network IPv6PrivacyExtensions=. Takes a boolean or the special values "prefer-public"
	// and "kernel". See for details in systemd.network. Defaults to "no".
	//
	// Added in version 254.
	IPv6PrivacyExtensions systemdconf.Value

	// Specifies the network- and protocol-independent default value for the same settings in [IPv6AcceptRA], [DHCPv4], and
	// [DHCPv6] sections below. Takes a boolean, or the special value route. See the same setting in systemd.network. Defaults
	// to "no".
	//
	// Added in version 256.
	UseDomains systemdconf.Value
}

// NetworkdIPv6AcceptRASection represents the default setting of the Neighbor Discovery
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BIPv6AcceptRA%5D%20Section%20Options for details).
type NetworkdIPv6AcceptRASection struct {
	systemdconf.Section

	// Specifies the network-independent default value for the same setting in the [IPv6AcceptRA] section in systemd.network.
	// Takes a boolean, or the special value route. When unspecified, the value specified in the [Network] section in networkd.conf,
	// which defaults to "no", will be used.
	//
	// Added in version 256.
	UseDomains systemdconf.Value
}

// NetworkdIPv6AddressLabelSection represents [IPv6AddressLabel] section
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BIPv6AddressLabel%5D%20Section%20Options for details).
type NetworkdIPv6AddressLabelSection struct {
	systemdconf.Section

	// The label for the prefix, an unsigned integer in the range 0…4294967294. 0xffffffff is reserved. This setting is mandatory.
	//
	// Added in version 257.
	Label systemdconf.Value

	// IPv6 prefix is an address with a prefix length, separated by a slash "/" character. This setting is mandatory.
	//
	// Added in version 257.
	Prefix systemdconf.Value
}

// NetworkdDHCPv4Section represents the DHCP Unique Identifier (DUID) value used by DHCP protocol
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BDHCPv4%5D%20Section%20Options for details).
type NetworkdDHCPv4Section struct {
	systemdconf.Section

	// Specifies how the DUID should be generated. See RFC 3315 for a description of all the options.
	//
	// This takes an integer in the range 0…65535, or one of the following string values:
	//
	//	vendor: If "DUIDType=vendor", then the DUID value will be generated using "43793" as the vendor identifier (systemd)
	//	and hashed contents of machine-id. This is the default if DUIDType= is not specified.
	//
	//	Added in version 230.
	//	uuid: If "DUIDType=uuid", and DUIDRawData= is not set, then the product UUID is used as a DUID value. If a system does not
	//	have valid product UUID, then an application-specific machine-id is used as a DUID value. About the application-specific
	//	machine ID, see sd_id128_get_machine_app_specific.
	//
	//	Added in version 230.
	//	link-layer-time[:TIME], link-layer: If "link-layer-time" or "link-layer" is specified, then the MAC address of the
	//	interface is used as a DUID value. The value "link-layer-time" can take additional time value after a colon, e.g. "link-layer-time:2018-01-23
	//	12:34:56 UTC". The default time value is "2000-01-01 00:00:00 UTC".
	//
	//	Added in version 240.
	//
	// In all cases, DUIDRawData= can be used to override the actual DUID value that is used.
	//
	// Added in version 230.
	DUIDType systemdconf.Value

	// Specifies the DHCP DUID value as a single newline-terminated, hexadecimal string, with each byte separated by ":". The
	// DUID that is sent is composed of the DUID type specified by DUIDType= and the value configured here.
	//
	// The DUID value specified here overrides the DUID that systemd-networkd.service generates from the machine ID. To configure
	// DUID per-network, see systemd.network. The configured DHCP DUID should conform to the specification in RFC 3315, RFC
	// 6355. To configure IAID, see systemd.network.
	//
	//	DUIDType=vendor DUIDRawData=00:00:ab:11:f9:2a:c2:77:29:f9:5c:00
	//
	// This specifies a 14 byte DUID, with the type DUID-EN ("00:02"), enterprise number 43793 ("00:00:ab:11"), and identifier
	// value "f9:2a:c2:77:29:f9:5c:00".
	//
	// Added in version 230.
	DUIDRawData systemdconf.Value

	// Same as the one in the [IPv6AcceptRA] section, but applied for DHCPv4 protocol.
	//
	// Added in version 256.
	UseDomains systemdconf.Value
}

// NetworkdDHCPv6Section represents the DHCP Unique Identifier (DUID) value used by DHCPv6 protocol
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BDHCPv6%5D%20Section%20Options for details).
type NetworkdDHCPv6Section struct {
	systemdconf.Section

	// As in the [DHCPv4] section.
	//
	// Added in version 249.
	DUIDType systemdconf.Value

	// As in the [DHCPv4] section.
	//
	// Added in version 249.
	DUIDRawData systemdconf.Value

	// As in the [DHCPv4] section.
	//
	// Added in version 256.
	UseDomains systemdconf.Value
}

// NetworkdDHCPServerSection represents the default setting of the DHCP server
// (see https://www.freedesktop.org/software/systemd/man/networkd.conf.html#%5BDHCPServer%5D%20Section%20Options for details).
type NetworkdDHCPServerSection struct {
	systemdconf.Section

	// Specifies the default value for per-network PersistLeases=. Takes a boolean. See for details in systemd.network. Defaults
	// to "yes".
	//
	// Added in version 256.
	PersistLeases systemdconf.Value
}
