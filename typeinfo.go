package systemdconf

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

var typeCache sync.Map // map[reflect.Type]*typeInfo

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

func getRawTypeInfo(t reflect.Type) *typeInfo {
	info := &typeInfo{
		strct: t,
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
			info.embeds = append(info.embeds, &fieldInfo{
				idx: sf.Index,
				typ: st,
			})
			if st.Kind() == reflect.Struct {
				for _, fi := range getRawTypeInfo(st).fields {
					fi.idx = append([]int{i}, fi.idx...)
					info.fields = append(info.fields, fi)
				}
				continue
			}
		}
		fi := &fieldInfo{
			strct: t,
			name:  sf.Name,
			idx:   sf.Index,
			typ:   sf.Type,
		}
		if tag != "" {
			fi.name = tag
			fi.tag = true
		}
		info.fields = append(info.fields, fi)
	}
	return info
}

type fieldInfo struct {
	strct reflect.Type
	name  string
	tag   bool
	idx   []int
	typ   reflect.Type
}

type typeInfo struct {
	strct  reflect.Type
	fields []*fieldInfo
	embeds []*fieldInfo
}

func (t *typeInfo) embedded(typ reflect.Type) *fieldInfo {
	for _, f := range t.embeds {
		if f.typ == typ {
			return f
		}
	}
	return nil
}

func (t *typeInfo) field(name string) *fieldInfo {
	for _, f := range t.fields {
		if f.name == name {
			return f
		}
	}
	return nil
}

func (t *typeInfo) Len() int {
	return len(t.fields)
}

func (t *typeInfo) Less(i, j int) bool {
	// Like encoding.json does
	if len(t.fields[i].idx) != len(t.fields[j].idx) {
		return len(t.fields[i].idx) < len(t.fields[j].idx)
	}
	if t.fields[i].tag != t.fields[j].tag {
		return t.fields[i].tag
	}
	for k, v := range t.fields[i].idx {
		if k >= len(t.fields[j].idx) {
			return false
		}
		if v != t.fields[j].idx[k] {
			return v < t.fields[j].idx[k]
		}
	}
	return len(t.fields[i].idx) < len(t.fields[j].idx)
}

func (t *typeInfo) Swap(i, j int) {
	t.fields[i], t.fields[j] = t.fields[j], t.fields[i]
}

func (t *typeInfo) info(name string) (*fieldInfo, error) {
	ti := typeInfo{
		strct: t.strct,
	}
	for _, f := range t.fields {
		if f.name == name {
			ti.fields = append(ti.fields, f)
		}
	}
	sort.Sort(&ti)
	if len(ti.fields) > 1 && len(ti.fields[0].idx) == len(ti.fields[1].idx) && ti.fields[0].tag == ti.fields[1].tag {
		f1 := t.strct.FieldByIndex(ti.fields[0].idx)
		f2 := t.strct.FieldByIndex(ti.fields[1].idx)
		return nil, &FieldConflictError{
			Struct1: ti.fields[0].strct,
			Struct2: ti.fields[1].strct,
			Field1:  f1.Name,
			Tag1:    f1.Tag.Get("systemd"),
			Field2:  f2.Name,
			Tag2:    f2.Tag.Get("systemd"),
		}
	}
	return ti.fields[0], nil
}

func (t *typeInfo) normalize() error {
	var fields []*fieldInfo
	names := map[string]bool{}
	for _, f := range t.fields {
		if _, ok := names[f.name]; ok {
			continue
		}
		fi, err := t.info(f.name)
		if err != nil {
			return err
		}
		fields = append(fields, fi)
		names[f.name] = true
	}
	t.fields = fields
	return nil
}

// FieldConflictError represents an error when marshalling/unmarshalling a struct
// is not possible because field names conflict.
type FieldConflictError struct {
	Struct1, Struct2 reflect.Type
	Field1, Field2   string
	Tag1, Tag2       string
}

func (e *FieldConflictError) Error() string {
	return fmt.Sprintf("field %s.%s with tag %q conflicts with field %s.%s with tag %q", e.Struct1, e.Field1, e.Tag1, e.Struct2, e.Field2, e.Tag2)
}
