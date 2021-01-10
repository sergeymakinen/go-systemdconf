// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf/v2"

// LinkFile represents systemd.link — Network device configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html for details)
type LinkFile struct {
	systemdconf.File

	Match LinkMatchSection // [Match] section, which determines if a given link file may be applied to a given device
	Link  LinkLinkSection  // [Link] section specifying how the device should be configured
}

// LinkMatchSection represents [Match] section, which determines if a given link file may be applied to a given device
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html#%5BMatch%5D%20Section%20Options for details)
type LinkMatchSection struct {
	systemdconf.Section

	// A whitespace-separated list of hardware addresses. Use full colon-, hyphen- or dot-delimited hexadecimal. See the example
	// below. This option may appear more than once, in which case the lists are merged. If the empty string is assigned to this option,
	// the list of hardware addresses defined prior to this is reset.
	//
	// Example:
	//
	// 	MACAddress=01:23:45:67:89:ab 00-11-22-33-44-55 AABB.CCDD.EEFF
	MACAddress systemdconf.Value

	// A whitespace-separated list of hardware's permanent addresses. While MACAddress= matches the device's current MAC
	// address, this matches the device's permanent MAC address, which may be different from the current one. Use full colon-,
	// hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists are merged. If the
	// empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	PermanentMACAddress systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the persistent path, as exposed by the udev property ID_PATH.
	Path systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the driver currently bound to the device, as exposed by the
	// udev property ID_NET_DRIVER of its parent device, or if that is not set, the driver as exposed by ethtool -i of the device
	// itself. If the list is prefixed with a "!", the test is inverted.
	Driver systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the device type, as exposed by networkctl status. If the list
	// is prefixed with a "!", the test is inverted.
	Type systemdconf.Value

	// A whitespace-separated list of udev property name with its value after a equal ("="). If multiple properties are specified,
	// the test results are ANDed. If the list is prefixed with a "!", the test is inverted. If a value contains white spaces, then
	// please quote whole key and value pair. If a value contains quotation, then please escape the quotation with "\".
	//
	// Example: if a .link file has the following:
	//
	// 	Property=ID_MODEL_ID=9999 "ID_VENDOR_FROM_DATABASE=vendor name" "KEY=with \"quotation\""
	//
	// then, the .link file matches only when an interface has all the above three properties.
	Property systemdconf.Value

	// A whitespace-separated list of shell-style globs matching the device name, as exposed by the udev property "INTERFACE".
	// This cannot be used to match on names that have already been changed from userspace. Caution is advised when matching on
	// kernel-assigned names, as they are known to be unstable between reboots.
	OriginalName systemdconf.Value

	// Matches against the hostname or machine ID of the host. See ConditionHost= in systemd.unit for details. When prefixed
	// with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is
	// cleared.
	Host systemdconf.Value

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation.
	// See ConditionVirtualization= in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result
	// is negated. If an empty string is assigned, then previously assigned value is cleared.
	Virtualization systemdconf.Value

	// Checks whether a specific kernel command line option is set. See ConditionKernelCommandLine= in systemd.unit for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned
	// value is cleared.
	KernelCommandLine systemdconf.Value

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See ConditionKernelVersion=
	// in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is
	// assigned, then previously assigned value is cleared.
	KernelVersion systemdconf.Value

	// Checks whether the system is running on a specific architecture. See ConditionArchitecture= in systemd.unit for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned
	// value is cleared.
	Architecture systemdconf.Value
}

