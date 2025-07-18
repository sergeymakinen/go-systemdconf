// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// RepartFile represents repart.d — Partition Definition Files for Automatic Boot-Time Repartitioning
// (see https://www.freedesktop.org/software/systemd/man/repart.d.html for details).
type RepartFile struct {
	systemdconf.File

	Partition RepartPartitionSection // basic properties of partitions of block devices of the local system
}

// RepartPartitionSection represents basic properties of partitions of block devices of the local system
// (see https://www.freedesktop.org/software/systemd/man/repart.d.html#%5BPartition%5D%20Section%20Options for details).
type RepartPartitionSection struct {
	systemdconf.Section

	// The GPT partition type UUID to match. This may be a GPT partition type UUID such as 4f68bce3-e8cd-4db1-96e7-fbcaf984b709,
	// or an identifier.
	//
	// The supported identifiers are:
	//
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	|        IDENTIFIER         |                                   EXPLANATION                                    |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//	| esp                       | EFI System Partition                                                             |
	//	| xbootldr                  | Extended Boot Loader Partition                                                   |
	//	| swap                      | Swap partition                                                                   |
	//	| home                      | Home (/home/) partition                                                          |
	//	| srv                       | Server data (/srv/) partition                                                    |
	//	| var                       | Variable data (/var/) partition                                                  |
	//	| tmp                       | Temporary data (/var/tmp/) partition                                             |
	//	| linux-generic             | Generic Linux file system partition                                              |
	//	| root                      | Root file system partition type appropriate for the local architecture (an       |
	//	|                           | alias for an architecture root file system partition type listed below, e.g.     |
	//	|                           | root-x86-64)                                                                     |
	//	| root-verity               | Verity data for the root file system partition for the local architecture        |
	//	| root-verity-sig           | Verity signature data for the root file system partition for the local           |
	//	|                           | architecture                                                                     |
	//	| root-secondary            | Root file system partition of the secondary architecture of the local            |
	//	|                           | architecture (usually the matching 32-bit architecture for the local 64-bit      |
	//	|                           | architecture)                                                                    |
	//	| root-secondary-verity     | Verity data for the root file system partition of the secondary architecture     |
	//	| root-secondary-verity-sig | Verity signature data for the root file system partition of the secondary        |
	//	|                           | architecture                                                                     |
	//	| root-{arch}               | Root file system partition of the given architecture (such as root-x86-64 or     |
	//	|                           | root-riscv64)                                                                    |
	//	| root-{arch}-verity        | Verity data for the root file system partition of the given architecture         |
	//	| root-{arch}-verity-sig    | Verity signature data for the root file system partition of the given            |
	//	|                           | architecture                                                                     |
	//	| usr                       | /usr/ file system partition type appropriate for the local architecture (an      |
	//	|                           | alias for an architecture /usr/ file system partition type listed below, e.g.    |
	//	|                           | usr-x86-64)                                                                      |
	//	| usr-verity                | Verity data for the /usr/ file system partition for the local architecture       |
	//	| usr-verity-sig            | Verity signature data for the /usr/ file system partition for the local          |
	//	|                           | architecture                                                                     |
	//	| usr-secondary             | /usr/ file system partition of the secondary architecture of the local           |
	//	|                           | architecture (usually the matching 32-bit architecture for the local 64-bit      |
	//	|                           | architecture)                                                                    |
	//	| usr-secondary-verity      | Verity data for the /usr/ file system partition of the secondary architecture    |
	//	| usr-secondary-verity-sig  | Verity signature data for the /usr/ file system partition of the secondary       |
	//	|                           | architecture                                                                     |
	//	| usr-{arch}                | /usr/ file system partition of the given architecture                            |
	//	| usr-{arch}-verity         | Verity data for the /usr/ file system partition of the given architecture        |
	//	| usr-{arch}-verity-sig     | Verity signature data for the /usr/ file system partition of the given           |
	//	|                           | architecture                                                                     |
	//	+---------------------------+----------------------------------------------------------------------------------+
	//
	// Architecture specific partition types can use one of these architecture identifiers: alpha, arc, arm (32-bit), arm64
	// (64-bit, aka aarch64), ia64, loongarch64, mips-le, mips64-le, parisc, ppc, ppc64, ppc64-le, riscv32, riscv64, s390,
	// s390x, tilegx, x86 (32-bit, aka i386) and x86-64 (64-bit, aka amd64).
	//
	// Most of the partition type UUIDs listed above are defined in the Discoverable Partitions Specification.
	//
	// Added in version 245.
	Type systemdconf.Value

	// The textual label to assign to the partition if none is assigned yet. Note that this setting is not used for matching. It is
	// also not used when a label is already set for an existing partition. It is thus only used when a partition is newly created
	// or when an existing one had a no label set (that is: an empty label). If not specified, a label derived from the partition type
	// is automatically used. Simple specifier expansion is supported, see below.
	//
	// Added in version 245.
	Label systemdconf.Value

	// The UUID to assign to the partition if none is assigned yet. Note that this setting is not used for matching. It is also not
	// used when a UUID is already set for an existing partition. It is thus only used when a partition is newly created or when an
	// existing one had a all-zero UUID set. If set to "null", the UUID is set to all zeroes. If not specified a UUID derived from the
	// partition type is automatically used.
	//
	// Added in version 246.
	UUID systemdconf.Value

	// A numeric priority to assign to this partition, in the range -2147483648…2147483647, with smaller values indicating
	// higher priority, and higher values indicating smaller priority. This priority is used in case the configured size constraints
	// on the defined partitions do not permit fitting all partitions onto the available disk space. If the partitions do not fit,
	// the highest numeric partition priority of all defined partitions is determined, and all defined partitions with this
	// priority are removed from the list of new partitions to create (which may be multiple, if the same priority is used for multiple
	// partitions). The fitting algorithm is then tried again. If the partitions still do not fit, the now highest numeric partition
	// priority is determined, and the matching partitions removed too, and so on. Partitions of a priority of 0 or lower are never
	// removed. If all partitions with a priority above 0 are removed and the partitions still do not fit on the device the operation
	// fails. Note that this priority has no effect on ordering partitions, for that use the alphabetical order of the filenames
	// of the partition definition files. Defaults to 0.
	//
	// Added in version 245.
	Priority systemdconf.Value

	// A numeric weight to assign to this partition in the range 0…1000000. Available disk space is assigned the defined partitions
	// according to their relative weights (subject to the size constraints configured with SizeMinBytes=, SizeMaxBytes=),
	// so that a partition with weight 2000 gets double the space as one with weight 1000, and a partition with weight 333 a third
	// of that. Defaults to 1000.
	//
	// The Weight= setting is used to distribute available disk space in an "elastic" fashion, based on the disk size and existing
	// partitions. If a partition shall have a fixed size use both SizeMinBytes= and SizeMaxBytes= with the same value in order
	// to fixate the size to one value, in which case the weight has no effect.
	//
	// Added in version 245.
	Weight systemdconf.Value

	// Similar to Weight=, but sets a weight for the free space after the partition (the "padding"). When distributing available
	// space the weights of all partitions and all defined padding is summed, and then each partition and padding gets the fraction
	// defined by its weight. Defaults to 0, i.e. by default no padding is applied.
	//
	// Padding is useful if empty space shall be left for later additions or a safety margin at the end of the device or between partitions.
	//
	// Added in version 245.
	PaddingWeight systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes. Takes the usual K, M, G, T, … suffixes (to the base of 1024). If SizeMinBytes=
	// is specified the partition is created at or grown to at least the specified size. If SizeMaxBytes= is specified the partition
	// is created at or grown to at most the specified size. The precise size is determined through the weight value configured
	// with Weight=, see above. When SizeMinBytes= is set equal to SizeMaxBytes= the configured weight has no effect as the partition
	// is explicitly sized to the specified fixed value. Note that partitions are never created smaller than 4096 bytes, and since
	// partitions are never shrunk the previous size of the partition (in case the partition already exists) is also enforced
	// as lower bound for the new size. The values should be specified as multiples of 4096 bytes, and are rounded upwards (in case
	// of SizeMinBytes=) or downwards (in case of SizeMaxBytes=) otherwise. If the backing device does not provide enough space
	// to fulfill the constraints placing the partition will fail. For partitions that shall be created, depending on the setting
	// of Priority= (see above) the partition might be dropped and the placing algorithm restarted. By default, a minimum size
	// constraint of 10M and no maximum size constraint is set.
	//
	// Added in version 245.
	SizeMinBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes. Takes the usual K, M, G, T, … suffixes (to the base of 1024). If SizeMinBytes=
	// is specified the partition is created at or grown to at least the specified size. If SizeMaxBytes= is specified the partition
	// is created at or grown to at most the specified size. The precise size is determined through the weight value configured
	// with Weight=, see above. When SizeMinBytes= is set equal to SizeMaxBytes= the configured weight has no effect as the partition
	// is explicitly sized to the specified fixed value. Note that partitions are never created smaller than 4096 bytes, and since
	// partitions are never shrunk the previous size of the partition (in case the partition already exists) is also enforced
	// as lower bound for the new size. The values should be specified as multiples of 4096 bytes, and are rounded upwards (in case
	// of SizeMinBytes=) or downwards (in case of SizeMaxBytes=) otherwise. If the backing device does not provide enough space
	// to fulfill the constraints placing the partition will fail. For partitions that shall be created, depending on the setting
	// of Priority= (see above) the partition might be dropped and the placing algorithm restarted. By default, a minimum size
	// constraint of 10M and no maximum size constraint is set.
	//
	// Added in version 245.
	SizeMaxBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes for the free space after the partition (the "padding"). Semantics
	// are similar to SizeMinBytes= and SizeMaxBytes=, except that unlike partition sizes free space can be shrunk and can be
	// as small as zero. By default, no size constraints on padding are set, so that only PaddingWeight= determines the size of
	// the padding applied.
	//
	// Added in version 245.
	PaddingMinBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes for the free space after the partition (the "padding"). Semantics
	// are similar to SizeMinBytes= and SizeMaxBytes=, except that unlike partition sizes free space can be shrunk and can be
	// as small as zero. By default, no size constraints on padding are set, so that only PaddingWeight= determines the size of
	// the padding applied.
	//
	// Added in version 245.
	PaddingMaxBytes systemdconf.Value

	// Takes a path to a regular file, block device node, char device node or directory, or the special value "auto". If specified
	// and the partition is newly created, the data from the specified path is written to the newly created partition, on the block
	// level. If a directory is specified, the backing block device of the file system the directory is on is determined, and the
	// data read directly from that. This option is useful to efficiently replicate existing file systems onto new partitions
	// on the block level — for example to build a simple OS installer or an OS image builder. Specify /dev/urandom as value to initialize
	// a partition with random data.
	//
	// If the special value "auto" is specified, the source to copy from is automatically picked up from the running system (or
	// the image specified with --image= — if used). A partition that matches both the configured partition type (as declared
	// with Type= described above), and the currently mounted directory appropriate for that partition type is determined.
	// For example, if the partition type is set to "root" the partition backing the root directory (/) is used as source to copy
	// from — if its partition type is set to "root" as well. If the declared type is "usr" the partition backing /usr/ is used as source
	// to copy blocks from — if its partition type is set to "usr" too. The logic is capable of automatically tracking down the backing
	// partitions for encrypted and Verity-enabled volumes. "CopyBlocks=auto" is useful for implementing "self-replicating"
	// systems, i.e. systems that are their own installer.
	//
	// The file specified here must have a size that is a multiple of the basic block size 512 and not be empty. If this option is used,
	// the size allocation algorithm is slightly altered: the partition is created at least as big as required to fit the data in,
	// i.e. the data size is an additional minimum size value taken into consideration for the allocation algorithm, similar
	// to and in addition to the SizeMin= value configured above.
	//
	// This option has no effect if the partition it is declared for already exists, i.e. existing data is never overwritten. Note
	// that the data is copied in before the partition table is updated, i.e. before the partition actually is persistently created.
	// This provides robustness: it is guaranteed that the partition either does not exist or exists fully populated; it is not
	// possible that the partition exists but is not or only partially populated.
	//
	// This option cannot be combined with Format= or CopyFiles=.
	//
	// Added in version 246.
	CopyBlocks systemdconf.Value

	// Takes a file system name, such as "ext4", "btrfs", "xfs", "vfat", "erofs", "squashfs" or the special value "swap". If specified
	// and the partition is newly created it is formatted with the specified file system (or as swap device). The file system UUID
	// and label are automatically derived from the partition UUID and label. If this option is used, the size allocation algorithm
	// is slightly altered: the partition is created at least as big as required for the minimal file system of the specified type
	// (or 4KiB if the minimal size is not known).
	//
	// This option has no effect if the partition already exists.
	//
	// Similarly to the behaviour of CopyBlocks=, the file system is formatted before the partition is created, ensuring that
	// the partition only ever exists with a fully initialized file system.
	//
	// This option cannot be combined with CopyBlocks=.
	//
	// Added in version 247.
	Format systemdconf.Value

	// Takes a pair of colon separated absolute file system paths. The first path refers to a source file or directory on the host,
	// the second path refers to a target in the file system of the newly created partition and formatted file system. This setting
	// may be used to copy files or directories from the host into the file system that is created due to the Format= option. If CopyFiles=
	// is used without Format= specified explicitly, "Format=" with a suitable default is implied (currently "vfat" for "ESP"
	// and "XBOOTLDR" partitions, and "ext4" otherwise, but this may change in the future). This option may be used multiple times
	// to copy multiple files or directories from host into the newly formatted file system. The colon and second path may be omitted
	// in which case the source path is also used as the target path (relative to the root of the newly created file system). If the
	// source path refers to a directory it is copied recursively.
	//
	// This option has no effect if the partition already exists: it cannot be used to copy additional files into an existing partition,
	// it may only be used to populate a file system created anew.
	//
	// The copy operation is executed before the file system is registered in the partition table, thus ensuring that a file system
	// populated this way only ever exists fully initialized.
	//
	// Note that CopyFiles= will skip copying files that are not supported by the target filesystem (e.g symlinks, fifos, sockets
	// and devices on vfat). When an unsupported file type is encountered, systemd-repart will skip copying this file and write
	// a log message about it.
	//
	// Note that systemd-repart does not change the UIDs/GIDs of any copied files and directories. When running systemd-repart
	// as an unprivileged user to build an image of files and directories owned by the same user, you can run systemd-repart in a
	// user namespace with the current user mapped to the root user to make sure the files and directories in the image are owned
	// by the root user.
	//
	// Note that when populating XFS filesystems with systemd-repart and loop devices are not available, populating XFS filesystems
	// with files containing spaces, tabs or newlines might fail on old versions of mkfs.xfs due to limitations of its protofile
	// format.
	//
	// Note that when populating XFS filesystems with systemd-repart and loop devices are not available, extended attributes
	// will not be copied into generated XFS filesystems due to limitations mkfs.xfs's protofile format.
	//
	// This option cannot be combined with CopyBlocks=.
	//
	// When systemd-repart is invoked with the --copy-source= command line switch the file paths are taken relative to the specified
	// directory. If --copy-source= is not used, but the --image= or --root= switches are used, the source paths are taken relative
	// to the specified root directory or disk image root.
	//
	// Added in version 247.
	CopyFiles systemdconf.Value

	// Takes one or more absolute paths, separated by whitespace, each referring to a source file or directory on the host. This
	// setting may be used to exclude files or directories from the host from being copied into the file system when CopyFiles=
	// is used. This option may be used multiple times to exclude multiple files or directories from host from being copied into
	// the newly formatted file system.
	//
	// If the path is a directory and ends with "/", only the directory's contents are excluded but not the directory itself. If
	// the path is a directory and does not end with "/", both the directory and its contents are excluded.
	//
	// ExcludeFilesTarget= is like ExcludeFiles= except that instead of excluding the path on the host from being copied into
	// the partition, it excludes any files and directories from being copied into the given path in the partition.
	//
	// When systemd-repart is invoked with the --image= or --root= command line switches the paths specified are taken relative
	// to the specified root directory or disk image root.
	//
	// Added in version 254.
	ExcludeFiles systemdconf.Value

	// Takes one or more absolute paths, separated by whitespace, each referring to a source file or directory on the host. This
	// setting may be used to exclude files or directories from the host from being copied into the file system when CopyFiles=
	// is used. This option may be used multiple times to exclude multiple files or directories from host from being copied into
	// the newly formatted file system.
	//
	// If the path is a directory and ends with "/", only the directory's contents are excluded but not the directory itself. If
	// the path is a directory and does not end with "/", both the directory and its contents are excluded.
	//
	// ExcludeFilesTarget= is like ExcludeFiles= except that instead of excluding the path on the host from being copied into
	// the partition, it excludes any files and directories from being copied into the given path in the partition.
	//
	// When systemd-repart is invoked with the --image= or --root= command line switches the paths specified are taken relative
	// to the specified root directory or disk image root.
	//
	// Added in version 254.
	ExcludeFilesTarget systemdconf.Value

	// Takes one or more absolute paths, separated by whitespace, each declaring a directory to create within the new file system.
	// Behaviour is similar to CopyFiles=, but instead of copying in a set of files this just creates the specified directories
	// with the default mode of 0755 owned by the root user and group, plus all their parent directories (with the same ownership
	// and access mode). To configure directories with different ownership or access mode, use CopyFiles= and specify a source
	// tree to copy containing appropriately owned/configured directories. This option may be used more than once to create
	// multiple directories. When CopyFiles= and MakeDirectories= are used together the former is applied first. If a directory
	// listed already exists no operation is executed (in particular, the ownership/access mode of the directories is left as
	// is).
	//
	// The primary use case for this option is to create a minimal set of directories that may be mounted over by other partitions
	// contained in the same disk image. For example, a disk image where the root file system is formatted at first boot might want
	// to automatically pre-create /usr/ in it this way, so that the "usr" partition may over-mount it.
	//
	// Consider using systemd-tmpfiles with its --image= option to pre-create other, more complex directory hierarchies (as
	// well as other inodes) with fine-grained control of ownership, access modes and other file attributes.
	//
	// Added in version 249.
	MakeDirectories systemdconf.Value

	// Takes one or more arguments, separated by whitespace, each declaring a symlink to create within the new file system. Each
	// argument is a pair of symlink source and target paths, separated by a colon. This option may be used more than once to create
	// multiple symlinks. When CopyFiles= and MakeSymlinks= are used together the former is applied first.
	//
	// The primary use case for this option is to create symlinks that need to exist before systemd-tmpfiles is executed. For example,
	// when using systemd-confext, this setting can be used to create symlinks in /var/lib/extensions.mutable to redirect
	// writes to mutable confexts to a custom location.
	//
	// Consider using systemd-tmpfiles with its --image= option to pre-create other symlinks (as well as other inodes) with
	// fine-grained control of ownership, access modes and other file attributes.
	//
	// Added in version 257.
	MakeSymlinks systemdconf.Value

	// Takes one or more absolute paths, separated by whitespace, each declaring a directory that should be a subvolume within
	// the new file system. Each path may optionally be followed by a colon and a list of comma-separated subvolume flags. The following
	// flags are understood:
	//
	//	+------+--------------------------------+
	//	| FLAG |            PURPOSE             |
	//	+------+--------------------------------+
	//	| "ro" | Make this subvolume read-only. |
	//	+------+--------------------------------+
	//
	// Note that this option does not create the directories themselves, that can be configured with MakeDirectories= and CopyFiles=.
	//
	// Note that this option only takes effect if the target filesystem supports subvolumes, such as btrfs.
	//
	// Note that this option is only supported in combination with --offline=yes since btrfs-progs 6.11 or newer.
	//
	// Added in version 255.
	Subvolumes systemdconf.Value

	// Takes an absolute path specifying the default subvolume within the new filesystem. Note that this setting does not create
	// the subvolume itself, that can be configured with Subvolumes=.
	//
	// Note that this option only takes effect if the target filesystem supports subvolumes, such as btrfs.
	//
	// Note that this option is only supported in combination with --offline=yes since btrfs-progs 6.11 or newer.
	//
	// Added in version 256.
	DefaultSubvolume systemdconf.Value

	// Takes one of "off", "key-file", "tpm2" and "key-file+tpm2" (alternatively, also accepts a boolean value, which is mapped
	// to "off" when false, and "key-file" when true). Defaults to "off". If not "off" the partition will be formatted with a LUKS2
	// superblock, before the blocks configured with CopyBlocks= are copied in or the file system configured with Format= is
	// created.
	//
	// The LUKS2 UUID is automatically derived from the partition UUID in a stable fashion. If "key-file" or "key-file+tpm2"
	// is used, a key is added to the LUKS2 superblock, configurable with the --key-file= option to systemd-repart. If "tpm2"
	// or "key-file+tpm2" is used, a key is added to the LUKS2 superblock that is enrolled to the local TPM2 chip, as configured
	// with the --tpm2-device= and --tpm2-pcrs= options to systemd-repart.
	//
	// When used this slightly alters the size allocation logic as the implicit, minimal size limits of Format= and CopyBlocks=
	// are increased by the space necessary for the LUKS2 superblock (see above).
	//
	// This option has no effect if the partition already exists.
	//
	// Added in version 247.
	Encrypt systemdconf.Value

	// Takes one of "off", "data", "hash" or "signature". Defaults to "off". If set to "off" or "data", the partition is populated
	// with content as specified by CopyBlocks= or CopyFiles=. If set to "hash", the partition will be populated with verity hashes
	// from the matching verity data partition. If set to "signature", the partition will be populated with a JSON object containing
	// a signature of the verity root hash of the matching verity hash partition.
	//
	// A matching verity partition is a partition with the same verity match key (as configured with VerityMatchKey=).
	//
	// If not explicitly configured, the data partition's UUID will be set to the first 128 bits of the verity root hash. Similarly,
	// if not configured, the hash partition's UUID will be set to the final 128 bits of the verity root hash. The verity root hash
	// itself will be included in the output of systemd-repart.
	//
	// This option has no effect if the partition already exists.
	//
	// Usage of this option in combination with Encrypt= is not supported.
	//
	// For each unique VerityMatchKey= value, a single verity data partition ("Verity=data") and a single verity hash partition
	// ("Verity=hash") must be defined.
	//
	// Added in version 252.
	Verity systemdconf.Value

	// Takes a short, user-chosen identifier string. This setting is used to find sibling verity partitions for the current verity
	// partition. See the description for Verity=.
	//
	// Added in version 252.
	VerityMatchKey systemdconf.Value

	// Configures the data block size of the generated verity hash partition. Must be between 512 and 4096 bytes and must be a power
	// of 2. Defaults to the sector size if configured explicitly, or the underlying block device sector size, or 4K if systemd-repart
	// is not operating on a block device.
	//
	// Added in version 255.
	VerityDataBlockSizeBytes systemdconf.Value

	// Configures the hash block size of the generated verity hash partition. Must be between 512 and 4096 bytes and must be a power
	// of 2. Defaults to the sector size if configured explicitly, or the underlying block device sector size, or 4K if systemd-repart
	// is not operating on a block device.
	//
	// Added in version 255.
	VerityHashBlockSizeBytes systemdconf.Value

	// Takes a boolean argument. If specified the partition is marked for removal during a factory reset operation. This functionality
	// is useful to implement schemes where images can be reset into their original state by removing partitions and creating
	// them anew. Defaults to off.
	//
	// Added in version 245.
	FactoryReset systemdconf.Value

	// Configures the 64-bit GPT partition flags field to set for the partition when creating it. This option has no effect if the
	// partition already exists. If not specified, the flags value is set to all zeroes, except for the three bits that can also
	// be configured via NoAuto=, ReadOnly= and GrowFileSystem=; see below for details on the defaults for these three flags.
	// Specify the flags value in hexadecimal (by prefixing it with "0x"), binary (prefix "0b") or decimal (no prefix).
	//
	// Added in version 249.
	Flags systemdconf.Value

	// Configures the No-Auto, Read-Only and Grow-File-System partition flags (bit 63, 60 and 59) of the partition table entry,
	// as defined by the Discoverable Partitions Specification. Only available for partition types supported by the specification.
	// This option is a friendly way to set bits 63, 60 and 59 of the partition flags value without setting any of the other bits, and
	// may be set via Flags= too, see above.
	//
	// If Flags= is used in conjunction with one or more of NoAuto=/ReadOnly=/GrowFileSystem= the latter control the value of
	// the relevant flags, i.e. the high-level settings NoAuto=/ReadOnly=/GrowFileSystem= override the relevant bits of
	// the low-level setting Flags=.
	//
	// Note that the three flags affect only automatic partition mounting, as implemented by systemd-gpt-auto-generator or
	// the --image= option of various commands (such as systemd-nspawn). It has no effect on explicit mounts, such as those done
	// via mount or fstab.
	//
	// If both bit 60 and 59 are set for a partition (i.e. the partition is marked both read-only and marked for file system growing)
	// the latter is typically without effect: the read-only flag takes precedence in most tools reading these flags, and since
	// growing the file system involves writing to the partition it is consequently ignored.
	//
	// NoAuto= defaults to off. ReadOnly= defaults to on for Verity partition types, and off for all others. GrowFileSystem=
	// defaults to on for all partition types that support it, except if the partition is marked read-only (and thus effectively,
	// defaults to off for Verity partitions).
	//
	// Added in version 249.
	NoAuto systemdconf.Value

	// Configures the No-Auto, Read-Only and Grow-File-System partition flags (bit 63, 60 and 59) of the partition table entry,
	// as defined by the Discoverable Partitions Specification. Only available for partition types supported by the specification.
	// This option is a friendly way to set bits 63, 60 and 59 of the partition flags value without setting any of the other bits, and
	// may be set via Flags= too, see above.
	//
	// If Flags= is used in conjunction with one or more of NoAuto=/ReadOnly=/GrowFileSystem= the latter control the value of
	// the relevant flags, i.e. the high-level settings NoAuto=/ReadOnly=/GrowFileSystem= override the relevant bits of
	// the low-level setting Flags=.
	//
	// Note that the three flags affect only automatic partition mounting, as implemented by systemd-gpt-auto-generator or
	// the --image= option of various commands (such as systemd-nspawn). It has no effect on explicit mounts, such as those done
	// via mount or fstab.
	//
	// If both bit 60 and 59 are set for a partition (i.e. the partition is marked both read-only and marked for file system growing)
	// the latter is typically without effect: the read-only flag takes precedence in most tools reading these flags, and since
	// growing the file system involves writing to the partition it is consequently ignored.
	//
	// NoAuto= defaults to off. ReadOnly= defaults to on for Verity partition types, and off for all others. GrowFileSystem=
	// defaults to on for all partition types that support it, except if the partition is marked read-only (and thus effectively,
	// defaults to off for Verity partitions).
	//
	// Added in version 249.
	ReadOnly systemdconf.Value

	// Configures the No-Auto, Read-Only and Grow-File-System partition flags (bit 63, 60 and 59) of the partition table entry,
	// as defined by the Discoverable Partitions Specification. Only available for partition types supported by the specification.
	// This option is a friendly way to set bits 63, 60 and 59 of the partition flags value without setting any of the other bits, and
	// may be set via Flags= too, see above.
	//
	// If Flags= is used in conjunction with one or more of NoAuto=/ReadOnly=/GrowFileSystem= the latter control the value of
	// the relevant flags, i.e. the high-level settings NoAuto=/ReadOnly=/GrowFileSystem= override the relevant bits of
	// the low-level setting Flags=.
	//
	// Note that the three flags affect only automatic partition mounting, as implemented by systemd-gpt-auto-generator or
	// the --image= option of various commands (such as systemd-nspawn). It has no effect on explicit mounts, such as those done
	// via mount or fstab.
	//
	// If both bit 60 and 59 are set for a partition (i.e. the partition is marked both read-only and marked for file system growing)
	// the latter is typically without effect: the read-only flag takes precedence in most tools reading these flags, and since
	// growing the file system involves writing to the partition it is consequently ignored.
	//
	// NoAuto= defaults to off. ReadOnly= defaults to on for Verity partition types, and off for all others. GrowFileSystem=
	// defaults to on for all partition types that support it, except if the partition is marked read-only (and thus effectively,
	// defaults to off for Verity partitions).
	//
	// Added in version 249.
	GrowFileSystem systemdconf.Value

	// Configures the suffix to append to split artifacts when the --split option of systemd-repart is used. Simple specifier
	// expansion is supported, see below. Defaults to "%t". To disable split artifact generation for a partition, set SplitName=
	// to "-".
	//
	// Added in version 252.
	SplitName systemdconf.Value

	// Takes one of "off", "best", and "guess" (alternatively, also accepts a boolean value, which is mapped to "off" when false,
	// and "best" when true). Defaults to "off". If set to "best", the partition will have the minimal size required to store the
	// sources configured with CopyFiles=. "best" is currently only supported for read-only filesystems. If set to "guess",
	// the partition is created at least as big as required to store the sources configured with CopyFiles=. Note that unless the
	// filesystem is a read-only filesystem, systemd-repart will have to populate the filesystem twice to guess the minimal
	// required size, so enabling this option might slow down repart when populating large partitions.
	//
	// Added in version 253.
	Minimize systemdconf.Value

	// Specifies where and how the partition should be mounted. Takes at least one and at most two fields separated with a colon
	// (":"). The first field specifies where the partition should be mounted. The second field specifies extra mount options
	// to append to the default mount options. These fields correspond to the second and fourth column of the fstab format. This
	// setting may be specified multiple times to mount the partition multiple times. This can be used to add mounts for different
	// btrfs subvolumes located on the same btrfs partition.
	//
	// Note that this setting is only taken into account when --generate-fstab= is specified on the systemd-repart command line.
	//
	// Added in version 256.
	MountPoint systemdconf.Value

	// Specifies how the encrypted partition should be set up. Takes at least one and at most three fields separated with a colon
	// (":"). The first field specifies the encrypted volume name under /dev/mapper/. If not specified, "luks-UUID" will be
	// used where "UUID" is the LUKS UUID. The second field specifies the keyfile to use following the same format as specified
	// in crypttab. The third field specifies a comma-delimited list of crypttab options. These fields correspond to the first,
	// third and fourth column of the crypttab format.
	//
	// Note that this setting is only taken into account when --generate-crypttab= is specified on the systemd-repart command
	// line.
	//
	// Added in version 256.
	EncryptedVolume systemdconf.Value

	// Specifies the compression algorithm to use for the filesystem configured with Format=. Takes a single argument specifying
	// the compression algorithm.
	//
	// Note that this setting is only taken into account when the filesystem configured with Format= supports compression ( btrfs,
	// squashfs, erofs). Here's an incomplete list of compression algorithms supported by the filesystems known to systemd-repart:
	//
	//	+-------------+--------------------------------+---------------+
	//	| FILE SYSTEM |     COMPRESSION ALGORITHMS     | DOCUMENTATION |
	//	+-------------+--------------------------------+---------------+
	//	| squashfs    | gzip, lzo, lz4, xz, zstd, lzma | mksquashfs    |
	//	| erofs       | lz4, lz4hc, lzma, deflate,     | mkfs.erofs    |
	//	|             | libdeflate, zstd               |               |
	//	+-------------+--------------------------------+---------------+
	//
	// Added in version 257.
	Compression systemdconf.Value

	// Specifies the compression level to use for the filesystem configured with Format=. Takes a single argument specifying
	// the compression level to use for the configured compression algorithm. The possible compression levels and their meaning
	// are filesystem specific (refer to the filesystem's documentation for the exact meaning of a particular compression level).
	//
	// Note that this setting is only taken into account when the filesystem configured with Format= supports compression and
	// the Compression= setting is configured explicitly.
	//
	// Added in version 257.
	CompressionLevel systemdconf.Value

	// Takes a partition definition name, such as "10-esp". If specified, systemd-repart will avoid creating this partition
	// and instead prefer to partially merge the two definitions. However, depending on the existing layout of partitions on
	// disk, systemd-repart may be forced to fall back onto un-merging the definitions and using them as originally written,
	// potentially creating this partition. Specifically, systemd-repart will fall back if this partition is found to already
	// exist on disk, or if the target partition already exists on disk but is too small, or if it cannot allocate space for the merged
	// partition for some other reason.
	//
	// The following fields are merged into the target definition in the specified ways: Weight= and PaddingWeight= are simply
	// overwritten; SizeMinBytes= and PaddingMinBytes= use the larger of the two values; SizeMaxBytes= and PaddingMaxBytes=
	// use the smaller value; and CopyFiles=, ExcludeFiles=, ExcludeFilesTarget=, MakeDirectories=, and Subvolumes= are
	// concatenated.
	//
	// Usage of this option in combination with CopyBlocks=, Encrypt=, or Verity= is not supported. The target definition cannot
	// set these settings either. A definition cannot simultaneously be a supplement and act as a target for some other supplement
	// definition. A target cannot have more than one supplement partition associated with it.
	//
	// For example, distributions can use this to implement $BOOT as defined in the Boot Loader Specification. Distributions
	// may prefer to use the ESP as $BOOT whenever possible, but to adhere to the spec XBOOTLDR must sometimes be used instead. So,
	// they should create two definitions: the first defining an ESP big enough to hold just the bootloader, and a second for the
	// XBOOTLDR that's sufficiently large to hold kernels and configured as a supplement for the ESP. Whenever possible, systemd-repart
	// will try to merge the two definitions to create one large ESP, but if that's not allowable due to the existing conditions
	// on disk a small ESP and a large XBOOTLDR will be created instead.
	//
	// As another example, distributions can also use this to seamlessly share a single /home partition in a multi-boot scenario,
	// while preferring to keep /home on the root partition by default. Having a /home partition separated from the root partition
	// entails some extra complexity: someone has to decide how to split the space between the two partitions. On the other hand,
	// it allows a user to share their home area between multiple installed OSs (i.e. via systemd-homed.service). Distributions
	// should create two definitions: the first for a root partition that takes up some relatively small percentage of the disk,
	// and the second as a supplement for the first to create a /home partition that takes up all the remaining free space. On first
	// boot, if systemd-repart finds an existing /home partition on disk, it'll un-merge the definitions and create just a small
	// root partition. Otherwise, the definitions will be merged and a single large root partition will be created.
	//
	// Added in version 257.
	SupplementFor systemdconf.Value
}
