// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// ScopeFile represents systemd.scope — Scope unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.scope.html for details).
type ScopeFile struct {
	systemdconf.File

	Unit  UnitSection  // generic information about the unit that is not dependent on the type of unit
	Scope ScopeSection // information about the scope and the units it contains
}

// ScopeSection represents information about the scope and the units it contains
// (see https://www.freedesktop.org/software/systemd/man/systemd.scope.html#Options for details).
type ScopeSection struct {
	systemdconf.Section
	KillOptions
	ResourceControlOptions

	// Configure the out-of-memory (OOM) killing policy for the kernel and the userspace OOM killer systemd-oomd.service.
	// On Linux, when memory becomes scarce to the point that the kernel has trouble allocating memory for itself, it might decide
	// to kill a running process in order to free up memory and reduce memory pressure. Note that systemd-oomd.service is a more
	// flexible solution that aims to prevent out-of-memory situations for the userspace too, not just the kernel, by attempting
	// to terminate services earlier, before the kernel would have to act.
	//
	// This setting takes one of continue, stop or kill. If set to continue and a process in the unit is killed by the OOM killer, this
	// is logged but the unit continues running. If set to stop the event is logged but the unit is terminated cleanly by the service
	// manager. If set to kill and one of the unit's processes is killed by the OOM killer the kernel is instructed to kill all remaining
	// processes of the unit too, by setting the memory.oom.group attribute to 1; also see kernel page Control Group v2.
	//
	// Defaults to the setting DefaultOOMPolicy= in systemd-system.conf is set to, except for units where Delegate= is turned
	// on, where it defaults to continue.
	//
	// Use the OOMScoreAdjust= setting to configure whether processes of the unit shall be considered preferred or less preferred
	// candidates for process termination by the Linux OOM killer logic. See systemd.exec for details.
	//
	// This setting also applies to systemd-oomd.service. Similarly to the kernel OOM kills performed by the kernel, this setting
	// determines the state of the unit after systemd-oomd kills a cgroup associated with it.
	//
	// Added in version 253.
	OOMPolicy systemdconf.Value

	// Configures a maximum time for the scope to run. If this is used and the scope has been active for longer than the specified
	// time it is terminated and put into a failure state. Pass "infinity" (the default) to configure no runtime limit.
	//
	// Added in version 244.
	RuntimeMaxSec systemdconf.Value

	// This option modifies RuntimeMaxSec= by increasing the maximum runtime by an evenly distributed duration between 0 and
	// the specified value (in seconds). If RuntimeMaxSec= is unspecified, then this feature will be disabled.
	//
	// Added in version 250.
	RuntimeRandomizedExtraSec systemdconf.Value
}
