// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package network

import "github.com/sergeymakinen/go-systemdconf"

// NetdevMatchSection represents parameters which determines if a given virtual network device may be created
type NetdevMatchSection struct {
	systemdconf.Section

	// Matches against the hostname or machine ID of the host. See "ConditionHost=" in systemd.unit for details. When prefixed
	// with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned value is
	// cleared.
	Host systemdconf.Value

	// Checks whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation.
	// See "ConditionVirtualization=" in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result
	// is negated. If an empty string is assigned, then previously assigned value is cleared.
	Virtualization systemdconf.Value

	// Checks whether a specific kernel command line option is set. See "ConditionKernelCommandLine=" in systemd.unit for
	// details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously
	// assigned value is cleared.
	KernelCommandLine systemdconf.Value

	// Checks whether the kernel version (as reported by uname -r) matches a certain expression. See "ConditionKernelVersion="
	// in systemd.unit for details. When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is
	// assigned, then previously assigned value is cleared.
	KernelVersion systemdconf.Value

	// Checks whether the system is running on a specific architecture. See "ConditionArchitecture=" in systemd.unit for details.
	// When prefixed with an exclamation mark ("!"), the result is negated. If an empty string is assigned, then previously assigned
	// value is cleared.
	Architecture systemdconf.Value
}

// NetdevNetDevSection represents Virtual Network Device parameters
type NetdevNetDevSection struct {
	systemdconf.Section

	// A free-form description of the netdev.
	Description systemdconf.Value

	// The interface name used when creating the netdev. This option is compulsory.
	Name systemdconf.Value

	// The netdev kind. This option is compulsory. See the "Supported netdev kinds" section for the valid keys.
	Kind systemdconf.Value

	// The maximum transmission unit in bytes to set for the device. The usual suffixes K, M, G, are supported and are understood
	// to the base of 1024. For "tun" or "tap" devices, MTUBytes= setting is not currently supported in "[NetDev]" section. Please
	// specify it in "[Link]" section of corresponding systemd.network files.
	MTUBytes systemdconf.Value

	// The MAC address to use for the device. For "tun" or "tap" devices, setting MACAddress= in the "[NetDev]" section is not supported.
	// Please specify it in "[Link]" section of the corresponding systemd.network file. If this option is not set, "vlan" devices
	// inherit the MAC address of the physical interface. For other kind of netdevs, if this option is not set, then MAC address
	// is generated based on the interface name and the machine-id.
	MACAddress systemdconf.Value
}

// NetdevBridgeSection represents netdev of kind "bridge" parameters
type NetdevBridgeSection struct {
	systemdconf.Section

	// HelloTimeSec specifies the number of seconds between two hello packets sent out by the root bridge and the designated bridges.
	// Hello packets are used to communicate information about the topology throughout the entire bridged local area network.
	HelloTimeSec systemdconf.Value

	// MaxAgeSec specifies the number of seconds of maximum message age. If the last seen (received) hello packet is more than
	// this number of seconds old, the bridge in question will start the takeover procedure in attempt to become the Root Bridge
	// itself.
	MaxAgeSec systemdconf.Value

	// ForwardDelaySec specifies the number of seconds spent in each of the Listening and Learning states before the Forwarding
	// state is entered.
	ForwardDelaySec systemdconf.Value

	// This specifies the number of seconds a MAC Address will be kept in the forwarding database after having a packet received
	// from this MAC Address.
	AgeingTimeSec systemdconf.Value

	// The priority of the bridge. An integer between 0 and 65535. A lower value means higher priority. The bridge having the lowest
	// priority will be elected as root bridge.
	Priority systemdconf.Value

	// A 16-bit bitmask represented as an integer which allows forwarding of link local frames with 802.1D reserved addresses
	// (01:80:C2:00:00:0X). A logical AND is performed between the specified bitmask and the exponentiation of 2^X, the lower
	// nibble of the last octet of the MAC address. For example, a value of 8 would allow forwarding of frames addressed to 01:80:C2:00:00:03
	// (802.1X PAE).
	GroupForwardMask systemdconf.Value

	// This specifies the default port VLAN ID of a newly attached bridge port. Set this to an integer in the range 1–4094 or "none"
	// to disable the PVID.
	DefaultPVID systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_QUERIER option in the kernel. If enabled, the kernel will send
	// general ICMP queries from a zero source address. This feature should allow faster convergence on startup, but it causes
	// some multicast-aware switches to misbehave and disrupt forwarding of multicast packets. When unset, the kernel's default
	// will be used.
	MulticastQuerier systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_MCAST_SNOOPING option in the kernel. If enabled, IGMP snooping monitors
	// the Internet Group Management Protocol (IGMP) traffic between hosts and multicast routers. When unset, the kernel's
	// default will be used.
	MulticastSnooping systemdconf.Value

	// Takes a boolean. This setting controls the IFLA_BR_VLAN_FILTERING option in the kernel. If enabled, the bridge will be
	// started in VLAN-filtering mode. When unset, the kernel's default will be used.
	VLANFiltering systemdconf.Value

	// Takes a boolean. This enables the bridge's Spanning Tree Protocol (STP). When unset, the kernel's default will be used.
	STP systemdconf.Value

	// Allows to change bridge's multicast Internet Group Management Protocol (IGMP) version. Takes an interger 2 or 3. When
	// unset, the kernel's default will be used.
	MulticastIGMPVersion systemdconf.Value
}

