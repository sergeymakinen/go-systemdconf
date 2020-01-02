// To avoid a cyclic dependency with systemdconf <-> go-systemdconf/unit

package systemdconf_test

import (
	"fmt"

	"github.com/sergeymakinen/go-systemdconf"
	"github.com/sergeymakinen/go-systemdconf/unit"
)

func ExampleMarshal() {
	service := unit.ServiceFile{
		Unit: unit.UnitSection{
			Description: systemdconf.Value{"Simple firewall"},
		},
		Service: unit.ServiceSection{
			Type:            systemdconf.Value{"oneshot"},
			RemainAfterExit: systemdconf.Value{"yes"},
			ExecStart: systemdconf.Value{
				"/usr/local/sbin/simple-firewall-start1",
				"/usr/local/sbin/simple-firewall-start2",
			},
			ExecStop: systemdconf.Value{"/usr/local/sbin/simple-firewall-stop"},
		},
		Install: unit.InstallSection{
			WantedBy: systemdconf.Value{"multi-user.target"},
		},
	}
	b, _ := systemdconf.Marshal(service)

	fmt.Println(string(b))
	// Output: [Unit]
	// Description=Simple firewall
	//
	// [Service]
	// Type=oneshot
	// RemainAfterExit=yes
	// ExecStart=/usr/local/sbin/simple-firewall-start1
	// ExecStart=/usr/local/sbin/simple-firewall-start2
	// ExecStop=/usr/local/sbin/simple-firewall-stop
	//
	// [Install]
	// WantedBy=multi-user.target
}

func ExampleUnmarshal() {
	file := `[Unit]
Description=Simple firewall

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/local/sbin/simple-firewall-start
ExecStop=/usr/local/sbin/simple-firewall-stop

[Install]
WantedBy=multi-user.target`

	var service unit.ServiceFile
	systemdconf.Unmarshal([]byte(file), &service)

	fmt.Println(service.Unit.Description)
	b, _ := service.Service.RemainAfterExit.Bool()
	if b {
		fmt.Println(b)
	}
	// Output: Simple firewall
	// true
}
