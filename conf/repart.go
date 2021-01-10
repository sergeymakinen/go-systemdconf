// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// RepartFile represents repart.d — Partition Definition Files for Automatic Boot-Time Repartitioning
// (see https://www.freedesktop.org/software/systemd/man/repart.d.html for details)
type RepartFile struct {
	systemdconf.File

	Partition RepartPartitionSection // Basic properties of partitions of block devices of the local system
}

// RepartPartitionSection represents basic properties of partitions of block devices of the local system
// (see https://www.freedesktop.org/software/systemd/man/repart.d.html#%5BPartition%5D%20Section%20Options for details)
type RepartPartitionSection struct {
	systemdconf.Section

	// The GPT partition type UUID to match. This may be a GPT partition type UUID such as 4f68bce3-e8cd-4db1-96e7-fbcaf984b709,
	// or one of the following special identifiers:
	//
	// 	+-----------------------+----------------------------------------------------------------------------------+
	// 	|      IDENTIFIER       |                                   EXPLANATION                                    |
	// 	+-----------------------+----------------------------------------------------------------------------------+
	// 	| esp                   | EFI System Partition                                                             |
	// 	| xbootldr              | Extended Boot Loader Partition                                                   |
	// 	| swap                  | Swap partition                                                                   |
	// 	| home                  | Home (/home/) partition                                                          |
	// 	| srv                   | Server data (/srv/) partition                                                    |
	// 	| var                   | Variable data (/var/) partition                                                  |
	// 	| tmp                   | Temporary data (/var/tmp/) partition                                             |
	// 	| linux-generic         | Generic Linux file system partition                                              |
	// 	| root                  | Root file system partition type appropriate for the local architecture (an       |
	// 	|                       | alias for an architecture root file system partition type listed below, e.g.     |
	// 	|                       | root-x86-64)                                                                     |
	// 	| root-verity           | Verity data for the root file system partition for the local architecture        |
	// 	| root-secondary        | Root file system partition of the secondary architecture of the local            |
	// 	|                       | architecture (usually the matching 32bit architecture for the local 64bit        |
	// 	|                       | architecture)                                                                    |
	// 	| root-secondary-verity | Verity data for the root file system partition of the secondary architecture     |
	// 	| root-x86              | Root file system partition for the x86 (32bit, aka i386) architecture            |
	// 	| root-x86-verity       | Verity data for the x86 (32bit) root file system partition                       |
	// 	| root-x86-64           | Root file system partition for the x86_64 (64bit, aka amd64) architecture        |
	// 	| root-x86-64-verity    | Verity data for the x86_64 (64bit) root file system partition                    |
	// 	| root-arm              | Root file system partition for the ARM (32bit) architecture                      |
	// 	| root-arm-verity       | Verity data for the ARM (32bit) root file system partition                       |
	// 	| root-arm64            | Root file system partition for the ARM (64bit, aka aarch64) architecture         |
	// 	| root-arm64-verity     | Verity data for the ARM (64bit, aka aarch64) root file system partition          |
	// 	| root-ia64             | Root file system partition for the ia64 architecture                             |
	// 	| root-ia64-verity      | Verity data for the ia64 root file system partition                              |
	// 	| root-riscv32          | Root file system partition for the RISC-V 32-bit architecture                    |
	// 	| root-riscv32-verity   | Verity data for the RISC-V 32-bit root file system partition                     |
	// 	| root-riscv64          | Root file system partition for the RISC-V 64-bit architecture                    |
	// 	| root-riscv64-verity   | Verity data for the RISC-V 64-bit root file system partition                     |
	// 	| usr                   | /usr/ file system partition type appropriate for the local architecture (an      |
	// 	|                       | alias for an architecture /usr/ file system partition type listed below, e.g.    |
	// 	|                       | usr-x86-64)                                                                      |
	// 	| usr-verity            | Verity data for the /usr/ file system partition for the local architecture       |
	// 	| usr-secondary         | /usr/ file system partition of the secondary architecture of the local           |
	// 	|                       | architecture (usually the matching 32bit architecture for the local 64bit        |
	// 	|                       | architecture)                                                                    |
	// 	| usr-secondary-verity  | Verity data for the /usr/ file system partition of the secondary architecture    |
	// 	| usr-x86               | /usr/ file system partition for the x86 (32bit, aka i386) architecture           |
	// 	| usr-x86-verity        | Verity data for the x86 (32bit) /usr/ file system partition                      |
	// 	| usr-x86-64            | /usr/ file system partition for the x86_64 (64bit, aka amd64) architecture       |
	// 	| usr-x86-64-verity     | Verity data for the x86_64 (64bit) /usr/ file system partition                   |
	// 	| usr-arm               | /usr/ file system partition for the ARM (32bit) architecture                     |
	// 	| usr-arm-verity        | Verity data for the ARM (32bit) /usr/ file system partition                      |
	// 	| usr-arm64             | /usr/ file system partition for the ARM (64bit, aka aarch64) architecture        |
	// 	| usr-arm64-verity      | Verity data for the ARM (64bit, aka aarch64) /usr/ file system partition         |
	// 	| usr-ia64              | /usr/ file system partition for the ia64 architecture                            |
	// 	| usr-ia64-verity       | Verity data for the ia64 /usr/ file system partition                             |
	// 	| usr-riscv32           | /usr/ file system partition for the RISC-V 32-bit architecture                   |
	// 	| usr-riscv32-verity    | Verity data for the RISC-V 32-bit /usr/ file system partition                    |
	// 	| usr-riscv64           | /usr/ file system partition for the RISC-V 64-bit architecture                   |
	// 	| usr-riscv64-verity    | Verity data for the RISC-V 64-bit /usr/ file system partition                    |
	// 	+-----------------------+----------------------------------------------------------------------------------+
	//
	// This setting defaults to linux-generic.
	//
	// Most of the partition type UUIDs listed above are defined in the Discoverable Partitions Specification.
	Type systemdconf.Value

	// The textual label to assign to the partition if none is assigned yet. Note that this setting is not used for matching. It is
	// also not used when a label is already set for an existing partition. It is thus only used when a partition is newly created
	// or when an existing one had a no label set (that is: an empty label). If not specified a label derived from the partition type
	// is automatically used. Simple specifier expansion is supported, see below.
	Label systemdconf.Value

	// The UUID to assign to the partition if none is assigned yet. Note that this setting is not used for matching. It is also not
	// used when a UUID is already set for an existing partition. It is thus only used when a partition is newly created or when an
	// existing one had a all-zero UUID set. If not specified a UUID derived from the partition type is automatically used.
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
	Priority systemdconf.Value

	// A numeric weight to assign to this partition in the range 0…1000000. Available disk space is assigned the defined partitions
	// according to their relative weights (subject to the size constraints configured with SizeMinBytes=, SizeMaxBytes=),
	// so that a partition with weight 2000 gets double the space as one with weight 1000, and a partition with weight 333 a third
	// of that. Defaults to 1000.
	//
	// The Weight= setting is used to distribute available disk space in an "elastic" fashion, based on the disk size and existing
	// partitions. If a partition shall have a fixed size use both SizeMinBytes= and SizeMaxBytes= with the same value in order
	// to fixate the size to one value, in which case the weight has no effect.
	Weight systemdconf.Value

	// Similar to Weight= but sets a weight for the free space after the partition (the "padding"). When distributing available
	// space the weights of all partitions and all defined padding is summed, and then each partition and padding gets the fraction
	// defined by its weight. Defaults to 0, i.e. by default no padding is applied.
	//
	// Padding is useful if empty space shall be left for later additions or a safety margin at the end of the device or between partitions.
	PaddingWeight systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes. Takes the usual K, M, G, T, … suffixes (to the base of 1024). If SizeMinBytes=
	// is specified the partition is created at or grown to at least the specified size. If SizeMaxBytes= is specified the partition
	// is created at or grown to at most the specified size. The precise size is determined through the weight value value configured
	// with Weight=, see above. When SizeMinBytes= is set equal to SizeMaxBytes= the configured weight has no effect as the partition
	// is explicitly sized to the specified fixed value. Note that partitions are never created smaller than 4096 bytes, and since
	// partitions are never shrunk the previous size of the partition (in case the partition already exists) is also enforced
	// as lower bound for the new size. The values should be specified as multiples of 4096 bytes, and are rounded upwards (in case
	// of SizeMinBytes=) or downwards (in case of SizeMaxBytes=) otherwise. If the backing device does not provide enough space
	// to fulfill the constraints placing the partition will fail. For partitions that shall be created, depending on the setting
	// of Priority= (see above) the partition might be dropped and the placing algorithm restarted. By default a minimum size
	// constraint of 10M and no maximum size constraint is set.
	SizeMinBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes. Takes the usual K, M, G, T, … suffixes (to the base of 1024). If SizeMinBytes=
	// is specified the partition is created at or grown to at least the specified size. If SizeMaxBytes= is specified the partition
	// is created at or grown to at most the specified size. The precise size is determined through the weight value value configured
	// with Weight=, see above. When SizeMinBytes= is set equal to SizeMaxBytes= the configured weight has no effect as the partition
	// is explicitly sized to the specified fixed value. Note that partitions are never created smaller than 4096 bytes, and since
	// partitions are never shrunk the previous size of the partition (in case the partition already exists) is also enforced
	// as lower bound for the new size. The values should be specified as multiples of 4096 bytes, and are rounded upwards (in case
	// of SizeMinBytes=) or downwards (in case of SizeMaxBytes=) otherwise. If the backing device does not provide enough space
	// to fulfill the constraints placing the partition will fail. For partitions that shall be created, depending on the setting
	// of Priority= (see above) the partition might be dropped and the placing algorithm restarted. By default a minimum size
	// constraint of 10M and no maximum size constraint is set.
	SizeMaxBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes for the free space after the partition (the "padding"). Semantics
	// are similar to SizeMinBytes= and SizeMaxBytes=, except that unlike partition sizes free space can be shrunk and can be
	// as small as zero. By default no size constraints on padding are set, so that only PaddingWeight= determines the size of the
	// padding applied.
	PaddingMinBytes systemdconf.Value

	// Specifies minimum and maximum size constraints in bytes for the free space after the partition (the "padding"). Semantics
	// are similar to SizeMinBytes= and SizeMaxBytes=, except that unlike partition sizes free space can be shrunk and can be
	// as small as zero. By default no size constraints on padding are set, so that only PaddingWeight= determines the size of the
	// padding applied.
	PaddingMaxBytes systemdconf.Value

	// Takes a path to a regular file, block device node or directory. If specified and the partition is newly created the data from
	// the specified path is written to the newly created partition, on the block level. If a directory is specified the backing
	// block device of the file system the directory is on is determined and the data read directly from that. This option is useful
	// to efficiently replicate existing file systems on the block level on a new partition, for example to build a simple OS installer
	// or OS image builder.
	//
	// The file specified here must have a size that is a multiple of the basic block size 512 and not be empty. If this option is used,
	// the size allocation algorithm is slightly altered: the partition is created as least as big as required to fit the data in,
	// i.e. the data size is an additional minimum size value taken into consideration for the allocation algorithm, similar
	// to and in addition to the SizeMin= value configured above.
	//
	// This option has no effect if the partition it is declared for already exists, i.e. existing data is never overwritten. Note
	// that the data is copied in before the partition table is updated, i.e. before the partition actually is persistently created.
	// This provides robustness: it is guaranteed that the partition either doesn't exist or exists fully populated; it is not
	// possible that the partition exists but is not or only partially populated.
	//
	// This option cannot be combined with Format= or CopyFiles=.
	CopyBlocks systemdconf.Value

	// Takes a file system name, such as "ext4", "btrfs", "xfs" or "vfat", or the special value "swap". If specified and the partition
	// is newly created it is formatted with the specified file system (or as swap device). The file system UUID and label are automatically
	// derived from the partition UUID and label. If this option is used, the size allocation algorithm is slightly altered: the
	// partition is created as least as big as required for the minimal file system of the specified type (or 4KiB if the minimal
	// size is not known).
	//
	// This option has no effect if the partition already exists.
	//
	// Similar to the behaviour of CopyBlocks= the file system is formatted before the partition is created, ensuring that the
	// partition only ever exists with a fully initialized file system.
	//
	// This option cannot be combined with CopyBlocks=.
	Format systemdconf.Value

	// Takes a pair of colon separated absolute file system paths. The first path refers to a source file or directory on the host,
	// the second path refers to a target in the file system of the newly created partition and formatted file system. This setting
	// may be used to copy files or directories from the host into the file system that is created due to the Format= option. If CopyFiles=
	// is used without Format= specified explicitly, "Format=" with a suitable default is implied (currently "ext4", but this
	// may change in the future). This option may be used multiple times to copy multiple files or directories from host into the
	// newly formatted file system. The colon and second path may be omitted in which case the source path is also used as the target
	// path (relative to the root of the newly created file system). If the source path refers to a directory it is copied recursively.
	//
	// This option has no effect if the partition already exists: it cannot be used to copy additional files into an existing partition,
	// it may only be used to populate a file system created anew.
	//
	// The copy operation is executed before the file system is registered in the partition table, thus ensuring that a file system
	// populated this way only ever exists fully initialized.
	//
	// This option cannot be combined with CopyBlocks=.
	CopyFiles systemdconf.Value

	// Takes one of "off", "key-file", "tpm2" and "key-file+tpm2" (alternatively, also accepts a boolean value, which is mapped
	// to "off" when false, and "key-file" when true). Defaults to "off". If not "off" the partition will be formatted with a LUKS2
	// superblock, before the blocks configured with CopyBlocks= are copied in or the file system configured with Format= is
	// created.
	//
	// The LUKS2 UUID is automatically derived from the partition UUID in a stable fashion. If "key-file" or "key-file+tpm2"
	// is used a key is added to the LUKS2 superblock, configurable with the --key-file= switch to systemd-repart. If "tpm2" or
	// "key-file+tpm2" is used a key is added to the LUKS2 superblock that is enrolled to the local TPM2 chip, as configured with
	// the --tpm2-device= and --tpm2-pcrs= options to systemd-repart.
	//
	// When used this slightly alters the size allocation logic as the implicit, minimal size limits of Format= and CopyBlocks=
	// are increased by the space necessary for the LUKS2 superblock (see above).
	//
	// This option has no effect if the partition already exists.
	Encrypt systemdconf.Value

	// Takes a boolean argument. If specified the partition is marked for removal during a factory reset operation. This functionality
	// is useful to implement schemes where images can be reset into their original state by removing partitions and creating
	// them anew. Defaults to off.
	FactoryReset systemdconf.Value
}
