// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// SleepFile represents systemd-sleep.conf, sleep.conf.d — Suspend and hibernation configuration file
// (see https://www.freedesktop.org/software/systemd/man/systemd-sleep.conf.html for details).
type SleepFile struct {
	systemdconf.File

	Sleep SleepSleepSection // [Sleep] section
}

// SleepSleepSection represents [Sleep] section
// (see https://www.freedesktop.org/software/systemd/man/systemd-sleep.conf.html#Options for details).
type SleepSleepSection struct {
	systemdconf.Section

	// By default, any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	//
	// Added in version 240.
	AllowSuspend systemdconf.Value

	// By default, any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	//
	// Added in version 240.
	AllowHibernation systemdconf.Value

	// By default, any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	//
	// Added in version 240.
	AllowHybridSleep systemdconf.Value

	// By default, any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	//
	// Added in version 240.
	AllowSuspendThenHibernate systemdconf.Value

	// The string to be written to /sys/power/state by systemd-suspend.service. More than one value can be specified by separating
	// multiple values with whitespace. They will be tried in turn, until one is written without error. If none of the writes succeed,
	// the operation will be aborted.
	//
	// The allowed set of values is determined by the kernel and is shown in the file itself (use cat /sys/power/state to display).
	// See  Basic sysfs Interfaces for System Suspend and Hibernation for more details.
	//
	// systemd-suspend-then-hibernate.service uses this value when suspending.
	//
	// Added in version 203.
	SuspendState systemdconf.Value

	// The string to be written to /sys/power/disk by systemd-hibernate.service. More than one value can be specified by separating
	// multiple values with whitespace. They will be tried in turn, until one is written without error. If none of the writes succeed,
	// the operation will be aborted.
	//
	// The allowed set of values is determined by the kernel and is shown in the file itself (use cat /sys/power/disk to display).
	// See the kernel documentation page  Basic sysfs Interfaces for System Suspend and Hibernation for more details.
	//
	// systemd-suspend-then-hibernate.service uses the value of HibernateMode= when hibernating.
	//
	// Added in version 203.
	HibernateMode systemdconf.Value

	// The string to be written to /sys/power/mem_sleep when SuspendState=mem or hybrid-sleep is used. More than one value can
	// be specified by separating multiple values with whitespace. They will be tried in turn, until one is written without error.
	// If none of the writes succeed, the operation will be aborted. Defaults to empty, i.e. the kernel default or kernel command
	// line option mem_sleep_default= is respected.
	//
	// The allowed set of values is determined by the kernel and is shown in the file itself (use cat /sys/power/mem_sleep to display).
	// See the kernel documentation page  Basic sysfs Interfaces for System Suspend and Hibernation for more details.
	//
	// Added in version 256.
	MemorySleepMode systemdconf.Value

	// The amount of time the system spends in suspend mode before the system is automatically put into hibernate mode. Only used
	// by systemd-suspend-then-hibernate.service. Refer to suspend-then-hibernate for details on how this option interacts
	// with other options/system battery state.
	//
	// Added in version 239.
	HibernateDelaySec systemdconf.Value

	// Whether to allow hibernation when the system has AC power. Only used by systemd-suspend-then-hibernate.service when
	// HibernateDelaySec= is set.
	//
	// If this option is disabled, the countdown of HibernateDelaySec= starts only after AC power is disconnected, keeping the
	// system in the suspend state otherwise.
	//
	// This option is only effective on systems with a battery.
	//
	// Added in version 257.
	HibernateOnACPower systemdconf.Value

	// The RTC alarm will wake the system after the specified timespan to measure the system battery capacity level and estimate
	// battery discharging rate. Only used by systemd-suspend-then-hibernate.service. Refer to suspend-then-hibernate
	// for details on how this option interacts with other options/system battery state.
	//
	// Added in version 253.
	SuspendEstimationSec systemdconf.Value
}