// NetdevVLANSection represents netdev of kind "vlan" parameters
type NetdevVLANSection struct {
	systemdconf.Section

	// The VLAN ID to use. An integer in the range 0–4094. This option is compulsory.
	Id systemdconf.Value

	// Takes a boolean. The Generic VLAN Registration Protocol (GVRP) is a protocol that allows automatic learning of VLANs on
	// a network. When unset, the kernel's default will be used.
	GVRP systemdconf.Value

	// Takes a boolean. Multiple VLAN Registration Protocol (MVRP) formerly known as GARP VLAN Registration Protocol (GVRP)
	// is a standards-based Layer 2 network protocol, for automatic configuration of VLAN information on switches. It was defined
	// in the 802.1ak amendment to 802.1Q-2005. When unset, the kernel's default will be used.
	MVRP systemdconf.Value

	// Takes a boolean. The VLAN loose binding mode, in which only the operational state is passed from the parent to the associated
	// VLANs, but the VLAN device state is not changed. When unset, the kernel's default will be used.
	LooseBinding systemdconf.Value

	// Takes a boolean. The VLAN reorder header is set VLAN interfaces behave like physical interfaces. When unset, the kernel's
	// default will be used.
	ReorderHeader systemdconf.Value
}

// NetdevMACVLANSection represents netdev of kind "macvlan" parameters
type NetdevMACVLANSection struct {
	systemdconf.Section

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", and "passthru".
	Mode systemdconf.Value
}

// NetdevMACVTAPSection represents netdev of kind "macvtap" parameters
type NetdevMACVTAPSection struct {
	systemdconf.Section

	// The MACVLAN mode to use. The supported options are "private", "vepa", "bridge", and "passthru".
	Mode systemdconf.Value
}

// NetdevIPVLANSection represents netdev of kind "ipvlan" parameters
type NetdevIPVLANSection struct {
	systemdconf.Section

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	Mode systemdconf.Value

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	Flags systemdconf.Value
}

// NetdevIPVTAPSection represents netdev of kind "ipvtap" parameters
type NetdevIPVTAPSection struct {
	systemdconf.Section

	// The IPVLAN mode to use. The supported options are "L2","L3" and "L3S".
	Mode systemdconf.Value

	// The IPVLAN flags to use. The supported options are "bridge","private" and "vepa".
	Flags systemdconf.Value
}

// NetdevVXLANSection represents netdev of kind "vxlan" parameters
type NetdevVXLANSection struct {
	systemdconf.Section

	// The VXLAN Network Identifier (or VXLAN Segment ID). Takes a number in the range 1-16777215.
	VNI systemdconf.Value

	// Configures destination IP address.
	Remote systemdconf.Value

	// Configures local IP address.
	Local systemdconf.Value

	// Configures VXLAN multicast group IP address. All members of a VXLAN must use the same multicast group address.
	Group systemdconf.Value

	// The Type Of Service byte value for a vxlan interface.
	TOS systemdconf.Value

	// A fixed Time To Live N on Virtual eXtensible Local Area Network packets. Takes "inherit" or a number in the range 0–255. 0
	// is a special value meaning inherit the inner protocol's TTL value. "inherit" means that it will inherit the outer protocol's
	// TTL value.
	TTL systemdconf.Value

	// Takes a boolean. When true, enables dynamic MAC learning to discover remote MAC addresses.
	MacLearning systemdconf.Value

	// The lifetime of Forwarding Database entry learnt by the kernel, in seconds.
	FDBAgeingSec systemdconf.Value

	// Configures maximum number of FDB entries.
	MaximumFDBEntries systemdconf.Value

	// Takes a boolean. When true, bridge-connected VXLAN tunnel endpoint answers ARP requests from the local bridge on behalf
	// of remote Distributed Overlay Virtual Ethernet  (DVOE) clients. Defaults to false.
	ReduceARPProxy systemdconf.Value

	// Takes a boolean. When true, enables netlink LLADDR miss notifications.
	L2MissNotification systemdconf.Value

	// Takes a boolean. When true, enables netlink IP address miss notifications.
	L3MissNotification systemdconf.Value

	// Takes a boolean. When true, route short circuiting is turned on.
	RouteShortCircuit systemdconf.Value

	// Takes a boolean. When true, transmitting UDP checksums when doing VXLAN/IPv4 is turned on.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, sending zero checksums in VXLAN/IPv6 is turned on.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, receiving zero checksums in VXLAN/IPv6 is turned on.
	UDP6ZeroChecksumRx systemdconf.Value

	// Takes a boolean. When true, remote transmit checksum offload of VXLAN is turned on.
	RemoteChecksumTx systemdconf.Value

	// Takes a boolean. When true, remote receive checksum offload in VXLAN is turned on.
	RemoteChecksumRx systemdconf.Value

	// Takes a boolean. When true, it enables Group Policy VXLAN extension security label mechanism across network peers based
	// on VXLAN. For details about the Group Policy VXLAN, see the  VXLAN Group Policy  document. Defaults to false.
	GroupPolicyExtension systemdconf.Value

	// Takes a boolean. When true, Generic Protocol Extension extends the existing VXLAN protocol to provide protocol typing,
	// OAM, and versioning capabilities. For details about the VXLAN GPE Header, see the  Generic Protocol Extension for VXLAN
	//  document. If destination port is not specified and Generic Protocol Extension is set then default port of 4790 is used.
	// Defaults to false.
	GenericProtocolExtension systemdconf.Value

	// Configures the default destination UDP port on a per-device basis. If destination port is not specified then Linux kernel
	// default will be used. Set destination port 4789 to get the IANA assigned value. If not set or if the destination port is assigned
	// the empty string the default port of 4789 is used.
	DestinationPort systemdconf.Value

	// Configures VXLAN port range. VXLAN bases source UDP port based on flow to help the receiver to be able to load balance based
	// on outer header flow. It restricts the port range to the normal UDP local ports, and allows overriding via configuration.
	PortRange systemdconf.Value

	// Specifies the flow label to use in outgoing packets. The valid range is 0-1048575.
	FlowLabel systemdconf.Value

	// Allows to set the IPv4 Do not Fragment (DF) bit in outgoing packets, or to inherit its value from the IPv4 inner header. Takes
	// a boolean value, or "inherit". Set to "inherit" if the encapsulated protocol is IPv6. When unset, the kernel's default
	// will be used.
	IPDoNotFragment systemdconf.Value
}

