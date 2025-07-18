//go:generate go run github.com/sergeymakinen/go-systemdconf/v3/cmd/internal/generatesdconf cmd/internal/generatesdconf/config.yml .

// Package systemdconf implements encoding and decoding of systemd configuration files
// as defined by systemd.syntax (see https://www.freedesktop.org/software/systemd/man/systemd.syntax.html for details).
//
// The mapping between systemd and Go is described
// in the documentation for the Marshal and Unmarshal functions.
package systemdconf

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	whitespaceChars               = " \t\n\r"
	infinity        time.Duration = 1<<63 - 1
	day                           = 24 * time.Hour
	week                          = 7 * day
	month                         = 2629800 * time.Second
	year                          = 31557600 * time.Second
)

var (
	ErrInvalidBool     = errors.New("invalid systemd boolean")
	ErrInvalidDuration = errors.New("invalid systemd duration")

	parseDurationMap = []struct {
		String   string
		Duration time.Duration
	}{
		{"seconds", time.Second},
		{"second", time.Second},
		{"sec", time.Second},
		{"s", time.Second},
		{"minutes", time.Minute},
		{"minute", time.Minute},
		{"min", time.Minute},
		{"months", month},
		{"month", month},
		{"M", month},
		{"msec", time.Millisecond},
		{"ms", time.Millisecond},
		{"m", time.Minute},
		{"hours", time.Hour},
		{"hour", time.Hour},
		{"hr", time.Hour},
		{"h", time.Hour},
		{"days", day},
		{"day", day},
		{"d", day},
		{"weeks", week},
		{"week", week},
		{"w", week},
		{"years", year},
		{"year", year},
		{"y", year},
		{"usec", time.Microsecond},
		{"us", time.Microsecond},
		{"µs", time.Microsecond},
	}
	formatDurationMap = []struct {
		String   string
		Duration time.Duration
	}{
		{"y", year},
		{"month", month},
		{"w", week},
		{"d", day},
		{"h", time.Hour},
		{"min", time.Minute},
		{"s", time.Second},
		{"ms", time.Millisecond},
		{"us", time.Microsecond},
	}
)

// ParseBool returns the boolean value represented by the string as defined in systemd.syntax.
func ParseBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "1", "yes", "y", "true", "t", "on":
		return true, nil
	case "0", "no", "n", "false", "f", "off":
		return false, nil
	default:
		return false, ErrInvalidBool
	}
}

// ParseDuration returns the time.Duration value represented by the string as defined in systemd.time (time spans).
func ParseDuration(s string, def time.Duration) (time.Duration, error) {
	s = strings.Trim(s, whitespaceChars)
	if s == "infinity" {
		return infinity, nil
	}
	var (
		d, m time.Duration
		x    int64
		rat  string
		err  error
	)
	for {
		s = strings.TrimLeft(s, whitespaceChars)
		if s == "" {
			if d == 0 {
				return 0, ErrInvalidDuration
			}
			break
		}
		if strings.HasPrefix(s, "-") {
			return 0, ErrInvalidDuration
		}
		if strings.HasPrefix(s, "+") {
			s = strings.TrimPrefix(s, "+")
		}
		x, rat, s, err = extractNumber(s)
		if err != nil {
			return 0, ErrInvalidDuration
		}
		m, s, err = extractUnit(s, def)
		if err != nil {
			return 0, ErrInvalidDuration
		}
		if time.Duration(x) >= infinity/m {
			return 0, ErrInvalidDuration
		}
		if time.Duration(x)*m >= infinity-d {
			return 0, ErrInvalidDuration
		}
		d += time.Duration(x) * m
		if rat != "" {
			m /= 10
			for _, r := range rat {
				ratd := time.Duration(r-'0') * m
				if ratd >= infinity-d {
					return 0, ErrInvalidDuration
				}
				d += ratd
				m /= 10
			}
		}
	}
	return d, nil
}

