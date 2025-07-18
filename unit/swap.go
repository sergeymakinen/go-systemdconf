// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package unit

import "github.com/sergeymakinen/go-systemdconf/v3"

// SwapFile represents systemd.swap — Swap unit configuration
// (see https://www.freedesktop.org/software/systemd/man/systemd.swap.html for details).
type SwapFile struct {
	systemdconf.File

	Unit    UnitSection    // generic information about the unit that is not dependent on the type of unit
	Swap    SwapSection    // information about the swap device it supervises
	Install InstallSection // installation information for the unit
}

// SwapSection represents information about the swap device it supervises
// (see https://www.freedesktop.org/software/systemd/man/systemd.swap.html#Options for details).
type SwapSection struct {
	systemdconf.Section
	ExecOptions
	KillOptions
	ResourceControlOptions

	// Takes an absolute path or a fstab-style identifier of a device node or file to use for paging. See swapon for details. If this
	// refers to a device node, a dependency on the respective device unit is automatically created. (See systemd.device for
	// more information.) If this refers to a file, a dependency on the respective mount unit is automatically created. (See systemd.mount
	// for more information.) This option is mandatory. Note that the usual specifier expansion is applied to this setting, literal
	// percent characters should hence be written as "%%".
	What systemdconf.Value

	// Swap priority to use when activating the swap device or file. This takes an integer. This setting is optional and ignored
	// when the priority is set by pri= in the Options= key.
	Priority systemdconf.Value

	// May contain an option string for the swap device. This may be used for controlling discard options among other functionality,
	// if the swap backing device supports the discard or trim operation. (See swapon for more information.) Note that the usual
	// specifier expansion is applied to this setting, literal percent characters should hence be written as "%%".
	//
	// Added in version 217.
	Options systemdconf.Value

	// Configures the time to wait for the swapon command to finish. If a command does not exit within the configured time, the swap
	// will be considered failed and be shut down again. All commands still running will be terminated forcibly via SIGTERM, and
	// after another delay of this time with SIGKILL. (See KillMode= in systemd.kill.) Takes a unit-less value in seconds, or
	// a time span value such as "5min 20s". Pass "0" to disable the timeout logic. Defaults to DefaultTimeoutStartSec= from the
	// manager configuration file (see systemd-system.conf).
	TimeoutSec systemdconf.Value
}
