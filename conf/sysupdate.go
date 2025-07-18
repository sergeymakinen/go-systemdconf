// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// SysupdateFile represents sysupdate.d — Transfer Definition Files for Automatic Updates
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.d.html for details).
type SysupdateFile struct {
	systemdconf.File

	Transfer SysupdateTransferSection // general properties of this transfer
	Source   SysupdateSourceSection   // properties of the transfer source
	Target   SysupdateTargetSection   // properties of the transfer target
}

// SysupdateTransferSection represents general properties of this transfer
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.d.html#%5BTransfer%5D%20Section%20Options for details).
type SysupdateTransferSection struct {
	systemdconf.Section

	// Specifies the minimum version to require for this transfer to take place. If the source or target patterns in this transfer
	// definition match files older than this version they will be considered obsolete, and never be considered for the update
	// operation.
	//
	// Added in version 251.
	MinVersion systemdconf.Value

	// Takes one or more version strings to mark as "protected". Protected versions are never removed while making room for new,
	// updated versions. This is useful to ensure that the currently booted OS version (or auxiliary resources associated with
	// it) is not replaced/overwritten during updates, in order to avoid runtime file system corruptions.
	//
	// Like many of the settings in these configuration files this setting supports specifier expansion. It's particularly
	// useful to set this setting to one of the "%A", "%B" or "%w" specifiers to automatically refer to the current OS version of
	// the running system. See below for details on supported specifiers.
	//
	// Added in version 251.
	ProtectVersion systemdconf.Value

	// Takes a boolean, defaults to yes. Controls whether to cryptographically verify downloaded resources (specifically:
	// validate the GPG signatures for downloaded SHA256SUMS manifest files, via their detached signature files SHA256SUMS.gpg
	// in combination with the system keyring /usr/lib/systemd/import-pubring.gpg or /etc/systemd/import-pubring.gpg).
	//
	// This option is essential to provide integrity guarantees for downloaded resources and thus should be left enabled, outside
	// of test environments.
	//
	// Note that the downloaded payload files are unconditionally checked against the SHA256 hashes listed in the manifest.
	// This option only controls whether the signatures of these manifests are verified.
	//
	// This option only has an effect if the source resource type is selected as url-file or url-tar, as integrity and authentication
	// checking is only available for transfers from remote sources.
	//
	// Added in version 251.
	Verify systemdconf.Value

	// Optionally takes a human-presentable URL to a website containing a change-log of the resource being updated.
	//
	// This may be set multiple times in a single transfer definition. If set multiple times, the values are gathered into a list
	// of URLs. Adding a value of the empty string will clear the existing list of all values.
	//
	// This setting supports specifier expansion. See below for details on supported specifiers. This setting will also expand
	// the "@v" wildcard pattern. See above for details.
	//
	// Added in version 257.
	ChangeLog systemdconf.Value

	// Optionally takes a URL to an AppStream catalog XML file. This may be used by software centers (such as GNOME Software or KDE
	// Discover) to present rich metadata about the resources being updated. This includes display names, changelogs, icons,
	// and more. The specified catalog must include special metadata to be correctly associated with systemd-sysupdate by the
	// software centers.
	//
	// This setting supports specifier expansion. See below for details on supported specifiers.
	//
	// Added in version 257.
	AppStream systemdconf.Value

	// A space-separated list of sysupdate.features that this transfer belongs to, by name. This option may be specified more
	// than once, in which case the specified list of features is merged. If the empty string is assigned to this option, the list
	// is reset and all prior assignments will have no effect. For example: "Features=foo bar" specifies that the transfer belongs
	// to "foo.feature" and "bar.feature".
	//
	// If the list of features is empty, then this transfer is always used. If this transfer belongs to more than one feature, then
	// it will be used if any one of the listed features is enabled. A name that does not correspond to a defined feature will resolve
	// to an implicit feature that is always disabled.
	//
	// Added in version 257.
	Features systemdconf.Value

	// This is like Features=, except that all features listed here must be enabled for this transfer to be enabled. If both options
	// are specified, then they both apply: the transfer will be enabled only if all features specified here are enabled, and at
	// least one feature listed in Features= is enabled.
	//
	// Added in version 257.
	RequisiteFeatures systemdconf.Value
}