// LinkLinkSection represents [Link] section specifying how the device should be configured
// (see https://www.freedesktop.org/software/systemd/man/systemd.link.html#%5BLink%5D%20Section%20Options for details)
type LinkLinkSection struct {
	systemdconf.Section

	// A description of the device.
	Description systemdconf.Value

	// The ifalias interface property is set to this value.
	Alias systemdconf.Value

	// The policy by which the MAC address should be set. The available policies are:
	//
	// 	persistent: If the hardware has a persistent MAC address, as most hardware should, and if it is used by the kernel, nothing
	// 	is done. Otherwise, a new MAC address is generated which is guaranteed to be the same on every boot for the given machine and
	// 	the given device, but which is otherwise random. This feature depends on ID_NET_NAME_* properties to exist for the link.
	// 	On hardware where these properties are not set, the generation of a persistent MAC address will fail.
	// 	random: If the kernel is using a random MAC address, nothing is done. Otherwise, a new address is randomly generated each
	// 	time the device appears, typically at boot. Either way, the random address will have the "unicast" and "locally administered"
	// 	bits set.
	// 	none: Keeps the MAC address assigned by the kernel. Or use the MAC address specified in MACAddress=.
	//
	// An empty string assignment is equivalent to setting "none".
	MACAddressPolicy systemdconf.Value

	// The interface MAC address to use. For this setting to take effect, MACAddressPolicy= must either be unset, empty, or "none".
	MACAddress systemdconf.Value

	// An ordered, space-separated list of policies by which the interface name should be set. NamePolicy= may be disabled by
	// specifying net.ifnames=0 on the kernel command line. Each of the policies may fail, and the first successful one is used.
	// The name is not set directly, but is exported to udev as the property ID_NET_NAME, which is, by default, used by a udev, rule
	// to set NAME. The available policies are:
	//
	// 	kernel: If the kernel claims that the name it has set for a device is predictable, then no renaming is performed.
	// 	database: The name is set based on entries in the udev's Hardware Database with the key ID_NET_NAME_FROM_DATABASE.
	// 	onboard: The name is set based on information given by the firmware for on-board devices, as exported by the udev property
	// 	ID_NET_NAME_ONBOARD. See systemd.net-naming-scheme.
	// 	slot: The name is set based on information given by the firmware for hot-plug devices, as exported by the udev property
	// 	ID_NET_NAME_SLOT. See systemd.net-naming-scheme.
	// 	path: The name is set based on the device's physical location, as exported by the udev property ID_NET_NAME_PATH. See
	// 	systemd.net-naming-scheme.
	// 	mac: The name is set based on the device's persistent MAC address, as exported by the udev property ID_NET_NAME_MAC. See
	// 	systemd.net-naming-scheme.
	// 	keep: If the device already had a name given by userspace (as part of creation of the device or a rename), keep it.
	NamePolicy systemdconf.Value

	// The interface name to use. This option has lower precedence than NamePolicy=, so for this setting to take effect, NamePolicy=
	// must either be unset, empty, disabled, or all policies configured there must fail. Also see the example below with "Name=dmz0".
	//
	// Note that specifying a name that the kernel might use for another interface (for example "eth0") is dangerous because the
	// name assignment done by udev will race with the assignment done by the kernel, and only one interface may use the name. Depending
	// on the order of operations, either udev or the kernel will win, making the naming unpredictable. It is best to use some different
	// prefix, for example "internal0"/"external0" or "lan0"/"lan1"/"lan3".
	Name systemdconf.Value

	// A space-separated list of policies by which the interface's alternative names should be set. Each of the policies may fail,
	// and all successful policies are used. The available policies are "database", "onboard", "slot", "path", and "mac". If
	// the kernel does not support the alternative names, then this setting will be ignored.
	AlternativeNamesPolicy systemdconf.Value

	// The alternative interface name to use. This option can be specified multiple times. If the empty string is assigned to this
	// option, the list is reset, and all prior assignments have no effect. If the kernel does not support the alternative names,
	// then this setting will be ignored.
	AlternativeName systemdconf.Value

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G, are supported and are understood
	// to the base of 1024.
	MTUBytes systemdconf.Value

	// The speed to set for the device, the value is rounded down to the nearest Mbps. The usual suffixes K, M, G, are supported and
	// are understood to the base of 1000.
	BitsPerSecond systemdconf.Value

	// The duplex mode to set for the device. The accepted values are half and full.
	Duplex systemdconf.Value

	// Takes a boolean. If set to yes, automatic negotiation of transmission parameters is enabled. Autonegotiation is a procedure
	// by which two connected ethernet devices choose common transmission parameters, such as speed, duplex mode, and flow control.
	// When unset, the kernel's default will be used.
	//
	// Note that if autonegotiation is enabled, speed and duplex settings are read-only. If autonegotiation is disabled, speed
	// and duplex settings are writable if the driver supports multiple link modes.
	AutoNegotiation systemdconf.Value

	// The Wake-on-LAN policy to set for the device. The supported values are:
	//
	// 	phy: Wake on PHY activity.
	// 	unicast: Wake on unicast messages.
	// 	multicast: Wake on multicast messages.
	// 	broadcast: Wake on broadcast messages.
	// 	arp: Wake on ARP.
	// 	magic: Wake on receipt of a magic packet.
	// 	secureon: Enable secureon(tm) password for MagicPacket(tm).
	// 	off: Never wake.
	//
	// Defaults to off.
	WakeOnLan systemdconf.Value

	// The port option is used to select the device port. The supported values are:
	//
	// 	tp: An Ethernet interface using Twisted-Pair cable as the medium.
	// 	aui: Attachment Unit Interface (AUI). Normally used with hubs.
	// 	bnc: An Ethernet interface using BNC connectors and co-axial cable.
	// 	mii: An Ethernet interface using a Media Independent Interface (MII).
	// 	fibre: An Ethernet interface using Optical Fibre as the medium.
	Port systemdconf.Value

	// This sets what speeds and duplex modes of operation are advertised for auto-negotiation. This implies "AutoNegotiation=yes".
	// The supported values are:
	//
	// 	+--------------------+--------------+-------------+
	// 	|     ADVERTISE      | SPEED (MBPS) | DUPLEX MODE |
	// 	+--------------------+--------------+-------------+
	// 	| 10baset-half       |           10 | half        |
	// 	| 10baset-full       |           10 | full        |
	// 	| 100baset-half      |          100 | half        |
	// 	| 100baset-full      |          100 | full        |
	// 	| 1000baset-half     |         1000 | half        |
	// 	| 1000baset-full     |         1000 | full        |
	// 	| 10000baset-full    |        10000 | full        |
	// 	| 2500basex-full     |         2500 | full        |
	// 	| 1000basekx-full    |         1000 | full        |
	// 	| 10000basekx4-full  |        10000 | full        |
	// 	| 10000basekr-full   |        10000 | full        |
	// 	| 10000baser-fec     |        10000 | full        |
	// 	| 20000basemld2-full |        20000 | full        |
	// 	| 20000basekr2-full  |        20000 | full        |
	// 	+--------------------+--------------+-------------+
	//
	// By default this is unset, i.e. all possible modes will be advertised. This option may be specified more than once, in which
	// case all specified speeds and modes are advertised. If the empty string is assigned to this option, the list is reset, and
	// all prior assignments have no effect.
	Advertise systemdconf.Value

	// Takes a boolean. If set to true, the hardware offload for checksumming of ingress network packets is enabled. When unset,
	// the kernel's default will be used.
	ReceiveChecksumOffload systemdconf.Value

	// Takes a boolean. If set to true, the hardware offload for checksumming of egress network packets is enabled. When unset,
	// the kernel's default will be used.
	TransmitChecksumOffload systemdconf.Value

	// Takes a boolean. If set to true, the TCP Segmentation Offload (TSO) is enabled. When unset, the kernel's default will be
	// used.
	TCPSegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, the TCP6 Segmentation Offload (tx-tcp6-segmentation) is enabled. When unset, the kernel's
	// default will be used.
	TCP6SegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, the Generic Segmentation Offload (GSO) is enabled. When unset, the kernel's default will
	// be used.
	GenericSegmentationOffload systemdconf.Value

	// Takes a boolean. If set to true, the Generic Receive Offload (GRO) is enabled. When unset, the kernel's default will be used.
	GenericReceiveOffload systemdconf.Value

	// Takes a boolean. If set to true, the Large Receive Offload (LRO) is enabled. When unset, the kernel's default will be used.
	LargeReceiveOffload systemdconf.Value

	// Sets the number of receive channels (a number between 1 and 4294967295) .
	RxChannels systemdconf.Value

	// Sets the number of transmit channels (a number between 1 and 4294967295).
	TxChannels systemdconf.Value

	// Sets the number of other channels (a number between 1 and 4294967295).
	OtherChannels systemdconf.Value

	// Sets the number of combined set channels (a number between 1 and 4294967295).
	CombinedChannels systemdconf.Value

	// Takes an integer. Specifies the maximum number of pending packets in the NIC receive buffer. When unset, the kernel's default
	// will be used.
	RxBufferSize systemdconf.Value

	// Takes an integer. Specifies the maximum number of pending packets in the NIC mini receive buffer. When unset, the kernel's
	// default will be used.
	RxMiniBufferSize systemdconf.Value

	// Takes an integer. Specifies the maximum number of pending packets in the NIC jumbo receive buffer. When unset, the kernel's
	// default will be used.
	RxJumboBufferSize systemdconf.Value

	// Takes an integer. Specifies the maximum number of pending packets in the NIC transmit buffer. When unset, the kernel's
	// default will be used.
	TxBufferSize systemdconf.Value

	// Takes a boolean. When set, enables the receive flow control, also known as the ethernet receive PAUSE message (generate
	// and send ethernet PAUSE frames). When unset, the kernel's default will be used.
	RxFlowControl systemdconf.Value

	// Takes a boolean. When set, enables the transmit flow control, also known as the ethernet transmit PAUSE message (respond
	// to received ethernet PAUSE frames). When unset, the kernel's default will be used.
	TxFlowControl systemdconf.Value

	// Takes a boolean. When set, the auto negotiation enables the interface to exchange state advertisements with the connected
	// peer so that the two devices can agree on the ethernet PAUSE configuration. When unset, the kernel's default will be used.
	AutoNegotiationFlowControl systemdconf.Value
}
