// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// NetdevFile represents systemd.netdev — Virtual Network Device configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html for details).
type NetdevFile struct {
	systemdconf.File

	Match                     NetdevMatchSection                     // a virtual network device
	NetDev                    NetdevNetDevSection                    // [NetDev] section
	Bridge                    NetdevBridgeSection                    // [Bridge] section, which only applies to netdevs of kind "bridge"
	VLAN                      NetdevVLANSection                      // [VLAN] section, which only applies to netdevs of kind "vlan"
	MACVLAN                   NetdevMACVLANSection                   // [MACVLAN] section, which only applies to netdevs of kind "macvlan"
	MACVTAP                   NetdevMACVTAPSection                   // [MACVTAP] section, which only applies to netdevs of kind "macvtap"
	IPVLAN                    NetdevIPVLANSection                    // [IPVLAN] section, which only applies to netdevs of kind "ipvlan"
	IPVTAP                    NetdevIPVTAPSection                    // [IPVTAP] section, which only applies to netdevs of kind "ipvtap"
	VXLAN                     NetdevVXLANSection                     // [VXLAN] section, which only applies to netdevs of kind "vxlan"
	GENEVE                    NetdevGENEVESection                    // [GENEVE] section, which only applies to netdevs of kind "geneve"
	BareUDP                   NetdevBareUDPSection                   // [BareUDP] section, which only applies to netdevs of kind "bareudp"
	L2TP                      NetdevL2TPSection                      // [L2TP] section, which only applies to netdevs of kind "l2tp"
	L2TPSession               NetdevL2TPSessionSection               // [L2TPSession] section, which only applies to netdevs of kind "l2tp"
	MACsec                    NetdevMACsecSection                    // [MACsec] section, which only applies to network devices of kind "macsec"
	MACsecReceiveChannel      NetdevMACsecReceiveChannelSection      // [MACsecReceiveChannel] section, which only applies to network devices of kind "macsec"
	MACsecTransmitAssociation NetdevMACsecTransmitAssociationSection // [MACsecTransmitAssociation] section, which only applies to network devices of kind "macsec"
	MACsecReceiveAssociation  NetdevMACsecReceiveAssociationSection  // [MACsecReceiveAssociation] section, which only applies to network devices of kind "macsec"
	Tunnel                    NetdevTunnelSection                    // [Tunnel] section, which only applies to netdevs of kind "ipip", "sit", "gre", "gretap", "ip6gre", "ip6gretap", "vti", "vti6", "ip6tnl", and "erspan"
	FooOverUDP                NetdevFooOverUDPSection                // [FooOverUDP] section, which only applies to netdevs of kind "fou"
	Peer                      NetdevPeerSection                      // [Peer] section, which only applies to netdevs of kind "veth"
	VXCAN                     NetdevVXCANSection                     // [VXCAN] section, which only applies to netdevs of kind "vxcan"
	Tun                       NetdevTunSection                       // [Tun] section, which only applies to netdevs of kind "tun"
	Tap                       NetdevTapSection                       // [Tap] section, which only applies to netdevs of kind "tap"
	WireGuard                 NetdevWireGuardSection                 // [WireGuard] section
	WireGuardPeer             []NetdevWireGuardPeerSection           // [WireGuardPeer] section
	Bond                      NetdevBondSection                      // [Bond] section
	Xfrm                      NetdevXfrmSection                      // [Xfrm] section
	VRF                       NetdevVRFSection                       // [VRF] section, which only applies to netdevs of kind "vrf"
	BatmanAdvanced            NetdevBatmanAdvancedSection            // [BatmanAdvanced] section, which only applies to netdevs of kind "batadv"
	IPoIB                     NetdevIPoIBSection                     // [IPoIB] section, which only applies to netdevs of kind "ipoib"
	WLAN                      NetdevWLANSection                      // [WLAN] section, which only applies to WLAN interfaces
}

