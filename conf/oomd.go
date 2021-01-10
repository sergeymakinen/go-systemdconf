// DO NOT EDIT. This file is generated from systemd 247 by generatesdconf

package conf

import "github.com/sergeymakinen/go-systemdconf/v2"

// OomdFile represents oomd.conf, oomd.conf.d — Global systemd-oomd configuration files
// (see https://www.freedesktop.org/software/systemd/man/oomd.conf.html for details)
type OomdFile struct {
	systemdconf.File

	OOM OomdOOMSection // Parameters of the systemd userspace out-of-memory (OOM) killer, systemd-oomd.service
}

// OomdOOMSection represents parameters of the systemd userspace out-of-memory (OOM) killer, systemd-oomd.service
// (see https://www.freedesktop.org/software/systemd/man/oomd.conf.html#%5BOOM%5D%20Section%20Options for details)
type OomdOOMSection struct {
	systemdconf.Section

	// Sets the limit for swap usage on the system before systemd-oomd will take action. If the percentage of swap used on the system
	// is more than what is defined here, systemd-oomd will act on eligible descendant cgroups, starting from the ones with the
	// highest swap usage to the lowest swap usage. Which cgroups are monitored and what action gets taken depends on what the unit
	// has configured for ManagedOOMSwap=. Takes a percentage value between 0% and 100%, inclusive. Defaults to 90%.
	SwapUsedLimitPercent systemdconf.Value

	// Sets the limit for memory pressure on the unit's cgroup before systemd-oomd will take action. A unit can override this value
	// with ManagedOOMMemoryPressureLimitPercent=. The memory pressure for this property represents the fraction of time
	// in a 10 second window in which all tasks in the cgroup were delayed. For each monitored cgroup, if the memory pressure on that
	// cgroup exceeds the limit set for more than 30 seconds, systemd-oomd will act on eligible descendant cgroups, starting
	// from the ones with the most reclaim activity to the least reclaim activity. Which cgroups are monitored and what action
	// gets taken depends on what the unit has configured for ManagedOOMMemoryPressure=. Takes a percentage value between 0%
	// and 100%, inclusive. Defaults to 60%.
	DefaultMemoryPressureLimitPercent systemdconf.Value
}