// NetdevGENEVESection represents netdev of kind "geneve" parameters
type NetdevGENEVESection struct {
	systemdconf.Section

	// Specifies the Virtual Network Identifier (VNI) to use. Ranges [0-16777215]. This field is mandatory.
	Id systemdconf.Value

	// Specifies the unicast destination IP address to use in outgoing packets.
	Remote systemdconf.Value

	// Specifies the TOS value to use in outgoing packets. Ranges [1-255].
	TOS systemdconf.Value

	// Accepts the same key in "[VXLAN]" section except when unset or set to 0, the kernel's default will be used meaning that packets
	// TTL will be set from /proc/sys/net/ipv4/ip_default_ttl.
	TTL systemdconf.Value

	// Takes a boolean. When true, specifies if UDP checksum is calculated for transmitted packets over IPv4.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	UDP6ZeroChecksumRx systemdconf.Value

	// Specifies destination port. Defaults to 6081. If not set or assigned the empty string, the default port of 6081 is used.
	DestinationPort systemdconf.Value

	// Specifies the flow label to use in outgoing packets.
	FlowLabel systemdconf.Value

	// Accepts the same key in "[VXLAN]" section.
	IPDoNotFragment systemdconf.Value
}

// NetdevL2TPSection represents netdev of kind "l2tp" parameters
type NetdevL2TPSection struct {
	systemdconf.Section

	// Specifies the tunnel id. The value used must match the "PeerTunnelId=" value being used at the peer. Ranges a number between
	// 1 and 4294967295). This option is compulsory.
	TunnelId systemdconf.Value

	// Specifies the peer tunnel id. The value used must match the "PeerTunnelId=" value being used at the peer. Ranges a number
	// between 1 and 4294967295). This option is compulsory.
	PeerTunnelId systemdconf.Value

	// Specifies the IP address of the remote peer. This option is compulsory.
	Remote systemdconf.Value

	// Specifies the IP address of the local interface. Takes an IP address, or the special values "auto", "static", or "dynamic".
	// When an address is set, then the local interface must have the address. If "auto", then one of the addresses on the local interface
	// is used. Similarly, if "static" or "dynamic" is set, then one of the static or dynamic addresses on the local interface is
	// used. Defaults to "auto".
	Local systemdconf.Value

	// Specifies the encapsulation type of the tunnel. Takes one of "udp" or "ip".
	EncapsulationType systemdconf.Value

	// Specifies the UDP source port to be used for the tunnel. When UDP encapsulation is selected it's mandotory. Ignored when
	// ip encapsulation is selected.
	UDPSourcePort systemdconf.Value

	// Specifies destination port. When UDP encapsulation is selected it's mandotory. Ignored when ip encapsulation is selected.
	DestinationPort systemdconf.Value

	// Takes a boolean. When true, specifies if UDP checksum is calculated for transmitted packets over IPv4.
	UDPChecksum systemdconf.Value

	// Takes a boolean. When true, skip UDP checksum calculation for transmitted packets over IPv6.
	UDP6ZeroChecksumTx systemdconf.Value

	// Takes a boolean. When true, allows incoming UDP packets over IPv6 with zero checksum field.
	UDP6ZeroChecksumRx systemdconf.Value
}

// NetdevL2TPSessionSection represents netdev of kind "l2tp" parameters
type NetdevL2TPSessionSection struct {
	systemdconf.Section

	// Specifies the name of the session. This option is compulsory.
	Name systemdconf.Value

	// Specifies the session id. The value used must match the "SessionId=" value being used at the peer. Ranges a number between
	// 1 and 4294967295). This option is compulsory.
	SessionId systemdconf.Value

	// Specifies the peer session id. The value used must match the "PeerSessionId=" value being used at the peer. Ranges a number
	// between 1 and 4294967295). This option is compulsory.
	PeerSessionId systemdconf.Value

	// Specifies layer2specific header type of the session. One of "none" or "default". Defaults to "default".
	Layer2SpecificHeader systemdconf.Value
}