// SysupdateSourceSection represents properties of the transfer source
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.d.html#%5BSource%5D%20Section%20Options for details).
type SysupdateSourceSection struct {
	systemdconf.Section

	// Specifies the resource type of the source for the transfer. Takes one of url-file, url-tar, tar, regular-file, directory
	// or subvolume. For details about the resource types, see above. This option is mandatory.
	//
	// Note that only certain combinations of source and target resource types are supported, see above.
	//
	// Added in version 251.
	Type systemdconf.Value

	// Specifies where to find source versions of this resource.
	//
	// If the source type is selected as url-file or url-tar this must be a HTTP/HTTPS URL. The URL is suffixed with /SHA256SUMS
	// to acquire the manifest file, with /SHA256SUMS.gpg to acquire the detached signature file for it, and with the file names
	// listed in the manifest file in case an update is executed and a resource shall be downloaded.
	//
	// For all other source resource types this must be a local path in the file system, referring to a local directory to find the
	// versions of this resource in.
	//
	// Added in version 251.
	Path systemdconf.Value

	// Specifies one or more file name match patterns that select the subset of files that are update candidates as source for this
	// transfer. See above for details on match patterns.
	//
	// This option is mandatory. Any pattern listed must contain at least the "@v" wildcard, so that a version identifier may be
	// extracted from the filename. All other wildcards are optional.
	//
	// If the source type is regular-file or directory, the pattern may contain slash characters. In this case, it will match the
	// file or directory in corresponding subdirectory. For example "MatchPattern=foo_@v/bar.efi" will match "bar.efi"
	// in directory "foo_1".
	//
	// Added in version 251.
	MatchPattern systemdconf.Value
}

