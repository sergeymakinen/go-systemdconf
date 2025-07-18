// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// ExecOptions represents systemd.exec — Execution environment configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.exec.html for details).
type ExecOptions struct {
	// Takes a colon separated list of absolute paths relative to which the executable used by the Exec*= (e.g. ExecStart=, ExecStop=,
	// etc.) properties can be found. ExecSearchPath= overrides $PATH if $PATH is not supplied by the user through Environment=,
	// EnvironmentFile= or PassEnvironment=. Assigning an empty string removes previous assignments and setting ExecSearchPath=
	// to a value multiple times will append to the previous setting.
	//
	// Added in version 250.
	ExecSearchPath systemdconf.Value

	// Takes a directory path relative to the service's root directory specified by RootDirectory=, or the special value "~".
	// Sets the working directory for executed processes. If set to "~", the home directory of the user specified in User= is used.
	// If not set, defaults to the root directory when systemd is running as a system instance and the respective user's home directory
	// if run as user. If the setting is prefixed with the "-" character, a missing working directory is not considered fatal. If
	// RootDirectory=/RootImage= is not set, then WorkingDirectory= is relative to the root of the system running the service
	// manager. Note that setting this parameter might result in additional dependencies to be added to the unit (see above).
	WorkingDirectory systemdconf.Value

	// Takes a directory path relative to the host's root directory (i.e. the root of the system running the service manager).
	// Sets the root directory for executed processes, with the pivot_root or chroot system call. If this is used, it must be ensured
	// that the process binary and all its auxiliary files are available in the new root. Note that setting this parameter might
	// result in additional dependencies to be added to the unit (see above).
	//
	// The MountAPIVFS= and PrivateUsers= settings are particularly useful in conjunction with RootDirectory=. For details,
	// see below.
	//
	// If RootDirectory=/RootImage= are used together with NotifyAccess= the notification socket is automatically mounted
	// from the host into the root environment, to ensure the notification interface can work correctly.
	//
	// Note that services using RootDirectory=/RootImage= will not be able to log via the syslog or journal protocols to the host
	// logging infrastructure, unless the relevant sockets are mounted from the host, specifically:
	//
	// The host's os-release file will be made available for the service (read-only) as /run/host/os-release. It will be updated
	// automatically on soft reboot (see: systemd-soft-reboot.service), in case the service is configured to survive it.
	//
	//	BindReadOnlyPaths=/dev/log /run/systemd/journal/socket /run/systemd/journal/stdout
	//
	// In place of the directory path a ".v/" versioned directory may be specified, see systemd.v for details.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	RootDirectory systemdconf.Value

	// Takes a path to a block device node or regular file as argument. This call is similar to RootDirectory= however mounts a file
	// system hierarchy from a block device node or loopback file instead of a directory. The device node or file system image file
	// needs to contain a file system without a partition table, or a file system within an MBR/MS-DOS or GPT partition table with
	// only a single Linux-compatible partition, or a set of file systems within a GPT partition table that follows the  Discoverable
	// Partitions Specification.
	//
	// When DevicePolicy= is set to "closed" or "strict", or set to "auto" and DeviceAllow= is set, then this setting adds /dev/loop-control
	// with rw mode, "block-loop" and "block-blkext" with rwm mode to DeviceAllow=. See systemd.resource-control for the details
	// about DevicePolicy= or DeviceAllow=. Also, see PrivateDevices= below, as it may change the setting of DevicePolicy=.
	//
	// Units making use of RootImage= automatically gain an After= dependency on systemd-udevd.service.
	//
	// The host's os-release file will be made available for the service (read-only) as /run/host/os-release. It will be updated
	// automatically on soft reboot (see: systemd-soft-reboot.service), in case the service is configured to survive it.
	//
	// In place of the image path a ".v/" versioned directory may be specified, see systemd.v for details.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 233.
	RootImage systemdconf.Value

	// Takes a comma-separated list of mount options that will be used on disk images specified by RootImage=. Optionally a partition
	// name can be prefixed, followed by colon, in case the image has multiple partitions, otherwise partition name "root" is
	// implied. Options for multiple partitions can be specified in a single line with space separators. Assigning an empty string
	// removes previous assignments. Duplicated options are ignored. For a list of valid mount options, please refer to mount.
	//
	// Valid partition names follow the  Discoverable Partitions Specification: root, usr, home, srv, esp, xbootldr, tmp, var.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 247.
	RootImageOptions systemdconf.Value

	// Takes a boolean argument. If enabled, executed processes will run in an ephemeral copy of the root directory or root image.
	// The ephemeral copy is placed in /var/lib/systemd/ephemeral-trees/ while the service is active and is cleaned up when
	// the service is stopped or restarted. If RootDirectory= is used and the root directory is a subvolume, the ephemeral copy
	// will be created by making a snapshot of the subvolume.
	//
	// To make sure making ephemeral copies can be made efficiently, the root directory or root image should be located on the same
	// filesystem as /var/lib/systemd/ephemeral-trees/. When using RootEphemeral= with root directories, btrfs should
	// be used as the filesystem and the root directory should ideally be a subvolume which systemd can snapshot to make the ephemeral
	// copy. For root images, a filesystem with support for reflinks should be used to ensure an efficient ephemeral copy.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 254.
	RootEphemeral systemdconf.Value

	// Takes a data integrity (dm-verity) root hash specified in hexadecimal, or the path to a file containing a root hash in ASCII
	// hexadecimal format. This option enables data integrity checks using dm-verity, if the used image contains the appropriate
	// integrity data (see above) or if RootVerity= is used. The specified hash must match the root hash of integrity data, and
	// is usually at least 256 bits (and hence 64 formatted hexadecimal characters) long (in case of SHA256 for example). If this
	// option is not specified, but the image file carries the "user.verity.roothash" extended file attribute (see xattr),
	// then the root hash is read from it, also as formatted hexadecimal characters. If the extended file attribute is not found
	// (or is not supported by the underlying file system), but a file with the .roothash suffix is found next to the image file,
	// bearing otherwise the same name (except if the image has the .raw suffix, in which case the root hash file must not have it
	// in its name), the root hash is read from it and automatically used, also as formatted hexadecimal characters.
	//
	// If the disk image contains a separate /usr/ partition it may also be Verity protected, in which case the root hash may configured
	// via an extended attribute "user.verity.usrhash" or a .usrhash file adjacent to the disk image. There's currently no option
	// to configure the root hash for the /usr/ file system via the unit file directly.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 246.
	RootHash systemdconf.Value

	// Takes a PKCS7 signature of the RootHash= option as a path to a DER-encoded signature file, or as an ASCII base64 string encoding
	// of a DER-encoded signature prefixed by "base64:". The dm-verity volume will only be opened if the signature of the root
	// hash is valid and signed by a public key present in the kernel keyring. If this option is not specified, but a file with the
	// .roothash.p7s suffix is found next to the image file, bearing otherwise the same name (except if the image has the .raw suffix,
	// in which case the signature file must not have it in its name), the signature is read from it and automatically used.
	//
	// If the disk image contains a separate /usr/ partition it may also be Verity protected, in which case the signature for the
	// root hash may configured via a .usrhash.p7s file adjacent to the disk image. There's currently no option to configure the
	// root hash signature for the /usr/ via the unit file directly.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 246.
	RootHashSignature systemdconf.Value

	// Takes the path to a data integrity (dm-verity) file. This option enables data integrity checks using dm-verity, if RootImage=
	// is used and a root-hash is passed and if the used image itself does not contain the integrity data. The integrity data must
	// be matched by the root hash. If this option is not specified, but a file with the .verity suffix is found next to the image file,
	// bearing otherwise the same name (except if the image has the .raw suffix, in which case the verity data file must not have
	// it in its name), the verity data is read from it and automatically used.
	//
	// This option is supported only for disk images that contain a single file system, without an enveloping partition table.
	// Images that contain a GPT partition table should instead include both root file system and matching Verity data in the same
	// image, implementing the  Discoverable Partitions Specification.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 246.
	RootVerity systemdconf.Value

	// Takes an image policy string as per systemd.image-policy to use when mounting the disk images (DDI) specified in RootImage=,
	// MountImage=, ExtensionImage=, respectively. If not specified the following policy string is the default for RootImagePolicy=
	// and MountImagePolicy:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent: \ home=encrypted+unprotected+absent:
	//	\ srv=encrypted+unprotected+absent: \ tmp=encrypted+unprotected+absent: \ var=encrypted+unprotected+absent
	//
	// The default policy for ExtensionImagePolicy= is:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent
	//
	// Added in version 254.
	RootImagePolicy systemdconf.Value

	// Takes an image policy string as per systemd.image-policy to use when mounting the disk images (DDI) specified in RootImage=,
	// MountImage=, ExtensionImage=, respectively. If not specified the following policy string is the default for RootImagePolicy=
	// and MountImagePolicy:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent: \ home=encrypted+unprotected+absent:
	//	\ srv=encrypted+unprotected+absent: \ tmp=encrypted+unprotected+absent: \ var=encrypted+unprotected+absent
	//
	// The default policy for ExtensionImagePolicy= is:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent
	//
	// Added in version 254.
	MountImagePolicy systemdconf.Value

	// Takes an image policy string as per systemd.image-policy to use when mounting the disk images (DDI) specified in RootImage=,
	// MountImage=, ExtensionImage=, respectively. If not specified the following policy string is the default for RootImagePolicy=
	// and MountImagePolicy:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent: \ home=encrypted+unprotected+absent:
	//	\ srv=encrypted+unprotected+absent: \ tmp=encrypted+unprotected+absent: \ var=encrypted+unprotected+absent
	//
	// The default policy for ExtensionImagePolicy= is:
	//
	//	root=verity+signed+encrypted+unprotected+absent: \ usr=verity+signed+encrypted+unprotected+absent
	//
	// Added in version 254.
	ExtensionImagePolicy systemdconf.Value

	// Takes a boolean argument. If on, a private mount namespace for the unit's processes is created and the API file systems /proc/,
	// /sys/, /dev/ and /run/ (as an empty "tmpfs") are mounted inside of it, unless they are already mounted. Note that this option
	// has no effect unless used in conjunction with RootDirectory=/RootImage= as these four mounts are generally mounted in
	// the host anyway, and unless the root directory is changed, the private mount namespace will be a 1:1 copy of the host's, and
	// include these four mounts. Note that the /dev/ file system of the host is bind mounted if this option is used without PrivateDevices=.
	// To run the service with a private, minimal version of /dev/, combine this option with PrivateDevices=.
	//
	// In order to allow propagating mounts at runtime in a safe manner, /run/systemd/propagate/ on the host will be used to set
	// up new mounts, and /run/host/incoming/ in the private namespace will be used as an intermediate step to store them before
	// being moved to the final mount point.
	//
	// Added in version 233.
	MountAPIVFS systemdconf.Value

	// Takes a boolean argument. If true, sockets from systemd-journald.socket will be bind mounted into the mount namespace.
	// This is particularly useful when a different instance of /run/ is employed, to make sure processes running in the namespace
	// can still make use of sd-journal.
	//
	// This option is implied when LogNamespace= is used, when MountAPIVFS=yes, or when PrivateDevices=yes is used in conjunction
	// with either RootDirectory= or RootImage=.
	//
	// Added in version 257.
	BindLogSockets systemdconf.Value

	// Takes one of "noaccess", "invisible", "ptraceable" or "default" (which it defaults to). When set, this controls the "hidepid="
	// mount option of the "procfs" instance for the unit that controls which directories with process metainformation (/proc/PID)
	// are visible and accessible: when set to "noaccess" the ability to access most of other users' process metadata in /proc/
	// is taken away for processes of the service. When set to "invisible" processes owned by other users are hidden from /proc/.
	// If "ptraceable" all processes that cannot be ptrace()'ed by a process are hidden to it. If "default" no restrictions on
	// /proc/ access or visibility are made. For further details see The /proc Filesystem. It is generally recommended to run
	// most system services with this option set to "invisible". This option is implemented via file system namespacing, and
	// thus cannot be used with services that shall be able to install mount points in the host file system hierarchy. Note that
	// the root user is unaffected by this option, so to be effective it has to be used together with User= or DynamicUser=yes, and
	// also without the "CAP_SYS_PTRACE" capability, which also allows a process to bypass this feature. It cannot be used for
	// services that need to access metainformation about other users' processes. This option implies MountAPIVFS=.
	//
	// If the kernel does not support per-mount point hidepid= mount options this setting remains without effect, and the unit's
	// processes will be able to access and see other process as if the option was not used.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 247.
	ProtectProc systemdconf.Value

	// Takes one of "all" (the default) and "pid". If "pid", all files and directories not directly associated with process management
	// and introspection are made invisible in the /proc/ file system configured for the unit's processes. This controls the
	// "subset=" mount option of the "procfs" instance for the unit. For further details see The /proc Filesystem. Note that Linux
	// exposes various kernel APIs via /proc/, which are made unavailable with this setting. Since these APIs are used frequently
	// this option is useful only in a few, specific cases, and is not suitable for most non-trivial programs.
	//
	// Much like ProtectProc= above, this is implemented via file system mount namespacing, and hence the same restrictions
	// apply: it is only available to system services, it disables mount propagation to the host mount table, and it implies MountAPIVFS=.
	// Also, like ProtectProc= this setting is gracefully disabled if the used kernel does not support the "subset=" mount option
	// of "procfs".
	//
	// Added in version 247.
	ProcSubset systemdconf.Value

	// Configures unit-specific bind mounts. A bind mount makes a particular file or directory available at an additional place
	// in the unit's view of the file system. Any bind mounts created with this option are specific to the unit, and are not visible
	// in the host's mount table. This option expects a whitespace separated list of bind mount definitions. Each definition
	// consists of a colon-separated triple of source path, destination path and option string, where the latter two are optional.
	// If only a source path is specified the source and destination is taken to be the same. The option string may be either "rbind"
	// or "norbind" for configuring a recursive or non-recursive bind mount. If the destination path is omitted, the option string
	// must be omitted too. Each bind mount definition may be prefixed with "-", in which case it will be ignored when its source
	// path does not exist.
	//
	// BindPaths= creates regular writable bind mounts (unless the source file system mount is already marked read-only), while
	// BindReadOnlyPaths= creates read-only bind mounts. These settings may be used more than once, each usage appends to the
	// unit's list of bind mounts. If the empty string is assigned to either of these two options the entire list of bind mounts defined
	// prior to this is reset. Note that in this case both read-only and regular bind mounts are reset, regardless which of the two
	// settings is used.
	//
	// Using this option implies that a mount namespace is allocated for the unit, i.e. it implies the effect of PrivateMounts=
	// (see below).
	//
	// This option is particularly useful when RootDirectory=/RootImage= is used. In this case the source path refers to a path
	// on the host file system, while the destination path refers to a path below the root directory of the unit.
	//
	// Note that the destination directory must exist or systemd must be able to create it. Thus, it is not possible to use those
	// options for mount points nested underneath paths specified in InaccessiblePaths=, or under /home/ and other protected
	// directories if ProtectHome=yes is specified. TemporaryFileSystem= with ":ro" or ProtectHome=tmpfs should be used
	// instead.
	//
	// Added in version 233.
	BindPaths systemdconf.Value

	// Configures unit-specific bind mounts. A bind mount makes a particular file or directory available at an additional place
	// in the unit's view of the file system. Any bind mounts created with this option are specific to the unit, and are not visible
	// in the host's mount table. This option expects a whitespace separated list of bind mount definitions. Each definition
	// consists of a colon-separated triple of source path, destination path and option string, where the latter two are optional.
	// If only a source path is specified the source and destination is taken to be the same. The option string may be either "rbind"
	// or "norbind" for configuring a recursive or non-recursive bind mount. If the destination path is omitted, the option string
	// must be omitted too. Each bind mount definition may be prefixed with "-", in which case it will be ignored when its source
	// path does not exist.
	//
	// BindPaths= creates regular writable bind mounts (unless the source file system mount is already marked read-only), while
	// BindReadOnlyPaths= creates read-only bind mounts. These settings may be used more than once, each usage appends to the
	// unit's list of bind mounts. If the empty string is assigned to either of these two options the entire list of bind mounts defined
	// prior to this is reset. Note that in this case both read-only and regular bind mounts are reset, regardless which of the two
	// settings is used.
	//
	// Using this option implies that a mount namespace is allocated for the unit, i.e. it implies the effect of PrivateMounts=
	// (see below).
	//
	// This option is particularly useful when RootDirectory=/RootImage= is used. In this case the source path refers to a path
	// on the host file system, while the destination path refers to a path below the root directory of the unit.
	//
	// Note that the destination directory must exist or systemd must be able to create it. Thus, it is not possible to use those
	// options for mount points nested underneath paths specified in InaccessiblePaths=, or under /home/ and other protected
	// directories if ProtectHome=yes is specified. TemporaryFileSystem= with ":ro" or ProtectHome=tmpfs should be used
	// instead.
	//
	// Added in version 233.
	BindReadOnlyPaths systemdconf.Value

	// This setting is similar to RootImage= in that it mounts a file system hierarchy from a block device node or loopback file,
	// but the destination directory can be specified as well as mount options. This option expects a whitespace separated list
	// of mount definitions. Each definition consists of a colon-separated tuple of source path and destination definitions,
	// optionally followed by another colon and a list of mount options.
	//
	// Mount options may be defined as a single comma-separated list of options, in which case they will be implicitly applied
	// to the root partition on the image, or a series of colon-separated tuples of partition name and mount options. Valid partition
	// names and mount options are the same as for RootImageOptions= setting described above.
	//
	// Each mount definition may be prefixed with "-", in which case it will be ignored when its source path does not exist. The source
	// argument is a path to a block device node or regular file. If source or destination contain a ":", it needs to be escaped as
	// "\:". The device node or file system image file needs to follow the same rules as specified for RootImage=. Any mounts created
	// with this option are specific to the unit, and are not visible in the host's mount table.
	//
	// These settings may be used more than once, each usage appends to the unit's list of mount paths. If the empty string is assigned,
	// the entire list of mount paths defined prior to this is reset.
	//
	// Note that the destination directory must exist or systemd must be able to create it. Thus, it is not possible to use those
	// options for mount points nested underneath paths specified in InaccessiblePaths=, or under /home/ and other protected
	// directories if ProtectHome=yes is specified.
	//
	// When DevicePolicy= is set to "closed" or "strict", or set to "auto" and DeviceAllow= is set, then this setting adds /dev/loop-control
	// with rw mode, "block-loop" and "block-blkext" with rwm mode to DeviceAllow=. See systemd.resource-control for the details
	// about DevicePolicy= or DeviceAllow=. Also, see PrivateDevices= below, as it may change the setting of DevicePolicy=.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 247.
	MountImages systemdconf.Value

	// This setting is similar to MountImages= in that it mounts a file system hierarchy from a block device node or loopback file,
	// but instead of providing a destination path, an overlay will be set up. This option expects a whitespace separated list
	// of mount definitions. Each definition consists of a source path, optionally followed by a colon and a list of mount options.
	//
	// A read-only OverlayFS will be set up on top of /usr/ and /opt/ hierarchies for sysext images and /etc/ hierarchy for confext
	// images. The order in which the images are listed will determine the order in which the overlay is laid down: images specified
	// first to last will result in overlayfs layers bottom to top.
	//
	// Mount options may be defined as a single comma-separated list of options, in which case they will be implicitly applied
	// to the root partition on the image, or a series of colon-separated tuples of partition name and mount options. Valid partition
	// names and mount options are the same as for RootImageOptions= setting described above.
	//
	// Each mount definition may be prefixed with "-", in which case it will be ignored when its source path does not exist. The source
	// argument is a path to a block device node or regular file. If the source path contains a ":", it needs to be escaped as "\:".
	// The device node or file system image file needs to follow the same rules as specified for RootImage=. Any mounts created
	// with this option are specific to the unit, and are not visible in the host's mount table.
	//
	// These settings may be used more than once, each usage appends to the unit's list of image paths. If the empty string is assigned,
	// the entire list of mount paths defined prior to this is reset.
	//
	// Each sysext image must carry a /usr/lib/extension-release.d/extension-release.IMAGE file while each confext image
	// must carry a /etc/extension-release.d/extension-release.IMAGE file, with the appropriate metadata which matches
	// RootImage=/RootDirectory= or the host. See: os-release. To disable the safety check that the extension-release file
	// name matches the image file name, the x-systemd.relax-extension-release-check mount option may be appended.
	//
	// When DevicePolicy= is set to "closed" or "strict", or set to "auto" and DeviceAllow= is set, then this setting adds /dev/loop-control
	// with rw mode, "block-loop" and "block-blkext" with rwm mode to DeviceAllow=. See systemd.resource-control for the details
	// about DevicePolicy= or DeviceAllow=. Also, see PrivateDevices= below, as it may change the setting of DevicePolicy=.
	//
	// In place of the image path a ".v/" versioned directory may be specified, see systemd.v for details.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 248.
	ExtensionImages systemdconf.Value

	// This setting is similar to BindReadOnlyPaths= in that it mounts a file system hierarchy from a directory, but instead of
	// providing a destination path, an overlay will be set up. This option expects a whitespace separated list of source directories.
	//
	// A read-only OverlayFS will be set up on top of /usr/ and /opt/ hierarchies for sysext images and /etc/ hierarchy for confext
	// images. The order in which the directories are listed will determine the order in which the overlay is laid down: directories
	// specified first to last will result in overlayfs layers bottom to top.
	//
	// Each directory listed in ExtensionDirectories= may be prefixed with "-", in which case it will be ignored when its source
	// path does not exist. Any mounts created with this option are specific to the unit, and are not visible in the host's mount
	// table.
	//
	// These settings may be used more than once, each usage appends to the unit's list of directories paths. If the empty string
	// is assigned, the entire list of mount paths defined prior to this is reset.
	//
	// Each sysext directory must contain a /usr/lib/extension-release.d/extension-release.IMAGE file while each confext
	// directory must carry a /etc/extension-release.d/extension-release.IMAGE file, with the appropriate metadata which
	// matches RootImage=/RootDirectory= or the host. See: os-release.
	//
	// Note that usage from user units requires overlayfs support in unprivileged user namespaces, which was first introduced
	// in kernel v5.11.
	//
	// In place of the directory path a ".v/" versioned directory may be specified, see systemd.v for details.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 251.
	ExtensionDirectories systemdconf.Value

	// Pass a credential to the unit. Credentials are limited-size binary or textual objects that may be passed to unit processes.
	// They are primarily used for passing cryptographic keys (both public and private) or certificates, user account information
	// or identity information from host to services. The data is accessible from the unit's processes via the file system, at
	// a read-only location that (if possible and permitted) is backed by non-swappable memory. The data is only accessible to
	// the user associated with the unit, via the User=/DynamicUser= settings (as well as the superuser). When available, the
	// location of credentials is exported as the $CREDENTIALS_DIRECTORY environment variable to the unit's processes.
	//
	// The LoadCredential= setting takes a textual ID to use as name for a credential plus a file system path, separated by a colon.
	// The ID must be a short ASCII string suitable as filename in the filesystem, and may be chosen freely by the user. If the specified
	// path is absolute it is opened as regular file and the credential data is read from it. If the absolute path refers to an AF_UNIX
	// stream socket in the file system a connection is made to it (only once at unit start-up) and the credential data read from
	// the connection, providing an easy IPC integration point for dynamically transferring credentials from other services.
	//
	// If the specified path is not absolute and itself qualifies as valid credential identifier it is attempted to find a credential
	// that the service manager itself received under the specified name — which may be used to propagate credentials from an invoking
	// environment (e.g. a container manager that invoked the service manager) into a service. If no matching system credential
	// is found, the directories /etc/credstore/, /run/credstore/ and /usr/lib/credstore/ are searched for files under the
	// credential's name — which hence are recommended locations for credential data on disk. If LoadCredentialEncrypted=
	// is used /run/credstore.encrypted/, /etc/credstore.encrypted/, and /usr/lib/credstore.encrypted/ are searched
	// as well.
	//
	// If the file system path is omitted it is chosen identical to the credential name, i.e. this is a terse way to declare credentials
	// to inherit from the service manager into a service. This option may be used multiple times, each time defining an additional
	// credential to pass to the unit.
	//
	// Note that if the path is not specified or a valid credential identifier is given, i.e. in the above two cases, a missing credential
	// is not considered fatal.
	//
	// If an absolute path referring to a directory is specified, every file in that directory (recursively) will be loaded as
	// a separate credential. The ID for each credential will be the provided ID suffixed with "_$FILENAME" (e.g., "Key_file1").
	// When loading from a directory, symlinks will be ignored.
	//
	// The contents of the file/socket may be arbitrary binary or textual data, including newline characters and NUL bytes.
	//
	// The LoadCredentialEncrypted= setting is identical to LoadCredential=, except that the credential data is decrypted
	// and authenticated before being passed on to the executed processes. Specifically, the referenced path should refer to
	// a file or socket with an encrypted credential, as implemented by systemd-creds. This credential is loaded, decrypted,
	// authenticated and then passed to the application in plaintext form, in the same way a regular credential specified via
	// LoadCredential= would be. A credential configured this way may be symmetrically encrypted/authenticated with a secret
	// key derived from the system's TPM2 security chip, or with a secret key stored in /var/lib/systemd/credentials.secret,
	// or with both. Using encrypted and authenticated credentials improves security as credentials are not stored in plaintext
	// and only authenticated and decrypted into plaintext the moment a service requiring them is started. Moreover, credentials
	// may be bound to the local hardware and installations, so that they cannot easily be analyzed offline, or be generated externally.
	// When DevicePolicy= is set to "closed" or "strict", or set to "auto" and DeviceAllow= is set, or PrivateDevices= is set,
	// then this setting adds /dev/tpmrm0 with rw mode to DeviceAllow=. See systemd.resource-control for the details about
	// DevicePolicy= or DeviceAllow=.
	//
	// Note that encrypted credentials targeted for services of the per-user service manager must be encrypted with systemd-creds
	// encrypt --user, and those for the system service manager without the --user switch. Encrypted credentials are always
	// targeted to a specific user or the system as a whole, and it is ensured that per-user service managers cannot decrypt secrets
	// intended for the system or for other users.
	//
	// The credential files/IPC sockets must be accessible to the service manager, but do not have to be directly accessible to
	// the unit's processes: the credential data is read and copied into separate, read-only copies for the unit that are accessible
	// to appropriately privileged processes. This is particularly useful in combination with DynamicUser= as this way privileged
	// data can be made available to processes running under a dynamic UID (i.e. not a previously known one) without having to open
	// up access to all users.
	//
	// In order to reference the path a credential may be read from within a ExecStart= command line use "${CREDENTIALS_DIRECTORY}/mycred",
	// e.g. "ExecStart=cat ${CREDENTIALS_DIRECTORY}/mycred". In order to reference the path a credential may be read from
	// within a Environment= line use "%d/mycred", e.g. "Environment=MYCREDPATH=%d/mycred". For system services the path
	// may also be referenced as "/run/credentials/UNITNAME" in cases where no interpolation is possible, e.g. configuration
	// files of software that does not yet support credentials natively. $CREDENTIALS_DIRECTORY is considered the primary
	// interface to look for credentials, though, since it also works for user services.
	//
	// Currently, an accumulated credential size limit of 1 MB per unit is enforced.
	//
	// The service manager itself may receive system credentials that can be propagated to services from a hosting container
	// manager or VM hypervisor. See the Container Interface documentation for details about the former. For the latter, pass
	// DMI/SMBIOS OEM string table entries (field type 11) with a prefix of "io.systemd.credential:" or "io.systemd.credential.binary:".
	// In both cases a key/value pair separated by "=" is expected, in the latter case the right-hand side is Base64 decoded when
	// parsed (thus permitting binary data to be passed in). Example qemu switch: "-smbios type=11,value=io.systemd.credential:xx=yy",
	// or "-smbios type=11,value=io.systemd.credential.binary:rick=TmV2ZXIgR29ubmEgR2l2ZSBZb3UgVXA=". Alternatively,
	// use the qemu "fw_cfg" node "opt/io.systemd.credentials/". Example qemu switch: "-fw_cfg name=opt/io.systemd.credentials/mycred,string=supersecret".
	// They may also be passed from the UEFI firmware environment via systemd-stub, from the initrd (see systemd), or be specified
	// on the kernel command line using the "systemd.set_credential=" and "systemd.set_credential_binary=" switches (see
	// systemd – this is not recommended since unprivileged userspace can read the kernel command line).
	//
	// If referencing an AF_UNIX stream socket to connect to, the connection will originate from an abstract namespace socket,
	// that includes information about the unit and the credential ID in its socket name. Use getpeername to query this information.
	// The returned socket name is formatted as NUL RANDOM "/unit/" UNIT "/" ID, i.e. a NUL byte (as required for abstract namespace
	// socket names), followed by a random string (consisting of alphadecimal characters), followed by the literal string "/unit/",
	// followed by the requesting unit name, followed by the literal character "/", followed by the textual credential ID requested.
	// Example: "\0adf9d86b6eda275e/unit/foobar.service/credx" in case the credential "credx" is requested for a unit "foobar.service".
	// This functionality is useful for using a single listening socket to serve credentials to multiple consumers.
	//
	// For further information see System and Service Credentials documentation.
	//
	// Added in version 247.
	LoadCredential systemdconf.Value

	// Pass a credential to the unit. Credentials are limited-size binary or textual objects that may be passed to unit processes.
	// They are primarily used for passing cryptographic keys (both public and private) or certificates, user account information
	// or identity information from host to services. The data is accessible from the unit's processes via the file system, at
	// a read-only location that (if possible and permitted) is backed by non-swappable memory. The data is only accessible to
	// the user associated with the unit, via the User=/DynamicUser= settings (as well as the superuser). When available, the
	// location of credentials is exported as the $CREDENTIALS_DIRECTORY environment variable to the unit's processes.
	//
	// The LoadCredential= setting takes a textual ID to use as name for a credential plus a file system path, separated by a colon.
	// The ID must be a short ASCII string suitable as filename in the filesystem, and may be chosen freely by the user. If the specified
	// path is absolute it is opened as regular file and the credential data is read from it. If the absolute path refers to an AF_UNIX
	// stream socket in the file system a connection is made to it (only once at unit start-up) and the credential data read from
	// the connection, providing an easy IPC integration point for dynamically transferring credentials from other services.
	//
	// If the specified path is not absolute and itself qualifies as valid credential identifier it is attempted to find a credential
	// that the service manager itself received under the specified name — which may be used to propagate credentials from an invoking
	// environment (e.g. a container manager that invoked the service manager) into a service. If no matching system credential
	// is found, the directories /etc/credstore/, /run/credstore/ and /usr/lib/credstore/ are searched for files under the
	// credential's name — which hence are recommended locations for credential data on disk. If LoadCredentialEncrypted=
	// is used /run/credstore.encrypted/, /etc/credstore.encrypted/, and /usr/lib/credstore.encrypted/ are searched
	// as well.
	//
	// If the file system path is omitted it is chosen identical to the credential name, i.e. this is a terse way to declare credentials
	// to inherit from the service manager into a service. This option may be used multiple times, each time defining an additional
	// credential to pass to the unit.
	//
	// Note that if the path is not specified or a valid credential identifier is given, i.e. in the above two cases, a missing credential
	// is not considered fatal.
	//
	// If an absolute path referring to a directory is specified, every file in that directory (recursively) will be loaded as
	// a separate credential. The ID for each credential will be the provided ID suffixed with "_$FILENAME" (e.g., "Key_file1").
	// When loading from a directory, symlinks will be ignored.
	//
	// The contents of the file/socket may be arbitrary binary or textual data, including newline characters and NUL bytes.
	//
	// The LoadCredentialEncrypted= setting is identical to LoadCredential=, except that the credential data is decrypted
	// and authenticated before being passed on to the executed processes. Specifically, the referenced path should refer to
	// a file or socket with an encrypted credential, as implemented by systemd-creds. This credential is loaded, decrypted,
	// authenticated and then passed to the application in plaintext form, in the same way a regular credential specified via
	// LoadCredential= would be. A credential configured this way may be symmetrically encrypted/authenticated with a secret
	// key derived from the system's TPM2 security chip, or with a secret key stored in /var/lib/systemd/credentials.secret,
	// or with both. Using encrypted and authenticated credentials improves security as credentials are not stored in plaintext
	// and only authenticated and decrypted into plaintext the moment a service requiring them is started. Moreover, credentials
	// may be bound to the local hardware and installations, so that they cannot easily be analyzed offline, or be generated externally.
	// When DevicePolicy= is set to "closed" or "strict", or set to "auto" and DeviceAllow= is set, or PrivateDevices= is set,
	// then this setting adds /dev/tpmrm0 with rw mode to DeviceAllow=. See systemd.resource-control for the details about
	// DevicePolicy= or DeviceAllow=.
	//
	// Note that encrypted credentials targeted for services of the per-user service manager must be encrypted with systemd-creds
	// encrypt --user, and those for the system service manager without the --user switch. Encrypted credentials are always
	// targeted to a specific user or the system as a whole, and it is ensured that per-user service managers cannot decrypt secrets
	// intended for the system or for other users.
	//
	// The credential files/IPC sockets must be accessible to the service manager, but do not have to be directly accessible to
	// the unit's processes: the credential data is read and copied into separate, read-only copies for the unit that are accessible
	// to appropriately privileged processes. This is particularly useful in combination with DynamicUser= as this way privileged
	// data can be made available to processes running under a dynamic UID (i.e. not a previously known one) without having to open
	// up access to all users.
	//
	// In order to reference the path a credential may be read from within a ExecStart= command line use "${CREDENTIALS_DIRECTORY}/mycred",
	// e.g. "ExecStart=cat ${CREDENTIALS_DIRECTORY}/mycred". In order to reference the path a credential may be read from
	// within a Environment= line use "%d/mycred", e.g. "Environment=MYCREDPATH=%d/mycred". For system services the path
	// may also be referenced as "/run/credentials/UNITNAME" in cases where no interpolation is possible, e.g. configuration
	// files of software that does not yet support credentials natively. $CREDENTIALS_DIRECTORY is considered the primary
	// interface to look for credentials, though, since it also works for user services.
	//
	// Currently, an accumulated credential size limit of 1 MB per unit is enforced.
	//
	// The service manager itself may receive system credentials that can be propagated to services from a hosting container
	// manager or VM hypervisor. See the Container Interface documentation for details about the former. For the latter, pass
	// DMI/SMBIOS OEM string table entries (field type 11) with a prefix of "io.systemd.credential:" or "io.systemd.credential.binary:".
	// In both cases a key/value pair separated by "=" is expected, in the latter case the right-hand side is Base64 decoded when
	// parsed (thus permitting binary data to be passed in). Example qemu switch: "-smbios type=11,value=io.systemd.credential:xx=yy",
	// or "-smbios type=11,value=io.systemd.credential.binary:rick=TmV2ZXIgR29ubmEgR2l2ZSBZb3UgVXA=". Alternatively,
	// use the qemu "fw_cfg" node "opt/io.systemd.credentials/". Example qemu switch: "-fw_cfg name=opt/io.systemd.credentials/mycred,string=supersecret".
	// They may also be passed from the UEFI firmware environment via systemd-stub, from the initrd (see systemd), or be specified
	// on the kernel command line using the "systemd.set_credential=" and "systemd.set_credential_binary=" switches (see
	// systemd – this is not recommended since unprivileged userspace can read the kernel command line).
	//
	// If referencing an AF_UNIX stream socket to connect to, the connection will originate from an abstract namespace socket,
	// that includes information about the unit and the credential ID in its socket name. Use getpeername to query this information.
	// The returned socket name is formatted as NUL RANDOM "/unit/" UNIT "/" ID, i.e. a NUL byte (as required for abstract namespace
	// socket names), followed by a random string (consisting of alphadecimal characters), followed by the literal string "/unit/",
	// followed by the requesting unit name, followed by the literal character "/", followed by the textual credential ID requested.
	// Example: "\0adf9d86b6eda275e/unit/foobar.service/credx" in case the credential "credx" is requested for a unit "foobar.service".
	// This functionality is useful for using a single listening socket to serve credentials to multiple consumers.
	//
	// For further information see System and Service Credentials documentation.
	//
	// Added in version 247.
	LoadCredentialEncrypted systemdconf.Value

	// Pass one or more credentials to the unit. Takes a credential name for which we will attempt to find a credential that the service
	// manager itself received under the specified name — which may be used to propagate credentials from an invoking environment
	// (e.g. a container manager that invoked the service manager) into a service. If the credential name is a glob, all credentials
	// matching the glob are passed to the unit. Matching credentials are searched for in the system credentials, the encrypted
	// system credentials, and under /etc/credstore/, /run/credstore/, /usr/lib/credstore/, /run/credstore.encrypted/,
	// /etc/credstore.encrypted/, and /usr/lib/credstore.encrypted/ in that order. When multiple credentials of the same
	// name are found, the first one found is used.
	//
	// The globbing expression implements a restrictive subset of glob: only a single trailing "*" wildcard may be specified.
	// Both "?" and "[]" wildcards are not permitted, nor are "*" wildcards anywhere except at the end of the glob expression.
	//
	// Optionally, the credential name or glob may be followed by a colon followed by a rename pattern. If specified, all credentials
	// matching the credential name or glob are renamed according to the given pattern. For example, if "ImportCredential=my.original.cred:my.renamed.cred"
	// is specified, the service manager will read the "my.original.cred" credential and make it available as the "my.renamed.cred"
	// credential to the service. Similarly, if "ImportCredential=my.original.*:my.renamed." is specified, the service
	// manager will read all credentials starting with "my.original." and make them available as "my.renamed.xxx" to the service.
	//
	// If ImportCredential= is specified multiple times and multiple credentials end up with the same name after renaming, the
	// first one is kept and later ones are dropped.
	//
	// When multiple credentials of the same name are found, credentials found by LoadCredential= and LoadCredentialEncrypted=
	// take priority over credentials found by ImportCredential=.
	//
	// Added in version 254.
	ImportCredential systemdconf.Value

	// The SetCredential= setting is similar to LoadCredential= but accepts a literal value to use as data for the credential,
	// instead of a file system path to read the data from. Do not use this option for data that is supposed to be secret, as it is accessible
	// to unprivileged processes via IPC. It's only safe to use this for user IDs, public key material and similar non-sensitive
	// data. For everything else use LoadCredential=. In order to embed binary data into the credential data use C-style escaping
	// (i.e. "\n" to embed a newline, or "\x00" to embed a NUL byte).
	//
	// The SetCredentialEncrypted= setting is identical to SetCredential= but expects an encrypted credential in literal
	// form as value. This allows embedding confidential credentials securely directly in unit files. Use systemd-creds' -p
	// switch to generate suitable SetCredentialEncrypted= lines directly from plaintext credentials. For further details
	// see LoadCredentialEncrypted= above.
	//
	// When multiple credentials of the same name are found, credentials found by LoadCredential=, LoadCredentialEncrypted=
	// and ImportCredential= take priority over credentials found by SetCredential=. As such, SetCredential= will act as default
	// if no credentials are found by any of the former. In this case not being able to retrieve the credential from the path specified
	// in LoadCredential= or LoadCredentialEncrypted= is not considered fatal.
	//
	// Added in version 247.
	SetCredential systemdconf.Value

	// The SetCredential= setting is similar to LoadCredential= but accepts a literal value to use as data for the credential,
	// instead of a file system path to read the data from. Do not use this option for data that is supposed to be secret, as it is accessible
	// to unprivileged processes via IPC. It's only safe to use this for user IDs, public key material and similar non-sensitive
	// data. For everything else use LoadCredential=. In order to embed binary data into the credential data use C-style escaping
	// (i.e. "\n" to embed a newline, or "\x00" to embed a NUL byte).
	//
	// The SetCredentialEncrypted= setting is identical to SetCredential= but expects an encrypted credential in literal
	// form as value. This allows embedding confidential credentials securely directly in unit files. Use systemd-creds' -p
	// switch to generate suitable SetCredentialEncrypted= lines directly from plaintext credentials. For further details
	// see LoadCredentialEncrypted= above.
	//
	// When multiple credentials of the same name are found, credentials found by LoadCredential=, LoadCredentialEncrypted=
	// and ImportCredential= take priority over credentials found by SetCredential=. As such, SetCredential= will act as default
	// if no credentials are found by any of the former. In this case not being able to retrieve the credential from the path specified
	// in LoadCredential= or LoadCredentialEncrypted= is not considered fatal.
	//
	// Added in version 247.
	SetCredentialEncrypted systemdconf.Value

	// Controls which capabilities to include in the capability bounding set for the executed process. See capabilities for
	// details. Takes a whitespace-separated list of capability names, e.g. CAP_SYS_ADMIN, CAP_DAC_OVERRIDE, CAP_SYS_PTRACE.
	// Capabilities listed will be included in the bounding set, all others are removed. If the list of capabilities is prefixed
	// with "~", all but the listed capabilities will be included, the effect of the assignment inverted. Note that this option
	// also affects the respective capabilities in the effective, permitted and inheritable capability sets. If this option
	// is not used, the capability bounding set is not modified on process execution, hence no limits on the capabilities of the
	// process are enforced. This option may appear more than once, in which case the bounding sets are merged by OR, or by AND if
	// the lines are prefixed with "~" (see below). If the empty string is assigned to this option, the bounding set is reset to the
	// empty capability set, and all prior settings have no effect. If set to "~" (without any further argument), the bounding
	// set is reset to the full set of available capabilities, also undoing any previous settings. This does not affect commands
	// prefixed with "+".
	//
	// Use systemd-analyze's capability command to retrieve a list of capabilities defined on the local system.
	//
	// Example: if a unit has the following,
	//
	//	CapabilityBoundingSet=CAP_A CAP_B CapabilityBoundingSet=CAP_B CAP_C
	//
	// then CAP_A, CAP_B, and CAP_C are set. If the second line is prefixed with "~", e.g.,
	//
	//	CapabilityBoundingSet=CAP_A CAP_B CapabilityBoundingSet=~CAP_B CAP_C
	//
	// then, only CAP_A is set.
	CapabilityBoundingSet systemdconf.Value

	// Controls which capabilities to include in the ambient capability set for the executed process. Takes a whitespace-separated
	// list of capability names, e.g. CAP_SYS_ADMIN, CAP_DAC_OVERRIDE, CAP_SYS_PTRACE. This option may appear more than once,
	// in which case the ambient capability sets are merged (see the above examples in CapabilityBoundingSet=). If the list of
	// capabilities is prefixed with "~", all but the listed capabilities will be included, the effect of the assignment inverted.
	// If the empty string is assigned to this option, the ambient capability set is reset to the empty capability set, and all prior
	// settings have no effect. If set to "~" (without any further argument), the ambient capability set is reset to the full set
	// of available capabilities, also undoing any previous settings. Note that adding capabilities to the ambient capability
	// set adds them to the process's inherited capability set.
	//
	// Ambient capability sets are useful if you want to execute a process as a non-privileged user but still want to give it some
	// capabilities. Note that in this case option keep-caps is automatically added to SecureBits= to retain the capabilities
	// over the user change. AmbientCapabilities= does not affect commands prefixed with "+".
	//
	// Added in version 229.
	AmbientCapabilities systemdconf.Value

	// Takes a boolean argument. If true, ensures that the service process and all its children can never gain new privileges through
	// execve() (e.g. via setuid or setgid bits, or filesystem capabilities). This is the simplest and most effective way to ensure
	// that a process and its children can never elevate privileges again. Defaults to false. In case the service will be run in
	// a new mount namespace anyway and SELinux is disabled, all file systems are mounted with MS_NOSUID flag. Also see No New Privileges
	// Flag.
	//
	// Note that this setting only has an effect on the unit's processes themselves (or any processes directly or indirectly forked
	// off them). It has no effect on processes potentially invoked on request of them through tools such as at, crontab, systemd-run,
	// or arbitrary IPC services.
	//
	// Added in version 187.
	NoNewPrivileges systemdconf.Value

	// Controls the secure bits set for the executed process. Takes a space-separated combination of options from the following
	// list: keep-caps, keep-caps-locked, no-setuid-fixup, no-setuid-fixup-locked, noroot, and noroot-locked. This option
	// may appear more than once, in which case the secure bits are ORed. If the empty string is assigned to this option, the bits
	// are reset to 0. This does not affect commands prefixed with "+". See capabilities for details.
	SecureBits systemdconf.Value

	// Set the SELinux security context of the executed process. If set, this will override the automated domain transition.
	// However, the policy still needs to authorize the transition. This directive is ignored if SELinux is disabled. If prefixed
	// by "-", failing to set the SELinux security context will be ignored, but it is still possible that the subsequent execve()
	// may fail if the policy does not allow the transition for the non-overridden context. This does not affect commands prefixed
	// with "+". See setexeccon for details.
	//
	// Added in version 209.
	SELinuxContext systemdconf.Value

	// Takes a profile name as argument. The process executed by the unit will switch to this profile when started. Profiles must
	// already be loaded in the kernel, or the unit will fail. If prefixed by "-", all errors will be ignored. This setting has no
	// effect if AppArmor is not enabled. This setting does not affect commands prefixed with "+".
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 210.
	AppArmorProfile systemdconf.Value

	// Takes a SMACK64 security label as argument. The process executed by the unit will be started under this label and SMACK will
	// decide whether the process is allowed to run or not, based on it. The process will continue to run under the label specified
	// here unless the executable has its own SMACK64EXEC label, in which case the process will transition to run under that label.
	// When not specified, the label that systemd is running under is used. This directive is ignored if SMACK is disabled.
	//
	// The value may be prefixed by "-", in which case all errors will be ignored. An empty value may be specified to unset previous
	// assignments. This does not affect commands prefixed with "+".
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 218.
	SmackProcessLabel systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitCPU systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitFSIZE systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitDATA systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitSTACK systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitCORE systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitRSS systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitNOFILE systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitAS systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitNPROC systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitMEMLOCK systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitLOCKS systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitSIGPENDING systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitMSGQUEUE systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitNICE systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitRTPRIO systemdconf.Value

	// Set soft and hard limits on various resources for executed processes. See setrlimit for details on the process resource
	// limit concept. Process resource limits may be specified in two formats: either as single value to set a specific soft and
	// hard limit to the same value, or as colon-separated pair soft:hard to set both limits individually (e.g. "LimitAS=4G:16G").
	// Use the string infinity to configure no limit on a specific resource. The multiplicative suffixes K, M, G, T, P and E (to the
	// base 1024) may be used for resource limits measured in bytes (e.g. "LimitAS=16G"). For the limits referring to time values,
	// the usual time units ms, s, min, h and so on may be used (see systemd.time for details). Note that if no time unit is specified
	// for LimitCPU= the default unit of seconds is implied, while for LimitRTTIME= the default unit of microseconds is implied.
	// Also, note that the effective granularity of the limits might influence their enforcement. For example, time limits specified
	// for LimitCPU= will be rounded up implicitly to multiples of 1s. For LimitNICE= the value may be specified in two syntaxes:
	// if prefixed with "+" or "-", the value is understood as regular Linux nice value in the range -20…19. If not prefixed like
	// this the value is understood as raw resource limit parameter in the range 0…40 (with 0 being equivalent to 1).
	//
	// Note that most process resource limits configured with these options are per-process, and processes may fork in order
	// to acquire a new set of resources that are accounted independently of the original process, and may thus escape limits set.
	// Also note that LimitRSS= is not implemented on Linux, and setting it has no effect. Often it is advisable to prefer the resource
	// controls listed in systemd.resource-control over these per-process limits, as they apply to services as a whole, may
	// be altered dynamically at runtime, and are generally more expressive. For example, MemoryMax= is a more powerful (and
	// working) replacement for LimitRSS=.
	//
	// Note that LimitNPROC= will limit the number of processes from one (real) UID and not the number of processes started (forked)
	// by the service. Therefore the limit is cumulative for all processes running under the same UID. Please also note that the
	// LimitNPROC= will not be enforced if the service is running as root (and not dropping privileges). Due to these limitations,
	// TasksMax= (see systemd.resource-control) is typically a better choice than LimitNPROC=.
	//
	// Resource limits not configured explicitly for a unit default to the value configured in the various DefaultLimitCPU=,
	// DefaultLimitFSIZE=, … options available in systemd-system.conf, and – if not configured there – the kernel or per-user
	// defaults, as defined by the OS (the latter only for user services, see below).
	//
	// For system units these resource limits may be chosen freely. When these settings are configured in a user service (i.e.
	// a service run by the per-user instance of the service manager) they cannot be used to raise the limits above those set for
	// the user manager itself when it was first invoked, as the user's service manager generally lacks the privileges to do so.
	// In user context these configuration options are hence only useful to lower the limits passed in or to raise the soft limit
	// to the maximum of the hard limit as configured for the user. To raise the user's limits further, the available configuration
	// mechanisms differ between operating systems, but typically require privileges. In most cases it is possible to configure
	// higher per-user resource limits via PAM or by setting limits on the system service encapsulating the user's service manager,
	// i.e. the user's instance of user@.service. After making such changes, make sure to restart the user's service manager.
	//
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	|    DIRECTIVE     | ULIMIT EQUIVALENT |            UNIT            |             NOTES              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	//	| LimitCPU=        | ulimit -t         | Seconds                    | -                              |
	//	| LimitFSIZE=      | ulimit -f         | Bytes                      | -                              |
	//	| LimitDATA=       | ulimit -d         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitSTACK=      | ulimit -s         | Bytes                      | -                              |
	//	| LimitCORE=       | ulimit -c         | Bytes                      | -                              |
	//	| LimitRSS=        | ulimit -m         | Bytes                      | Do not use. No effect on       |
	//	|                  |                   |                            | Linux.                         |
	//	| LimitNOFILE=     | ulimit -n         | Number of File Descriptors | Do not use. Be careful         |
	//	|                  |                   |                            | when raising the soft limit    |
	//	|                  |                   |                            | above 1024, since select       |
	//	|                  |                   |                            | cannot function with file      |
	//	|                  |                   |                            | descriptors above 1023 on      |
	//	|                  |                   |                            | Linux. Nowadays, the hard      |
	//	|                  |                   |                            | limit defaults to 524288, a    |
	//	|                  |                   |                            | very high value compared to    |
	//	|                  |                   |                            | historical defaults. Typically |
	//	|                  |                   |                            | applications should increase   |
	//	|                  |                   |                            | their soft limit to the hard   |
	//	|                  |                   |                            | limit on their own, if they    |
	//	|                  |                   |                            | are OK with working with file  |
	//	|                  |                   |                            | descriptors above 1023, i.e.   |
	//	|                  |                   |                            | do not use select. Note that   |
	//	|                  |                   |                            | file descriptors are nowadays  |
	//	|                  |                   |                            | accounted like any other form  |
	//	|                  |                   |                            | of memory, thus there should   |
	//	|                  |                   |                            | not be any need to lower the   |
	//	|                  |                   |                            | hard limit. Use MemoryMax= to  |
	//	|                  |                   |                            | control overall service memory |
	//	|                  |                   |                            | use, including file descriptor |
	//	|                  |                   |                            | memory.                        |
	//	| LimitAS=         | ulimit -v         | Bytes                      | Do not use. This limits        |
	//	|                  |                   |                            | the allowed address range,     |
	//	|                  |                   |                            | not memory use! Defaults       |
	//	|                  |                   |                            | to unlimited and should not    |
	//	|                  |                   |                            | be lowered. To limit memory    |
	//	|                  |                   |                            | use, see MemoryMax= in         |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitNPROC=      | ulimit -u         | Number of Processes        | This limit is enforced based   |
	//	|                  |                   |                            | on the number of processes     |
	//	|                  |                   |                            | belonging to the user.         |
	//	|                  |                   |                            | Typically it is better to      |
	//	|                  |                   |                            | track processes per service,   |
	//	|                  |                   |                            | i.e. use TasksMax=, see        |
	//	|                  |                   |                            | systemd.resource-control.      |
	//	| LimitMEMLOCK=    | ulimit -l         | Bytes                      | -                              |
	//	| LimitLOCKS=      | ulimit -x         | Number of Locks            | -                              |
	//	| LimitSIGPENDING= | ulimit -i         | Number of Queued Signals   | -                              |
	//	| LimitMSGQUEUE=   | ulimit -q         | Bytes                      | -                              |
	//	| LimitNICE=       | ulimit -e         | Nice Level                 | -                              |
	//	| LimitRTPRIO=     | ulimit -r         | Realtime Priority          | -                              |
	//	| LimitRTTIME=     | ulimit -R         | Microseconds               | -                              |
	//	+------------------+-------------------+----------------------------+--------------------------------+
	LimitRTTIME systemdconf.Value

	// Controls the file mode creation mask. Takes an access mode in octal notation. See umask for details. Defaults to 0022 for
	// system units. For user units the default value is inherited from the per-user service manager (whose default is in turn
	// inherited from the system service manager, and thus typically also is 0022 — unless overridden by a PAM module). In order
	// to change the per-user mask for all user services, consider setting the UMask= setting of the user's user@.service system
	// service instance. The per-user umask may also be set via the umask field of a user's JSON User Record (for users managed by
	// systemd-homed.service this field may be controlled via homectl --umask=). It may also be set via a PAM module, such as pam_umask.
	UMask systemdconf.Value

	// Controls which types of memory mappings will be saved if the process dumps core (using the /proc/pid/coredump_filter
	// file). Takes a whitespace-separated combination of mapping type names or numbers (with the default base 16). Mapping
	// type names are private-anonymous, shared-anonymous, private-file-backed, shared-file-backed, elf-headers, private-huge,
	// shared-huge, private-dax, shared-dax, and the special values all (all types) and default (the kernel default of "private-anonymous
	// shared-anonymous elf-headers private-huge"). See core for the meaning of the mapping types. When specified multiple
	// times, all specified masks are ORed. When not set, or if the empty value is assigned, the inherited value is not changed.
	//
	//	CoredumpFilter=default private-dax shared-dax
	//
	// Added in version 246.
	CoredumpFilter systemdconf.Value

	// Controls how the kernel session keyring is set up for the service (see session-keyring for details on the session keyring).
	// Takes one of inherit, private, shared. If set to inherit no special keyring setup is done, and the kernel's default behaviour
	// is applied. If private is used a new session keyring is allocated when a service process is invoked, and it is not linked up
	// with any user keyring. This is the recommended setting for system services, as this ensures that multiple services running
	// under the same system user ID (in particular the root user) do not share their key material among each other. If shared is
	// used a new session keyring is allocated as for private, but the user keyring of the user configured with User= is linked into
	// it, so that keys assigned to the user may be requested by the unit's processes. In this mode multiple units running processes
	// under the same user ID may share key material. Unless inherit is selected the unique invocation ID for the unit (see below)
	// is added as a protected key by the name "invocation_id" to the newly created session keyring. Defaults to private for services
	// of the system service manager and to inherit for non-service units and for services of the user service manager.
	//
	// Added in version 235.
	KeyringMode systemdconf.Value

	// Sets the adjustment value for the Linux kernel's Out-Of-Memory (OOM) killer score for executed processes. Takes an integer
	// between -1000 (to disable OOM killing of processes of this unit) and 1000 (to make killing of processes of this unit under
	// memory pressure very likely). See The /proc Filesystem for details. If not specified defaults to the OOM score adjustment
	// level of the service manager itself, which is normally at 0.
	//
	// Use the OOMPolicy= setting of service units to configure how the service manager shall react to the kernel OOM killer or
	// systemd-oomd terminating a process of the service. See systemd.service for details.
	OOMScoreAdjust systemdconf.Value

	// Sets the timer slack in nanoseconds for the executed processes. The timer slack controls the accuracy of wake-ups triggered
	// by timers. See prctl for more information. Note that in contrast to most other time span definitions this parameter takes
	// an integer value in nano-seconds if no unit is specified. The usual time units are understood too.
	TimerSlackNSec systemdconf.Value

	// Controls which kernel architecture uname shall report, when invoked by unit processes. Takes one of the architecture
	// identifiers arm64, arm64-be, arm, arm-be, x86, x86-64, ppc, ppc-le, ppc64, ppc64-le, s390 or s390x. Which personality
	// architectures are supported depends on the kernel's native architecture. Usually the 64-bit versions of the various
	// system architectures support their immediate 32-bit personality architecture counterpart, but no others. For example,
	// x86-64 systems support the x86-64 and x86 personalities but no others. The personality feature is useful when running
	// 32-bit services on a 64-bit host system. If not specified, the personality is left unmodified and thus reflects the personality
	// of the host system's kernel. This option is not useful on architectures for which only one native word width was ever available,
	// such as m68k (32-bit only) or alpha (64-bit only).
	//
	// Added in version 209.
	Personality systemdconf.Value

	// Takes a boolean argument. If true, SIGPIPE is ignored in the executed process. Defaults to true since SIGPIPE is generally
	// only useful in shell pipelines.
	IgnoreSIGPIPE systemdconf.Value

	// Sets the default nice level (scheduling priority) for executed processes. Takes an integer between -20 (highest priority)
	// and 19 (lowest priority). In case of resource contention, smaller values mean more resources will be made available to
	// the unit's processes, larger values mean less resources will be made available. See setpriority for details.
	Nice systemdconf.Value

	// Sets the CPU scheduling policy for executed processes. Takes one of other, batch, idle, fifo or rr. See sched_setscheduler
	// for details.
	CPUSchedulingPolicy systemdconf.Value

	// Sets the CPU scheduling priority for executed processes. The available priority range depends on the selected CPU scheduling
	// policy (see above). For real-time scheduling policies an integer between 1 (lowest priority) and 99 (highest priority)
	// can be used. In case of CPU resource contention, smaller values mean less CPU time is made available to the service, larger
	// values mean more. See sched_setscheduler for details.
	CPUSchedulingPriority systemdconf.Value

	// Takes a boolean argument. If true, elevated CPU scheduling priorities and policies will be reset when the executed processes
	// call fork, and can hence not leak into child processes. See sched_setscheduler for details. Defaults to false.
	CPUSchedulingResetOnFork systemdconf.Value

	// Controls the CPU affinity of the executed processes. Takes a list of CPU indices or ranges separated by either whitespace
	// or commas. Alternatively, takes a special "numa" value in which case systemd automatically derives allowed CPU range
	// based on the value of NUMAMask= option. CPU ranges are specified by the lower and upper CPU indices separated by a dash. This
	// option may be specified more than once, in which case the specified CPU affinity masks are merged. If the empty string is
	// assigned, the mask is reset, all assignments prior to this will have no effect. See sched_setaffinity for details.
	CPUAffinity systemdconf.Value

	// Controls the NUMA memory policy of the executed processes. Takes a policy type, one of: default, preferred, bind, interleave
	// and local. A list of NUMA nodes that should be associated with the policy must be specified in NUMAMask=. For more details
	// on each policy please see, set_mempolicy. For overall overview of NUMA support in Linux see, numa.
	//
	// Added in version 243.
	NUMAPolicy systemdconf.Value

	// Controls the NUMA node list which will be applied alongside with selected NUMA policy. Takes a list of NUMA nodes and has
	// the same syntax as a list of CPUs for CPUAffinity= option or special "all" value which will include all available NUMA nodes
	// in the mask. Note that the list of NUMA nodes is not required for default and local policies and for preferred policy we expect
	// a single NUMA node.
	//
	// Added in version 243.
	NUMAMask systemdconf.Value

	// Sets the I/O scheduling class for executed processes. Takes one of the strings realtime, best-effort or idle. The kernel's
	// default scheduling class is best-effort at a priority of 4. If the empty string is assigned to this option, all prior assignments
	// to both IOSchedulingClass= and IOSchedulingPriority= have no effect. See ioprio_set for details.
	IOSchedulingClass systemdconf.Value

	// Sets the I/O scheduling priority for executed processes. Takes an integer between 0 (highest priority) and 7 (lowest priority).
	// In case of I/O contention, smaller values mean more I/O bandwidth is made available to the unit's processes, larger values
	// mean less bandwidth. The available priorities depend on the selected I/O scheduling class (see above). If the empty string
	// is assigned to this option, all prior assignments to both IOSchedulingClass= and IOSchedulingPriority= have no effect.
	// For the kernel's default scheduling class (best-effort) this defaults to 4. See ioprio_set for details.
	IOSchedulingPriority systemdconf.Value

	// Takes a boolean argument or the special values "full" or "strict". If true, mounts the /usr/ and the boot loader directories
	// (/boot and /efi) read-only for processes invoked by this unit. If set to "full", the /etc/ directory is mounted read-only,
	// too. If set to "strict" the entire file system hierarchy is mounted read-only, except for the API file system subtrees /dev/,
	// /proc/ and /sys/ (protect these directories using PrivateDevices=, ProtectKernelTunables=, ProtectControlGroups=).
	// This setting ensures that any modification of the vendor-supplied operating system (and optionally its configuration,
	// and local mounts) is prohibited for the service. It is recommended to enable this setting for all long-running services,
	// unless they are involved with system updates or need to modify the operating system in other ways. If this option is used,
	// ReadWritePaths= may be used to exclude specific directories from being made read-only. Similar, StateDirectory=, LogsDirectory=,
	// … and related directory settings (see below) also exclude the specific directories from the effect of ProtectSystem=.
	// This setting is implied if DynamicUser= is set. This setting cannot ensure protection in all cases. In general it has the
	// same limitations as ReadOnlyPaths=, see below. Defaults to off.
	//
	// Note that if ProtectSystem= is set to "strict" and PrivateTmp= is enabled, then /tmp/ and /var/tmp/ will be writable.
	//
	// Added in version 214.
	ProtectSystem systemdconf.Value

	// Takes a boolean argument or the special values "read-only" or "tmpfs". If true, the directories /home/, /root, and /run/user
	// are made inaccessible and empty for processes invoked by this unit. If set to "read-only", the three directories are made
	// read-only instead. If set to "tmpfs", temporary file systems are mounted on the three directories in read-only mode. The
	// value "tmpfs" is useful to hide home directories not relevant to the processes invoked by the unit, while still allowing
	// necessary directories to be made visible when listed in BindPaths= or BindReadOnlyPaths=.
	//
	// Setting this to "yes" is mostly equivalent to setting the three directories in InaccessiblePaths=. Similarly, "read-only"
	// is mostly equivalent to ReadOnlyPaths=, and "tmpfs" is mostly equivalent to TemporaryFileSystem= with ":ro".
	//
	// It is recommended to enable this setting for all long-running services (in particular network-facing ones), to ensure
	// they cannot get access to private user data, unless the services actually require access to the user's private data. This
	// setting is implied if DynamicUser= is set. This setting cannot ensure protection in all cases. In general it has the same
	// limitations as ReadOnlyPaths=, see below.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 214.
	ProtectHome systemdconf.Value

	// These options take a whitespace-separated list of directory names. The specified directory names must be relative, and
	// may not include "..". If set, when the unit is started, one or more directories by the specified names will be created (including
	// their parents) below the locations defined in the following table. Also, the corresponding environment variable will
	// be defined with the full paths of the directories. If multiple directories are set, then in the environment variable the
	// paths are concatenated with colon (":").
	//
	// If DynamicUser= is used, and if the kernel version supports id-mapped mounts, the specified directories will be owned
	// by "nobody" in the host namespace and will be mapped to (and will be owned by) the service's UID/GID in its own namespace.
	// For backward compatibility, existing directories created without id-mapped mounts will be kept untouched.
	//
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	|        DIRECTORY        | BELOW PATH FOR SYSTEM UNITS | BELOW PATH FOR USER UNITS | ENVIRONMENT VARIABLE SET |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	| RuntimeDirectory=       | /run/                       | $XDG_RUNTIME_DIR          | $RUNTIME_DIRECTORY       |
	//	| StateDirectory=         | /var/lib/                   | $XDG_STATE_HOME           | $STATE_DIRECTORY         |
	//	| CacheDirectory=         | /var/cache/                 | $XDG_CACHE_HOME           | $CACHE_DIRECTORY         |
	//	| LogsDirectory=          | /var/log/                   | $XDG_STATE_HOME/log/      | $LOGS_DIRECTORY          |
	//	| ConfigurationDirectory= | /etc/                       | $XDG_CONFIG_HOME          | $CONFIGURATION_DIRECTORY |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//
	// In case of RuntimeDirectory= the innermost subdirectories are removed when the unit is stopped. It is possible to preserve
	// the specified directories in this case if RuntimeDirectoryPreserve= is configured to restart or yes (see below). The
	// directories specified with StateDirectory=, CacheDirectory=, LogsDirectory=, ConfigurationDirectory= are not
	// removed when the unit is stopped.
	//
	// Except in case of ConfigurationDirectory=, the innermost specified directories will be owned by the user and group specified
	// in User= and Group=. If the specified directories already exist and their owning user or group do not match the configured
	// ones, all files and directories below the specified directories as well as the directories themselves will have their
	// file ownership recursively changed to match what is configured. As an optimization, if the specified directories are
	// already owned by the right user and group, files and directories below of them are left as-is, even if they do not match what
	// is requested. The innermost specified directories will have their access mode adjusted to the what is specified in RuntimeDirectoryMode=,
	// StateDirectoryMode=, CacheDirectoryMode=, LogsDirectoryMode= and ConfigurationDirectoryMode=.
	//
	// These options imply BindPaths= for the specified paths. When combined with RootDirectory= or RootImage= these paths
	// always reside on the host and are mounted from there into the unit's file system namespace.
	//
	// If DynamicUser= is used, the logic for CacheDirectory=, LogsDirectory= and StateDirectory= is slightly altered: the
	// directories are created below /var/cache/private, /var/log/private and /var/lib/private, respectively, which are
	// host directories made inaccessible to unprivileged users, which ensures that access to these directories cannot be gained
	// through dynamic user ID recycling. Symbolic links are created to hide this difference in behaviour. Both from perspective
	// of the host and from inside the unit, the relevant directories hence always appear directly below /var/cache, /var/log
	// and /var/lib.
	//
	// Use RuntimeDirectory= to manage one or more runtime directories for the unit and bind their lifetime to the daemon runtime.
	// This is particularly useful for unprivileged daemons that cannot create runtime directories in /run/ due to lack of privileges,
	// and to make sure the runtime directory is cleaned up automatically after use. For runtime directories that require more
	// complex or different configuration or lifetime guarantees, please consider using tmpfiles.d.
	//
	// RuntimeDirectory=, StateDirectory=, CacheDirectory= and LogsDirectory= optionally support two more parameters,
	// separated by ":". The second parameter will be interpreted as a destination path that will be created as a symlink to the
	// directory. The symlinks will be created after any BindPaths= or TemporaryFileSystem= options have been set up, to make
	// ephemeral symlinking possible. The same source can have multiple symlinks, by using the same first parameter, but a different
	// second parameter. The third parameter is a flags field, and since v257 can take a value of ro to make the directory read only
	// for the service. This is also supported for ConfigurationDirectory=. If multiple symlinks are set up, the directory will
	// be read only if at least one is configured to be read only. To pass a flag without a destination symlink, the second parameter
	// can be empty, for example:
	//
	//	ConfigurationDirectory=foo::ro
	//
	// The directories defined by these options are always created under the standard paths used by systemd (/var/, /run/, /etc/,
	// …). If the service needs directories in a different location, a different mechanism has to be used to create them.
	//
	// tmpfiles.d provides functionality that overlaps with these options. Using these options is recommended, because the
	// lifetime of the directories is tied directly to the lifetime of the unit, and it is not necessary to ensure that the tmpfiles.d
	// configuration is executed before the unit is started.
	//
	// To remove any of the directories created by these settings, use the systemctl clean … command on the relevant units, see
	// systemctl for details.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar baz
	//
	// the service manager creates /run/foo (if it does not exist), /run/foo/bar, and /run/baz. The directories /run/foo/bar
	// and /run/baz except /run/foo are owned by the user and group specified in User= and Group=, and removed when the service
	// is stopped.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar StateDirectory=aaa/bbb ccc
	//
	// then the environment variable "RUNTIME_DIRECTORY" is set with "/run/foo/bar", and "STATE_DIRECTORY" is set with "/var/lib/aaa/bbb:/var/lib/ccc".
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo:bar foo:baz
	//
	// the service manager creates /run/foo (if it does not exist), and /run/bar plus /run/baz as symlinks to /run/foo.
	//
	// Added in version 211.
	RuntimeDirectory systemdconf.Value

	// These options take a whitespace-separated list of directory names. The specified directory names must be relative, and
	// may not include "..". If set, when the unit is started, one or more directories by the specified names will be created (including
	// their parents) below the locations defined in the following table. Also, the corresponding environment variable will
	// be defined with the full paths of the directories. If multiple directories are set, then in the environment variable the
	// paths are concatenated with colon (":").
	//
	// If DynamicUser= is used, and if the kernel version supports id-mapped mounts, the specified directories will be owned
	// by "nobody" in the host namespace and will be mapped to (and will be owned by) the service's UID/GID in its own namespace.
	// For backward compatibility, existing directories created without id-mapped mounts will be kept untouched.
	//
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	|        DIRECTORY        | BELOW PATH FOR SYSTEM UNITS | BELOW PATH FOR USER UNITS | ENVIRONMENT VARIABLE SET |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	| RuntimeDirectory=       | /run/                       | $XDG_RUNTIME_DIR          | $RUNTIME_DIRECTORY       |
	//	| StateDirectory=         | /var/lib/                   | $XDG_STATE_HOME           | $STATE_DIRECTORY         |
	//	| CacheDirectory=         | /var/cache/                 | $XDG_CACHE_HOME           | $CACHE_DIRECTORY         |
	//	| LogsDirectory=          | /var/log/                   | $XDG_STATE_HOME/log/      | $LOGS_DIRECTORY          |
	//	| ConfigurationDirectory= | /etc/                       | $XDG_CONFIG_HOME          | $CONFIGURATION_DIRECTORY |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//
	// In case of RuntimeDirectory= the innermost subdirectories are removed when the unit is stopped. It is possible to preserve
	// the specified directories in this case if RuntimeDirectoryPreserve= is configured to restart or yes (see below). The
	// directories specified with StateDirectory=, CacheDirectory=, LogsDirectory=, ConfigurationDirectory= are not
	// removed when the unit is stopped.
	//
	// Except in case of ConfigurationDirectory=, the innermost specified directories will be owned by the user and group specified
	// in User= and Group=. If the specified directories already exist and their owning user or group do not match the configured
	// ones, all files and directories below the specified directories as well as the directories themselves will have their
	// file ownership recursively changed to match what is configured. As an optimization, if the specified directories are
	// already owned by the right user and group, files and directories below of them are left as-is, even if they do not match what
	// is requested. The innermost specified directories will have their access mode adjusted to the what is specified in RuntimeDirectoryMode=,
	// StateDirectoryMode=, CacheDirectoryMode=, LogsDirectoryMode= and ConfigurationDirectoryMode=.
	//
	// These options imply BindPaths= for the specified paths. When combined with RootDirectory= or RootImage= these paths
	// always reside on the host and are mounted from there into the unit's file system namespace.
	//
	// If DynamicUser= is used, the logic for CacheDirectory=, LogsDirectory= and StateDirectory= is slightly altered: the
	// directories are created below /var/cache/private, /var/log/private and /var/lib/private, respectively, which are
	// host directories made inaccessible to unprivileged users, which ensures that access to these directories cannot be gained
	// through dynamic user ID recycling. Symbolic links are created to hide this difference in behaviour. Both from perspective
	// of the host and from inside the unit, the relevant directories hence always appear directly below /var/cache, /var/log
	// and /var/lib.
	//
	// Use RuntimeDirectory= to manage one or more runtime directories for the unit and bind their lifetime to the daemon runtime.
	// This is particularly useful for unprivileged daemons that cannot create runtime directories in /run/ due to lack of privileges,
	// and to make sure the runtime directory is cleaned up automatically after use. For runtime directories that require more
	// complex or different configuration or lifetime guarantees, please consider using tmpfiles.d.
	//
	// RuntimeDirectory=, StateDirectory=, CacheDirectory= and LogsDirectory= optionally support two more parameters,
	// separated by ":". The second parameter will be interpreted as a destination path that will be created as a symlink to the
	// directory. The symlinks will be created after any BindPaths= or TemporaryFileSystem= options have been set up, to make
	// ephemeral symlinking possible. The same source can have multiple symlinks, by using the same first parameter, but a different
	// second parameter. The third parameter is a flags field, and since v257 can take a value of ro to make the directory read only
	// for the service. This is also supported for ConfigurationDirectory=. If multiple symlinks are set up, the directory will
	// be read only if at least one is configured to be read only. To pass a flag without a destination symlink, the second parameter
	// can be empty, for example:
	//
	//	ConfigurationDirectory=foo::ro
	//
	// The directories defined by these options are always created under the standard paths used by systemd (/var/, /run/, /etc/,
	// …). If the service needs directories in a different location, a different mechanism has to be used to create them.
	//
	// tmpfiles.d provides functionality that overlaps with these options. Using these options is recommended, because the
	// lifetime of the directories is tied directly to the lifetime of the unit, and it is not necessary to ensure that the tmpfiles.d
	// configuration is executed before the unit is started.
	//
	// To remove any of the directories created by these settings, use the systemctl clean … command on the relevant units, see
	// systemctl for details.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar baz
	//
	// the service manager creates /run/foo (if it does not exist), /run/foo/bar, and /run/baz. The directories /run/foo/bar
	// and /run/baz except /run/foo are owned by the user and group specified in User= and Group=, and removed when the service
	// is stopped.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar StateDirectory=aaa/bbb ccc
	//
	// then the environment variable "RUNTIME_DIRECTORY" is set with "/run/foo/bar", and "STATE_DIRECTORY" is set with "/var/lib/aaa/bbb:/var/lib/ccc".
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo:bar foo:baz
	//
	// the service manager creates /run/foo (if it does not exist), and /run/bar plus /run/baz as symlinks to /run/foo.
	//
	// Added in version 211.
	StateDirectory systemdconf.Value

	// These options take a whitespace-separated list of directory names. The specified directory names must be relative, and
	// may not include "..". If set, when the unit is started, one or more directories by the specified names will be created (including
	// their parents) below the locations defined in the following table. Also, the corresponding environment variable will
	// be defined with the full paths of the directories. If multiple directories are set, then in the environment variable the
	// paths are concatenated with colon (":").
	//
	// If DynamicUser= is used, and if the kernel version supports id-mapped mounts, the specified directories will be owned
	// by "nobody" in the host namespace and will be mapped to (and will be owned by) the service's UID/GID in its own namespace.
	// For backward compatibility, existing directories created without id-mapped mounts will be kept untouched.
	//
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	|        DIRECTORY        | BELOW PATH FOR SYSTEM UNITS | BELOW PATH FOR USER UNITS | ENVIRONMENT VARIABLE SET |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	| RuntimeDirectory=       | /run/                       | $XDG_RUNTIME_DIR          | $RUNTIME_DIRECTORY       |
	//	| StateDirectory=         | /var/lib/                   | $XDG_STATE_HOME           | $STATE_DIRECTORY         |
	//	| CacheDirectory=         | /var/cache/                 | $XDG_CACHE_HOME           | $CACHE_DIRECTORY         |
	//	| LogsDirectory=          | /var/log/                   | $XDG_STATE_HOME/log/      | $LOGS_DIRECTORY          |
	//	| ConfigurationDirectory= | /etc/                       | $XDG_CONFIG_HOME          | $CONFIGURATION_DIRECTORY |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//
	// In case of RuntimeDirectory= the innermost subdirectories are removed when the unit is stopped. It is possible to preserve
	// the specified directories in this case if RuntimeDirectoryPreserve= is configured to restart or yes (see below). The
	// directories specified with StateDirectory=, CacheDirectory=, LogsDirectory=, ConfigurationDirectory= are not
	// removed when the unit is stopped.
	//
	// Except in case of ConfigurationDirectory=, the innermost specified directories will be owned by the user and group specified
	// in User= and Group=. If the specified directories already exist and their owning user or group do not match the configured
	// ones, all files and directories below the specified directories as well as the directories themselves will have their
	// file ownership recursively changed to match what is configured. As an optimization, if the specified directories are
	// already owned by the right user and group, files and directories below of them are left as-is, even if they do not match what
	// is requested. The innermost specified directories will have their access mode adjusted to the what is specified in RuntimeDirectoryMode=,
	// StateDirectoryMode=, CacheDirectoryMode=, LogsDirectoryMode= and ConfigurationDirectoryMode=.
	//
	// These options imply BindPaths= for the specified paths. When combined with RootDirectory= or RootImage= these paths
	// always reside on the host and are mounted from there into the unit's file system namespace.
	//
	// If DynamicUser= is used, the logic for CacheDirectory=, LogsDirectory= and StateDirectory= is slightly altered: the
	// directories are created below /var/cache/private, /var/log/private and /var/lib/private, respectively, which are
	// host directories made inaccessible to unprivileged users, which ensures that access to these directories cannot be gained
	// through dynamic user ID recycling. Symbolic links are created to hide this difference in behaviour. Both from perspective
	// of the host and from inside the unit, the relevant directories hence always appear directly below /var/cache, /var/log
	// and /var/lib.
	//
	// Use RuntimeDirectory= to manage one or more runtime directories for the unit and bind their lifetime to the daemon runtime.
	// This is particularly useful for unprivileged daemons that cannot create runtime directories in /run/ due to lack of privileges,
	// and to make sure the runtime directory is cleaned up automatically after use. For runtime directories that require more
	// complex or different configuration or lifetime guarantees, please consider using tmpfiles.d.
	//
	// RuntimeDirectory=, StateDirectory=, CacheDirectory= and LogsDirectory= optionally support two more parameters,
	// separated by ":". The second parameter will be interpreted as a destination path that will be created as a symlink to the
	// directory. The symlinks will be created after any BindPaths= or TemporaryFileSystem= options have been set up, to make
	// ephemeral symlinking possible. The same source can have multiple symlinks, by using the same first parameter, but a different
	// second parameter. The third parameter is a flags field, and since v257 can take a value of ro to make the directory read only
	// for the service. This is also supported for ConfigurationDirectory=. If multiple symlinks are set up, the directory will
	// be read only if at least one is configured to be read only. To pass a flag without a destination symlink, the second parameter
	// can be empty, for example:
	//
	//	ConfigurationDirectory=foo::ro
	//
	// The directories defined by these options are always created under the standard paths used by systemd (/var/, /run/, /etc/,
	// …). If the service needs directories in a different location, a different mechanism has to be used to create them.
	//
	// tmpfiles.d provides functionality that overlaps with these options. Using these options is recommended, because the
	// lifetime of the directories is tied directly to the lifetime of the unit, and it is not necessary to ensure that the tmpfiles.d
	// configuration is executed before the unit is started.
	//
	// To remove any of the directories created by these settings, use the systemctl clean … command on the relevant units, see
	// systemctl for details.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar baz
	//
	// the service manager creates /run/foo (if it does not exist), /run/foo/bar, and /run/baz. The directories /run/foo/bar
	// and /run/baz except /run/foo are owned by the user and group specified in User= and Group=, and removed when the service
	// is stopped.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar StateDirectory=aaa/bbb ccc
	//
	// then the environment variable "RUNTIME_DIRECTORY" is set with "/run/foo/bar", and "STATE_DIRECTORY" is set with "/var/lib/aaa/bbb:/var/lib/ccc".
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo:bar foo:baz
	//
	// the service manager creates /run/foo (if it does not exist), and /run/bar plus /run/baz as symlinks to /run/foo.
	//
	// Added in version 211.
	CacheDirectory systemdconf.Value

	// These options take a whitespace-separated list of directory names. The specified directory names must be relative, and
	// may not include "..". If set, when the unit is started, one or more directories by the specified names will be created (including
	// their parents) below the locations defined in the following table. Also, the corresponding environment variable will
	// be defined with the full paths of the directories. If multiple directories are set, then in the environment variable the
	// paths are concatenated with colon (":").
	//
	// If DynamicUser= is used, and if the kernel version supports id-mapped mounts, the specified directories will be owned
	// by "nobody" in the host namespace and will be mapped to (and will be owned by) the service's UID/GID in its own namespace.
	// For backward compatibility, existing directories created without id-mapped mounts will be kept untouched.
	//
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	|        DIRECTORY        | BELOW PATH FOR SYSTEM UNITS | BELOW PATH FOR USER UNITS | ENVIRONMENT VARIABLE SET |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	| RuntimeDirectory=       | /run/                       | $XDG_RUNTIME_DIR          | $RUNTIME_DIRECTORY       |
	//	| StateDirectory=         | /var/lib/                   | $XDG_STATE_HOME           | $STATE_DIRECTORY         |
	//	| CacheDirectory=         | /var/cache/                 | $XDG_CACHE_HOME           | $CACHE_DIRECTORY         |
	//	| LogsDirectory=          | /var/log/                   | $XDG_STATE_HOME/log/      | $LOGS_DIRECTORY          |
	//	| ConfigurationDirectory= | /etc/                       | $XDG_CONFIG_HOME          | $CONFIGURATION_DIRECTORY |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//
	// In case of RuntimeDirectory= the innermost subdirectories are removed when the unit is stopped. It is possible to preserve
	// the specified directories in this case if RuntimeDirectoryPreserve= is configured to restart or yes (see below). The
	// directories specified with StateDirectory=, CacheDirectory=, LogsDirectory=, ConfigurationDirectory= are not
	// removed when the unit is stopped.
	//
	// Except in case of ConfigurationDirectory=, the innermost specified directories will be owned by the user and group specified
	// in User= and Group=. If the specified directories already exist and their owning user or group do not match the configured
	// ones, all files and directories below the specified directories as well as the directories themselves will have their
	// file ownership recursively changed to match what is configured. As an optimization, if the specified directories are
	// already owned by the right user and group, files and directories below of them are left as-is, even if they do not match what
	// is requested. The innermost specified directories will have their access mode adjusted to the what is specified in RuntimeDirectoryMode=,
	// StateDirectoryMode=, CacheDirectoryMode=, LogsDirectoryMode= and ConfigurationDirectoryMode=.
	//
	// These options imply BindPaths= for the specified paths. When combined with RootDirectory= or RootImage= these paths
	// always reside on the host and are mounted from there into the unit's file system namespace.
	//
	// If DynamicUser= is used, the logic for CacheDirectory=, LogsDirectory= and StateDirectory= is slightly altered: the
	// directories are created below /var/cache/private, /var/log/private and /var/lib/private, respectively, which are
	// host directories made inaccessible to unprivileged users, which ensures that access to these directories cannot be gained
	// through dynamic user ID recycling. Symbolic links are created to hide this difference in behaviour. Both from perspective
	// of the host and from inside the unit, the relevant directories hence always appear directly below /var/cache, /var/log
	// and /var/lib.
	//
	// Use RuntimeDirectory= to manage one or more runtime directories for the unit and bind their lifetime to the daemon runtime.
	// This is particularly useful for unprivileged daemons that cannot create runtime directories in /run/ due to lack of privileges,
	// and to make sure the runtime directory is cleaned up automatically after use. For runtime directories that require more
	// complex or different configuration or lifetime guarantees, please consider using tmpfiles.d.
	//
	// RuntimeDirectory=, StateDirectory=, CacheDirectory= and LogsDirectory= optionally support two more parameters,
	// separated by ":". The second parameter will be interpreted as a destination path that will be created as a symlink to the
	// directory. The symlinks will be created after any BindPaths= or TemporaryFileSystem= options have been set up, to make
	// ephemeral symlinking possible. The same source can have multiple symlinks, by using the same first parameter, but a different
	// second parameter. The third parameter is a flags field, and since v257 can take a value of ro to make the directory read only
	// for the service. This is also supported for ConfigurationDirectory=. If multiple symlinks are set up, the directory will
	// be read only if at least one is configured to be read only. To pass a flag without a destination symlink, the second parameter
	// can be empty, for example:
	//
	//	ConfigurationDirectory=foo::ro
	//
	// The directories defined by these options are always created under the standard paths used by systemd (/var/, /run/, /etc/,
	// …). If the service needs directories in a different location, a different mechanism has to be used to create them.
	//
	// tmpfiles.d provides functionality that overlaps with these options. Using these options is recommended, because the
	// lifetime of the directories is tied directly to the lifetime of the unit, and it is not necessary to ensure that the tmpfiles.d
	// configuration is executed before the unit is started.
	//
	// To remove any of the directories created by these settings, use the systemctl clean … command on the relevant units, see
	// systemctl for details.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar baz
	//
	// the service manager creates /run/foo (if it does not exist), /run/foo/bar, and /run/baz. The directories /run/foo/bar
	// and /run/baz except /run/foo are owned by the user and group specified in User= and Group=, and removed when the service
	// is stopped.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar StateDirectory=aaa/bbb ccc
	//
	// then the environment variable "RUNTIME_DIRECTORY" is set with "/run/foo/bar", and "STATE_DIRECTORY" is set with "/var/lib/aaa/bbb:/var/lib/ccc".
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo:bar foo:baz
	//
	// the service manager creates /run/foo (if it does not exist), and /run/bar plus /run/baz as symlinks to /run/foo.
	//
	// Added in version 211.
	LogsDirectory systemdconf.Value

	// These options take a whitespace-separated list of directory names. The specified directory names must be relative, and
	// may not include "..". If set, when the unit is started, one or more directories by the specified names will be created (including
	// their parents) below the locations defined in the following table. Also, the corresponding environment variable will
	// be defined with the full paths of the directories. If multiple directories are set, then in the environment variable the
	// paths are concatenated with colon (":").
	//
	// If DynamicUser= is used, and if the kernel version supports id-mapped mounts, the specified directories will be owned
	// by "nobody" in the host namespace and will be mapped to (and will be owned by) the service's UID/GID in its own namespace.
	// For backward compatibility, existing directories created without id-mapped mounts will be kept untouched.
	//
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	|        DIRECTORY        | BELOW PATH FOR SYSTEM UNITS | BELOW PATH FOR USER UNITS | ENVIRONMENT VARIABLE SET |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//	| RuntimeDirectory=       | /run/                       | $XDG_RUNTIME_DIR          | $RUNTIME_DIRECTORY       |
	//	| StateDirectory=         | /var/lib/                   | $XDG_STATE_HOME           | $STATE_DIRECTORY         |
	//	| CacheDirectory=         | /var/cache/                 | $XDG_CACHE_HOME           | $CACHE_DIRECTORY         |
	//	| LogsDirectory=          | /var/log/                   | $XDG_STATE_HOME/log/      | $LOGS_DIRECTORY          |
	//	| ConfigurationDirectory= | /etc/                       | $XDG_CONFIG_HOME          | $CONFIGURATION_DIRECTORY |
	//	+-------------------------+-----------------------------+---------------------------+--------------------------+
	//
	// In case of RuntimeDirectory= the innermost subdirectories are removed when the unit is stopped. It is possible to preserve
	// the specified directories in this case if RuntimeDirectoryPreserve= is configured to restart or yes (see below). The
	// directories specified with StateDirectory=, CacheDirectory=, LogsDirectory=, ConfigurationDirectory= are not
	// removed when the unit is stopped.
	//
	// Except in case of ConfigurationDirectory=, the innermost specified directories will be owned by the user and group specified
	// in User= and Group=. If the specified directories already exist and their owning user or group do not match the configured
	// ones, all files and directories below the specified directories as well as the directories themselves will have their
	// file ownership recursively changed to match what is configured. As an optimization, if the specified directories are
	// already owned by the right user and group, files and directories below of them are left as-is, even if they do not match what
	// is requested. The innermost specified directories will have their access mode adjusted to the what is specified in RuntimeDirectoryMode=,
	// StateDirectoryMode=, CacheDirectoryMode=, LogsDirectoryMode= and ConfigurationDirectoryMode=.
	//
	// These options imply BindPaths= for the specified paths. When combined with RootDirectory= or RootImage= these paths
	// always reside on the host and are mounted from there into the unit's file system namespace.
	//
	// If DynamicUser= is used, the logic for CacheDirectory=, LogsDirectory= and StateDirectory= is slightly altered: the
	// directories are created below /var/cache/private, /var/log/private and /var/lib/private, respectively, which are
	// host directories made inaccessible to unprivileged users, which ensures that access to these directories cannot be gained
	// through dynamic user ID recycling. Symbolic links are created to hide this difference in behaviour. Both from perspective
	// of the host and from inside the unit, the relevant directories hence always appear directly below /var/cache, /var/log
	// and /var/lib.
	//
	// Use RuntimeDirectory= to manage one or more runtime directories for the unit and bind their lifetime to the daemon runtime.
	// This is particularly useful for unprivileged daemons that cannot create runtime directories in /run/ due to lack of privileges,
	// and to make sure the runtime directory is cleaned up automatically after use. For runtime directories that require more
	// complex or different configuration or lifetime guarantees, please consider using tmpfiles.d.
	//
	// RuntimeDirectory=, StateDirectory=, CacheDirectory= and LogsDirectory= optionally support two more parameters,
	// separated by ":". The second parameter will be interpreted as a destination path that will be created as a symlink to the
	// directory. The symlinks will be created after any BindPaths= or TemporaryFileSystem= options have been set up, to make
	// ephemeral symlinking possible. The same source can have multiple symlinks, by using the same first parameter, but a different
	// second parameter. The third parameter is a flags field, and since v257 can take a value of ro to make the directory read only
	// for the service. This is also supported for ConfigurationDirectory=. If multiple symlinks are set up, the directory will
	// be read only if at least one is configured to be read only. To pass a flag without a destination symlink, the second parameter
	// can be empty, for example:
	//
	//	ConfigurationDirectory=foo::ro
	//
	// The directories defined by these options are always created under the standard paths used by systemd (/var/, /run/, /etc/,
	// …). If the service needs directories in a different location, a different mechanism has to be used to create them.
	//
	// tmpfiles.d provides functionality that overlaps with these options. Using these options is recommended, because the
	// lifetime of the directories is tied directly to the lifetime of the unit, and it is not necessary to ensure that the tmpfiles.d
	// configuration is executed before the unit is started.
	//
	// To remove any of the directories created by these settings, use the systemctl clean … command on the relevant units, see
	// systemctl for details.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar baz
	//
	// the service manager creates /run/foo (if it does not exist), /run/foo/bar, and /run/baz. The directories /run/foo/bar
	// and /run/baz except /run/foo are owned by the user and group specified in User= and Group=, and removed when the service
	// is stopped.
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo/bar StateDirectory=aaa/bbb ccc
	//
	// then the environment variable "RUNTIME_DIRECTORY" is set with "/run/foo/bar", and "STATE_DIRECTORY" is set with "/var/lib/aaa/bbb:/var/lib/ccc".
	//
	// Example: if a system service unit has the following,
	//
	//	RuntimeDirectory=foo:bar foo:baz
	//
	// the service manager creates /run/foo (if it does not exist), and /run/bar plus /run/baz as symlinks to /run/foo.
	//
	// Added in version 211.
	ConfigurationDirectory systemdconf.Value

	// Specifies the access mode of the directories specified in RuntimeDirectory=, StateDirectory=, CacheDirectory=, LogsDirectory=,
	// or ConfigurationDirectory=, respectively, as an octal number. Defaults to 0755. See "Permissions" in path_resolution
	// for a discussion of the meaning of permission bits.
	//
	// Added in version 234.
	RuntimeDirectoryMode systemdconf.Value

	// Specifies the access mode of the directories specified in RuntimeDirectory=, StateDirectory=, CacheDirectory=, LogsDirectory=,
	// or ConfigurationDirectory=, respectively, as an octal number. Defaults to 0755. See "Permissions" in path_resolution
	// for a discussion of the meaning of permission bits.
	//
	// Added in version 234.
	StateDirectoryMode systemdconf.Value

	// Specifies the access mode of the directories specified in RuntimeDirectory=, StateDirectory=, CacheDirectory=, LogsDirectory=,
	// or ConfigurationDirectory=, respectively, as an octal number. Defaults to 0755. See "Permissions" in path_resolution
	// for a discussion of the meaning of permission bits.
	//
	// Added in version 234.
	CacheDirectoryMode systemdconf.Value

	// Specifies the access mode of the directories specified in RuntimeDirectory=, StateDirectory=, CacheDirectory=, LogsDirectory=,
	// or ConfigurationDirectory=, respectively, as an octal number. Defaults to 0755. See "Permissions" in path_resolution
	// for a discussion of the meaning of permission bits.
	//
	// Added in version 234.
	LogsDirectoryMode systemdconf.Value

	// Specifies the access mode of the directories specified in RuntimeDirectory=, StateDirectory=, CacheDirectory=, LogsDirectory=,
	// or ConfigurationDirectory=, respectively, as an octal number. Defaults to 0755. See "Permissions" in path_resolution
	// for a discussion of the meaning of permission bits.
	//
	// Added in version 234.
	ConfigurationDirectoryMode systemdconf.Value

	// Takes a boolean argument or restart. If set to no (the default), the directories specified in RuntimeDirectory= are always
	// removed when the service stops. If set to restart the directories are preserved when the service is both automatically
	// and manually restarted. Here, the automatic restart means the operation specified in Restart=, and manual restart means
	// the one triggered by systemctl restart foo.service. If set to yes, then the directories are not removed when the service
	// is stopped. Note that since the runtime directory /run/ is a mount point of "tmpfs", then for system services the directories
	// specified in RuntimeDirectory= are removed when the system is rebooted.
	//
	// Added in version 235.
	RuntimeDirectoryPreserve systemdconf.Value

	// Configures a timeout on the clean-up operation requested through systemctl clean …, see systemctl for details. Takes
	// the usual time values and defaults to infinity, i.e. by default no timeout is applied. If a timeout is configured the clean
	// operation will be aborted forcibly when the timeout is reached, potentially leaving resources on disk.
	//
	// Added in version 244.
	TimeoutCleanSec systemdconf.Value

	// Sets up a new file system namespace for executed processes. These options may be used to limit access a process has to the
	// file system. Each setting takes a space-separated list of paths relative to the host's root directory (i.e. the system
	// running the service manager). Note that if paths contain symlinks, they are resolved relative to the root directory set
	// with RootDirectory=/RootImage=.
	//
	// Paths listed in ReadWritePaths= are accessible from within the namespace with the same access modes as from outside of
	// it. Paths listed in ReadOnlyPaths= are accessible for reading only, writing will be refused even if the usual file access
	// controls would permit this. Nest ReadWritePaths= inside of ReadOnlyPaths= in order to provide writable subdirectories
	// within read-only directories. Use ReadWritePaths= in order to allow-list specific paths for write access if ProtectSystem=strict
	// is used. Note that ReadWritePaths= cannot be used to gain write access to a file system whose superblock is mounted read-only.
	// On Linux, for each mount point write access is granted only if the mount point itself and the file system superblock backing
	// it are not marked read-only. ReadWritePaths= only controls the former, not the latter, hence a read-only file system superblock
	// remains protected.
	//
	// Paths listed in InaccessiblePaths= will be made inaccessible for processes inside the namespace along with everything
	// below them in the file system hierarchy. This may be more restrictive than desired, because it is not possible to nest ReadWritePaths=,
	// ReadOnlyPaths=, BindPaths=, or BindReadOnlyPaths= inside it. For a more flexible option, see TemporaryFileSystem=.
	//
	// Content in paths listed in NoExecPaths= are not executable even if the usual file access controls would permit this. Nest
	// ExecPaths= inside of NoExecPaths= in order to provide executable content within non-executable directories.
	//
	// Non-directory paths may be specified as well. These options may be specified more than once, in which case all paths listed
	// will have limited access from within the namespace. If the empty string is assigned to this option, the specific list is
	// reset, and all prior assignments have no effect.
	//
	// Paths in ReadWritePaths=, ReadOnlyPaths=, InaccessiblePaths=, ExecPaths= and NoExecPaths= may be prefixed with "-",
	// in which case they will be ignored when they do not exist. If prefixed with "+" the paths are taken relative to the root directory
	// of the unit, as configured with RootDirectory=/RootImage=, instead of relative to the root directory of the host (see
	// above). When combining "-" and "+" on the same path make sure to specify "-" first, and "+" second.
	//
	// Note that these settings will disconnect propagation of mounts from the unit's processes to the host. This means that this
	// setting may not be used for services which shall be able to install mount points in the main mount namespace. For ReadWritePaths=
	// and ReadOnlyPaths=, propagation in the other direction is not affected, i.e. mounts created on the host generally appear
	// in the unit processes' namespace, and mounts removed on the host also disappear there too. In particular, note that mount
	// propagation from host to unit will result in unmodified mounts to be created in the unit's namespace, i.e. writable mounts
	// appearing on the host will be writable in the unit's namespace too, even when propagated below a path marked with ReadOnlyPaths=!
	// Restricting access with these options hence does not extend to submounts of a directory that are created later on. This
	// means the lock-down offered by that setting is not complete, and does not offer full protection.
	//
	// Note that the effect of these settings may be undone by privileged processes. In order to set up an effective sandboxed environment
	// for a unit it is thus recommended to combine these settings with either CapabilityBoundingSet=~CAP_SYS_ADMIN or SystemCallFilter=~@mount.
	//
	// Please be extra careful when applying these options to API file systems (a list of them could be found in MountAPIVPS=),
	// since they may be required for basic system functionalities. Moreover, /run/ needs to be writable for setting up mount
	// namespace and propagation.
	//
	// Simple allow-list example using these directives:
	//
	//	[Service] ReadOnlyPaths=/ ReadWritePaths=/var /run InaccessiblePaths=-/lost+found NoExecPaths=/ ExecPaths=/usr/sbin/my_daemon
	//	/usr/lib /usr/lib64
	//
	// These options are only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 231.
	ReadWritePaths systemdconf.Value

	// Sets up a new file system namespace for executed processes. These options may be used to limit access a process has to the
	// file system. Each setting takes a space-separated list of paths relative to the host's root directory (i.e. the system
	// running the service manager). Note that if paths contain symlinks, they are resolved relative to the root directory set
	// with RootDirectory=/RootImage=.
	//
	// Paths listed in ReadWritePaths= are accessible from within the namespace with the same access modes as from outside of
	// it. Paths listed in ReadOnlyPaths= are accessible for reading only, writing will be refused even if the usual file access
	// controls would permit this. Nest ReadWritePaths= inside of ReadOnlyPaths= in order to provide writable subdirectories
	// within read-only directories. Use ReadWritePaths= in order to allow-list specific paths for write access if ProtectSystem=strict
	// is used. Note that ReadWritePaths= cannot be used to gain write access to a file system whose superblock is mounted read-only.
	// On Linux, for each mount point write access is granted only if the mount point itself and the file system superblock backing
	// it are not marked read-only. ReadWritePaths= only controls the former, not the latter, hence a read-only file system superblock
	// remains protected.
	//
	// Paths listed in InaccessiblePaths= will be made inaccessible for processes inside the namespace along with everything
	// below them in the file system hierarchy. This may be more restrictive than desired, because it is not possible to nest ReadWritePaths=,
	// ReadOnlyPaths=, BindPaths=, or BindReadOnlyPaths= inside it. For a more flexible option, see TemporaryFileSystem=.
	//
	// Content in paths listed in NoExecPaths= are not executable even if the usual file access controls would permit this. Nest
	// ExecPaths= inside of NoExecPaths= in order to provide executable content within non-executable directories.
	//
	// Non-directory paths may be specified as well. These options may be specified more than once, in which case all paths listed
	// will have limited access from within the namespace. If the empty string is assigned to this option, the specific list is
	// reset, and all prior assignments have no effect.
	//
	// Paths in ReadWritePaths=, ReadOnlyPaths=, InaccessiblePaths=, ExecPaths= and NoExecPaths= may be prefixed with "-",
	// in which case they will be ignored when they do not exist. If prefixed with "+" the paths are taken relative to the root directory
	// of the unit, as configured with RootDirectory=/RootImage=, instead of relative to the root directory of the host (see
	// above). When combining "-" and "+" on the same path make sure to specify "-" first, and "+" second.
	//
	// Note that these settings will disconnect propagation of mounts from the unit's processes to the host. This means that this
	// setting may not be used for services which shall be able to install mount points in the main mount namespace. For ReadWritePaths=
	// and ReadOnlyPaths=, propagation in the other direction is not affected, i.e. mounts created on the host generally appear
	// in the unit processes' namespace, and mounts removed on the host also disappear there too. In particular, note that mount
	// propagation from host to unit will result in unmodified mounts to be created in the unit's namespace, i.e. writable mounts
	// appearing on the host will be writable in the unit's namespace too, even when propagated below a path marked with ReadOnlyPaths=!
	// Restricting access with these options hence does not extend to submounts of a directory that are created later on. This
	// means the lock-down offered by that setting is not complete, and does not offer full protection.
	//
	// Note that the effect of these settings may be undone by privileged processes. In order to set up an effective sandboxed environment
	// for a unit it is thus recommended to combine these settings with either CapabilityBoundingSet=~CAP_SYS_ADMIN or SystemCallFilter=~@mount.
	//
	// Please be extra careful when applying these options to API file systems (a list of them could be found in MountAPIVPS=),
	// since they may be required for basic system functionalities. Moreover, /run/ needs to be writable for setting up mount
	// namespace and propagation.
	//
	// Simple allow-list example using these directives:
	//
	//	[Service] ReadOnlyPaths=/ ReadWritePaths=/var /run InaccessiblePaths=-/lost+found NoExecPaths=/ ExecPaths=/usr/sbin/my_daemon
	//	/usr/lib /usr/lib64
	//
	// These options are only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 231.
	ReadOnlyPaths systemdconf.Value

	// Sets up a new file system namespace for executed processes. These options may be used to limit access a process has to the
	// file system. Each setting takes a space-separated list of paths relative to the host's root directory (i.e. the system
	// running the service manager). Note that if paths contain symlinks, they are resolved relative to the root directory set
	// with RootDirectory=/RootImage=.
	//
	// Paths listed in ReadWritePaths= are accessible from within the namespace with the same access modes as from outside of
	// it. Paths listed in ReadOnlyPaths= are accessible for reading only, writing will be refused even if the usual file access
	// controls would permit this. Nest ReadWritePaths= inside of ReadOnlyPaths= in order to provide writable subdirectories
	// within read-only directories. Use ReadWritePaths= in order to allow-list specific paths for write access if ProtectSystem=strict
	// is used. Note that ReadWritePaths= cannot be used to gain write access to a file system whose superblock is mounted read-only.
	// On Linux, for each mount point write access is granted only if the mount point itself and the file system superblock backing
	// it are not marked read-only. ReadWritePaths= only controls the former, not the latter, hence a read-only file system superblock
	// remains protected.
	//
	// Paths listed in InaccessiblePaths= will be made inaccessible for processes inside the namespace along with everything
	// below them in the file system hierarchy. This may be more restrictive than desired, because it is not possible to nest ReadWritePaths=,
	// ReadOnlyPaths=, BindPaths=, or BindReadOnlyPaths= inside it. For a more flexible option, see TemporaryFileSystem=.
	//
	// Content in paths listed in NoExecPaths= are not executable even if the usual file access controls would permit this. Nest
	// ExecPaths= inside of NoExecPaths= in order to provide executable content within non-executable directories.
	//
	// Non-directory paths may be specified as well. These options may be specified more than once, in which case all paths listed
	// will have limited access from within the namespace. If the empty string is assigned to this option, the specific list is
	// reset, and all prior assignments have no effect.
	//
	// Paths in ReadWritePaths=, ReadOnlyPaths=, InaccessiblePaths=, ExecPaths= and NoExecPaths= may be prefixed with "-",
	// in which case they will be ignored when they do not exist. If prefixed with "+" the paths are taken relative to the root directory
	// of the unit, as configured with RootDirectory=/RootImage=, instead of relative to the root directory of the host (see
	// above). When combining "-" and "+" on the same path make sure to specify "-" first, and "+" second.
	//
	// Note that these settings will disconnect propagation of mounts from the unit's processes to the host. This means that this
	// setting may not be used for services which shall be able to install mount points in the main mount namespace. For ReadWritePaths=
	// and ReadOnlyPaths=, propagation in the other direction is not affected, i.e. mounts created on the host generally appear
	// in the unit processes' namespace, and mounts removed on the host also disappear there too. In particular, note that mount
	// propagation from host to unit will result in unmodified mounts to be created in the unit's namespace, i.e. writable mounts
	// appearing on the host will be writable in the unit's namespace too, even when propagated below a path marked with ReadOnlyPaths=!
	// Restricting access with these options hence does not extend to submounts of a directory that are created later on. This
	// means the lock-down offered by that setting is not complete, and does not offer full protection.
	//
	// Note that the effect of these settings may be undone by privileged processes. In order to set up an effective sandboxed environment
	// for a unit it is thus recommended to combine these settings with either CapabilityBoundingSet=~CAP_SYS_ADMIN or SystemCallFilter=~@mount.
	//
	// Please be extra careful when applying these options to API file systems (a list of them could be found in MountAPIVPS=),
	// since they may be required for basic system functionalities. Moreover, /run/ needs to be writable for setting up mount
	// namespace and propagation.
	//
	// Simple allow-list example using these directives:
	//
	//	[Service] ReadOnlyPaths=/ ReadWritePaths=/var /run InaccessiblePaths=-/lost+found NoExecPaths=/ ExecPaths=/usr/sbin/my_daemon
	//	/usr/lib /usr/lib64
	//
	// These options are only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 231.
	InaccessiblePaths systemdconf.Value

	// Sets up a new file system namespace for executed processes. These options may be used to limit access a process has to the
	// file system. Each setting takes a space-separated list of paths relative to the host's root directory (i.e. the system
	// running the service manager). Note that if paths contain symlinks, they are resolved relative to the root directory set
	// with RootDirectory=/RootImage=.
	//
	// Paths listed in ReadWritePaths= are accessible from within the namespace with the same access modes as from outside of
	// it. Paths listed in ReadOnlyPaths= are accessible for reading only, writing will be refused even if the usual file access
	// controls would permit this. Nest ReadWritePaths= inside of ReadOnlyPaths= in order to provide writable subdirectories
	// within read-only directories. Use ReadWritePaths= in order to allow-list specific paths for write access if ProtectSystem=strict
	// is used. Note that ReadWritePaths= cannot be used to gain write access to a file system whose superblock is mounted read-only.
	// On Linux, for each mount point write access is granted only if the mount point itself and the file system superblock backing
	// it are not marked read-only. ReadWritePaths= only controls the former, not the latter, hence a read-only file system superblock
	// remains protected.
	//
	// Paths listed in InaccessiblePaths= will be made inaccessible for processes inside the namespace along with everything
	// below them in the file system hierarchy. This may be more restrictive than desired, because it is not possible to nest ReadWritePaths=,
	// ReadOnlyPaths=, BindPaths=, or BindReadOnlyPaths= inside it. For a more flexible option, see TemporaryFileSystem=.
	//
	// Content in paths listed in NoExecPaths= are not executable even if the usual file access controls would permit this. Nest
	// ExecPaths= inside of NoExecPaths= in order to provide executable content within non-executable directories.
	//
	// Non-directory paths may be specified as well. These options may be specified more than once, in which case all paths listed
	// will have limited access from within the namespace. If the empty string is assigned to this option, the specific list is
	// reset, and all prior assignments have no effect.
	//
	// Paths in ReadWritePaths=, ReadOnlyPaths=, InaccessiblePaths=, ExecPaths= and NoExecPaths= may be prefixed with "-",
	// in which case they will be ignored when they do not exist. If prefixed with "+" the paths are taken relative to the root directory
	// of the unit, as configured with RootDirectory=/RootImage=, instead of relative to the root directory of the host (see
	// above). When combining "-" and "+" on the same path make sure to specify "-" first, and "+" second.
	//
	// Note that these settings will disconnect propagation of mounts from the unit's processes to the host. This means that this
	// setting may not be used for services which shall be able to install mount points in the main mount namespace. For ReadWritePaths=
	// and ReadOnlyPaths=, propagation in the other direction is not affected, i.e. mounts created on the host generally appear
	// in the unit processes' namespace, and mounts removed on the host also disappear there too. In particular, note that mount
	// propagation from host to unit will result in unmodified mounts to be created in the unit's namespace, i.e. writable mounts
	// appearing on the host will be writable in the unit's namespace too, even when propagated below a path marked with ReadOnlyPaths=!
	// Restricting access with these options hence does not extend to submounts of a directory that are created later on. This
	// means the lock-down offered by that setting is not complete, and does not offer full protection.
	//
	// Note that the effect of these settings may be undone by privileged processes. In order to set up an effective sandboxed environment
	// for a unit it is thus recommended to combine these settings with either CapabilityBoundingSet=~CAP_SYS_ADMIN or SystemCallFilter=~@mount.
	//
	// Please be extra careful when applying these options to API file systems (a list of them could be found in MountAPIVPS=),
	// since they may be required for basic system functionalities. Moreover, /run/ needs to be writable for setting up mount
	// namespace and propagation.
	//
	// Simple allow-list example using these directives:
	//
	//	[Service] ReadOnlyPaths=/ ReadWritePaths=/var /run InaccessiblePaths=-/lost+found NoExecPaths=/ ExecPaths=/usr/sbin/my_daemon
	//	/usr/lib /usr/lib64
	//
	// These options are only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 231.
	ExecPaths systemdconf.Value

	// Sets up a new file system namespace for executed processes. These options may be used to limit access a process has to the
	// file system. Each setting takes a space-separated list of paths relative to the host's root directory (i.e. the system
	// running the service manager). Note that if paths contain symlinks, they are resolved relative to the root directory set
	// with RootDirectory=/RootImage=.
	//
	// Paths listed in ReadWritePaths= are accessible from within the namespace with the same access modes as from outside of
	// it. Paths listed in ReadOnlyPaths= are accessible for reading only, writing will be refused even if the usual file access
	// controls would permit this. Nest ReadWritePaths= inside of ReadOnlyPaths= in order to provide writable subdirectories
	// within read-only directories. Use ReadWritePaths= in order to allow-list specific paths for write access if ProtectSystem=strict
	// is used. Note that ReadWritePaths= cannot be used to gain write access to a file system whose superblock is mounted read-only.
	// On Linux, for each mount point write access is granted only if the mount point itself and the file system superblock backing
	// it are not marked read-only. ReadWritePaths= only controls the former, not the latter, hence a read-only file system superblock
	// remains protected.
	//
	// Paths listed in InaccessiblePaths= will be made inaccessible for processes inside the namespace along with everything
	// below them in the file system hierarchy. This may be more restrictive than desired, because it is not possible to nest ReadWritePaths=,
	// ReadOnlyPaths=, BindPaths=, or BindReadOnlyPaths= inside it. For a more flexible option, see TemporaryFileSystem=.
	//
	// Content in paths listed in NoExecPaths= are not executable even if the usual file access controls would permit this. Nest
	// ExecPaths= inside of NoExecPaths= in order to provide executable content within non-executable directories.
	//
	// Non-directory paths may be specified as well. These options may be specified more than once, in which case all paths listed
	// will have limited access from within the namespace. If the empty string is assigned to this option, the specific list is
	// reset, and all prior assignments have no effect.
	//
	// Paths in ReadWritePaths=, ReadOnlyPaths=, InaccessiblePaths=, ExecPaths= and NoExecPaths= may be prefixed with "-",
	// in which case they will be ignored when they do not exist. If prefixed with "+" the paths are taken relative to the root directory
	// of the unit, as configured with RootDirectory=/RootImage=, instead of relative to the root directory of the host (see
	// above). When combining "-" and "+" on the same path make sure to specify "-" first, and "+" second.
	//
	// Note that these settings will disconnect propagation of mounts from the unit's processes to the host. This means that this
	// setting may not be used for services which shall be able to install mount points in the main mount namespace. For ReadWritePaths=
	// and ReadOnlyPaths=, propagation in the other direction is not affected, i.e. mounts created on the host generally appear
	// in the unit processes' namespace, and mounts removed on the host also disappear there too. In particular, note that mount
	// propagation from host to unit will result in unmodified mounts to be created in the unit's namespace, i.e. writable mounts
	// appearing on the host will be writable in the unit's namespace too, even when propagated below a path marked with ReadOnlyPaths=!
	// Restricting access with these options hence does not extend to submounts of a directory that are created later on. This
	// means the lock-down offered by that setting is not complete, and does not offer full protection.
	//
	// Note that the effect of these settings may be undone by privileged processes. In order to set up an effective sandboxed environment
	// for a unit it is thus recommended to combine these settings with either CapabilityBoundingSet=~CAP_SYS_ADMIN or SystemCallFilter=~@mount.
	//
	// Please be extra careful when applying these options to API file systems (a list of them could be found in MountAPIVPS=),
	// since they may be required for basic system functionalities. Moreover, /run/ needs to be writable for setting up mount
	// namespace and propagation.
	//
	// Simple allow-list example using these directives:
	//
	//	[Service] ReadOnlyPaths=/ ReadWritePaths=/var /run InaccessiblePaths=-/lost+found NoExecPaths=/ ExecPaths=/usr/sbin/my_daemon
	//	/usr/lib /usr/lib64
	//
	// These options are only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 231.
	NoExecPaths systemdconf.Value

	// Takes a space-separated list of mount points for temporary file systems (tmpfs). If set, a new file system namespace is
	// set up for executed processes, and a temporary file system is mounted on each mount point. This option may be specified more
	// than once, in which case temporary file systems are mounted on all listed mount points. If the empty string is assigned to
	// this option, the list is reset, and all prior assignments have no effect. Each mount point may optionally be suffixed with
	// a colon (":") and mount options such as "size=10%" or "ro". By default, each temporary file system is mounted with "nodev,strictatime,mode=0755".
	// These can be disabled by explicitly specifying the corresponding mount options, e.g., "dev" or "nostrictatime".
	//
	// This is useful to hide files or directories not relevant to the processes invoked by the unit, while necessary files or directories
	// can be still accessed by combining with BindPaths= or BindReadOnlyPaths=:
	//
	// Example: if a unit has the following,
	//
	//	TemporaryFileSystem=/var:ro BindReadOnlyPaths=/var/lib/systemd
	//
	// then the invoked processes by the unit cannot see any files or directories under /var/ except for /var/lib/systemd or its
	// contents.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 238.
	TemporaryFileSystem systemdconf.Value

	// Takes a boolean argument, or "disconnected". If enabled, a new file system namespace will be set up for the executed processes,
	// and /tmp/ and /var/tmp/ directories inside it are not shared with processes outside of the namespace, plus all temporary
	// files created by a service in these directories will be removed after the service is stopped. If "true", the backing storage
	// of the private temporary directories will remain on the host's /tmp/ and /var/tmp/ directories. If "disconnected", the
	// directories will be backed by a completely new tmpfs instance, meaning that the storage is fully disconnected from the
	// host namespace. Defaults to false.
	//
	// This setting is useful to secure access to temporary files of the process, but makes sharing between processes via /tmp/
	// or /var/tmp/ impossible. If not set to "disconnected", it is possible to run two or more units within the same private /tmp/
	// and /var/tmp/ namespace by using the JoinsNamespaceOf= directive, see systemd.unit for details. This setting is implied
	// if DynamicUser= is set. For this setting, the same restrictions regarding mount propagation and privileges apply as for
	// ReadOnlyPaths= and related calls, see above. If set to "true" (as opposed to "disconnected"), this has the side effect
	// of adding Requires= and After= dependencies on all mount units necessary to access /tmp/ and /var/tmp/ on the host. Moreover
	// an implicitly After= ordering on systemd-tmpfiles-setup.service is added.
	//
	// Note that the implementation of this setting might be impossible (for example if mount namespaces are not available),
	// and the unit should be written in a way that does not solely rely on this setting for security.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	PrivateTmp systemdconf.Value

	// Takes a boolean argument. If true, sets up a new /dev/ mount for the executed processes and only adds API pseudo devices such
	// as /dev/null, /dev/zero or /dev/random (as well as the pseudo TTY subsystem) to it, but no physical devices such as /dev/sda,
	// system memory /dev/mem, system ports /dev/port and others. This is useful to turn off physical device access by the executed
	// process. Defaults to false.
	//
	// Enabling this option will install a system call filter to block low-level I/O system calls that are grouped in the @raw-io
	// set, remove CAP_MKNOD and CAP_SYS_RAWIO from the capability bounding set for the unit, and set DevicePolicy=closed (see
	// systemd.resource-control for details). Note that using this setting will disconnect propagation of mounts from the
	// service to the host (propagation in the opposite direction continues to work). This means that this setting may not be used
	// for services which shall be able to install mount points in the main mount namespace. The new /dev/ will be mounted read-only
	// and 'noexec'. The latter may break old programs which try to set up executable memory by using mmap of /dev/zero instead
	// of using MAP_ANON. For this setting the same restrictions regarding mount propagation and privileges apply as for ReadOnlyPaths=
	// and related calls, see above.
	//
	// Note that the implementation of this setting might be impossible (for example if mount namespaces are not available),
	// and the unit should be written in a way that does not solely rely on this setting for security.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// When access to some but not all devices must be possible, the DeviceAllow= setting might be used instead. See systemd.resource-control.
	//
	// Added in version 209.
	PrivateDevices systemdconf.Value

	// Takes a boolean argument. If true, sets up a new network namespace for the executed processes and configures only the loopback
	// network device "lo" inside it. No other network devices will be available to the executed process. This is useful to turn
	// off network access by the executed process. Defaults to false. It is possible to run two or more units within the same private
	// network namespace by using the JoinsNamespaceOf= directive, see systemd.unit for details. Note that this option will
	// disconnect all socket families from the host, including AF_NETLINK and AF_UNIX. Effectively, for AF_NETLINK this means
	// that device configuration events received from systemd-udevd.service are not delivered to the unit's processes. And
	// for AF_UNIX this has the effect that AF_UNIX sockets in the abstract socket namespace of the host will become unavailable
	// to the unit's processes (however, those located in the file system will continue to be accessible).
	//
	// Note that the implementation of this setting might be impossible (for example if network namespaces are not available),
	// and the unit should be written in a way that does not solely rely on this setting for security.
	//
	// When this option is enabled, PrivateMounts= is implied unless it is explicitly disabled, and /sys will be remounted to
	// associate it with the new network namespace.
	//
	// When this option is used on a socket unit any sockets bound on behalf of this unit will be bound within a private network namespace.
	// This may be combined with JoinsNamespaceOf= to listen on sockets inside of network namespaces of other services.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	PrivateNetwork systemdconf.Value

	// Takes an absolute file system path referring to a Linux network namespace pseudo-file (i.e. a file like /proc/$PID/ns/net
	// or a bind mount or symlink to one). When set the invoked processes are added to the network namespace referenced by that path.
	// The path has to point to a valid namespace file at the moment the processes are forked off. If this option is used PrivateNetwork=
	// has no effect. If this option is used together with JoinsNamespaceOf= then it only has an effect if this unit is started before
	// any of the listed units that have PrivateNetwork= or NetworkNamespacePath= configured, as otherwise the network namespace
	// of those units is reused.
	//
	// When this option is enabled, PrivateMounts= is implied unless it is explicitly disabled, and /sys will be remounted to
	// associate it with the new network namespace.
	//
	// When this option is used on a socket unit any sockets bound on behalf of this unit will be bound within the specified network
	// namespace.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 242.
	NetworkNamespacePath systemdconf.Value

	// Takes a boolean argument. If true, sets up a new IPC namespace for the executed processes. Each IPC namespace has its own
	// set of System V IPC identifiers and its own POSIX message queue file system. This is useful to avoid name clash of IPC identifiers.
	// Defaults to false. It is possible to run two or more units within the same private IPC namespace by using the JoinsNamespaceOf=
	// directive, see systemd.unit for details.
	//
	// Note that IPC namespacing does not have an effect on AF_UNIX sockets, which are the most common form of IPC used on Linux.
	// Instead, AF_UNIX sockets in the file system are subject to mount namespacing, and those in the abstract namespace are subject
	// to network namespacing. IPC namespacing only has an effect on SysV IPC (which is mostly legacy) as well as POSIX message
	// queues (for which AF_UNIX/SOCK_SEQPACKET sockets are typically a better replacement). IPC namespacing also has no effect
	// on POSIX shared memory (which is subject to mount namespacing) either. See ipc_namespaces for the details.
	//
	// Note that the implementation of this setting might be impossible (for example if IPC namespaces are not available), and
	// the unit should be written in a way that does not solely rely on this setting for security.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 248.
	PrivateIPC systemdconf.Value

	// Takes an absolute file system path referring to a Linux IPC namespace pseudo-file (i.e. a file like /proc/$PID/ns/ipc
	// or a bind mount or symlink to one). When set the invoked processes are added to the network namespace referenced by that path.
	// The path has to point to a valid namespace file at the moment the processes are forked off. If this option is used PrivateIPC=
	// has no effect. If this option is used together with JoinsNamespaceOf= then it only has an effect if this unit is started before
	// any of the listed units that have PrivateIPC= or IPCNamespacePath= configured, as otherwise the network namespace of
	// those units is reused.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 248.
	IPCNamespacePath systemdconf.Value

	// Takes a boolean argument. When set, it enables KSM (kernel samepage merging) for the processes. KSM is a memory-saving
	// de-duplication feature. Anonymous memory pages with identical content can be replaced by a single write-protected page.
	// This feature should only be enabled for jobs that share the same security domain. For details, see Kernel Samepage Merging
	// in the kernel documentation.
	//
	// Note that this functionality might not be available, for example if KSM is disabled in the kernel, or the kernel does not
	// support controlling KSM at the process level through prctl.
	//
	// Added in version 254.
	MemoryKSM systemdconf.Value

	// Takes a boolean argument. Defaults to false. If enabled, sets up a new PID namespace for the executed processes. Each executed
	// process is now PID 1 - the init process - in the new namespace. /proc/ is mounted such that only processes in the PID namespace
	// are visible. If PrivatePIDs= is set, MountAPIVFS=yes is implied.
	//
	// PrivatePIDs= is only supported for service units. This setting is not supported with Type=forking since the kernel will
	// kill all processes in the PID namespace if the init process terminates.
	//
	// This setting will be ignored if the kernel does not support PID namespaces.
	//
	// Note unprivileged user services (i.e. a service run by the per-user instance of the service manager) will fail with PrivatePIDs=yes
	// if /proc/ is masked (i.e. /proc/kmsg is over-mounted with tmpfs like systemd-nspawn does). This is due to a kernel restriction
	// not allowing unprivileged user namespaces to mount a less restrictive instance of /proc/.
	//
	// Added in version 257.
	PrivatePIDs systemdconf.Value

	// Takes a boolean argument or one of "self" or "identity". Defaults to false. If enabled, sets up a new user namespace for the
	// executed processes and configures a user and group mapping. If set to a true value or "self", a minimal user and group mapping
	// is configured that maps the "root" user and group as well as the unit's own user and group to themselves and everything else
	// to the "nobody" user and group. This is useful to securely detach the user and group databases used by the unit from the rest
	// of the system, and thus to create an effective sandbox environment. All files, directories, processes, IPC objects and
	// other resources owned by users/groups not equaling "root" or the unit's own will stay visible from within the unit but appear
	// owned by the "nobody" user and group.
	//
	// If the parameter is "identity", user namespacing is set up with an identity mapping for the first 65536 UIDs/GIDs. Any UIDs/GIDs
	// above 65536 will be mapped to the "nobody" user and group, respectively. While this does not provide UID/GID isolation,
	// since all UIDs/GIDs are chosen identically it does provide process capability isolation, and hence is often a good choice
	// if proper user namespacing with distinct UID maps is not appropriate.
	//
	// If this mode is enabled, all unit processes are run without privileges in the host user namespace (regardless if the unit's
	// own user/group is "root" or not). Specifically this means that the process will have zero process capabilities on the host's
	// user namespace, but full capabilities within the service's user namespace. Settings such as CapabilityBoundingSet=
	// will affect only the latter, and there's no way to acquire additional capabilities in the host's user namespace.
	//
	// When this setting is set up by a per-user instance of the service manager, the mapping of the "root" user and group to itself
	// is omitted (unless the user manager is root). Additionally, in the per-user instance manager case, the user namespace
	// will be set up before most other namespaces. This means that combining PrivateUsers=true with other namespaces will enable
	// use of features not normally supported by the per-user instances of the service manager.
	//
	// This setting is particularly useful in conjunction with RootDirectory=/RootImage=, as the need to synchronize the user
	// and group databases in the root directory and on the host is reduced, as the only users and groups who need to be matched are
	// "root", "nobody" and the unit's own user and group.
	//
	// Note that the implementation of this setting might be impossible (for example if user namespaces are not available), and
	// the unit should be written in a way that does not solely rely on this setting for security.
	//
	// Added in version 232.
	PrivateUsers systemdconf.Value

	// Takes a boolean argument. When set, sets up a new UTS namespace for the executed processes. In addition, changing hostname
	// or domainname is prevented. Defaults to off.
	//
	// Note that the implementation of this setting might be impossible (for example if UTS namespaces are not available), and
	// the unit should be written in a way that does not solely rely on this setting for security.
	//
	// Note that when this option is enabled for a service hostname changes no longer propagate from the system into the service,
	// it is hence not suitable for services that need to take notice of system hostname changes dynamically.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 242.
	ProtectHostname systemdconf.Value

	// Takes a boolean argument. If set, writes to the hardware clock or system clock will be denied. Defaults to off. Enabling
	// this option removes CAP_SYS_TIME and CAP_WAKE_ALARM from the capability bounding set for this unit, installs a system
	// call filter to block calls that can set the clock, and DeviceAllow=char-rtc r is implied. Note that the system calls are
	// blocked altogether, the filter does not take into account that some of the calls can be used to read the clock state with some
	// parameter combinations. Effectively, /dev/rtc0, /dev/rtc1, etc. are made read-only to the service. See systemd.resource-control
	// for the details about DeviceAllow=.
	//
	// It is recommended to turn this on for most services that do not need modify the clock or check its state.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 245.
	ProtectClock systemdconf.Value

	// Takes a boolean argument. If true, kernel variables accessible through /proc/sys/, /sys/, /proc/sysrq-trigger, /proc/latency_stats,
	// /proc/acpi, /proc/timer_stats, /proc/fs and /proc/irq will be made read-only and /proc/kallsyms as well as /proc/kcore
	// will be inaccessible to all processes of the unit. Usually, tunable kernel variables should be initialized only at boot-time,
	// for example with the sysctl.d mechanism. Few services need to write to these at runtime; it is hence recommended to turn
	// this on for most services. For this setting the same restrictions regarding mount propagation and privileges apply as
	// for ReadOnlyPaths= and related calls, see above. Defaults to off. Note that this option does not prevent indirect changes
	// to kernel tunables affected by IPC calls to other processes. However, InaccessiblePaths= may be used to make relevant
	// IPC file system objects inaccessible. If ProtectKernelTunables= is set, MountAPIVFS=yes is implied.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 232.
	ProtectKernelTunables systemdconf.Value

	// Takes a boolean argument. If true, explicit module loading will be denied. This allows module load and unload operations
	// to be turned off on modular kernels. It is recommended to turn this on for most services that do not need special file systems
	// or extra kernel modules to work. Defaults to off. Enabling this option removes CAP_SYS_MODULE from the capability bounding
	// set for the unit, and installs a system call filter to block module system calls, also /usr/lib/modules is made inaccessible.
	// For this setting the same restrictions regarding mount propagation and privileges apply as for ReadOnlyPaths= and related
	// calls, see above. Note that limited automatic module loading due to user configuration or kernel mapping tables might
	// still happen as side effect of requested user operations, both privileged and unprivileged. To disable module auto-load
	// feature please see sysctl.d kernel.modules_disabled mechanism and /proc/sys/kernel/modules_disabled documentation.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 232.
	ProtectKernelModules systemdconf.Value

	// Takes a boolean argument. If true, access to the kernel log ring buffer will be denied. It is recommended to turn this on for
	// most services that do not need to read from or write to the kernel log ring buffer. Enabling this option removes CAP_SYSLOG
	// from the capability bounding set for this unit, and installs a system call filter to block the syslog system call (not to
	// be confused with the libc API syslog for userspace logging). The kernel exposes its log buffer to userspace via /dev/kmsg
	// and /proc/kmsg. If enabled, these are made inaccessible to all the processes in the unit.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 244.
	ProtectKernelLogs systemdconf.Value

	// Takes a boolean argument or the special values "private" or "strict". If true, the Linux Control Groups (cgroups) hierarchies
	// accessible through /sys/fs/cgroup/ will be made read-only to all processes of the unit. If set to "private", the unit will
	// run in a cgroup namespace with a private writable mount of /sys/fs/cgroup/. If set to "strict", the unit will run in a cgroup
	// namespace with a private read-only mount of /sys/fs/cgroup/. Defaults to off. If ProtectControlGroups= is set, MountAPIVFS=yes
	// is implied. Note "private" and "strict" are downgraded to false and true respectively unless the system is using the unified
	// control group hierarchy and the kernel supports cgroup namespaces.
	//
	// Except for container managers no services should require write access to the control groups hierarchies; it is hence recommended
	// to set ProtectControlGroups= to true or "strict" for most services. For this setting the same restrictions regarding
	// mount propagation and privileges apply as for ReadOnlyPaths= and related settings, see above.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 232.
	ProtectControlGroups systemdconf.Value

	// Restricts the set of socket address families accessible to the processes of this unit. Takes "none", or a space-separated
	// list of address family names to allow-list, such as AF_UNIX, AF_INET or AF_INET6. When "none" is specified, then all address
	// families will be denied. When prefixed with "~" the listed address families will be applied as deny list, otherwise as allow
	// list.
	//
	// By default, no restrictions apply, all address families are accessible to processes. If assigned the empty string, any
	// previous address family restriction changes are undone. This setting does not affect commands prefixed with "+".
	//
	// Use this option to limit exposure of processes to remote access, in particular via exotic and sensitive network protocols,
	// such as AF_PACKET. Note that in most cases, the local AF_UNIX address family should be included in the configured allow
	// list as it is frequently used for local communication, including for syslog logging.
	//
	// Note that this restricts access to the socket system call only. Sockets passed into the process by other means (for example,
	// by using socket activation with socket units, see systemd.socket) are unaffected. Also, sockets created with socketpair()
	// (which creates connected AF_UNIX sockets) or the io_uring functions, are not affected. Thus, it is recommended to combined
	// this setting with SystemCallFilter=@service, to only allow a limited subset of system calls.
	//
	// Note that this option is limited to some ABIs, in particular x86-64, but currently has no effect on 32-bit x86, s390, s390x,
	// mips, mips-le, ppc, ppc-le, ppc64, or ppc64-le, and is ignored. On systems supporting multiple ABIs (such as x86/x86-64)
	// it is recommended to turn off alternative ABIs for services, so that they cannot be used to circumvent the restrictions
	// of this option. Specifically, it is recommended to combine this option with SystemCallArchitectures=native or similar.
	//
	// Added in version 211.
	RestrictAddressFamilies systemdconf.Value

	// Restricts the set of filesystems processes of this unit can open files on. Takes a space-separated list of filesystem names.
	// Any filesystem listed is made accessible to the unit's processes, access to filesystem types not listed is prohibited
	// (allow-listing). If the first character of the list is "~", the effect is inverted: access to the filesystems listed is
	// prohibited (deny-listing). If the empty string is assigned, access to filesystems is not restricted.
	//
	// If you specify both types of this option (i.e. allow-listing and deny-listing), the first encountered will take precedence
	// and will dictate the default action (allow access to the filesystem or deny it). Then the next occurrences of this option
	// will add or delete the listed filesystems from the set of the restricted filesystems, depending on its type and the default
	// action.
	//
	// Example: if a unit has the following,
	//
	//	RestrictFileSystems=ext4 tmpfs RestrictFileSystems=ext2 ext4
	//
	// then access to ext4, tmpfs, and ext2 is allowed and access to other filesystems is denied.
	//
	// Example: if a unit has the following,
	//
	//	RestrictFileSystems=ext4 tmpfs RestrictFileSystems=~ext4
	//
	// then only access tmpfs is allowed.
	//
	// Example: if a unit has the following,
	//
	//	RestrictFileSystems=~ext4 tmpfs RestrictFileSystems=ext4
	//
	// then only access to tmpfs is denied.
	//
	// As the number of possible filesystems is large, predefined sets of filesystems are provided. A set starts with "@" character,
	// followed by name of the set.
	//
	//	+-------------------+----------------------------------------------------------------------------------+
	//	|        SET        |                                   DESCRIPTION                                    |
	//	+-------------------+----------------------------------------------------------------------------------+
	//	| @basic-api        | Basic filesystem API.                                                            |
	//	| @auxiliary-api    | Auxiliary filesystem API.                                                        |
	//	| @common-block     | Common block device filesystems.                                                 |
	//	| @historical-block | Historical block device filesystems.                                             |
	//	| @network          | Well-known network filesystems.                                                  |
	//	| @privileged-api   | Privileged filesystem API.                                                       |
	//	| @temporary        | Temporary filesystems: tmpfs, ramfs.                                             |
	//	| @known            | All known filesystems defined by the kernel. This list is defined statically in  |
	//	|                   | systemd based on a kernel version that was available when this systemd version   |
	//	|                   | was released. It will become progressively more out-of-date as the kernel is     |
	//	|                   | updated.                                                                         |
	//	+-------------------+----------------------------------------------------------------------------------+
	//
	// Use systemd-analyze's filesystems command to retrieve a list of filesystems defined on the local system.
	//
	// Note that this setting might not be supported on some systems (for example if the LSM eBPF hook is not enabled in the underlying
	// kernel or if not using the unified control group hierarchy). In that case this setting has no effect.
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 250.
	RestrictFileSystems systemdconf.Value

	// Restricts access to Linux namespace functionality for the processes of this unit. For details about Linux namespaces,
	// see namespaces. Either takes a boolean argument, or a space-separated list of namespace type identifiers. If false (the
	// default), no restrictions on namespace creation and switching are made. If true, access to any kind of namespacing is prohibited.
	// Otherwise, a space-separated list of namespace type identifiers must be specified, consisting of any combination of:
	// cgroup, ipc, net, mnt, pid, user and uts. Any namespace type listed is made accessible to the unit's processes, access to
	// namespace types not listed is prohibited (allow-listing). By prepending the list with a single tilde character ("~")
	// the effect may be inverted: only the listed namespace types will be made inaccessible, all unlisted ones are permitted
	// (deny-listing). If the empty string is assigned, the default namespace restrictions are applied, which is equivalent
	// to false. This option may appear more than once, in which case the namespace types are merged by OR, or by AND if the lines are
	// prefixed with "~" (see examples below). Internally, this setting limits access to the unshare, clone and setns system
	// calls, taking the specified flags parameters into account. Note that — if this option is used — in addition to restricting
	// creation and switching of the specified types of namespaces (or all of them, if true) access to the setns() system call with
	// a zero flags parameter is prohibited. This setting is only supported on x86, x86-64, mips, mips-le, mips64, mips64-le,
	// mips64-n32, mips64-le-n32, ppc64, ppc64-le, s390 and s390x, and enforces no restrictions on other architectures.
	//
	// Example: if a unit has the following,
	//
	//	RestrictNamespaces=cgroup ipc RestrictNamespaces=cgroup net
	//
	// then cgroup, ipc, and net are set. If the second line is prefixed with "~", e.g.,
	//
	//	RestrictNamespaces=cgroup ipc RestrictNamespaces=~cgroup net
	//
	// then, only ipc is set.
	//
	// Added in version 233.
	RestrictNamespaces systemdconf.Value

	// Takes a boolean argument. If set, locks down the personality system call so that the kernel execution domain may not be changed
	// from the default or the personality selected with Personality= directive. This may be useful to improve security, because
	// odd personality emulations may be poorly tested and source of vulnerabilities.
	//
	// Added in version 235.
	LockPersonality systemdconf.Value

	// Takes a boolean argument. If set, attempts to create memory mappings that are writable and executable at the same time,
	// or to change existing memory mappings to become executable, or mapping shared memory segments as executable, are prohibited.
	// Specifically, a system call filter is added (or preferably, an equivalent kernel check is enabled with prctl) that rejects
	// mmap system calls with both PROT_EXEC and PROT_WRITE set, mprotect or pkey_mprotect system calls with PROT_EXEC set and
	// shmat system calls with SHM_EXEC set. Note that this option is incompatible with programs and libraries that generate
	// program code dynamically at runtime, including JIT execution engines, executable stacks, and code "trampoline" feature
	// of various C compilers. This option improves service security, as it makes harder for software exploits to change running
	// code dynamically. However, the protection can be circumvented, if the service can write to a filesystem, which is not mounted
	// with noexec (such as /dev/shm), or it can use memfd_create(). This can be prevented by making such file systems inaccessible
	// to the service (e.g. InaccessiblePaths=/dev/shm) and installing further system call filters (SystemCallFilter=~memfd_create).
	// Note that this feature is fully available on x86-64, and partially on x86. Specifically, the shmat() protection is not
	// available on x86. Note that on systems supporting multiple ABIs (such as x86/x86-64) it is recommended to turn off alternative
	// ABIs for services, so that they cannot be used to circumvent the restrictions of this option. Specifically, it is recommended
	// to combine this option with SystemCallArchitectures=native or similar.
	//
	// Added in version 231.
	MemoryDenyWriteExecute systemdconf.Value

	// Takes a boolean argument. If set, any attempts to enable realtime scheduling in a process of the unit are refused. This restricts
	// access to realtime task scheduling policies such as SCHED_FIFO, SCHED_RR or SCHED_DEADLINE. See sched for details about
	// these scheduling policies. Realtime scheduling policies may be used to monopolize CPU time for longer periods of time,
	// and may hence be used to lock up or otherwise trigger Denial-of-Service situations on the system. It is hence recommended
	// to restrict access to realtime scheduling to the few programs that actually require them. Defaults to off.
	//
	// Added in version 231.
	RestrictRealtime systemdconf.Value

	// Takes a boolean argument. If set, any attempts to set the set-user-ID (SUID) or set-group-ID (SGID) bits on files or directories
	// will be denied (for details on these bits see inode). As the SUID/SGID bits are mechanisms to elevate privileges, and allow
	// users to acquire the identity of other users, it is recommended to restrict creation of SUID/SGID files to the few programs
	// that actually require them. Note that this restricts marking of any type of file system object with these bits, including
	// both regular files and directories (where the SGID is a different meaning than for files, see documentation). This option
	// is implied if DynamicUser= is enabled. Defaults to off.
	//
	// Added in version 242.
	RestrictSUIDSGID systemdconf.Value

	// Takes a boolean parameter. If set, all System V and POSIX IPC objects owned by the user and group the processes of this unit
	// are run as are removed when the unit is stopped. This setting only has an effect if at least one of User=, Group= and DynamicUser=
	// are used. It has no effect on IPC objects owned by the root user. Specifically, this removes System V semaphores, as well
	// as System V and POSIX shared memory segments and message queues. If multiple units use the same user or group the IPC objects
	// are removed when the last of these units is stopped. This setting is implied if DynamicUser= is set.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 232.
	RemoveIPC systemdconf.Value

	// Takes a boolean parameter. If set, the processes of this unit will be run in their own private file system (mount) namespace
	// with all mount propagation from the processes towards the host's main file system namespace turned off. This means any
	// file system mount points established or removed by the unit's processes will be private to them and not be visible to the
	// host. However, file system mount points established or removed on the host will be propagated to the unit's processes.
	// See mount_namespaces for details on file system namespaces. Defaults to off.
	//
	// When turned on, this executes three operations for each invoked process: a new CLONE_NEWNS namespace is created, after
	// which all existing mounts are remounted to MS_SLAVE to disable propagation from the unit's processes to the host (but leaving
	// propagation in the opposite direction in effect). Finally, the mounts are remounted again to the propagation mode configured
	// with MountFlags=, see below.
	//
	// File system namespaces are set up individually for each process forked off by the service manager. Mounts established
	// in the namespace of the process created by ExecStartPre= will hence be cleaned up automatically as soon as that process
	// exits and will not be available to subsequent processes forked off for ExecStart= (and similar applies to the various other
	// commands configured for units). Similarly, JoinsNamespaceOf= does not permit sharing kernel mount namespaces between
	// units, it only enables sharing of the /tmp/ and /var/tmp/ directories.
	//
	// Other file system namespace unit settings — PrivateTmp=, PrivateDevices=, ProtectSystem=, ProtectHome=, ReadOnlyPaths=,
	// InaccessiblePaths=, ReadWritePaths=, BindPaths=, BindReadOnlyPaths=, … — also enable file system namespacing in
	// a fashion equivalent to this option. Hence it is primarily useful to explicitly request this behaviour if none of the other
	// settings are used.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	//
	// Added in version 239.
	PrivateMounts systemdconf.Value

	// Takes a mount propagation setting: shared, slave or private, which controls whether file system mount points in the file
	// system namespaces set up for this unit's processes will receive or propagate mounts and unmounts from other file system
	// namespaces. See mount for details on mount propagation, and the three propagation flags in particular.
	//
	// This setting only controls the final propagation setting in effect on all mount points of the file system namespace created
	// for each process of this unit. Other file system namespacing unit settings (see the discussion in PrivateMounts= above)
	// will implicitly disable mount and unmount propagation from the unit's processes towards the host by changing the propagation
	// setting of all mount points in the unit's file system namespace to slave first. Setting this option to shared does not reestablish
	// propagation in that case.
	//
	// If not set – but file system namespaces are enabled through another file system namespace unit setting – shared mount propagation
	// is used, but — as mentioned — as slave is applied first, propagation from the unit's processes to the host is still turned
	// off.
	//
	// It is not recommended to use private mount propagation for units, as this means temporary mounts (such as removable media)
	// of the host will stay mounted and thus indefinitely busy in forked off processes, as unmount propagation events will not
	// be received by the file system namespace of the unit.
	//
	// Usually, it is best to leave this setting unmodified, and use higher level file system namespacing options instead, in
	// particular PrivateMounts=, see above.
	//
	// This option is only available for system services, or for services running in per-user instances of the service manager
	// in which case PrivateUsers= is implicitly enabled (requires unprivileged user namespaces support to be enabled in the
	// kernel via the "kernel.unprivileged_userns_clone=" sysctl).
	MountFlags systemdconf.Value

	// Takes a space-separated list of system call names or system call groups. If this setting is used, system calls executed
	// by the unit processes except for the listed ones will result in the system call being denied (allow-listing). If the first
	// character of the list is "~", the effect is inverted: only the listed system calls will be denied (deny-listing). This option
	// may be specified more than once, in which case the filter masks are merged. If the empty string is assigned, the filter is
	// reset, all prior assignments will have no effect.
	//
	// Commands prefixed with "+" are not subject to filtering. The execve(), exit(), exit_group(), getrlimit(), rt_sigreturn(),
	// sigreturn() system calls and the system calls for querying time and sleeping are implicitly allow-listed and do not need
	// to be listed explicitly.
	//
	// The default action when a system call is denied is to terminate the processes with a SIGSYS signal. This can changed using
	// SystemCallErrorNumber=, see below. In addition, deny-listed system calls and system call groups may optionally be suffixed
	// with a colon (":") and an argument in the same format as SystemCallErrorNumber=, to take this action when the matching system
	// call is made. This takes precedence over the action specified in SystemCallErrorNumber=.
	//
	// This feature makes use of the Secure Computing Mode 2 interfaces of the kernel ('seccomp filtering') and is useful for enforcing
	// a minimal sandboxing environment.
	//
	// Note that on systems supporting multiple ABIs (such as x86/x86-64) it is recommended to turn off alternative ABIs for services,
	// so that they cannot be used to circumvent the restrictions of this option. Specifically, it is recommended to combine this
	// option with SystemCallArchitectures=native or similar.
	//
	// Note that strict system call filters may impact execution and error handling code paths of the service invocation. Specifically,
	// access to the execve() system call is required for the execution of the service binary — if it is blocked service invocation
	// will necessarily fail. Also, if execution of the service binary fails for some reason (for example: missing service executable),
	// the error handling logic might require access to an additional set of system calls in order to process and log this failure
	// correctly. It might be necessary to temporarily disable system call filters in order to allow debugging of such failures.
	//
	// If you specify both types of this option (i.e. allow-listing and deny-listing), the first encountered will take precedence
	// and will dictate the default action (termination or approval of a system call). Then the next occurrences of this option
	// will add or delete the listed system calls from the set of the filtered system calls, depending of its type and the default
	// action. (For example, if you have started with an allow list rule for read() and write(), and right after it add a deny list
	// rule for write(), then write() will be removed from the set.)
	//
	// As the number of possible system calls is large, predefined groups of system calls are provided. A group starts with "@"
	// character, followed by name of the set.
	//
	//	+-----------------+----------------------------------------------------------------------------------+
	//	|       SET       |                                   DESCRIPTION                                    |
	//	+-----------------+----------------------------------------------------------------------------------+
	//	| @aio            | Asynchronous I/O (io_setup, io_submit, and related calls)                        |
	//	| @basic-io       | System calls for basic I/O: reading, writing, seeking, file descriptor           |
	//	|                 | duplication and closing (read, write, and related calls)                         |
	//	| @chown          | Changing file ownership (chown, fchownat, and related calls)                     |
	//	| @clock          | System calls for changing the system clock (adjtimex, settimeofday, and related  |
	//	|                 | calls)                                                                           |
	//	| @cpu-emulation  | System calls for CPU emulation functionality (vm86 and related calls)            |
	//	| @debug          | Debugging, performance monitoring and tracing functionality (ptrace,             |
	//	|                 | perf_event_open and related calls)                                               |
	//	| @file-system    | File system operations: opening, creating files and directories for read and     |
	//	|                 | write, renaming and removing them, reading file properties, or creating hard and |
	//	|                 | symbolic links                                                                   |
	//	| @io-event       | Event loop system calls (poll, select, epoll, eventfd and related calls)         |
	//	| @ipc            | Pipes, SysV IPC, POSIX Message Queues and other IPC (mq_overview, svipc)         |
	//	| @keyring        | Kernel keyring access (keyctl and related calls)                                 |
	//	| @memlock        | Locking of memory in RAM (mlock, mlockall and related calls)                     |
	//	| @module         | Loading and unloading of kernel modules (init_module, delete_module and related  |
	//	|                 | calls)                                                                           |
	//	| @mount          | Mounting and unmounting of file systems (mount, chroot, and related calls)       |
	//	| @network-io     | Socket I/O (including local AF_UNIX): socket, unix                               |
	//	| @obsolete       | Unusual, obsolete or unimplemented (create_module, gtty, …)                      |
	//	| @pkey           | System calls that deal with memory protection keys (pkeys)                       |
	//	| @privileged     | All system calls which need super-user capabilities (capabilities)               |
	//	| @process        | Process control, execution, namespacing operations (clone, kill, namespaces, …)  |
	//	| @raw-io         | Raw I/O port access (ioperm, iopl, pciconfig_read(), …)                          |
	//	| @reboot         | System calls for rebooting and reboot preparation (reboot, kexec(), …)           |
	//	| @resources      | System calls for changing resource limits, memory and scheduling parameters      |
	//	|                 | (setrlimit, setpriority, …)                                                      |
	//	| @sandbox        | System calls for sandboxing programs (seccomp, Landlock system calls, …)         |
	//	| @setuid         | System calls for changing user ID and group ID credentials, (setuid, setgid,     |
	//	|                 | setresuid, …)                                                                    |
	//	| @signal         | System calls for manipulating and handling process signals (signal, sigprocmask, |
	//	|                 | …)                                                                               |
	//	| @swap           | System calls for enabling/disabling swap devices (swapon, swapoff)               |
	//	| @sync           | Synchronizing files and memory to disk (fsync, msync, and related calls)         |
	//	| @system-service | A reasonable set of system calls used by common system services, excluding any   |
	//	|                 | special purpose calls. This is the recommended starting point for allow-listing  |
	//	|                 | system calls for system services, as it contains what is typically needed by     |
	//	|                 | system services, but excludes overly specific interfaces. For example, the       |
	//	|                 | following APIs are excluded: "@clock", "@mount", "@swap", "@reboot".             |
	//	| @timer          | System calls for scheduling operations by time (alarm, timer_create, …)          |
	//	| @known          | All system calls defined by the kernel. This list is defined statically in       |
	//	|                 | systemd based on a kernel version that was available when this systemd version   |
	//	|                 | was released. It will become progressively more out-of-date as the kernel is     |
	//	|                 | updated.                                                                         |
	//	+-----------------+----------------------------------------------------------------------------------+
	//
	// Note, that as new system calls are added to the kernel, additional system calls might be added to the groups above. Contents
	// of the sets may also change between systemd versions. In addition, the list of system calls depends on the kernel version
	// and architecture for which systemd was compiled. Use systemd-analyze syscall-filter to list the actual list of system
	// calls in each filter.
	//
	// Generally, allow-listing system calls (rather than deny-listing) is the safer mode of operation. It is recommended to
	// enforce system call allow lists for all long-running system services. Specifically, the following lines are a relatively
	// safe basic choice for the majority of system services:
	//
	//	[Service] SystemCallFilter=@system-service SystemCallErrorNumber=EPERM
	//
	// Note that various kernel system calls are defined redundantly: there are multiple system calls for executing the same
	// operation. For example, the pidfd_send_signal() system call may be used to execute operations similar to what can be done
	// with the older kill() system call, hence blocking the latter without the former only provides weak protection. Since new
	// system calls are added regularly to the kernel as development progresses, keeping system call deny lists comprehensive
	// requires constant work. It is thus recommended to use allow-listing instead, which offers the benefit that new system
	// calls are by default implicitly blocked until the allow list is updated.
	//
	// Also note that a number of system calls are required to be accessible for the dynamic linker to work. The dynamic linker is
	// required for running most regular programs (specifically: all dynamic ELF binaries, which is how most distributions
	// build packaged programs). This means that blocking these system calls (which include open(), openat() or mmap()) will
	// make most programs typically shipped with generic distributions unusable.
	//
	// It is recommended to combine the file system namespacing related options with SystemCallFilter=~@mount, in order to
	// prohibit the unit's processes to undo the mappings. Specifically these are the options PrivateTmp=, PrivateDevices=,
	// ProtectSystem=, ProtectHome=, ProtectKernelTunables=, ProtectControlGroups=, ProtectKernelLogs=, ProtectClock=,
	// ReadOnlyPaths=, InaccessiblePaths= and ReadWritePaths=.
	//
	// Added in version 187.
	SystemCallFilter systemdconf.Value

	// Takes an "errno" error number (between 1 and 4095) or errno name such as EPERM, EACCES or EUCLEAN, to return when the system
	// call filter configured with SystemCallFilter= is triggered, instead of terminating the process immediately. See errno
	// for a full list of error codes. When this setting is not used, or when the empty string or the special setting "kill" is assigned,
	// the process will be terminated immediately when the filter is triggered.
	//
	// Added in version 209.
	SystemCallErrorNumber systemdconf.Value

	// Takes a space-separated list of architecture identifiers to include in the system call filter. The known architecture
	// identifiers are the same as for ConditionArchitecture= described in systemd.unit, as well as x32, mips64-n32, mips64-le-n32,
	// and the special identifier native. The special identifier native implicitly maps to the native architecture of the system
	// (or more precisely: to the architecture the system manager is compiled for). By default, this option is set to the empty
	// list, i.e. no filtering is applied.
	//
	// If this setting is used, processes of this unit will only be permitted to call native system calls, and system calls of the
	// specified architectures. For the purposes of this option, the x32 architecture is treated as including x86-64 system
	// calls. However, this setting still fulfills its purpose, as explained below, on x32.
	//
	// System call filtering is not equally effective on all architectures. For example, on x86 filtering of network socket-related
	// calls is not possible, due to ABI limitations — a limitation that x86-64 does not have, however. On systems supporting multiple
	// ABIs at the same time — such as x86/x86-64 — it is hence recommended to limit the set of permitted system call architectures
	// so that secondary ABIs may not be used to circumvent the restrictions applied to the native ABI of the system. In particular,
	// setting SystemCallArchitectures=native is a good choice for disabling non-native ABIs.
	//
	// System call architectures may also be restricted system-wide via the SystemCallArchitectures= option in the global
	// configuration. See systemd-system.conf for details.
	//
	// Added in version 209.
	SystemCallArchitectures systemdconf.Value

	// Takes a space-separated list of system call names. If this setting is used, all system calls executed by the unit processes
	// for the listed ones will be logged. If the first character of the list is "~", the effect is inverted: all system calls except
	// the listed system calls will be logged. This feature makes use of the Secure Computing Mode 2 interfaces of the kernel ('seccomp
	// filtering') and is useful for auditing or setting up a minimal sandboxing environment. This option may be specified more
	// than once, in which case the filter masks are merged. If the empty string is assigned, the filter is reset, all prior assignments
	// will have no effect. This does not affect commands prefixed with "+".
	//
	// Added in version 247.
	SystemCallLog systemdconf.Value

	// Sets environment variables for executed processes. Each line is unquoted using the rules described in "Quoting" section
	// in systemd.syntax and becomes a list of variable assignments. If you need to assign a value containing spaces or the equals
	// sign to a variable, put quotes around the whole assignment. Variable expansion is not performed inside the strings and
	// the "$" character has no special meaning. Specifier expansion is performed, see the "Specifiers" section in systemd.unit.
	//
	// This option may be specified more than once, in which case all listed variables will be set. If the same variable is listed
	// twice, the later setting will override the earlier setting. If the empty string is assigned to this option, the list of environment
	// variables is reset, all prior assignments have no effect.
	//
	// The names of the variables can contain ASCII letters, digits, and the underscore character. Variable names cannot be empty
	// or start with a digit. In variable values, most characters are allowed, but non-printable characters are currently rejected.
	//
	// Example:
	//
	//	Environment="VAR1=word1 word2" VAR2=word3 "VAR3=$word 5 6"
	//
	// gives three variables "VAR1", "VAR2", "VAR3" with the values "word1 word2", "word3", "$word 5 6".
	//
	// See environ for details about environment variables.
	//
	// Note that environment variables are not suitable for passing secrets (such as passwords, key material, …) to service processes.
	// Environment variables set for a unit are exposed to unprivileged clients via D-Bus IPC, and generally not understood as
	// being data that requires protection. Moreover, environment variables are propagated down the process tree, including
	// across security boundaries (such as setuid/setgid executables), and hence might leak to processes that should not have
	// access to the secret data. Use LoadCredential=, LoadCredentialEncrypted= or SetCredentialEncrypted= (see below)
	// to pass data to unit processes securely.
	Environment systemdconf.Value

	// Similar to Environment=, but reads the environment variables from a text file. The text file should contain newline-separated
	// variable assignments. Empty lines, lines without an "=" separator, or lines starting with ";" or "#" will be ignored, which
	// may be used for commenting. The file must be encoded with UTF-8. Valid characters are unicode scalar values other than unicode
	// noncharacters, U+0000 NUL, and U+FEFF unicode byte order mark. Control codes other than NUL are allowed.
	//
	// In the file, an unquoted value after the "=" is parsed with the same backslash-escape rules as POSIX shell unquoted text,
	// but unlike in a shell, interior whitespace is preserved and quotes after the first non-whitespace character are preserved.
	// Leading and trailing whitespace (space, tab, carriage return) is discarded, but interior whitespace within the line
	// is preserved verbatim. A line ending with a backslash will be continued to the following one, with the newline itself discarded.
	// A backslash "\" followed by any character other than newline will preserve the following character, so that "\\" will become
	// the value "\".
	//
	// In the file, a "'"-quoted value after the "=" can span multiple lines and contain any character verbatim other than single
	// quote, like POSIX shell single-quoted text. No backslash-escape sequences are recognized. Leading and trailing whitespace
	// outside of the single quotes is discarded.
	//
	// In the file, a """-quoted value after the "=" can span multiple lines, and the same escape sequences are recognized as in
	// POSIX shell double-quoted text. Backslash ("\") followed by any of ""\`$" will preserve that character. A backslash followed
	// by newline is a line continuation, and the newline itself is discarded. A backslash followed by any other character is ignored;
	// both the backslash and the following character are preserved verbatim. Leading and trailing whitespace outside of the
	// double quotes is discarded.
	//
	// The argument passed should be an absolute filename or wildcard expression, optionally prefixed with "-", which indicates
	// that if the file does not exist, it will not be read and no error or warning message is logged. This option may be specified
	// more than once in which case all specified files are read. If the empty string is assigned to this option, the list of file
	// to read is reset, all prior assignments have no effect.
	//
	// The files listed with this directive will be read shortly before the process is executed (more specifically, after all
	// processes from a previous unit state terminated. This means you can generate these files in one unit state, and read it with
	// this option in the next. The files are read from the file system of the service manager, before any file system changes like
	// bind mounts take place).
	//
	// Settings from these files override settings made with Environment=. If the same variable is set twice from these files,
	// the files will be read in the order they are specified and the later setting will override the earlier setting.
	EnvironmentFile systemdconf.Value

	// Pass environment variables set for the system service manager to executed processes. Takes a space-separated list of
	// variable names. This option may be specified more than once, in which case all listed variables will be passed. If the empty
	// string is assigned to this option, the list of environment variables to pass is reset, all prior assignments have no effect.
	// Variables specified that are not set for the system manager will not be passed and will be silently ignored. Note that this
	// option is only relevant for the system service manager, as system services by default do not automatically inherit any
	// environment variables set for the service manager itself. However, in case of the user service manager all environment
	// variables are passed to the executed processes anyway, hence this option is without effect for the user service manager.
	//
	// Variables set for invoked processes due to this setting are subject to being overridden by those configured with Environment=
	// or EnvironmentFile=.
	//
	// Example:
	//
	//	PassEnvironment=VAR1 VAR2 VAR3
	//
	// passes three variables "VAR1", "VAR2", "VAR3" with the values set for those variables in PID1.
	//
	// See environ for details about environment variables.
	//
	// Added in version 228.
	PassEnvironment systemdconf.Value

	// Explicitly unset environment variable assignments that would normally be passed from the service manager to invoked
	// processes of this unit. Takes a space-separated list of variable names or variable assignments. This option may be specified
	// more than once, in which case all listed variables/assignments will be unset. If the empty string is assigned to this option,
	// the list of environment variables/assignments to unset is reset. If a variable assignment is specified (that is: a variable
	// name, followed by "=", followed by its value), then any environment variable matching this precise assignment is removed.
	// If a variable name is specified (that is a variable name without any following "=" or value), then any assignment matching
	// the variable name, regardless of its value is removed. Note that the effect of UnsetEnvironment= is applied as final step
	// when the environment list passed to executed processes is compiled. That means it may undo assignments from any configuration
	// source, including assignments made through Environment= or EnvironmentFile=, inherited from the system manager's
	// global set of environment variables, inherited via PassEnvironment=, set by the service manager itself (such as $NOTIFY_SOCKET
	// and such), or set by a PAM module (in case PAMName= is used).
	//
	// See "Environment Variables in Spawned Processes" below for a description of how those settings combine to form the inherited
	// environment. See environ for general information about environment variables.
	//
	// Added in version 235.
	UnsetEnvironment systemdconf.Value

	// Controls where file descriptor 0 (STDIN) of the executed processes is connected to. Takes one of null, tty, tty-force,
	// tty-fail, data, file:path, socket or fd:name.
	//
	// If null is selected, standard input will be connected to /dev/null, i.e. all read attempts by the process will result in
	// immediate EOF.
	//
	// If tty is selected, standard input is connected to a TTY (as configured by TTYPath=, see below) and the executed process
	// becomes the controlling process of the terminal. If the terminal is already being controlled by another process, the executed
	// process waits until the current controlling process releases the terminal.
	//
	// tty-force is similar to tty, but the executed process is forcefully and immediately made the controlling process of the
	// terminal, potentially removing previous controlling processes from the terminal.
	//
	// tty-fail is similar to tty, but if the terminal already has a controlling process start-up of the executed process fails.
	//
	// The data option may be used to configure arbitrary textual or binary data to pass via standard input to the executed process.
	// The data to pass is configured via StandardInputText=/StandardInputData= (see below). Note that the actual file descriptor
	// type passed (memory file, regular file, UNIX pipe, …) might depend on the kernel and available privileges. In any case,
	// the file descriptor is read-only, and when read returns the specified data followed by EOF.
	//
	// The file:path option may be used to connect a specific file system object to standard input. An absolute path following
	// the ":" character is expected, which may refer to a regular file, a FIFO or special file. If an AF_UNIX socket in the file system
	// is specified, a stream socket is connected to it. The latter is useful for connecting standard input of processes to arbitrary
	// system services.
	//
	// The socket option is valid in socket-activated services only, and requires the relevant socket unit file (see systemd.socket
	// for details) to have Accept=yes set, or to specify a single socket only. If this option is set, standard input will be connected
	// to the socket the service was activated from, which is primarily useful for compatibility with daemons designed for use
	// with the traditional inetd socket activation daemon ($LISTEN_FDS (and related) environment variables are not passed
	// when socket value is configured).
	//
	// The fd:name option connects standard input to a specific, named file descriptor provided by a socket unit. The name may
	// be specified as part of this option, following a ":" character (e.g. "fd:foobar"). If no name is specified, the name "stdin"
	// is implied (i.e. "fd" is equivalent to "fd:stdin"). At least one socket unit defining the specified name must be provided
	// via the Sockets= option, and the file descriptor name may differ from the name of its containing socket unit. If multiple
	// matches are found, the first one will be used. See FileDescriptorName= in systemd.socket for more details about named
	// file descriptors and their ordering.
	//
	// This setting defaults to null, unless StandardInputText=/StandardInputData= are set, in which case it defaults to data.
	StandardInput systemdconf.Value

	// Controls where file descriptor 1 (stdout) of the executed processes is connected to. Takes one of inherit, null, tty, journal,
	// kmsg, journal+console, kmsg+console, file:path, append:path, truncate:path, socket or fd:name.
	//
	// inherit duplicates the file descriptor of standard input for standard output.
	//
	// null connects standard output to /dev/null, i.e. everything written to it will be lost.
	//
	// tty connects standard output to a tty (as configured via TTYPath=, see below). If the TTY is used for output only, the executed
	// process will not become the controlling process of the terminal, and will not fail or wait for other processes to release
	// the terminal. Note: if a unit tries to print multiple lines to a TTY during bootup or shutdown, then there's a chance that
	// those lines will be broken up by status messages. SetShowStatus() can be used to prevent this problem. See org.freedesktop.systemd1
	// for details.
	//
	// journal connects standard output with the journal, which is accessible via journalctl. Note that everything that is written
	// to kmsg (see below) is implicitly stored in the journal as well, the specific option listed below is hence a superset of this
	// one. (Also note that any external, additional syslog daemons receive their log data from the journal, too, hence this is
	// the option to use when logging shall be processed with such a daemon.)
	//
	// kmsg connects standard output with the kernel log buffer which is accessible via dmesg, in addition to the journal. The
	// journal daemon might be configured to send all logs to kmsg anyway, in which case this option is no different from journal.
	//
	// journal+console and kmsg+console work in a similar way as the two options above but copy the output to the system console
	// as well.
	//
	// The file:path option may be used to connect a specific file system object to standard output. The semantics are similar
	// to the same option of StandardInput=, see above. If path refers to a regular file on the filesystem, it is opened (created
	// if it does not exist yet using privileges of the user executing the systemd process) for writing at the beginning of the file,
	// but without truncating it. If standard input and output are directed to the same file path, it is opened only once — for reading
	// as well as writing — and duplicated. This is particularly useful when the specified path refers to an AF_UNIX socket in the
	// file system, as in that case only a single stream connection is created for both input and output.
	//
	// append:path is similar to file:path above, but it opens the file in append mode.
	//
	// truncate:path is similar to file:path above, but it truncates the file when opening it. For units with multiple command
	// lines, e.g. Type=oneshot services with multiple ExecStart=, or services with ExecCondition=, ExecStartPre= or ExecStartPost=,
	// the output file is reopened and therefore re-truncated for each command line. If the output file is truncated while another
	// process still has the file open, e.g. by an ExecReload= running concurrently with an ExecStart=, and the other process
	// continues writing to the file without adjusting its offset, then the space between the file pointers of the two processes
	// may be filled with NUL bytes, producing a sparse file. Thus, truncate:path is typically only useful for units where only
	// one process runs at a time, such as services with a single ExecStart= and no ExecStartPost=, ExecReload=, ExecStop= or
	// similar.
	//
	// socket connects standard output to a socket acquired via socket activation. The semantics are similar to the same option
	// of StandardInput=, see above.
	//
	// The fd:name option connects standard output to a specific, named file descriptor provided by a socket unit. A name may be
	// specified as part of this option, following a ":" character (e.g. "fd:foobar"). If no name is specified, the name "stdout"
	// is implied (i.e. "fd" is equivalent to "fd:stdout"). At least one socket unit defining the specified name must be provided
	// via the Sockets= option, and the file descriptor name may differ from the name of its containing socket unit. If multiple
	// matches are found, the first one will be used. See FileDescriptorName= in systemd.socket for more details about named
	// descriptors and their ordering.
	//
	// If the standard output (or error output, see below) of a unit is connected to the journal or the kernel log buffer, the unit
	// will implicitly gain a dependency of type After= on systemd-journald.socket (also see the "Implicit Dependencies" section
	// above). Also note that in this case stdout (or stderr, see below) will be an AF_UNIX stream socket, and not a pipe or FIFO that
	// can be reopened. This means when executing shell scripts the construct echo "hello" > /dev/stderr for writing text to stderr
	// will not work. To mitigate this use the construct echo "hello" >&2 instead, which is mostly equivalent and avoids this pitfall.
	//
	// If StandardInput= is set to one of tty, tty-force, tty-fail, socket, or fd:name, this setting defaults to inherit.
	//
	// In other cases, this setting defaults to the value set with DefaultStandardOutput= in systemd-system.conf, which defaults
	// to journal. Note that setting this parameter might result in additional dependencies to be added to the unit (see above).
	StandardOutput systemdconf.Value

	// Controls where file descriptor 2 (stderr) of the executed processes is connected to. The available options are identical
	// to those of StandardOutput=, with some exceptions: if set to inherit the file descriptor used for standard output is duplicated
	// for standard error, while fd:name will use a default file descriptor name of "stderr".
	//
	// This setting defaults to the value set with DefaultStandardError= in systemd-system.conf, which defaults to inherit.
	// Note that setting this parameter might result in additional dependencies to be added to the unit (see above).
	StandardError systemdconf.Value

	// Configures arbitrary textual or binary data to pass via file descriptor 0 (STDIN) to the executed processes. These settings
	// have no effect unless StandardInput= is set to data (which is the default if StandardInput= is not set otherwise, but StandardInputText=/StandardInputData=
	// is). Use this option to embed process input data directly in the unit file.
	//
	// StandardInputText= accepts arbitrary textual data. C-style escapes for special characters as well as the usual "%"-specifiers
	// are resolved. Each time this setting is used the specified text is appended to the per-unit data buffer, followed by a newline
	// character (thus every use appends a new line to the end of the buffer). Note that leading and trailing whitespace of lines
	// configured with this option is removed. If an empty line is specified the buffer is cleared (hence, in order to insert an
	// empty line, add an additional "\n" to the end or beginning of a line).
	//
	// StandardInputData= accepts arbitrary binary data, encoded in Base64. No escape sequences or specifiers are resolved.
	// Any whitespace in the encoded version is ignored during decoding.
	//
	// Note that StandardInputText= and StandardInputData= operate on the same data buffer, and may be mixed in order to configure
	// both binary and textual data for the same input stream. The textual or binary data is joined strictly in the order the settings
	// appear in the unit file. Assigning an empty string to either will reset the data buffer.
	//
	// Please keep in mind that in order to maintain readability long unit file settings may be split into multiple lines, by suffixing
	// each line (except for the last) with a "\" character (see systemd.unit for details). This is particularly useful for large
	// data configured with these two options. Example:
	//
	//	… StandardInput=data StandardInputData=V2XigLJyZSBubyBzdHJhbmdlcnMgdG8gbG92ZQpZb3Uga25vdyB0aGUgcnVsZXMgYW5kIHNvIGRv
	//	\ IEkKQSBmdWxsIGNvbW1pdG1lbnQncyB3aGF0IEnigLJtIHRoaW5raW5nIG9mCllvdSB3b3VsZG4n \ dCBnZXQgdGhpcyBmcm9tIGFueSBvdGhlciBndXkKSSBqdXN0IHdhbm5hIHRlbGwgeW91IGhvdyBJ
	//	\ J20gZmVlbGluZwpHb3R0YSBtYWtlIHlvdSB1bmRlcnN0YW5kCgpOZXZlciBnb25uYSBnaXZlIHlv \ dSB1cApOZXZlciBnb25uYSBsZXQgeW91IGRvd24KTmV2ZXIgZ29ubmEgcnVuIGFyb3VuZCBhbmQg
	//	\ ZGVzZXJ0IHlvdQpOZXZlciBnb25uYSBtYWtlIHlvdSBjcnkKTmV2ZXIgZ29ubmEgc2F5IGdvb2Ri \ eWUKTmV2ZXIgZ29ubmEgdGVsbCBhIGxpZSBhbmQgaHVydCB5b3UK
	//	…
	//
	// Added in version 236.
	StandardInputText systemdconf.Value

	// Configures arbitrary textual or binary data to pass via file descriptor 0 (STDIN) to the executed processes. These settings
	// have no effect unless StandardInput= is set to data (which is the default if StandardInput= is not set otherwise, but StandardInputText=/StandardInputData=
	// is). Use this option to embed process input data directly in the unit file.
	//
	// StandardInputText= accepts arbitrary textual data. C-style escapes for special characters as well as the usual "%"-specifiers
	// are resolved. Each time this setting is used the specified text is appended to the per-unit data buffer, followed by a newline
	// character (thus every use appends a new line to the end of the buffer). Note that leading and trailing whitespace of lines
	// configured with this option is removed. If an empty line is specified the buffer is cleared (hence, in order to insert an
	// empty line, add an additional "\n" to the end or beginning of a line).
	//
	// StandardInputData= accepts arbitrary binary data, encoded in Base64. No escape sequences or specifiers are resolved.
	// Any whitespace in the encoded version is ignored during decoding.
	//
	// Note that StandardInputText= and StandardInputData= operate on the same data buffer, and may be mixed in order to configure
	// both binary and textual data for the same input stream. The textual or binary data is joined strictly in the order the settings
	// appear in the unit file. Assigning an empty string to either will reset the data buffer.
	//
	// Please keep in mind that in order to maintain readability long unit file settings may be split into multiple lines, by suffixing
	// each line (except for the last) with a "\" character (see systemd.unit for details). This is particularly useful for large
	// data configured with these two options. Example:
	//
	//	… StandardInput=data StandardInputData=V2XigLJyZSBubyBzdHJhbmdlcnMgdG8gbG92ZQpZb3Uga25vdyB0aGUgcnVsZXMgYW5kIHNvIGRv
	//	\ IEkKQSBmdWxsIGNvbW1pdG1lbnQncyB3aGF0IEnigLJtIHRoaW5raW5nIG9mCllvdSB3b3VsZG4n \ dCBnZXQgdGhpcyBmcm9tIGFueSBvdGhlciBndXkKSSBqdXN0IHdhbm5hIHRlbGwgeW91IGhvdyBJ
	//	\ J20gZmVlbGluZwpHb3R0YSBtYWtlIHlvdSB1bmRlcnN0YW5kCgpOZXZlciBnb25uYSBnaXZlIHlv \ dSB1cApOZXZlciBnb25uYSBsZXQgeW91IGRvd24KTmV2ZXIgZ29ubmEgcnVuIGFyb3VuZCBhbmQg
	//	\ ZGVzZXJ0IHlvdQpOZXZlciBnb25uYSBtYWtlIHlvdSBjcnkKTmV2ZXIgZ29ubmEgc2F5IGdvb2Ri \ eWUKTmV2ZXIgZ29ubmEgdGVsbCBhIGxpZSBhbmQgaHVydCB5b3UK
	//	…
	//
	// Added in version 236.
	StandardInputData systemdconf.Value

	// Configures filtering by log level of log messages generated by this unit. Takes a syslog log level, one of emerg (lowest
	// log level, only highest priority messages), alert, crit, err, warning, notice, info, debug (highest log level, also lowest
	// priority messages). See syslog for details. By default no filtering is applied (i.e. the default maximum log level is debug).
	// Use this option to configure the logging system to drop log messages of a specific service above the specified level. For
	// example, set LogLevelMax=info in order to turn off debug logging of a particularly chatty unit. Note that the configured
	// level is applied to any log messages written by any of the processes belonging to this unit, as well as any log messages written
	// by the system manager process (PID 1) in reference to this unit, sent via any supported logging protocol. The filtering
	// is applied early in the logging pipeline, before any kind of further processing is done. Moreover, messages which pass
	// through this filter successfully might still be dropped by filters applied at a later stage in the logging subsystem. For
	// example, MaxLevelStore= configured in journald.conf might prohibit messages of higher log levels to be stored on disk,
	// even though the per-unit LogLevelMax= permitted it to be processed.
	//
	// Added in version 236.
	LogLevelMax systemdconf.Value

	// Configures additional log metadata fields to include in all log records generated by processes associated with this unit,
	// including systemd. This setting takes one or more journal field assignments in the format "FIELD=VALUE" separated by
	// whitespace. See systemd.journal-fields for details on the journal field concept. Even though the underlying journal
	// implementation permits binary field values, this setting accepts only valid UTF-8 values. To include space characters
	// in a journal field value, enclose the assignment in double quotes ("). The usual specifiers are expanded in all assignments
	// (see below). Note that this setting is not only useful for attaching additional metadata to log records of a unit, but given
	// that all fields and values are indexed may also be used to implement cross-unit log record matching. Assign an empty string
	// to reset the list.
	//
	// Note that this functionality is currently only available in system services, not in per-user services.
	//
	// Added in version 236.
	LogExtraFields systemdconf.Value

	// Configures the rate limiting that is applied to log messages generated by this unit. If, in the time interval defined by
	// LogRateLimitIntervalSec=, more messages than specified in LogRateLimitBurst= are logged by a service, all further
	// messages within the interval are dropped until the interval is over. A message about the number of dropped messages is generated.
	// The time specification for LogRateLimitIntervalSec= may be specified in the following units: "s", "min", "h", "ms",
	// "us". See systemd.time for details. The default settings are set by RateLimitIntervalSec= and RateLimitBurst= configured
	// in journald.conf. Note that this only applies to log messages that are processed by the logging subsystem, i.e. by systemd-journald.service.
	// This means that if you connect a service's stderr directly to a file via StandardOutput=file:… or a similar setting, the
	// rate limiting will not be applied to messages written that way (but it will be enforced for messages generated via syslog
	// and similar functions).
	//
	// Added in version 240.
	LogRateLimitIntervalSec systemdconf.Value

	// Configures the rate limiting that is applied to log messages generated by this unit. If, in the time interval defined by
	// LogRateLimitIntervalSec=, more messages than specified in LogRateLimitBurst= are logged by a service, all further
	// messages within the interval are dropped until the interval is over. A message about the number of dropped messages is generated.
	// The time specification for LogRateLimitIntervalSec= may be specified in the following units: "s", "min", "h", "ms",
	// "us". See systemd.time for details. The default settings are set by RateLimitIntervalSec= and RateLimitBurst= configured
	// in journald.conf. Note that this only applies to log messages that are processed by the logging subsystem, i.e. by systemd-journald.service.
	// This means that if you connect a service's stderr directly to a file via StandardOutput=file:… or a similar setting, the
	// rate limiting will not be applied to messages written that way (but it will be enforced for messages generated via syslog
	// and similar functions).
	//
	// Added in version 240.
	LogRateLimitBurst systemdconf.Value

	// Define an extended regular expression to filter log messages based on the MESSAGE= field of the structured message. If
	// the first character of the pattern is "~", log entries matching the pattern should be discarded. This option takes a single
	// pattern as an argument but can be used multiple times to create a list of allowed and denied patterns. If the empty string
	// is assigned, the filter is reset, and all prior assignments will have no effect.
	//
	// Because the "~" character is used to define denied patterns, it must be replaced with "\x7e" to allow a message starting
	// with "~". For example, "~foobar" would add a pattern matching "foobar" to the deny list, while "\x7efoobar" would add a
	// pattern matching "~foobar" to the allow list.
	//
	// Log messages are tested against denied patterns (if any), then against allowed patterns (if any). If a log message matches
	// any of the denied patterns, it is discarded immediately without considering allowed patterns. Remaining log messages
	// are tested against allowed patterns. Messages matching against none of the allowed pattern are discarded. If no allowed
	// patterns are defined, then all messages are processed directly after going through denied filters.
	//
	// Filtering is based on the unit for which LogFilterPatterns= is defined, meaning log messages coming from systemd about
	// the unit are not taken into account. Filtered log messages will not be forwarded to traditional syslog daemons, the kernel
	// log buffer (kmsg), the systemd console, or sent as wall messages to all logged-in users.
	//
	// Note that this functionality is currently only available in system services, not in per-user services.
	//
	// Added in version 253.
	LogFilterPatterns systemdconf.Value

	// Run the unit's processes in the specified journal namespace. Expects a short user-defined string identifying the namespace.
	// If not used the processes of the service are run in the default journal namespace, i.e. their log stream is collected and
	// processed by systemd-journald.service. If this option is used any log data generated by processes of this unit (regardless
	// if via the syslog(), journal native logging or stdout/stderr logging) is collected and processed by an instance of the
	// systemd-journald@.service template unit, which manages the specified namespace. The log data is stored in a data store
	// independent from the default log namespace's data store. See systemd-journald.service for details about journal namespaces.
	//
	// Internally, journal namespaces are implemented through Linux mount namespacing and over-mounting the directory that
	// contains the relevant AF_UNIX sockets used for logging in the unit's mount namespace. Since mount namespaces are used
	// this setting disconnects propagation of mounts from the unit's processes to the host, similarly to how ReadOnlyPaths=
	// and similar settings describe above work. Journal namespaces may hence not be used for services that need to establish
	// mount points on the host.
	//
	// When this option is used the unit will automatically gain ordering and requirement dependencies on the two socket units
	// associated with the systemd-journald@.service instance so that they are automatically established prior to the unit
	// starting up. Note that when this option is used log output of this service does not appear in the regular journalctl output,
	// unless the --namespace= option is used.
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 245.
	LogNamespace systemdconf.Value

	// Sets the process name ("syslog tag") to prefix log lines sent to the logging system or the kernel log buffer with. If not set,
	// defaults to the process name of the executed process. This option is only useful when StandardOutput= or StandardError=
	// are set to journal or kmsg (or to the same settings in combination with +console) and only applies to log messages written
	// to stdout or stderr.
	SyslogIdentifier systemdconf.Value

	// Sets the syslog facility identifier to use when logging. One of kern, user, mail, daemon, auth, syslog, lpr, news, uucp,
	// cron, authpriv, ftp, local0, local1, local2, local3, local4, local5, local6 or local7. See syslog for details. This option
	// is only useful when StandardOutput= or StandardError= are set to journal or kmsg (or to the same settings in combination
	// with +console), and only applies to log messages written to stdout or stderr. Defaults to daemon.
	SyslogFacility systemdconf.Value

	// The default syslog log level to use when logging to the logging system or the kernel log buffer. One of emerg, alert, crit,
	// err, warning, notice, info, debug. See syslog for details. This option is only useful when StandardOutput= or StandardError=
	// are set to journal or kmsg (or to the same settings in combination with +console), and only applies to log messages written
	// to stdout or stderr. Note that individual lines output by executed processes may be prefixed with a different log level
	// which can be used to override the default log level specified here. The interpretation of these prefixes may be disabled
	// with SyslogLevelPrefix=, see below. For details, see sd-daemon. Defaults to info.
	SyslogLevel systemdconf.Value

	// Takes a boolean argument. If true and StandardOutput= or StandardError= are set to journal or kmsg (or to the same settings
	// in combination with +console), log lines written by the executed process that are prefixed with a log level will be processed
	// with this log level set but the prefix removed. If set to false, the interpretation of these prefixes is disabled and the
	// logged lines are passed on as-is. This only applies to log messages written to stdout or stderr. For details about this prefixing
	// see sd-daemon. Defaults to true.
	SyslogLevelPrefix systemdconf.Value

	// Sets the terminal device node to use if standard input, output, or error are connected to a TTY (see above). Defaults to /dev/console.
	TTYPath systemdconf.Value

	// Reset the terminal device specified with TTYPath= before and after execution. This does not erase the screen (see TTYVTDisallocate=
	// below for that). Defaults to "no".
	TTYReset systemdconf.Value

	// Disconnect all clients which have opened the terminal device specified with TTYPath= before and after execution. Defaults
	// to "no".
	TTYVHangup systemdconf.Value

	// Configure the size of the TTY specified with TTYPath=. If unset or set to the empty string, it is attempted to retrieve the
	// dimensions of the terminal screen via ANSI sequences, and if that fails the kernel defaults (typically 80x24) are used.
	//
	// Added in version 250.
	TTYColumns systemdconf.Value

	// Configure the size of the TTY specified with TTYPath=. If unset or set to the empty string, it is attempted to retrieve the
	// dimensions of the terminal screen via ANSI sequences, and if that fails the kernel defaults (typically 80x24) are used.
	//
	// Added in version 250.
	TTYRows systemdconf.Value

	// If the terminal device specified with TTYPath= is a virtual console terminal, try to deallocate the TTY before and after
	// execution. This ensures that the screen and scrollback buffer is cleared. If the terminal device is of any other type of
	// TTY an attempt is made to clear the screen via ANSI sequences. Defaults to "no".
	TTYVTDisallocate systemdconf.Value

	// Takes a four character identifier string for an utmp and wtmp entry for this service. This should only be set for services
	// such as getty implementations (such as agetty) where utmp/wtmp entries must be created and cleared before and after execution,
	// or for services that shall be executed as if they were run by a getty process (see below). If the configured string is longer
	// than four characters, it is truncated and the terminal four characters are used. This setting interprets %I style string
	// replacements. This setting is unset by default, i.e. no utmp/wtmp entries are created or cleaned up for this service.
	UtmpIdentifier systemdconf.Value

	// Takes one of "init", "login" or "user". If UtmpIdentifier= is set, controls which type of utmp/wtmp entries for this service
	// are generated. This setting has no effect unless UtmpIdentifier= is set too. If "init" is set, only an INIT_PROCESS entry
	// is generated and the invoked process must implement a getty-compatible utmp/wtmp logic. If "login" is set, first an INIT_PROCESS
	// entry, followed by a LOGIN_PROCESS entry is generated. In this case, the invoked process must implement a login-compatible
	// utmp/wtmp logic. If "user" is set, first an INIT_PROCESS entry, then a LOGIN_PROCESS entry and finally a USER_PROCESS
	// entry is generated. In this case, the invoked process may be any process that is suitable to be run as session leader. Defaults
	// to "init".
	//
	// Added in version 225.
	UtmpMode systemdconf.Value
}

