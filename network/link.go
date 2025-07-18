// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// LinkFile represents systemd.link — Network device configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html for details).
type LinkFile struct {
	systemdconf.File

	Match LinkMatchSection   // [Match] section, which determines if a given link file may be applied to a given device
	Link  LinkLinkSection    // [Link] section, which specifies how the device should be configured
	SRIOV []LinkSRIOVSection `systemd:"SR-IOV"` // [SR-IOV] section, which provides the ability to partition a single physical PCI resource into virtual PCI functions which can then be e.g. injected into a VM
}

// LinkMatchSection represents [Match] section, which determines if a given link file may be applied to a given device
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html#%5BMatch%5D%20Section%20Options for details).
type LinkMatchSection struct {
	systemdconf.Section

	// A whitespace-separated list of hardware addresses. The acceptable formats are:
	//
	//	colon-delimited hexadecimal: Each field must be one byte. E.g. "12:34:56:78:90:ab" or "AA:BB:CC:DD:EE:FF".
	//
	//	Added in version 250.
	//	hyphen-delimited hexadecimal: Each field must be one byte. E.g. "12-34-56-78-90-ab" or "AA-BB-CC-DD-EE-FF".
	//
	//	Added in version 250.
	//	dot-delimited hexadecimal: Each field must be two bytes. E.g. "1234.5678.90ab" or "AABB.CCDD.EEFF".
	//
	//	Added in version 250.
	//	IPv4 address format: E.g. "127.0.0.1" or "192.168.0.1".
	//
	//	Added in version 250.
	//	IPv6 address format: E.g. "2001:0db8:85a3::8a2e:0370:7334" or "::1".
	//
	//	Added in version 250.
	//
	// The total length of each MAC address must be 4 (for IPv4 tunnel), 6 (for Ethernet), 16 (for IPv6 tunnel), or 20 (for InfiniBand).
	// This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option,
	// the list of hardware addresses defined prior to this is reset. Defaults to unset.
	//
	// Added in version 211.
	MACAddress systemdconf.Value

	// A whitespace-separated list of hardware's permanent addresses. While MACAddress= matches the device's current MAC
	// address, this matches the device's permanent MAC address, which may be different from the current one. Use full colon-,
	// hyphen- or dot-delimited hexadecimal, or IPv4 or IPv6 address format. This option may appear more than once, in which case
	// the lists are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this
	// is reset. Defaults to unset.
	//
	// Added in version 245.
	PermanentMACAddress systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the persistent path, as exposed by the udev property ID_PATH.
	//
	// Added in version 211.
	Path systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the driver currently bound to the device, as exposed by the
	// udev property ID_NET_DRIVER of its parent device, or if that is not set, the driver as exposed by ethtool -i of the device
	// itself. If the list is prefixed with a "!", the test is inverted.
	//
	// Added in version 211.
	Driver systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the device type, as exposed by networkctl list. If the list
	// is prefixed with a "!", the test is inverted. Some valid values are "ether", "loopback", "wlan", "wwan". Valid types are
	// named either from the udev "DEVTYPE" attribute, or "ARPHRD_" macros in linux/if_arp.h, so this is not comprehensive.
	//
	// Added in version 211.
	Type systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the device kind, as exposed by networkctl status INTERFACE
	// or ip -d link show INTERFACE. If the list is prefixed with a "!", the test is inverted. Some valid values are "bond", "bridge",
	// "gre", "tun", "veth". Valid kinds are given by netlink's "IFLA_INFO_KIND" attribute, so this is not comprehensive.
	//
	// Added in version 251.
	Kind systemdconf.Value

	// A whitespace-separated list of udev property names with their values after equals sign ("="). If multiple properties
	// are specified, the test results are ANDed. If the list is prefixed with a "!", the test is inverted. If a value contains white
	// spaces, then please quote whole key and value pair. If a value contains quotation, then please escape the quotation with
	// "\".
	//
	// Example: if a .link file has the following:
	//
	//	Property=ID_MODEL_ID=9999 "ID_VENDOR_FROM_DATABASE=vendor name" "KEY=with \"quotation\""
	//
	// then, the .link file matches only when an interface has all the above three properties.
	//
	// Added in version 243.
	Property systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the device name, as exposed by the udev property "INTERFACE".
	// This cannot be used to match on names that have already been changed from userspace. Caution is advised when matching on
	// kernel-assigned names, as they are known to be unstable between reboots.
	//
	// Added in version 218.
	OriginalName systemdconf.Value

	// Matches against the hostname or machine ID of the host. See ConditionHost= in systemd.unit for details. When prefixed
	// with an exclamation mark ("!"), the result is negated. If an empty string is assigned, the previously assigned value is
	// cleared.
	//
	// Added in version 211.
	Host systemdconf.Value

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation.
	// See ConditionVirtualization= in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result
	// is negated. If an empty string is assigned, the previously assigned value is cleared.
	//
	// Added in version 211.
	Virtualization systemdconf.Value

	// Checks whether a specific kernel command line option is set. See ConditionKernelCommandLine= in systemd.unit for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, the previously assigned
	// value is cleared.
	//
	// Added in version 211.
	KernelCommandLine systemdconf.Value

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See ConditionKernelVersion=
	// in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is
	// assigned, the previously assigned value is cleared.
	//
	// Added in version 237.
	KernelVersion systemdconf.Value

	// Checks whether the specified credential was passed to the systemd-udevd.service service. See System and Service Credentials
	// for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, the previously
	// assigned value is cleared.
	//
	// Added in version 252.
	Credential systemdconf.Value

	// Checks whether the system is running on a specific architecture. See ConditionArchitecture= in systemd.unit for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, the previously assigned
	// value is cleared.
	//
	// Added in version 211.
	Architecture systemdconf.Value

	// Checks whether the system is running on a machine with the specified firmware. See ConditionFirmware= in systemd.unit
	// for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, the previously
	// assigned value is cleared.
	//
	// Added in version 249.
	Firmware systemdconf.Value
}

