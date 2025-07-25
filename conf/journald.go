// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// JournaldFile represents journald.conf, journald.conf.d, journald@.conf — Journal service configuration files
// (see https://www.freedesktop.org/software/systemd/man/journald.conf.html for details).
type JournaldFile struct {
	systemdconf.File

	Journal JournaldJournalSection // parameters of the systemd journal service, systemd-journald.service
}

// JournaldJournalSection represents parameters of the systemd journal service, systemd-journald.service
// (see https://www.freedesktop.org/software/systemd/man/journald.conf.html#Options for details).
type JournaldJournalSection struct {
	systemdconf.Section

	// Controls where to store journal data. One of "volatile", "persistent", "auto" and "none". If "volatile", journal log
	// data will be stored only in memory, i.e. below the /run/log/journal hierarchy (which is created if needed). If "persistent",
	// data will be stored preferably on disk, i.e. below the /var/log/journal hierarchy (which is created if needed), with a
	// fallback to /run/log/journal (which is created if needed), during early boot and if the disk is not writable. "auto" behaves
	// like "persistent" if the /var/log/journal directory exists, and "volatile" otherwise (the existence of the directory
	// controls the storage mode). "none" turns off all storage, all log data received will be dropped (but forwarding to other
	// targets, such as the console, the kernel log buffer, or a syslog socket will still work). Defaults to "auto" in the default
	// journal namespace, and "persistent" in all others.
	//
	// Note that journald will initially use volatile storage, until a call to journalctl --flush (or sending SIGUSR1 to journald)
	// will cause it to switch to persistent logging (under the conditions mentioned above). This is done automatically on boot
	// via "systemd-journal-flush.service".
	//
	// Note that when this option is changed to "volatile", existing persistent data is not removed. In the other direction, journalctl
	// with the --flush option may be used to move volatile data to persistent storage.
	//
	// When journal namespacing (see LogNamespace= in systemd.exec) is used, setting Storage= to "volatile" or "auto" will
	// not have an effect on the creation of the per-namespace logs directory in /var/log/journal/, as the systemd-journald@.service
	// service file by default carries LogsDirectory=. To turn that off, add a unit file drop-in file that sets LogsDirectory=
	// to an empty string.
	//
	// Note that per-user journal files are not supported unless persistent storage is enabled, thus making journalctl --user
	// unavailable.
	//
	// The storage to use can also be specified via the "journal.storage" credential. Values configured via configuration files
	// take priority over values configured via the credential.
	//
	// Added in version 186.
	Storage systemdconf.Value

	// Can take a boolean value. If enabled (the default), data objects that shall be stored in the journal and are larger than the
	// default threshold of 512 bytes are compressed before they are written to the file system. It can also be set to a number of
	// bytes to specify the compression threshold directly. Suffixes like K, M, and G can be used to specify larger units.
	Compress systemdconf.Value

	// Takes a boolean value. If enabled (the default), and a sealing key is available (as created by journalctl's --setup-keys
	// command), Forward Secure Sealing (FSS) for all persistent journal files is enabled. FSS is based on Seekable Sequential
	// Key Generators by G. A. Marson and B. Poettering (doi:10.1007/978-3-642-40203-6_7) and may be used to protect journal
	// files from unnoticed alteration.
	//
	// Added in version 189.
	Seal systemdconf.Value

	// Controls whether to split up journal files per user, either "uid" or "none". Split journal files are primarily useful for
	// access control: on UNIX/Linux access control is managed per file, and the journal daemon will assign users read access
	// to their journal files. If "uid", all regular users (with UID outside the range of system users, dynamic service users,
	// and the nobody user) will each get their own journal files, and system users will log to the system journal. See Users, Groups,
	// UIDs and GIDs on systemd systems for more details about UID ranges. If "none", journal files are not split up by user and all
	// messages are instead stored in the single system journal. In this mode unprivileged users generally do not have access
	// to their own log data. Note that splitting up journal files by user is only available for journals stored persistently.
	// If journals are stored on volatile storage (see Storage= above), only a single journal file is used. Defaults to "uid".
	//
	// Added in version 190.
	SplitMode systemdconf.Value

	// Configures the rate limiting that is applied to all messages generated on the system. If, in the time interval defined by
	// RateLimitIntervalSec=, more messages than specified in RateLimitBurst= are logged by a service, all further messages
	// within the interval are dropped until the interval is over. A message about the number of dropped messages is generated.
	// This rate limiting is applied per-service, so that two services which log do not interfere with each other's limits. Defaults
	// to 10000 messages in 30s. The time specification for RateLimitIntervalSec= may be specified in the following units: "s",
	// "min", "h", "ms", "us". To turn off any kind of rate limiting, set either value to 0.
	//
	// Note that the effective rate limit is multiplied by a factor derived from the available free disk space for the journal.
	// Currently, this factor is calculated using the base 2 logarithm.
	//
	//	+----------------------+------------------+
	//	| AVAILABLE DISK SPACE | BURST MULTIPLIER |
	//	+----------------------+------------------+
	//	| <= 1MB               |                1 |
	//	| <= 16MB              |                2 |
	//	| <= 256MB             |                3 |
	//	| <= 4GB               |                4 |
	//	| <= 64GB              |                5 |
	//	| <= 1TB               |                6 |
	//	+----------------------+------------------+
	//
	// If a service provides rate limits for itself through LogRateLimitIntervalSec= and/or LogRateLimitBurst= in systemd.exec,
	// those values will override the settings specified here.
	RateLimitIntervalSec systemdconf.Value

	// Configures the rate limiting that is applied to all messages generated on the system. If, in the time interval defined by
	// RateLimitIntervalSec=, more messages than specified in RateLimitBurst= are logged by a service, all further messages
	// within the interval are dropped until the interval is over. A message about the number of dropped messages is generated.
	// This rate limiting is applied per-service, so that two services which log do not interfere with each other's limits. Defaults
	// to 10000 messages in 30s. The time specification for RateLimitIntervalSec= may be specified in the following units: "s",
	// "min", "h", "ms", "us". To turn off any kind of rate limiting, set either value to 0.
	//
	// Note that the effective rate limit is multiplied by a factor derived from the available free disk space for the journal.
	// Currently, this factor is calculated using the base 2 logarithm.
	//
	//	+----------------------+------------------+
	//	| AVAILABLE DISK SPACE | BURST MULTIPLIER |
	//	+----------------------+------------------+
	//	| <= 1MB               |                1 |
	//	| <= 16MB              |                2 |
	//	| <= 256MB             |                3 |
	//	| <= 4GB               |                4 |
	//	| <= 64GB              |                5 |
	//	| <= 1TB               |                6 |
	//	+----------------------+------------------+
	//
	// If a service provides rate limits for itself through LogRateLimitIntervalSec= and/or LogRateLimitBurst= in systemd.exec,
	// those values will override the settings specified here.
	RateLimitBurst systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	SystemMaxUse systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	SystemKeepFree systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	SystemMaxFileSize systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	SystemMaxFiles systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	RuntimeMaxUse systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	RuntimeKeepFree systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	RuntimeMaxFileSize systemdconf.Value

	// Enforce size limits on the journal files stored. The options prefixed with "System" apply to the journal files when stored
	// on a persistent file system, more specifically /var/log/journal. The options prefixed with "Runtime" apply to the journal
	// files when stored on a volatile in-memory file system, more specifically /run/log/journal. The former is used only when
	// /var/ is mounted, writable, and the directory /var/log/journal exists. Otherwise, only the latter applies. Note that
	// this means that during early boot and if the administrator disabled persistent logging, only the latter options apply,
	// while the former apply if persistent logging is enabled and the system is fully booted up. journalctl and systemd-journald
	// ignore all files with names not ending with ".journal" or ".journal~", so only such files, located in the appropriate directories,
	// are taken into account when calculating current disk usage.
	//
	// SystemMaxUse= and RuntimeMaxUse= control how much disk space the journal may use up at most. SystemKeepFree= and RuntimeKeepFree=
	// control how much disk space systemd-journald shall leave free for other uses. systemd-journald will respect both limits
	// and use the smaller of the two values.
	//
	// The first pair defaults to 10% and the second to 15% of the size of the respective file system, but each of the calculated default
	// values is capped to 4G. If the file system is nearly full and either SystemKeepFree= or RuntimeKeepFree= are violated when
	// systemd-journald is started, the limit will be raised to the percentage that is actually free. This means that if there
	// was enough free space before and journal files were created, and subsequently something else causes the file system to
	// fill up, journald will stop using more space, but it will not be removing existing files to reduce the footprint again, either.
	// Also note that only archived files are deleted to reduce the space occupied by journal files. This means that, in effect,
	// there might still be more space used than SystemMaxUse= or RuntimeMaxUse= limit after a vacuuming operation is complete.
	//
	// SystemMaxFileSize= and RuntimeMaxFileSize= control how large individual journal files may grow at most. This influences
	// the granularity in which disk space is made available through rotation, i.e. deletion of historic data. Defaults to one
	// eighth of the values configured with SystemMaxUse= and RuntimeMaxUse= capped to 128M, so that usually seven rotated journal
	// files are kept as history. If the journal compact mode is enabled (enabled by default), the maximum file size is capped to
	// 4G.
	//
	// Specify values in bytes or use K, M, G, T, P, E as units for the specified sizes (equal to 1024, 1024², … bytes). Note that size
	// limits are enforced synchronously when journal files are extended, and no explicit rotation step triggered by time is
	// needed.
	//
	// SystemMaxFiles= and RuntimeMaxFiles= control how many individual journal files to keep at most. Note that only archived
	// files are deleted to reduce the number of files until this limit is reached; active files will stay around. This means that,
	// in effect, there might still be more journal files around in total than this limit after a vacuuming operation is complete.
	// This setting defaults to 100.
	RuntimeMaxFiles systemdconf.Value

	// The maximum time to store entries in a single journal file before rotating to the next one. Normally, time-based rotation
	// should not be required as size-based rotation with options such as SystemMaxFileSize= should be sufficient to ensure
	// that journal files do not grow without bounds. However, to ensure that not too much data is lost at once when old journal files
	// are deleted, it might make sense to change this value from the default of one month. Set to 0 to turn off this feature. This
	// setting takes time values which may be suffixed with the units "year", "month", "week", "day", "h" or "m" to override the
	// default time unit of seconds.
	//
	// Added in version 195.
	MaxFileSec systemdconf.Value

	// The maximum time to store journal entries. This controls whether journal files containing entries older than the specified
	// time span are deleted. Normally, time-based deletion of old journal files should not be required as size-based deletion
	// with options such as SystemMaxUse= should be sufficient to ensure that journal files do not grow without bounds. However,
	// to enforce data retention policies, it might make sense to change this value from the default of 0 (which turns off this feature).
	// This setting also takes time values which may be suffixed with the units "year", "month", "week", "day", "h" or " m" to override
	// the default time unit of seconds.
	//
	// Added in version 195.
	MaxRetentionSec systemdconf.Value

	// The timeout before synchronizing journal files to disk. After syncing, journal files are placed in the OFFLINE state.
	// Note that syncing is unconditionally done immediately after a log message of priority CRIT, ALERT or EMERG has been logged.
	// This setting hence applies only to messages of the levels ERR, WARNING, NOTICE, INFO, DEBUG. The default timeout is 5 minutes.
	//
	// Added in version 199.
	SyncIntervalSec systemdconf.Value

	// Control whether log messages received by the journal daemon shall be forwarded to a traditional syslog daemon, to the kernel
	// log buffer (kmsg), to the system console, sent as wall messages to all logged-in users or sent over a socket. These options
	// take boolean arguments except for "ForwardToSocket=" which takes an address instead. If forwarding to syslog is enabled
	// but nothing reads messages from the socket, forwarding to syslog has no effect. By default, only forwarding to wall is enabled.
	// These settings may be overridden at boot time with the kernel command line options "systemd.journald.forward_to_syslog",
	// "systemd.journald.forward_to_kmsg", "systemd.journald.forward_to_console", and "systemd.journald.forward_to_wall".
	// If the option name is specified without "=" and the following argument, true is assumed. Otherwise, the argument is parsed
	// as a boolean.
	//
	// The socket forwarding address can be specified with the credential "journal.forward_to_socket". The following socket
	// types are supported:
	//
	// AF_INET (e.g. "192.168.0.11:4444"), AF_INET6 (e.g. "[2001:db8::ff00:42:8329]:4444"), AF_UNIX (e.g. "/run/host/journal/socket"),
	// AF_VSOCK (e.g. "vsock:2:1234")
	//
	// When forwarding to the console, the TTY to log to can be changed with TTYPath=, described below.
	//
	// When forwarding to the kernel log buffer (kmsg), make sure to select a suitably large size for the log buffer, for example
	// by adding "log_buf_len=8M" to the kernel command line. systemd will automatically disable kernel's rate-limiting applied
	// to userspace processes (equivalent to setting "printk.devkmsg=on").
	//
	// When forwarding over a socket the  Journal Export Format is used when sending over the wire. Notably this includes the metadata
	// field __REALTIME_TIMESTAMP so that systemd-journal-remote (see systemd-journal-remote.service) can be used to receive
	// the forwarded journal entries.
	//
	// Note: Forwarding is performed synchronously within journald, and may significantly affect its performance. This is
	// particularly relevant when using ForwardToConsole=yes in cloud environments, where the console is often a slow, virtual
	// serial port. Since journald is implemented as a conventional single-process daemon, forwarding to a completely hung
	// console will block journald. This can have a cascading effect resulting in any services synchronously logging to the blocked
	// journal also becoming blocked. Unless actively debugging/developing something, it is generally preferable to setup
	// a journalctl --follow style service redirected to the console, instead of ForwardToConsole=yes, for production use.
	ForwardToSyslog systemdconf.Value

	// Control whether log messages received by the journal daemon shall be forwarded to a traditional syslog daemon, to the kernel
	// log buffer (kmsg), to the system console, sent as wall messages to all logged-in users or sent over a socket. These options
	// take boolean arguments except for "ForwardToSocket=" which takes an address instead. If forwarding to syslog is enabled
	// but nothing reads messages from the socket, forwarding to syslog has no effect. By default, only forwarding to wall is enabled.
	// These settings may be overridden at boot time with the kernel command line options "systemd.journald.forward_to_syslog",
	// "systemd.journald.forward_to_kmsg", "systemd.journald.forward_to_console", and "systemd.journald.forward_to_wall".
	// If the option name is specified without "=" and the following argument, true is assumed. Otherwise, the argument is parsed
	// as a boolean.
	//
	// The socket forwarding address can be specified with the credential "journal.forward_to_socket". The following socket
	// types are supported:
	//
	// AF_INET (e.g. "192.168.0.11:4444"), AF_INET6 (e.g. "[2001:db8::ff00:42:8329]:4444"), AF_UNIX (e.g. "/run/host/journal/socket"),
	// AF_VSOCK (e.g. "vsock:2:1234")
	//
	// When forwarding to the console, the TTY to log to can be changed with TTYPath=, described below.
	//
	// When forwarding to the kernel log buffer (kmsg), make sure to select a suitably large size for the log buffer, for example
	// by adding "log_buf_len=8M" to the kernel command line. systemd will automatically disable kernel's rate-limiting applied
	// to userspace processes (equivalent to setting "printk.devkmsg=on").
	//
	// When forwarding over a socket the  Journal Export Format is used when sending over the wire. Notably this includes the metadata
	// field __REALTIME_TIMESTAMP so that systemd-journal-remote (see systemd-journal-remote.service) can be used to receive
	// the forwarded journal entries.
	//
	// Note: Forwarding is performed synchronously within journald, and may significantly affect its performance. This is
	// particularly relevant when using ForwardToConsole=yes in cloud environments, where the console is often a slow, virtual
	// serial port. Since journald is implemented as a conventional single-process daemon, forwarding to a completely hung
	// console will block journald. This can have a cascading effect resulting in any services synchronously logging to the blocked
	// journal also becoming blocked. Unless actively debugging/developing something, it is generally preferable to setup
	// a journalctl --follow style service redirected to the console, instead of ForwardToConsole=yes, for production use.
	ForwardToKMsg systemdconf.Value

	// Control whether log messages received by the journal daemon shall be forwarded to a traditional syslog daemon, to the kernel
	// log buffer (kmsg), to the system console, sent as wall messages to all logged-in users or sent over a socket. These options
	// take boolean arguments except for "ForwardToSocket=" which takes an address instead. If forwarding to syslog is enabled
	// but nothing reads messages from the socket, forwarding to syslog has no effect. By default, only forwarding to wall is enabled.
	// These settings may be overridden at boot time with the kernel command line options "systemd.journald.forward_to_syslog",
	// "systemd.journald.forward_to_kmsg", "systemd.journald.forward_to_console", and "systemd.journald.forward_to_wall".
	// If the option name is specified without "=" and the following argument, true is assumed. Otherwise, the argument is parsed
	// as a boolean.
	//
	// The socket forwarding address can be specified with the credential "journal.forward_to_socket". The following socket
	// types are supported:
	//
	// AF_INET (e.g. "192.168.0.11:4444"), AF_INET6 (e.g. "[2001:db8::ff00:42:8329]:4444"), AF_UNIX (e.g. "/run/host/journal/socket"),
	// AF_VSOCK (e.g. "vsock:2:1234")
	//
	// When forwarding to the console, the TTY to log to can be changed with TTYPath=, described below.
	//
	// When forwarding to the kernel log buffer (kmsg), make sure to select a suitably large size for the log buffer, for example
	// by adding "log_buf_len=8M" to the kernel command line. systemd will automatically disable kernel's rate-limiting applied
	// to userspace processes (equivalent to setting "printk.devkmsg=on").
	//
	// When forwarding over a socket the  Journal Export Format is used when sending over the wire. Notably this includes the metadata
	// field __REALTIME_TIMESTAMP so that systemd-journal-remote (see systemd-journal-remote.service) can be used to receive
	// the forwarded journal entries.
	//
	// Note: Forwarding is performed synchronously within journald, and may significantly affect its performance. This is
	// particularly relevant when using ForwardToConsole=yes in cloud environments, where the console is often a slow, virtual
	// serial port. Since journald is implemented as a conventional single-process daemon, forwarding to a completely hung
	// console will block journald. This can have a cascading effect resulting in any services synchronously logging to the blocked
	// journal also becoming blocked. Unless actively debugging/developing something, it is generally preferable to setup
	// a journalctl --follow style service redirected to the console, instead of ForwardToConsole=yes, for production use.
	ForwardToConsole systemdconf.Value

	// Control whether log messages received by the journal daemon shall be forwarded to a traditional syslog daemon, to the kernel
	// log buffer (kmsg), to the system console, sent as wall messages to all logged-in users or sent over a socket. These options
	// take boolean arguments except for "ForwardToSocket=" which takes an address instead. If forwarding to syslog is enabled
	// but nothing reads messages from the socket, forwarding to syslog has no effect. By default, only forwarding to wall is enabled.
	// These settings may be overridden at boot time with the kernel command line options "systemd.journald.forward_to_syslog",
	// "systemd.journald.forward_to_kmsg", "systemd.journald.forward_to_console", and "systemd.journald.forward_to_wall".
	// If the option name is specified without "=" and the following argument, true is assumed. Otherwise, the argument is parsed
	// as a boolean.
	//
	// The socket forwarding address can be specified with the credential "journal.forward_to_socket". The following socket
	// types are supported:
	//
	// AF_INET (e.g. "192.168.0.11:4444"), AF_INET6 (e.g. "[2001:db8::ff00:42:8329]:4444"), AF_UNIX (e.g. "/run/host/journal/socket"),
	// AF_VSOCK (e.g. "vsock:2:1234")
	//
	// When forwarding to the console, the TTY to log to can be changed with TTYPath=, described below.
	//
	// When forwarding to the kernel log buffer (kmsg), make sure to select a suitably large size for the log buffer, for example
	// by adding "log_buf_len=8M" to the kernel command line. systemd will automatically disable kernel's rate-limiting applied
	// to userspace processes (equivalent to setting "printk.devkmsg=on").
	//
	// When forwarding over a socket the  Journal Export Format is used when sending over the wire. Notably this includes the metadata
	// field __REALTIME_TIMESTAMP so that systemd-journal-remote (see systemd-journal-remote.service) can be used to receive
	// the forwarded journal entries.
	//
	// Note: Forwarding is performed synchronously within journald, and may significantly affect its performance. This is
	// particularly relevant when using ForwardToConsole=yes in cloud environments, where the console is often a slow, virtual
	// serial port. Since journald is implemented as a conventional single-process daemon, forwarding to a completely hung
	// console will block journald. This can have a cascading effect resulting in any services synchronously logging to the blocked
	// journal also becoming blocked. Unless actively debugging/developing something, it is generally preferable to setup
	// a journalctl --follow style service redirected to the console, instead of ForwardToConsole=yes, for production use.
	ForwardToWall systemdconf.Value

	// Control whether log messages received by the journal daemon shall be forwarded to a traditional syslog daemon, to the kernel
	// log buffer (kmsg), to the system console, sent as wall messages to all logged-in users or sent over a socket. These options
	// take boolean arguments except for "ForwardToSocket=" which takes an address instead. If forwarding to syslog is enabled
	// but nothing reads messages from the socket, forwarding to syslog has no effect. By default, only forwarding to wall is enabled.
	// These settings may be overridden at boot time with the kernel command line options "systemd.journald.forward_to_syslog",
	// "systemd.journald.forward_to_kmsg", "systemd.journald.forward_to_console", and "systemd.journald.forward_to_wall".
	// If the option name is specified without "=" and the following argument, true is assumed. Otherwise, the argument is parsed
	// as a boolean.
	//
	// The socket forwarding address can be specified with the credential "journal.forward_to_socket". The following socket
	// types are supported:
	//
	// AF_INET (e.g. "192.168.0.11:4444"), AF_INET6 (e.g. "[2001:db8::ff00:42:8329]:4444"), AF_UNIX (e.g. "/run/host/journal/socket"),
	// AF_VSOCK (e.g. "vsock:2:1234")
	//
	// When forwarding to the console, the TTY to log to can be changed with TTYPath=, described below.
	//
	// When forwarding to the kernel log buffer (kmsg), make sure to select a suitably large size for the log buffer, for example
	// by adding "log_buf_len=8M" to the kernel command line. systemd will automatically disable kernel's rate-limiting applied
	// to userspace processes (equivalent to setting "printk.devkmsg=on").
	//
	// When forwarding over a socket the  Journal Export Format is used when sending over the wire. Notably this includes the metadata
	// field __REALTIME_TIMESTAMP so that systemd-journal-remote (see systemd-journal-remote.service) can be used to receive
	// the forwarded journal entries.
	//
	// Note: Forwarding is performed synchronously within journald, and may significantly affect its performance. This is
	// particularly relevant when using ForwardToConsole=yes in cloud environments, where the console is often a slow, virtual
	// serial port. Since journald is implemented as a conventional single-process daemon, forwarding to a completely hung
	// console will block journald. This can have a cascading effect resulting in any services synchronously logging to the blocked
	// journal also becoming blocked. Unless actively debugging/developing something, it is generally preferable to setup
	// a journalctl --follow style service redirected to the console, instead of ForwardToConsole=yes, for production use.
	ForwardToSocket systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelStore systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelSyslog systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelKMsg systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelConsole systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelWall systemdconf.Value

	// Controls the maximum log level of messages that are stored in the journal, forwarded to syslog, kmsg, the console, the wall,
	// or a socket (if that is enabled, see above). As argument, takes one of "emerg", "alert", "crit", "err", "warning", "notice",
	// "info", "debug", or integer values in the range of 0–7 (corresponding to the same levels). Messages equal or below the log
	// level specified are stored/forwarded, messages above are dropped. Defaults to "debug" for MaxLevelStore=, MaxLevelSyslog=
	// and MaxLevelSocket=, to ensure that the all messages are stored in the journal, forwarded to syslog and the socket if one
	// exists. Defaults to "notice" for MaxLevelKMsg=, "info" for MaxLevelConsole=, and "emerg" for MaxLevelWall=. These
	// settings may be overridden at boot time with the kernel command line options "systemd.journald.max_level_store=",
	// "systemd.journald.max_level_syslog=", "systemd.journald.max_level_kmsg=", "systemd.journald.max_level_console=",
	// "systemd.journald.max_level_wall=", "systemd.journald.max_level_socket=".
	//
	// Added in version 185.
	MaxLevelSocket systemdconf.Value

	// Takes a boolean value. If enabled systemd-journal processes /dev/kmsg messages generated by the kernel. In the default
	// journal namespace this option is enabled by default, it is disabled in all others.
	//
	// Added in version 235.
	ReadKMsg systemdconf.Value

	// Takes a boolean value. If enabled systemd-journald will turn on kernel auditing on start-up. If disabled it will turn it
	// off. If unset it will neither enable nor disable it, leaving the previous state unchanged. This means if another tool turns
	// on auditing even if systemd-journald left it off, it will still collect the generated messages. Defaults to on.
	//
	// Note that this option does not control whether systemd-journald collects generated audit records, it just controls whether
	// it tells the kernel to generate them. If you need to prevent systemd-journald from collecting the generated messages,
	// the socket unit "systemd-journald-audit.socket" can be disabled and, in this case, this setting is without effect.
	//
	// Added in version 246.
	Audit systemdconf.Value

	// Change the console TTY to use if ForwardToConsole=yes is used. Defaults to /dev/console.
	//
	// Added in version 185.
	TTYPath systemdconf.Value

	// The maximum line length to permit when converting stream logs into record logs. When a systemd unit's standard output/error
	// are connected to the journal via a stream socket, the data read is split into individual log records at newline ("\n", ASCII
	// 10) and NUL characters. If no such delimiter is read for the specified number of bytes a hard log record boundary is artificially
	// inserted, breaking up overly long lines into multiple log records. Selecting overly large values increases the possible
	// memory usage of the Journal daemon for each stream client, as in the worst case the journal daemon needs to buffer the specified
	// number of bytes in memory before it can flush a new log record to disk. Also note that permitting overly large line maximum
	// line lengths affects compatibility with traditional log protocols as log records might not fit anymore into a single AF_UNIX
	// or AF_INET datagram. Takes a size in bytes. If the value is suffixed with K, M, G or T, the specified size is parsed as Kilobytes,
	// Megabytes, Gigabytes, or Terabytes (with the base 1024), respectively. Defaults to 48K, which is relatively large but
	// still small enough so that log records likely fit into network datagrams along with extra room for metadata. Note that values
	// below 79 are not accepted and will be bumped to 79.
	//
	// Added in version 235.
	LineMax systemdconf.Value
}
