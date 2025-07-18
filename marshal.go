package systemdconf

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/sergeymakinen/go-systemdconf/v3/ast"
)

var (
	durationType  = reflect.TypeOf((*time.Duration)(nil)).Elem()
	marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
	fileType      = reflect.TypeOf((*File)(nil)).Elem()
	sectionType   = reflect.TypeOf((*Section)(nil)).Elem()
)

// Marshaler is the interface implemented by types that can marshal themselves into a systemd Value.
type Marshaler interface {
	MarshalSystemd() (Value, error)
}

// FieldError represents an error when marshalling/unmarshalling struct fields.
type FieldError struct {
	Struct reflect.Type // type of the struct containing the field
	Field  string       // name of the field
	Type   reflect.Type // type of the field
	Err    error        // the actual error
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("failed to convert struct field %s.%s of type %s: %s", e.Struct, e.Field, e.Type, e.Err)
}

// Marshal encodes the systemd configuration file of v.
//
// 1) If the value v is not a struct (or pointer to struct), Marshal returns an error.
// Marshal encodes sections: each exported struct field becomes a section,
// using the field name as the section name, unless the field is omitted,
// and the field value as the section entries for the next step.
//
// If the struct embeds the File struct, any unknown or extra section will be extracted from it.
//
// 2) If the field value from the previous step is a struct (or pointer to struct),
// Marshal encodes entries from it: each exported struct field becomes an entry,
// using the field name as the section key name, unless the field is omitted,
// and the field value as the section value for the next step.
//
// If the struct embeds the Section struct, any unknown or extra value will be extracted from it.
//
// 3) If the section value implements the Marshaler interface
// and is not a nil pointer, Marshal calls its MarshalSystemd method
// to produce a Value.
// If no MarshalSystemd method is present, the value must be:
//   - boolean
//   - slice of booleans
//   - string
//   - slice of strings
//   - time.Duration
//   - slice of time.Durations
//   - pointer to type above
//
// The marshalling of each struct field can be customized by the format string
// stored under the "systemd" key in the struct field's tag.
// The format string gives the name of the field. If the field tag is "-", the field is always omitted.
//
// Anonymous struct fields are marshaled as if their inner exported fields
// were fields in the outer struct, subject to the usual Go visibility rules.
//
// Pointer values resolve as the value pointed to.
//
// Interface values resolve as the value contained in the interface.
func Marshal(v any) ([]byte, error) {
	val := indirect(reflect.ValueOf(v))
	if !val.IsValid() {
		return nil, errors.New("expected value, got nil")
	}
	file, err := marshalFile(val)
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	if err = (&ast.Printer{}).Fprint(&b, file); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func marshalFile(v reflect.Value) (*ast.File, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.New("expected struct, got " + v.Kind().String())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}
	file := &ast.File{}
	for _, field := range info.fields {
		sv := indirect(v.FieldByIndex(field.idx))
		if !sv.IsValid() {
			continue
		}
		sections, err := marshalSections(field, sv)
		if err != nil {
			return nil, err
		}
		for _, section := range sections {
			file.Children = append(file.Children, section)
		}
	}
	if embedded := info.embedded(fileType); embedded != nil {
		fileVal := indirect(v.FieldByIndex(embedded.idx))
		if fileVal.IsValid() {
			for _, section := range fileToAstSections(fileVal.Interface().(File)) {
				file.Children = append(file.Children, section)
			}
		}
	}
	return file, nil
}

func marshalSections(field *fieldInfo, v reflect.Value) ([]*ast.Section, error) {
	switch v.Kind() {
	case reflect.Struct:
		section, err := marshalSection(field, v)
		if err != nil {
			return nil, err
		}
		if len(section.Children) == 0 {
			return nil, nil
		}
		return []*ast.Section{section}, nil
	case reflect.Array, reflect.Slice:
		var sections []*ast.Section
		for i := 0; i < v.Len(); i++ {
			ev := indirect(v.Index(i))
			if !ev.IsValid() {
				continue
			}
			section, err := marshalSection(field, ev)
			if err != nil {
				return nil, err
			}
			if len(section.Children) == 0 {
				continue
			}
			sections = append(sections, section)
		}
		return sections, nil
	default:
		return nil, errors.New("expected struct, slice, or array, got " + v.Kind().String())
	}
}