// NetdevMatchSection represents a virtual network device
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMatch%5D%20Section%20Options for details).
type NetdevMatchSection struct {
	systemdconf.Section

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

// NetdevNetDevSection represents [NetDev] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BNetDev%5D%20Section%20Options for details).
type NetdevNetDevSection struct {
	systemdconf.Section

	// A free-form description of the netdev.
	//
	// Added in version 215.
	Description systemdconf.Value

	// The interface name used when creating the netdev. This setting is compulsory.
	//
	// Added in version 211.
	Name systemdconf.Value

	// The netdev kind. This setting is compulsory. See the "Supported netdev kinds" section for the valid keys.
	//
	// Added in version 211.
	Kind systemdconf.Value

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G are supported and are understood
	// to the base of 1024. For "tun" or "tap" devices, MTUBytes= setting is not currently supported in [NetDev] section. Please
	// specify it in [Link] section of corresponding systemd.network files.
	//
	// Added in version 215.
	MTUBytes systemdconf.Value

	// Specifies the MAC address to use for the device, or takes the special value "none". When "none", systemd-networkd does
	// not request the MAC address for the device, and the kernel will assign a random MAC address. For "tun", "tap", or "l2tp" devices,
	// the MACAddress= setting in the [NetDev] section is not supported and will be ignored. Please specify it in the [Link] section
	// of the corresponding systemd.network file. If this option is not set, "vlan" device inherits the MAC address of the master
	// interface. For other kind of netdevs, if this option is not set, then the MAC address is generated based on the interface
	// name and the machine-id.
	//
	// Note, even if "none" is specified, systemd-udevd will assign the persistent MAC address for the device, as 99-default.link
	// has MACAddressPolicy=persistent. So, it is also necessary to create a custom .link file for the device, if the MAC address
	// assignment is not desired.
	//
	// Added in version 215.
	MACAddress systemdconf.Value
}

// NetdevBridgeSection represents [Bridge] section, which only applies to netdevs of kind "bridge"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BBridge%5D%20Section%20Options for details).
type NetdevBridgeSection struct {
	systemdconf.Section

	// HelloTimeSec specifies the number of seconds between two hello packets sent out by the root bridge and the designated bridges.
	// Hello packets are used to communicate information about the topology throughout the entire bridged local area network.
	//
	// Added in version 227.
	HelloTimeSec systemdconf.Value

	// MaxAgeSec specifies the number of seconds of maximum message age. If the last seen (received) hello packet is more than
	// this number of seconds old, the bridge in question will start the takeover procedure in attempt to become the Root Bridge
	// itself.
	//
	// Added in version 227.
	MaxAgeSec systemdconf.Value

	// ForwardDelaySec specifies the number of seconds spent in each of the Listening and Learning states before the Forwarding
	// state is entered.
	//
	// Added in version 227.
	ForwardDelaySec systemdconf.Value

	// This specifies the number of seconds a MAC Address will be kept in the forwarding database after having a packet received
	// from this MAC Address.
	//
	// Added in version 232.
	AgeingTimeSec systemdconf.Value

	// The priority of the bridge. An integer between 0 and 65535. A lower value means higher priority. The bridge having the lowest
	// priority will be elected as root bridge.
	//
	// Added in version 232.
	Priority systemdconf.Value

	// A 16-bit bitmask represented as an integer which allows forwarding of link local frames with 802.1D reserved addresses
	// (01:80:C2:00:00:0X). A logical AND is performed between the specified bitmask and the exponentiation of 2^X, the lower
	// nibble of the last octet of the MAC address. For example, a value of 8 would allow forwarding of frames addressed to 01:80:C2:00:00:03
	// (802.1X PAE).
	//
	// Added in version 235.
	GroupForwardMask systemdconf.Value

	// This specifies the default port VLAN ID of a newly attached bridge port. Set this to an integer in the range 1…4094 or "none"
	// to disable the PVID.
	//
	// Added in version 232.
	DefaultPVID systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_QUERIER option in the kernel. If enabled, the kernel will send
	// general ICMP queries from a zero source address. This feature should allow faster convergence on startup, but it causes
	// some multicast-aware switches to misbehave and disrupt forwarding of multicast packets. When unset, the kernel's default
	// will be used.
	//
	// Added in version 230.
	MulticastQuerier systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_SNOOPING option in the kernel. If enabled, IGMP snooping monitors
	// the Internet Group Management Protocol (IGMP) traffic between hosts and multicast routers. When unset, the kernel's
	// default will be used.
	//
	// Added in version 230.
	MulticastSnooping systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_VLAN_FILTERING option in the kernel. If enabled, the bridge will be
	// started in VLAN-filtering mode. When unset, the kernel's default will be used.
	//
	// Added in version 231.
	VLANFiltering systemdconf.Value

	// Allows setting the protocol used for VLAN filtering. Takes 802.1q or, 802.1ad, and defaults to unset and kernel's default
	// is used.
	//
	// Added in version 246.
	VLANProtocol systemdconf.Value

	// Takes a boolean. This enables the bridge's Spanning Tree Protocol (STP). When unset, the kernel's default will be used.
	//
	// Added in version 232.
	STP systemdconf.Value

	// Allows changing bridge's multicast Internet Group Management Protocol (IGMP) version. Takes an integer 2 or 3. When unset,
	// the kernel's default will be used.
	//
	// Added in version 243.
	MulticastIGMPVersion systemdconf.Value

	// Specifies the maximum number of learned Ethernet addresses for the bridge. When the limit is reached, no more addresses
	// are learned. When unset, the kernel's default will be used. 0 disables the limit.
	//
	// Added in version 257.
	FDBMaxLearned systemdconf.Value
}

// NetdevVLANSection represents [VLAN] section, which only applies to netdevs of kind "vlan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BVLAN%5D%20Section%20Options for details).
type NetdevVLANSection struct {
	systemdconf.Section

	// The VLAN ID to use. An integer in the range 0…4094. This setting is compulsory.
	//
	// Added in version 211.
	Id systemdconf.Value

	// Allows setting the protocol used for the VLAN interface. Takes "802.1q" or, "802.1ad", and defaults to unset and kernel's
	// default is used.
	//
	// Added in version 248.
	Protocol systemdconf.Value

	// Takes a boolean. The Generic VLAN Registration Protocol (GVRP) is a protocol that allows automatic learning of VLANs on
	// a network. When unset, the kernel's default will be used.
	//
	// Added in version 234.
	GVRP systemdconf.Value

	// Takes a boolean. Multiple VLAN Registration Protocol (MVRP) formerly known as GARP VLAN Registration Protocol (GVRP)
	// is a standards-based Layer 2 network protocol, for automatic configuration of VLAN information on switches. It was defined
	// in the 802.1ak amendment to 802.1Q-2005. When unset, the kernel's default will be used.
	//
	// Added in version 234.
	MVRP systemdconf.Value

	// Takes a boolean. The VLAN loose binding mode, in which only the operational state is passed from the parent to the associated
	// VLANs, but the VLAN device state is not changed. When unset, the kernel's default will be used.
	//
	// Added in version 234.
	LooseBinding systemdconf.Value

	// Takes a boolean. When enabled, the VLAN reorder header is used and VLAN interfaces behave like physical interfaces. When
	// unset, the kernel's default will be used.
	//
	// Added in version 234.
	ReorderHeader systemdconf.Value

	// Defines a mapping of Linux internal packet priority (SO_PRIORITY) to VLAN header PCP field for outgoing and incoming frames,
	// respectively. Takes a whitespace-separated list of integer pairs, where each integer must be in the range 1…4294967294,
	// in the format "from"-"to", e.g., "21-7 45-5". Note that "from" must be greater than or equal to "to". When unset, the kernel's
	// default will be used.
	//
	// Added in version 248.
	EgressQOSMaps systemdconf.Value

	// Defines a mapping of Linux internal packet priority (SO_PRIORITY) to VLAN header PCP field for outgoing and incoming frames,
	// respectively. Takes a whitespace-separated list of integer pairs, where each integer must be in the range 1…4294967294,
	// in the format "from"-"to", e.g., "21-7 45-5". Note that "from" must be greater than or equal to "to". When unset, the kernel's
	// default will be used.
	//
	// Added in version 248.
	IngressQOSMaps systemdconf.Value
}

// NetdevMACVLANSection represents [MACVLAN] section, which only applies to netdevs of kind "macvlan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACVLAN%5D%20Section%20Options for details).
type NetdevMACVLANSection struct {
	systemdconf.Section

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", "passthru", and "source".
	//
	// Added in version 211.
	Mode systemdconf.Value

	// A whitespace-separated list of remote hardware addresses allowed on the MACVLAN. This option only has an effect in source
	// mode. Use full colon-, hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists
	// are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	// Defaults to unset.
	//
	// Added in version 246.
	SourceMACAddress systemdconf.Value

	// Specifies the length of the receive queue for broadcast/multicast packets. An unsigned integer in the range 0…4294967294.
	// Defaults to unset.
	//
	// Added in version 248.
	BroadcastMulticastQueueLength systemdconf.Value

	// Controls the threshold for broadcast queueing of the macvlan device. Takes the special value "no", or an integer in the
	// range 0…2147483647. When "no" is specified, the broadcast queueing is disabled altogether. When an integer is specified,
	// a multicast address will be queued as broadcast if the number of devices using the macvlan is greater than the given value.
	// Defaults to unset, and the kernel default will be used.
	//
	// Added in version 256.
	BroadcastQueueThreshold systemdconf.Value
}

// NetdevMACVTAPSection represents [MACVTAP] section, which only applies to netdevs of kind "macvtap"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACVLAN%5D%20Section%20Options for details).
type NetdevMACVTAPSection struct {
	systemdconf.Section

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", "passthru", and "source".
	//
	// Added in version 211.
	Mode systemdconf.Value

	// A whitespace-separated list of remote hardware addresses allowed on the MACVLAN. This option only has an effect in source
	// mode. Use full colon-, hyphen- or dot-delimited hexadecimal. This option may appear more than once, in which case the lists
	// are merged. If the empty string is assigned to this option, the list of hardware addresses defined prior to this is reset.
	// Defaults to unset.
	//
	// Added in version 246.
	SourceMACAddress systemdconf.Value

	// Specifies the length of the receive queue for broadcast/multicast packets. An unsigned integer in the range 0…4294967294.
	// Defaults to unset.
	//
	// Added in version 248.
	BroadcastMulticastQueueLength systemdconf.Value

	// Controls the threshold for broadcast queueing of the macvlan device. Takes the special value "no", or an integer in the
	// range 0…2147483647. When "no" is specified, the broadcast queueing is disabled altogether. When an integer is specified,
	// a multicast address will be queued as broadcast if the number of devices using the macvlan is greater than the given value.
	// Defaults to unset, and the kernel default will be used.
	//
	// Added in version 256.
	BroadcastQueueThreshold systemdconf.Value
}

// NetdevIPVLANSection represents [IPVLAN] section, which only applies to netdevs of kind "ipvlan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BIPVLAN%5D%20Section%20Options for details).
type NetdevIPVLANSection struct {
	systemdconf.Section

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	//
	// Added in version 219.
	Mode systemdconf.Value

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	//
	// Added in version 237.
	Flags systemdconf.Value
}

// NetdevIPVTAPSection represents [IPVTAP] section, which only applies to netdevs of kind "ipvtap"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BIPVLAN%5D%20Section%20Options for details).
type NetdevIPVTAPSection struct {
	systemdconf.Section

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	//
	// Added in version 219.
	Mode systemdconf.Value

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	//
	// Added in version 237.
	Flags systemdconf.Value
}

// NetdevVXLANSection represents [VXLAN] section, which only applies to netdevs of kind "vxlan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BVXLAN%5D%20Section%20Options for details).
type NetdevVXLANSection struct {
	systemdconf.Section

	// The VXLAN Network Identifier (or VXLAN Segment ID). Takes a number in the range 1…16777215.
	//
	// Added in version 243.
	VNI systemdconf.Value

	// Configures destination IP address.
	//
	// Added in version 233.
	Remote systemdconf.Value

	// Configures local IP address. It must be an address on the underlying interface of the VXLAN interface, or one of the special
	// values "ipv4_link_local", "ipv6_link_local", "dhcp4", "dhcp6", and "slaac". If one of the special values is specified,
	// an address which matches the corresponding type on the underlying interface will be used. Defaults to unset.
	//
	// Added in version 233.
	Local systemdconf.Value

	// Configures VXLAN multicast group IP address. All members of a VXLAN must use the same multicast group address.
	//
	// Added in version 243.
	Group systemdconf.Value

	// The Type Of Service byte value for a vxlan interface.
	//
	// Added in version 215.
	TOS systemdconf.Value

	// A fixed Time To Live N on Virtual eXtensible Local Area Network packets. Takes "inherit" or a number in the range 0…255. 0
	// is a special value meaning inherit the inner protocol's TTL value. "inherit" means that it will inherit the outer protocol's
	// TTL value.
	//
	// Added in version 215.
	TTL systemdconf.Value

	// Takes a boolean. When true, enables dynamic MAC learning to discover remote MAC addresses.
	//
	// Added in version 215.
	MacLearning systemdconf.Value

	// The lifetime of Forwarding Database entry learnt by the kernel, in seconds.
	//
	// Added in version 218.
	FDBAgeingSec systemdconf.Value

	// Configures maximum number of FDB entries.
	//
	// Added in version 228.
	MaximumFDBEntries systemdconf.Value

	// Takes a boolean. When true, bridge-connected VXLAN tunnel endpoint answers ARP requests from the local bridge on behalf
	// of remote  Distributed Overlay Virtual Ethernet (DOVE) clients. Defaults to false.
	//
	// Added in version 233.
	ReduceARPProxy systemdconf.Value

	// Takes a boolean. When true, enables netlink LLADDR miss notifications.
	//
	// Added in version 218.
	L2MissNotification systemdconf.Value

	// Takes a boolean. When true, enables netlink IP address miss notifications.
	//
	// Added in version 218.
	L3MissNotification systemdconf.Value

	// Takes a boolean. When true, route short circuiting is turned on.
	//
	// Added in version 218.
	RouteShortCircuit systemdconf.Value

	// Takes a boolean. When true, transmitting UDP checksums when doing VXLAN/IPv4 is turned on.
	//
	// Added in version 220.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, sending zero checksums in VXLAN/IPv6 is turned on.
	//
	// Added in version 220.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, receiving zero checksums in VXLAN/IPv6 is turned on.
	//
	// Added in version 220.
	UDP6ZeroChecksumRx systemdconf.Value

	// Takes a boolean. When true, remote transmit checksum offload of VXLAN is turned on.
	//
	// Added in version 232.
	RemoteChecksumTx systemdconf.Value

	// Takes a boolean. When true, remote receive checksum offload in VXLAN is turned on.
	//
	// Added in version 232.
	RemoteChecksumRx systemdconf.Value

	// Takes a boolean. When true, it enables Group Policy VXLAN extension security label mechanism across network peers based
	// on VXLAN. For details about the Group Policy VXLAN, see the  VXLAN Group Policy  document. Defaults to false.
	//
	// Added in version 224.
	GroupPolicyExtension systemdconf.Value

	// Takes a boolean. When true, Generic Protocol Extension extends the existing VXLAN protocol to provide protocol typing,
	// OAM, and versioning capabilities. For details about the VXLAN GPE Header, see the  Generic Protocol Extension for VXLAN
	//
	//	document. If destination port is not specified and Generic Protocol Extension is set, the default port of 4790 is used.
	//
	// Defaults to false.
	//
	// Added in version 243.
	GenericProtocolExtension systemdconf.Value

	// Configures the default destination UDP port. If the destination port is not specified, the Linux kernel default will be
	// used. Set to 4789 to get the IANA assigned value.
	//
	// Added in version 229.
	DestinationPort systemdconf.Value

	// Configures the source port range for the VXLAN. The kernel assigns the source UDP port based on the flow to help the receiver
	// to do load balancing. When this option is not set, the normal range of local UDP ports is used.
	//
	// Added in version 229.
	PortRange systemdconf.Value

	// Specifies the flow label to use in outgoing packets. The valid range is 0-1048575.
	//
	// Added in version 234.
	FlowLabel systemdconf.Value

	// Allows setting the IPv4 Do not Fragment (DF) bit in outgoing packets, or to inherit its value from the IPv4 inner header.
	// Takes a boolean value, or "inherit". Set to "inherit" if the encapsulated protocol is IPv6. When unset, the kernel's default
	// will be used.
	//
	// Added in version 243.
	IPDoNotFragment systemdconf.Value

	// Takes a boolean. When true, the vxlan interface is created without any underlying network interface. Defaults to false,
	// which means that a .network file that requests this VXLAN interface using VXLAN= is required for the VXLAN to be created.
	//
	// Added in version 247.
	Independent systemdconf.Value
}

// NetdevGENEVESection represents [GENEVE] section, which only applies to netdevs of kind "geneve"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BGENEVE%5D%20Section%20Options for details).
type NetdevGENEVESection struct {
	systemdconf.Section

	// Specifies the Virtual Network Identifier (VNI) to use, a number between 0 and 16777215. This field is mandatory.
	//
	// Added in version 234.
	Id systemdconf.Value

	// Specifies the unicast destination IP address to use in outgoing packets.
	//
	// Added in version 234.
	Remote systemdconf.Value

	// Specifies the TOS value to use in outgoing packets. Takes a number between 1 and 255.
	//
	// Added in version 234.
	TOS systemdconf.Value

	// Accepts the same values as in the [VXLAN] section, except that when unset or set to 0, the kernel's default will be used, meaning
	// that packet TTL will be set from /proc/sys/net/ipv4/ip_default_ttl.
	//
	// Added in version 234.
	TTL systemdconf.Value

	// Takes a boolean. When true, specifies that UDP checksum is calculated for transmitted packets over IPv4.
	//
	// Added in version 234.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	//
	// Added in version 234.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	//
	// Added in version 234.
	UDP6ZeroChecksumRx systemdconf.Value

	// Specifies destination port. Defaults to 6081. If not set or assigned the empty string, the default port of 6081 is used.
	//
	// Added in version 234.
	DestinationPort systemdconf.Value

	// Specifies the flow label to use in outgoing packets.
	//
	// Added in version 234.
	FlowLabel systemdconf.Value

	// Accepts the same key as in [VXLAN] section.
	//
	// Added in version 243.
	IPDoNotFragment systemdconf.Value

	// Takes a boolean. When true, inner Layer 3 protocol is set as Protocol Type in the GENEVE header instead of Ethernet. Defaults
	// to false.
	//
	// Added in version 254.
	InheritInnerProtocol systemdconf.Value
}

// NetdevBareUDPSection represents [BareUDP] section, which only applies to netdevs of kind "bareudp"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BBareUDP%5D%20Section%20Options for details).
type NetdevBareUDPSection struct {
	systemdconf.Section

	// Specifies the destination UDP port (in range 1…65535). This is mandatory.
	//
	// Added in version 247.
	DestinationPort systemdconf.Value

	// Specifies the L3 protocol. Takes one of "ipv4", "ipv6", "mpls-uc" or "mpls-mc". This is mandatory.
	//
	// Added in version 247.
	EtherType systemdconf.Value

	// Specifies the lowest value of the UDP tunnel source port (in range 1…65535). Defaults to unset.
	//
	// Added in version 257.
	MinSourcePort systemdconf.Value
}

// NetdevL2TPSection represents [L2TP] section, which only applies to netdevs of kind "l2tp"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BL2TP%5D%20Section%20Options for details).
type NetdevL2TPSection struct {
	systemdconf.Section

	// Specifies the tunnel identifier. Takes an number in the range 1…4294967295. The value used must match the "PeerTunnelId="
	// value being used at the peer. This setting is compulsory.
	//
	// Added in version 242.
	TunnelId systemdconf.Value

	// Specifies the peer tunnel id. Takes a number in the range 1…4294967295. The value used must match the "TunnelId=" value
	// being used at the peer. This setting is compulsory.
	//
	// Added in version 242.
	PeerTunnelId systemdconf.Value

	// Specifies the IP address of the remote peer. This setting is compulsory.
	//
	// Added in version 242.
	Remote systemdconf.Value

	// Specifies the IP address of a local interface. Takes an IP address, or the special values "auto", "static", or "dynamic".
	// Optionally a name of a local interface can be specified after "@", e.g. "192.168.0.1@eth0" or "auto@eth0". When an address
	// is specified, then a local or specified interface must have the address, and the remote address must be accessible through
	// the local address. If "auto", then one of the addresses on a local or specified interface which is accessible to the remote
	// address will be used. Similarly, if "static" or "dynamic" is set, then one of the static or dynamic addresses will be used.
	// Defaults to "auto".
	//
	// Added in version 242.
	Local systemdconf.Value

	// Specifies the encapsulation type of the tunnel. Takes one of "udp" or "ip".
	//
	// Added in version 242.
	EncapsulationType systemdconf.Value

	// Specifies the UDP source port to be used for the tunnel. When UDP encapsulation is selected it is mandatory. Ignored when
	// IP encapsulation is selected.
	//
	// Added in version 242.
	UDPSourcePort systemdconf.Value

	// Specifies destination port. When UDP encapsulation is selected it is mandatory. Ignored when IP encapsulation is selected.
	//
	// Added in version 245.
	UDPDestinationPort systemdconf.Value

	// Takes a boolean. When true, specifies that UDP checksum is calculated for transmitted packets over IPv4.
	//
	// Added in version 242.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	//
	// Added in version 242.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	//
	// Added in version 242.
	UDP6ZeroChecksumRx systemdconf.Value
}

// NetdevL2TPSessionSection represents [L2TPSession] section, which only applies to netdevs of kind "l2tp"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BL2TPSession%5D%20Section%20Options for details).
type NetdevL2TPSessionSection struct {
	systemdconf.Section

	// Specifies the name of the session. This setting is compulsory.
	//
	// Added in version 242.
	Name systemdconf.Value

	// Specifies the session identifier. Takes an number in the range 1…4294967295. The value used must match the "SessionId="
	// value being used at the peer. This setting is compulsory.
	//
	// Added in version 242.
	SessionId systemdconf.Value

	// Specifies the peer session identifier. Takes an number in the range 1…4294967295. The value used must match the "PeerSessionId="
	// value being used at the peer. This setting is compulsory.
	//
	// Added in version 242.
	PeerSessionId systemdconf.Value

	// Specifies layer2specific header type of the session. One of "none" or "default". Defaults to "default".
	//
	// Added in version 242.
	Layer2SpecificHeader systemdconf.Value
}

// NetdevMACsecSection represents [MACsec] section, which only applies to network devices of kind "macsec"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACsec%5D%20Section%20Options for details).
type NetdevMACsecSection struct {
	systemdconf.Section

	// Specifies the port to be used for the MACsec transmit channel. The port is used to make secure channel identifier (SCI).
	// Takes a value between 1 and 65535. Defaults to unset.
	//
	// Added in version 243.
	Port systemdconf.Value

	// Takes a boolean. When true, enable encryption. Defaults to unset.
	//
	// Added in version 243.
	Encrypt systemdconf.Value
}

// NetdevMACsecReceiveChannelSection represents [MACsecReceiveChannel] section, which only applies to network devices of kind "macsec"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACsecReceiveChannel%5D%20Section%20Options for details).
type NetdevMACsecReceiveChannelSection struct {
	systemdconf.Section

	// Specifies the port to be used for the MACsec receive channel. The port is used to make secure channel identifier (SCI). Takes
	// a value between 1 and 65535. This option is compulsory, and is not set by default.
	//
	// Added in version 243.
	Port systemdconf.Value

	// Specifies the MAC address to be used for the MACsec receive channel. The MAC address used to make secure channel identifier
	// (SCI). This setting is compulsory, and is not set by default.
	//
	// Added in version 243.
	MACAddress systemdconf.Value
}

// NetdevMACsecTransmitAssociationSection represents [MACsecTransmitAssociation] section, which only applies to network devices of kind "macsec"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACsecTransmitAssociation%5D%20Section%20Options for details).
type NetdevMACsecTransmitAssociationSection struct {
	systemdconf.Section

	// Specifies the packet number to be used for replay protection and the construction of the initialization vector (along
	// with the secure channel identifier [SCI]). Takes a value between 1-4,294,967,295. Defaults to unset.
	//
	// Added in version 243.
	PacketNumber systemdconf.Value

	// Specifies the identification for the key. Takes a number between 0-255. This option is compulsory, and is not set by default.
	//
	// Added in version 243.
	KeyId systemdconf.Value

	// Specifies the encryption key used in the transmission channel. The same key must be configured on the peer’s matching receive
	// channel. This setting is compulsory, and is not set by default. Takes a 128-bit key encoded in a hexadecimal string, for
	// example "dffafc8d7b9a43d5b9a3dfbbf6a30c16".
	//
	// Added in version 243.
	Key systemdconf.Value

	// Takes an absolute path to a file which contains a 128-bit key encoded in a hexadecimal string, which will be used in the transmission
	// channel. When this option is specified, Key= is ignored. Note that the file must be readable by the user "systemd-network",
	// so it should be, e.g., owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket
	// in the file system a connection is made to it and the key read from it.
	//
	// Added in version 243.
	KeyFile systemdconf.Value

	// Takes a boolean. If enabled, then the security association is activated. Defaults to unset.
	//
	// Added in version 243.
	Activate systemdconf.Value

	// Takes a boolean. If enabled, then the security association is used for encoding. Only one [MACsecTransmitAssociation]
	// section can enable this option. When enabled, Activate=yes is implied. Defaults to unset.
	//
	// Added in version 243.
	UseForEncoding systemdconf.Value
}

// NetdevMACsecReceiveAssociationSection represents [MACsecReceiveAssociation] section, which only applies to network devices of kind "macsec"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BMACsecReceiveAssociation%5D%20Section%20Options for details).
type NetdevMACsecReceiveAssociationSection struct {
	systemdconf.Section

	// Accepts the same key as in [MACsecReceiveChannel] section.
	//
	// Added in version 243.
	Port systemdconf.Value

	// Accepts the same key as in [MACsecReceiveChannel] section.
	//
	// Added in version 243.
	MACAddress systemdconf.Value

	// Accepts the same key as in [MACsecTransmitAssociation] section.
	//
	// Added in version 243.
	PacketNumber systemdconf.Value

	// Accepts the same key as in [MACsecTransmitAssociation] section.
	//
	// Added in version 243.
	KeyId systemdconf.Value

	// Accepts the same key as in [MACsecTransmitAssociation] section.
	//
	// Added in version 243.
	Key systemdconf.Value

	// Accepts the same key as in [MACsecTransmitAssociation] section.
	//
	// Added in version 243.
	KeyFile systemdconf.Value

	// Accepts the same key as in [MACsecTransmitAssociation] section.
	//
	// Added in version 243.
	Activate systemdconf.Value
}

// NetdevTunnelSection represents [Tunnel] section, which only applies to netdevs of kind "ipip", "sit", "gre", "gretap", "ip6gre", "ip6gretap", "vti", "vti6", "ip6tnl", and "erspan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BTunnel%5D%20Section%20Options for details).
type NetdevTunnelSection struct {
	systemdconf.Section

	// Takes a boolean value. When true, then the tunnel is externally controlled, which is also known as collect metadata mode,
	// and most settings below like Local= or Remote= are ignored. This implies Independent=. Defaults to false.
	//
	// Added in version 251.
	External systemdconf.Value

	// A static local address for tunneled packets. It must be an address on another interface of this host, or one of the special
	// values "any", "ipv4_link_local", "ipv6_link_local", "dhcp4", "dhcp6", and "slaac". If one of the special values except
	// for "any" is specified, an address which matches the corresponding type on the underlying interface will be used. Defaults
	// to "any".
	//
	// Added in version 215.
	Local systemdconf.Value

	// The remote endpoint of the tunnel. Takes an IP address or the special value "any".
	//
	// Added in version 215.
	Remote systemdconf.Value

	// The Type Of Service byte value for a tunnel interface. For details about the TOS, see the  Type of Service in the Internet Protocol
	// Suite  document.
	//
	// Added in version 215.
	TOS systemdconf.Value

	// A fixed Time To Live N on tunneled packets. N is a number in the range 1…255. 0 is a special value meaning that packets inherit
	// the TTL value. The default value for IPv4 tunnels is 0 (inherit). The default value for IPv6 tunnels is 64.
	//
	// Added in version 215.
	TTL systemdconf.Value

	// Takes a boolean. When true, enables Path MTU Discovery on the tunnel. When IgnoreDontFragment= is enabled, defaults to
	// false. Otherwise, defaults to true.
	//
	// Added in version 215.
	DiscoverPathMTU systemdconf.Value

	// Takes a boolean. When true, enables IPv4 Don't Fragment (DF) suppression on the tunnel. Defaults to false. Note that if
	// IgnoreDontFragment= is set to true, DiscoverPathMTU= cannot be set to true. Only applicable to GRE, GRETAP, and ERSPAN
	// tunnels.
	//
	// Added in version 254.
	IgnoreDontFragment systemdconf.Value

	// Configures the 20-bit flow label (see  RFC 6437) field in the IPv6 header (see  RFC 2460), which is used by a node to label packets
	// of a flow. It is only used for IPv6 tunnels. A flow label of zero is used to indicate packets that have not been labeled. It can
	// be configured to a value in the range 0…0xFFFFF, or be set to "inherit", in which case the original flowlabel is used.
	//
	// Added in version 223.
	IPv6FlowLabel systemdconf.Value

	// Takes a boolean. When true, the Differentiated Service Code Point (DSCP) field will be copied to the inner header from outer
	// header during the decapsulation of an IPv6 tunnel packet. DSCP is a field in an IP packet that enables different levels of
	// service to be assigned to network traffic. Defaults to "no".
	//
	// Added in version 223.
	CopyDSCP systemdconf.Value

	// The Tunnel Encapsulation Limit option specifies how many additional levels of encapsulation are permitted to be prepended
	// to the packet. For example, a Tunnel Encapsulation Limit option containing a limit value of zero means that a packet carrying
	// that option may not enter another tunnel before exiting the current tunnel. (see  RFC 2473). The valid range is 0…255 and
	// "none". Defaults to 4.
	//
	// Added in version 226.
	EncapsulationLimit systemdconf.Value

	// The Key= parameter specifies the same key to use in both directions (InputKey= and OutputKey=). The Key= is either a number
	// or an IPv4 address-like dotted quad. It is used as mark-configured SAD/SPD entry as part of the lookup key (both in data and
	// control path) in IP XFRM (framework used to implement IPsec protocol). See  ip-xfrm — transform configuration for details.
	// It is only used for VTI/VTI6, GRE, GRETAP, and ERSPAN tunnels.
	//
	// Added in version 231.
	Key systemdconf.Value

	// The InputKey= parameter specifies the key to use for input. The format is same as Key=. It is only used for VTI/VTI6, GRE,
	// GRETAP, and ERSPAN tunnels.
	//
	// Added in version 231.
	InputKey systemdconf.Value

	// The OutputKey= parameter specifies the key to use for output. The format is same as Key=. It is only used for VTI/VTI6, GRE,
	// GRETAP, and ERSPAN tunnels.
	//
	// Added in version 231.
	OutputKey systemdconf.Value

	// An "ip6tnl" tunnel can be in one of three modes "ip6ip6" for IPv6 over IPv6, "ipip6" for IPv4 over IPv6 or "any" for either.
	//
	// Added in version 219.
	Mode systemdconf.Value

	// Takes a boolean. When false (the default), the tunnel is always created over some network device, and a .network file that
	// requests this tunnel using Tunnel= is required for the tunnel to be created. When true, the tunnel is created independently
	// of any network as "tunnel@NONE".
	//
	// Added in version 235.
	Independent systemdconf.Value

	// Takes a boolean. If set to "yes", the loopback interface "lo" is used as the underlying device of the tunnel interface. Defaults
	// to "no".
	//
	// Added in version 243.
	AssignToLoopback systemdconf.Value

	// Takes a boolean. When true allows tunnel traffic on ip6tnl devices where the remote endpoint is a local host address. When
	// unset, the kernel's default will be used.
	//
	// Added in version 237.
	AllowLocalRemote systemdconf.Value

	// Takes a boolean. Specifies whether FooOverUDP= tunnel is to be configured. Defaults to false. This takes effects only
	// for IPIP, SIT, GRE, and GRETAP tunnels. For more detail information see Foo over UDP
	//
	// Added in version 240.
	FooOverUDP systemdconf.Value

	// This setting specifies the UDP destination port for encapsulation. This field is mandatory when FooOverUDP=yes, and
	// is not set by default.
	//
	// Added in version 240.
	FOUDestinationPort systemdconf.Value

	// This setting specifies the UDP source port for encapsulation. Defaults to 0 — that is, the source port for packets is left
	// to the network stack to decide.
	//
	// Added in version 240.
	FOUSourcePort systemdconf.Value

	// Accepts the same key as in the [FooOverUDP] section.
	//
	// Added in version 240.
	Encapsulation systemdconf.Value

	// Reconfigure the tunnel for IPv6 Rapid Deployment, also known as 6rd. The value is an ISP-specific IPv6 prefix with a non-zero
	// length. Only applicable to SIT tunnels.
	//
	// Added in version 240.
	IPv6RapidDeploymentPrefix systemdconf.Value

	// Takes a boolean. If set, configures the tunnel as Intra-Site Automatic Tunnel Addressing Protocol (ISATAP) tunnel. Only
	// applicable to SIT tunnels. When unset, the kernel's default will be used.
	//
	// Added in version 240.
	ISATAP systemdconf.Value

	// Takes a boolean. If set to yes, then packets are serialized. Only applies for GRE, GRETAP, and ERSPAN tunnels. When unset,
	// the kernel's default will be used.
	//
	// Added in version 240.
	SerializeTunneledPackets systemdconf.Value

	// Specifies the ERSPAN version number. Takes 0 for version 0 (a.k.a. type I), 1 for version 1 (a.k.a. type II), or 2 for version
	// 2 (a.k.a. type III). Defaults to 1.
	//
	// Added in version 252.
	ERSPANVersion systemdconf.Value

	// Specifies the ERSPAN v1 index field for the interface. Takes an integer in the range 0…1048575, which is associated with
	// the ERSPAN traffic's source port and direction. Only used when ERSPANVersion=1. Defaults to 0.
	//
	// Added in version 240.
	ERSPANIndex systemdconf.Value

	// Specifies the ERSPAN v2 mirrored traffic's direction. Takes "ingress" or "egress". Only used when ERSPANVersion=2.
	// Defaults to "ingress".
	//
	// Added in version 252.
	ERSPANDirection systemdconf.Value

	// Specifies an unique identifier of the ERSPAN v2 engine. Takes an integer in the range 0…63. Only used when ERSPANVersion=2.
	// Defaults to 0.
	//
	// Added in version 252.
	ERSPANHardwareId systemdconf.Value
}

// NetdevFooOverUDPSection represents [FooOverUDP] section, which only applies to netdevs of kind "fou"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BFooOverUDP%5D%20Section%20Options for details).
type NetdevFooOverUDPSection struct {
	systemdconf.Section

	// Specifies the encapsulation mechanism used to store networking packets of various protocols inside the UDP packets.
	// Supports the following values: "FooOverUDP" provides the simplest no-frills model of UDP encapsulation, it simply encapsulates
	// packets directly in the UDP payload. "GenericUDPEncapsulation" is a generic and extensible encapsulation, it allows
	// encapsulation of packets for any IP protocol and optional data as part of the encapsulation. For more detailed information
	// see Generic UDP Encapsulation. Defaults to "FooOverUDP".
	//
	// Added in version 240.
	Encapsulation systemdconf.Value

	// Specifies the port number where the encapsulated packets will arrive. Those packets will be removed and manually fed back
	// into the network stack with the encapsulation removed to be sent to the real destination. This option is mandatory.
	//
	// Added in version 240.
	Port systemdconf.Value

	// Specifies the peer port number. Defaults to unset. Note that when peer port is set "Peer=" address is mandatory.
	//
	// Added in version 243.
	PeerPort systemdconf.Value

	// The Protocol= specifies the protocol number of the packets arriving at the UDP port. When Encapsulation=FooOverUDP,
	// this field is mandatory and is not set by default. Takes an IP protocol name such as "gre" or "ipip", or an integer within the
	// range 1…255. When Encapsulation=GenericUDPEncapsulation, this must not be specified.
	//
	// Added in version 240.
	Protocol systemdconf.Value

	// Configures peer IP address. Note that when peer address is set "PeerPort=" is mandatory.
	//
	// Added in version 243.
	Peer systemdconf.Value

	// Configures local IP address.
	//
	// Added in version 243.
	Local systemdconf.Value
}

// NetdevPeerSection represents [Peer] section, which only applies to netdevs of kind "veth"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BPeer%5D%20Section%20Options for details).
type NetdevPeerSection struct {
	systemdconf.Section

	// The interface name used when creating the netdev. This setting is compulsory.
	//
	// Added in version 215.
	Name systemdconf.Value

	// The peer MACAddress, if not set, it is generated in the same way as the MAC address of the main interface.
	//
	// Added in version 215.
	MACAddress systemdconf.Value
}

// NetdevVXCANSection represents [VXCAN] section, which only applies to netdevs of kind "vxcan"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BVXCAN%5D%20Section%20Options for details).
type NetdevVXCANSection struct {
	systemdconf.Section

	// The peer interface name used when creating the netdev. This setting is compulsory.
	//
	// Added in version 236.
	Peer systemdconf.Value
}

// NetdevTunSection represents [Tun] section, which only applies to netdevs of kind "tun"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BTun%5D%20Section%20Options for details).
type NetdevTunSection struct {
	systemdconf.Section

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving.
	// Defaults to "no".
	//
	// Added in version 215.
	MultiQueue systemdconf.Value

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol
	// bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	//
	// Added in version 215.
	PacketInfo systemdconf.Value

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation
	// Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	//
	// Added in version 223.
	VNetHeader systemdconf.Value

	// User to grant access to the /dev/net/tun device.
	//
	// Added in version 215.
	User systemdconf.Value

	// Group to grant access to the /dev/net/tun device.
	//
	// Added in version 215.
	Group systemdconf.Value

	// Takes a boolean. If enabled, to make the interface maintain its carrier status, the file descriptor of the interface is
	// kept open. This may be useful to keep the interface in running state, for example while the backing process is temporarily
	// shutdown. Defaults to "no".
	//
	// Added in version 252.
	KeepCarrier systemdconf.Value
}

// NetdevTapSection represents [Tap] section, which only applies to netdevs of kind "tap"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BTun%5D%20Section%20Options for details).
type NetdevTapSection struct {
	systemdconf.Section

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving.
	// Defaults to "no".
	//
	// Added in version 215.
	MultiQueue systemdconf.Value

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol
	// bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	//
	// Added in version 215.
	PacketInfo systemdconf.Value

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation
	// Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	//
	// Added in version 223.
	VNetHeader systemdconf.Value

	// User to grant access to the /dev/net/tun device.
	//
	// Added in version 215.
	User systemdconf.Value

	// Group to grant access to the /dev/net/tun device.
	//
	// Added in version 215.
	Group systemdconf.Value

	// Takes a boolean. If enabled, to make the interface maintain its carrier status, the file descriptor of the interface is
	// kept open. This may be useful to keep the interface in running state, for example while the backing process is temporarily
	// shutdown. Defaults to "no".
	//
	// Added in version 252.
	KeepCarrier systemdconf.Value
}

// NetdevWireGuardSection represents [WireGuard] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BWireGuard%5D%20Section%20Options for details).
type NetdevWireGuardSection struct {
	systemdconf.Section

	// The Base64 encoded private key for the interface. It can be generated using the wg genkey command (see wg). Specially, if
	// the specified key is prefixed with "@", it is interpreted as the name of the credential from which the actual key shall be
	// read. systemd-networkd.service automatically imports credentials matching "network.wireguard.*". For more details
	// on credentials, refer to systemd.exec. A private key is mandatory to use WireGuard. If not set, the credential "network.wireguard.private.netdev"
	// is used if exists. I.e. for 50-foobar.netdev, "network.wireguard.private.50-foobar" is tried.
	//
	// Note that because this information is secret, it is strongly recommended to use an (encrypted) credential. Alternatively,
	// you may want to set the permissions of the .netdev file to be owned by "root:systemd-network" with a "0640" file mode.
	//
	// Added in version 237.
	PrivateKey systemdconf.Value

	// Takes an absolute path to a file which contains the Base64 encoded private key for the interface. When this option is specified,
	// then PrivateKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g., owned
	// by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system a connection
	// is made to it and the key read from it.
	//
	// Added in version 242.
	PrivateKeyFile systemdconf.Value

	// Sets UDP port for listening. Takes either value between 1 and 65535 or "auto". If "auto" is specified, the port is automatically
	// generated based on interface name. Defaults to "auto".
	//
	// Added in version 237.
	ListenPort systemdconf.Value

	// Sets a firewall mark on outgoing WireGuard packets from this interface. Takes a number between 1 and 4294967295.
	//
	// Added in version 243.
	FirewallMark systemdconf.Value

	// The table identifier for the routes to the addresses specified in the AllowedIPs=. Takes a negative boolean value, one
	// of the predefined names "default", "main", and "local", names defined in RouteTable= in networkd.conf, or a number in
	// the range 1…4294967295. When "off" the routes to the addresses specified in the AllowedIPs= setting will not be configured.
	// Defaults to false. This setting will be ignored when the same setting is specified in the [WireGuardPeer] section.
	//
	// Added in version 250.
	RouteTable systemdconf.Value

	// The priority of the routes to the addresses specified in the AllowedIPs=. Takes an integer in the range 0…4294967295. Defaults
	// to 0 for IPv4 addresses, and 1024 for IPv6 addresses. This setting will be ignored when the same setting is specified in the
	// [WireGuardPeer] section.
	//
	// Added in version 250.
	RouteMetric systemdconf.Value
}

// NetdevWireGuardPeerSection represents [WireGuardPeer] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BWireGuardPeer%5D%20Section%20Options for details).
type NetdevWireGuardPeerSection struct {
	systemdconf.Section

	// Sets a Base64 encoded public key calculated by wg pubkey (see wg) from a private key, and usually transmitted out of band
	// to the author of the configuration file. This option honors the "@" prefix in the same way as the PrivateKey= setting of the
	// [WireGuard] section. This option is mandatory for this section.
	//
	// Added in version 237.
	PublicKey systemdconf.Value

	// Takes an absolute path to a file which contains the Base64 encoded public key for the peer. When this option is specified,
	// then PublicKey= will be ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g.,
	// owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system
	// a connection is made to it and the key read from it.
	//
	// Added in version 257.
	PublicKeyFile systemdconf.Value

	// Optional preshared key for the interface. It can be generated by the wg genpsk command. This option adds an additional layer
	// of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	// This option honors the "@" prefix in the same way as the PrivateKey= setting of the [WireGuard] section.
	//
	// Note that because this information is secret, it is strongly recommended to use an (encrypted) credential. Alternatively,
	// you may want to set the permissions of the .netdev file to be owned by "root:systemd-network" with a "0640" file mode.
	//
	// Added in version 237.
	PresharedKey systemdconf.Value

	// Takes an absolute path to a file which contains the Base64 encoded preshared key for the peer. When this option is specified,
	// then PresharedKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g.,
	// owned by "root:systemd-network" with a "0640" file mode. If the path refers to an AF_UNIX stream socket in the file system
	// a connection is made to it and the key read from it.
	//
	// Added in version 242.
	PresharedKeyFile systemdconf.Value

	// Sets a comma-separated list of IP (v4 or v6) addresses with CIDR masks from which this peer is allowed to send incoming traffic
	// and to which outgoing traffic for this peer is directed. This setting can be specified multiple times. If an empty string
	// is assigned, then the all previous assignments are cleared.
	//
	// The catch-all 0.0.0.0/0 may be specified for matching all IPv4 addresses, and ::/0 may be specified for matching all IPv6
	// addresses.
	//
	// Note that this only affects routing inside the network interface itself, i.e. the packets that pass through the tunnel
	// itself. To cause packets to be sent via the tunnel in the first place, an appropriate route needs to be added as well — either
	// in the "[Routes]" section on the ".network" matching the wireguard interface, or externally to systemd-networkd.
	//
	// Added in version 237.
	AllowedIPs systemdconf.Value

	// Sets an endpoint IP address or hostname, followed by a colon, and then a port number. IPv6 address must be in the square brackets.
	// For example, "111.222.333.444:51820" for IPv4 and "[1111:2222::3333]:51820" for IPv6 address. This endpoint will
	// be updated automatically once to the most recent source IP address and port of correctly authenticated packets from the
	// peer at configuration time.
	//
	// This option honors the "@" prefix in the same way as the PrivateKey= setting of the [WireGuard] section.
	//
	// Added in version 237.
	Endpoint systemdconf.Value

	// Sets a seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for
	// the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely
	// sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from
	// having a persistent keepalive interval of 25 seconds. If set to 0 or "off", this option is disabled. By default or when unspecified,
	// this option is off. Most users will not need this.
	//
	// Added in version 237.
	PersistentKeepalive systemdconf.Value

	// The table identifier for the routes to the addresses specified in the AllowedIPs=. Takes a negative boolean value, one
	// of the predefined names "default", "main", and "local", names defined in RouteTable= in networkd.conf, or a number in
	// the range 1…4294967295. Defaults to unset, and the value specified in the same setting in the [WireGuard] section will
	// be used.
	//
	// Added in version 250.
	RouteTable systemdconf.Value

	// The priority of the routes to the addresses specified in the AllowedIPs=. Takes an integer in the range 0…4294967295. Defaults
	// to unset, and the value specified in the same setting in the [WireGuard] section will be used.
	//
	// Added in version 250.
	RouteMetric systemdconf.Value
}

// NetdevBondSection represents [Bond] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BBond%5D%20Section%20Options for details).
type NetdevBondSection struct {
	systemdconf.Section

	// Specifies one of the bonding policies. The default is "balance-rr" (round robin). Possible values are "balance-rr",
	// "active-backup", "balance-xor", "broadcast", "802.3ad", "balance-tlb", and "balance-alb".
	//
	// Added in version 216.
	Mode systemdconf.Value

	// Selects the transmit hash policy to use for slave selection in balance-xor, 802.3ad, and tlb modes. Possible values are
	// "layer2", "layer3+4", "layer2+3", "encap2+3", and "encap3+4".
	//
	// Added in version 216.
	TransmitHashPolicy systemdconf.Value

	// Specifies the rate with which link partner transmits Link Aggregation Control Protocol Data Unit packets in 802.3ad mode.
	// Possible values are "slow", which requests partner to transmit LACPDUs every 30 seconds, and "fast", which requests partner
	// to transmit LACPDUs every second. The default value is "slow".
	//
	// Added in version 216.
	LACPTransmitRate systemdconf.Value

	// Specifies the frequency that Media Independent Interface link monitoring will occur. A value of zero disables MII link
	// monitoring. This value is rounded down to the nearest millisecond. The default value is 0.
	//
	// Added in version 216.
	MIIMonitorSec systemdconf.Value

	// Specifies the number of seconds the delay between each peer notification (gratuitous ARP and unsolicited IPv6 Neighbor
	// Advertisement) when they are issued after a failover event. This delay should be a multiple of the MII link monitor interval
	// (miimon). The valid range is 0...300s. The default value is 0, which means to match the value of the MIIMonitorSec=.
	//
	// Added in version 256.
	PeerNotifyDelaySec systemdconf.Value

	// Specifies the delay before a link is enabled after a link up status has been detected. This value is rounded down to a multiple
	// of MIIMonitorSec=. The default value is 0.
	//
	// Added in version 216.
	UpDelaySec systemdconf.Value

	// Specifies the delay before a link is disabled after a link down status has been detected. This value is rounded down to a multiple
	// of MIIMonitorSec=. The default value is 0.
	//
	// Added in version 216.
	DownDelaySec systemdconf.Value

	// Specifies the number of seconds between instances where the bonding driver sends learning packets to each slave peer switch.
	// The valid range is 1…0x7fffffff; the default value is 1. This option has an effect only for the balance-tlb and balance-alb
	// modes.
	//
	// Added in version 220.
	LearnPacketIntervalSec systemdconf.Value

	// Specifies the 802.3ad aggregation selection logic to use. Possible values are "stable", "bandwidth" and "count".
	//
	// Added in version 220.
	AdSelect systemdconf.Value

	// Specifies the 802.3ad actor system priority. Takes a number in the range 1…65535.
	//
	// Added in version 240.
	AdActorSystemPriority systemdconf.Value

	// Specifies the 802.3ad user defined portion of the port key. Takes a number in the range 0…1023.
	//
	// Added in version 240.
	AdUserPortKey systemdconf.Value

	// Specifies the 802.3ad system MAC address. This cannot be a null or multicast address.
	//
	// Added in version 240.
	AdActorSystem systemdconf.Value

	// Specifies whether the active-backup mode should set all slaves to the same MAC address at the time of enslavement or, when
	// enabled, to perform special handling of the bond's MAC address in accordance with the selected policy. The default policy
	// is none. Possible values are "none", "active" and "follow".
	//
	// Added in version 220.
	FailOverMACPolicy systemdconf.Value

	// Specifies whether or not ARP probes and replies should be validated in any mode that supports ARP monitoring, or whether
	// non-ARP traffic should be filtered (disregarded) for link monitoring purposes. Possible values are "none", "active",
	// "backup" and "all".
	//
	// Added in version 220.
	ARPValidate systemdconf.Value

	// Specifies the ARP link monitoring frequency. A value of 0 disables ARP monitoring. The default value is 0, and the default
	// unit seconds.
	//
	// Added in version 220.
	ARPIntervalSec systemdconf.Value

	// Specifies the IP addresses to use as ARP monitoring peers when ARPIntervalSec= is greater than 0. These are the targets
	// of the ARP request sent to determine the health of the link to the targets. Specify these values in IPv4 dotted decimal format.
	// At least one IP address must be given for ARP monitoring to function. The maximum number of targets that can be specified
	// is 16. The default value is no IP addresses.
	//
	// Added in version 220.
	ARPIPTargets systemdconf.Value

	// Specifies the quantity of ARPIPTargets= that must be reachable in order for the ARP monitor to consider a slave as being
	// up. This option affects only active-backup mode for slaves with ARPValidate enabled. Possible values are "any" and "all".
	//
	// Added in version 220.
	ARPAllTargets systemdconf.Value

	// Specifies the reselection policy for the primary slave. This affects how the primary slave is chosen to become the active
	// slave when failure of the active slave or recovery of the primary slave occurs. This option is designed to prevent flip-flopping
	// between the primary slave and other slaves. Possible values are "always", "better" and "failure".
	//
	// Added in version 220.
	PrimaryReselectPolicy systemdconf.Value

	// Specifies the number of IGMP membership reports to be issued after a failover event. One membership report is issued immediately
	// after the failover, subsequent packets are sent in each 200ms interval. The valid range is 0…255. Defaults to 1. A value
	// of 0 prevents the IGMP membership report from being issued in response to the failover event.
	//
	// Added in version 220.
	ResendIGMP systemdconf.Value

	// Specify the number of packets to transmit through a slave before moving to the next one. When set to 0, then a slave is chosen
	// at random. The valid range is 0…65535. Defaults to 1. This option only has effect when in balance-rr mode.
	//
	// Added in version 220.
	PacketsPerSlave systemdconf.Value

	// Specify the number of peer notifications (gratuitous ARPs and unsolicited IPv6 Neighbor Advertisements) to be issued
	// after a failover event. As soon as the link is up on the new slave, a peer notification is sent on the bonding device and each
	// VLAN sub-device. This is repeated at each link monitor interval (ARPIntervalSec or MIIMonitorSec, whichever is active)
	// if the number is greater than 1. The valid range is 0…255. The default value is 1. These options affect only the active-backup
	// mode.
	//
	// Added in version 220.
	GratuitousARP systemdconf.Value

	// Takes a boolean. Specifies that duplicate frames (received on inactive ports) should be dropped when false, or delivered
	// when true. Normally, bonding will drop duplicate frames (received on inactive ports), which is desirable for most users.
	// But there are some times it is nice to allow duplicate frames to be delivered. The default value is false (drop duplicate
	// frames received on inactive ports).
	//
	// Added in version 220.
	AllSlavesActive systemdconf.Value

	// Takes a boolean. Specifies if dynamic shuffling of flows is enabled. Applies only for balance-tlb mode. Defaults to unset.
	//
	// Added in version 240.
	DynamicTransmitLoadBalancing systemdconf.Value

	// Specifies the minimum number of links that must be active before asserting carrier. The default value is 0.
	//
	// Added in version 220.
	MinLinks systemdconf.Value

	// Specify the maximum number of arp interval monitor cycle for missed ARP replies. If this number is exceeded, link is reported
	// as down. Defaults to unset.
	//
	// Added in version 256.
	ARPMissedMax systemdconf.Value
}

// NetdevXfrmSection represents [Xfrm] section
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BXfrm%5D%20Section%20Options for details).
type NetdevXfrmSection struct {
	systemdconf.Section

	// Sets the ID/key of the xfrm interface which needs to be associated with a SA/policy. Can be decimal or hexadecimal, valid
	// range is 1-0xffffffff. This is mandatory.
	//
	// Added in version 243.
	InterfaceId systemdconf.Value

	// Takes a boolean. If false (the default), the xfrm interface must have an underlying device which can be used for hardware
	// offloading.
	//
	// Added in version 243.
	Independent systemdconf.Value
}

