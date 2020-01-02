// DO NOT EDIT. This file is generated from systemd 244 by generatesdconf

package unit

import "github.com/sergeymakinen/go-systemdconf"

// UnitSection represents generic information about the unit that is not dependent on the type of unit
type UnitSection struct {
	systemdconf.Section

	// A human readable name for the unit. This is used by systemd (and other UIs) as the label for the unit, so this string should
	// identify the unit rather than describe it, despite the name. "Apache2 Web Server" is a good example. Bad examples are "high-performance
	// light-weight HTTP server" (too generic) or "Apache2" (too specific and meaningless for people who do not know Apache).
	// systemd will use this string as a noun in status messages ("Starting description...", "Started description.", "Reached
	// target description.", "Failed to start description."), so it should be capitalized, and should not be a full sentence
	// or a phrase with a continuous verb. Bad examples include "exiting the container" or "updating the database once per day.".
	Description systemdconf.Value

	// A space-separated list of URIs referencing documentation for this unit or its configuration. Accepted are only URIs of
	// the types "http://", "https://", "file:", "info:", "man:". For more information about the syntax of these URIs, see uri.
	// The URIs should be listed in order of relevance, starting with the most relevant. It is a good idea to first reference documentation
	// that explains what the unit's purpose is, followed by how it is configured, followed by any other related documentation.
	// This option may be specified more than once, in which case the specified list of URIs is merged. If the empty string is assigned
	// to this option, the list is reset and all prior assignments will have no effect.
	Documentation systemdconf.Value

	// Configures requirement dependencies on other units. This option may be specified more than once or multiple space-separated
	// units may be specified in one option in which case dependencies for all listed names will be created. Dependencies of this
	// type may also be configured outside of the unit configuration file by adding a symlink to a .wants/ directory accompanying
	// the unit file. For details, see above.
	//
	// Units listed in this option will be started if the configuring unit is. However, if the listed units fail to start or cannot
	// be added to the transaction, this has no impact on the validity of the transaction as a whole, and this unit will still be started.
	// This is the recommended way to hook start-up of one unit to the start-up of another unit.
	//
	// Note that requirement dependencies do not influence the order in which services are started or stopped. This has to be configured
	// independently with the After= or Before= options. If unit foo.service pulls in unit bar.service as configured with Wants=
	// and no ordering is configured with After= or Before=, then both units will be started simultaneously and without any delay
	// between them if foo.service is activated.
	Wants systemdconf.Value

	// Similar to Wants=, but declares a stronger dependency. Dependencies of this type may also be configured by adding a symlink
	// to a .requires/ directory accompanying the unit file.
	//
	// If this unit gets activated, the units listed will be activated as well. If one of the other units fails to activate, and an
	// ordering dependency After= on the failing unit is set, this unit will not be started. Besides, with or without specifying
	// After=, this unit will be stopped if one of the other units is explicitly stopped.
	//
	// Often, it is a better choice to use Wants= instead of Requires= in order to achieve a system that is more robust when dealing
	// with failing services.
	//
	// Note that this dependency type does not imply that the other unit always has to be in active state when this unit is running.
	// Specifically: failing condition checks (such as ConditionPathExists=, ConditionPathIsSymbolicLink=, … — see below)
	// do not cause the start job of a unit with a Requires= dependency on it to fail. Also, some unit types may deactivate on their
	// own (for example, a service process may decide to exit cleanly, or a device may be unplugged by the user), which is not propagated
	// to units having a Requires= dependency. Use the BindsTo= dependency type together with After= to ensure that a unit may
	// never be in active state without a specific other unit also in active state (see below).
	Requires systemdconf.Value

	// Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting
	// of this unit will fail immediately. Requisite= does not imply an ordering dependency, even if both units are started in
	// the same transaction. Hence this setting should usually be combined with After=, to ensure this unit is not started before
	// the other unit.
	//
	// When Requisite=b.service is used on a.service, this dependency will show as RequisiteOf=a.service in property listing
	// of b.service. RequisiteOf= dependency cannot be specified directly.
	Requisite systemdconf.Value

	// Configures requirement dependencies, very similar in style to Requires=. However, this dependency type is stronger:
	// in addition to the effect of Requires= it declares that if the unit bound to is stopped, this unit will be stopped too. This
	// means a unit bound to another unit that suddenly enters inactive state will be stopped too. Units can suddenly, unexpectedly
	// enter inactive state for different reasons: the main process of a service unit might terminate on its own choice, the backing
	// device of a device unit might be unplugged or the mount point of a mount unit might be unmounted without involvement of the
	// system and service manager.
	//
	// When used in conjunction with After= on the same unit the behaviour of BindsTo= is even stronger. In this case, the unit bound
	// to strictly has to be in active state for this unit to also be in active state. This not only means a unit bound to another unit
	// that suddenly enters inactive state, but also one that is bound to another unit that gets skipped due to a failed condition
	// check (such as ConditionPathExists=, ConditionPathIsSymbolicLink=, … — see below) will be stopped, should it be running.
	// Hence, in many cases it is best to combine BindsTo= with After=.
	//
	// When BindsTo=b.service is used on a.service, this dependency will show as BoundBy=a.service in property listing of b.service.
	// BoundBy= dependency cannot be specified directly.
	BindsTo systemdconf.Value

	// Configures dependencies similar to Requires=, but limited to stopping and restarting of units. When systemd stops or
	// restarts the units listed here, the action is propagated to this unit. Note that this is a one-way dependency — changes to
	// this unit do not affect the listed units.
	//
	// When PartOf=b.service is used on a.service, this dependency will show as ConsistsOf=a.service in property listing of
	// b.service. ConsistsOf= dependency cannot be specified directly.
	PartOf systemdconf.Value

	// A space-separated list of unit names. Configures negative requirement dependencies. If a unit has a Conflicts= setting
	// on another unit, starting the former will stop the latter and vice versa.
	//
	// Note that this setting does not imply an ordering dependency, similarly to the Wants= and Requires= dependencies described
	// above. This means that to ensure that the conflicting unit is stopped before the other unit is started, an After= or Before=
	// dependency must be declared. It doesn't matter which of the two ordering dependencies is used, because stop jobs are always
	// ordered before start jobs, see the discussion in Before=/After= below.
	//
	// If unit A that conflicts with unit B is scheduled to be started at the same time as B, the transaction will either fail (in case
	// both are required parts of the transaction) or be modified to be fixed (in case one or both jobs are not a required part of the
	// transaction). In the latter case, the job that is not required will be removed, or in case both are not required, the unit
	// that conflicts will be started and the unit that is conflicted is stopped.
	Conflicts systemdconf.Value

	// These two settings expect a space-separated list of unit names. They may be specified more than once, in which case dependencies
	// for all listed names are created.
	//
	// Those two setttings configure ordering dependencies between units. If unit foo.service contains the setting Before=bar.service
	// and both units are being started, bar.service's start-up is delayed until foo.service has finished starting up. After=
	// is the inverse of Before=, i.e. while Before= ensures that the configured unit is started before the listed unit begins
	// starting up, After= ensures the opposite, that the listed unit is fully started up before the configured unit is started.
	//
	// When two units with an ordering dependency between them are shut down, the inverse of the start-up order is applied. i.e.
	// if a unit is configured with After= on another unit, the former is stopped before the latter if both are shut down. Given two
	// units with any ordering dependency between them, if one unit is shut down and the other is started up, the shutdown is ordered
	// before the start-up. It doesn't matter if the ordering dependency is After= or Before=, in this case. It also doesn't matter
	// which of the two is shut down, as long as one is shut down and the other is started up; the shutdown is ordered before the start-up
	// in all cases. If two units have no ordering dependencies between them, they are shut down or started up simultaneously,
	// and no ordering takes place. It depends on the unit type when precisely a unit has finished starting up. Most importantly,
	// for service units start-up is considered completed for the purpose of Before=/After= when all its configured start-up
	// commands have been invoked and they either failed or reported start-up success.
	//
	// Note that those settings are independent of and orthogonal to the requirement dependencies as configured by Requires=,
	// Wants=, Requisite=, or BindsTo=. It is a common pattern to include a unit name in both the After= and Wants= options, in which
	// case the unit listed will be started before the unit that is configured with these options.
	Before systemdconf.Value

	// These two settings expect a space-separated list of unit names. They may be specified more than once, in which case dependencies
	// for all listed names are created.
	//
	// Those two setttings configure ordering dependencies between units. If unit foo.service contains the setting Before=bar.service
	// and both units are being started, bar.service's start-up is delayed until foo.service has finished starting up. After=
	// is the inverse of Before=, i.e. while Before= ensures that the configured unit is started before the listed unit begins
	// starting up, After= ensures the opposite, that the listed unit is fully started up before the configured unit is started.
	//
	// When two units with an ordering dependency between them are shut down, the inverse of the start-up order is applied. i.e.
	// if a unit is configured with After= on another unit, the former is stopped before the latter if both are shut down. Given two
	// units with any ordering dependency between them, if one unit is shut down and the other is started up, the shutdown is ordered
	// before the start-up. It doesn't matter if the ordering dependency is After= or Before=, in this case. It also doesn't matter
	// which of the two is shut down, as long as one is shut down and the other is started up; the shutdown is ordered before the start-up
	// in all cases. If two units have no ordering dependencies between them, they are shut down or started up simultaneously,
	// and no ordering takes place. It depends on the unit type when precisely a unit has finished starting up. Most importantly,
	// for service units start-up is considered completed for the purpose of Before=/After= when all its configured start-up
	// commands have been invoked and they either failed or reported start-up success.
	//
	// Note that those settings are independent of and orthogonal to the requirement dependencies as configured by Requires=,
	// Wants=, Requisite=, or BindsTo=. It is a common pattern to include a unit name in both the After= and Wants= options, in which
	// case the unit listed will be started before the unit that is configured with these options.
	After systemdconf.Value

	// A space-separated list of one or more units that are activated when this unit enters the "failed" state. A service unit using
	// Restart= enters the failed state only after the start limits are reached.
	OnFailure systemdconf.Value

	// A space-separated list of one or more units where reload requests on this unit will be propagated to, or reload requests
	// on the other unit will be propagated to this unit, respectively. Issuing a reload request on a unit will automatically also
	// enqueue a reload request on all units that the reload request shall be propagated to via these two settings.
	PropagatesReloadTo systemdconf.Value

	// A space-separated list of one or more units where reload requests on this unit will be propagated to, or reload requests
	// on the other unit will be propagated to this unit, respectively. Issuing a reload request on a unit will automatically also
	// enqueue a reload request on all units that the reload request shall be propagated to via these two settings.
	ReloadPropagatedFrom systemdconf.Value

	// For units that start processes (such as service units), lists one or more other units whose network and/or temporary file
	// namespace to join. This only applies to unit types which support the PrivateNetwork=, NetworkNamespacePath= and PrivateTmp=
	// directives (see systemd.exec for details). If a unit that has this setting set is started, its processes will see the same
	// /tmp, /var/tmp and network namespace as one listed unit that is started. If multiple listed units are already started,
	// it is not defined which namespace is joined. Note that this setting only has an effect if PrivateNetwork=/NetworkNamespacePath=
	// and/or PrivateTmp= is enabled for both the unit that joins the namespace and the unit whose namespace is joined.
	JoinsNamespaceOf systemdconf.Value

	// Takes a space-separated list of absolute paths. Automatically adds dependencies of type Requires= and After= for all
	// mount units required to access the specified path.
	//
	// Mount points marked with noauto are not mounted automatically through local-fs.target, but are still honored for the
	// purposes of this option, i.e. they will be pulled in by this unit.
	RequiresMountsFor systemdconf.Value

	// Takes a value of "fail", "replace", "replace-irreversibly", "isolate", "flush", "ignore-dependencies" or "ignore-requirements".
	// Defaults to "replace". Specifies how the units listed in OnFailure= will be enqueued. See systemctl's --job-mode= option
	// for details on the possible values. If this is set to "isolate", only a single unit may be listed in OnFailure=..
	OnFailureJobMode systemdconf.Value

	// Takes a boolean argument. If true, this unit will not be stopped when isolating another unit. Defaults to false for service,
	// target, socket, busname, timer, and path units, and true for slice, scope, device, swap, mount, and automount units.
	IgnoreOnIsolate systemdconf.Value

	// Takes a boolean argument. If true, this unit will be stopped when it is no longer used. Note that, in order to minimize the
	// work to be executed, systemd will not stop units by default unless they are conflicting with other units, or the user explicitly
	// requested their shut down. If this option is set, a unit will be automatically cleaned up if no other active unit requires
	// it. Defaults to false.
	StopWhenUnneeded systemdconf.Value

	// Takes a boolean argument. If true, this unit can only be activated or deactivated indirectly. In this case, explicit start-up
	// or termination requested by the user is denied, however if it is started or stopped as a dependency of another unit, start-up
	// or termination will succeed. This is mostly a safety feature to ensure that the user does not accidentally activate units
	// that are not intended to be activated explicitly, and not accidentally deactivate units that are not intended to be deactivated.
	// These options default to false.
	RefuseManualStart systemdconf.Value

	// Takes a boolean argument. If true, this unit can only be activated or deactivated indirectly. In this case, explicit start-up
	// or termination requested by the user is denied, however if it is started or stopped as a dependency of another unit, start-up
	// or termination will succeed. This is mostly a safety feature to ensure that the user does not accidentally activate units
	// that are not intended to be activated explicitly, and not accidentally deactivate units that are not intended to be deactivated.
	// These options default to false.
	RefuseManualStop systemdconf.Value

	// Takes a boolean argument. If true, this unit may be used with the systemctl isolate command. Otherwise, this will be refused.
	// It probably is a good idea to leave this disabled except for target units that shall be used similar to runlevels in SysV init
	// systems, just as a precaution to avoid unusable system states. This option defaults to false.
	AllowIsolate systemdconf.Value

	// Takes a boolean argument. If yes, (the default), a few default dependencies will implicitly be created for the unit. The
	// actual dependencies created depend on the unit type. For example, for service units, these dependencies ensure that the
	// service is started only after basic system initialization is completed and is properly terminated on system shutdown.
	// See the respective man pages for details. Generally, only services involved with early boot or late shutdown should set
	// this option to no. It is highly recommended to leave this option enabled for the majority of common units. If set to no, this
	// option does not disable all implicit dependencies, just non-essential ones.
	DefaultDependencies systemdconf.Value

	// Tweaks the "garbage collection" algorithm for this unit. Takes one of inactive or inactive-or-failed. If set to inactive
	// the unit will be unloaded if it is in the inactive state and is not referenced by clients, jobs or other units — however it is
	// not unloaded if it is in the failed state. In failed mode, failed units are not unloaded until the user invoked systemctl
	// reset-failed on them to reset the failed state, or an equivalent command. This behaviour is altered if this option is set
	// to inactive-or-failed: in this case the unit is unloaded even if the unit is in a failed state, and thus an explicitly resetting
	// of the failed state is not necessary. Note that if this mode is used unit results (such as exit codes, exit signals, consumed
	// resources, …) are flushed out immediately after the unit completed, except for what is stored in the logging subsystem.
	// Defaults to inactive.
	CollectMode systemdconf.Value

	// Configure the action to take when the unit stops and enters a failed state or inactive state. Takes one of none, reboot, reboot-force,
	// reboot-immediate, poweroff, poweroff-force, poweroff-immediate, exit, and exit-force. In system mode, all options
	// are allowed. In user mode, only none, exit, and exit-force are allowed. Both options default to none.
	//
	// If none is set, no action will be triggered. reboot causes a reboot following the normal shutdown procedure (i.e. equivalent
	// to systemctl reboot). reboot-force causes a forced reboot which will terminate all processes forcibly but should cause
	// no dirty file systems on reboot (i.e. equivalent to systemctl reboot -f) and reboot-immediate causes immediate execution
	// of the reboot system call, which might result in data loss (i.e. equivalent to systemctl reboot -ff). Similarly, poweroff,
	// poweroff-force, poweroff-immediate have the effect of powering down the system with similar semantics. exit causes
	// the manager to exit following the normal shutdown procedure, and exit-force causes it terminate without shutting down
	// services. When exit or exit-force is used by default the exit status of the main process of the unit (if this applies) is returned
	// from the service manager. However, this may be overridden with FailureActionExitStatus=/SuccessActionExitStatus=,
	// see below.
	FailureAction systemdconf.Value

	// Configure the action to take when the unit stops and enters a failed state or inactive state. Takes one of none, reboot, reboot-force,
	// reboot-immediate, poweroff, poweroff-force, poweroff-immediate, exit, and exit-force. In system mode, all options
	// are allowed. In user mode, only none, exit, and exit-force are allowed. Both options default to none.
	//
	// If none is set, no action will be triggered. reboot causes a reboot following the normal shutdown procedure (i.e. equivalent
	// to systemctl reboot). reboot-force causes a forced reboot which will terminate all processes forcibly but should cause
	// no dirty file systems on reboot (i.e. equivalent to systemctl reboot -f) and reboot-immediate causes immediate execution
	// of the reboot system call, which might result in data loss (i.e. equivalent to systemctl reboot -ff). Similarly, poweroff,
	// poweroff-force, poweroff-immediate have the effect of powering down the system with similar semantics. exit causes
	// the manager to exit following the normal shutdown procedure, and exit-force causes it terminate without shutting down
	// services. When exit or exit-force is used by default the exit status of the main process of the unit (if this applies) is returned
	// from the service manager. However, this may be overridden with FailureActionExitStatus=/SuccessActionExitStatus=,
	// see below.
	SuccessAction systemdconf.Value

	// Controls the exit status to propagate back to an invoking container manager (in case of a system service) or service manager
	// (in case of a user manager) when the FailureAction=/SuccessAction= are set to exit or exit-force and the action is triggered.
	// By default the exit status of the main process of the triggering unit (if this applies) is propagated. Takes a value in the
	// range 0…255 or the empty string to request default behaviour.
	FailureActionExitStatus systemdconf.Value

	// Controls the exit status to propagate back to an invoking container manager (in case of a system service) or service manager
	// (in case of a user manager) when the FailureAction=/SuccessAction= are set to exit or exit-force and the action is triggered.
	// By default the exit status of the main process of the triggering unit (if this applies) is propagated. Takes a value in the
	// range 0…255 or the empty string to request default behaviour.
	SuccessActionExitStatus systemdconf.Value

	// When a job for this unit is queued, a timeout JobTimeoutSec= may be configured. Similarly, JobRunningTimeoutSec= starts
	// counting when the queued job is actually started. If either time limit is reached, the job will be cancelled, the unit however
	// will not change state or even enter the "failed" mode. This value defaults to "infinity" (job timeouts disabled), except
	// for device units (JobRunningTimeoutSec= defaults to DefaultTimeoutStartSec=). NB: this timeout is independent from
	// any unit-specific timeout (for example, the timeout set with TimeoutStartSec= in service units) as the job timeout has
	// no effect on the unit itself, only on the job that might be pending for it. Or in other words: unit-specific timeouts are useful
	// to abort unit state changes, and revert them. The job timeout set with this option however is useful to abort only the job
	// waiting for the unit state to change.
	JobTimeoutSec systemdconf.Value

	// When a job for this unit is queued, a timeout JobTimeoutSec= may be configured. Similarly, JobRunningTimeoutSec= starts
	// counting when the queued job is actually started. If either time limit is reached, the job will be cancelled, the unit however
	// will not change state or even enter the "failed" mode. This value defaults to "infinity" (job timeouts disabled), except
	// for device units (JobRunningTimeoutSec= defaults to DefaultTimeoutStartSec=). NB: this timeout is independent from
	// any unit-specific timeout (for example, the timeout set with TimeoutStartSec= in service units) as the job timeout has
	// no effect on the unit itself, only on the job that might be pending for it. Or in other words: unit-specific timeouts are useful
	// to abort unit state changes, and revert them. The job timeout set with this option however is useful to abort only the job
	// waiting for the unit state to change.
	JobRunningTimeoutSec systemdconf.Value

	// JobTimeoutAction= optionally configures an additional action to take when the timeout is hit, see description of JobTimeoutSec=
	// and JobRunningTimeoutSec= above. It takes the same values as StartLimitAction=. Defaults to none. JobTimeoutRebootArgument=
	// configures an optional reboot string to pass to the reboot system call.
	JobTimeoutAction systemdconf.Value

	// JobTimeoutAction= optionally configures an additional action to take when the timeout is hit, see description of JobTimeoutSec=
	// and JobRunningTimeoutSec= above. It takes the same values as StartLimitAction=. Defaults to none. JobTimeoutRebootArgument=
	// configures an optional reboot string to pass to the reboot system call.
	JobTimeoutRebootArgument systemdconf.Value

	// Configure unit start rate limiting. Units which are started more than burst times within an interval time interval are
	// not permitted to start any more. Use StartLimitIntervalSec= to configure the checking interval (defaults to DefaultStartLimitIntervalSec=
	// in manager configuration file, set it to 0 to disable any kind of rate limiting). Use StartLimitBurst= to configure how
	// many starts per interval are allowed (defaults to DefaultStartLimitBurst= in manager configuration file). These configuration
	// options are particularly useful in conjunction with the service setting Restart= (see systemd.service); however, they
	// apply to all kinds of starts (including manual), not just those triggered by the Restart= logic. Note that units which are
	// configured for Restart= and which reach the start limit are not attempted to be restarted anymore; however, they may still
	// be restarted manually at a later point, after the interval has passed. From this point on, the restart logic is activated
	// again. Note that systemctl reset-failed will cause the restart rate counter for a service to be flushed, which is useful
	// if the administrator wants to manually start a unit and the start limit interferes with that. Note that this rate-limiting
	// is enforced after any unit condition checks are executed, and hence unit activations with failing conditions do not count
	// towards this rate limit. This setting does not apply to slice, target, device, and scope units, since they are unit types
	// whose activation may either never fail, or may succeed only a single time.
	//
	// When a unit is unloaded due to the garbage collection logic (see above) its rate limit counters are flushed out too. This
	// means that configuring start rate limiting for a unit that is not referenced continuously has no effect.
	StartLimitIntervalSec systemdconf.Value

	// Configure unit start rate limiting. Units which are started more than burst times within an interval time interval are
	// not permitted to start any more. Use StartLimitIntervalSec= to configure the checking interval (defaults to DefaultStartLimitIntervalSec=
	// in manager configuration file, set it to 0 to disable any kind of rate limiting). Use StartLimitBurst= to configure how
	// many starts per interval are allowed (defaults to DefaultStartLimitBurst= in manager configuration file). These configuration
	// options are particularly useful in conjunction with the service setting Restart= (see systemd.service); however, they
	// apply to all kinds of starts (including manual), not just those triggered by the Restart= logic. Note that units which are
	// configured for Restart= and which reach the start limit are not attempted to be restarted anymore; however, they may still
	// be restarted manually at a later point, after the interval has passed. From this point on, the restart logic is activated
	// again. Note that systemctl reset-failed will cause the restart rate counter for a service to be flushed, which is useful
	// if the administrator wants to manually start a unit and the start limit interferes with that. Note that this rate-limiting
	// is enforced after any unit condition checks are executed, and hence unit activations with failing conditions do not count
	// towards this rate limit. This setting does not apply to slice, target, device, and scope units, since they are unit types
	// whose activation may either never fail, or may succeed only a single time.
	//
	// When a unit is unloaded due to the garbage collection logic (see above) its rate limit counters are flushed out too. This
	// means that configuring start rate limiting for a unit that is not referenced continuously has no effect.
	StartLimitBurst systemdconf.Value

	// Configure an additional action to take if the rate limit configured with StartLimitIntervalSec= and StartLimitBurst=
	// is hit. Takes the same values as the setting FailureAction=/SuccessAction= settings and executes the same actions. If
	// none is set, hitting the rate limit will trigger no action besides that the start will not be permitted. Defaults to none.
	StartLimitAction systemdconf.Value

	// Configure the optional argument for the reboot system call if StartLimitAction= or FailureAction= is a reboot action.
	// This works just like the optional argument to systemctl reboot command.
	RebootArgument systemdconf.Value

	// A path to a configuration file this unit has been generated from. This is primarily useful for implementation of generator
	// tools that convert configuration from an external configuration file format into native unit files. This functionality
	// should not be used in normal units.
	SourcePath systemdconf.Value

	// Check whether the system is running on a specific architecture. Takes one of "x86", "x86-64", "ppc", "ppc-le", "ppc64",
	// "ppc64-le", "ia64", "parisc", "parisc64", "s390", "s390x", "sparc", "sparc64", "mips", "mips-le", "mips64", "mips64-le",
	// "alpha", "arm", "arm-be", "arm64", "arm64-be", "sh", "sh64", "m68k", "tilegx", "cris", "arc", "arc-be", or "native".
	//
	// The architecture is determined from the information returned by uname and is thus subject to personality. Note that a Personality=
	// setting in the same unit file has no effect on this condition. A special architecture name "native" is mapped to the architecture
	// the system manager itself is compiled for. The test may be negated by prepending an exclamation mark.
	ConditionArchitecture systemdconf.Value

	// Check whether the system is executed in a virtualized environment and optionally test whether it is a specific implementation.
	// Takes either boolean value to check if being executed in any virtualized environment, or one of "vm" and "container" to
	// test against a generic type of virtualization solution, or one of "qemu", "kvm", "zvm", "vmware", "microsoft", "oracle",
	// "xen", "bochs", "uml", "bhyve", "qnx", "openvz", "lxc", "lxc-libvirt", "systemd-nspawn", "docker", "podman", "rkt",
	// "wsl", "acrn" to test against a specific implementation, or "private-users" to check whether we are running in a user namespace.
	// See systemd-detect-virt for a full list of known virtualization technologies and their identifiers. If multiple virtualization
	// technologies are nested, only the innermost is considered. The test may be negated by prepending an exclamation mark.
	ConditionVirtualization systemdconf.Value

	// ConditionHost= may be used to match against the hostname or machine ID of the host. This either takes a hostname string (optionally
	// with shell style globs) which is tested against the locally set hostname as returned by gethostname, or a machine ID formatted
	// as string (see machine-id). The test may be negated by prepending an exclamation mark.
	ConditionHost systemdconf.Value

	// ConditionKernelCommandLine= may be used to check whether a specific kernel command line option is set (or if prefixed
	// with the exclamation mark — unset). The argument must either be a single word, or an assignment (i.e. two words, separated
	// by "="). In the former case the kernel command line is searched for the word appearing as is, or as left hand side of an assignment.
	// In the latter case, the exact assignment is looked for with right and left hand side matching.
	ConditionKernelCommandLine systemdconf.Value

	// ConditionKernelVersion= may be used to check whether the kernel version (as reported by uname -r) matches a certain expression
	// (or if prefixed with the exclamation mark does not match it). The argument must be a list of (potentially quoted) expressions.
	// For each of the expressions, if it starts with one of "<", "<=", "=", "!=", ">=", ">" a relative version comparison is done,
	// otherwise the specified string is matched with shell-style globs.
	//
	// Note that using the kernel version string is an unreliable way to determine which features are supported by a kernel, because
	// of the widespread practice of backporting drivers, features, and fixes from newer upstream kernels into older versions
	// provided by distributions. Hence, this check is inherently unportable and should not be used for units which may be used
	// on different distributions.
	ConditionKernelVersion systemdconf.Value

	// ConditionSecurity= may be used to check whether the given security technology is enabled on the system. Currently, the
	// recognized values are "selinux", "apparmor", "tomoyo", "ima", "smack", "audit" and "uefi-secureboot". The test may
	// be negated by prepending an exclamation mark.
	ConditionSecurity systemdconf.Value

	// Check whether the given capability exists in the capability bounding set of the service manager (i.e. this does not check
	// whether capability is actually available in the permitted or effective sets, see capabilities for details). Pass a capability
	// name such as "CAP_MKNOD", possibly prefixed with an exclamation mark to negate the check.
	ConditionCapability systemdconf.Value

	// Check whether the system has AC power, or is exclusively battery powered at the time of activation of the unit. This takes
	// a boolean argument. If set to "true", the condition will hold only if at least one AC connector of the system is connected
	// to a power source, or if no AC connectors are known. Conversely, if set to "false", the condition will hold only if there is
	// at least one AC connector known and all AC connectors are disconnected from a power source.
	ConditionACPower systemdconf.Value

	// Takes one of /var or /etc as argument, possibly prefixed with a "!" (to inverting the condition). This condition may be used
	// to conditionalize units on whether the specified directory requires an update because /usr's modification time is newer
	// than the stamp file .updated in the specified directory. This is useful to implement offline updates of the vendor operating
	// system resources in /usr that require updating of /etc or /var on the next following boot. Units making use of this condition
	// should order themselves before systemd-update-done.service, to make sure they run before the stamp file's modification
	// time gets reset indicating a completed update.
	ConditionNeedsUpdate systemdconf.Value

	// Takes a boolean argument. This condition may be used to conditionalize units on whether the system is booting up with an
	// unpopulated /etc directory (specifically: an /etc with no /etc/machine-id). This may be used to populate /etc on the first
	// boot after factory reset, or when a new system instance boots up for the first time.
	ConditionFirstBoot systemdconf.Value

	// Check for the exists of a file. If the specified absolute path name does not exist, the condition will fail. If the absolute
	// path name passed to ConditionPathExists= is prefixed with an exclamation mark ("!"), the test is negated, and the unit
	// is only started if the path does not exist.
	ConditionPathExists systemdconf.Value

	// ConditionPathExistsGlob= is similar to ConditionPathExists=, but checks for the existence of at least one file or directory
	// matching the specified globbing pattern.
	ConditionPathExistsGlob systemdconf.Value

	// ConditionPathIsDirectory= is similar to ConditionPathExists= but verifies that a certain path exists and is a directory.
	ConditionPathIsDirectory systemdconf.Value

	// ConditionPathIsSymbolicLink= is similar to ConditionPathExists= but verifies that a certain path exists and is a symbolic
	// link.
	ConditionPathIsSymbolicLink systemdconf.Value

	// ConditionPathIsMountPoint= is similar to ConditionPathExists= but verifies that a certain path exists and is a mount
	// point.
	ConditionPathIsMountPoint systemdconf.Value

	// ConditionPathIsReadWrite= is similar to ConditionPathExists= but verifies that the underlying file system is readable
	// and writable (i.e. not mounted read-only).
	ConditionPathIsReadWrite systemdconf.Value

	// ConditionDirectoryNotEmpty= is similar to ConditionPathExists= but verifies that a certain path exists and is a non-empty
	// directory.
	ConditionDirectoryNotEmpty systemdconf.Value

	// ConditionFileNotEmpty= is similar to ConditionPathExists= but verifies that a certain path exists and refers to a regular
	// file with a non-zero size.
	ConditionFileNotEmpty systemdconf.Value

	// ConditionFileIsExecutable= is similar to ConditionPathExists= but verifies that a certain path exists, is a regular
	// file, and marked executable.
	ConditionFileIsExecutable systemdconf.Value

	// ConditionUser= takes a numeric "UID", a UNIX user name, or the special value "@system". This condition may be used to check
	// whether the service manager is running as the given user. The special value "@system" can be used to check if the user id is
	// within the system user range. This option is not useful for system services, as the system manager exclusively runs as the
	// root user, and thus the test result is constant.
	ConditionUser systemdconf.Value

	// ConditionGroup= is similar to ConditionUser= but verifies that the service manager's real or effective group, or any
	// of its auxiliary groups, match the specified group or GID. This setting does not support the special value "@system".
	ConditionGroup systemdconf.Value

	// Verify that the given cgroup controller (eg. "cpu") is available for use on the system. For example, a particular controller
	// may not be available if it was disabled on the kernel command line with cgroup_disable=controller. Multiple controllers
	// may be passed with a space separating them; in this case the condition will only pass if all listed controllers are available
	// for use. Controllers unknown to systemd are ignored. Valid controllers are "cpu", "cpuacct", "io", "blkio", "memory",
	// "devices", and "pids".
	ConditionControlGroupController systemdconf.Value

	// Verify that the specified amount of system memory is available to the current system. Takes a memory size in bytes as argument,
	// optionally prefixed with a comparison operator "<", "<=", "=", "!=", ">=", ">". On bare-metal systems compares the amount
	// of physical memory in the system with the specified size, adhering to the specified comparison operator. In containers
	// compares the amount of memory assigned to the container instead.
	ConditionMemory systemdconf.Value

	// Verify that the specified number of CPUs is available to the current system. Takes a number of CPUs as argument, optionally
	// prefixed with a comparison operator "<", "<=", "=", "!=", ">=", ">". Compares the number of CPUs in the CPU affinity mask
	// configured of the service manager itself with the specified number, adhering to the specified comparison operator. On
	// physical systems the number of CPUs in the affinity mask of the service manager usually matches the number of physical CPUs,
	// but in special and virtual environments might differ. In particular, in containers the affinity mask usually matches
	// the number of CPUs assigned to the container and not the physically available ones.
	ConditionCPUs systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertArchitecture systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertVirtualization systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertHost systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertKernelCommandLine systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertKernelVersion systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertSecurity systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertCapability systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertACPower systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertNeedsUpdate systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertFirstBoot systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathExists systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathExistsGlob systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathIsDirectory systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathIsSymbolicLink systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathIsMountPoint systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertPathIsReadWrite systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertDirectoryNotEmpty systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertFileNotEmpty systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertFileIsExecutable systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertUser systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertGroup systemdconf.Value

	// Similar to the ConditionArchitecture=, ConditionVirtualization=, …, condition settings described above, these settings
	// add assertion checks to the start-up of the unit. However, unlike the conditions settings, any assertion setting that
	// is not met results in failure of the start job (which means this is logged loudly). Note that hitting a configured assertion
	// does not cause the unit to enter the "failed" state (or in fact result in any state change of the unit), it affects only the
	// job queued for it. Use assertion expressions for units that cannot operate when specific requirements are not met, and
	// when this is something the administrator or user should look into.
	AssertControlGroupController systemdconf.Value
}