func extractNumber(s string) (x int64, rat, rem string, err error) {
	var (
		isRat          bool
		xIndex, rIndex int
	)
	for i, r := range s {
		if r == '.' && !isRat {
			xIndex = i
			isRat = true
		} else if unicode.IsDigit(r) {
			if isRat {
				rIndex = i + 1
			} else {
				xIndex = i + 1
			}
		} else {
			break
		}
	}
	if xIndex > 0 {
		if x, err = strconv.ParseInt(s[:xIndex], 10, 64); err != nil {
			return
		}
		rem = s[xIndex:]
	} else if rIndex == 0 {
		err = ErrInvalidDuration
		return
	}
	if rIndex > 0 {
		rat = s[xIndex+1 : rIndex]
		rem = s[rIndex:]
	}
	return
}

func extractUnit(s string, def time.Duration) (d time.Duration, rem string, err error) {
	d = def
	rem = strings.TrimLeft(s, whitespaceChars)
	for _, v := range parseDurationMap {
		if strings.HasPrefix(rem, v.String) {
			d = v.Duration
			rem = strings.TrimPrefix(rem, v.String)
			return
		}
	}
	if rem != "" && rem == s {
		err = ErrInvalidDuration
	}
	return
}

// FormatDuration returns a string representing the duration as defined in systemd.time (time spans).
func FormatDuration(d time.Duration) string {
	if d == infinity {
		return "infinity"
	}
	if d <= 0 {
		return "0"
	}
	var b strings.Builder
	for _, v := range formatDurationMap {
		if d <= 0 {
			break
		}
		if d < v.Duration {
			continue
		}
		div := d / v.Duration
		rem := d % v.Duration
		if b.Len() > 0 {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprintf("%d%s", div, v.String))
		d = rem
	}
	return b.String()
}

// File embeds unknown and extra sections of a systemd configuration file.
type File struct {
	sections []*Section
}

// Unknown returns unknown sections of a systemd configuration file.
func (f *File) Unknown() []*Section {
	var sections []*Section
	for _, section := range f.sections {
		if !isExtra(section.name) {
			sections = append(sections, section)
		}
	}
	return sections
}

// Extra returns extra sections of a systemd configuration file.
func (f *File) Extra() []*Section {
	var sections []*Section
	for _, section := range f.sections {
		if isExtra(section.name) {
			sections = append(sections, section)
		}
	}
	return sections
}

// Section embeds unknown and extra entries of a systemd configuration file section.
type Section struct {
	name    string
	entries []*Entry
}

// Name returns a section name.
func (s *Section) Name() string { return s.name }

// Unknown returns unknown entries of the section.
func (s *Section) Unknown() []*Entry {
	var entries []*Entry
	for _, entry := range s.entries {
		if !isExtra(entry.Key) {
			entries = append(entries, entry)
		}
	}
	return entries
}

// Extra returns extra entries of the section.
func (s *Section) Extra() []*Entry {
	var entries []*Entry
	for _, entry := range s.entries {
		if isExtra(entry.Key) {
			entries = append(entries, entry)
		}
	}
	return entries
}

func isExtra(name string) bool {
	return len(name) > 2 && strings.HasPrefix(name, "X-")
}

// Entry represents a key-value configuration entry used by File and Section.
type Entry struct {
	Key   string // key name
	Value Value  // entry value
}

// Value represents a generic systemd value as a list of strings.
type Value []string

// Bool returns a first defined value parsed as a boolean.
func (v Value) Bool() (bool, error) { return ParseBool(v.String()) }

// BoolSlice returns a slice of defined values parsed as booleans.
func (v Value) BoolSlice() ([]bool, error) {
	var list []bool
	for _, val := range v {
		b, err := ParseBool(val)
		if err != nil {
			return nil, err
		}
		list = append(list, b)
	}
	return list, nil
}

// Duration returns a first defined value parsed as a time.Duration.
func (v Value) Duration() (time.Duration, error) { return ParseDuration(v.String(), time.Second) }

// DurationSlice returns a slice of defined values parsed as time.Durations.
func (v Value) DurationSlice() ([]time.Duration, error) {
	var list []time.Duration
	for _, val := range v {
		d, err := ParseDuration(val, time.Second)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}

// String returns a first defined value, or an empty string if the value does not exist.
func (v Value) String() string {
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