// KillOptions represents systemd.kill — Process killing procedure
// (see https://www.freedesktop.org/software/systemd/man/systemd.kill.html#Options for details).
type KillOptions struct {
	// Specifies how processes of this unit shall be killed. One of control-group, mixed, process, none.
	//
	// If set to control-group, all remaining processes in the control group of this unit will be killed on unit stop (for services:
	// after the stop command is executed, as configured with ExecStop=). If set to mixed, the SIGTERM signal (see below) is sent
	// to the main process while the subsequent SIGKILL signal (see below) is sent to all remaining processes of the unit's control
	// group. If set to process, only the main process itself is killed (not recommended!). If set to none, no process is killed
	// (strongly recommended against!). In this case, only the stop command will be executed on unit stop, but no process will
	// be killed otherwise. Processes remaining alive after stop are left in their control group and the control group continues
	// to exist after stop unless empty.
	//
	// Note that it is not recommended to set KillMode= to process or even none, as this allows processes to escape the service manager's
	// lifecycle and resource management, and to remain running even while their service is considered stopped and is assumed
	// to not consume any resources.
	//
	// Processes will first be terminated via SIGTERM (unless the signal to send is changed via KillSignal= or RestartKillSignal=).
	// Optionally, this is immediately followed by a SIGHUP (if enabled with SendSIGHUP=). If processes still remain after:
	//
	// the main process of a unit has exited (applies to KillMode=: mixed)
	//
	// the delay configured via the TimeoutStopSec= has passed (applies to KillMode=: control-group, mixed, process)
	//
	// the termination request is repeated with the SIGKILL signal or the signal specified via FinalKillSignal= (unless this
	// is disabled via the SendSIGKILL= option). See kill for more information.
	//
	// Defaults to control-group.
	//
	// Added in version 187.
	KillMode systemdconf.Value

	// Specifies which signal to use when stopping a service. This controls the signal that is sent as first step of shutting down
	// a unit (see above), and is usually followed by SIGKILL (see above and below). For a list of valid signals, see signal. Defaults
	// to SIGTERM.
	//
	// Note that, right after sending the signal specified in this setting, systemd will always send SIGCONT, to ensure that even
	// suspended tasks can be terminated cleanly.
	//
	// Added in version 187.
	KillSignal systemdconf.Value

	// Specifies which signal to use when restarting a service. The same as KillSignal= described above, with the exception that
	// this setting is used in a restart job. Not set by default, and the value of KillSignal= is used.
	//
	// Added in version 244.
	RestartKillSignal systemdconf.Value

	// Specifies whether to send SIGHUP to remaining processes immediately after sending the signal configured with KillSignal=.
	// This is useful to indicate to shells and shell-like programs that their connection has been severed. Takes a boolean value.
	// Defaults to "no".
	//
	// Added in version 207.
	SendSIGHUP systemdconf.Value

	// Specifies whether to send SIGKILL (or the signal specified by FinalKillSignal=) to remaining processes after a timeout,
	// if the normal shutdown procedure left processes of the service around. When disabled, a KillMode= of control-group or
	// mixed service will not restart if processes from prior services exist within the control group. Takes a boolean value.
	// Defaults to "yes".
	//
	// Added in version 187.
	SendSIGKILL systemdconf.Value

	// Specifies which signal to send to remaining processes after a timeout if SendSIGKILL= is enabled. The signal configured
	// here should be one that is not typically caught and processed by services (SIGTERM is not suitable). Developers can find
	// it useful to use this to generate a coredump to troubleshoot why a service did not terminate upon receiving the initial SIGTERM
	// signal. This can be achieved by configuring LimitCORE= and setting FinalKillSignal= to either SIGQUIT or SIGABRT. Defaults
	// to SIGKILL.
	//
	// Added in version 240.
	FinalKillSignal systemdconf.Value

	// Specifies which signal to use to terminate the service when the watchdog timeout expires (enabled through WatchdogSec=).
	// Defaults to SIGABRT.
	//
	// Added in version 240.
	WatchdogSignal systemdconf.Value
}