// NetdevVRFSection represents [VRF] section, which only applies to netdevs of kind "vrf"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BVRF%5D%20Section%20Options for details).
type NetdevVRFSection struct {
	systemdconf.Section

	// The numeric routing table identifier. This setting is compulsory.
	//
	// Added in version 243.
	Table systemdconf.Value
}

// NetdevBatmanAdvancedSection represents [BatmanAdvanced] section, which only applies to netdevs of kind "batadv"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BBatmanAdvanced%5D%20Section%20Options for details).
type NetdevBatmanAdvancedSection struct {
	systemdconf.Section

	// Takes one of "off", "server", or "client". A batman-adv node can either run in server mode (sharing its internet connection
	// with the mesh) or in client mode (searching for the most suitable internet connection in the mesh) or having the gateway
	// support turned off entirely (which is the default setting).
	//
	// Added in version 248.
	GatewayMode systemdconf.Value

	// Takes a boolean value. Enables or disables aggregation of originator messages. Defaults to true.
	//
	// Added in version 248.
	Aggregation systemdconf.Value

	// Takes a boolean value. Enables or disables avoidance of loops on bridges. Defaults to true.
	//
	// Added in version 248.
	BridgeLoopAvoidance systemdconf.Value

	// Takes a boolean value. Enables or disables the distributed ARP table. Defaults to true.
	//
	// Added in version 248.
	DistributedArpTable systemdconf.Value

	// Takes a boolean value. Enables or disables fragmentation. Defaults to true.
	//
	// Added in version 248.
	Fragmentation systemdconf.Value

	// The hop penalty setting allows one to modify batctl preference for multihop routes vs. short routes. This integer value
	// is applied to the TQ (Transmit Quality) of each forwarded OGM (Originator Message), thereby propagating the cost of an
	// extra hop (the packet has to be received and retransmitted which costs airtime). A higher hop penalty will make it more unlikely
	// that other nodes will choose this node as intermediate hop towards any given destination. The default hop penalty of '15'
	// is a reasonable value for most setups and probably does not need to be changed. However, mobile nodes could choose a value
	// of 255 (maximum value) to avoid being chosen as a router by other nodes. The minimum value is 0.
	//
	// Added in version 248.
	HopPenalty systemdconf.Value

	// The value specifies the interval in seconds, unless another time unit is specified in which batman-adv floods the network
	// with its protocol information. See systemd.time for more information.
	//
	// Added in version 248.
	OriginatorIntervalSec systemdconf.Value

	// If the node is a server, this parameter is used to inform other nodes in the network about this node's internet connection
	// download bandwidth in bits per second. Just enter any number suffixed with K, M, G or T (base 1000) and the batman-adv module
	// will propagate the entered value in the mesh.
	//
	// Added in version 248.
	GatewayBandwidthDown systemdconf.Value

	// If the node is a server, this parameter is used to inform other nodes in the network about this node's internet connection
	// upload bandwidth in bits per second. Just enter any number suffixed with K, M, G or T (base 1000) and the batman-adv module
	// will propagate the entered value in the mesh.
	//
	// Added in version 248.
	GatewayBandwidthUp systemdconf.Value

	// This can be either "batman-v" or "batman-iv" and describes which routing_algo of batctl to use. The algorithm cannot be
	// changed after interface creation. Defaults to "batman-v".
	//
	// Added in version 248.
	RoutingAlgorithm systemdconf.Value
}

