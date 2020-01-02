// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// SwapSection represents information about the swap device it supervises
type SwapSection struct {
	systemdconf.Section
	ExecOptions
	KillOptions
	ResourceControlOptions

	// Takes an absolute path of a device node or file to use for paging. See swapon for details. If this refers to a device node, a
	// dependency on the respective device unit is automatically created. (See systemd.device for more information.) If this
	// refers to a file, a dependency on the respective mount unit is automatically created. (See systemd.mount for more information.)
	// This option is mandatory. Note that the usual specifier expansion is applied to this setting, literal percent characters
	// should hence be written as "%%".
	What systemdconf.Value

	// Swap priority to use when activating the swap device or file. This takes an integer. This setting is optional and ignored
	// when the priority is set by pri= in the Options= key.
	Priority systemdconf.Value

	// May contain an option string for the swap device. This may be used for controlling discard options among other functionality,
	// if the swap backing device supports the discard or trim operation. (See swapon for more information.) Note that the usual
	// specifier expansion is applied to this setting, literal percent characters should hence be written as "%%".
	Options systemdconf.Value

	// Configures the time to wait for the swapon command to finish. If a command does not exit within the configured time, the swap
	// will be considered failed and be shut down again. All commands still running will be terminated forcibly via SIGTERM, and
	// after another delay of this time with SIGKILL. (See KillMode= in systemd.kill.) Takes a unit-less value in seconds, or
	// a time span value such as "5min 20s". Pass "0" to disable the timeout logic. Defaults to DefaultTimeoutStartSec= from the
	// manager configuration file (see systemd-system.conf).
	TimeoutSec systemdconf.Value
}

// SwapFile represents information about a swap device or file for memory paging controlled and supervised by systemd
type SwapFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Swap    SwapSection    // Information about the swap device it supervises
	Install InstallSection // Installation information for the unit
}
