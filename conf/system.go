// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// SystemFile represents systemd-system.conf, system.conf.d, systemd-user.conf, user.conf.d — System and session service manager configuration files
// (see https://www.freedesktop.org/software/systemd/man/systemd-system.conf.html for details)
type SystemFile struct {
	systemdconf.File

	Manager SystemManagerSection // Few settings controlling basic manager operations
}

// SystemManagerSection represents few settings controlling basic manager operations
// (see https://www.freedesktop.org/software/systemd/man/systemd-system.conf.html#Options for details)
type SystemManagerSection struct {
	systemdconf.Section

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	LogColor systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	LogLevel systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	LogLocation systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	LogTarget systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	LogTime systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	DumpCore systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	CrashChangeVT systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	CrashShell systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	CrashReboot systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	ShowStatus systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	DefaultStandardOutput systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	DefaultStandardError systemdconf.Value

	// Defines what action will be performed if user presses Ctrl-Alt-Delete more than 7 times in 2s. Can be set to "reboot-force",
	// "poweroff-force", "reboot-immediate", "poweroff-immediate" or disabled with "none". Defaults to "reboot-force".
	CtrlAltDelBurstAction systemdconf.Value

	// Configures the CPU affinity for the service manager as well as the default CPU affinity for all forked off processes. Takes
	// a list of CPU indices or ranges separated by either whitespace or commas. CPU ranges are specified by the lower and upper
	// CPU indices separated by a dash. This option may be specified more than once, in which case the specified CPU affinity masks
	// are merged. If the empty string is assigned, the mask is reset, all assignments prior to this will have no effect. Individual
	// services may override the CPU affinity for their processes with the CPUAffinity= setting in unit files, see systemd.exec.
	CPUAffinity systemdconf.Value

	// Configures the NUMA memory policy for the service manager and the default NUMA memory policy for all forked off processes.
	// Individual services may override the default policy with the NUMAPolicy= setting in unit files, see systemd.exec.
	NUMAPolicy systemdconf.Value

	// Configures the NUMA node mask that will be associated with the selected NUMA policy. Note that default and local NUMA policies
	// don't require explicit NUMA node mask and value of the option can be empty. Similarly to NUMAPolicy=, value can be overridden
	// by individual services in unit files, see systemd.exec.
	NUMAMask systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"). If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog
	// or the path specified with WatchdogDevice= or the kernel option systemd.watchdog-device=) will be programmed to automatically
	// reboot the system if it is not contacted within the specified timeout interval. The system manager will ensure to contact
	// it at least once in half the specified timeout interval. This feature requires a hardware watchdog device to be present,
	// as it is commonly the case in embedded and server systems. Not all hardware watchdogs allow configuration of all possible
	// reboot timeout values, in which case the closest available timeout is picked. RebootWatchdogSec= may be used to configure
	// the hardware watchdog when the system is asked to reboot. It works as a safety net to ensure that the reboot takes place even
	// if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout applies only to the second phase of the reboot,
	// i.e. after all regular services are already terminated, and after the system and service manager process (PID 1) got replaced
	// by the systemd-shutdown binary, see system bootup for details. During the first phase of the shutdown operation the system
	// and service manager remains running and hence RuntimeWatchdogSec= is still honoured. In order to define a timeout on this
	// first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction= in the [Unit] section of the shutdown.target
	// unit. By default RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec= to 10min. KExecWatchdogSec= may be
	// used to additionally enable the watchdog when kexec is being executed rather than when rebooting. Note that if the kernel
	// does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this case the watchdog might
	// not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec= is also enabled
	// at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec= is also enabled.
	// These settings have no effect if a hardware watchdog is not available.
	RuntimeWatchdogSec systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"). If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog
	// or the path specified with WatchdogDevice= or the kernel option systemd.watchdog-device=) will be programmed to automatically
	// reboot the system if it is not contacted within the specified timeout interval. The system manager will ensure to contact
	// it at least once in half the specified timeout interval. This feature requires a hardware watchdog device to be present,
	// as it is commonly the case in embedded and server systems. Not all hardware watchdogs allow configuration of all possible
	// reboot timeout values, in which case the closest available timeout is picked. RebootWatchdogSec= may be used to configure
	// the hardware watchdog when the system is asked to reboot. It works as a safety net to ensure that the reboot takes place even
	// if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout applies only to the second phase of the reboot,
	// i.e. after all regular services are already terminated, and after the system and service manager process (PID 1) got replaced
	// by the systemd-shutdown binary, see system bootup for details. During the first phase of the shutdown operation the system
	// and service manager remains running and hence RuntimeWatchdogSec= is still honoured. In order to define a timeout on this
	// first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction= in the [Unit] section of the shutdown.target
	// unit. By default RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec= to 10min. KExecWatchdogSec= may be
	// used to additionally enable the watchdog when kexec is being executed rather than when rebooting. Note that if the kernel
	// does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this case the watchdog might
	// not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec= is also enabled
	// at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec= is also enabled.
	// These settings have no effect if a hardware watchdog is not available.
	RebootWatchdogSec systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"). If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog
	// or the path specified with WatchdogDevice= or the kernel option systemd.watchdog-device=) will be programmed to automatically
	// reboot the system if it is not contacted within the specified timeout interval. The system manager will ensure to contact
	// it at least once in half the specified timeout interval. This feature requires a hardware watchdog device to be present,
	// as it is commonly the case in embedded and server systems. Not all hardware watchdogs allow configuration of all possible
	// reboot timeout values, in which case the closest available timeout is picked. RebootWatchdogSec= may be used to configure
	// the hardware watchdog when the system is asked to reboot. It works as a safety net to ensure that the reboot takes place even
	// if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout applies only to the second phase of the reboot,
	// i.e. after all regular services are already terminated, and after the system and service manager process (PID 1) got replaced
	// by the systemd-shutdown binary, see system bootup for details. During the first phase of the shutdown operation the system
	// and service manager remains running and hence RuntimeWatchdogSec= is still honoured. In order to define a timeout on this
	// first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction= in the [Unit] section of the shutdown.target
	// unit. By default RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec= to 10min. KExecWatchdogSec= may be
	// used to additionally enable the watchdog when kexec is being executed rather than when rebooting. Note that if the kernel
	// does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this case the watchdog might
	// not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec= is also enabled
	// at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec= is also enabled.
	// These settings have no effect if a hardware watchdog is not available.
	KExecWatchdogSec systemdconf.Value

	// Configure the hardware watchdog device that the runtime and shutdown watchdog timers will open and use. Defaults to /dev/watchdog.
	// This setting has no effect if a hardware watchdog is not available.
	WatchdogDevice systemdconf.Value

	// Controls which capabilities to include in the capability bounding set for PID 1 and its children. See capabilities for
	// details. Takes a whitespace-separated list of capability names as read by cap_from_name. Capabilities listed will be
	// included in the bounding set, all others are removed. If the list of capabilities is prefixed with ~, all but the listed capabilities
	// will be included, the effect of the assignment inverted. Note that this option also affects the respective capabilities
	// in the effective, permitted and inheritable capability sets. The capability bounding set may also be individually configured
	// for units using the CapabilityBoundingSet= directive for units, but note that capabilities dropped for PID 1 cannot be
	// regained in individual units, they are lost for good.
	CapabilityBoundingSet systemdconf.Value

	// Takes a boolean argument. If true, ensures that PID 1 and all its children can never gain new privileges through execve (e.g.
	// via setuid or setgid bits, or filesystem capabilities). Defaults to false. General purpose distributions commonly rely
	// on executables with setuid or setgid bits and will thus not function properly with this option enabled. Individual units
	// cannot disable this option. Also see No New Privileges Flag.
	NoNewPrivileges systemdconf.Value

	// Takes a space-separated list of architecture identifiers. Selects from which architectures system calls may be invoked
	// on this system. This may be used as an effective way to disable invocation of non-native binaries system-wide, for example
	// to prohibit execution of 32-bit x86 binaries on 64-bit x86-64 systems. This option operates system-wide, and acts similar
	// to the SystemCallArchitectures= setting of unit files, see systemd.exec for details. This setting defaults to the empty
	// list, in which case no filtering of system calls based on architecture is applied. Known architecture identifiers are
	// "x86", "x86-64", "x32", "arm" and the special identifier "native". The latter implicitly maps to the native architecture
	// of the system (or more specifically, the architecture the system manager was compiled for). Set this setting to "native"
	// to prohibit execution of any non-native binaries. When a binary executes a system call of an architecture that is not listed
	// in this setting, it will be immediately terminated with the SIGSYS signal.
	SystemCallArchitectures systemdconf.Value

	// Sets the timer slack in nanoseconds for PID 1, which is inherited by all executed processes, unless overridden individually,
	// for example with the TimerSlackNSec= setting in service units (for details see systemd.exec). The timer slack controls
	// the accuracy of wake-ups triggered by system timers. See prctl for more information. Note that in contrast to most other
	// time span definitions this parameter takes an integer value in nano-seconds if no unit is specified. The usual time units
	// are understood too.
	TimerSlackNSec systemdconf.Value

	// Takes either name or description as the value. If name, the system manager will use unit names in status messages, instead
	// of the longer and more informative descriptions set with Description=, see systemd.unit.
	StatusUnitFormat systemdconf.Value

	// Sets the default accuracy of timer units. This controls the global default for the AccuracySec= setting of timer units,
	// see systemd.timer for details. AccuracySec= set in individual units override the global default for the specific unit.
	// Defaults to 1min. Note that the accuracy of timer units is also affected by the configured timer slack for PID 1, see TimerSlackNSec=
	// above.
	DefaultTimerAccuracySec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). Disabled by default, when service with Type=oneshot
	// is used. For non-service units, DefaultTimeoutStartSec= sets the default TimeoutSec= value. DefaultTimeoutStartSec=
	// and DefaultTimeoutStopSec= default to 90s. DefaultTimeoutAbortSec= is not set by default so that all units fall back
	// to TimeoutStopSec=. DefaultRestartSec= defaults to 100ms.
	DefaultTimeoutStartSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). Disabled by default, when service with Type=oneshot
	// is used. For non-service units, DefaultTimeoutStartSec= sets the default TimeoutSec= value. DefaultTimeoutStartSec=
	// and DefaultTimeoutStopSec= default to 90s. DefaultTimeoutAbortSec= is not set by default so that all units fall back
	// to TimeoutStopSec=. DefaultRestartSec= defaults to 100ms.
	DefaultTimeoutStopSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). Disabled by default, when service with Type=oneshot
	// is used. For non-service units, DefaultTimeoutStartSec= sets the default TimeoutSec= value. DefaultTimeoutStartSec=
	// and DefaultTimeoutStopSec= default to 90s. DefaultTimeoutAbortSec= is not set by default so that all units fall back
	// to TimeoutStopSec=. DefaultRestartSec= defaults to 100ms.
	DefaultTimeoutAbortSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). Disabled by default, when service with Type=oneshot
	// is used. For non-service units, DefaultTimeoutStartSec= sets the default TimeoutSec= value. DefaultTimeoutStartSec=
	// and DefaultTimeoutStopSec= default to 90s. DefaultTimeoutAbortSec= is not set by default so that all units fall back
	// to TimeoutStopSec=. DefaultRestartSec= defaults to 100ms.
	DefaultRestartSec systemdconf.Value

	// Configure the default unit start rate limiting, as configured per-service by StartLimitIntervalSec= and StartLimitBurst=.
	// See systemd.service for details on the per-service settings. DefaultStartLimitIntervalSec= defaults to 10s. DefaultStartLimitBurst=
	// defaults to 5.
	DefaultStartLimitIntervalSec systemdconf.Value

	// Configure the default unit start rate limiting, as configured per-service by StartLimitIntervalSec= and StartLimitBurst=.
	// See systemd.service for details on the per-service settings. DefaultStartLimitIntervalSec= defaults to 10s. DefaultStartLimitBurst=
	// defaults to 5.
	DefaultStartLimitBurst systemdconf.Value

	// Sets manager environment variables passed to all executed processes. Takes a space-separated list of variable assignments.
	// See environ for details about environment variables.
	//
	// Example:
	//
	// 	DefaultEnvironment="VAR1=word1 word2" VAR2=word3 "VAR3=word 5 6"
	//
	// Sets three variables "VAR1", "VAR2", "VAR3".
	DefaultEnvironment systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultCPUAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultBlockIOAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultMemoryAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultTasksAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultIOAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, BlockIOAccounting=,
	// MemoryAccounting=, TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details
	// on the per-unit settings. DefaultTasksAccounting= defaults to yes, DefaultMemoryAccounting= to yes. DefaultCPUAccounting=
	// defaults to yes if enabling CPU accounting doesn't require the CPU controller to be enabled (Linux 4.15+ using the unified
	// hierarchy for resource control), otherwise it defaults to no. The other three settings default to no.
	DefaultIPAccounting systemdconf.Value

	// Configure the default value for the per-unit TasksMax= setting. See systemd.resource-control for details. This setting
	// applies to all unit types that support resource control settings, with the exception of slice units. Defaults to 15% of
	// the sysctl setting kernel.pid_max= or root cgroup pids.max. Kernel has a default value for kernel.pid_max= and an algorithm
	// of counting in case of more than 32 cores. For example with the default kernel.pid_max=, DefaultTasksMax= defaults to
	// 4915, but might be greater in other systems or smaller in OS containers.
	DefaultTasksMax systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitCPU systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitFSIZE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitDATA systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitSTACK systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitCORE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitRSS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitNOFILE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitAS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitNPROC systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitMEMLOCK systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitLOCKS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitSIGPENDING systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitMSGQUEUE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitNICE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitRTPRIO systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	DefaultLimitRTTIME systemdconf.Value

	// Configure the default policy for reacting to processes being killed by the Linux Out-Of-Memory (OOM) killer. This may
	// be used to pick a global default for the per-unit OOMPolicy= setting. See systemd.service for details. Note that this default
	// is not used for services that have Delegate= turned on.
	DefaultOOMPolicy systemdconf.Value
}
