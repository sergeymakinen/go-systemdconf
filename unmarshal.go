package systemdconf

import (
	"bytes"
	"errors"
	"reflect"
	"time"

	"github.com/sergeymakinen/go-systemdconf/v3/ast"
)

var unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()

// Unmarshaler is the interface implemented by types that can unmarshal a systemd Value of themselves.
// UnmarshalSystemd must copy the value if it wishes to retain the value after returning.
type Unmarshaler interface {
	UnmarshalSystemd(value Value) error
}

// Unmarshal decodes the systemd configuration file in data and stores the result
// in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an error.
//
// Unmarshal uses the inverse of the rules that
// Marshal uses, allocating slices and pointers as necessary.
//
// To unmarshal data into a pointer, Unmarshal unmarshals data into
// the value pointed at by the pointer. If the pointer is nil, Unmarshal
// allocates a new value for it to point to.
func Unmarshal(data []byte, v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		kind := "nil"
		if val.IsValid() && val.Kind() != reflect.Ptr {
			kind = val.Kind().String()
		}
		return errors.New("expected non-nil pointer, got " + kind)
	}
	f, err := ast.Parse(bytes.NewReader(data))
	if err != nil {
		return err
	}
	groups := map[string]*[]*Section{}
	var sections []*[]*Section
	for _, n := range f.Children {
		s, ok := n.(*ast.Section)
		if !ok {
			continue
		}
		section := &Section{name: s.Name}
		if group, ok := groups[s.Name]; ok {
			group2 := append(*group, section)
			*groups[s.Name] = group2
		} else {
			group2 := []*Section{section}
			groups[s.Name] = &group2
			sections = append(sections, &group2)
		}
		entries := map[string]*Entry{}
		for _, n := range s.Children {
			e, ok := n.(*ast.Entry)
			if !ok {
				continue
			}
			var entry *Entry
			if entry2, ok := entries[e.Key]; ok {
				entry = entry2
			} else {
				entry = &Entry{Key: e.Key}
				entries[entry.Key] = entry
				section.entries = append(section.entries, entry)
			}
			entry.Value = append(entry.Value, e.Value)
		}
	}
	if err = unmarshalFile(sections, unmarshalIndirect(val)); err != nil {
		return err
	}
	return nil
}

func unmarshalFile(sections []*[]*Section, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		return errors.New("expected struct, got " + v.Kind().String())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}
	for _, group := range sections {
		field := info.field((*group)[0].name)
		if field != nil {
			if err := unmarshalGroup(*group, unmarshalIndirect(v.FieldByIndex(field.idx))); err != nil {
				return err
			}
		} else if embedded := info.embedded(fileType); embedded != nil {
			file := unmarshalIndirect(v.FieldByIndex(embedded.idx)).Addr().Interface().(*File)
			for _, section := range *group {
				file.sections = append(file.sections, section)
			}
		}
	}
	return nil
}

func unmarshalGroup(group []*Section, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Struct:
		var entries []*Entry
		entryGroups := map[string]*Entry{}
		for _, section := range group {
			for _, entry := range section.entries {
				if entry2, ok := entryGroups[entry.Key]; ok {
					entry2.Value = append(entry2.Value, entry.Value...)
				} else {
					entryGroups[entry.Key] = entry
					entries = append(entries, entry)
				}
			}
		}
		return unmarshalSection(group[0].name, entries, v)
	case reflect.Array, reflect.Slice:
		n := len(group)
		if v.Kind() == reflect.Slice {
			if n > v.Cap() {
				slice := reflect.MakeSlice(v.Type(), n, n)
				reflect.Copy(slice, v)
				v.Set(slice)
			}
			v.SetLen(n)
		}
		if v.Len() < n {
			n = v.Len()
		}
		for i := 0; i < n; i++ {
			ev := unmarshalIndirect(v.Index(i))
			if err := unmarshalSection(group[i].name, group[i].entries, ev); err != nil {
				return err
			}
		}
		if v.Kind() == reflect.Array && len(group) < v.Len() {
			zero := reflect.Zero(v.Type().Elem())
			for i := len(group); i < v.Len(); i++ {
				v.Index(i).Set(zero)
			}
		}
		return nil
	default:
		return errors.New("expected struct, slice, or array, got " + v.Kind().String())
	}
}

func unmarshalSection(name string, entries []*Entry, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		return errors.New("expected struct, got " + v.Kind().String())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return err
	}
	for _, entry := range entries {
		field := info.field(entry.Key)
		if field != nil {
			if err := unmarshalEntry(entry, field, unmarshalIndirect(v.FieldByIndex(field.idx))); err != nil {
				return err
			}
		} else if embedded := info.embedded(sectionType); embedded != nil {
			section := unmarshalIndirect(v.FieldByIndex(embedded.idx)).Addr().Interface().(*Section)
			section.name = name
			section.entries = append(section.entries, entry)
		}
	}
	return nil
}

func unmarshalEntry(entry *Entry, field *fieldInfo, v reflect.Value) error {
	ft := indirectType(field.typ)
	if v.CanInterface() && ft.Implements(unmarshalerType) {
		if err := v.Interface().(Unmarshaler).UnmarshalSystemd(entry.Value); err != nil {
			return &FieldError{
				Struct: field.strct,
				Field:  field.name,
				Type:   field.typ,
				Err:    err,
			}
		}
		return nil
	}
	if v.CanAddr() {
		a := v.Addr()
		if a.CanInterface() && a.Type().Implements(unmarshalerType) {
			if err := a.Interface().(Unmarshaler).UnmarshalSystemd(entry.Value); err != nil {
				return &FieldError{
					Struct: field.strct,
					Field:  field.name,
					Type:   field.typ,
					Err:    err,
				}
			}
			return nil
		}
	}
	if ft == durationType {
		d, err := entry.Value.Duration()
		if err != nil {
			return &FieldError{
				Struct: field.strct,
				Field:  field.name,
				Type:   field.typ,
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
				Struct: field.strct,
				Field:  field.name,
				Type:   field.typ,
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
						Struct: field.strct,
						Field:  field.name,
						Type:   field.typ,
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
						Struct: field.strct,
						Field:  field.name,
						Type:   field.typ,
						Err:    err,
					}
				}
				ev.SetBool(b)
			case reflect.String:
				ev.SetString(entry.Value[i])
			default:
				return &FieldError{
					Struct: field.strct,
					Field:  field.name,
					Type:   field.typ,
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
			Struct: field.strct,
			Field:  field.name,
			Type:   field.typ,
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