// NetdevMACsecSection represents netdev of kind "macsec" parameters
type NetdevMACsecSection struct {
	systemdconf.Section

	// Specifies the port to be used for the MACsec transmit channel. The port is used to make secure channel identifier (SCI).
	// Takes a value between 1 and 65535. Defaults to unset.
	Port systemdconf.Value

	// Takes a boolean. When true, enable encryption. Defaults to unset.
	Encrypt systemdconf.Value
}

// NetdevMACsecReceiveChannelSection represents netdev of kind "macsec" parameters
type NetdevMACsecReceiveChannelSection struct {
	systemdconf.Section

	// Specifies the port to be used for the MACsec receive channel. The port is used to make secure channel identifier (SCI). Takes
	// a value between 1 and 65535. This option is compulsory, and is not set by default.
	Port systemdconf.Value

	// Specifies the MAC address to be used for the MACsec receive channel. The MAC address used to make secure channel identifier
	// (SCI). This option is compulsory, and is not set by default.
	MACAddress systemdconf.Value
}

// NetdevMACsecTransmitAssociationSection represents netdev of kind "macsec" parameters
type NetdevMACsecTransmitAssociationSection struct {
	systemdconf.Section

	// Specifies the packet number to be used for replay protection and the construction of the initialization vector (along
	// with the secure channel identifier [SCI]). Takes a value between 1-4,294,967,295. Defaults to unset.
	PacketNumber systemdconf.Value

	// Specifies the identification for the key. Takes a number between 0-255. This option is compulsory, and is not set by default.
	KeyId systemdconf.Value

	// Specifies the encryption key used in the transmission channel. The same key must be configured on the peer’s matching receive
	// channel. This option is compulsory, and is not set by default. Takes a 128-bit key encoded in a hexadecimal string, for example
	// "dffafc8d7b9a43d5b9a3dfbbf6a30c16".
	Key systemdconf.Value

	// Takes a absolute path to a file which contains a 128-bit key encoded in a hexadecimal string, which will be used in the transmission
	// channel. When this option is specified, Key= is ignored. Note that the file must be readable by the user "systemd-network",
	// so it should be, e.g., owned by "root:systemd-network" with a "0640" file mode.
	KeyFile systemdconf.Value

	// Takes a boolean. If enabled, then the security association is activated. Defaults to unset.
	Activate systemdconf.Value

	// Takes a boolean. If enabled, then the security association is used for encoding. Only one "[MACsecTransmitAssociation]"
	// section can enable this option. When enabled, Activate=yes is implied. Defaults to unset.
	UseForEncoding systemdconf.Value
}

// NetdevMACsecReceiveAssociationSection represents netdev of kind "macsec" parameters
type NetdevMACsecReceiveAssociationSection struct {
	systemdconf.Section

	// Accepts the same key in "[MACsecReceiveChannel]" section.
	Port systemdconf.Value

	// Accepts the same key in "[MACsecReceiveChannel]" section.
	MACAddress systemdconf.Value

	// Accepts the same key in "[MACsecTransmitAssociation]" section.
	PacketNumber systemdconf.Value

	// Accepts the same key in "[MACsecTransmitAssociation]" section.
	KeyId systemdconf.Value

	// Accepts the same key in "[MACsecTransmitAssociation]" section.
	Key systemdconf.Value

	// Accepts the same key in "[MACsecTransmitAssociation]" section.
	KeyFile systemdconf.Value

	// Accepts the same key in "[MACsecTransmitAssociation]" section.
	Activate systemdconf.Value
}

