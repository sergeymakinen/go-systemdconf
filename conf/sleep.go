// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf"

// SleepFile represents systemd-sleep.conf, sleep.conf.d — Suspend and hibernation configuration file
// (see https://www.freedesktop.org/software/systemd/man/systemd-sleep.conf.html for details)
type SleepFile struct {
	systemdconf.File

	Sleep SleepSleepSection // What strings will be written to /sys/power/disk and /sys/power/state by systemd-sleep when systemd attempts to suspend or hibernate the machine
}

// SleepSleepSection represents what strings will be written to /sys/power/disk and /sys/power/state by systemd-sleep when systemd attempts to suspend or hibernate the machine
// (see https://www.freedesktop.org/software/systemd/man/systemd-sleep.conf.html#Options for details)
type SleepSleepSection struct {
	systemdconf.Section

	// By default any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	AllowSuspend systemdconf.Value

	// By default any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	AllowHibernation systemdconf.Value

	// By default any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	AllowSuspendThenHibernate systemdconf.Value

	// By default any power-saving mode is advertised if possible (i.e. the kernel supports that mode, the necessary resources
	// are available). Those switches can be used to disable specific modes.
	//
	// If AllowHibernation=no or AllowSuspend=no is used, this implies AllowSuspendThenHibernate=no and AllowHybridSleep=no,
	// since those methods use both suspend and hibernation internally. AllowSuspendThenHibernate=yes and AllowHybridSleep=yes
	// can be used to override and enable those specific modes.
	AllowHybridSleep systemdconf.Value

	// The string to be written to /sys/power/disk by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	SuspendMode systemdconf.Value

	// The string to be written to /sys/power/disk by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	HibernateMode systemdconf.Value

	// The string to be written to /sys/power/disk by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	HybridSleepMode systemdconf.Value

	// The string to be written to /sys/power/state by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	SuspendState systemdconf.Value

	// The string to be written to /sys/power/state by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	HibernateState systemdconf.Value

	// The string to be written to /sys/power/state by, respectively, systemd-suspend.service, systemd-hibernate.service,
	// systemd-hybrid-sleep.service, or systemd-suspend-then-hibernate.service. More than one value can be specified
	// by separating multiple values with whitespace. They will be tried in turn, until one is written without error. If neither
	// succeeds, the operation will be aborted.
	HybridSleepState systemdconf.Value

	// The amount of time the system spends in suspend mode before the system is automatically put into hibernate mode, when using
	// systemd-suspend-then-hibernate.service. Defaults to 2h.
	HibernateDelaySec systemdconf.Value
}