// NetdevIPoIBSection represents [IPoIB] section, which only applies to netdevs of kind "ipoib"
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BIPoIB%5D%20Section%20Options for details).
type NetdevIPoIBSection struct {
	systemdconf.Section

	// Takes an integer in the range 1…0xffff, except for 0x8000. Defaults to unset, and the kernel's default is used.
	//
	// Added in version 250.
	PartitionKey systemdconf.Value

	// Takes one of the special values "datagram" or "connected". Defaults to unset, and the kernel's default is used.
	//
	// When "datagram", the Infiniband unreliable datagram (UD) transport is used, and so the interface MTU is equal to the IB
	// L2 MTU minus the IPoIB encapsulation header (4 bytes). For example, in a typical IB fabric with a 2K MTU, the IPoIB MTU will
	// be 2048 - 4 = 2044 bytes.
	//
	// When "connected", the Infiniband reliable connected (RC) transport is used. Connected mode takes advantage of the connected
	// nature of the IB transport and allows an MTU up to the maximal IP packet size of 64K, which reduces the number of IP packets
	// needed for handling large UDP datagrams, TCP segments, etc and increases the performance for large messages.
	//
	// Added in version 250.
	Mode systemdconf.Value

	// Takes an boolean value. When true, the kernel ignores multicast groups handled by userspace. Defaults to unset, and the
	// kernel's default is used.
	//
	// Added in version 250.
	IgnoreUserspaceMulticastGroup systemdconf.Value
}

// NetdevWLANSection represents [WLAN] section, which only applies to WLAN interfaces
// (see https://www.freedesktop.org/software/systemd/man/systemd.netdev.html#%5BWLAN%5D%20Section%20Options for details).
type NetdevWLANSection struct {
	systemdconf.Section

	// Specifies the name or index of the physical WLAN device (e.g. "0" or "phy0"). The list of the physical WLAN devices that exist
	// on the host can be obtained by iw phy command. This option is mandatory.
	//
	// Added in version 251.
	PhysicalDevice systemdconf.Value

	// Specifies the type of the interface. Takes one of the "ad-hoc", "station", "ap", "ap-vlan", "wds", "monitor", "mesh-point",
	// "p2p-client", "p2p-go", "p2p-device", "ocb", and "nan". This option is mandatory.
	//
	// Added in version 251.
	Type systemdconf.Value

	// Enables the Wireless Distribution System (WDS) mode on the interface. The mode is also known as the "4 address mode". Takes
	// a boolean value. Defaults to unset, and the kernel's default will be used.
	//
	// Added in version 251.
	WDS systemdconf.Value
}
