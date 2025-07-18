// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// SysupdateFeaturesFile represents sysupdate.features — Definition Files for Optional Features
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.features.html for details).
type SysupdateFeaturesFile struct {
	systemdconf.File

	Feature SysupdateFeaturesFeatureSection // general properties of this feature
}

// SysupdateFeaturesFeatureSection represents general properties of this feature
// (see https://www.freedesktop.org/software/systemd/man/sysupdate.features.html#%5BFeature%5D%20Section%20Options for details).
type SysupdateFeaturesFeatureSection struct {
	systemdconf.Section

	// A short human readable description of this feature. This may be used as a label for this feature, so the string should meaningfully
	// identify the feature among the features available in sysupdate.d/.
	//
	// Added in version 257.
	Description systemdconf.Value

	// A user-presentable URL to documentation about this feature. This setting supports specifier expansion; see below for
	// details on supported specifiers.
	//
	// Added in version 257.
	Documentation systemdconf.Value

	// A URL to an AppStream catalog XML file. This may be used by software centers (such as GNOME Software or KDE Discover) to present
	// rich metadata about this feature. This includes display names, chagnelogs, icons, and more. This setting supports specifier
	// expansion; see below for details on supported specifiers.
	//
	// Added in version 257.
	AppStream systemdconf.Value

	// Whether or not this feature is enabled. If unspecified, the feature is disabled by default.
	//
	// Added in version 257.
	Enabled systemdconf.Value
}
