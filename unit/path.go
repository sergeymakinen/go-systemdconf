// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// PathSection represents information about the path(s) it monitors
type PathSection struct {
	systemdconf.Section

	// Defines paths to monitor for certain changes: PathExists= may be used to watch the mere existence of a file or directory.
	// If the file specified exists, the configured unit is activated. PathExistsGlob= works similar, but checks for the existence
	// of at least one file matching the globbing pattern specified. PathChanged= may be used to watch a file or directory and activate
	// the configured unit whenever it changes. It is not activated on every write to the watched file but it is activated if the
	// file which was open for writing gets closed. PathModified= is similar, but additionally it is activated also on simple
	// writes to the watched file. DirectoryNotEmpty= may be used to watch a directory and activate the configured unit whenever
	// it contains at least one file.
	//
	// The arguments of these directives must be absolute file system paths.
	//
	// Multiple directives may be combined, of the same and of different types, to watch multiple paths. If the empty string is
	// assigned to any of these options, the list of paths to watch is reset, and any prior assignments of these options will not
	// have any effect.
	//
	// If a path already exists (in case of PathExists= and PathExistsGlob=) or a directory already is not empty (in case of DirectoryNotEmpty=)
	// at the time the path unit is activated, then the configured unit is immediately activated as well. Something similar does
	// not apply to PathChanged= and PathModified=.
	//
	// If the path itself or any of the containing directories are not accessible, systemd will watch for permission changes and
	// notice that conditions are satisfied when permissions allow that.
	PathExists systemdconf.Value

	// Defines paths to monitor for certain changes: PathExists= may be used to watch the mere existence of a file or directory.
	// If the file specified exists, the configured unit is activated. PathExistsGlob= works similar, but checks for the existence
	// of at least one file matching the globbing pattern specified. PathChanged= may be used to watch a file or directory and activate
	// the configured unit whenever it changes. It is not activated on every write to the watched file but it is activated if the
	// file which was open for writing gets closed. PathModified= is similar, but additionally it is activated also on simple
	// writes to the watched file. DirectoryNotEmpty= may be used to watch a directory and activate the configured unit whenever
	// it contains at least one file.
	//
	// The arguments of these directives must be absolute file system paths.
	//
	// Multiple directives may be combined, of the same and of different types, to watch multiple paths. If the empty string is
	// assigned to any of these options, the list of paths to watch is reset, and any prior assignments of these options will not
	// have any effect.
	//
	// If a path already exists (in case of PathExists= and PathExistsGlob=) or a directory already is not empty (in case of DirectoryNotEmpty=)
	// at the time the path unit is activated, then the configured unit is immediately activated as well. Something similar does
	// not apply to PathChanged= and PathModified=.
	//
	// If the path itself or any of the containing directories are not accessible, systemd will watch for permission changes and
	// notice that conditions are satisfied when permissions allow that.
	PathExistsGlob systemdconf.Value

	// Defines paths to monitor for certain changes: PathExists= may be used to watch the mere existence of a file or directory.
	// If the file specified exists, the configured unit is activated. PathExistsGlob= works similar, but checks for the existence
	// of at least one file matching the globbing pattern specified. PathChanged= may be used to watch a file or directory and activate
	// the configured unit whenever it changes. It is not activated on every write to the watched file but it is activated if the
	// file which was open for writing gets closed. PathModified= is similar, but additionally it is activated also on simple
	// writes to the watched file. DirectoryNotEmpty= may be used to watch a directory and activate the configured unit whenever
	// it contains at least one file.
	//
	// The arguments of these directives must be absolute file system paths.
	//
	// Multiple directives may be combined, of the same and of different types, to watch multiple paths. If the empty string is
	// assigned to any of these options, the list of paths to watch is reset, and any prior assignments of these options will not
	// have any effect.
	//
	// If a path already exists (in case of PathExists= and PathExistsGlob=) or a directory already is not empty (in case of DirectoryNotEmpty=)
	// at the time the path unit is activated, then the configured unit is immediately activated as well. Something similar does
	// not apply to PathChanged= and PathModified=.
	//
	// If the path itself or any of the containing directories are not accessible, systemd will watch for permission changes and
	// notice that conditions are satisfied when permissions allow that.
	PathChanged systemdconf.Value

	// Defines paths to monitor for certain changes: PathExists= may be used to watch the mere existence of a file or directory.
	// If the file specified exists, the configured unit is activated. PathExistsGlob= works similar, but checks for the existence
	// of at least one file matching the globbing pattern specified. PathChanged= may be used to watch a file or directory and activate
	// the configured unit whenever it changes. It is not activated on every write to the watched file but it is activated if the
	// file which was open for writing gets closed. PathModified= is similar, but additionally it is activated also on simple
	// writes to the watched file. DirectoryNotEmpty= may be used to watch a directory and activate the configured unit whenever
	// it contains at least one file.
	//
	// The arguments of these directives must be absolute file system paths.
	//
	// Multiple directives may be combined, of the same and of different types, to watch multiple paths. If the empty string is
	// assigned to any of these options, the list of paths to watch is reset, and any prior assignments of these options will not
	// have any effect.
	//
	// If a path already exists (in case of PathExists= and PathExistsGlob=) or a directory already is not empty (in case of DirectoryNotEmpty=)
	// at the time the path unit is activated, then the configured unit is immediately activated as well. Something similar does
	// not apply to PathChanged= and PathModified=.
	//
	// If the path itself or any of the containing directories are not accessible, systemd will watch for permission changes and
	// notice that conditions are satisfied when permissions allow that.
	PathModified systemdconf.Value

	// Defines paths to monitor for certain changes: PathExists= may be used to watch the mere existence of a file or directory.
	// If the file specified exists, the configured unit is activated. PathExistsGlob= works similar, but checks for the existence
	// of at least one file matching the globbing pattern specified. PathChanged= may be used to watch a file or directory and activate
	// the configured unit whenever it changes. It is not activated on every write to the watched file but it is activated if the
	// file which was open for writing gets closed. PathModified= is similar, but additionally it is activated also on simple
	// writes to the watched file. DirectoryNotEmpty= may be used to watch a directory and activate the configured unit whenever
	// it contains at least one file.
	//
	// The arguments of these directives must be absolute file system paths.
	//
	// Multiple directives may be combined, of the same and of different types, to watch multiple paths. If the empty string is
	// assigned to any of these options, the list of paths to watch is reset, and any prior assignments of these options will not
	// have any effect.
	//
	// If a path already exists (in case of PathExists= and PathExistsGlob=) or a directory already is not empty (in case of DirectoryNotEmpty=)
	// at the time the path unit is activated, then the configured unit is immediately activated as well. Something similar does
	// not apply to PathChanged= and PathModified=.
	//
	// If the path itself or any of the containing directories are not accessible, systemd will watch for permission changes and
	// notice that conditions are satisfied when permissions allow that.
	DirectoryNotEmpty systemdconf.Value

	// The unit to activate when any of the configured paths changes. The argument is a unit name, whose suffix is not ".path". If
	// not specified, this value defaults to a service that has the same name as the path unit, except for the suffix. (See above.)
	// It is recommended that the unit name that is activated and the unit name of the path unit are named identical, except for the
	// suffix.
	Unit systemdconf.Value

	// Takes a boolean argument. If true, the directories to watch are created before watching. This option is ignored for PathExists=
	// settings. Defaults to false.
	MakeDirectory systemdconf.Value

	// If MakeDirectory= is enabled, use the mode specified here to create the directories in question. Takes an access mode in
	// octal notation. Defaults to 0755.
	DirectoryMode systemdconf.Value
}

// PathFile represents information about a path monitored by systemd, for path-based activation
type PathFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Path    PathSection    // Information about the path(s) it monitors
	Install InstallSection // Installation information for the unit
}