// SysupdateTargetSection represents properties of the transfer target
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.d.html#%5BTarget%5D%20Section%20Options for details).
type SysupdateTargetSection struct {
	systemdconf.Section

	// Specifies the resource type of the target for the transfer. Takes one of partition, regular-file, directory or subvolume.
	// For details about the resource types, see above. This option is mandatory.
	//
	// Note that only certain combinations of source and target resource types are supported, see above.
	//
	// Added in version 251.
	Type systemdconf.Value

	// Specifies a file system path where to look for already installed versions or place newly downloaded versions of this configured
	// resource. If Type= is set to partition, expects a path to a (whole) block device node, or the special string "auto" in which
	// case the block device which contains the root file system of the currently booted system is automatically determined and
	// used. If Type= is set to regular-file, directory or subvolume, must refer to a path in the local file system referencing
	// the directory to find or place the version files or directories under.
	//
	// Note that this mechanism cannot be used to create or remove partitions, in case Type= is set to partition. Partitions must
	// exist already, and a special partition label "_empty" is used to indicate empty partitions. To automatically generate
	// suitable partitions on first boot, use a tool such as systemd-repart.
	//
	// Added in version 251.
	Path systemdconf.Value

	// Specifies what anchor point Path= should be relative to. Takes one of root, esp, xbootldr, boot or directory. If unspecified,
	// defaults to root.
	//
	// If set to root, esp, xbootldr, the specified Path= will be resolved relative to the mount point of the corresponding partition,
	// as defined by the Boot Loader Specification.
	//
	// If set to boot, the specified Path= will be resolved relative to the mount point of the $BOOT partition (i.e. the ESP or XBOOTLDR),
	// as defined by the Boot Loader Specification.
	//
	// If set to explicit, the specified Path= will be resolved relative to the directory specified with --transfer-source=
	// when invoking systemd-sysupdate.
	//
	// The values esp, xbootldr, and boot are only supported when Type= is set to regular-file or directory.
	//
	// Added in version 254.
	PathRelativeTo systemdconf.Value

	// Specifies one or more file name or partition label match patterns that select the subset of files or partitions that are
	// update candidates as targets for this transfer. See above for details on match patterns.
	//
	// This option is mandatory. Any pattern listed must contain at least the "@v" wildcard, so that a version identifier may be
	// extracted from the filename. All other wildcards are optional.
	//
	// This pattern is both used for matching existing installed versions and for determining the name of new versions to install.
	// If multiple patterns are specified, the first specified is used for naming newly installed versions.
	//
	// If the target type is regular-file or directory, the pattern may contain slash characters. In this case, it will match the
	// file or directory in corresponding subdirectory. For example "MatchPattern=foo_@v/bar.efi" will match "bar.efi"
	// in directory "foo_1". Directories in the path will be created when file is installed. Empty directories will be removed
	// when file is removed.
	//
	// Added in version 251.
	MatchPattern systemdconf.Value

	// When the target Type= is chosen as partition, specifies the GPT partition type to look for. Only partitions of this type
	// are considered, all other partitions are ignored. If not specified, the GPT partition type linux-generic is used. Accepts
	// either a literal type UUID or a symbolic type identifier. For a list of supported type identifiers, see the Type= setting
	// in repart.d.
	//
	// Added in version 251.
	MatchPartitionType systemdconf.Value

	// When the target Type= is picked as partition, selects the GPT partition UUID and partition flags to use for the updated partition.
	// Expects a valid UUID string, a hexadecimal integer, or booleans, respectively. If not set, but the source match pattern
	// includes wildcards for these fields (i.e. "@u", "@f", "@a", or "@g"), the values from the patterns are used. If neither
	// configured with wildcards or these explicit settings, the values are left untouched. If both the overall PartitionFlags=
	// flags setting and the individual flag settings PartitionNoAuto= and PartitionGrowFileSystem= are used (or the wildcards
	// for them), then the latter override the former, i.e. the individual flag bit overrides the overall flags value. See Discoverable
	// Partitions Specification for details about these flags.
	//
	// Note that these settings are not used for matching, they only have effect on newly written partitions in case a transfer
	// takes place.
	//
	// Added in version 251.
	PartitionUUID systemdconf.Value

	// When the target Type= is picked as partition, selects the GPT partition UUID and partition flags to use for the updated partition.
	// Expects a valid UUID string, a hexadecimal integer, or booleans, respectively. If not set, but the source match pattern
	// includes wildcards for these fields (i.e. "@u", "@f", "@a", or "@g"), the values from the patterns are used. If neither
	// configured with wildcards or these explicit settings, the values are left untouched. If both the overall PartitionFlags=
	// flags setting and the individual flag settings PartitionNoAuto= and PartitionGrowFileSystem= are used (or the wildcards
	// for them), then the latter override the former, i.e. the individual flag bit overrides the overall flags value. See Discoverable
	// Partitions Specification for details about these flags.
	//
	// Note that these settings are not used for matching, they only have effect on newly written partitions in case a transfer
	// takes place.
	//
	// Added in version 251.
	PartitionFlags systemdconf.Value

	// When the target Type= is picked as partition, selects the GPT partition UUID and partition flags to use for the updated partition.
	// Expects a valid UUID string, a hexadecimal integer, or booleans, respectively. If not set, but the source match pattern
	// includes wildcards for these fields (i.e. "@u", "@f", "@a", or "@g"), the values from the patterns are used. If neither
	// configured with wildcards or these explicit settings, the values are left untouched. If both the overall PartitionFlags=
	// flags setting and the individual flag settings PartitionNoAuto= and PartitionGrowFileSystem= are used (or the wildcards
	// for them), then the latter override the former, i.e. the individual flag bit overrides the overall flags value. See Discoverable
	// Partitions Specification for details about these flags.
	//
	// Note that these settings are not used for matching, they only have effect on newly written partitions in case a transfer
	// takes place.
	//
	// Added in version 251.
	PartitionNoAuto systemdconf.Value

	// When the target Type= is picked as partition, selects the GPT partition UUID and partition flags to use for the updated partition.
	// Expects a valid UUID string, a hexadecimal integer, or booleans, respectively. If not set, but the source match pattern
	// includes wildcards for these fields (i.e. "@u", "@f", "@a", or "@g"), the values from the patterns are used. If neither
	// configured with wildcards or these explicit settings, the values are left untouched. If both the overall PartitionFlags=
	// flags setting and the individual flag settings PartitionNoAuto= and PartitionGrowFileSystem= are used (or the wildcards
	// for them), then the latter override the former, i.e. the individual flag bit overrides the overall flags value. See Discoverable
	// Partitions Specification for details about these flags.
	//
	// Note that these settings are not used for matching, they only have effect on newly written partitions in case a transfer
	// takes place.
	//
	// Added in version 251.
	PartitionGrowFileSystem systemdconf.Value

	// Controls whether to mark the resulting file, subvolume or partition read-only. If the target type is partition this controls
	// the ReadOnly partition flag, as per Discoverable Partitions Specification, similar to the PartitionNoAuto= and PartitionGrowFileSystem=
	// flags described above. If the target type is regular-file, the writable bit is removed from the access mode. If the target
	// type is subvolume, the subvolume will be marked read-only as a whole. Finally, if the target Type= is selected as directory,
	// the "immutable" file attribute is set, see chattr for details.
	//
	// Added in version 251.
	ReadOnly systemdconf.Value

	// The UNIX file access mode to use for newly created files in case the target resource type is picked as regular-file. Expects
	// an octal integer, in typical UNIX fashion. If not set, but the source match pattern includes a wildcard for this field (i.e.
	// "@t"), the value from the pattern is used.
	//
	// Note that this setting is not used for matching, it only has an effect on newly written files when a transfer takes place.
	//
	// Added in version 251.
	Mode systemdconf.Value

	// These options take positive, decimal integers, and control the number of attempts done and left for this file. These settings
	// are useful for managing kernel images, following the scheme defined in Automatic Boot Assessment, and only have an effect
	// if the target pattern includes the "@d" or "@l" wildcards.
	//
	// Added in version 251.
	TriesDone systemdconf.Value

	// These options take positive, decimal integers, and control the number of attempts done and left for this file. These settings
	// are useful for managing kernel images, following the scheme defined in Automatic Boot Assessment, and only have an effect
	// if the target pattern includes the "@d" or "@l" wildcards.
	//
	// Added in version 251.
	TriesLeft systemdconf.Value

	// Takes a decimal integer equal to or greater than 2. This configures how many concurrent versions of the resource to keep.
	// Whenever a new update is initiated it is made sure that no more than the number of versions specified here minus one exist
	// in the target. Any excess versions are deleted (in case the target Type= of regular-file, directory, subvolume is used)
	// or emptied (in case the target Type= of partition is used; emptying in this case simply means to set the partition label to
	// the special string "_empty"; note that no partitions are actually removed). After an update is completed the number of
	// concurrent versions of the target resources is equal to or below the number specified here.
	//
	// Note that this setting may be set differently for each transfer. However, it generally is advisable to keep this setting
	// the same for all transfers, since otherwise incomplete combinations of files or partitions will be left installed.
	//
	// If the target Type= is selected as partition, the number of concurrent versions to keep is additionally restricted by the
	// number of partition slots of the right type in the partition table. I.e. if there are only 2 partition slots for the selected
	// partition type, setting this value larger than 2 is without effect, since no more than 2 concurrent versions could be stored
	// in the image anyway.
	//
	// Added in version 251.
	InstancesMax systemdconf.Value

	// Takes a boolean argument. If this option is enabled (which is the default) before initiating an update, all left-over,
	// incomplete updates from a previous attempt are removed from the target directory. This only has an effect if the target
	// resource Type= is selected as regular-file, directory or subvolume.
	//
	// Added in version 251.
	RemoveTemporary systemdconf.Value

	// Takes a symlink name as argument. If this option is used, as the last step of the update a symlink under the specified name
	// is created/updated pointing to the completed update. This is useful in to provide a stable name always pointing to the newest
	// version of the resource. This is only supported if the target resource Type= is selected as regular-file, directory or
	// subvolume.
	//
	// Added in version 251.
	CurrentSymlink systemdconf.Value
}