// NetdevTunnelSection represents netdev of kind "ipip", "sit", "gre", "gretap", "ip6gre", "ip6gretap", "vti", "vti6", "ip6tnl", and "erspan" parameters
type NetdevTunnelSection struct {
	systemdconf.Section

	// A static local address for tunneled packets. It must be an address on another interface of this host, or the special value
	// "any".
	Local systemdconf.Value

	// The remote endpoint of the tunnel. Takes an IP address or the special value "any".
	Remote systemdconf.Value

	// The Type Of Service byte value for a tunnel interface. For details about the TOS, see the  Type of Service in the Internet Protocol
	// Suite  document.
	TOS systemdconf.Value

	// A fixed Time To Live N on tunneled packets. N is a number in the range 1–255. 0 is a special value meaning that packets inherit
	// the TTL value. The default value for IPv4 tunnels is: inherit. The default value for IPv6 tunnels is 64.
	TTL systemdconf.Value

	// Takes a boolean. When true, enables Path MTU Discovery on the tunnel.
	DiscoverPathMTU systemdconf.Value

	// Configures the 20-bit flow label (see  RFC 6437) field in the IPv6 header (see  RFC 2460), which is used by a node to label packets
	// of a flow. It is only used for IPv6 tunnels. A flow label of zero is used to indicate packets that have not been labeled. It can
	// be configured to a value in the range 0–0xFFFFF, or be set to "inherit", in which case the original flowlabel is used.
	IPv6FlowLabel systemdconf.Value

	// Takes a boolean. When true, the Differentiated Service Code Point (DSCP) field will be copied to the inner header from outer
	// header during the decapsulation of an IPv6 tunnel packet. DSCP is a field in an IP packet that enables different levels of
	// service to be assigned to network traffic. Defaults to "no".
	CopyDSCP systemdconf.Value

	// The Tunnel Encapsulation Limit option specifies how many additional levels of encapsulation are permitted to be prepended
	// to the packet. For example, a Tunnel Encapsulation Limit option containing a limit value of zero means that a packet carrying
	// that option may not enter another tunnel before exiting the current tunnel. (see  RFC 2473). The valid range is 0–255 and
	// "none". Defaults to 4.
	EncapsulationLimit systemdconf.Value

	// The Key= parameter specifies the same key to use in both directions (InputKey= and OutputKey=). The Key= is either a number
	// or an IPv4 address-like dotted quad. It is used as mark-configured SAD/SPD entry as part of the lookup key (both in data and
	// control path) in ip xfrm (framework used to implement IPsec protocol). See  ip-xfrm — transform configuration for details.
	// It is only used for VTI/VTI6, GRE, GRETAP, and ERSPAN tunnels.
	Key systemdconf.Value

	// The InputKey= parameter specifies the key to use for input. The format is same as Key=. It is only used for VTI/VTI6, GRE,
	// GRETAP, and ERSPAN tunnels.
	InputKey systemdconf.Value

	// The OutputKey= parameter specifies the key to use for output. The format is same as Key=. It is only used for VTI/VTI6, GRE,
	// GRETAP, and ERSPAN tunnels.
	OutputKey systemdconf.Value

	// An "ip6tnl" tunnel can be in one of three modes "ip6ip6" for IPv6 over IPv6, "ipip6" for IPv4 over IPv6 or "any" for either.
	Mode systemdconf.Value

	// Takes a boolean. When true tunnel does not require .network file. Created as "tunnel@NONE". Defaults to "false".
	Independent systemdconf.Value

	// Takes a boolean. If set to "yes", the loopback interface "lo" is used as the underlying device of the tunnel interface. Defaults
	// to "no".
	AssignToLoopback systemdconf.Value

	// Takes a boolean. When true allows tunnel traffic on ip6tnl devices where the remote endpoint is a local host address. When
	// unset, the kernel's default will be used.
	AllowLocalRemote systemdconf.Value

	// Takes a boolean. Specifies whether FooOverUDP= tunnel is to be configured. Defaults to false. This takes effects only
	// for IPIP, SIT, GRE, and GRETAP tunnels. For more detail information see Foo over UDP
	FooOverUDP systemdconf.Value

	// This setting specifies the UDP destination port for encapsulation. This field is mandatory when FooOverUDP=yes, and
	// is not set by default.
	FOUDestinationPort systemdconf.Value

	// This setting specifies the UDP source port for encapsulation. Defaults to 0 — that is, the source port for packets is left
	// to the network stack to decide.
	FOUSourcePort systemdconf.Value

	// Accepts the same key as in the "[FooOverUDP]" section.
	Encapsulation systemdconf.Value

	// Reconfigure the tunnel for IPv6 Rapid Deployment, also known as 6rd. The value is an ISP-specific IPv6 prefix with a non-zero
	// length. Only applicable to SIT tunnels.
	IPv6RapidDeploymentPrefix systemdconf.Value

	// Takes a boolean. If set, configures the tunnel as Intra-Site Automatic Tunnel Addressing Protocol (ISATAP) tunnel. Only
	// applicable to SIT tunnels. When unset, the kernel's default will be used.
	ISATAP systemdconf.Value

	// Takes a boolean. If set to yes, then packets are serialized. Only applies for GRE, GRETAP, and ERSPAN tunnels. When unset,
	// the kernel's default will be used.
	SerializeTunneledPackets systemdconf.Value

	// Specifies the ERSPAN index field for the interface, an integer in the range 1-1048575 associated with the ERSPAN traffic's
	// source port and direction. This field is mandatory.
	ERSPANIndex systemdconf.Value
}

// NetdevFooOverUDPSection represents netdev of kind "fou" parameters
type NetdevFooOverUDPSection struct {
	systemdconf.Section

	// Specifies the encapsulation mechanism used to store networking packets of various protocols inside the UDP packets.
	// Supports the following values: "FooOverUDP" provides the simplest no frills model of UDP encapsulation, it simply encapsulates
	// packets directly in the UDP payload. "GenericUDPEncapsulation" is a generic and extensible encapsulation, it allows
	// encapsulation of packets for any IP protocol and optional data as part of the encapsulation. For more detailed information
	// see Generic UDP Encapsulation. Defaults to "FooOverUDP".
	Encapsulation systemdconf.Value

	// Specifies the port number, where the IP encapsulation packets will arrive. Please take note that the packets will arrive
	// with the encapsulation will be removed. Then they will be manually fed back into the network stack, and sent ahead for delivery
	// to the real destination. This option is mandatory.
	Port systemdconf.Value

	// Specifies the peer port number. Defaults to unset. Note that when peer port is set "Peer=" address is mandotory.
	PeerPort systemdconf.Value

	// The Protocol= specifies the protocol number of the packets arriving at the UDP port. When Encapsulation=FooOverUDP,
	// this field is mandatory and is not set by default. Takes an IP protocol name such as "gre" or "ipip", or an integer within the
	// range 1-255. When Encapsulation=GenericUDPEncapsulation, this must not be specified.
	Protocol systemdconf.Value

	// Configures peer IP address. Note that when peer address is set "PeerPort=" is mandotory.
	Peer systemdconf.Value

	// Configures local IP address.
	Local systemdconf.Value
}