// LinkLinkSection represents [Link] section, which specifies how the device should be configured
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html#%5BLink%5D%20Section%20Options for details).
type LinkLinkSection struct {
	systemdconf.Section

	// A description of the device.
	//
	// Added in version 211.
	Description systemdconf.Value

	// Set specified udev properties. This takes space separated list of key-value pairs concatenated with equal sign ("=").
	// Example:
	//
	//	Property=HOGE=foo BAR=baz
	//
	// This option supports simple specifier expansion, see the Specifiers section below. This option can be specified multiple
	// times. If an empty string is assigned, then the all previous assignments are cleared.
	//
	// This setting is useful to configure the "ID_NET_MANAGED_BY=" property which declares which network management service
	// shall manage the interface, which is respected by systemd-networkd and others. Use
	//
	//	Property=ID_NET_MANAGED_BY=io.systemd.Network
	//
	// to declare explicitly that systemd-networkd shall manage the interface, or set the property to something else to declare
	// explicitly it shall not do so. See systemd.network for details how this property is used to match interface names.
	//
	// Added in version 256.
	Property systemdconf.Value

	// Import specified udev properties from the saved database. This takes space separated list of property names. Example:
	//
	//	ImportProperty=HOGE BAR
	//
	// This option supports simple specifier expansion, see the Specifiers section below. This option can be specified multiple
	// times. If an empty string is assigned, then the all previous assignments are cleared.
	//
	// If the same property is also set in Property= in the above, then the imported property value will be overridden by the value
	// specified in Property=.
	//
	// Added in version 256.
	ImportProperty systemdconf.Value

	// Unset specified udev properties. This takes space separated list of property names. Example:
	//
	//	ImportProperty=HOGE BAR
	//
	// This option supports simple specifier expansion, see the Specifiers section below. This option can be specified multiple
	// times. If an empty string is assigned, then the all previous assignments are cleared.
	//
	// This setting is applied after ImportProperty= and Property= are applied. Hence, if the same property is specified in ImportProperty=
	// or Property=, then the imported or specified property value will be ignored, and the property will be unset.
	//
	// Added in version 256.
	UnsetProperty systemdconf.Value

	// The ifalias interface property is set to this value.
	//
	// Added in version 211.
	Alias systemdconf.Value

	// The policy by which the MAC address should be set. The available policies are:
	//
	//	persistent: If the hardware has a persistent MAC address, as most hardware should, and if it is used by the kernel, nothing
	//	is done. Otherwise, a new MAC address is generated which is guaranteed to be the same on every boot for the given machine and
	//	the given device, but which is otherwise random. This feature depends on ID_NET_NAME_* properties to exist for the link.
	//	On hardware where these properties are not set, the generation of a persistent MAC address will fail.
	//
	//	Added in version 211.
	//	random: If the kernel is using a random MAC address, nothing is done. Otherwise, a new address is randomly generated each
	//	time the device appears, typically at boot. Either way, the random address will have the "unicast" and "locally administered"
	//	bits set.
	//
	//	Added in version 211.
	//	none: Keeps the MAC address assigned by the kernel. Or use the MAC address specified in MACAddress=.
	//
	//	Added in version 227.
	//
	// An empty string assignment is equivalent to setting "none".
	//
	// Added in version 211.
	MACAddressPolicy systemdconf.Value

	// The interface MAC address to use. For this setting to take effect, MACAddressPolicy= must either be unset, empty, or "none".
	//
	// Added in version 211.
	MACAddress systemdconf.Value

	// An ordered, space-separated list of policies by which the interface name should be set. NamePolicy= may be disabled by
	// specifying net.ifnames=0 on the kernel command line. Each of the policies may fail, and the first successful one is used.
	// The name is not set directly, but is exported to udev as the property ID_NET_NAME, which is, by default, used by a udev, rule
	// to set NAME. The available policies are:
	//
	//	kernel: If the kernel claims that the name it has set for a device is predictable, then no renaming is performed.
	//
	//	Added in version 216.
	//	database: The name is set based on entries in the udev's Hardware Database with the key ID_NET_NAME_FROM_DATABASE.
	//
	//	Added in version 211.
	//	onboard: The name is set based on information given by the firmware for on-board devices, as exported by the udev property
	//	ID_NET_NAME_ONBOARD. See systemd.net-naming-scheme.
	//
	//	Added in version 211.
	//	slot: The name is set based on information given by the firmware for hot-plug devices, as exported by the udev property
	//	ID_NET_NAME_SLOT. See systemd.net-naming-scheme.
	//
	//	Added in version 211.
	//	path: The name is set based on the device's physical location, as exported by the udev property ID_NET_NAME_PATH. See
	//	systemd.net-naming-scheme.
	//
	//	Added in version 211.
	//	mac: The name is set based on the device's persistent MAC address, as exported by the udev property ID_NET_NAME_MAC. See
	//	systemd.net-naming-scheme.
	//
	//	Added in version 211.
	//	keep: If the device already had a name given by userspace (as part of creation of the device or a rename), keep it.
	//
	//	Added in version 241.
	//
	// Added in version 211.
	NamePolicy systemdconf.Value

	// The interface name to use. This option has lower precedence than NamePolicy=, so for this setting to take effect, NamePolicy=
	// must either be unset, empty, disabled, or all policies configured there must fail. Also see the example below with "Name=dmz0".
	//
	// Note that specifying a name that the kernel might use for another interface (for example "eth0") is dangerous because the
	// name assignment done by udev will race with the assignment done by the kernel, and only one interface may use the name. Depending
	// on the order of operations, either udev or the kernel will win, making the naming unpredictable. It is best to use some different
	// prefix, for example "internal0"/"external0" or "lan0"/"lan1"/"lan3".
	//
	// Interface names must have a minimum length of 1 character and a maximum length of 15 characters, and may contain any 7bit
	// ASCII character, with the exception of control characters, ":", "/" and "%". While "." is an allowed character, it is recommended
	// to avoid it when naming interfaces as various tools (such as resolvconf) use it as separator character. Also, fully numeric
	// interface names are not allowed (in order to avoid ambiguity with interface specification by numeric indexes), nor are
	// the special strings ".", "..", "all" and "default".
	//
	// Added in version 211.
	Name systemdconf.Value

	// A space-separated list of policies by which the interface's alternative names should be set. Each of the policies may fail,
	// and all successful policies are used. The available policies are "database", "onboard", "slot", "path", and "mac". If
	// the kernel does not support the alternative names, then this setting will be ignored.
	//
	// Added in version 245.
	AlternativeNamesPolicy systemdconf.Value

	// The alternative interface name to use. This option can be specified multiple times. If the empty string is assigned to this
	// option, the list is reset, and all prior assignments have no effect. If the kernel does not support the alternative names,
	// then this setting will be ignored.
	//
	// Alternative interface names may be used to identify interfaces in various tools. In contrast to the primary name (as configured
	// with Name= above) there may be multiple alternative names referring to the same interface. Alternative names may have
	// a maximum length of 127 characters, in contrast to the 15 allowed for the primary interface name, but otherwise are subject
	// to the same naming constraints.
	//
	// Added in version 245.
	AlternativeName systemdconf.Value

	// Specifies the device's number of transmit queues. An integer in the range 1…4096. When unset, the kernel's default will
	// be used.
	//
	// Added in version 248.
	TransmitQueues systemdconf.Value

	// Specifies the device's number of receive queues. An integer in the range 1…4096. When unset, the kernel's default will
	// be used.
	//
	// Added in version 248.
	ReceiveQueues systemdconf.Value

	// Specifies the transmit queue length of the device in number of packets. An unsigned integer in the range 0…4294967294.
	// When unset, the kernel's default will be used.
	//
	// Added in version 248.
	TransmitQueueLength systemdconf.Value

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G are supported and are understood
	// to the base of 1024.
	//
	// Added in version 211.
	MTUBytes systemdconf.Value

	// The speed to set for the device, the value is rounded down to the nearest Mbps. The usual suffixes K, M, G are supported and
	// are understood to the base of 1000.
	//
	// Added in version 211.
	BitsPerSecond systemdconf.Value

	// The duplex mode to set for the device. The accepted values are half and full.
	//
	// Added in version 211.
	Duplex systemdconf.Value

	// Takes a boolean. If set to yes, automatic negotiation of transmission parameters is enabled. Autonegotiation is a procedure
	// by which two connected ethernet devices choose common transmission parameters, such as speed, duplex mode, and flow control.
	// When unset, the kernel's default will be used.
	//
	// Note that if autonegotiation is enabled, speed and duplex settings are read-only. If autonegotiation is disabled, speed
	// and duplex settings are writable if the driver supports multiple link modes.
	//
	// Added in version 233.
	AutoNegotiation systemdconf.Value

	// The Wake-on-LAN policy to set for the device. Takes the special value "off" which disables Wake-on-LAN, or space separated
	// list of the following words:
	//
	//	phy: Wake on PHY activity.
	//
	//	Added in version 211.
	//	unicast: Wake on unicast messages.
	//
	//	Added in version 235.
	//	multicast: Wake on multicast messages.
	//
	//	Added in version 235.
	//	broadcast: Wake on broadcast messages.
	//
	//	Added in version 235.
	//	arp: Wake on ARP.
	//
	//	Added in version 235.
	//	magic: Wake on receipt of a magic packet.
	//
	//	Added in version 211.
	//	secureon: Enable SecureOn password for MagicPacket. Implied when WakeOnLanPassword= is specified. If specified without
	//	WakeOnLanPassword= option, then the password is read from the credential "LINK.link.wol.password" (e.g., "60-foo.link.wol.password"),
	//	and if the credential not found, then read from "wol.password". See ImportCredential=/LoadCredential=/SetCredential=
	//	in systemd.exec for details. The password in the credential, must be 6 bytes in hex format with each byte separated by a colon
	//	(":") like an Ethernet MAC address, e.g., "aa:bb:cc:dd:ee:ff".
	//
	//	Added in version 235.
	//
	// Defaults to unset, and the device's default will be used. This setting can be specified multiple times. If an empty string
	// is assigned, then the all previous assignments are cleared.
	//
	// Added in version 211.
	WakeOnLan systemdconf.Value

	// Specifies the SecureOn password for MagicPacket. Takes an absolute path to a regular file or an AF_UNIX stream socket,
	// or the plain password. When a path to a regular file is specified, the password is read from it. When an AF_UNIX stream socket
	// is specified, a connection is made to it and the password is read from it. The password must be 6 bytes in hex format with each
	// byte separated by a colon (":") like an Ethernet MAC address, e.g., "aa:bb:cc:dd:ee:ff". This implies WakeOnLan=secureon.
	// Defaults to unset, and the current value will not be changed.
	//
	// Added in version 250.
	WakeOnLanPassword systemdconf.Value

	// The port option is used to select the device port. The supported values are:
	//
	//	tp: An Ethernet interface using Twisted-Pair cable as the medium.
	//
	//	Added in version 234.
	//	aui: Attachment Unit Interface (AUI). Normally used with hubs.
	//
	//	Added in version 234.
	//	bnc: An Ethernet interface using BNC connectors and co-axial cable.
	//
	//	Added in version 234.
	//	mii: An Ethernet interface using a Media Independent Interface (MII).
	//
	//	Added in version 234.
	//	fibre: An Ethernet interface using Optical Fibre as the medium.
	//
	//	Added in version 234.
	//
	// Added in version 234.
	Port systemdconf.Value

	// This sets what speeds and duplex modes of operation are advertised for auto-negotiation. This implies "AutoNegotiation=yes".
	// The supported values are:
	//
	//	+----------------------------+--------------+-------------+
	//	|         ADVERTISE          | SPEED (MBPS) | DUPLEX MODE |
	//	+----------------------------+--------------+-------------+
	//	| 10baset-full               |           10 | full        |
	//	| 10baset1brr-full           |           10 | full        |
	//	| 10baset1l-full             |           10 | full        |
	//	| 10baset1s-full             |           10 | full        |
	//	| 10baset-half               |           10 | half        |
	//	| 10baset1s-half             |           10 | half        |
	//	| 10baset1s-p2mp-half        |           10 | half        |
	//	| 100basefx-full             |          100 | full        |
	//	| 100baset-full              |          100 | full        |
	//	| 100baset1-full             |          100 | full        |
	//	| 100basefx-half             |          100 | half        |
	//	| 100baset-half              |          100 | half        |
	//	| 1000basekx-full            |         1000 | full        |
	//	| 1000baset-full             |         1000 | full        |
	//	| 1000baset1-full            |         1000 | full        |
	//	| 1000basex-full             |         1000 | full        |
	//	| 1000baset-half             |         1000 | half        |
	//	| 2500baset-full             |         2500 | full        |
	//	| 2500basex-full             |         2500 | full        |
	//	| 5000baset-full             |         5000 | full        |
	//	| 10000baser-fec             |        10000 |             |
	//	| 10000basecr-full           |        10000 | full        |
	//	| 10000baseer-full           |        10000 | full        |
	//	| 10000basekr-full           |        10000 | full        |
	//	| 10000basekx4-full          |        10000 | full        |
	//	| 10000baselr-full           |        10000 | full        |
	//	| 10000baselrm-full          |        10000 | full        |
	//	| 10000basesr-full           |        10000 | full        |
	//	| 10000baset-full            |        10000 | full        |
	//	| 20000basekr2-full          |        20000 | full        |
	//	| 20000basemld2-full         |        20000 | full        |
	//	| 25000basecr-full           |        25000 | full        |
	//	| 25000basekr-full           |        25000 | full        |
	//	| 25000basesr-full           |        25000 | full        |
	//	| 40000basecr4-full          |        40000 | full        |
	//	| 40000basekr4-full          |        40000 | full        |
	//	| 40000baselr4-full          |        40000 | full        |
	//	| 40000basesr4-full          |        40000 | full        |
	//	| 50000basecr-full           |        50000 | full        |
	//	| 50000basecr2-full          |        50000 | full        |
	//	| 50000basedr-full           |        50000 | full        |
	//	| 50000basekr-full           |        50000 | full        |
	//	| 50000basekr2-full          |        50000 | full        |
	//	| 50000baselr-er-fr-full     |        50000 | full        |
	//	| 50000basesr-full           |        50000 | full        |
	//	| 50000basesr2-full          |        50000 | full        |
	//	| 56000basecr4-full          |        56000 | full        |
	//	| 56000basekr4-full          |        56000 | full        |
	//	| 56000baselr4-full          |        56000 | full        |
	//	| 56000basesr4-full          |        56000 | full        |
	//	| 100000basecr-full          |       100000 | full        |
	//	| 100000basecr2-full         |       100000 | full        |
	//	| 100000basecr4-full         |       100000 | full        |
	//	| 100000basedr-full          |       100000 | full        |
	//	| 100000basedr2-full         |       100000 | full        |
	//	| 100000basekr-full          |       100000 | full        |
	//	| 100000basekr2-full         |       100000 | full        |
	//	| 100000basekr4-full         |       100000 | full        |
	//	| 100000baselr-er-fr-full    |       100000 | full        |
	//	| 100000baselr2-er2-fr2-full |       100000 | full        |
	//	| 100000baselr4-er4-full     |       100000 | full        |
	//	| 100000basesr-full          |       100000 | full        |
	//	| 100000basesr2-full         |       100000 | full        |
	//	| 100000basesr4-full         |       100000 | full        |
	//	| 200000basecr2-full         |       200000 | full        |
	//	| 200000basecr4-full         |       200000 | full        |
	//	| 200000basedr2-full         |       200000 | full        |
	//	| 200000basedr4-full         |       200000 | full        |
	//	| 200000basekr2-full         |       200000 | full        |
	//	| 200000basekr4-full         |       200000 | full        |
	//	| 200000baselr2-er2-fr2-full |       200000 | full        |
	//	| 200000baselr4-er4-fr4-full |       200000 | full        |
	//	| 200000basesr2-full         |       200000 | full        |
	//	| 200000basesr4-full         |       200000 | full        |
	//	| 400000basecr4-full         |       400000 | full        |
	//	| 400000basecr8-full         |       400000 | full        |
	//	| 400000basedr4-full         |       400000 | full        |
	//	| 400000basedr8-full         |       400000 | full        |
	//	| 400000basekr4-full         |       400000 | full        |
	//	| 400000basekr8-full         |       400000 | full        |
	//	| 400000baselr4-er4-fr4-full |       400000 | full        |
	//	| 400000baselr8-er8-fr8-full |       400000 | full        |
	//	| 400000basesr4-full         |       400000 | full        |
	//	| 400000basesr8-full         |       400000 | full        |
	//	| 800000basecr8-full         |       800000 | full        |
	//	| 800000basedr8-2-full       |       800000 | full        |
	//	| 800000basedr8-full         |       800000 | full        |
	//	| 800000basekr8-full         |       800000 | full        |
	//	| 800000basesr8-full         |       800000 | full        |
	//	| 800000basevr8-full         |       800000 | full        |
	//	| asym-pause                 |              |             |
	//	| aui                        |              |             |
	//	| autonegotiation            |              |             |
	//	| backplane                  |              |             |
	//	| bnc                        |              |             |
	//	| fec-baser                  |              |             |
	//	| fec-llrs                   |              |             |
	//	| fec-none                   |              |             |
	//	| fec-rs                     |              |             |
	//	| fibre                      |              |             |
	//	| mii                        |              |             |
	//	| pause                      |              |             |
	//	| tp                         |              |             |
	//	+----------------------------+--------------+-------------+
	//
	// By default, this is unset, i.e. all possible modes will be advertised. This option may be specified more than once, in which
	// case all specified speeds and modes are advertised. If the empty string is assigned to this option, the list is reset, and
	// all prior assignments have no effect.
	//
	// Added in version 240.
	Advertise systemdconf.Value

	// Takes a boolean. If set to true, hardware offload for checksumming of ingress network packets is enabled. When unset, the
	// kernel's default will be used.
	//
	// Added in version 245.
	ReceiveChecksumOffload systemdconf.Value

	// Takes a boolean. If set to true, hardware offload for checksumming of egress network packets is enabled. When unset, the
	// kernel's default will be used.
	//
	// Added in version 245.
	TransmitChecksumOffload systemdconf.Value

	// Takes a boolean. If set to true, TCP Segmentation Offload (TSO) is enabled. When unset, the kernel's default will be used.
	//
	// Added in version 232.
	TCPSegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, TCP6 Segmentation Offload (tx-tcp6-segmentation) is enabled. When unset, the kernel's
	// default will be used.
	//
	// Added in version 235.
	TCP6SegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, Generic Segmentation Offload (GSO) is enabled. When unset, the kernel's default will be
	// used.
	//
	// Added in version 232.
	GenericSegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, Generic Receive Offload (GRO) is enabled. When unset, the kernel's default will be used.
	//
	// Added in version 232.
	GenericReceiveOffload systemdconf.Value

	// Takes a boolean. If set to true, hardware accelerated Generic Receive Offload (GRO) is enabled. When unset, the kernel's
	// default will be used.
	//
	// Added in version 250.
	GenericReceiveOffloadHardware systemdconf.Value

	// Takes a boolean. If set to true, Large Receive Offload (LRO) is enabled. When unset, the kernel's default will be used.
	//
	// Added in version 232.
	LargeReceiveOffload systemdconf.Value

	// Configures Receive Packet Steering (RPS) list of CPUs to which RPS may forward traffic. Takes a list of CPU indices or ranges
	// separated by either whitespace or commas. Alternatively, takes the special value "all", which will include all available
	// CPUs in the mask. CPU ranges are specified by the lower and upper CPU indices separated by a dash (e.g. "2-6"). This option
	// may be specified more than once, in which case the specified list of CPU ranges are merged. If an empty string is assigned,
	// the list is reset, all assignments prior to this will have no effect. Defaults to unset and RPS CPU list is unchanged. To disable
	// RPS when it was previously enabled, use the special value "disable".
	//
	// Added in version 256.
	ReceivePacketSteeringCPUMask systemdconf.Value

	// Takes a boolean. If set to true, receive VLAN CTAG hardware acceleration is enabled. When unset, the kernel's default will
	// be used.
	//
	// Added in version 250.
	ReceiveVLANCTAGHardwareAcceleration systemdconf.Value

	// Takes a boolean. If set to true, transmit VLAN CTAG hardware acceleration is enabled. When unset, the kernel's default
	// will be used.
	//
	// Added in version 250.
	TransmitVLANCTAGHardwareAcceleration systemdconf.Value

	// Takes a boolean. If set to true, receive filtering on VLAN CTAGs is enabled. When unset, the kernel's default will be used.
	//
	// Added in version 250.
	ReceiveVLANCTAGFilter systemdconf.Value

	// Takes a boolean. If set to true, transmit VLAN STAG hardware acceleration is enabled. When unset, the kernel's default
	// will be used.
	//
	// Added in version 250.
	TransmitVLANSTAGHardwareAcceleration systemdconf.Value

	// Takes a boolean. If set to true, receive N-tuple filters and actions are enabled. When unset, the kernel's default will
	// be used.
	//
	// Added in version 250.
	NTupleFilter systemdconf.Value

	// Specifies the number of receive, transmit, other, or combined channels, respectively. Takes an unsigned integer in the
	// range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the
	// number will not be changed. Defaults to unset.
	//
	// Added in version 239.
	RxChannels systemdconf.Value

	// Specifies the number of receive, transmit, other, or combined channels, respectively. Takes an unsigned integer in the
	// range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the
	// number will not be changed. Defaults to unset.
	//
	// Added in version 239.
	TxChannels systemdconf.Value

	// Specifies the number of receive, transmit, other, or combined channels, respectively. Takes an unsigned integer in the
	// range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the
	// number will not be changed. Defaults to unset.
	//
	// Added in version 239.
	OtherChannels systemdconf.Value

	// Specifies the number of receive, transmit, other, or combined channels, respectively. Takes an unsigned integer in the
	// range 1…4294967295 or "max". If set to "max", the advertised maximum value of the hardware will be used. When unset, the
	// number will not be changed. Defaults to unset.
	//
	// Added in version 239.
	CombinedChannels systemdconf.Value

	// Specifies the maximum number of pending packets in the NIC receive buffer, mini receive buffer, jumbo receive buffer,
	// or transmit buffer, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised
	// maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	//
	// Added in version 244.
	RxBufferSize systemdconf.Value

	// Specifies the maximum number of pending packets in the NIC receive buffer, mini receive buffer, jumbo receive buffer,
	// or transmit buffer, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised
	// maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	//
	// Added in version 244.
	RxMiniBufferSize systemdconf.Value

	// Specifies the maximum number of pending packets in the NIC receive buffer, mini receive buffer, jumbo receive buffer,
	// or transmit buffer, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised
	// maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	//
	// Added in version 244.
	RxJumboBufferSize systemdconf.Value

	// Specifies the maximum number of pending packets in the NIC receive buffer, mini receive buffer, jumbo receive buffer,
	// or transmit buffer, respectively. Takes an unsigned integer in the range 1…4294967295 or "max". If set to "max", the advertised
	// maximum value of the hardware will be used. When unset, the number will not be changed. Defaults to unset.
	//
	// Added in version 244.
	TxBufferSize systemdconf.Value

	// Takes a boolean. When set, enables receive flow control, also known as the ethernet receive PAUSE message (generate and
	// send ethernet PAUSE frames). When unset, the kernel's default will be used.
	//
	// Added in version 246.
	RxFlowControl systemdconf.Value

	// Takes a boolean. When set, enables transmit flow control, also known as the ethernet transmit PAUSE message (respond to
	// received ethernet PAUSE frames). When unset, the kernel's default will be used.
	//
	// Added in version 246.
	TxFlowControl systemdconf.Value

	// Takes a boolean. When set, auto negotiation enables the interface to exchange state advertisements with the connected
	// peer so that the two devices can agree on the ethernet PAUSE configuration. When unset, the kernel's default will be used.
	//
	// Added in version 246.
	AutoNegotiationFlowControl systemdconf.Value

	// Specifies the maximum size of a Generic Segment Offload (GSO) packet the device should accept. The usual suffixes K, M,
	// G are supported and are understood to the base of 1024. An unsigned integer in the range 1…65536. Defaults to unset.
	//
	// Added in version 248.
	GenericSegmentOffloadMaxBytes systemdconf.Value

	// Specifies the maximum number of Generic Segment Offload (GSO) segments the device should accept. An unsigned integer
	// in the range 1…65535. Defaults to unset.
	//
	// Added in version 248.
	GenericSegmentOffloadMaxSegments systemdconf.Value

	// Boolean properties that, when set, enable/disable adaptive Rx/Tx coalescing if the hardware supports it. When unset,
	// the kernel's default will be used.
	//
	// Added in version 250.
	UseAdaptiveRxCoalesce systemdconf.Value

	// Boolean properties that, when set, enable/disable adaptive Rx/Tx coalescing if the hardware supports it. When unset,
	// the kernel's default will be used.
	//
	// Added in version 250.
	UseAdaptiveTxCoalesce systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxCoalesceSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxCoalesceIrqSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxCoalesceLowSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxCoalesceHighSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxCoalesceSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxCoalesceIrqSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxCoalesceLowSec systemdconf.Value

	// These properties configure the delay before Rx/Tx interrupts are generated after a packet is sent/received. The "Irq"
	// properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect when the
	// packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if adaptive
	// Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxCoalesceHighSec systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxMaxCoalescedFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxMaxCoalescedIrqFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxMaxCoalescedLowFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	RxMaxCoalescedHighFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxMaxCoalescedFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxMaxCoalescedIrqFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxMaxCoalescedLowFrames systemdconf.Value

	// These properties configure the maximum number of frames that are sent/received before a Rx/Tx interrupt is generated.
	// The "Irq" properties come into effect when the host is servicing an IRQ. The "Low" and "High" properties come into effect
	// when the packet rate drops below the low packet rate threshold or exceeds the high packet rate threshold respectively if
	// adaptive Rx/Tx coalescing is enabled. When unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	TxMaxCoalescedHighFrames systemdconf.Value

	// These properties configure the low and high packet rate (expressed in packets per second) threshold respectively and
	// are used to determine when the corresponding coalescing settings for low and high packet rates come into effect if adaptive
	// Rx/Tx coalescing is enabled. If unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	CoalescePacketRateLow systemdconf.Value

	// These properties configure the low and high packet rate (expressed in packets per second) threshold respectively and
	// are used to determine when the corresponding coalescing settings for low and high packet rates come into effect if adaptive
	// Rx/Tx coalescing is enabled. If unset, the kernel's defaults will be used.
	//
	// Added in version 250.
	CoalescePacketRateHigh systemdconf.Value

	// Configures how often to sample the packet rate used for adaptive Rx/Tx coalescing. This property cannot be zero. This lowest
	// time granularity supported by this property is seconds. Partial seconds will be rounded up before being passed to the kernel.
	// If unset, the kernel's default will be used.
	//
	// Added in version 250.
	CoalescePacketRateSampleIntervalSec systemdconf.Value

	// How long to delay driver in-memory statistics block updates. If the driver does not have an in-memory statistic block,
	// this property is ignored. This property cannot be zero. If unset, the kernel's default will be used.
	//
	// Added in version 250.
	StatisticsBlockCoalesceSec systemdconf.Value

	// Specifies the medium dependent interface (MDI) mode for the interface. A MDI describes the interface from a physical layer
	// implementation to the physical medium used to carry the transmission. Takes one of the following words: "straight" (or
	// equivalently: "mdi"), "crossover" (or equivalently: "mdi-x", "mdix"), and "auto". When "straight", the MDI straight
	// through mode will be used. When "crossover", the MDI crossover (MDI-X) mode will be used. When "auto", the MDI status is
	// automatically detected. Defaults to unset, and the kernel's default will be used.
	//
	// Added in version 251.
	MDI systemdconf.Value

	// Specifies the number of SR-IOV virtual functions. Takes an integer in the range 0…2147483647. Defaults to unset, and automatically
	// determined from the values specified in the VirtualFunction= settings in the [SR-IOV] sections.
	//
	// Added in version 251.
	SRIOVVirtualFunctions systemdconf.Value `systemd:"SR-IOVVirtualFunctions"`
}

