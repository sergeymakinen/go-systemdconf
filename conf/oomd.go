// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// OomdFile represents oomd.conf, oomd.conf.d — Global systemd-oomd configuration files
// (see https://www.freedesktop.org/software/systemd/man/oomd.conf.html for details).
type OomdFile struct {
	systemdconf.File

	OOM OomdOOMSection // parameters of the systemd userspace out-of-memory (OOM) killer, systemd-oomd.service
}

// OomdOOMSection represents parameters of the systemd userspace out-of-memory (OOM) killer, systemd-oomd.service
// (see https://www.freedesktop.org/software/systemd/man/oomd.conf.html#%5BOOM%5D%20Section%20Options for details).
type OomdOOMSection struct {
	systemdconf.Section

	// Sets the limit for memory and swap usage on the system before systemd-oomd will take action. If the fraction of memory used
	// and the fraction of swap used on the system are both more than what is defined here, systemd-oomd will act on eligible descendant
	// control groups with swap usage greater than 5% of total swap, starting from the ones with the highest swap usage. Which control
	// groups are monitored and what action gets taken depends on what the unit has configured for ManagedOOMSwap=. Takes a value
	// specified in percent (when suffixed with "%"), permille ("‰") or permyriad ("‱"), between 0% and 100%, inclusive. Defaults
	// to 90%.
	//
	// Added in version 247.
	SwapUsedLimit systemdconf.Value

	// Sets the limit for memory pressure on the unit's control group before systemd-oomd will take action. A unit can override
	// this value with ManagedOOMMemoryPressureLimit=. The memory pressure for this property represents the fraction of time
	// in a 10 second window in which all tasks in the control group were delayed. For each monitored control group, if the memory
	// pressure on that control group exceeds the limit set for longer than the duration set by DefaultMemoryPressureDurationSec=,
	// systemd-oomd will act on eligible descendant control groups, starting from the ones with the most reclaim activity to
	// the least reclaim activity. Which control groups are monitored and what action gets taken depends on what the unit has configured
	// for ManagedOOMMemoryPressure=. Takes a fraction specified in the same way as SwapUsedLimit= above. Defaults to 60%.
	//
	// Added in version 247.
	DefaultMemoryPressureLimit systemdconf.Value

	// Sets the amount of time a unit's control group needs to have exceeded memory pressure limits before systemd-oomd will take
	// action. A unit can override this value with ManagedOOMMemoryPressureDurationSec=. Memory pressure limits are defined
	// by DefaultMemoryPressureLimit= and ManagedOOMMemoryPressureLimit=. Must be set to 0, or at least 1 second. Defaults
	// to 30 seconds when unset or 0.
	//
	// Added in version 248.
	DefaultMemoryPressureDurationSec systemdconf.Value
}