// NetdevPeerSection represents netdev of kind "veth" parameters
type NetdevPeerSection struct {
	systemdconf.Section

	// The interface name used when creating the netdev. This option is compulsory.
	Name systemdconf.Value

	// The peer MACAddress, if not set, it is generated in the same way as the MAC address of the main interface.
	MACAddress systemdconf.Value
}

// NetdevVXCANSection represents netdev of kind "vxcan" parameters
type NetdevVXCANSection struct {
	systemdconf.Section

	// The peer interface name used when creating the netdev. This option is compulsory.
	Peer systemdconf.Value
}

// NetdevTunSection represents netdev of kind "tun" parameters
type NetdevTunSection struct {
	systemdconf.Section

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving.
	// Defaults to "no".
	MultiQueue systemdconf.Value

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol
	// bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	PacketInfo systemdconf.Value

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation
	// Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	VNetHeader systemdconf.Value

	// User to grant access to the /dev/net/tun device.
	User systemdconf.Value

	// Group to grant access to the /dev/net/tun device.
	Group systemdconf.Value
}

// NetdevTapSection represents netdev of kind "tap" parameters
type NetdevTapSection struct {
	systemdconf.Section

	// Takes a boolean. Configures whether to use multiple file descriptors (queues) to parallelize packets sending and receiving.
	// Defaults to "no".
	MultiQueue systemdconf.Value

	// Takes a boolean. Configures whether packets should be prepended with four extra bytes (two flag bytes and two protocol
	// bytes). If disabled, it indicates that the packets will be pure IP packets. Defaults to "no".
	PacketInfo systemdconf.Value

	// Takes a boolean. Configures IFF_VNET_HDR flag for a tun or tap device. It allows sending and receiving larger Generic Segmentation
	// Offload (GSO) packets. This may increase throughput significantly. Defaults to "no".
	VNetHeader systemdconf.Value

	// User to grant access to the /dev/net/tun device.
	User systemdconf.Value

	// Group to grant access to the /dev/net/tun device.
	Group systemdconf.Value
}

// NetdevWireGuardSection represents WireGuard parameters
type NetdevWireGuardSection struct {
	systemdconf.Section

	// The Base64 encoded private key for the interface. It can be generated using the wg genkey command (see wg). This option or
	// PrivateKeyFile= is mandatory to use WireGuard. Note that because this information is secret, you may want to set the permissions
	// of the .netdev file to be owned by "root:systemd-network" with a "0640" file mode.
	PrivateKey systemdconf.Value

	// Takes an absolute path to a file which contains the Base64 encoded private key for the interface. When this option is specified,
	// then PrivateKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g., owned
	// by "root:systemd-network" with a "0640" file mode.
	PrivateKeyFile systemdconf.Value

	// Sets UDP port for listening. Takes either value between 1 and 65535 or "auto". If "auto" is specified, the port is automatically
	// generated based on interface name. Defaults to "auto".
	ListenPort systemdconf.Value

	// Sets a firewall mark on outgoing WireGuard packets from this interface. Takes a number between 1 and 4294967295.
	FirewallMark systemdconf.Value
}

// NetdevWireGuardPeerSection represents WireGuardPeer peer parameters
type NetdevWireGuardPeerSection struct {
	systemdconf.Section

	// Sets a Base64 encoded public key calculated by wg pubkey (see wg) from a private key, and usually transmitted out of band
	// to the author of the configuration file. This option is mandatory for this section.
	PublicKey systemdconf.Value

	// Optional preshared key for the interface. It can be generated by the wg genpsk command. This option adds an additional layer
	// of symmetric-key cryptography to be mixed into the already existing public-key cryptography, for post-quantum resistance.
	// Note that because this information is secret, you may want to set the permissions of the .netdev file to be owned by "root:systemd-networkd"
	// with a "0640" file mode.
	PresharedKey systemdconf.Value

	// Takes an absolute path to a file which contains the Base64 encoded preshared key for the peer. When this option is specified,
	// then PresharedKey= is ignored. Note that the file must be readable by the user "systemd-network", so it should be, e.g.,
	// owned by "root:systemd-network" with a "0640" file mode.
	PresharedKeyFile systemdconf.Value

	// Sets a comma-separated list of IP (v4 or v6) addresses with CIDR masks from which this peer is allowed to send incoming traffic
	// and to which outgoing traffic for this peer is directed. The catch-all 0.0.0.0/0 may be specified for matching all IPv4
	// addresses, and ::/0 may be specified for matching all IPv6 addresses.
	AllowedIPs systemdconf.Value

	// Sets an endpoint IP address or hostname, followed by a colon, and then a port number. This endpoint will be updated automatically
	// once to the most recent source IP address and port of correctly authenticated packets from the peer at configuration time.
	Endpoint systemdconf.Value

	// Sets a seconds interval, between 1 and 65535 inclusive, of how often to send an authenticated empty packet to the peer for
	// the purpose of keeping a stateful firewall or NAT mapping valid persistently. For example, if the interface very rarely
	// sends traffic, but it might at anytime receive traffic from a peer, and it is behind NAT, the interface might benefit from
	// having a persistent keepalive interval of 25 seconds. If set to 0 or "off", this option is disabled. By default or when unspecified,
	// this option is off. Most users will not need this.
	PersistentKeepalive systemdconf.Value
}

