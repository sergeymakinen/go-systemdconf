// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// UkifyFile represents ukify — Combine components into a signed Unified Kernel Image for UEFI systems
// (see https://www.freedesktop.org/software/systemd/man/ukify.html for details).
type UkifyFile struct {
	systemdconf.File

	UKI UkifyUKISection // [UKI] section
}

// UkifyUKISection represents [UKI] section
// (see https://www.freedesktop.org/software/systemd/man/ukify.html#%5BUKI%5D%20section for details).
type UkifyUKISection struct {
	systemdconf.Section

	// A path to the kernel binary.
	//
	// Added in version 254.
	Linux systemdconf.Value

	// The os-release description (the ".osrel" section). The argument may be a literal string, or "@" followed by a path name.
	// If not specified, the os-release file will be picked up from the host system.
	//
	// Added in version 253.
	OSRelease systemdconf.Value

	// The kernel command line (the ".cmdline" section). The argument may be a literal string, or "@" followed by a path name. If
	// not specified, no command line will be embedded.
	//
	// Added in version 253.
	Cmdline systemdconf.Value

	// Zero or more initrd paths. In the configuration file, items are separated by whitespace. The initrds are combined in the
	// order of specification, with the initrds specified in the config file first.
	//
	// Added in version 254.
	Initrd systemdconf.Value

	// Path to initrd containing microcode updates. If not specified, the section will not be present.
	//
	// Added in version 256.
	Microcode systemdconf.Value

	// A picture to display during boot (the ".splash" section). The argument is a path to a BMP file. If not specified, the section
	// will not be present.
	//
	// Added in version 253.
	Splash systemdconf.Value

	// The devicetree description (the ".dtb" section). The argument is a path to a compiled binary DeviceTree file. If not specified,
	// the section will not be present.
	//
	// Added in version 253.
	DeviceTree systemdconf.Value

	// Zero or more automatically selectable DeviceTree files. In the configuration file, items are separated by whitespace.
	// Each DeviceTree will be in a separate ".dtbauto" section.
	//
	// Added in version 257.
	DeviceTreeAuto systemdconf.Value

	// The hardware ID device table (the ".hwids" section). The argument is a path to a directory with JSON HWID device description
	// files. Each file needs to contain a single JSON object with a "name", "compatible" and "hwids" keys. The "name" and "compatible"
	// keys must have string values and the "hwids" key must have a list of strings as value, where the strings must be valid UUIDs
	// that represent CHIDs/HWIDs. Example:
	//
	//	{ "name": "Example Laptop 16 Gen 7", "compatible": "example,laptop-16-g7", "hwids": [ "5dc05bf4-01f6-4089-b464-a08c47ea9295",
	//	"3e3f8f3c-2003-46f2-811c-85554f7d5952" ] }
	//
	// Here "Example Laptop 16 Gen 7" is the device "name" (as defined by the manufacturer), "example,laptop-16-g7" is the "compatible"
	// (as defined by the kernel) and "hwids" is an array of CHIDs/HWIDs (extracted i.e. from fwupdtool hwids output). If not specified,
	// the section will not be present. It is recommended to specify this parameter if automatically selectable DeviceTrees
	// are to be used.
	//
	// Added in version 257.
	HWIDs systemdconf.Value

	// Specify the kernel version (as in uname -r, the ".uname" section). If not specified, an attempt will be made to extract the
	// version string from the kernel image. It is recommended to pass this explicitly if known, because the extraction is based
	// on heuristics and not very reliable. If not specified and extraction fails, the section will not be present.
	//
	// Added in version 253.
	Uname systemdconf.Value

	// SBAT metadata associated with the UKI or addon. SBAT policies are useful to revoke whole groups of UKIs or addons with a single,
	// static policy update that does not take space in DBX/MOKX. If not specified manually, a default metadata entry consisting
	// of
	//
	//	uki,1,UKI,uki,1,https://uapi-group.org/specifications/specs/unified_kernel_image/
	//
	// for UKIs and
	//
	//	uki-addon,1,UKI Addon,addon,1,https://www.freedesktop.org/software/systemd/man/latest/systemd-stub.html
	//
	// for addons will be used, to ensure it is always possible to revoke them. For more information on SBAT see Shim documentation.
	//
	// Added in version 254.
	SBAT systemdconf.Value

	// A path to a public key to embed in the ".pcrpkey" section. If not specified, and there's exactly one PCRPublicKey=/--pcr-public-key=
	// argument, that key will be used. Otherwise, the section will not be present.
	//
	// Added in version 253.
	PCRPKey systemdconf.Value

	// A path to a UKI profile to place in an ".profile" section. This option is useful for creating multi-profile UKIs, and is typically
	// used in combination with --join-profile=, to extend the specified UKI with an additional profile.
	//
	// Added in version 257.
	Profile systemdconf.Value

	// A comma or space-separated list of PCR banks to sign a policy for. If not present, all known banks will be used ("sha1", "sha256",
	// "sha384", "sha512"), which will fail if not supported by the system.
	//
	// Added in version 253.
	PCRBanks systemdconf.Value

	// Whether to use "sbsign", "pesign", or "systemd-sbsign". Depending on this choice, different parameters are required
	// in order to sign an image. Defaults to "sbsign".
	//
	// Added in version 254.
	SecureBootSigningTool systemdconf.Value

	// A path to a private key to use for signing of the resulting binary. If the SigningEngine=/--signing-engine= or SigningProvider=/--signing-provider=
	// option is used, this may also be an engine or provider specific designation. This option is required by SecureBootSigningTool=sbsign/--signtool=sbsign.
	//
	// Added in version 253.
	SecureBootPrivateKey systemdconf.Value

	// A path to a certificate to use for signing of the resulting binary. If the SigningEngine=/--signing-engine= or SigningProvider=/--signing-provider=
	// option is used, this may also be an engine or provider specific designation. This option is required by SecureBootSigningTool=sbsign/--signtool=sbsign.
	//
	// Added in version 253.
	SecureBootCertificate systemdconf.Value

	// A path to a nss certificate database directory to use for signing of the resulting binary. Takes effect when SecureBootSigningTool=pesign/--signtool=pesign
	// is used. Defaults to /etc/pki/pesign.
	//
	// Added in version 254.
	SecureBootCertificateDir systemdconf.Value

	// The name of the nss certificate database entry to use for signing of the resulting binary. This option is required by SecureBootSigningTool=pesign/--signtool=pesign.
	//
	// Added in version 254.
	SecureBootCertificateName systemdconf.Value

	// Period of validity (in days) for a certificate created by genkey. Defaults to 3650, i.e. 10 years.
	//
	// Added in version 254.
	SecureBootCertificateValidity systemdconf.Value

	// An OpenSSL engine to be used for signing the resulting binary and PCR measurements, see openssl-engine.
	//
	// Added in version 253.
	SigningEngine systemdconf.Value

	// An OpenSSL provider to be used for signing the resulting binary and PCR measurements, see provider. This option can only
	// be used when systemd-sbsign is used as the signing tool.
	//
	// Added in version 257.
	SigningProvider systemdconf.Value

	// An OpenSSL provider to be used for loading the certificate used to sign the resulting binary and PCR measurements, see provider.
	// This option can only be used when systemd-sbsign is used as the signing tool.
	//
	// Added in version 257.
	CertificateProvider systemdconf.Value

	// Override the detection of whether to sign the Linux binary itself before it is embedded in the combined image. If not specified,
	// it will be signed if a SecureBoot signing key is provided via the SecureBootPrivateKey=/--secureboot-private-key=
	// option and the binary has not already been signed. If SignKernel=/--sign-kernel is true, and the binary has already been
	// signed, the signature will be appended anyway.
	//
	// Added in version 253.
	SignKernel systemdconf.Value
}

// PCRSignatureSection represents [PCRSignature:NAME] section
// (see https://www.freedesktop.org/software/systemd/man/ukify.html#%5BPCRSignature:NAME%5D%20section for details).
type PCRSignatureSection struct {
	systemdconf.Section

	// A private key to use for signing PCR policies. On the command line, this option may be specified more than once, in which case
	// multiple signatures will be made.
	//
	// Added in version 253.
	PCRPrivateKey systemdconf.Value

	// A public key to use for signing PCR policies.
	//
	// On the command line, this option may be specified more than once, similarly to the --pcr-private-key= option. If not present,
	// the public keys will be extracted from the private keys. On the command line, if present, this option must be specified the
	// same number of times as the --pcr-private-key= option.
	//
	// Added in version 253.
	PCRPublicKey systemdconf.Value

	// A comma or space-separated list of colon-separated phase paths to sign a policy for. Each set of boot phase paths will be
	// signed with the corresponding private key. If not present, the default of systemd-measure will be used.
	//
	// On the command line, when this argument is present, it must appear the same number of times as the --pcr-private-key= option.
	//
	// Added in version 253.
	Phases systemdconf.Value
}
