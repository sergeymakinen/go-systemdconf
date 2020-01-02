package systemdconf

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sergeymakinen/go-systemdconf/ast"
)

var (
	durationType  = reflect.TypeOf((*time.Duration)(nil)).Elem()
	marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()
	fileType      = reflect.TypeOf((*File)(nil)).Elem()
	sectionType   = reflect.TypeOf((*Section)(nil)).Elem()
)

// Marshaler is the interface implemented by types that can marshal themselves into a systemd Value
type Marshaler interface {
	MarshalSystemd() (Value, error)
}

// FieldError represents an error when marshalling/unmarshalling struct fields
type FieldError struct {
	Struct reflect.Type // Type of the struct containing the field
	Field  string       // Name of the field
	Type   reflect.Type // Type of the field
	Err    error        // The actual error
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
// to produce a Value. If no MarshalSystemd method is present but the
// value implements encoding.TextMarshaler instead, Marshal calls
// its MarshalText method and uses the result as a single value.
// If no MarshalText method is present, the value must be:
// 	- boolean
// 	- slice of booleans
//	- string
//	- slice of strings
//	- time.Duration
//	- slice of time.Durations
//	- pointer to type above
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
// Interface values resolve as the value contained in the interface
func Marshal(v interface{}) ([]byte, error) {
	val := indirect(reflect.ValueOf(v))
	if !val.IsValid() {
		return nil, errors.New("expected value, got nil")
	}
	file, err := marshalFile(val)
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	if err = (&ast.Serializer{}).Serialize(file, &b); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func marshalFile(v reflect.Value) (*ast.File, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.Errorf("expected struct, got %s", v.Kind())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}
	file := &ast.File{}
	for _, field := range info.Fields {
		sv := indirect(v.FieldByIndex(field.Index))
		if !sv.IsValid() {
			continue
		}
		section, err := marshalSection(field, sv)
		if err != nil {
			return nil, err
		}
		file.Children = append(file.Children, section)
	}
	if embedded := info.Embedded(fileType); embedded != nil {
		fileVal := indirect(v.FieldByIndex(embedded.Index))
		if fileVal.IsValid() {
			for _, section := range fileToAstSections(fileVal.Interface().(File)) {
				file.Children = append(file.Children, section)
			}
		}
	}
	return file, nil
}

func marshalSection(field *fieldInfo, v reflect.Value) (*ast.Section, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.Errorf("expected struct, got %s", v.Kind())
	}
	info, err := getTypeInfo(v.Type())
	if err != nil {
		return nil, err
	}
	section := &ast.Section{
		Name: field.Name,
	}
	for _, field := range info.Fields {
		sv := indirect(v.FieldByIndex(field.Index))
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
	if embedded := info.Embedded(sectionType); embedded != nil {
		sectionVal := indirect(v.FieldByIndex(embedded.Index))
		if sectionVal.IsValid() {
			for _, entry := range sectionToAstEntries(sectionVal.Interface().(Section)) {
				section.Children = append(section.Children, entry)
			}
		}
	}
	return section, nil
}

func marshalEntries(field *fieldInfo, v reflect.Value) ([]*ast.Entry, error) {
	if v.CanInterface() && field.Type.Implements(marshalerType) {
		value, err := v.Interface().(Marshaler).MarshalSystemd()
		if err != nil {
			return nil, &FieldError{
				Struct: field.Struct,
				Field:  field.Name,
				Type:   field.Type,
				Err:    errors.Wrap(err, "cannot marshal"),
			}
		}
		var entries []*ast.Entry
		for _, s := range value {
			entries = append(entries, &ast.Entry{
				Key:   field.Name,
				Value: s,
			})
		}
		return entries, nil
	}
	if field.Type == durationType {
		return []*ast.Entry{{
			Key:   field.Name,
			Value: FormatDuration(v.Interface().(time.Duration)),
		}}, nil
	}
	switch field.Type.Kind() {
	case reflect.Bool:
		return []*ast.Entry{{
			Key:   field.Name,
			Value: strconv.FormatBool(v.Bool()),
		}}, nil
	case reflect.String:
		return []*ast.Entry{{
			Key:   field.Name,
			Value: v.String(),
		}}, nil
	case reflect.Array, reflect.Slice:
		var entries []*ast.Entry
		for i := 0; i < v.Len(); i++ {
			lt := indirectType(field.Type.Elem())
			lv := indirect(v.Index(i))
			if !lv.IsValid() {
				continue
			}
			if lt == durationType {
				entries = append(entries, &ast.Entry{
					Key:   field.Name,
					Value: FormatDuration(lv.Interface().(time.Duration)),
				})
				continue
			}
			switch lt.Kind() {
			case reflect.Bool:
				entries = append(entries, &ast.Entry{
					Key:   field.Name,
					Value: strconv.FormatBool(lv.Bool()),
				})
			case reflect.String:
				entries = append(entries, &ast.Entry{
					Key:   field.Name,
					Value: lv.String(),
				})
			default:
				return nil, &FieldError{
					Struct: field.Struct,
					Field:  field.Name,
					Type:   field.Type,
					Err:    errors.New("unsupported type"),
				}
			}
		}
		return entries, nil
	default:
		return nil, &FieldError{
			Struct: field.Struct,
			Field:  field.Name,
			Type:   field.Type,
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
