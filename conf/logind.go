// Code generated from systemd 257 by generatesdconf. DO NOT EDIT.

package conf

import "github.com/sergeymakinen/go-systemdconf/v3"

// LogindFile represents logind.conf, logind.conf.d — Login manager configuration files
// (see https://www.freedesktop.org/software/systemd/man/logind.conf.html for details).
type LogindFile struct {
	systemdconf.File

	Login LogindLoginSection // parameters of the systemd login manager, systemd-logind.service
}

// LogindLoginSection represents parameters of the systemd login manager, systemd-logind.service
// (see https://www.freedesktop.org/software/systemd/man/logind.conf.html#Options for details).
type LogindLoginSection struct {
	systemdconf.Section

	// Takes a positive integer. Configures how many virtual terminals (VTs) to allocate by default that, when switched to and
	// are previously unused, "autovt" services are automatically spawned on. These services are instantiated from the template
	// unit autovt@.service for the respective VT TTY name, for example, autovt@tty4.service. By default, autovt@.service
	// is linked to getty@.service. In other words, login prompts are started dynamically as the user switches to unused virtual
	// terminals. Hence, this parameter controls how many login "gettys" are available on the VTs. If a VT is already used by some
	// other subsystem (for example, a graphical login), this kind of activation will not be attempted. Note that the VT configured
	// in ReserveVT= is always subject to this kind of activation, even if it is not one of the VTs configured with the NAutoVTs=
	// directive. Defaults to 6. When set to 0, automatic spawning of "autovt" services is disabled.
	NAutoVTs systemdconf.Value

	// Takes a positive integer. Identifies one virtual terminal that shall unconditionally be reserved for autovt@.service
	// activation (see above). The VT selected with this option will be marked busy unconditionally, so that no other subsystem
	// will allocate it. This functionality is useful to ensure that, regardless of how many VTs are allocated by other subsystems,
	// one login "getty" is always available. Defaults to 6 (in other words, there will always be a "getty" available on Alt-F6.).
	// When set to 0, VT reservation is disabled.
	//
	// Added in version 190.
	ReserveVT systemdconf.Value

	// Takes a boolean argument. Configures whether the processes of a user should be killed when the user logs out. If true, the
	// scope unit corresponding to the session and all processes inside that scope will be terminated. If false, the scope is "abandoned",
	// see systemd.scope, and processes are not killed. Defaults to "yes", but see the options KillOnlyUsers= and KillExcludeUsers=
	// below.
	//
	// In addition to session processes, user process may run under the user manager unit user@.service. Depending on the linger
	// settings, this may allow users to run processes independent of their login sessions. See the description of enable-linger
	// in loginctl.
	//
	// Note that setting KillUserProcesses=yes will break tools like screen and tmux, unless they are moved out of the session
	// scope. See example in systemd-run.
	KillUserProcesses systemdconf.Value

	// These settings take space-separated lists of usernames that override the KillUserProcesses= setting. A user name may
	// be added to KillExcludeUsers= to exclude the processes in the session scopes of that user from being killed even if KillUserProcesses=yes
	// is set. If KillExcludeUsers= is not set, the "root" user is excluded by default. KillExcludeUsers= may be set to an empty
	// value to override this default. If a user is not excluded, KillOnlyUsers= is checked next. If this setting is specified,
	// only the processes in the session scopes of those users will be killed. Otherwise, users are subject to the KillUserProcesses=yes
	// setting.
	KillOnlyUsers systemdconf.Value

	// These settings take space-separated lists of usernames that override the KillUserProcesses= setting. A user name may
	// be added to KillExcludeUsers= to exclude the processes in the session scopes of that user from being killed even if KillUserProcesses=yes
	// is set. If KillExcludeUsers= is not set, the "root" user is excluded by default. KillExcludeUsers= may be set to an empty
	// value to override this default. If a user is not excluded, KillOnlyUsers= is checked next. If this setting is specified,
	// only the processes in the session scopes of those users will be killed. Otherwise, users are subject to the KillUserProcesses=yes
	// setting.
	KillExcludeUsers systemdconf.Value

	// Configures the action to take when the system is idle. Takes one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend",
	// "hibernate", "hybrid-sleep", "suspend-then-hibernate", "sleep", and "lock". Defaults to "ignore".
	//
	// Note that this requires that user sessions correctly report the idle status to the system. The system will execute the action
	// after all sessions report that they are idle, no idle inhibitor lock is active, and subsequently, the time configured with
	// IdleActionSec= (see below) has expired.
	//
	// Added in version 198.
	IdleAction systemdconf.Value

	// Configures the delay after which the action configured in IdleAction= (see above) is taken after the system is idle.
	//
	// Added in version 198.
	IdleActionSec systemdconf.Value

	// Specifies the maximum time a system shutdown or sleep request is delayed due to an inhibitor lock of type "delay" being active
	// before the inhibitor is ignored and the operation executes anyway. Defaults to 5.
	InhibitDelayMaxSec systemdconf.Value

	// Specifies how long to keep the user record and per-user service user@.service around for a user after they logged out fully.
	// If set to zero, the per-user service is terminated immediately when the last session of the user has ended. If this option
	// is configured to non-zero rapid logout/login cycles are sped up, as the user's service manager is not constantly restarted.
	// If set to "infinity" the per-user service for a user is never terminated again after first login, and continues to run until
	// system shutdown. Defaults to 10s.
	//
	// Added in version 240.
	UserStopDelaySec systemdconf.Value

	// Takes a list of sleep operations. Possible values are "suspend", "hibernate", "hybrid-sleep", and "suspend-then-hibernate".
	// Controls the candidate sleep operations for the "sleep" action. When "sleep" action is performed, the specified sleep
	// operations are checked in a fixed order ("suspend-then-hibernate" → "hybrid-sleep" → "suspend" → "hibernate"), and
	// the first one supported by the machine is used to put the system into sleep. Defaults to "suspend-then-hibernate suspend
	// hibernate".
	//
	// Added in version 256.
	SleepOperation systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandlePowerKey systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandlePowerKeyLongPress systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleRebootKey systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleRebootKeyLongPress systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleSuspendKey systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleSuspendKeyLongPress systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleHibernateKey systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleHibernateKeyLongPress systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleLidSwitch systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleLidSwitchExternalPower systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleLidSwitchDocked systemdconf.Value

	// Controls how logind shall handle the system power, reboot and sleep keys and the lid switch to trigger actions such as system
	// power-off, reboot or suspend. Can be one of "ignore", "poweroff", "reboot", "halt", "kexec", "suspend", "hibernate",
	// "hybrid-sleep", "suspend-then-hibernate", "sleep", "lock", and "factory-reset", "secure-attention-key". If "ignore",
	// systemd-logind will never handle these keys. If "lock", all running sessions will be screen-locked; otherwise, the specified
	// action will be taken in the respective event. Only input devices with the "power-switch" udev tag will be watched for key/lid
	// switch events.
	//
	// HandlePowerKey= defaults to "poweroff", HandleRebootKey= defaults to "reboot", HandleSuspendKey= defaults to "suspend",
	// HandleHibernateKey= defaults to "hibernate", HandlePowerKeyLongPress= defaults to "ignore", HandleRebootKeyLongPress=
	// defaults to "poweroff", HandleSuspendKeyLongPress= defaults to "hibernate", HandleHibernateKeyLongPress= defaults
	// to "ignore". HandleLidSwitch= defaults to "suspend". HandleLidSwitchExternalPower= is completely ignored by default
	// (for backwards compatibility) — an explicit value must be set before it will be used to determine behaviour. HandleLidSwitchDocked=
	// defaults to "ignore". If the system is inserted in a docking station, or if more than one display is connected, the action
	// specified by HandleLidSwitchDocked= occurs; if the system is on external power the action (if any) specified by HandleLidSwitchExternalPower=
	// occurs; otherwise the HandleLidSwitch= action occurs. HandleSecureAttentionKey= defaults to "secure-attention-key"
	//
	// A different application may disable logind's handling of system power and sleep keys and the lid switch by taking a low-level
	// inhibitor lock ("handle-power-key", "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key").
	// This is most commonly used by graphical desktop environments to take over suspend and hibernation handling, and to use
	// their own configuration mechanisms. If a low-level inhibitor lock is taken, logind will not take any action when that key
	// or switch is triggered and the Handle*= settings are irrelevant, except for HandleSecureAttentionKey=, which is always
	// handled since its addition in v257.
	//
	// Added in version 184.
	HandleSecureAttentionKey systemdconf.Value

	// Controls whether actions that systemd-logind takes when the power, reboot and sleep keys and the lid switch are triggered
	// are subject to high-level inhibitor locks ("shutdown", "reboot", "sleep", "idle"). Low level inhibitor locks ("handle-power-key",
	// "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key"), are always honored,
	// irrespective of this setting.
	//
	// These settings take boolean arguments. If "no", the inhibitor locks taken by applications are respected. If "yes", "shutdown",
	// "reboot" "sleep", and "idle" inhibitor locks are ignored. PowerKeyIgnoreInhibited=, SuspendKeyIgnoreInhibited=,
	// HibernateKeyIgnoreInhibited= and RebootKeyIgnoreInhibited= default to "no". LidSwitchIgnoreInhibited= defaults
	// to "yes". This means that when systemd-logind is handling events by itself (no low level inhibitor locks are taken by another
	// application), the lid switch does not respect suspend blockers by default, but the power and sleep keys do.
	//
	// Added in version 190.
	PowerKeyIgnoreInhibited systemdconf.Value

	// Controls whether actions that systemd-logind takes when the power, reboot and sleep keys and the lid switch are triggered
	// are subject to high-level inhibitor locks ("shutdown", "reboot", "sleep", "idle"). Low level inhibitor locks ("handle-power-key",
	// "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key"), are always honored,
	// irrespective of this setting.
	//
	// These settings take boolean arguments. If "no", the inhibitor locks taken by applications are respected. If "yes", "shutdown",
	// "reboot" "sleep", and "idle" inhibitor locks are ignored. PowerKeyIgnoreInhibited=, SuspendKeyIgnoreInhibited=,
	// HibernateKeyIgnoreInhibited= and RebootKeyIgnoreInhibited= default to "no". LidSwitchIgnoreInhibited= defaults
	// to "yes". This means that when systemd-logind is handling events by itself (no low level inhibitor locks are taken by another
	// application), the lid switch does not respect suspend blockers by default, but the power and sleep keys do.
	//
	// Added in version 190.
	SuspendKeyIgnoreInhibited systemdconf.Value

	// Controls whether actions that systemd-logind takes when the power, reboot and sleep keys and the lid switch are triggered
	// are subject to high-level inhibitor locks ("shutdown", "reboot", "sleep", "idle"). Low level inhibitor locks ("handle-power-key",
	// "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key"), are always honored,
	// irrespective of this setting.
	//
	// These settings take boolean arguments. If "no", the inhibitor locks taken by applications are respected. If "yes", "shutdown",
	// "reboot" "sleep", and "idle" inhibitor locks are ignored. PowerKeyIgnoreInhibited=, SuspendKeyIgnoreInhibited=,
	// HibernateKeyIgnoreInhibited= and RebootKeyIgnoreInhibited= default to "no". LidSwitchIgnoreInhibited= defaults
	// to "yes". This means that when systemd-logind is handling events by itself (no low level inhibitor locks are taken by another
	// application), the lid switch does not respect suspend blockers by default, but the power and sleep keys do.
	//
	// Added in version 190.
	HibernateKeyIgnoreInhibited systemdconf.Value

	// Controls whether actions that systemd-logind takes when the power, reboot and sleep keys and the lid switch are triggered
	// are subject to high-level inhibitor locks ("shutdown", "reboot", "sleep", "idle"). Low level inhibitor locks ("handle-power-key",
	// "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key"), are always honored,
	// irrespective of this setting.
	//
	// These settings take boolean arguments. If "no", the inhibitor locks taken by applications are respected. If "yes", "shutdown",
	// "reboot" "sleep", and "idle" inhibitor locks are ignored. PowerKeyIgnoreInhibited=, SuspendKeyIgnoreInhibited=,
	// HibernateKeyIgnoreInhibited= and RebootKeyIgnoreInhibited= default to "no". LidSwitchIgnoreInhibited= defaults
	// to "yes". This means that when systemd-logind is handling events by itself (no low level inhibitor locks are taken by another
	// application), the lid switch does not respect suspend blockers by default, but the power and sleep keys do.
	//
	// Added in version 190.
	LidSwitchIgnoreInhibited systemdconf.Value

	// Controls whether actions that systemd-logind takes when the power, reboot and sleep keys and the lid switch are triggered
	// are subject to high-level inhibitor locks ("shutdown", "reboot", "sleep", "idle"). Low level inhibitor locks ("handle-power-key",
	// "handle-suspend-key", "handle-hibernate-key", "handle-lid-switch", "handle-reboot-key"), are always honored,
	// irrespective of this setting.
	//
	// These settings take boolean arguments. If "no", the inhibitor locks taken by applications are respected. If "yes", "shutdown",
	// "reboot" "sleep", and "idle" inhibitor locks are ignored. PowerKeyIgnoreInhibited=, SuspendKeyIgnoreInhibited=,
	// HibernateKeyIgnoreInhibited= and RebootKeyIgnoreInhibited= default to "no". LidSwitchIgnoreInhibited= defaults
	// to "yes". This means that when systemd-logind is handling events by itself (no low level inhibitor locks are taken by another
	// application), the lid switch does not respect suspend blockers by default, but the power and sleep keys do.
	//
	// Added in version 190.
	RebootKeyIgnoreInhibited systemdconf.Value

	// Specifies a period of time after system startup or system resume in which systemd will hold off on reacting to lid events.
	// This is required for the system to properly detect any hotplugged devices so systemd can ignore lid events if external monitors,
	// or docks, are connected. If set to 0, systemd will always react immediately, possibly before the kernel fully probed all
	// hotplugged devices. This is safe, as long as you do not care for systemd to account for devices that have been plugged or unplugged
	// while the system was off. Defaults to 30s.
	//
	// Added in version 220.
	HoldoffTimeoutSec systemdconf.Value

	// Sets the size limit on the $XDG_RUNTIME_DIR runtime directory for each user who logs in. Takes a size in bytes, optionally
	// suffixed with the usual K, G, M, and T suffixes, to the base 1024 (IEC). Alternatively, a numerical percentage suffixed
	// by "%" may be specified, which sets the size limit relative to the amount of physical RAM. Defaults to 10%. Note that this
	// size is a safety limit only. As each runtime directory is a tmpfs file system, it will only consume as much memory as is needed.
	//
	// Added in version 211.
	RuntimeDirectorySize systemdconf.Value

	// Sets the limit on number of inodes for the $XDG_RUNTIME_DIR runtime directory for each user who logs in. Takes a number,
	// optionally suffixed with the usual K, G, M, and T suffixes, to the base 1024 (IEC). Defaults to RuntimeDirectorySize= divided
	// by 4096. Note that this size is a safety limit only. As each runtime directory is a tmpfs file system, it will only consume
	// as much memory as is needed.
	//
	// Added in version 246.
	RuntimeDirectoryInodesMax systemdconf.Value

	// Controls the maximum number of concurrent inhibitors to permit. Defaults to 8192 (8K).
	//
	// Added in version 230.
	InhibitorsMax systemdconf.Value

	// Controls the maximum number of concurrent user sessions to manage. Defaults to 8192 (8K). Depending on how the pam_systemd.so
	// module is included in the PAM stack configuration, further login sessions will either be refused, or permitted but not
	// tracked by systemd-logind.
	//
	// Added in version 230.
	SessionsMax systemdconf.Value

	// Controls whether System V and POSIX IPC objects belonging to the user shall be removed when the user fully logs out. Takes
	// a boolean argument. If enabled, the user may not consume IPC resources after the last of the user's sessions terminated.
	// This covers System V semaphores, shared memory and message queues, as well as POSIX shared memory and message queues. Note
	// that IPC objects of the root user and other system users are excluded from the effect of this setting. Defaults to "yes".
	//
	// Added in version 212.
	RemoveIPC systemdconf.Value

	// Specifies a timeout in seconds, or a time span value after which systemd-logind checks the idle state of all sessions. Every
	// session that is idle for longer than the timeout will be stopped. Note that this option does not apply to "greeter" or "lock-screen"
	// sessions. Defaults to "infinity" (systemd-logind is not checking the idle state of sessions). For details about the syntax
	// of time spans, see systemd.time.
	//
	// Added in version 252.
	StopIdleSessionSec systemdconf.Value

	// Specifies a default calendar event for scheduled shutdowns. So when using e.g. the command shutdown -r to reboot the system
	// without specifying a timeout, logind would use the configured calendar event instead. For details about the syntax of
	// calendar events, see systemd.time.
	//
	// Added in version 257.
	DesignatedMaintenanceTime systemdconf.Value
}
