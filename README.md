# systemdconf

[![Test](https://github.com/sergeymakinen/go-systemdconf/actions/workflows/test.yml/badge.svg)](https://github.com/sergeymakinen/go-systemdconf/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sergeymakinen/go-systemdconf.svg)](https://pkg.go.dev/github.com/sergeymakinen/go-systemdconf/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/sergeymakinen/go-systemdconf)](https://goreportcard.com/report/github.com/sergeymakinen/go-systemdconf)
[![codecov](https://codecov.io/gh/sergeymakinen/go-systemdconf/branch/master/graph/badge.svg)](https://codecov.io/gh/sergeymakinen/go-systemdconf)

Package systemdconf implements encoding and decoding of systemd configuration files
as defined by systemd.syntax (see https://www.freedesktop.org/software/systemd/man/systemd.syntax.html for details).

The mapping between systemd and Go is described
in the documentation for the Marshal and Unmarshal functions

## Installation

Use go get:

```bash
go get github.com/sergeymakinen/go-systemdconf/v3
```

Then import the package into your own code:

```go
import "github.com/sergeymakinen/go-systemdconf/v3"
```


## Example

### Marshal

```go
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
// Output: 
// [Unit]
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
```

### Unmarshal

```go
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
fmt.Println(b)
// Output: 
// Simple firewall
// true
```

## License

BSD 3-Clause
