package systemdconf

import (
	"bytes"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/sergeymakinen/go-systemdconf/ast"
)

var unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()

// Unmarshaler is the interface implemented by types that can unmarshal a systemd Value of themselves.
// UnmarshalSystemd must copy the value if it wishes to retain the value after returning
type Unmarshaler interface {
	UnmarshalSystemd(value Value) error
}

// Unmarshal decodes the systemd configuration file in data and stores the result
// in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an error.
//
// Unmarshal uses the inverse of the rules that
// Marshal uses, allocating maps, slices, and pointers as necessary.
//
// To unmarshal data into a pointer, Unmarshal unmarshals data into
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
// allocates a new value for it to point to
func Unmarshal(data []byte, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("expected value or pointer to value, got nil")
	}
	var sections []*Section
	sectionNames := map[string]*Section{}
	type entryKey struct {
		Section, Entry string
	}
	entryNames := map[entryKey]*Entry{}
	f, err := ast.NewParser(bytes.NewReader(data)).Parse()
	if err != nil {
		return err
	}
	for _, n := range f.Children {
		if section, ok := n.(*ast.Section); ok {
			var s *Section
			if s2, ok := sectionNames[section.Name]; ok {
				s = s2
			} else {
				s = &Section{
					name: section.Name,
				}
				sectionNames[section.Name] = s
				sections = append(sections, s)
			}
			for _, n := range section.Children {
				if entry, ok := n.(*ast.Entry); ok {
					var e *Entry
					if e2, ok := entryNames[entryKey{section.Name, entry.Key}]; ok {
						e = e2
					} else {
						e = &Entry{
							Key: entry.Key,
						}
						entryNames[entryKey{section.Name, entry.Key}] = e
						s.entries = append(s.entries, e)
					}
					e.Value = append(e.Value, entry.Value)
				}
			}
		}
	}
	if err = unmarshalFile(sections, unmarshalIndirect(val)); err != nil {
		return err
	}
	return nil
}

func unmarshalFile(sections []*Section, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		return errors.Errorf("expected struct, got %s", v.Kind())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}
	for _, section := range sections {
		field := info.Field(section.name)
		if field != nil {
			if err := unmarshalSection(section, unmarshalIndirect(v.FieldByIndex(field.Index))); err != nil {
				return err
			}
		} else if embedded := info.Embedded(fileType); embedded != nil {
			fileStruct := unmarshalIndirect(v.FieldByIndex(embedded.Index)).Addr().Interface().(*File)
			fileStruct.sections = append(fileStruct.sections, section)
		}
	}
	return nil
}

func unmarshalSection(section *Section, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		return errors.Errorf("expected struct, got %s", v.Kind())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}
	for _, entry := range section.entries {
		field := info.Field(entry.Key)
		if field != nil {
			if err := unmarshalEntry(entry, field, unmarshalIndirect(v.FieldByIndex(field.Index))); err != nil {
				return err
			}
		} else if embedded := info.Embedded(sectionType); embedded != nil {
			sectionStruct := unmarshalIndirect(v.FieldByIndex(embedded.Index)).Addr().Interface().(*Section)
			sectionStruct.name = section.name
			sectionStruct.entries = append(sectionStruct.entries, entry)
		}
	}
	return nil
}

func unmarshalEntry(entry *Entry, field *fieldInfo, v reflect.Value) error {
	ft := indirectType(field.Type)
	if v.CanInterface() && ft.Implements(unmarshalerType) {
		if err := v.Interface().(Unmarshaler).UnmarshalSystemd(entry.Value); err != nil {
			return &FieldError{
				Struct: field.Struct,
				Field:  field.Name,
				Type:   field.Type,
				Err:    errors.Wrap(err, "cannot unmarshal"),
			}
		}
		return nil
	}
	if v.CanAddr() {
		a := v.Addr()
		if a.CanInterface() && a.Type().Implements(unmarshalerType) {
			if err := a.Interface().(Unmarshaler).UnmarshalSystemd(entry.Value); err != nil {
				return &FieldError{
					Struct: field.Struct,
					Field:  field.Name,
					Type:   field.Type,
					Err:    errors.Wrap(err, "cannot unmarshal"),
				}
			}
			return nil
		}
	}
	if ft == durationType {
		d, err := entry.Value.Duration()
		if err != nil {
			return &FieldError{
				Struct: field.Struct,
				Field:  field.Name,
				Type:   field.Type,
				Err:    err,
			}
		}
		v.Set(reflect.ValueOf(d))
		return nil
	}
	switch ft.Kind() {
	case reflect.Bool:
		b, err := entry.Value.Bool()
		if err != nil {
			return &FieldError{
				Struct: field.Struct,
				Field:  field.Name,
				Type:   field.Type,
				Err:    err,
			}
		}
		v.SetBool(b)
	case reflect.String:
		v.SetString(entry.Value.String())
	case reflect.Array, reflect.Slice:
		n := len(entry.Value)
		if ft.Kind() == reflect.Slice {
			if n > v.Cap() {
				slice := reflect.MakeSlice(ft, n, n)
				reflect.Copy(slice, v)
				v.Set(slice)
			}
			v.SetLen(n)
		}
		if v.Len() < n {
			n = v.Len()
		}
		for i := 0; i < n; i++ {
			et := indirectType(ft.Elem())
			ev := unmarshalIndirect(v.Index(i))
			if et == durationType {
				d, err := ParseDuration(entry.Value[i], time.Second)
				if err != nil {
					return &FieldError{
						Struct: field.Struct,
						Field:  field.Name,
						Type:   field.Type,
						Err:    err,
					}
				}
				ev.Set(reflect.ValueOf(d))
				continue
			}
			switch et.Kind() {
			case reflect.Bool:
				b, err := ParseBool(entry.Value[i])
				if err != nil {
					return &FieldError{
						Struct: field.Struct,
						Field:  field.Name,
						Type:   field.Type,
						Err:    err,
					}
				}
				ev.SetBool(b)
			case reflect.String:
				ev.SetString(entry.Value[i])
			default:
				return &FieldError{
					Struct: field.Struct,
					Field:  field.Name,
					Type:   field.Type,
					Err:    errors.New("unsupported type"),
				}
			}
		}
		if ft.Kind() == reflect.Array && len(entry.Value) < v.Len() {
			zero := reflect.Zero(ft.Elem())
			for i := len(entry.Value); i < v.Len(); i++ {
				v.Index(i).Set(zero)
			}
		}
	default:
		return &FieldError{
			Struct: field.Struct,
			Field:  field.Name,
			Type:   field.Type,
			Err:    errors.New("unsupported type"),
		}
	}
	return nil
}

func unmarshalIndirect(v reflect.Value) reflect.Value {
	var done bool
	for !done {
		switch v.Kind() {
		case reflect.Interface:
			done = true
			if !v.IsNil() {
				e := v.Elem()
				if e.Kind() == reflect.Ptr && !e.IsNil() {
					v = e
					done = false
				}
			}
		case reflect.Ptr:
			if v.IsNil() {
				v.Set(reflect.New(v.Type().Elem()))
			}
			v = v.Elem()
		default:
			done = true
		}
	}
	return v
}
