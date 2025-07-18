// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// SystemFile represents systemd-system.conf, system.conf.d, systemd-user.conf, user.conf.d — System and session service manager configuration files
// (see https://www.freedesktop.org/software/systemd/man/systemd-system.conf.html for details).
type SystemFile struct {
	systemdconf.File

	Manager SystemManagerSection // settings controlling basic manager operations
}

// SystemManagerSection represents settings controlling basic manager operations
// (see https://www.freedesktop.org/software/systemd/man/systemd-system.conf.html for details).
type SystemManagerSection struct {
	systemdconf.Section

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	LogColor systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	LogLevel systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	LogLocation systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	LogTarget systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	LogTime systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	DumpCore systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	CrashChangeVT systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	CrashShell systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	CrashAction systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	ShowStatus systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	DefaultStandardOutput systemdconf.Value

	// Configures various parameters of basic manager operation. These options may be overridden by the respective process
	// and kernel command line arguments. See systemd for details.
	//
	// Added in version 198.
	DefaultStandardError systemdconf.Value

	// Defines what action will be performed if user presses Ctrl-Alt-Delete more than 7 times in 2s. Can be set to "reboot-force",
	// "poweroff-force", "reboot-immediate", "poweroff-immediate" or disabled with "none". Defaults to "reboot-force".
	//
	// Added in version 232.
	CtrlAltDelBurstAction systemdconf.Value

	// Takes name, description or combined as the value. If name, the system manager will use unit names in status messages (e.g.
	// "systemd-journald.service"), instead of the longer and more informative descriptions set with Description= (e.g.
	// "Journal Logging Service"). If combined, the system manager will use both unit names and descriptions in status messages
	// (e.g. "systemd-journald.service - Journal Logging Service").
	//
	// See systemd.unit for details about unit names and Description=.
	//
	// Added in version 243.
	StatusUnitFormat systemdconf.Value

	// Sets the default accuracy of timer units. This controls the global default for the AccuracySec= setting of timer units,
	// see systemd.timer for details. AccuracySec= set in individual units override the global default for the specific unit.
	// Defaults to 1min. Note that the accuracy of timer units is also affected by the configured timer slack for PID 1, see TimerSlackNSec=
	// above.
	//
	// Added in version 212.
	DefaultTimerAccuracySec systemdconf.Value

	// Sets the timer slack in nanoseconds for PID 1, which is inherited by all executed processes, unless overridden individually,
	// for example with the TimerSlackNSec= setting in service units (for details see systemd.exec). The timer slack controls
	// the accuracy of wake-ups triggered by system timers. See prctl for more information. Note that in contrast to most other
	// time span definitions this parameter takes an integer value in nano-seconds if no unit is specified. The usual time units
	// are understood too.
	//
	// Added in version 198.
	TimerSlackNSec systemdconf.Value

	// Configures the CPU affinity for the service manager as well as the default CPU affinity for all forked off processes. Takes
	// a list of CPU indices or ranges separated by either whitespace or commas. CPU ranges are specified by the lower and upper
	// CPU indices separated by a dash. This option may be specified more than once, in which case the specified CPU affinity masks
	// are merged. If the empty string is assigned, the mask is reset, all assignments prior to this will have no effect. Individual
	// services may override the CPU affinity for their processes with the CPUAffinity= setting in unit files, see systemd.exec.
	//
	// Added in version 198.
	CPUAffinity systemdconf.Value

	// Configures the NUMA memory policy for the service manager and the default NUMA memory policy for all forked off processes.
	// Individual services may override the default policy with the NUMAPolicy= setting in unit files, see systemd.exec.
	//
	// Added in version 243.
	NUMAPolicy systemdconf.Value

	// Configures the NUMA node mask that will be associated with the selected NUMA policy. Note that default and local NUMA policies
	// do not require explicit NUMA node mask and value of the option can be empty. Similarly to NUMAPolicy=, value can be overridden
	// by individual services in unit files, see systemd.exec.
	//
	// Added in version 243.
	NUMAMask systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, MemoryAccounting=,
	// TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details on the per-unit settings.
	//
	// DefaultCPUAccounting= defaults to yes when running on kernel ≥4.15, and no on older versions. DefaultMemoryAccounting=
	// defaults to yes. DefaultTasksAccounting= defaults to yes. The other settings default to no.
	//
	// Added in version 211.
	DefaultCPUAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, MemoryAccounting=,
	// TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details on the per-unit settings.
	//
	// DefaultCPUAccounting= defaults to yes when running on kernel ≥4.15, and no on older versions. DefaultMemoryAccounting=
	// defaults to yes. DefaultTasksAccounting= defaults to yes. The other settings default to no.
	//
	// Added in version 211.
	DefaultMemoryAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, MemoryAccounting=,
	// TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details on the per-unit settings.
	//
	// DefaultCPUAccounting= defaults to yes when running on kernel ≥4.15, and no on older versions. DefaultMemoryAccounting=
	// defaults to yes. DefaultTasksAccounting= defaults to yes. The other settings default to no.
	//
	// Added in version 211.
	DefaultTasksAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, MemoryAccounting=,
	// TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details on the per-unit settings.
	//
	// DefaultCPUAccounting= defaults to yes when running on kernel ≥4.15, and no on older versions. DefaultMemoryAccounting=
	// defaults to yes. DefaultTasksAccounting= defaults to yes. The other settings default to no.
	//
	// Added in version 211.
	DefaultIOAccounting systemdconf.Value

	// Configure the default resource accounting settings, as configured per-unit by CPUAccounting=, MemoryAccounting=,
	// TasksAccounting=, IOAccounting= and IPAccounting=. See systemd.resource-control for details on the per-unit settings.
	//
	// DefaultCPUAccounting= defaults to yes when running on kernel ≥4.15, and no on older versions. DefaultMemoryAccounting=
	// defaults to yes. DefaultTasksAccounting= defaults to yes. The other settings default to no.
	//
	// Added in version 211.
	DefaultIPAccounting systemdconf.Value

	// Configure the default value for the per-unit TasksMax= setting. See systemd.resource-control for details. This setting
	// applies to all unit types that support resource control settings, with the exception of slice units. Defaults to 15% of
	// the minimum of kernel.pid_max=, kernel.threads-max= and root cgroup pids.max. Kernel has a default value for kernel.pid_max=
	// and an algorithm of counting in case of more than 32 cores. For example, with the default kernel.pid_max=, DefaultTasksMax=
	// defaults to 4915, but might be greater in other systems or smaller in OS containers.
	//
	// Added in version 228.
	DefaultTasksMax systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitCPU systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitFSIZE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitDATA systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitSTACK systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitCORE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitRSS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitNOFILE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitAS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitNPROC systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitMEMLOCK systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitLOCKS systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitSIGPENDING systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitMSGQUEUE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitNICE systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitRTPRIO systemdconf.Value

	// These settings control various default resource limits for processes executed by units. See setrlimit for details. These
	// settings may be overridden in individual units using the corresponding LimitXXX= directives and they accept the same
	// parameter syntax, see systemd.exec for details. Note that these resource limits are only defaults for units, they are
	// not applied to the service manager process (i.e. PID 1) itself.
	//
	// Most of these settings are unset, which means the resource limits are inherited from the kernel or, if invoked in a container,
	// from the container manager. However, the following have defaults:
	//
	// DefaultLimitNOFILE= defaults to 1024:524288.
	//
	// DefaultLimitMEMLOCK= defaults to 8M.
	//
	// DefaultLimitCORE= does not have a default but it is worth mentioning that RLIMIT_CORE is set to "infinity" by PID 1 which
	// is inherited by its children.
	//
	// Note that the service manager internally in PID 1 bumps RLIMIT_NOFILE and RLIMIT_MEMLOCK to higher values, however the
	// limit is reverted to the mentioned defaults for all child processes forked off.
	//
	// Added in version 198.
	DefaultLimitRTTIME systemdconf.Value

	// Configure the default policy for reacting to processes being killed by the Linux Out-Of-Memory (OOM) killer or systemd-oomd.
	// This may be used to pick a global default for the per-unit OOMPolicy= setting. See systemd.service for details. Note that
	// this default is not used for services that have Delegate= turned on.
	//
	// Added in version 243.
	DefaultOOMPolicy systemdconf.Value

	// Configures the default OOM score adjustments of processes run by the service manager. This defaults to unset (meaning
	// the forked off processes inherit the service manager's OOM score adjustment value), except if the service manager is run
	// for an unprivileged user, in which case this defaults to the service manager's OOM adjustment value plus 100 (this makes
	// service processes slightly more likely to be killed under memory pressure than the manager itself). This may be used to
	// pick a global default for the per-unit OOMScoreAdjust= setting. See systemd.exec for details. Note that this setting
	// has no effect on the OOM score adjustment value of the service manager process itself, it retains the original value set
	// during its invocation.
	//
	// Added in version 250.
	DefaultOOMScoreAdjust systemdconf.Value

	// Configures the default settings for the per-unit MemoryPressureWatch= and MemoryPressureThresholdSec= settings.
	// See systemd.resource-control for details. Defaults to "auto" and "200ms", respectively. This also sets the memory pressure
	// monitoring threshold for the service manager itself.
	//
	// Added in version 254.
	DefaultMemoryPressureWatch systemdconf.Value

	// Configures the default settings for the per-unit MemoryPressureWatch= and MemoryPressureThresholdSec= settings.
	// See systemd.resource-control for details. Defaults to "auto" and "200ms", respectively. This also sets the memory pressure
	// monitoring threshold for the service manager itself.
	//
	// Added in version 254.
	DefaultMemoryPressureThresholdSec systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"), or the special strings "off" or "default". If set to "off" (alternatively: "0") the watchdog
	// logic is disabled: no watchdog device is opened, configured, or pinged. If set to the special string "default" the watchdog
	// is opened and pinged in regular intervals, but the timeout is not changed from the default. If set to any other time value
	// the watchdog timeout is configured to the specified value (or a value close to it, depending on hardware capabilities).
	//
	// If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog0 or the path specified with WatchdogDevice=
	// or the kernel option systemd.watchdog_device=) will be programmed to automatically reboot the system if it is not contacted
	// within the specified timeout interval. The system manager will ensure to contact it at least once in half the specified
	// timeout interval. This feature requires a hardware watchdog device to be present, as it is commonly the case in embedded
	// and server systems. Not all hardware watchdogs allow configuration of all possible reboot timeout values, in which case
	// the closest available timeout is picked.
	//
	// RebootWatchdogSec= may be used to configure the hardware watchdog when the system is asked to reboot. It works as a safety
	// net to ensure that the reboot takes place even if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout
	// applies only to the second phase of the reboot, i.e. after all regular services are already terminated, and after the system
	// and service manager process (PID 1) got replaced by the systemd-shutdown binary, see system bootup for details. During
	// the first phase of the shutdown operation the system and service manager remains running and hence RuntimeWatchdogSec=
	// is still honoured. In order to define a timeout on this first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction=
	// in the [Unit] section of the shutdown.target unit. By default, RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec=
	// to 10min.
	//
	// KExecWatchdogSec= may be used to additionally enable the watchdog when kexec is being executed rather than when rebooting.
	// Note that if the kernel does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this
	// case the watchdog might not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec=
	// is also enabled at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec=
	// is also enabled.
	//
	// These settings have no effect if a hardware watchdog is not available.
	//
	// Added in version 198.
	RuntimeWatchdogSec systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"), or the special strings "off" or "default". If set to "off" (alternatively: "0") the watchdog
	// logic is disabled: no watchdog device is opened, configured, or pinged. If set to the special string "default" the watchdog
	// is opened and pinged in regular intervals, but the timeout is not changed from the default. If set to any other time value
	// the watchdog timeout is configured to the specified value (or a value close to it, depending on hardware capabilities).
	//
	// If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog0 or the path specified with WatchdogDevice=
	// or the kernel option systemd.watchdog_device=) will be programmed to automatically reboot the system if it is not contacted
	// within the specified timeout interval. The system manager will ensure to contact it at least once in half the specified
	// timeout interval. This feature requires a hardware watchdog device to be present, as it is commonly the case in embedded
	// and server systems. Not all hardware watchdogs allow configuration of all possible reboot timeout values, in which case
	// the closest available timeout is picked.
	//
	// RebootWatchdogSec= may be used to configure the hardware watchdog when the system is asked to reboot. It works as a safety
	// net to ensure that the reboot takes place even if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout
	// applies only to the second phase of the reboot, i.e. after all regular services are already terminated, and after the system
	// and service manager process (PID 1) got replaced by the systemd-shutdown binary, see system bootup for details. During
	// the first phase of the shutdown operation the system and service manager remains running and hence RuntimeWatchdogSec=
	// is still honoured. In order to define a timeout on this first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction=
	// in the [Unit] section of the shutdown.target unit. By default, RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec=
	// to 10min.
	//
	// KExecWatchdogSec= may be used to additionally enable the watchdog when kexec is being executed rather than when rebooting.
	// Note that if the kernel does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this
	// case the watchdog might not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec=
	// is also enabled at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec=
	// is also enabled.
	//
	// These settings have no effect if a hardware watchdog is not available.
	//
	// Added in version 198.
	RebootWatchdogSec systemdconf.Value

	// Configure the hardware watchdog at runtime and at reboot. Takes a timeout value in seconds (or in other time units if suffixed
	// with "ms", "min", "h", "d", "w"), or the special strings "off" or "default". If set to "off" (alternatively: "0") the watchdog
	// logic is disabled: no watchdog device is opened, configured, or pinged. If set to the special string "default" the watchdog
	// is opened and pinged in regular intervals, but the timeout is not changed from the default. If set to any other time value
	// the watchdog timeout is configured to the specified value (or a value close to it, depending on hardware capabilities).
	//
	// If RuntimeWatchdogSec= is set to a non-zero value, the watchdog hardware (/dev/watchdog0 or the path specified with WatchdogDevice=
	// or the kernel option systemd.watchdog_device=) will be programmed to automatically reboot the system if it is not contacted
	// within the specified timeout interval. The system manager will ensure to contact it at least once in half the specified
	// timeout interval. This feature requires a hardware watchdog device to be present, as it is commonly the case in embedded
	// and server systems. Not all hardware watchdogs allow configuration of all possible reboot timeout values, in which case
	// the closest available timeout is picked.
	//
	// RebootWatchdogSec= may be used to configure the hardware watchdog when the system is asked to reboot. It works as a safety
	// net to ensure that the reboot takes place even if a clean reboot attempt times out. Note that the RebootWatchdogSec= timeout
	// applies only to the second phase of the reboot, i.e. after all regular services are already terminated, and after the system
	// and service manager process (PID 1) got replaced by the systemd-shutdown binary, see system bootup for details. During
	// the first phase of the shutdown operation the system and service manager remains running and hence RuntimeWatchdogSec=
	// is still honoured. In order to define a timeout on this first phase of system shutdown, configure JobTimeoutSec= and JobTimeoutAction=
	// in the [Unit] section of the shutdown.target unit. By default, RuntimeWatchdogSec= defaults to 0 (off), and RebootWatchdogSec=
	// to 10min.
	//
	// KExecWatchdogSec= may be used to additionally enable the watchdog when kexec is being executed rather than when rebooting.
	// Note that if the kernel does not reset the watchdog on kexec (depending on the specific hardware and/or driver), in this
	// case the watchdog might not get disabled after kexec succeeds and thus the system might get rebooted, unless RuntimeWatchdogSec=
	// is also enabled at the same time. For this reason it is recommended to enable KExecWatchdogSec= only if RuntimeWatchdogSec=
	// is also enabled.
	//
	// These settings have no effect if a hardware watchdog is not available.
	//
	// Added in version 198.
	KExecWatchdogSec systemdconf.Value

	// Configure the hardware watchdog device pre-timeout value. Takes a timeout value in seconds (or in other time units similar
	// to RuntimeWatchdogSec=). A watchdog pre-timeout is a notification generated by the watchdog before the watchdog reset
	// might occur in the event the watchdog has not been serviced. This notification is handled by the kernel and can be configured
	// to take an action (i.e. generate a kernel panic) using RuntimeWatchdogPreGovernor=. Not all watchdog hardware or drivers
	// support generating a pre-timeout and depending on the state of the system, the kernel may be unable to take the configured
	// action before the watchdog reboot. The watchdog will be configured to generate the pre-timeout event at the amount of time
	// specified by RuntimeWatchdogPreSec= before the runtime watchdog timeout (set by RuntimeWatchdogSec=). For example,
	// if the we have RuntimeWatchdogSec=30 and RuntimeWatchdogPreSec=10, then the pre-timeout event will occur if the watchdog
	// has not pinged for 20s (10s before the watchdog would fire). By default, RuntimeWatchdogPreSec= defaults to 0 (off). The
	// value set for RuntimeWatchdogPreSec= must be smaller than the timeout value for RuntimeWatchdogSec=. This setting has
	// no effect if a hardware watchdog is not available or the hardware watchdog does not support a pre-timeout and will be ignored
	// by the kernel if the setting is greater than the actual watchdog timeout.
	//
	// Added in version 251.
	RuntimeWatchdogPreSec systemdconf.Value

	// Configure the action taken by the hardware watchdog device when the pre-timeout expires. The default action for the pre-timeout
	// event depends on the kernel configuration, but it is usually to log a kernel message. For a list of valid actions available
	// for a given watchdog device, check the content of the /sys/class/watchdog/watchdogX/pretimeout_available_governors
	// file. Typically, available governor types are noop and panic. Availability, names and functionality might vary depending
	// on the specific device driver in use. If the pretimeout_available_governors sysfs file is empty, the governor might be
	// built as a kernel module and might need to be manually loaded (e.g. pretimeout_noop.ko), or the watchdog device might not
	// support pre-timeouts.
	//
	// Added in version 251.
	RuntimeWatchdogPreGovernor systemdconf.Value

	// Configure the hardware watchdog device that the runtime and shutdown watchdog timers will open and use. Defaults to /dev/watchdog0.
	// This setting has no effect if a hardware watchdog is not available.
	//
	// Added in version 236.
	WatchdogDevice systemdconf.Value

	// Controls which capabilities to include in the capability bounding set for PID 1 and its children. See capabilities for
	// details. Takes a whitespace-separated list of capability names as read by cap_from_name. Capabilities listed will be
	// included in the bounding set, all others are removed. If the list of capabilities is prefixed with ~, all but the listed capabilities
	// will be included, the effect of the assignment inverted. Note that this option also affects the respective capabilities
	// in the effective, permitted and inheritable capability sets. The capability bounding set may also be individually configured
	// for units using the CapabilityBoundingSet= directive for units, but note that capabilities dropped for PID 1 cannot be
	// regained in individual units, they are lost for good.
	//
	// Added in version 198.
	CapabilityBoundingSet systemdconf.Value

	// Takes a boolean argument. If true, ensures that PID 1 and all its children can never gain new privileges through execve (e.g.
	// via setuid or setgid bits, or filesystem capabilities). Defaults to false. General purpose distributions commonly rely
	// on executables with setuid or setgid bits and will thus not function properly with this option enabled. Individual units
	// cannot disable this option. Also see No New Privileges Flag.
	//
	// Added in version 239.
	NoNewPrivileges systemdconf.Value

	// Takes a boolean argument or the string "auto". If set to true this will remount /usr/ read-only. If set to "auto" (the default)
	// and running in an initrd equivalent to true, otherwise false. This implements a restricted subset of the per-unit setting
	// of the same name, see systemd.exec for details: currently, the "full" or "strict" values are not supported.
	//
	// Added in version 256.
	ProtectSystem systemdconf.Value

	// Takes a space-separated list of architecture identifiers. Selects from which architectures system calls may be invoked
	// on this system. This may be used as an effective way to disable invocation of non-native binaries system-wide, for example
	// to prohibit execution of 32-bit x86 binaries on 64-bit x86-64 systems. This option operates system-wide, and acts similar
	// to the SystemCallArchitectures= setting of unit files, see systemd.exec for details. This setting defaults to the empty
	// list, in which case no filtering of system calls based on architecture is applied. Known architecture identifiers are
	// "x86", "x86-64", "x32", "arm" and the special identifier "native". The latter implicitly maps to the native architecture
	// of the system (or more specifically, the architecture the system manager was compiled for). Set this setting to "native"
	// to prohibit execution of any non-native binaries. When a binary executes a system call of an architecture that is not listed
	// in this setting, it will be immediately terminated with the SIGSYS signal.
	//
	// Added in version 209.
	SystemCallArchitectures systemdconf.Value

	// Takes a SMACK64 security label as the argument. The process executed by a unit will be started under this label if SmackProcessLabel=
	// is not set in the unit. See systemd.exec for the details.
	//
	// If the value is "/", only labels specified with SmackProcessLabel= are assigned and the compile-time default is ignored.
	//
	// Added in version 252.
	DefaultSmackProcessLabel systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). For non-service units, DefaultTimeoutStartSec=
	// sets the default TimeoutSec= value.
	//
	// DefaultTimeoutStartSec= and DefaultTimeoutStopSec= default to 90 s in the system manager and 90 s in the user manager.
	// DefaultTimeoutAbortSec= is not set by default so that all units fall back to TimeoutStopSec=. DefaultRestartSec= defaults
	// to 100 ms.
	//
	// Added in version 209.
	DefaultTimeoutStartSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). For non-service units, DefaultTimeoutStartSec=
	// sets the default TimeoutSec= value.
	//
	// DefaultTimeoutStartSec= and DefaultTimeoutStopSec= default to 90 s in the system manager and 90 s in the user manager.
	// DefaultTimeoutAbortSec= is not set by default so that all units fall back to TimeoutStopSec=. DefaultRestartSec= defaults
	// to 100 ms.
	//
	// Added in version 209.
	DefaultTimeoutStopSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). For non-service units, DefaultTimeoutStartSec=
	// sets the default TimeoutSec= value.
	//
	// DefaultTimeoutStartSec= and DefaultTimeoutStopSec= default to 90 s in the system manager and 90 s in the user manager.
	// DefaultTimeoutAbortSec= is not set by default so that all units fall back to TimeoutStopSec=. DefaultRestartSec= defaults
	// to 100 ms.
	//
	// Added in version 209.
	DefaultTimeoutAbortSec systemdconf.Value

	// Configures the default timeouts for starting, stopping and aborting of units, as well as the default time to sleep between
	// automatic restarts of units, as configured per-unit in TimeoutStartSec=, TimeoutStopSec=, TimeoutAbortSec= and RestartSec=
	// (for services, see systemd.service for details on the per-unit settings). For non-service units, DefaultTimeoutStartSec=
	// sets the default TimeoutSec= value.
	//
	// DefaultTimeoutStartSec= and DefaultTimeoutStopSec= default to 90 s in the system manager and 90 s in the user manager.
	// DefaultTimeoutAbortSec= is not set by default so that all units fall back to TimeoutStopSec=. DefaultRestartSec= defaults
	// to 100 ms.
	//
	// Added in version 209.
	DefaultRestartSec systemdconf.Value

	// Configures the default timeout for waiting for devices. It can be changed per device via the x-systemd.device-timeout=
	// option in /etc/fstab and /etc/crypttab (see systemd.mount, crypttab). Defaults to 90 s in the system manager and 90 s in
	// the user manager.
	//
	// Added in version 252.
	DefaultDeviceTimeoutSec systemdconf.Value

	// Configure the default unit start rate limiting, as configured per-service by StartLimitIntervalSec= and StartLimitBurst=.
	// See systemd.service for details on the per-service settings. DefaultStartLimitIntervalSec= defaults to 10s. DefaultStartLimitBurst=
	// defaults to 5.
	//
	// Added in version 209.
	DefaultStartLimitIntervalSec systemdconf.Value

	// Configure the default unit start rate limiting, as configured per-service by StartLimitIntervalSec= and StartLimitBurst=.
	// See systemd.service for details on the per-service settings. DefaultStartLimitIntervalSec= defaults to 10s. DefaultStartLimitBurst=
	// defaults to 5.
	//
	// Added in version 209.
	DefaultStartLimitBurst systemdconf.Value

	// Rate limiting for daemon-reload and (since v256) daemon-reexec requests. The setting applies to both operations, but
	// the rate limits are tracked separately. Defaults to unset, and any number of operations can be requested at any time. ReloadLimitIntervalSec=
	// takes a value in seconds to configure the rate limit window, and ReloadLimitBurst= takes a positive integer to configure
	// the maximum allowed number of operations within the configured time window.
	//
	// Added in version 253.
	ReloadLimitIntervalSec systemdconf.Value

	// Rate limiting for daemon-reload and (since v256) daemon-reexec requests. The setting applies to both operations, but
	// the rate limits are tracked separately. Defaults to unset, and any number of operations can be requested at any time. ReloadLimitIntervalSec=
	// takes a value in seconds to configure the rate limit window, and ReloadLimitBurst= takes a positive integer to configure
	// the maximum allowed number of operations within the configured time window.
	//
	// Added in version 253.
	ReloadLimitBurst systemdconf.Value

	// Takes the same arguments as DefaultEnvironment=, see above. Sets environment variables for the manager process itself.
	// These variables are inherited by processes spawned by user managers, but not the system manager - use DefaultEnvironment=
	// for that. Note that these variables are merged into the existing environment block. In particular, in case of the system
	// manager, this includes variables set by the kernel based on the kernel command line. As with DefaultEnvironment=, this
	// environment block is internal, and changes are not reflected in the manager's /proc/PID/environ.
	//
	// Setting environment variables for the manager process may be useful to modify its behaviour. See Known Environment Variables
	// for a descriptions of some variables understood by systemd.
	//
	// Simple "%"-specifier expansion is supported, see below for a list of supported specifiers.
	//
	// Added in version 248.
	ManagerEnvironment systemdconf.Value

	// Configures environment variables passed to all executed processes. Takes a space-separated list of variable assignments.
	// See environ for details about environment variables.
	//
	// Simple "%"-specifier expansion is supported, see below for a list of supported specifiers.
	//
	// Example:
	//
	//	DefaultEnvironment="VAR1=word1 word2" VAR2=word3 "VAR3=word 5 6"
	//
	// Sets three variables "VAR1", "VAR2", "VAR3".
	//
	// Added in version 205.
	DefaultEnvironment systemdconf.Value
}