// NetdevBondSection represents Bond parameters
type NetdevBondSection struct {
	systemdconf.Section

	// Specifies one of the bonding policies. The default is "balance-rr" (round robin). Possible values are "balance-rr",
	// "active-backup", "balance-xor", "broadcast", "802.3ad", "balance-tlb", and "balance-alb".
	Mode systemdconf.Value

	// Selects the transmit hash policy to use for slave selection in balance-xor, 802.3ad, and tlb modes. Possible values are
	// "layer2", "layer3+4", "layer2+3", "encap2+3", and "encap3+4".
	TransmitHashPolicy systemdconf.Value

	// Specifies the rate with which link partner transmits Link Aggregation Control Protocol Data Unit packets in 802.3ad mode.
	// Possible values are "slow", which requests partner to transmit LACPDUs every 30 seconds, and "fast", which requests partner
	// to transmit LACPDUs every second. The default value is "slow".
	LACPTransmitRate systemdconf.Value

	// Specifies the frequency that Media Independent Interface link monitoring will occur. A value of zero disables MII link
	// monitoring. This value is rounded down to the nearest millisecond. The default value is 0.
	MIIMonitorSec systemdconf.Value

	// Specifies the delay before a link is enabled after a link up status has been detected. This value is rounded down to a multiple
	// of MIIMonitorSec. The default value is 0.
	UpDelaySec systemdconf.Value

	// Specifies the delay before a link is disabled after a link down status has been detected. This value is rounded down to a multiple
	// of MIIMonitorSec. The default value is 0.
	DownDelaySec systemdconf.Value

	// Specifies the number of seconds between instances where the bonding driver sends learning packets to each slave peer switch.
	// The valid range is 1–0x7fffffff; the default value is 1. This option has an effect only for the balance-tlb and balance-alb
	// modes.
	LearnPacketIntervalSec systemdconf.Value

	// Specifies the 802.3ad aggregation selection logic to use. Possible values are "stable", "bandwidth" and "count".
	AdSelect systemdconf.Value

	// Specifies the 802.3ad actor system priority. Ranges [1-65535].
	AdActorSystemPriority systemdconf.Value

	// Specifies the 802.3ad user defined portion of the port key. Ranges [0-1023].
	AdUserPortKey systemdconf.Value

	// Specifies the 802.3ad system mac address. This can not be either NULL or Multicast.
	AdActorSystem systemdconf.Value

	// Specifies whether the active-backup mode should set all slaves to the same MAC address at the time of enslavement or, when
	// enabled, to perform special handling of the bond's MAC address in accordance with the selected policy. The default policy
	// is none. Possible values are "none", "active" and "follow".
	FailOverMACPolicy systemdconf.Value

	// Specifies whether or not ARP probes and replies should be validated in any mode that supports ARP monitoring, or whether
	// non-ARP traffic should be filtered (disregarded) for link monitoring purposes. Possible values are "none", "active",
	// "backup" and "all".
	ARPValidate systemdconf.Value

	// Specifies the ARP link monitoring frequency. A value of 0 disables ARP monitoring. The default value is 0, and the default
	// unit seconds.
	ARPIntervalSec systemdconf.Value

	// Specifies the IP addresses to use as ARP monitoring peers when ARPIntervalSec is greater than 0. These are the targets of
	// the ARP request sent to determine the health of the link to the targets. Specify these values in IPv4 dotted decimal format.
	// At least one IP address must be given for ARP monitoring to function. The maximum number of targets that can be specified
	// is 16. The default value is no IP addresses.
	ARPIPTargets systemdconf.Value

	// Specifies the quantity of ARPIPTargets that must be reachable in order for the ARP monitor to consider a slave as being up.
	// This option affects only active-backup mode for slaves with ARPValidate enabled. Possible values are "any" and "all".
	ARPAllTargets systemdconf.Value

	// Specifies the reselection policy for the primary slave. This affects how the primary slave is chosen to become the active
	// slave when failure of the active slave or recovery of the primary slave occurs. This option is designed to prevent flip-flopping
	// between the primary slave and other slaves. Possible values are "always", "better" and "failure".
	PrimaryReselectPolicy systemdconf.Value

	// Specifies the number of IGMP membership reports to be issued after a failover event. One membership report is issued immediately
	// after the failover, subsequent packets are sent in each 200ms interval. The valid range is 0–255. Defaults to 1. A value
	// of 0 prevents the IGMP membership report from being issued in response to the failover event.
	ResendIGMP systemdconf.Value

	// Specify the number of packets to transmit through a slave before moving to the next one. When set to 0, then a slave is chosen
	// at random. The valid range is 0–65535. Defaults to 1. This option only has effect when in balance-rr mode.
	PacketsPerSlave systemdconf.Value

	// Specify the number of peer notifications (gratuitous ARPs and unsolicited IPv6 Neighbor Advertisements) to be issued
	// after a failover event. As soon as the link is up on the new slave, a peer notification is sent on the bonding device and each
	// VLAN sub-device. This is repeated at each link monitor interval (ARPIntervalSec or MIIMonitorSec, whichever is active)
	// if the number is greater than 1. The valid range is 0–255. The default value is 1. These options affect only the active-backup
	// mode.
	GratuitousARP systemdconf.Value

	// Takes a boolean. Specifies that duplicate frames (received on inactive ports) should be dropped when false, or delivered
	// when true. Normally, bonding will drop duplicate frames (received on inactive ports), which is desirable for most users.
	// But there are some times it is nice to allow duplicate frames to be delivered. The default value is false (drop duplicate
	// frames received on inactive ports).
	AllSlavesActive systemdconf.Value

	// Takes a boolean. Specifies if dynamic shuffling of flows is enabled. Applies only for balance-tlb mode. Defaults to unset.
	DynamicTransmitLoadBalancing systemdconf.Value

	// Specifies the minimum number of links that must be active before asserting carrier. The default value is 0.
	MinLinks systemdconf.Value
}

