// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package network

import "github.com/sergeymakinen/go-systemdconf/v3"

// ResolvedFile represents resolved.conf, resolved.conf.d — Network Name Resolution configuration files
// (see https://www.freedesktop.org/software/systemd/man/resolved.conf.html for details).
type ResolvedFile struct {
	systemdconf.File

	Resolve ResolvedResolveSection // [Resolve] section
}

// ResolvedResolveSection represents [Resolve] section
// (see https://www.freedesktop.org/software/systemd/man/resolved.conf.html#Options for details).
type ResolvedResolveSection struct {
	systemdconf.Section

	// A space-separated list of IPv4 and IPv6 addresses to use as system DNS servers. Each address can optionally take a port number
	// separated with ":", a network interface name or index separated with "%", and a Server Name Indication (SNI) separated
	// with "#". When IPv6 address is specified with a port number, then the address must be in the square brackets. That is, the
	// acceptable full formats are "111.222.333.444:9953%ifname#example.com" for IPv4 and "[1111:2222::3333]:9953%ifname#example.com"
	// for IPv6. DNS requests are sent to one of the listed DNS servers in parallel to suitable per-link DNS servers acquired from
	// systemd-networkd.service or set at runtime by external applications. For compatibility reasons, if this setting is
	// not specified, the DNS servers listed in /etc/resolv.conf are used instead, if that file exists and any servers are configured
	// in it. This setting defaults to the empty list.
	//
	// Added in version 213.
	DNS systemdconf.Value

	// A space-separated list of IPv4 and IPv6 addresses to use as the fallback DNS servers. Please see DNS= for acceptable format
	// of addresses. Any per-link DNS servers obtained from systemd-networkd.service take precedence over this setting, as
	// do any servers set via DNS= above or /etc/resolv.conf. This setting is hence only used if no other DNS server information
	// is known. If this option is not given, a compiled-in list of DNS servers is used instead.
	//
	// Added in version 216.
	FallbackDNS systemdconf.Value

	// A space-separated list of domains, optionally prefixed with "~", used for two distinct purposes described below. Defaults
	// to the empty list.
	//
	// Any domains not prefixed with "~" are used as search suffixes when resolving single-label hostnames (domain names which
	// contain no dot), in order to qualify them into fully-qualified domain names (FQDNs). These "search domains" are strictly
	// processed in the order they are specified in, until the name with the suffix appended is found. For compatibility reasons,
	// if this setting is not specified, the search domains listed in /etc/resolv.conf with the search keyword are used instead,
	// if that file exists and any domains are configured in it.
	//
	// The domains prefixed with "~" are called "route-only domains". All domains listed here (both search domains and route-only
	// domains after removing the "~" prefix) define a search path that preferably directs DNS queries to this interface. This
	// search path has an effect only when suitable per-link DNS servers are known. Such servers may be defined through the DNS=
	// setting (see above) and dynamically at run time, for example from DHCP leases. If no per-link DNS servers are known, route-only
	// domains have no effect.
	//
	// Use the construct "~." (which is composed from "~" to indicate a route-only domain and "." to indicate the DNS root domain
	// that is the implied suffix of all DNS domains) to use the DNS servers defined for this link preferably for all domains.
	//
	// See "Protocols and Routing" in systemd-resolved.service for details of how search and route-only domains are used.
	//
	// Note that configuring the MulticastDNS domain "local" as search or routing domain has the effect of routing lookups for
	// this domain to classic unicast DNS. This may be used to provide compatibility with legacy installations that use this domain
	// in a unicast DNS context, against the IANA assignment of this domain to pure MulticastDNS purposes. Search and routing
	// domains are a unicast DNS concept, they cannot be used to resolve single-label lookups via MulticastDNS.
	//
	// Added in version 229.
	Domains systemdconf.Value

	// Takes a boolean argument or "resolve". Controls Link-Local Multicast Name Resolution support (RFC 4795) on the local
	// host. If true, enables full LLMNR responder and resolver support. If false, disables both. If set to "resolve", only resolution
	// support is enabled, but responding is disabled. Note that systemd-networkd.service also maintains per-link LLMNR settings.
	// LLMNR will be enabled on a link only if the per-link and the global setting is on.
	//
	// Added in version 216.
	LLMNR systemdconf.Value

	// Takes a boolean argument or "resolve". Controls Multicast DNS support (RFC 6762) on the local host. If true, enables full
	// Multicast DNS responder and resolver support. If false, disables both. If set to "resolve", only resolution support is
	// enabled, but responding is disabled. Note that systemd-networkd.service also maintains per-link Multicast DNS settings.
	// Multicast DNS will be enabled on a link only if the per-link and the global setting is on.
	//
	// Added in version 234.
	MulticastDNS systemdconf.Value

	// Takes a boolean argument or "allow-downgrade".
	//
	// If set to true, all DNS lookups are DNSSEC-validated locally (excluding LLMNR and Multicast DNS). If the response to a lookup
	// request is detected to be invalid a lookup failure is returned to applications. Note that this mode requires a DNS server
	// that supports DNSSEC. If the DNS server does not properly support DNSSEC all validations will fail.
	//
	// If set to "allow-downgrade", DNSSEC validation is attempted, but if the server does not support DNSSEC properly, DNSSEC
	// mode is automatically disabled. Note that this mode makes DNSSEC validation vulnerable to "downgrade" attacks, where
	// an attacker might be able to trigger a downgrade to non-DNSSEC mode by synthesizing a DNS response that suggests DNSSEC
	// was not supported.
	//
	// If set to false, DNS lookups are not DNSSEC validated.
	//
	// Note that DNSSEC validation requires retrieval of additional DNS data, and thus results in a small DNS lookup time penalty.
	//
	// DNSSEC requires knowledge of "trust anchors" to prove data integrity. The trust anchor for the Internet root domain is
	// built into the resolver, additional trust anchors may be defined with dnssec-trust-anchors.d. Trust anchors may change
	// at regular intervals, and old trust anchors may be revoked. In such a case DNSSEC validation is not possible until new trust
	// anchors are configured locally or the resolver software package is updated with the new root trust anchor. In effect, when
	// the built-in trust anchor is revoked and DNSSEC= is true, all further lookups will fail, as it cannot be proved anymore whether
	// lookups are correctly signed, or validly unsigned. If DNSSEC= is set to "allow-downgrade" the resolver will automatically
	// turn off DNSSEC validation in such a case.
	//
	// Client programs looking up DNS data will be informed whether lookups could be verified using DNSSEC, or whether the returned
	// data could not be verified (either because the data was found unsigned in the DNS, or the DNS server did not support DNSSEC
	// or no appropriate trust anchors were known). In the latter case, it is assumed that client programs employ a secondary scheme
	// to validate the returned DNS data, should this be required.
	//
	// It is recommended to set DNSSEC= to true on systems where it is known that the DNS server supports DNSSEC correctly, and where
	// software or trust anchor updates happen regularly. On other systems it is recommended to set DNSSEC= to "allow-downgrade".
	//
	// In addition to this global DNSSEC setting systemd-networkd.service also maintains per-link DNSSEC settings. For system
	// DNS servers (see above), only the global DNSSEC setting is in effect. For per-link DNS servers the per-link setting is in
	// effect, unless it is unset in which case the global setting is used instead.
	//
	// Site-private DNS zones generally conflict with DNSSEC operation, unless a negative (if the private zone is not signed)
	// or positive (if the private zone is signed) trust anchor is configured for them. If "allow-downgrade" mode is selected,
	// it is attempted to detect site-private DNS zones using top-level domains (TLDs) that are not known by the DNS root server.
	// This logic does not work in all private zone setups.
	//
	// Defaults to "allow-downgrade".
	//
	// Added in version 229.
	DNSSEC systemdconf.Value

	// Takes a boolean argument or "opportunistic". If true all connections to the server will be encrypted. Note that this mode
	// requires a DNS server that supports DNS-over-TLS and has a valid certificate. If the hostname was specified in DNS= by using
	// the format "address#server_name" it is used to validate its certificate and also to enable Server Name Indication (SNI)
	// when opening a TLS connection. Otherwise the certificate is checked against the server's IP. If the DNS server does not
	// support DNS-over-TLS all DNS requests will fail.
	//
	// When set to "opportunistic" DNS request are attempted to send encrypted with DNS-over-TLS. If the DNS server does not support
	// TLS, DNS-over-TLS is disabled. Note that this mode makes DNS-over-TLS vulnerable to "downgrade" attacks, where an attacker
	// might be able to trigger a downgrade to non-encrypted mode by synthesizing a response that suggests DNS-over-TLS was not
	// supported. If set to false, DNS lookups are send over UDP.
	//
	// Note that DNS-over-TLS requires additional data to be send for setting up an encrypted connection, and thus results in
	// a small DNS look-up time penalty.
	//
	// Note that in "opportunistic" mode the resolver is not capable of authenticating the server, so it is vulnerable to "man-in-the-middle"
	// attacks.
	//
	// In addition to this global DNSOverTLS= setting systemd-networkd.service also maintains per-link DNSOverTLS= settings.
	// For system DNS servers (see above), only the global DNSOverTLS= setting is in effect. For per-link DNS servers the per-link
	// setting is in effect, unless it is unset in which case the global setting is used instead.
	//
	// Defaults to "no".
	//
	// Added in version 239.
	DNSOverTLS systemdconf.Value

	// Takes a boolean or "no-negative" as argument. If "yes" (the default), resolving a domain name which already got queried
	// earlier will return the previous result as long as it is still valid, and thus does not result in a new network request. Be
	// aware that turning off caching comes at a performance penalty, which is particularly high when DNSSEC is used. If "no-negative",
	// only positive answers are cached.
	//
	// Note that caching is turned off by default for host-local DNS servers. See CacheFromLocalhost= for details.
	//
	// Added in version 231.
	Cache systemdconf.Value

	// Takes a boolean as argument. If "no" (the default), and response cames from host-local IP address (such as 127.0.0.1 or
	// ::1), the result would not be cached in order to avoid potential duplicate local caching.
	//
	// Added in version 248.
	CacheFromLocalhost systemdconf.Value

	// Takes a boolean argument or one of "udp" and "tcp". If "udp", a DNS stub resolver will listen for UDP requests on addresses
	// 127.0.0.53 and 127.0.0.54, port 53. If "tcp", the stub will listen for TCP requests on the same addresses and port. If "yes"
	// (the default), the stub listens for both UDP and TCP requests. If "no", the stub listener is disabled.
	//
	// The DNS stub resolver on 127.0.0.53 provides the full feature set of the local resolver, which includes offering LLMNR/MulticastDNS
	// resolution. The DNS stub resolver on 127.0.0.54 provides a more limited resolver, that operates in "proxy" mode only,
	// i.e. it will pass most DNS messages relatively unmodified to the current upstream DNS servers and back, but not try to process
	// the messages locally, and hence does not validate DNSSEC, or offer up LLMNR/MulticastDNS. (It will translate to DNS-over-TLS
	// communication if needed however.)
	//
	// Note that the DNS stub listener is turned off implicitly when its listening address and port are already in use.
	//
	// Added in version 232.
	DNSStubListener systemdconf.Value

	// Takes an IPv4 or IPv6 address to listen on. The address may be optionally prefixed with a protocol name ("udp" or "tcp") separated
	// with ":". If the protocol is not specified, the service will listen on both UDP and TCP. It may be also optionally suffixed
	// by a numeric port number with separator ":". When an IPv6 address is specified with a port number, then the address must be
	// in the square brackets. If the port is not specified, then the service uses port 53. Note that this is independent of the primary
	// DNS stub configured with DNSStubListener=, and only configures additional sockets to listen on. This option can be specified
	// multiple times. If an empty string is assigned, then the all previous assignments are cleared. Defaults to unset.
	//
	// Examples:
	//
	//	DNSStubListenerExtra=192.168.10.10 DNSStubListenerExtra=2001:db8:0:f102::10 DNSStubListenerExtra=192.168.10.11:9953
	//	DNSStubListenerExtra=[2001:db8:0:f102::11]:9953 DNSStubListenerExtra=tcp:192.168.10.12 DNSStubListenerExtra=udp:2001:db8:0:f102::12
	//	DNSStubListenerExtra=tcp:192.168.10.13:9953 DNSStubListenerExtra=udp:[2001:db8:0:f102::13]:9953
	//
	// Added in version 247.
	DNSStubListenerExtra systemdconf.Value

	// Takes a boolean argument. If "yes" (the default), systemd-resolved will read /etc/hosts, and try to resolve hosts or address
	// by using the entries in the file before sending query to DNS servers.
	//
	// Added in version 240.
	ReadEtcHosts systemdconf.Value

	// Takes a boolean argument. When false (the default), systemd-resolved will not resolve A and AAAA queries for single-label
	// names over classic DNS. Note that such names may still be resolved if search domains are specified (see Domains= above),
	// or using other mechanisms, in particular via LLMNR or from /etc/hosts. When true, queries for single-label names will
	// be forwarded to global DNS servers even if no search domains are defined.
	//
	// This option is provided for compatibility with configurations where public DNS servers are not used. Forwarding single-label
	// names to servers not under your control is not standard-conformant, see IAB Statement, and may create a privacy and security
	// risk.
	//
	// Added in version 246.
	ResolveUnicastSingleLabel systemdconf.Value

	// Takes a duration value, which determines the length of time DNS resource records can be retained in the cache beyond their
	// Time To Live (TTL). This allows these records to be returned as stale records. By default, this value is set to zero, meaning
	// that DNS resource records are not stored in the cache after their TTL expires.
	//
	// This is useful when a DNS server failure occurs or becomes unreachable. In such cases, systemd-resolved continues to use
	// the stale records to answer DNS queries, particularly when no valid response can be obtained from the upstream DNS servers.
	// However, this does not apply to NXDOMAIN responses, as those are still perfectly valid responses. This feature enhances
	// resilience against DNS infrastructure failures and outages.
	//
	// systemd-resolved always attempts to reach the upstream DNS servers first, before providing the client application with
	// any stale data. If this feature is enabled, cache will not be flushed when changing servers.
	//
	// Added in version 254.
	StaleRetentionSec systemdconf.Value
}
