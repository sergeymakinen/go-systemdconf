// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// IocostFile represents iocost.conf — Configuration files for the iocost solution manager
// (see https://www.freedesktop.org/software/systemd/man/iocost.conf.html for details).
type IocostFile struct {
	systemdconf.File

	IOCost IocostIOCostSection // behavior of "iocost", a tool mostly used by systemd-udevd rules to automatically apply I/O cost solutions to /sys/fs/cgroup/io.cost.*
}

// IocostIOCostSection represents behavior of "iocost", a tool mostly used by systemd-udevd rules to automatically apply I/O cost solutions to /sys/fs/cgroup/io.cost.*
// (see https://www.freedesktop.org/software/systemd/man/iocost.conf.html#Options for details).
type IocostIOCostSection struct {
	systemdconf.Section

	// Chooses which I/O cost solution (identified by named string) should be used for the devices in this system. The known solutions
	// can be queried from the udev metadata attached to the devices. If a device does not have the specified solution, the first
	// one listed in IOCOST_SOLUTIONS is used instead.
	//
	// E.g. "TargetSolution=isolated-bandwidth".
	//
	// Added in version 254.
	TargetSolution systemdconf.Value
}
