// Package systemdconf implements encoding and decoding of systemd configuration files
// as defined by systemd.syntax (see https://www.freedesktop.org/software/systemd/man/systemd.syntax.html for details).
//
// The mapping between systemd and Go is described
// in the documentation for the Marshal and Unmarshal functions
package systemdconf

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/pkg/errors"
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

	typeCache        sync.Map // map[reflect.Type]*typeInfo
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

// ParseBool returns the boolean value represented by the string as defined in systemd.syntax
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

// ParseDuration returns the time.Duration value represented by the string as defined in systemd.time (time spans)
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

// FormatDuration returns a string representing the duration as defined in systemd.time (time spans)
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

// File embeds unknown and extra sections of a systemd configuration file
type File struct {
	sections []*Section
}

// Unknown returns unknown sections of a systemd configuration file
func (f *File) Unknown() []*Section {
	var sections []*Section
	for _, section := range f.sections {
		if !isExtra(section.name) {
			sections = append(sections, section)
		}
	}
	return sections
}

// Extra returns extra sections of a systemd configuration file
func (f *File) Extra() []*Section {
	var sections []*Section
	for _, section := range f.sections {
		if isExtra(section.name) {
			sections = append(sections, section)
		}
	}
	return sections
}

// Section embeds unknown and extra entries of a systemd configuration file section
type Section struct {
	name    string
	entries []*Entry
}

// Name returns a section name
func (s *Section) Name() string { return s.name }

// Unknown returns unknown entries of the section
func (s *Section) Unknown() []*Entry {
	var entries []*Entry
	for _, entry := range s.entries {
		if !isExtra(entry.Key) {
			entries = append(entries, entry)
		}
	}
	return entries
}

// Extra returns extra entries of the section
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

// Entry represents a key-value configuration entry used by File and Section
type Entry struct {
	Key   string // Key name
	Value Value  // Entry value
}

// Value represents a generic systemd value as list of strings
type Value []string

// Bool returns a first defined value parsed as a boolean
func (v Value) Bool() (bool, error) { return ParseBool(v.String()) }

// BoolSlice returns a slice of defined values parsed as booleans
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

// Duration returns a first defined value parsed as a time.Duration
func (v Value) Duration() (time.Duration, error) { return ParseDuration(v.String(), time.Second) }

// DurationSlice returns a slice of defined values parsed as time.Durations
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

// String returns a first defined value, or an empty string if the value does not exist
func (v Value) String() string {
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

// FieldConflictError represents an error when marshalling/unmarshalling a struct
// is not possible because field names conflict
type FieldConflictError struct {
	Struct1, Struct2 reflect.Type
	Field1, Field2   string
	Tag1, Tag2       string
}

func (e *FieldConflictError) Error() string {
	return fmt.Sprintf("field %s.%s with tag %q conflicts with field %s.%s with tag %q", e.Struct1, e.Field1, e.Tag1, e.Struct2, e.Field2, e.Tag2)
}

type fieldInfo struct {
	Struct reflect.Type
	Name   string
	Tag    bool
	Index  []int
	Type   reflect.Type
}

type typeInfo struct {
	Struct reflect.Type
	Fields []*fieldInfo
	Embeds []*fieldInfo
}

func (t *typeInfo) Embedded(typ reflect.Type) *fieldInfo {
	for _, f := range t.Embeds {
		if f.Type == typ {
			return f
		}
	}
	return nil
}

func (t *typeInfo) Field(name string) *fieldInfo {
	for _, f := range t.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (t *typeInfo) Len() int {
	return len(t.Fields)
}

func (t *typeInfo) Less(i, j int) bool {
	// Like encoding.json does
	if len(t.Fields[i].Index) != len(t.Fields[j].Index) {
		return len(t.Fields[i].Index) < len(t.Fields[j].Index)
	}
	if t.Fields[i].Tag != t.Fields[j].Tag {
		return t.Fields[i].Tag
	}
	for k, v := range t.Fields[i].Index {
		if k >= len(t.Fields[j].Index) {
			return false
		}
		if v != t.Fields[j].Index[k] {
			return v < t.Fields[j].Index[k]
		}
	}
	return len(t.Fields[i].Index) < len(t.Fields[j].Index)
}

func (t *typeInfo) Swap(i, j int) {
	t.Fields[i], t.Fields[j] = t.Fields[j], t.Fields[i]
}

func (t *typeInfo) field(name string) (*fieldInfo, error) {
	ti := typeInfo{
		Struct: t.Struct,
	}
	for _, f := range t.Fields {
		if f.Name == name {
			ti.Fields = append(ti.Fields, f)
		}
	}
	sort.Sort(&ti)
	if len(ti.Fields) > 1 && len(ti.Fields[0].Index) == len(ti.Fields[1].Index) && ti.Fields[0].Tag == ti.Fields[1].Tag {
		f1 := t.Struct.FieldByIndex(ti.Fields[0].Index)
		f2 := t.Struct.FieldByIndex(ti.Fields[1].Index)
		return nil, &FieldConflictError{
			Struct1: ti.Fields[0].Struct,
			Struct2: ti.Fields[1].Struct,
			Field1:  f1.Name,
			Tag1:    f1.Tag.Get("systemd"),
			Field2:  f2.Name,
			Tag2:    f2.Tag.Get("systemd"),
		}
	}
	return ti.Fields[0], nil
}

func (t *typeInfo) normalize() error {
	var fields []*fieldInfo
	names := map[string]bool{}
	for _, f := range t.Fields {
		if _, ok := names[f.Name]; ok {
			continue
		}
		fi, err := t.field(f.Name)
		if err != nil {
			return err
		}
		fields = append(fields, fi)
		names[f.Name] = true
	}
	t.Fields = fields
	return nil
}

func getRawTypeInfo(t reflect.Type) *typeInfo {
	info := &typeInfo{
		Struct: t,
	}
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("systemd")
		if (sf.PkgPath != "" && !sf.Anonymous) || tag == "-" {
			continue
		}
		if sf.Anonymous {
			st := sf.Type
			if st.Kind() == reflect.Ptr {
				st = st.Elem()
			}
			info.Embeds = append(info.Embeds, &fieldInfo{
				Index: sf.Index,
				Type:  st,
			})
			if st.Kind() == reflect.Struct {
				for _, fi := range getRawTypeInfo(st).Fields {
					fi.Index = append([]int{i}, fi.Index...)
					info.Fields = append(info.Fields, fi)
				}
				continue
			}
		}
		fi := &fieldInfo{
			Struct: t,
			Name:   sf.Name,
			Index:  sf.Index,
			Type:   sf.Type,
		}
		if tag != "" {
			fi.Name = tag
			fi.Tag = true
		}
		info.Fields = append(info.Fields, fi)
	}
	return info
}

func getTypeInfo(t reflect.Type) (*typeInfo, error) {
	if f, ok := typeCache.Load(t); ok {
		return f.(*typeInfo), nil
	}
	info := getRawTypeInfo(t)
	if err := info.normalize(); err != nil {
		return nil, err
	}
	f, _ := typeCache.LoadOrStore(t, info)
	return f.(*typeInfo), nil
}