// LinkSRIOVSection represents [SR-IOV] section, which provides the ability to partition a single physical PCI resource into virtual PCI functions which can then be e.g. injected into a VM
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html#%5BSR-IOV%5D%20Section%20Options for details).
type LinkSRIOVSection struct {
	systemdconf.Section

	// Specifies a Virtual Function (VF), lightweight PCIe function designed solely to move data in and out. Takes an integer
	// in the range 0…2147483646. This option is compulsory.
	//
	// Added in version 251.
	VirtualFunction systemdconf.Value

	// Specifies VLAN ID of the virtual function. Takes an integer in the range 1…4095.
	//
	// Added in version 251.
	VLANId systemdconf.Value

	// Specifies quality of service of the virtual function. Takes an integer in the range 1…4294967294.
	//
	// Added in version 251.
	QualityOfService systemdconf.Value

	// Specifies VLAN protocol of the virtual function. Takes "802.1Q" or "802.1ad".
	//
	// Added in version 251.
	VLANProtocol systemdconf.Value

	// Takes a boolean. Controls the MAC spoof checking. When unset, the kernel's default will be used.
	//
	// Added in version 251.
	MACSpoofCheck systemdconf.Value

	// Takes a boolean. Toggle the ability of querying the receive side scaling (RSS) configuration of the virtual function (VF).
	// The VF RSS information like RSS hash key may be considered sensitive on some devices where this information is shared between
	// VF and the physical function (PF). When unset, the kernel's default will be used.
	//
	// Added in version 251.
	QueryReceiveSideScaling systemdconf.Value

	// Takes a boolean. Allows one to set trust mode of the virtual function (VF). When set, VF users can set a specific feature which
	// may impact security and/or performance. When unset, the kernel's default will be used.
	//
	// Added in version 251.
	Trust systemdconf.Value

	// Allows one to set the link state of the virtual function (VF). Takes a boolean or a special value "auto". Setting to "auto"
	// means a reflection of the physical function (PF) link state, "yes" lets the VF to communicate with other VFs on this host
	// even if the PF link state is down, "no" causes the hardware to drop any packets sent by the VF. When unset, the kernel's default
	// will be used.
	//
	// Added in version 251.
	LinkState systemdconf.Value

	// Specifies the MAC address for the virtual function.
	//
	// Added in version 251.
	MACAddress systemdconf.Value
}