// NetdevXfrmSection represents xfrm interface parameters
type NetdevXfrmSection struct {
	systemdconf.Section

	// Sets the ID/key of the xfrm interface which needs to be associated with a SA/policy. Can be decimal or hexadecimal, valid
	// range is 0-0xffffffff, defaults to 0.
	InterfaceId systemdconf.Value

	// Takes a boolean. If set to "no", the xfrm interface should have an underlying device which can be used for hardware offloading.
	// Defaults to "no". See systemd.network for how to configure the underlying device.
	Independent systemdconf.Value
}

// NetdevVRFSection represents netdev of kind "vrf" parameters
type NetdevVRFSection struct {
	systemdconf.Section

	// The numeric routing table identifier. This option is compulsory.
	Table systemdconf.Value
}

// NetdevFile represents Virtual Network Device parameters
type NetdevFile struct {
	systemdconf.File

	Match                     NetdevMatchSection                     // Parameters which determines if a given virtual network device may be created
	NetDev                    NetdevNetDevSection                    // Virtual Network Device parameters
	Bridge                    NetdevBridgeSection                    // Netdev of kind "bridge" parameters
	VLAN                      NetdevVLANSection                      // Netdev of kind "vlan" parameters
	MACVLAN                   NetdevMACVLANSection                   // Netdev of kind "macvlan" parameters
	MACVTAP                   NetdevMACVTAPSection                   // Netdev of kind "macvtap" parameters
	IPVLAN                    NetdevIPVLANSection                    // Netdev of kind "ipvlan" parameters
	IPVTAP                    NetdevIPVTAPSection                    // Netdev of kind "ipvtap" parameters
	VXLAN                     NetdevVXLANSection                     // Netdev of kind "vxlan" parameters
	GENEVE                    NetdevGENEVESection                    // Netdev of kind "geneve" parameters
	L2TP                      NetdevL2TPSection                      // Netdev of kind "l2tp" parameters
	L2TPSession               NetdevL2TPSessionSection               // Netdev of kind "l2tp" parameters
	MACsec                    NetdevMACsecSection                    // Netdev of kind "macsec" parameters
	MACsecReceiveChannel      NetdevMACsecReceiveChannelSection      // Netdev of kind "macsec" parameters
	MACsecTransmitAssociation NetdevMACsecTransmitAssociationSection // Netdev of kind "macsec" parameters
	MACsecReceiveAssociation  NetdevMACsecReceiveAssociationSection  // Netdev of kind "macsec" parameters
	Tunnel                    NetdevTunnelSection                    // Netdev of kind "ipip", "sit", "gre", "gretap", "ip6gre", "ip6gretap", "vti", "vti6", "ip6tnl", and "erspan" parameters
	FooOverUDP                NetdevFooOverUDPSection                // Netdev of kind "fou" parameters
	Peer                      NetdevPeerSection                      // Netdev of kind "veth" parameters
	VXCAN                     NetdevVXCANSection                     // Netdev of kind "vxcan" parameters
	Tun                       NetdevTunSection                       // Netdev of kind "tun" parameters
	Tap                       NetdevTapSection                       // Netdev of kind "tap" parameters
	WireGuard                 NetdevWireGuardSection                 // WireGuard parameters
	WireGuardPeer             NetdevWireGuardPeerSection             // WireGuardPeer peer parameters
	Bond                      NetdevBondSection                      // Bond parameters
	Xfrm                      NetdevXfrmSection                      // Xfrm interface parameters
	VRF                       NetdevVRFSection                       // Netdev of kind "vrf" parameters
}