// InstallSection represents installation information for the unit
type InstallSection struct {
	systemdconf.Section

	// A space-separated list of additional names this unit shall be installed under. The names listed here must have the same
	// suffix (i.e. type) as the unit filename. This option may be specified more than once, in which case all listed names are used.
	// At installation time, systemctl enable will create symlinks from these names to the unit filename. Note that not all unit
	// types support such alias names, and this setting is not supported for them. Specifically, mount, slice, swap, and automount
	// units do not support aliasing.
	Alias systemdconf.Value

	// This option may be used more than once, or a space-separated list of unit names may be given. A symbolic link is created in
	// the .wants/ or .requires/ directory of each of the listed units when this unit is installed by systemctl enable. This has
	// the effect that a dependency of type Wants= or Requires= is added from the listed unit to the current unit. The primary result
	// is that the current unit will be started when the listed unit is started. See the description of Wants= and Requires= in the
	// [Unit] section for details.
	//
	// WantedBy=foo.service in a service bar.service is mostly equivalent to Alias=foo.service.wants/bar.service in the
	// same file. In case of template units, systemctl enable must be called with an instance name, and this instance will be added
	// to the .wants/ or .requires/ list of the listed unit. E.g. WantedBy=getty.target in a service getty@.service will result
	// in systemctl enable getty@tty2.service creating a getty.target.wants/getty@tty2.service link to getty@.service.
	WantedBy systemdconf.Value

	// This option may be used more than once, or a space-separated list of unit names may be given. A symbolic link is created in
	// the .wants/ or .requires/ directory of each of the listed units when this unit is installed by systemctl enable. This has
	// the effect that a dependency of type Wants= or Requires= is added from the listed unit to the current unit. The primary result
	// is that the current unit will be started when the listed unit is started. See the description of Wants= and Requires= in the
	// [Unit] section for details.
	//
	// WantedBy=foo.service in a service bar.service is mostly equivalent to Alias=foo.service.wants/bar.service in the
	// same file. In case of template units, systemctl enable must be called with an instance name, and this instance will be added
	// to the .wants/ or .requires/ list of the listed unit. E.g. WantedBy=getty.target in a service getty@.service will result
	// in systemctl enable getty@tty2.service creating a getty.target.wants/getty@tty2.service link to getty@.service.
	RequiredBy systemdconf.Value

	// Additional units to install/deinstall when this unit is installed/deinstalled. If the user requests installation/deinstallation
	// of a unit with this option configured, systemctl enable and systemctl disable will automatically install/uninstall
	// units listed in this option as well.
	//
	// This option may be used more than once, or a space-separated list of unit names may be given.
	Also systemdconf.Value

	// In template unit files, this specifies for which instance the unit shall be enabled if the template is enabled without any
	// explicitly set instance. This option has no effect in non-template unit files. The specified string must be usable as instance
	// identifier.
	DefaultInstance systemdconf.Value
}
