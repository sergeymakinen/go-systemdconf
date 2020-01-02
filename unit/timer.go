// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// TimerSection represents information about the timer it defines
type TimerSection struct {
	systemdconf.Section

	// Defines monotonic timers relative to different starting points:
	//
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	|      SETTING       |                                     MEANING                                      |
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	| OnActiveSec=       | Defines a timer relative to the moment the timer unit itself is activated.       |
	// 	| OnBootSec=         | Defines a timer relative to when the machine was booted up. In containers,       |
	// 	|                    | for the system manager instance, this is mapped to OnStartupSec=, making both    |
	// 	|                    | equivalent.                                                                      |
	// 	| OnStartupSec=      | Defines a timer relative to when the service manager was first started. For      |
	// 	|                    | system timer units this is very similar to OnBootSec= as the system service      |
	// 	|                    | manager is generally started very early at boot. It's primarily useful when      |
	// 	|                    | configured in units running in the per-user service manager, as the user service |
	// 	|                    | manager is generally started on first login only, not already during boot.       |
	// 	| OnUnitActiveSec=   | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | activated.                                                                       |
	// 	| OnUnitInactiveSec= | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | deactivated.                                                                     |
	// 	+--------------------+----------------------------------------------------------------------------------+
	//
	// Multiple directives may be combined of the same and of different types, in which case the timer unit will trigger whenever
	// any of the specified timer expressions elapse. For example, by combining OnBootSec= and OnUnitActiveSec=, it is possible
	// to define a timer that elapses in regular intervals and activates a specific service each time. Moreover, both monotonic
	// time expressions and OnCalendar= calendar expressions may be combined in the same timer unit.
	//
	// The arguments to the directives are time spans configured in seconds. Example: "OnBootSec=50" means 50s after boot-up.
	// The argument may also include time units. Example: "OnBootSec=5h 30min" means 5 hours and 30 minutes after boot-up. For
	// details about the syntax of time spans, see systemd.time.
	//
	// If a timer configured with OnBootSec= or OnStartupSec= is already in the past when the timer unit is activated, it will immediately
	// elapse and the configured unit is started. This is not the case for timers defined in the other directives.
	//
	// These are monotonic timers, independent of wall-clock time and timezones. If the computer is temporarily suspended,
	// the monotonic clock pauses, too.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both monotonic timers and OnCalendar=
	// timers, see below), and all prior assignments will have no effect.
	//
	// Note that timers do not necessarily expire at the precise time configured with these settings, as they are subject to the
	// AccuracySec= setting below.
	OnActiveSec systemdconf.Value

	// Defines monotonic timers relative to different starting points:
	//
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	|      SETTING       |                                     MEANING                                      |
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	| OnActiveSec=       | Defines a timer relative to the moment the timer unit itself is activated.       |
	// 	| OnBootSec=         | Defines a timer relative to when the machine was booted up. In containers,       |
	// 	|                    | for the system manager instance, this is mapped to OnStartupSec=, making both    |
	// 	|                    | equivalent.                                                                      |
	// 	| OnStartupSec=      | Defines a timer relative to when the service manager was first started. For      |
	// 	|                    | system timer units this is very similar to OnBootSec= as the system service      |
	// 	|                    | manager is generally started very early at boot. It's primarily useful when      |
	// 	|                    | configured in units running in the per-user service manager, as the user service |
	// 	|                    | manager is generally started on first login only, not already during boot.       |
	// 	| OnUnitActiveSec=   | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | activated.                                                                       |
	// 	| OnUnitInactiveSec= | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | deactivated.                                                                     |
	// 	+--------------------+----------------------------------------------------------------------------------+
	//
	// Multiple directives may be combined of the same and of different types, in which case the timer unit will trigger whenever
	// any of the specified timer expressions elapse. For example, by combining OnBootSec= and OnUnitActiveSec=, it is possible
	// to define a timer that elapses in regular intervals and activates a specific service each time. Moreover, both monotonic
	// time expressions and OnCalendar= calendar expressions may be combined in the same timer unit.
	//
	// The arguments to the directives are time spans configured in seconds. Example: "OnBootSec=50" means 50s after boot-up.
	// The argument may also include time units. Example: "OnBootSec=5h 30min" means 5 hours and 30 minutes after boot-up. For
	// details about the syntax of time spans, see systemd.time.
	//
	// If a timer configured with OnBootSec= or OnStartupSec= is already in the past when the timer unit is activated, it will immediately
	// elapse and the configured unit is started. This is not the case for timers defined in the other directives.
	//
	// These are monotonic timers, independent of wall-clock time and timezones. If the computer is temporarily suspended,
	// the monotonic clock pauses, too.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both monotonic timers and OnCalendar=
	// timers, see below), and all prior assignments will have no effect.
	//
	// Note that timers do not necessarily expire at the precise time configured with these settings, as they are subject to the
	// AccuracySec= setting below.
	OnBootSec systemdconf.Value

	// Defines monotonic timers relative to different starting points:
	//
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	|      SETTING       |                                     MEANING                                      |
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	| OnActiveSec=       | Defines a timer relative to the moment the timer unit itself is activated.       |
	// 	| OnBootSec=         | Defines a timer relative to when the machine was booted up. In containers,       |
	// 	|                    | for the system manager instance, this is mapped to OnStartupSec=, making both    |
	// 	|                    | equivalent.                                                                      |
	// 	| OnStartupSec=      | Defines a timer relative to when the service manager was first started. For      |
	// 	|                    | system timer units this is very similar to OnBootSec= as the system service      |
	// 	|                    | manager is generally started very early at boot. It's primarily useful when      |
	// 	|                    | configured in units running in the per-user service manager, as the user service |
	// 	|                    | manager is generally started on first login only, not already during boot.       |
	// 	| OnUnitActiveSec=   | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | activated.                                                                       |
	// 	| OnUnitInactiveSec= | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | deactivated.                                                                     |
	// 	+--------------------+----------------------------------------------------------------------------------+
	//
	// Multiple directives may be combined of the same and of different types, in which case the timer unit will trigger whenever
	// any of the specified timer expressions elapse. For example, by combining OnBootSec= and OnUnitActiveSec=, it is possible
	// to define a timer that elapses in regular intervals and activates a specific service each time. Moreover, both monotonic
	// time expressions and OnCalendar= calendar expressions may be combined in the same timer unit.
	//
	// The arguments to the directives are time spans configured in seconds. Example: "OnBootSec=50" means 50s after boot-up.
	// The argument may also include time units. Example: "OnBootSec=5h 30min" means 5 hours and 30 minutes after boot-up. For
	// details about the syntax of time spans, see systemd.time.
	//
	// If a timer configured with OnBootSec= or OnStartupSec= is already in the past when the timer unit is activated, it will immediately
	// elapse and the configured unit is started. This is not the case for timers defined in the other directives.
	//
	// These are monotonic timers, independent of wall-clock time and timezones. If the computer is temporarily suspended,
	// the monotonic clock pauses, too.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both monotonic timers and OnCalendar=
	// timers, see below), and all prior assignments will have no effect.
	//
	// Note that timers do not necessarily expire at the precise time configured with these settings, as they are subject to the
	// AccuracySec= setting below.
	OnStartupSec systemdconf.Value

	// Defines monotonic timers relative to different starting points:
	//
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	|      SETTING       |                                     MEANING                                      |
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	| OnActiveSec=       | Defines a timer relative to the moment the timer unit itself is activated.       |
	// 	| OnBootSec=         | Defines a timer relative to when the machine was booted up. In containers,       |
	// 	|                    | for the system manager instance, this is mapped to OnStartupSec=, making both    |
	// 	|                    | equivalent.                                                                      |
	// 	| OnStartupSec=      | Defines a timer relative to when the service manager was first started. For      |
	// 	|                    | system timer units this is very similar to OnBootSec= as the system service      |
	// 	|                    | manager is generally started very early at boot. It's primarily useful when      |
	// 	|                    | configured in units running in the per-user service manager, as the user service |
	// 	|                    | manager is generally started on first login only, not already during boot.       |
	// 	| OnUnitActiveSec=   | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | activated.                                                                       |
	// 	| OnUnitInactiveSec= | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | deactivated.                                                                     |
	// 	+--------------------+----------------------------------------------------------------------------------+
	//
	// Multiple directives may be combined of the same and of different types, in which case the timer unit will trigger whenever
	// any of the specified timer expressions elapse. For example, by combining OnBootSec= and OnUnitActiveSec=, it is possible
	// to define a timer that elapses in regular intervals and activates a specific service each time. Moreover, both monotonic
	// time expressions and OnCalendar= calendar expressions may be combined in the same timer unit.
	//
	// The arguments to the directives are time spans configured in seconds. Example: "OnBootSec=50" means 50s after boot-up.
	// The argument may also include time units. Example: "OnBootSec=5h 30min" means 5 hours and 30 minutes after boot-up. For
	// details about the syntax of time spans, see systemd.time.
	//
	// If a timer configured with OnBootSec= or OnStartupSec= is already in the past when the timer unit is activated, it will immediately
	// elapse and the configured unit is started. This is not the case for timers defined in the other directives.
	//
	// These are monotonic timers, independent of wall-clock time and timezones. If the computer is temporarily suspended,
	// the monotonic clock pauses, too.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both monotonic timers and OnCalendar=
	// timers, see below), and all prior assignments will have no effect.
	//
	// Note that timers do not necessarily expire at the precise time configured with these settings, as they are subject to the
	// AccuracySec= setting below.
	OnUnitActiveSec systemdconf.Value

	// Defines monotonic timers relative to different starting points:
	//
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	|      SETTING       |                                     MEANING                                      |
	// 	+--------------------+----------------------------------------------------------------------------------+
	// 	| OnActiveSec=       | Defines a timer relative to the moment the timer unit itself is activated.       |
	// 	| OnBootSec=         | Defines a timer relative to when the machine was booted up. In containers,       |
	// 	|                    | for the system manager instance, this is mapped to OnStartupSec=, making both    |
	// 	|                    | equivalent.                                                                      |
	// 	| OnStartupSec=      | Defines a timer relative to when the service manager was first started. For      |
	// 	|                    | system timer units this is very similar to OnBootSec= as the system service      |
	// 	|                    | manager is generally started very early at boot. It's primarily useful when      |
	// 	|                    | configured in units running in the per-user service manager, as the user service |
	// 	|                    | manager is generally started on first login only, not already during boot.       |
	// 	| OnUnitActiveSec=   | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | activated.                                                                       |
	// 	| OnUnitInactiveSec= | Defines a timer relative to when the unit the timer unit is activating was last  |
	// 	|                    | deactivated.                                                                     |
	// 	+--------------------+----------------------------------------------------------------------------------+
	//
	// Multiple directives may be combined of the same and of different types, in which case the timer unit will trigger whenever
	// any of the specified timer expressions elapse. For example, by combining OnBootSec= and OnUnitActiveSec=, it is possible
	// to define a timer that elapses in regular intervals and activates a specific service each time. Moreover, both monotonic
	// time expressions and OnCalendar= calendar expressions may be combined in the same timer unit.
	//
	// The arguments to the directives are time spans configured in seconds. Example: "OnBootSec=50" means 50s after boot-up.
	// The argument may also include time units. Example: "OnBootSec=5h 30min" means 5 hours and 30 minutes after boot-up. For
	// details about the syntax of time spans, see systemd.time.
	//
	// If a timer configured with OnBootSec= or OnStartupSec= is already in the past when the timer unit is activated, it will immediately
	// elapse and the configured unit is started. This is not the case for timers defined in the other directives.
	//
	// These are monotonic timers, independent of wall-clock time and timezones. If the computer is temporarily suspended,
	// the monotonic clock pauses, too.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both monotonic timers and OnCalendar=
	// timers, see below), and all prior assignments will have no effect.
	//
	// Note that timers do not necessarily expire at the precise time configured with these settings, as they are subject to the
	// AccuracySec= setting below.
	OnUnitInactiveSec systemdconf.Value

	// Defines realtime (i.e. wallclock) timers with calendar event expressions. See systemd.time for more information on
	// the syntax of calendar event expressions. Otherwise, the semantics are similar to OnActiveSec= and related settings.
	//
	// Note that timers do not necessarily expire at the precise time configured with this setting, as it is subject to the AccuracySec=
	// setting below.
	//
	// May be specified more than once, in which case the timer unit will trigger whenever any of the specified expressions elapse.
	// Moreover calendar timers and monotonic timers (see above) may be combined within the same timer unit.
	//
	// If the empty string is assigned to any of these options, the list of timers is reset (both OnCalendar= timers and monotonic
	// timers, see above), and all prior assignments will have no effect.
	OnCalendar systemdconf.Value

	// Specify the accuracy the timer shall elapse with. Defaults to 1min. The timer is scheduled to elapse within a time window
	// starting with the time specified in OnCalendar=, OnActiveSec=, OnBootSec=, OnStartupSec=, OnUnitActiveSec= or OnUnitInactiveSec=
	// and ending the time configured with AccuracySec= later. Within this time window, the expiry time will be placed at a host-specific,
	// randomized, but stable position that is synchronized between all local timer units. This is done in order to optimize power
	// consumption to suppress unnecessary CPU wake-ups. To get best accuracy, set this option to 1us. Note that the timer is still
	// subject to the timer slack configured via systemd-system.conf's TimerSlackNSec= setting. See prctl for details. To
	// optimize power consumption, make sure to set this value as high as possible and as low as necessary.
	AccuracySec systemdconf.Value

	// Delay the timer by a randomly selected, evenly distributed amount of time between 0 and the specified time value. Defaults
	// to 0, indicating that no randomized delay shall be applied. Each timer unit will determine this delay randomly before each
	// iteration, and the delay will simply be added on top of the next determined elapsing time. This is useful to stretch dispatching
	// of similarly configured timer events over a certain amount time, to avoid that they all fire at the same time, possibly resulting
	// in resource congestion. Note the relation to AccuracySec= above: the latter allows the service manager to coalesce timer
	// events within a specified time range in order to minimize wakeups, the former does the opposite: it stretches timer events
	// over a time range, to make it unlikely that they fire simultaneously. If RandomizedDelaySec= and AccuracySec= are used
	// in conjunction, first the randomized delay is added, and then the result is possibly further shifted to coalesce it with
	// other timer events happening on the system. As mentioned above AccuracySec= defaults to 1min and RandomizedDelaySec=
	// to 0, thus encouraging coalescing of timer events. In order to optimally stretch timer events over a certain range of time,
	// make sure to set RandomizedDelaySec= to a higher value, and AccuracySec=1us.
	RandomizedDelaySec systemdconf.Value

	// These options take boolean arguments. When true, the service unit will be triggered when the system clock (CLOCK_REALTIME)
	// jumps relative to the monotonic clock (CLOCK_MONOTONIC), or when the local system timezone is modified. These options
	// can be used alone or in combination with other timer expressions (see above) within the same timer unit. These options default
	// to false.
	OnClockChange systemdconf.Value

	// These options take boolean arguments. When true, the service unit will be triggered when the system clock (CLOCK_REALTIME)
	// jumps relative to the monotonic clock (CLOCK_MONOTONIC), or when the local system timezone is modified. These options
	// can be used alone or in combination with other timer expressions (see above) within the same timer unit. These options default
	// to false.
	OnTimezoneChange systemdconf.Value

	// The unit to activate when this timer elapses. The argument is a unit name, whose suffix is not ".timer". If not specified,
	// this value defaults to a service that has the same name as the timer unit, except for the suffix. (See above.) It is recommended
	// that the unit name that is activated and the unit name of the timer unit are named identically, except for the suffix.
	Unit systemdconf.Value

	// Takes a boolean argument. If true, the time when the service unit was last triggered is stored on disk. When the timer is activated,
	// the service unit is triggered immediately if it would have been triggered at least once during the time when the timer was
	// inactive. This is useful to catch up on missed runs of the service when the system was powered down. Note that this setting
	// only has an effect on timers configured with OnCalendar=. Defaults to false.
	//
	// Use systemctl clean --what=state … on the timer unit to remove the timestamp file maintained by this option from disk. In
	// particular, use this command before uninstalling a timer unit. See systemctl for details.
	Persistent systemdconf.Value

	// Takes a boolean argument. If true, an elapsing timer will cause the system to resume from suspend, should it be suspended
	// and if the system supports this. Note that this option will only make sure the system resumes on the appropriate times, it
	// will not take care of suspending it again after any work that is to be done is finished. Defaults to false.
	//
	// Note that this functionality requires privileges and is thus generally only available in the system service manager.
	WakeSystem systemdconf.Value

	// Takes a boolean argument. If true, an elapsed timer will stay loaded, and its state remains queryable. If false, an elapsed
	// timer unit that cannot elapse anymore is unloaded. Turning this off is particularly useful for transient timer units that
	// shall disappear after they first elapse. Note that this setting has an effect on repeatedly starting a timer unit that only
	// elapses once: if RemainAfterElapse= is on, it will not be started again, and is guaranteed to elapse only once. However,
	// if RemainAfterElapse= is off, it might be started again if it is already elapsed, and thus be triggered multiple times.
	// Defaults to yes.
	RemainAfterElapse systemdconf.Value
}

// TimerFile represents information about a timer controlled and supervised by systemd, for timer-based activation
type TimerFile struct {
	systemdconf.File

	Unit    UnitSection    // Generic information about the unit that is not dependent on the type of unit
	Timer   TimerSection   // Information about the timer it defines
	Install InstallSection // Installation information for the unit
}