// ResourceControlOptions represents systemd.resource-control — Resource control unit settings
// (see https://www.freedesktop.org/software/systemd/man/systemd.resource-control.html for details).
type ResourceControlOptions struct {
	// Turn on CPU usage accounting for this unit. Takes a boolean argument. Note that turning on CPU accounting for one unit will
	// also implicitly turn it on for all units contained in the same slice and for all its parent slices and the units contained
	// therein. The system default for this setting may be controlled with DefaultCPUAccounting= in systemd-system.conf.
	//
	// Under the unified cgroup hierarchy, CPU accounting is available for all units and this setting has no effect.
	//
	// Added in version 208.
	CPUAccounting systemdconf.Value

	// These settings control the cpu controller in the unified hierarchy.
	//
	// These options accept an integer value or the special string "idle":
	//
	// If set to an integer value, assign the specified CPU time weight to the processes executed, if the unified control group
	// hierarchy is used on the system. These options control the "cpu.weight" control group attribute. The allowed range is
	// 1 to 10000. Defaults to unset, but the kernel default is 100. For details about this control group attribute, see Control
	// Groups v2 and CFS Scheduler. The available CPU time is split up among all units within one slice relative to their CPU time
	// weight. A higher weight means more CPU time, a lower weight means less.
	//
	// If set to the special string "idle", mark the cgroup for "idle scheduling", which means that it will get CPU resources only
	// when there are no processes not marked in this way to execute in this cgroup or its siblings. This setting corresponds to
	// the "cpu.idle" cgroup attribute.
	//
	// Note that this value only has an effect on cgroup-v2, for cgroup-v1 it is equivalent to the minimum weight.
	//
	// While StartupCPUWeight= applies to the startup and shutdown phases of the system, CPUWeight= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupCPUWeight= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// In addition to the resource allocation performed by the cpu controller, the kernel may automatically divide resources
	// based on session-id grouping, see "The autogroup feature" in sched. The effect of this feature is similar to the cpu controller
	// with no explicit configuration, so users should be careful to not mistake one for the other.
	//
	// Added in version 232.
	CPUWeight systemdconf.Value

	// These settings control the cpu controller in the unified hierarchy.
	//
	// These options accept an integer value or the special string "idle":
	//
	// If set to an integer value, assign the specified CPU time weight to the processes executed, if the unified control group
	// hierarchy is used on the system. These options control the "cpu.weight" control group attribute. The allowed range is
	// 1 to 10000. Defaults to unset, but the kernel default is 100. For details about this control group attribute, see Control
	// Groups v2 and CFS Scheduler. The available CPU time is split up among all units within one slice relative to their CPU time
	// weight. A higher weight means more CPU time, a lower weight means less.
	//
	// If set to the special string "idle", mark the cgroup for "idle scheduling", which means that it will get CPU resources only
	// when there are no processes not marked in this way to execute in this cgroup or its siblings. This setting corresponds to
	// the "cpu.idle" cgroup attribute.
	//
	// Note that this value only has an effect on cgroup-v2, for cgroup-v1 it is equivalent to the minimum weight.
	//
	// While StartupCPUWeight= applies to the startup and shutdown phases of the system, CPUWeight= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupCPUWeight= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// In addition to the resource allocation performed by the cpu controller, the kernel may automatically divide resources
	// based on session-id grouping, see "The autogroup feature" in sched. The effect of this feature is similar to the cpu controller
	// with no explicit configuration, so users should be careful to not mistake one for the other.
	//
	// Added in version 232.
	StartupCPUWeight systemdconf.Value

	// This setting controls the cpu controller in the unified hierarchy.
	//
	// Assign the specified CPU time quota to the processes executed. Takes a percentage value, suffixed with "%". The percentage
	// specifies how much CPU time the unit shall get at maximum, relative to the total CPU time available on one CPU. Use values
	// > 100% for allotting CPU time on more than one CPU. This controls the "cpu.max" attribute on the unified control group hierarchy
	// and "cpu.cfs_quota_us" on legacy. For details about these control group attributes, see Control Groups v2 and CFS Bandwidth
	// Control. Setting CPUQuota= to an empty value unsets the quota.
	//
	// Example: CPUQuota=20% ensures that the executed processes will never get more than 20% CPU time on one CPU.
	//
	// Added in version 213.
	CPUQuota systemdconf.Value

	// This setting controls the cpu controller in the unified hierarchy.
	//
	// Assign the duration over which the CPU time quota specified by CPUQuota= is measured. Takes a time duration value in seconds,
	// with an optional suffix such as "ms" for milliseconds (or "s" for seconds.) The default setting is 100ms. The period is clamped
	// to the range supported by the kernel, which is [1ms, 1000ms]. Additionally, the period is adjusted up so that the quota interval
	// is also at least 1ms. Setting CPUQuotaPeriodSec= to an empty value resets it to the default.
	//
	// This controls the second field of "cpu.max" attribute on the unified control group hierarchy and "cpu.cfs_period_us"
	// on legacy. For details about these control group attributes, see Control Groups v2 and CFS Scheduler.
	//
	// Example: CPUQuotaPeriodSec=10ms to request that the CPU quota is measured in periods of 10ms.
	//
	// Added in version 242.
	CPUQuotaPeriodSec systemdconf.Value

	// This setting controls the cpuset controller in the unified hierarchy.
	//
	// Restrict processes to be executed on specific CPUs. Takes a list of CPU indices or ranges separated by either whitespace
	// or commas. CPU ranges are specified by the lower and upper CPU indices separated by a dash.
	//
	// Setting AllowedCPUs= or StartupAllowedCPUs= does not guarantee that all of the CPUs will be used by the processes as it
	// may be limited by parent units. The effective configuration is reported as EffectiveCPUs=.
	//
	// While StartupAllowedCPUs= applies to the startup and shutdown phases of the system, AllowedCPUs= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupAllowedCPUs= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// This setting is supported only with the unified control group hierarchy.
	//
	// Added in version 244.
	AllowedCPUs systemdconf.Value

	// This setting controls the cpuset controller in the unified hierarchy.
	//
	// Restrict processes to be executed on specific CPUs. Takes a list of CPU indices or ranges separated by either whitespace
	// or commas. CPU ranges are specified by the lower and upper CPU indices separated by a dash.
	//
	// Setting AllowedCPUs= or StartupAllowedCPUs= does not guarantee that all of the CPUs will be used by the processes as it
	// may be limited by parent units. The effective configuration is reported as EffectiveCPUs=.
	//
	// While StartupAllowedCPUs= applies to the startup and shutdown phases of the system, AllowedCPUs= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupAllowedCPUs= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// This setting is supported only with the unified control group hierarchy.
	//
	// Added in version 244.
	StartupAllowedCPUs systemdconf.Value

	// This setting controls the memory controller in the unified hierarchy.
	//
	// Turn on process and kernel memory accounting for this unit. Takes a boolean argument. Note that turning on memory accounting
	// for one unit will also implicitly turn it on for all units contained in the same slice and for all its parent slices and the
	// units contained therein. The system default for this setting may be controlled with DefaultMemoryAccounting= in systemd-system.conf.
	//
	// Added in version 208.
	MemoryAccounting systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the memory usage protection of the executed processes in this unit. When reclaiming memory, the unit is treated
	// as if it was using less memory resulting in memory to be preferentially reclaimed from unprotected units. Using MemoryLow=
	// results in a weaker protection where memory may still be reclaimed to avoid invoking the OOM killer in case there is no other
	// reclaimable memory.
	//
	// For a protection to be effective, it is generally required to set a corresponding allocation on all ancestors, which is
	// then distributed between children (with the exception of the root slice). Any MemoryMin= or MemoryLow= allocation that
	// is not explicitly distributed to specific children is used to create a shared protection for all children. As this is a shared
	// protection, the children will freely compete for the memory.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", all available
	// memory is protected, which may be useful in order to always inherit all of the protection afforded by ancestors. This controls
	// the "memory.min" or "memory.low" control group attribute. For details about this control group attribute, see Memory
	// Interface Files.
	//
	// Units may have their children use a default "memory.min" or "memory.low" value by specifying DefaultMemoryMin= or DefaultMemoryLow=,
	// which has the same semantics as MemoryMin= and MemoryLow=, or DefaultStartupMemoryLow= which has the same semantics
	// as StartupMemoryLow=. This setting does not affect "memory.min" or "memory.low" in the unit itself. Using it to set a default
	// child allocation is only useful on kernels older than 5.7, which do not support the "memory_recursiveprot" cgroup2 mount
	// option.
	//
	// While StartupMemoryLow= applies to the startup and shutdown phases of the system, MemoryMin= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryLow= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 240.
	MemoryMin systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the memory usage protection of the executed processes in this unit. When reclaiming memory, the unit is treated
	// as if it was using less memory resulting in memory to be preferentially reclaimed from unprotected units. Using MemoryLow=
	// results in a weaker protection where memory may still be reclaimed to avoid invoking the OOM killer in case there is no other
	// reclaimable memory.
	//
	// For a protection to be effective, it is generally required to set a corresponding allocation on all ancestors, which is
	// then distributed between children (with the exception of the root slice). Any MemoryMin= or MemoryLow= allocation that
	// is not explicitly distributed to specific children is used to create a shared protection for all children. As this is a shared
	// protection, the children will freely compete for the memory.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", all available
	// memory is protected, which may be useful in order to always inherit all of the protection afforded by ancestors. This controls
	// the "memory.min" or "memory.low" control group attribute. For details about this control group attribute, see Memory
	// Interface Files.
	//
	// Units may have their children use a default "memory.min" or "memory.low" value by specifying DefaultMemoryMin= or DefaultMemoryLow=,
	// which has the same semantics as MemoryMin= and MemoryLow=, or DefaultStartupMemoryLow= which has the same semantics
	// as StartupMemoryLow=. This setting does not affect "memory.min" or "memory.low" in the unit itself. Using it to set a default
	// child allocation is only useful on kernels older than 5.7, which do not support the "memory_recursiveprot" cgroup2 mount
	// option.
	//
	// While StartupMemoryLow= applies to the startup and shutdown phases of the system, MemoryMin= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryLow= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 240.
	MemoryLow systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the memory usage protection of the executed processes in this unit. When reclaiming memory, the unit is treated
	// as if it was using less memory resulting in memory to be preferentially reclaimed from unprotected units. Using MemoryLow=
	// results in a weaker protection where memory may still be reclaimed to avoid invoking the OOM killer in case there is no other
	// reclaimable memory.
	//
	// For a protection to be effective, it is generally required to set a corresponding allocation on all ancestors, which is
	// then distributed between children (with the exception of the root slice). Any MemoryMin= or MemoryLow= allocation that
	// is not explicitly distributed to specific children is used to create a shared protection for all children. As this is a shared
	// protection, the children will freely compete for the memory.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", all available
	// memory is protected, which may be useful in order to always inherit all of the protection afforded by ancestors. This controls
	// the "memory.min" or "memory.low" control group attribute. For details about this control group attribute, see Memory
	// Interface Files.
	//
	// Units may have their children use a default "memory.min" or "memory.low" value by specifying DefaultMemoryMin= or DefaultMemoryLow=,
	// which has the same semantics as MemoryMin= and MemoryLow=, or DefaultStartupMemoryLow= which has the same semantics
	// as StartupMemoryLow=. This setting does not affect "memory.min" or "memory.low" in the unit itself. Using it to set a default
	// child allocation is only useful on kernels older than 5.7, which do not support the "memory_recursiveprot" cgroup2 mount
	// option.
	//
	// While StartupMemoryLow= applies to the startup and shutdown phases of the system, MemoryMin= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryLow= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 240.
	StartupMemoryLow systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the memory usage protection of the executed processes in this unit. When reclaiming memory, the unit is treated
	// as if it was using less memory resulting in memory to be preferentially reclaimed from unprotected units. Using MemoryLow=
	// results in a weaker protection where memory may still be reclaimed to avoid invoking the OOM killer in case there is no other
	// reclaimable memory.
	//
	// For a protection to be effective, it is generally required to set a corresponding allocation on all ancestors, which is
	// then distributed between children (with the exception of the root slice). Any MemoryMin= or MemoryLow= allocation that
	// is not explicitly distributed to specific children is used to create a shared protection for all children. As this is a shared
	// protection, the children will freely compete for the memory.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", all available
	// memory is protected, which may be useful in order to always inherit all of the protection afforded by ancestors. This controls
	// the "memory.min" or "memory.low" control group attribute. For details about this control group attribute, see Memory
	// Interface Files.
	//
	// Units may have their children use a default "memory.min" or "memory.low" value by specifying DefaultMemoryMin= or DefaultMemoryLow=,
	// which has the same semantics as MemoryMin= and MemoryLow=, or DefaultStartupMemoryLow= which has the same semantics
	// as StartupMemoryLow=. This setting does not affect "memory.min" or "memory.low" in the unit itself. Using it to set a default
	// child allocation is only useful on kernels older than 5.7, which do not support the "memory_recursiveprot" cgroup2 mount
	// option.
	//
	// While StartupMemoryLow= applies to the startup and shutdown phases of the system, MemoryMin= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryLow= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 240.
	DefaultStartupMemoryLow systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the throttling limit on memory usage of the executed processes in this unit. Memory usage may go above the limit
	// if unavoidable, but the processes are heavily slowed down and memory is taken away aggressively in such cases. This is the
	// main mechanism to control memory usage of a unit.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", no memory throttling
	// is applied. This controls the "memory.high" control group attribute. For details about this control group attribute,
	// see Memory Interface Files. The effective configuration is reported as EffectiveMemoryHigh= (see also EffectiveMemoryMax=).
	//
	// While StartupMemoryHigh= applies to the startup and shutdown phases of the system, MemoryHigh= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryHigh= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 231.
	MemoryHigh systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the throttling limit on memory usage of the executed processes in this unit. Memory usage may go above the limit
	// if unavoidable, but the processes are heavily slowed down and memory is taken away aggressively in such cases. This is the
	// main mechanism to control memory usage of a unit.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", no memory throttling
	// is applied. This controls the "memory.high" control group attribute. For details about this control group attribute,
	// see Memory Interface Files. The effective configuration is reported as EffectiveMemoryHigh= (see also EffectiveMemoryMax=).
	//
	// While StartupMemoryHigh= applies to the startup and shutdown phases of the system, MemoryHigh= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryHigh= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 231.
	StartupMemoryHigh systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on memory usage of the executed processes in this unit. If memory usage cannot be contained under
	// the limit, out-of-memory killer is invoked inside the unit. It is recommended to use MemoryHigh= as the main control mechanism
	// and use MemoryMax= as the last line of defense.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", no memory limit
	// is applied. This controls the "memory.max" control group attribute. For details about this control group attribute,
	// see Memory Interface Files. The effective configuration is reported as EffectiveMemoryMax= (the value is the most stringent
	// limit of the unit and parent slices and it is capped by physical memory).
	//
	// While StartupMemoryMax= applies to the startup and shutdown phases of the system, MemoryMax= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryMax= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 231.
	MemoryMax systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on memory usage of the executed processes in this unit. If memory usage cannot be contained under
	// the limit, out-of-memory killer is invoked inside the unit. It is recommended to use MemoryHigh= as the main control mechanism
	// and use MemoryMax= as the last line of defense.
	//
	// Takes a memory size in bytes. If the value is suffixed with K, M, G or T, the specified memory size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the installed physical memory on the system. If assigned the special value "infinity", no memory limit
	// is applied. This controls the "memory.max" control group attribute. For details about this control group attribute,
	// see Memory Interface Files. The effective configuration is reported as EffectiveMemoryMax= (the value is the most stringent
	// limit of the unit and parent slices and it is capped by physical memory).
	//
	// While StartupMemoryMax= applies to the startup and shutdown phases of the system, MemoryMax= applies to normal runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryMax= allows prioritizing
	// specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 231.
	StartupMemoryMax systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on swap usage of the executed processes in this unit.
	//
	// Takes a swap size in bytes. If the value is suffixed with K, M, G or T, the specified swap size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the specified swap size on the system. If assigned the special value "infinity", no swap limit is applied.
	// These settings control the "memory.swap.max" control group attribute. For details about this control group attribute,
	// see Memory Interface Files.
	//
	// While StartupMemorySwapMax= applies to the startup and shutdown phases of the system, MemorySwapMax= applies to normal
	// runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemorySwapMax=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 232.
	MemorySwapMax systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on swap usage of the executed processes in this unit.
	//
	// Takes a swap size in bytes. If the value is suffixed with K, M, G or T, the specified swap size is parsed as Kilobytes, Megabytes,
	// Gigabytes, or Terabytes (with the base 1024), respectively. Alternatively, a percentage value may be specified, which
	// is taken relative to the specified swap size on the system. If assigned the special value "infinity", no swap limit is applied.
	// These settings control the "memory.swap.max" control group attribute. For details about this control group attribute,
	// see Memory Interface Files.
	//
	// While StartupMemorySwapMax= applies to the startup and shutdown phases of the system, MemorySwapMax= applies to normal
	// runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemorySwapMax=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 232.
	StartupMemorySwapMax systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on zswap usage of the processes in this unit. Zswap is a lightweight compressed cache for swap
	// pages. It takes pages that are in the process of being swapped out and attempts to compress them into a dynamically allocated
	// RAM-based memory pool. If the limit specified is hit, no entries from this unit will be stored in the pool until existing
	// entries are faulted back or written out to disk. See the kernel's Zswap documentation for more details.
	//
	// Takes a size in bytes. If the value is suffixed with K, M, G or T, the specified size is parsed as Kilobytes, Megabytes, Gigabytes,
	// or Terabytes (with the base 1024), respectively. If assigned the special value "infinity", no limit is applied. These
	// settings control the "memory.zswap.max" control group attribute. For details about this control group attribute, see
	// Memory Interface Files.
	//
	// While StartupMemoryZSwapMax= applies to the startup and shutdown phases of the system, MemoryZSwapMax= applies to normal
	// runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryZSwapMax=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 253.
	MemoryZSwapMax systemdconf.Value

	// These settings control the memory controller in the unified hierarchy.
	//
	// Specify the absolute limit on zswap usage of the processes in this unit. Zswap is a lightweight compressed cache for swap
	// pages. It takes pages that are in the process of being swapped out and attempts to compress them into a dynamically allocated
	// RAM-based memory pool. If the limit specified is hit, no entries from this unit will be stored in the pool until existing
	// entries are faulted back or written out to disk. See the kernel's Zswap documentation for more details.
	//
	// Takes a size in bytes. If the value is suffixed with K, M, G or T, the specified size is parsed as Kilobytes, Megabytes, Gigabytes,
	// or Terabytes (with the base 1024), respectively. If assigned the special value "infinity", no limit is applied. These
	// settings control the "memory.zswap.max" control group attribute. For details about this control group attribute, see
	// Memory Interface Files.
	//
	// While StartupMemoryZSwapMax= applies to the startup and shutdown phases of the system, MemoryZSwapMax= applies to normal
	// runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupMemoryZSwapMax=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// Added in version 253.
	StartupMemoryZSwapMax systemdconf.Value

	// This setting controls the memory controller in the unified hierarchy.
	//
	// Takes a boolean argument. When true, pages stored in the Zswap cache are permitted to be written to the backing storage,
	// false otherwise. Defaults to true. This allows disabling writeback of swap pages for IO-intensive applications, while
	// retaining the ability to store compressed pages in Zswap. See the kernel's Zswap documentation for more details.
	//
	// Added in version 256.
	MemoryZSwapWriteback systemdconf.Value

	// These settings control the cpuset controller in the unified hierarchy.
	//
	// Restrict processes to be executed on specific memory NUMA nodes. Takes a list of memory NUMA nodes indices or ranges separated
	// by either whitespace or commas. Memory NUMA nodes ranges are specified by the lower and upper NUMA nodes indices separated
	// by a dash.
	//
	// Setting AllowedMemoryNodes= or StartupAllowedMemoryNodes= does not guarantee that all of the memory NUMA nodes will
	// be used by the processes as it may be limited by parent units. The effective configuration is reported as EffectiveMemoryNodes=.
	//
	// While StartupAllowedMemoryNodes= applies to the startup and shutdown phases of the system, AllowedMemoryNodes= applies
	// to normal runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupAllowedMemoryNodes=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// This setting is supported only with the unified control group hierarchy.
	//
	// Added in version 244.
	AllowedMemoryNodes systemdconf.Value

	// These settings control the cpuset controller in the unified hierarchy.
	//
	// Restrict processes to be executed on specific memory NUMA nodes. Takes a list of memory NUMA nodes indices or ranges separated
	// by either whitespace or commas. Memory NUMA nodes ranges are specified by the lower and upper NUMA nodes indices separated
	// by a dash.
	//
	// Setting AllowedMemoryNodes= or StartupAllowedMemoryNodes= does not guarantee that all of the memory NUMA nodes will
	// be used by the processes as it may be limited by parent units. The effective configuration is reported as EffectiveMemoryNodes=.
	//
	// While StartupAllowedMemoryNodes= applies to the startup and shutdown phases of the system, AllowedMemoryNodes= applies
	// to normal runtime of the system, and if the former is not set also to the startup and shutdown phases. Using StartupAllowedMemoryNodes=
	// allows prioritizing specific services at boot-up and shutdown differently than during normal runtime.
	//
	// This setting is supported only with the unified control group hierarchy.
	//
	// Added in version 244.
	StartupAllowedMemoryNodes systemdconf.Value

	// This setting controls the pids controller in the unified hierarchy.
	//
	// Turn on task accounting for this unit. Takes a boolean argument. If enabled, the kernel will keep track of the total number
	// of tasks in the unit and its children. This number includes both kernel threads and userspace processes, with each thread
	// counted individually. Note that turning on tasks accounting for one unit will also implicitly turn it on for all units contained
	// in the same slice and for all its parent slices and the units contained therein. The system default for this setting may be
	// controlled with DefaultTasksAccounting= in systemd-system.conf.
	//
	// Added in version 227.
	TasksAccounting systemdconf.Value

	// This setting controls the pids controller in the unified hierarchy.
	//
	// Specify the maximum number of tasks that may be created in the unit. This ensures that the number of tasks accounted for the
	// unit (see above) stays below a specific limit. This either takes an absolute number of tasks or a percentage value that is
	// taken relative to the configured maximum number of tasks on the system. If assigned the special value "infinity", no tasks
	// limit is applied. This controls the "pids.max" control group attribute. For details about this control group attribute,
	// the pids controller . The effective configuration is reported as EffectiveTasksMax=.
	//
	// The system default for this setting may be controlled with DefaultTasksMax= in systemd-system.conf.
	//
	// Added in version 227.
	TasksMax systemdconf.Value

	// This setting controls the io controller in the unified hierarchy.
	//
	// Turn on Block I/O accounting for this unit, if the unified control group hierarchy is used on the system. Takes a boolean
	// argument. Note that turning on block I/O accounting for one unit will also implicitly turn it on for all units contained
	// in the same slice and all for its parent slices and the units contained therein. The system default for this setting may be
	// controlled with DefaultIOAccounting= in systemd-system.conf.
	//
	// Added in version 230.
	IOAccounting systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the default overall block I/O weight for the executed processes, if the unified control group hierarchy is used on the
	// system. Takes a single weight value (between 1 and 10000) to set the default block I/O weight. This controls the "io.weight"
	// control group attribute, which defaults to 100. For details about this control group attribute, see IO Interface Files.
	// The available I/O bandwidth is split up among all units within one slice relative to their block I/O weight. A higher weight
	// means more I/O bandwidth, a lower weight means less.
	//
	// While StartupIOWeight= applies to the startup and shutdown phases of the system, IOWeight= applies to the later runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. This allows prioritizing specific services
	// at boot-up and shutdown differently than during runtime.
	//
	// Added in version 230.
	IOWeight systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the default overall block I/O weight for the executed processes, if the unified control group hierarchy is used on the
	// system. Takes a single weight value (between 1 and 10000) to set the default block I/O weight. This controls the "io.weight"
	// control group attribute, which defaults to 100. For details about this control group attribute, see IO Interface Files.
	// The available I/O bandwidth is split up among all units within one slice relative to their block I/O weight. A higher weight
	// means more I/O bandwidth, a lower weight means less.
	//
	// While StartupIOWeight= applies to the startup and shutdown phases of the system, IOWeight= applies to the later runtime
	// of the system, and if the former is not set also to the startup and shutdown phases. This allows prioritizing specific services
	// at boot-up and shutdown differently than during runtime.
	//
	// Added in version 230.
	StartupIOWeight systemdconf.Value

	// This setting controls the io controller in the unified hierarchy.
	//
	// Set the per-device overall block I/O weight for the executed processes, if the unified control group hierarchy is used
	// on the system. Takes a space-separated pair of a file path and a weight value to specify the device specific weight value,
	// between 1 and 10000. (Example: "/dev/sda 1000"). The file path may be specified as path to a block device node or as any other
	// file, in which case the backing block device of the file system of the file is determined. This controls the "io.weight"
	// control group attribute, which defaults to 100. Use this option multiple times to set weights for multiple devices. For
	// details about this control group attribute, see IO Interface Files.
	//
	// The specified device node should reference a block device that has an I/O scheduler associated, i.e. should not refer to
	// partition or loopback block devices, but to the originating, physical device. When a path to a regular file or directory
	// is specified it is attempted to discover the correct originating device backing the file system of the specified path.
	// This works correctly only for simpler cases, where the file system is directly placed on a partition or physical block device,
	// or where simple 1:1 encryption using dm-crypt/LUKS is used. This discovery does not cover complex storage and in particular
	// RAID and volume management storage devices.
	//
	// Added in version 230.
	IODeviceWeight systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the per-device overall block I/O bandwidth maximum limit for the executed processes, if the unified control group
	// hierarchy is used on the system. This limit is not work-conserving and the executed processes are not allowed to use more
	// even if the device has idle capacity. Takes a space-separated pair of a file path and a bandwidth value (in bytes per second)
	// to specify the device specific bandwidth. The file path may be a path to a block device node, or as any other file in which case
	// the backing block device of the file system of the file is used. If the bandwidth is suffixed with K, M, G, or T, the specified
	// bandwidth is parsed as Kilobytes, Megabytes, Gigabytes, or Terabytes, respectively, to the base of 1000. (Example: "/dev/disk/by-path/pci-0000:00:1f.2-scsi-0:0:0:0
	// 5M"). This controls the "io.max" control group attributes. Use this option multiple times to set bandwidth limits for
	// multiple devices. For details about this control group attribute, see IO Interface Files.
	//
	// Similar restrictions on block device discovery as for IODeviceWeight= apply, see above.
	//
	// Added in version 230.
	IOReadBandwidthMax systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the per-device overall block I/O bandwidth maximum limit for the executed processes, if the unified control group
	// hierarchy is used on the system. This limit is not work-conserving and the executed processes are not allowed to use more
	// even if the device has idle capacity. Takes a space-separated pair of a file path and a bandwidth value (in bytes per second)
	// to specify the device specific bandwidth. The file path may be a path to a block device node, or as any other file in which case
	// the backing block device of the file system of the file is used. If the bandwidth is suffixed with K, M, G, or T, the specified
	// bandwidth is parsed as Kilobytes, Megabytes, Gigabytes, or Terabytes, respectively, to the base of 1000. (Example: "/dev/disk/by-path/pci-0000:00:1f.2-scsi-0:0:0:0
	// 5M"). This controls the "io.max" control group attributes. Use this option multiple times to set bandwidth limits for
	// multiple devices. For details about this control group attribute, see IO Interface Files.
	//
	// Similar restrictions on block device discovery as for IODeviceWeight= apply, see above.
	//
	// Added in version 230.
	IOWriteBandwidthMax systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the per-device overall block I/O IOs-Per-Second maximum limit for the executed processes, if the unified control
	// group hierarchy is used on the system. This limit is not work-conserving and the executed processes are not allowed to use
	// more even if the device has idle capacity. Takes a space-separated pair of a file path and an IOPS value to specify the device
	// specific IOPS. The file path may be a path to a block device node, or as any other file in which case the backing block device
	// of the file system of the file is used. If the IOPS is suffixed with K, M, G, or T, the specified IOPS is parsed as KiloIOPS, MegaIOPS,
	// GigaIOPS, or TeraIOPS, respectively, to the base of 1000. (Example: "/dev/disk/by-path/pci-0000:00:1f.2-scsi-0:0:0:0
	// 1K"). This controls the "io.max" control group attributes. Use this option multiple times to set IOPS limits for multiple
	// devices. For details about this control group attribute, see IO Interface Files.
	//
	// Similar restrictions on block device discovery as for IODeviceWeight= apply, see above.
	//
	// Added in version 230.
	IOReadIOPSMax systemdconf.Value

	// These settings control the io controller in the unified hierarchy.
	//
	// Set the per-device overall block I/O IOs-Per-Second maximum limit for the executed processes, if the unified control
	// group hierarchy is used on the system. This limit is not work-conserving and the executed processes are not allowed to use
	// more even if the device has idle capacity. Takes a space-separated pair of a file path and an IOPS value to specify the device
	// specific IOPS. The file path may be a path to a block device node, or as any other file in which case the backing block device
	// of the file system of the file is used. If the IOPS is suffixed with K, M, G, or T, the specified IOPS is parsed as KiloIOPS, MegaIOPS,
	// GigaIOPS, or TeraIOPS, respectively, to the base of 1000. (Example: "/dev/disk/by-path/pci-0000:00:1f.2-scsi-0:0:0:0
	// 1K"). This controls the "io.max" control group attributes. Use this option multiple times to set IOPS limits for multiple
	// devices. For details about this control group attribute, see IO Interface Files.
	//
	// Similar restrictions on block device discovery as for IODeviceWeight= apply, see above.
	//
	// Added in version 230.
	IOWriteIOPSMax systemdconf.Value

	// This setting controls the io controller in the unified hierarchy.
	//
	// Set the per-device average target I/O latency for the executed processes, if the unified control group hierarchy is used
	// on the system. Takes a file path and a timespan separated by a space to specify the device specific latency target. (Example:
	// "/dev/sda 25ms"). The file path may be specified as path to a block device node or as any other file, in which case the backing
	// block device of the file system of the file is determined. This controls the "io.latency" control group attribute. Use
	// this option multiple times to set latency target for multiple devices. For details about this control group attribute,
	// see IO Interface Files.
	//
	// Implies "IOAccounting=yes".
	//
	// These settings are supported only if the unified control group hierarchy is used.
	//
	// Similar restrictions on block device discovery as for IODeviceWeight= apply, see above.
	//
	// Added in version 240.
	IODeviceLatencyTargetSec systemdconf.Value

	// Takes a boolean argument. If true, turns on IPv4 and IPv6 network traffic accounting for packets sent or received by the
	// unit. When this option is turned on, all IPv4 and IPv6 sockets created by any process of the unit are accounted for.
	//
	// When this option is used in socket units, it applies to all IPv4 and IPv6 sockets associated with it (including both listening
	// and connection sockets where this applies). Note that for socket-activated services, this configuration setting and
	// the accounting data of the service unit and the socket unit are kept separate, and displayed separately. No propagation
	// of the setting and the collected statistics is done, in either direction. Moreover, any traffic sent or received on any
	// of the socket unit's sockets is accounted to the socket unit — and never to the service unit it might have activated, even
	// if the socket is used by it.
	//
	// The system default for this setting may be controlled with DefaultIPAccounting= in systemd-system.conf.
	//
	// Note that this functionality is currently only available for system services, not for per-user services.
	//
	// Added in version 235.
	IPAccounting systemdconf.Value

	// Turn on network traffic filtering for IP packets sent and received over AF_INET and AF_INET6 sockets. Both directives
	// take a space separated list of IPv4 or IPv6 addresses, each optionally suffixed with an address prefix length in bits after
	// a "/" character. If the suffix is omitted, the address is considered a host address, i.e. the filter covers the whole address
	// (32 bits for IPv4, 128 bits for IPv6).
	//
	// The access lists configured with this option are applied to all sockets created by processes of this unit (or in the case
	// of socket units, associated with it). The lists are implicitly combined with any lists configured for any of the parent
	// slice units this unit might be a member of. By default, both access lists are empty. Both ingress and egress traffic is filtered
	// by these settings. In case of ingress traffic the source IP address is checked against these access lists, in case of egress
	// traffic the destination IP address is checked. The following rules are applied in turn:
	//
	// Access is granted when the checked IP address matches an entry in the IPAddressAllow= list.
	//
	// Otherwise, access is denied when the checked IP address matches an entry in the IPAddressDeny= list.
	//
	// Otherwise, access is granted.
	//
	// In order to implement an allow-listing IP firewall, it is recommended to use a IPAddressDeny=any setting on an upper-level
	// slice unit (such as the root slice -.slice or the slice containing all system services system.slice – see systemd.special
	// for details on these slice units), plus individual per-service IPAddressAllow= lines permitting network access to relevant
	// services, and only them.
	//
	// Note that for socket-activated services, the IP access list configured on the socket unit applies to all sockets associated
	// with it directly, but not to any sockets created by the ultimately activated services for it. Conversely, the IP access
	// list configured for the service is not applied to any sockets passed into the service via socket activation. Thus, it is
	// usually a good idea to replicate the IP access lists on both the socket and the service unit. Nevertheless, it may make sense
	// to maintain one list more open and the other one more restricted, depending on the use case.
	//
	// If these settings are used multiple times in the same unit the specified lists are combined. If an empty string is assigned
	// to these settings the specific access list is reset and all previous settings undone.
	//
	// In place of explicit IPv4 or IPv6 address and prefix length specifications a small set of symbolic names may be used. The
	// following names are defined:
	//
	//	+---------------+--------------------------+--------------------------------+
	//	| SYMBOLIC NAME |        DEFINITION        |            MEANING             |
	//	+---------------+--------------------------+--------------------------------+
	//	| any           | 0.0.0.0/0 ::/0           | Any host                       |
	//	| localhost     | 127.0.0.0/8 ::1/128      | All addresses on the local     |
	//	|               |                          | loopback                       |
	//	| link-local    | 169.254.0.0/16 fe80::/64 | All link-local IP addresses    |
	//	| multicast     | 224.0.0.0/4 ff00::/8     | All IP multicasting addresses  |
	//	+---------------+--------------------------+--------------------------------+
	//
	// Note that these settings might not be supported on some systems (for example if eBPF control group support is not enabled
	// in the underlying kernel or container manager). These settings will have no effect in that case. If compatibility with
	// such systems is desired it is hence recommended to not exclusively rely on them for IP security.
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 235.
	IPAddressAllow systemdconf.Value

	// Turn on network traffic filtering for IP packets sent and received over AF_INET and AF_INET6 sockets. Both directives
	// take a space separated list of IPv4 or IPv6 addresses, each optionally suffixed with an address prefix length in bits after
	// a "/" character. If the suffix is omitted, the address is considered a host address, i.e. the filter covers the whole address
	// (32 bits for IPv4, 128 bits for IPv6).
	//
	// The access lists configured with this option are applied to all sockets created by processes of this unit (or in the case
	// of socket units, associated with it). The lists are implicitly combined with any lists configured for any of the parent
	// slice units this unit might be a member of. By default, both access lists are empty. Both ingress and egress traffic is filtered
	// by these settings. In case of ingress traffic the source IP address is checked against these access lists, in case of egress
	// traffic the destination IP address is checked. The following rules are applied in turn:
	//
	// Access is granted when the checked IP address matches an entry in the IPAddressAllow= list.
	//
	// Otherwise, access is denied when the checked IP address matches an entry in the IPAddressDeny= list.
	//
	// Otherwise, access is granted.
	//
	// In order to implement an allow-listing IP firewall, it is recommended to use a IPAddressDeny=any setting on an upper-level
	// slice unit (such as the root slice -.slice or the slice containing all system services system.slice – see systemd.special
	// for details on these slice units), plus individual per-service IPAddressAllow= lines permitting network access to relevant
	// services, and only them.
	//
	// Note that for socket-activated services, the IP access list configured on the socket unit applies to all sockets associated
	// with it directly, but not to any sockets created by the ultimately activated services for it. Conversely, the IP access
	// list configured for the service is not applied to any sockets passed into the service via socket activation. Thus, it is
	// usually a good idea to replicate the IP access lists on both the socket and the service unit. Nevertheless, it may make sense
	// to maintain one list more open and the other one more restricted, depending on the use case.
	//
	// If these settings are used multiple times in the same unit the specified lists are combined. If an empty string is assigned
	// to these settings the specific access list is reset and all previous settings undone.
	//
	// In place of explicit IPv4 or IPv6 address and prefix length specifications a small set of symbolic names may be used. The
	// following names are defined:
	//
	//	+---------------+--------------------------+--------------------------------+
	//	| SYMBOLIC NAME |        DEFINITION        |            MEANING             |
	//	+---------------+--------------------------+--------------------------------+
	//	| any           | 0.0.0.0/0 ::/0           | Any host                       |
	//	| localhost     | 127.0.0.0/8 ::1/128      | All addresses on the local     |
	//	|               |                          | loopback                       |
	//	| link-local    | 169.254.0.0/16 fe80::/64 | All link-local IP addresses    |
	//	| multicast     | 224.0.0.0/4 ff00::/8     | All IP multicasting addresses  |
	//	+---------------+--------------------------+--------------------------------+
	//
	// Note that these settings might not be supported on some systems (for example if eBPF control group support is not enabled
	// in the underlying kernel or container manager). These settings will have no effect in that case. If compatibility with
	// such systems is desired it is hence recommended to not exclusively rely on them for IP security.
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 235.
	IPAddressDeny systemdconf.Value

	// Configures restrictions on the ability of unit processes to invoke bind on a socket. Both allow and deny rules to be defined
	// that restrict which addresses a socket may be bound to.
	//
	// bind-rule describes socket properties such as address-family, transport-protocol and ip-ports.
	//
	// bind-rule := { [address-family:][transport-protocol:][ip-ports] | any }
	//
	// address-family := { ipv4 | ipv6 }
	//
	// transport-protocol := { tcp | udp }
	//
	// ip-ports := { ip-port | ip-port-range }
	//
	// An optional address-family expects ipv4 or ipv6 values. If not specified, a rule will be matched for both IPv4 and IPv6 addresses
	// and applied depending on other socket fields, e.g. transport-protocol, ip-port.
	//
	// An optional transport-protocol expects tcp or udp transport protocol names. If not specified, a rule will be matched for
	// any transport protocol.
	//
	// An optional ip-port value must lie within 1…65535 interval inclusively, i.e. dynamic port 0 is not allowed. A range of sequential
	// ports is described by ip-port-range := ip-port-low-ip-port-high, where ip-port-low is smaller than or equal to ip-port-high
	// and both are within 1…65535 inclusively.
	//
	// A special value any can be used to apply a rule to any address family, transport protocol and any port with a positive value.
	//
	// To allow multiple rules assign SocketBindAllow= or SocketBindDeny= multiple times. To clear the existing assignments
	// pass an empty SocketBindAllow= or SocketBindDeny= assignment.
	//
	// For each of SocketBindAllow= and SocketBindDeny=, maximum allowed number of assignments is 128.
	//
	// Binding to a socket is allowed when a socket address matches an entry in the SocketBindAllow= list.
	//
	// Otherwise, binding is denied when the socket address matches an entry in the SocketBindDeny= list.
	//
	// Otherwise, binding is allowed.
	//
	// The feature is implemented with cgroup/bind4 and cgroup/bind6 cgroup-bpf hooks.
	//
	// Note that these settings apply to any bind system call invocation by the unit processes, regardless in which network namespace
	// they are placed. Or in other words: changing the network namespace is not a suitable mechanism for escaping these restrictions
	// on bind().
	//
	// Examples:
	//
	//	… # Allow binding IPv6 socket addresses with a port greater than or equal to 10000. [Service] SocketBindAllow=ipv6:10000-65535
	//	SocketBindDeny=any … # Allow binding IPv4 and IPv6 socket addresses with 1234 and 4321 ports. [Service] SocketBindAllow=1234
	//	SocketBindAllow=4321 SocketBindDeny=any … # Deny binding IPv6 socket addresses. [Service] SocketBindDeny=ipv6 …
	//	# Deny binding IPv4 and IPv6 socket addresses. [Service] SocketBindDeny=any … # Allow binding only over TCP [Service]
	//	SocketBindAllow=tcp SocketBindDeny=any … # Allow binding only over IPv6/TCP [Service] SocketBindAllow=ipv6:tcp
	//	SocketBindDeny=any … # Allow binding ports within 10000-65535 range over IPv4/UDP. [Service] SocketBindAllow=ipv4:udp:10000-65535
	//	SocketBindDeny=any …
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 249.
	SocketBindAllow systemdconf.Value

	// Configures restrictions on the ability of unit processes to invoke bind on a socket. Both allow and deny rules to be defined
	// that restrict which addresses a socket may be bound to.
	//
	// bind-rule describes socket properties such as address-family, transport-protocol and ip-ports.
	//
	// bind-rule := { [address-family:][transport-protocol:][ip-ports] | any }
	//
	// address-family := { ipv4 | ipv6 }
	//
	// transport-protocol := { tcp | udp }
	//
	// ip-ports := { ip-port | ip-port-range }
	//
	// An optional address-family expects ipv4 or ipv6 values. If not specified, a rule will be matched for both IPv4 and IPv6 addresses
	// and applied depending on other socket fields, e.g. transport-protocol, ip-port.
	//
	// An optional transport-protocol expects tcp or udp transport protocol names. If not specified, a rule will be matched for
	// any transport protocol.
	//
	// An optional ip-port value must lie within 1…65535 interval inclusively, i.e. dynamic port 0 is not allowed. A range of sequential
	// ports is described by ip-port-range := ip-port-low-ip-port-high, where ip-port-low is smaller than or equal to ip-port-high
	// and both are within 1…65535 inclusively.
	//
	// A special value any can be used to apply a rule to any address family, transport protocol and any port with a positive value.
	//
	// To allow multiple rules assign SocketBindAllow= or SocketBindDeny= multiple times. To clear the existing assignments
	// pass an empty SocketBindAllow= or SocketBindDeny= assignment.
	//
	// For each of SocketBindAllow= and SocketBindDeny=, maximum allowed number of assignments is 128.
	//
	// Binding to a socket is allowed when a socket address matches an entry in the SocketBindAllow= list.
	//
	// Otherwise, binding is denied when the socket address matches an entry in the SocketBindDeny= list.
	//
	// Otherwise, binding is allowed.
	//
	// The feature is implemented with cgroup/bind4 and cgroup/bind6 cgroup-bpf hooks.
	//
	// Note that these settings apply to any bind system call invocation by the unit processes, regardless in which network namespace
	// they are placed. Or in other words: changing the network namespace is not a suitable mechanism for escaping these restrictions
	// on bind().
	//
	// Examples:
	//
	//	… # Allow binding IPv6 socket addresses with a port greater than or equal to 10000. [Service] SocketBindAllow=ipv6:10000-65535
	//	SocketBindDeny=any … # Allow binding IPv4 and IPv6 socket addresses with 1234 and 4321 ports. [Service] SocketBindAllow=1234
	//	SocketBindAllow=4321 SocketBindDeny=any … # Deny binding IPv6 socket addresses. [Service] SocketBindDeny=ipv6 …
	//	# Deny binding IPv4 and IPv6 socket addresses. [Service] SocketBindDeny=any … # Allow binding only over TCP [Service]
	//	SocketBindAllow=tcp SocketBindDeny=any … # Allow binding only over IPv6/TCP [Service] SocketBindAllow=ipv6:tcp
	//	SocketBindDeny=any … # Allow binding ports within 10000-65535 range over IPv4/UDP. [Service] SocketBindAllow=ipv4:udp:10000-65535
	//	SocketBindDeny=any …
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 249.
	SocketBindDeny systemdconf.Value

	// Takes a list of space-separated network interface names. This option restricts the network interfaces that processes
	// of this unit can use. By default, processes can only use the network interfaces listed (allow-list). If the first character
	// of the rule is "~", the effect is inverted: the processes can only use network interfaces not listed (deny-list).
	//
	// This option can appear multiple times, in which case the network interface names are merged. If the empty string is assigned
	// the set is reset, all prior assignments will have not effect.
	//
	// If you specify both types of this option (i.e. allow-listing and deny-listing), the first encountered will take precedence
	// and will dictate the default action (allow vs deny). Then the next occurrences of this option will add or delete the listed
	// network interface names from the set, depending of its type and the default action.
	//
	// The loopback interface ("lo") is not treated in any special way, you have to configure it explicitly in the unit file.
	//
	// Example 1: allow-list
	//
	//	RestrictNetworkInterfaces=eth1 RestrictNetworkInterfaces=eth2
	//
	// Programs in the unit will be only able to use the eth1 and eth2 network interfaces.
	//
	// Example 2: deny-list
	//
	//	RestrictNetworkInterfaces=~eth1 eth2
	//
	// Programs in the unit will be able to use any network interface but eth1 and eth2.
	//
	// Example 3: mixed
	//
	//	RestrictNetworkInterfaces=eth1 eth2 RestrictNetworkInterfaces=~eth1
	//
	// Programs in the unit will be only able to use the eth2 network interface.
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 250.
	RestrictNetworkInterfaces systemdconf.Value

	// This setting provides a method for integrating dynamic cgroup, user and group IDs into firewall rules with NFT sets. The
	// benefit of using this setting is to be able to use the IDs as selectors in firewall rules easily and this in turn allows more
	// fine grained filtering. NFT rules for cgroup matching use numeric cgroup IDs, which change every time a service is restarted,
	// making them hard to use in systemd environment otherwise. Dynamic and random IDs used by DynamicUser= can be also integrated
	// with this setting.
	//
	// This option expects a whitespace separated list of NFT set definitions. Each definition consists of a colon-separated
	// tuple of source type (one of "cgroup", "user" or "group"), NFT address family (one of "arp", "bridge", "inet", "ip", "ip6",
	// or "netdev"), table name and set name. The names of tables and sets must conform to lexical restrictions of NFT table names.
	// The type of the element used in the NFT filter must match the type implied by the directive ("cgroup", "user" or "group")
	// as shown in the table below. When a control group or a unit is realized, the corresponding ID will be appended to the NFT sets
	// and it will be be removed when the control group or unit is removed. systemd only inserts elements to (or removes from) the
	// sets, so the related NFT rules, tables and sets must be prepared elsewhere in advance. Failures to manage the sets will be
	// ignored.
	//
	//	+-------------+------------------+-----------------------------+
	//	| SOURCE TYPE |   DESCRIPTION    | CORRESPONDING NFT TYPE NAME |
	//	+-------------+------------------+-----------------------------+
	//	| "cgroup"    | control group ID | "cgroupsv2"                 |
	//	| "user"      | user ID          | "meta skuid"                |
	//	| "group"     | group ID         | "meta skgid"                |
	//	+-------------+------------------+-----------------------------+
	//
	// If the firewall rules are reinstalled so that the contents of NFT sets are destroyed, command systemctl daemon-reload
	// can be used to refill the sets.
	//
	// Example:
	//
	//	[Unit] NFTSet=cgroup:inet:filter:my_service user:inet:filter:serviceuser
	//
	// Corresponding NFT rules:
	//
	//	table inet filter { set my_service { type cgroupsv2 } set serviceuser { typeof meta skuid } chain x { socket cgroupv2 level
	//	2 @my_service accept drop } chain y { meta skuid @serviceuser accept drop } }
	//
	// This option is only available for system services and is not supported for services running in per-user instances of the
	// service manager.
	//
	// Added in version 255.
	NFTSet systemdconf.Value

	// Add custom network traffic filters implemented as BPF programs, applying to all IP packets sent and received over AF_INET
	// and AF_INET6 sockets. Takes an absolute path to a pinned BPF program in the BPF virtual filesystem (/sys/fs/bpf/).
	//
	// The filters configured with this option are applied to all sockets created by processes of this unit (or in the case of socket
	// units, associated with it). The filters are loaded in addition to filters any of the parent slice units this unit might be
	// a member of as well as any IPAddressAllow= and IPAddressDeny= filters in any of these units. By default, there are no filters
	// specified.
	//
	// If these settings are used multiple times in the same unit all the specified programs are attached. If an empty string is
	// assigned to these settings the program list is reset and all previous specified programs ignored.
	//
	// If the path BPF_FS_PROGRAM_PATH in IPIngressFilterPath= assignment is already being handled by BPFProgram= ingress
	// hook, e.g. BPFProgram=ingress:BPF_FS_PROGRAM_PATH, the assignment will be still considered valid and the program
	// will be attached to a cgroup. Same for IPEgressFilterPath= path and egress hook.
	//
	// Note that for socket-activated services, the IP filter programs configured on the socket unit apply to all sockets associated
	// with it directly, but not to any sockets created by the ultimately activated services for it. Conversely, the IP filter
	// programs configured for the service are not applied to any sockets passed into the service via socket activation. Thus,
	// it is usually a good idea, to replicate the IP filter programs on both the socket and the service unit, however it often makes
	// sense to maintain one configuration more open and the other one more restricted, depending on the use case.
	//
	// Note that these settings might not be supported on some systems (for example if eBPF control group support is not enabled
	// in the underlying kernel or container manager). These settings will fail the service in that case. If compatibility with
	// such systems is desired it is hence recommended to attach your filter manually (requires Delegate=yes) instead of using
	// this setting.
	//
	// Added in version 243.
	IPIngressFilterPath systemdconf.Value

	// Add custom network traffic filters implemented as BPF programs, applying to all IP packets sent and received over AF_INET
	// and AF_INET6 sockets. Takes an absolute path to a pinned BPF program in the BPF virtual filesystem (/sys/fs/bpf/).
	//
	// The filters configured with this option are applied to all sockets created by processes of this unit (or in the case of socket
	// units, associated with it). The filters are loaded in addition to filters any of the parent slice units this unit might be
	// a member of as well as any IPAddressAllow= and IPAddressDeny= filters in any of these units. By default, there are no filters
	// specified.
	//
	// If these settings are used multiple times in the same unit all the specified programs are attached. If an empty string is
	// assigned to these settings the program list is reset and all previous specified programs ignored.
	//
	// If the path BPF_FS_PROGRAM_PATH in IPIngressFilterPath= assignment is already being handled by BPFProgram= ingress
	// hook, e.g. BPFProgram=ingress:BPF_FS_PROGRAM_PATH, the assignment will be still considered valid and the program
	// will be attached to a cgroup. Same for IPEgressFilterPath= path and egress hook.
	//
	// Note that for socket-activated services, the IP filter programs configured on the socket unit apply to all sockets associated
	// with it directly, but not to any sockets created by the ultimately activated services for it. Conversely, the IP filter
	// programs configured for the service are not applied to any sockets passed into the service via socket activation. Thus,
	// it is usually a good idea, to replicate the IP filter programs on both the socket and the service unit, however it often makes
	// sense to maintain one configuration more open and the other one more restricted, depending on the use case.
	//
	// Note that these settings might not be supported on some systems (for example if eBPF control group support is not enabled
	// in the underlying kernel or container manager). These settings will fail the service in that case. If compatibility with
	// such systems is desired it is hence recommended to attach your filter manually (requires Delegate=yes) instead of using
	// this setting.
	//
	// Added in version 243.
	IPEgressFilterPath systemdconf.Value

	// BPFProgram= allows attaching custom BPF programs to the cgroup of a unit. (This generalizes the functionality exposed
	// via IPEgressFilterPath= and IPIngressFilterPath= for other hooks.) Cgroup-bpf hooks in the form of BPF programs loaded
	// to the BPF filesystem are attached with cgroup-bpf attach flags determined by the unit. For details about attachment types
	// and flags see bpf.h. Also refer to the general BPF documentation.
	//
	// The specification of BPF program consists of a pair of BPF program type and program path in the file system, with ":" as the
	// separator: type:program-path.
	//
	// The BPF program type is equivalent to the BPF attach type used in bpftool It may be one of egress, ingress, sock_create, sock_ops,
	// device, bind4, bind6, connect4, connect6, post_bind4, post_bind6, sendmsg4, sendmsg6, sysctl, recvmsg4, recvmsg6,
	// getsockopt, or setsockopt.
	//
	// The specified program path must be an absolute path referencing a BPF program inode in the bpffs file system (which generally
	// means it must begin with /sys/fs/bpf/). If a specified program does not exist (i.e. has not been uploaded to the BPF subsystem
	// of the kernel yet), it will not be installed but unit activation will continue (a warning will be printed to the logs).
	//
	// Setting BPFProgram= to an empty value makes previous assignments ineffective.
	//
	// Multiple assignments of the same program type/path pair have the same effect as a single assignment: the program will be
	// attached just once.
	//
	// If BPF egress pinned to program-path path is already being handled by IPEgressFilterPath=, BPFProgram= assignment will
	// be considered valid and BPFProgram= will be attached to a cgroup. Similarly for ingress hook and IPIngressFilterPath=
	// assignment.
	//
	// BPF programs passed with BPFProgram= are attached to the cgroup of a unit with BPF attach flag multi, that allows further
	// attachments of the same type within cgroup hierarchy topped by the unit cgroup.
	//
	// Examples:
	//
	//	BPFProgram=egress:/sys/fs/bpf/egress-hook BPFProgram=bind6:/sys/fs/bpf/sock-addr-hook
	//
	// Added in version 249.
	BPFProgram systemdconf.Value

	// Control access to specific device nodes by the executed processes. Takes two space-separated strings: a device node specifier
	// followed by a combination of r, w, m to control reading, writing, or creation of the specific device nodes by the unit (mknod),
	// respectively. This functionality is implemented using eBPF filtering.
	//
	// When access to all physical devices should be disallowed, PrivateDevices= may be used instead. See systemd.exec.
	//
	// The device node specifier is either a path to a device node in the file system, starting with /dev/, or a string starting with
	// either "char-" or "block-" followed by a device group name, as listed in /proc/devices. The latter is useful to allow-list
	// all current and future devices belonging to a specific device group at once. The device group is matched according to filename
	// globbing rules, you may hence use the "*" and "?" wildcards. (Note that such globbing wildcards are not available for device
	// node path specifications!) In order to match device nodes by numeric major/minor, use device node paths in the /dev/char/
	// and /dev/block/ directories. However, matching devices by major/minor is generally not recommended as assignments
	// are neither stable nor portable between systems or different kernel versions.
	//
	// Examples: /dev/sda5 is a path to a device node, referring to an ATA or SCSI block device. "char-pts" and "char-alsa" are
	// specifiers for all pseudo TTYs and all ALSA sound devices, respectively. "char-cpu/*" is a specifier matching all CPU
	// related device groups.
	//
	// Note that allow lists defined this way should only reference device groups which are resolvable at the time the unit is started.
	// Any device groups not resolvable then are not added to the device allow list. In order to work around this limitation, consider
	// extending service units with a pair of After=modprobe@xyz.service and Wants=modprobe@xyz.service lines that load
	// the necessary kernel module implementing the device group if missing. Example:
	//
	//	… [Unit] Wants=modprobe@loop.service After=modprobe@loop.service [Service] DeviceAllow=block-loop DeviceAllow=/dev/loop-control
	//	…
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 208.
	DeviceAllow systemdconf.Value

	// Control the policy for allowing device access:
	//
	//	strict: means to only allow types of access that are explicitly specified.
	//
	//	Added in version 208.
	//	closed: in addition, allows access to standard pseudo devices including /dev/null, /dev/zero, /dev/full, /dev/random,
	//	and /dev/urandom.
	//
	//	Added in version 208.
	//	auto: in addition, allows access to all devices if no explicit DeviceAllow= is present. This is the default.
	//
	//	Added in version 208.
	//
	// This option cannot be bypassed by prefixing "+" to the executable path in the service unit, as it applies to the whole control
	// group.
	//
	// Added in version 208.
	DevicePolicy systemdconf.Value

	// The name of the slice unit to place the unit in. Defaults to system.slice for all non-instantiated units of all unit types
	// (except for slice units themselves see below). Instance units are by default placed in a subslice of system.slice that
	// is named after the template name.
	//
	// This option may be used to arrange systemd units in a hierarchy of slices each of which might have resource settings applied.
	//
	// For units of type slice, the only accepted value for this setting is the parent slice. Since the name of a slice unit implies
	// the parent slice, it is hence redundant to ever set this parameter directly for slice units.
	//
	// Special care should be taken when relying on the default slice assignment in templated service units that have DefaultDependencies=no
	// set, see systemd.service, section "Default Dependencies" for details.
	//
	// Added in version 208.
	Slice systemdconf.Value

	// Turns on delegation of further resource control partitioning to processes of the unit. Units where this is enabled may
	// create and manage their own private subhierarchy of control groups below the control group of the unit itself. For unprivileged
	// services (i.e. those using the User= setting) the unit's control group will be made accessible to the relevant user.
	//
	// When enabled the service manager will refrain from manipulating control groups or moving processes below the unit's control
	// group, so that a clear concept of ownership is established: the control group tree at the level of the unit's control group
	// and above (i.e. towards the root control group) is owned and managed by the service manager of the host, while the control
	// group tree below the unit's control group is owned and managed by the unit itself.
	//
	// Takes either a boolean argument or a (possibly empty) list of control group controller names. If true, delegation is turned
	// on, and all supported controllers are enabled for the unit, making them available to the unit's processes for management.
	// If false, delegation is turned off entirely (and no additional controllers are enabled). If set to a list of controllers,
	// delegation is turned on, and the specified controllers are enabled for the unit. Assigning the empty string will enable
	// delegation, but reset the list of controllers, and all assignments prior to this will have no effect. Note that additional
	// controllers other than the ones specified might be made available as well, depending on configuration of the containing
	// slice unit or other units contained in it. Defaults to false.
	//
	// Note that controller delegation to less privileged code is only safe on the unified control group hierarchy. Accordingly,
	// access to the specified controllers will not be granted to unprivileged services on the legacy hierarchy, even when requested.
	//
	// The following controller names may be specified: cpu, cpuacct, cpuset, io, blkio, memory, devices, pids, bpf-firewall,
	// and bpf-devices.
	//
	// Not all of these controllers are available on all kernels however, and some are specific to the unified hierarchy while
	// others are specific to the legacy hierarchy. Also note that the kernel might support further controllers, which are not
	// covered here yet, as delegation is either not supported at all for them or not defined cleanly.
	//
	// Note that because of the hierarchical nature of cgroup hierarchy, any controllers that are delegated will be enabled for
	// the parent and sibling units of the unit with delegation.
	//
	// For further details on the delegation model consult Control Group APIs and Delegation.
	//
	// Added in version 218.
	Delegate systemdconf.Value

	// Place unit processes in the specified subgroup of the unit's control group. Takes a valid control group name (not a path!)
	// as parameter, or an empty string to turn this feature off. Defaults to off. The control group name must be usable as filename
	// and avoid conflicts with the kernel's control group attribute files (i.e. cgroup.procs is not an acceptable name, since
	// the kernel exposes a native control group attribute file by that name). This option has no effect unless control group delegation
	// is turned on via Delegate=, see above. Note that this setting only applies to "main" processes of a unit, i.e. for services
	// to ExecStart=, but not for ExecReload= and similar. If delegation is enabled, the latter are always placed inside a subgroup
	// named .control. The specified subgroup is automatically created (and potentially ownership is passed to the unit's configured
	// user/group) when a process is started in it.
	//
	// This option is useful to avoid manually moving the invoked process into a subgroup after it has been started. Since no processes
	// should live in inner nodes of the control group tree it is almost always necessary to run the main ("supervising") process
	// of a unit that has delegation turned on in a subgroup.
	//
	// Added in version 254.
	DelegateSubgroup systemdconf.Value

	// Disables controllers from being enabled for a unit's children. If a controller listed is already in use in its subtree,
	// the controller will be removed from the subtree. This can be used to avoid configuration in child units from being able to
	// implicitly or explicitly enable a controller. Defaults to empty.
	//
	// Multiple controllers may be specified, separated by spaces. You may also pass DisableControllers= multiple times, in
	// which case each new instance adds another controller to disable. Passing DisableControllers= by itself with no controller
	// name present resets the disabled controller list.
	//
	// It may not be possible to disable a controller after units have been started, if the unit or any child of the unit in question
	// delegates controllers to its children, as any delegated subtree of the cgroup hierarchy is unmanaged by systemd.
	//
	// The following controller names may be specified: cpu, cpuacct, cpuset, io, blkio, memory, devices, pids, bpf-firewall,
	// and bpf-devices.
	//
	// Added in version 240.
	DisableControllers systemdconf.Value

	// Specifies how systemd-oomd.service will act on this unit's cgroups. Defaults to auto.
	//
	// When set to kill, the unit becomes a candidate for monitoring by systemd-oomd. If the cgroup passes the limits set by oomd.conf
	// or the unit configuration, systemd-oomd will select a descendant cgroup and send SIGKILL to all of the processes under
	// it. You can find more details on candidates and kill behavior at systemd-oomd.service and oomd.conf.
	//
	// Setting either of these properties to kill will also result in After= and Wants= dependencies on systemd-oomd.service
	// unless DefaultDependencies=no.
	//
	// When set to auto, systemd-oomd will not actively use this cgroup's data for monitoring and detection. However, if an ancestor
	// cgroup has one of these properties set to kill, a unit with auto can still be a candidate for systemd-oomd to terminate.
	//
	// Added in version 247.
	ManagedOOMSwap systemdconf.Value

	// Specifies how systemd-oomd.service will act on this unit's cgroups. Defaults to auto.
	//
	// When set to kill, the unit becomes a candidate for monitoring by systemd-oomd. If the cgroup passes the limits set by oomd.conf
	// or the unit configuration, systemd-oomd will select a descendant cgroup and send SIGKILL to all of the processes under
	// it. You can find more details on candidates and kill behavior at systemd-oomd.service and oomd.conf.
	//
	// Setting either of these properties to kill will also result in After= and Wants= dependencies on systemd-oomd.service
	// unless DefaultDependencies=no.
	//
	// When set to auto, systemd-oomd will not actively use this cgroup's data for monitoring and detection. However, if an ancestor
	// cgroup has one of these properties set to kill, a unit with auto can still be a candidate for systemd-oomd to terminate.
	//
	// Added in version 247.
	ManagedOOMMemoryPressure systemdconf.Value

	// Overrides the default memory pressure limit set by oomd.conf for the cgroup of this unit. Takes a percentage value between
	// 0% and 100%, inclusive. Defaults to 0%, which means to use the default set by oomd.conf. This property is ignored unless
	// ManagedOOMMemoryPressure=kill.
	//
	// Added in version 247.
	ManagedOOMMemoryPressureLimit systemdconf.Value

	// Overrides the default memory pressure duration set by oomd.conf for the cgroup of this unit. The specified value supports
	// a time unit such as "ms" or "μs", see systemd.time for details on the permitted syntax. Must be set to either empty or a value
	// of at least 1s. Defaults to empty, which means to use the default set by oomd.conf. This property is ignored unless ManagedOOMMemoryPressure=kill.
	//
	// Added in version 257.
	ManagedOOMMemoryPressureDurationSec systemdconf.Value

	// Allows deprioritizing or omitting this unit's cgroup as a candidate when systemd-oomd needs to act. Requires support
	// for extended attributes (see xattr) in order to use avoid or omit.
	//
	// When calculating candidates to relieve swap usage, systemd-oomd will only respect these extended attributes if the unit's
	// cgroup is owned by root.
	//
	// When calculating candidates to relieve memory pressure, systemd-oomd will only respect these extended attributes if
	// the unit's cgroup is owned by root, or if the unit's cgroup owner, and the owner of the monitored ancestor cgroup are the same.
	// For example, if systemd-oomd is calculating candidates for -.slice, then extended attributes set on descendants of /user.slice/user-1000.slice/user@1000.service/
	// will be ignored because the descendants are owned by UID 1000, and -.slice is owned by UID 0. But, if calculating candidates
	// for /user.slice/user-1000.slice/user@1000.service/, then extended attributes set on the descendants would be respected.
	//
	// If this property is set to avoid, the service manager will convey this to systemd-oomd, which will only select this cgroup
	// if there are no other viable candidates.
	//
	// If this property is set to omit, the service manager will convey this to systemd-oomd, which will ignore this cgroup as a
	// candidate and will not perform any actions on it.
	//
	// It is recommended to use avoid and omit sparingly, as it can adversely affect systemd-oomd's kill behavior. Also note that
	// these extended attributes are not applied recursively to cgroups under this unit's cgroup.
	//
	// Defaults to none which means systemd-oomd will rank this unit's cgroup as defined in systemd-oomd.service and oomd.conf.
	//
	// Added in version 248.
	ManagedOOMPreference systemdconf.Value

	// Controls memory pressure monitoring for invoked processes. Takes a boolean or one of "auto" and "skip". If "no", tells
	// the service not to watch for memory pressure events, by setting the $MEMORY_PRESSURE_WATCH environment variable to the
	// literal string /dev/null. If "yes", tells the service to watch for memory pressure events. This enables memory accounting
	// for the service, and ensures the memory.pressure cgroup attribute file is accessible for reading and writing by the service's
	// user. It then sets the $MEMORY_PRESSURE_WATCH environment variable for processes invoked by the unit to the file system
	// path to this file. The threshold information configured with MemoryPressureThresholdSec= is encoded in the $MEMORY_PRESSURE_WRITE
	// environment variable. If the "auto" value is set the protocol is enabled if memory accounting is anyway enabled for the
	// unit, and disabled otherwise. If set to "skip" the logic is neither enabled, nor disabled and the two environment variables
	// are not set.
	//
	// Note that services are free to use the two environment variables, but it is unproblematic if they ignore them. Memory pressure
	// handling must be implemented individually in each service, and usually means different things for different software.
	// For further details on memory pressure handling see Memory Pressure Handling in systemd.
	//
	// Services implemented using sd-event may use sd_event_add_memory_pressure to watch for and handle memory pressure events.
	//
	// If not explicit set, defaults to the DefaultMemoryPressureWatch= setting in systemd-system.conf.
	//
	// Added in version 254.
	MemoryPressureWatch systemdconf.Value

	// Sets the memory pressure threshold time for memory pressure monitor as configured via MemoryPressureWatch=. Specifies
	// the maximum allocation latency before a memory pressure event is signalled to the service, per 2s window. If not specified,
	// defaults to the DefaultMemoryPressureThresholdSec= setting in systemd-system.conf (which in turn defaults to 200ms).
	// The specified value expects a time unit such as "ms" or "μs", see systemd.time for details on the permitted syntax.
	//
	// Added in version 254.
	MemoryPressureThresholdSec systemdconf.Value

	// Takes a boolean argument. This setting is used to enable coredump forwarding for containers that belong to this unit's
	// cgroup. Units with CoredumpReceive=yes must also be configured with Delegate=yes. Defaults to false.
	//
	// When systemd-coredump is handling a coredump for a process from a container, if the container's leader process is a descendant
	// of a cgroup with CoredumpReceive=yes and Delegate=yes, then systemd-coredump will attempt to forward the coredump to
	// systemd-coredump within the container. See also systemd-coredump.
	//
	// Added in version 255.
	CoredumpReceive systemdconf.Value
}