func marshalSection(field *fieldInfo, v reflect.Value) (*ast.Section, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.New("expected struct, got " + v.Kind().String())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}
	section := &ast.Section{
		Name: field.name,
	}
	for _, field := range info.fields {
		sv := indirect(v.FieldByIndex(field.idx))
		if !sv.IsValid() {
			continue
		}
		entries, err := marshalEntries(field, sv)
		if err != nil {
			return nil, err
		}
		for _, e := range entries {
			section.Children = append(section.Children, e)
		}
	}
	if embedded := info.embedded(sectionType); embedded != nil {
		sectionVal := indirect(v.FieldByIndex(embedded.idx))
		if sectionVal.IsValid() {
			for _, entry := range sectionToAstEntries(sectionVal.Interface().(Section)) {
				section.Children = append(section.Children, entry)
			}
		}
	}
	return section, nil
}

func marshalEntries(field *fieldInfo, v reflect.Value) ([]*ast.Entry, error) {
	ft := indirectType(field.typ)
	if v.CanInterface() && ft.Implements(marshalerType) {
		value, err := v.Interface().(Marshaler).MarshalSystemd()
		if err != nil {
			return nil, &FieldError{
				Struct: field.strct,
				Field:  field.name,
				Type:   field.typ,
				Err:    err,
			}
		}
		var entries []*ast.Entry
		for _, s := range value {
			entries = append(entries, &ast.Entry{
				Key:   field.name,
				Value: s,
			})
		}
		return entries, nil
	}
	if ft == durationType {
		return []*ast.Entry{{
			Key:   field.name,
			Value: FormatDuration(v.Interface().(time.Duration)),
		}}, nil
	}
	switch ft.Kind() {
	case reflect.Bool:
		return []*ast.Entry{{
			Key:   field.name,
			Value: strconv.FormatBool(v.Bool()),
		}}, nil
	case reflect.String:
		return []*ast.Entry{{
			Key:   field.name,
			Value: v.String(),
		}}, nil
	case reflect.Array, reflect.Slice:
		var entries []*ast.Entry
		for i := 0; i < v.Len(); i++ {
			et := indirectType(ft.Elem())
			ev := indirect(v.Index(i))
			if !ev.IsValid() {
				continue
			}
			if et == durationType {
				entries = append(entries, &ast.Entry{
					Key:   field.name,
					Value: FormatDuration(ev.Interface().(time.Duration)),
				})
				continue
			}
			switch et.Kind() {
			case reflect.Bool:
				entries = append(entries, &ast.Entry{
					Key:   field.name,
					Value: strconv.FormatBool(ev.Bool()),
				})
			case reflect.String:
				entries = append(entries, &ast.Entry{
					Key:   field.name,
					Value: ev.String(),
				})
			default:
				return nil, &FieldError{
					Struct: field.strct,
					Field:  field.name,
					Type:   field.typ,
					Err:    errors.New("unsupported type"),
				}
			}
		}
		return entries, nil
	default:
		return nil, &FieldError{
			Struct: field.strct,
			Field:  field.name,
			Type:   field.typ,
			Err:    errors.New("unsupported type"),
		}
	}
}

func indirect(v reflect.Value) reflect.Value {
	for {
		if !v.IsValid() {
			return reflect.Value{}
		}
		typ := v.Type()
		switch typ.Kind() {
		case reflect.Interface, reflect.Ptr:
			if v.IsNil() {
				return reflect.Value{}
			}
			v = v.Elem()
		default:
			return v
		}
	}
}

func indirectType(typ reflect.Type) reflect.Type {
	for {
		switch typ.Kind() {
		case reflect.Ptr:
			typ = typ.Elem()
		default:
			return typ
		}
	}
}

func fileToAstSections(file File) []*ast.Section {
	var sections []*ast.Section
	for _, section := range file.sections {
		s := &ast.Section{
			Name: section.name,
		}
		for _, entry := range sectionToAstEntries(*section) {
			s.Children = append(s.Children, entry)
		}
		sections = append(sections, s)
	}
	return sections
}

func sectionToAstEntries(section Section) []*ast.Entry {
	var entries []*ast.Entry
	for _, entry := range section.entries {
		for _, s := range entry.Value {
			entries = append(entries, &ast.Entry{
				Key:   entry.Key,
				Value: s,
			})
		}
	}
	return entries
}
