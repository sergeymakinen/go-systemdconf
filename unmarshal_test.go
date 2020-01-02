package systemdconf

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

type fileTest struct {
	File
	Section1 *struct {
		*Section
		Key1, Key2 Value
	}
	Section3 struct {
		Key Value
	}
	XSection1 struct {
		Key  Value
		XKey Value `systemd:"X-Key"`
	} `systemd:"X-Section1"`
	XSection2 struct {
		Key  Value
		XKey Value `systemd:"X-Key"`
	} `systemd:"-"`
	ActualXSection2 struct {
		Section
		Key  Value
		XKey Value `systemd:"X-Key"`
	} `systemd:"X-Section2"`
	Bool struct {
		True, False bool
		Key         []bool
	}
	Duration struct {
		Key []time.Duration
	}
}

func TestUnmarshalFile(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/unmarshal.conf")
	if err != nil {
		panic("Error reading testdata/unmarshal.conf: " + err.Error())
	}
	expected := fileTest{
		File: File{
			sections: []*Section{
				{
					name: "Section2",
					entries: []*Entry{
						{
							Key:   "Key1",
							Value: Value{"value 1"},
						},
						{
							Key:   "Key2",
							Value: Value{"value 2", "value 3"},
						},
					},
				},
				{
					name: "X-Section3",
					entries: []*Entry{
						{
							Key:   "X-Key",
							Value: Value{"value 2", "value 3"},
						},
						{
							Key:   "Key",
							Value: Value{"value 1"},
						},
					},
				},
			},
		},
		Section1: &struct {
			*Section
			Key1, Key2 Value
		}{
			Section: &Section{
				name: "Section1",
				entries: []*Entry{
					{
						Key:   "X-Key1",
						Value: Value{"value 3"},
					},
					{
						Key:   "X-Key2",
						Value: Value{"value 4"},
					},
				},
			},
			Key1: Value{"value 1"},
			Key2: Value{"value 2"},
		},
		Section3: struct {
			Key Value
		}{
			Key: Value{"value 1"},
		},
		XSection1: struct {
			Key  Value
			XKey Value `systemd:"X-Key"`
		}{
			Key:  Value{"value 1"},
			XKey: Value{"value 2"},
		},
		XSection2: struct {
			Key  Value
			XKey Value `systemd:"X-Key"`
		}{},
		ActualXSection2: struct {
			Section
			Key  Value
			XKey Value `systemd:"X-Key"`
		}{
			Key:  Value{"value 1"},
			XKey: Value{"value 2"},
		},
		Bool: struct {
			True, False bool
			Key         []bool
		}{
			True:  true,
			False: false,
			Key: []bool{
				true,
				true,
				true,
				true,
				true,
				true,
				false,
				false,
				false,
				false,
				false,
				false,
			},
		},
		Duration: struct {
			Key []time.Duration
		}{
			Key: []time.Duration{
				3*time.Hour + 3*time.Minute,
				3 * time.Second,
				3 * time.Second,
				3 * time.Second,
				3 * time.Second,
				3 * time.Minute,
				3 * time.Minute,
				3 * time.Minute,
				3 * month,
				3 * month,
				3 * month,
				3 * time.Millisecond,
				3 * time.Millisecond,
				3 * time.Minute,
				3 * time.Hour,
				3 * time.Hour,
				3 * time.Hour,
				3 * time.Hour,
				3 * day,
				3 * day,
				3 * day,
				3 * week,
				3 * week,
				3 * week,
				3 * year,
				3 * year,
				3 * year,
				3 * time.Microsecond,
				3 * time.Microsecond,
				3 * time.Microsecond,
				infinity,
			},
		},
	}
	var s fileTest
	err = Unmarshal(b, &s)
	if err != nil {
		t.Fatalf("Unmarshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(expected, s, cmp.AllowUnexported(File{}, Section{})); diff != "" {
		t.Errorf("Unmarshal() mismatch (-want +got):\n%s", diff)
	}
	expectedSections := []*Section{
		{
			name: "Section2",
			entries: []*Entry{
				{
					Key:   "Key1",
					Value: Value{"value 1"},
				},
				{
					Key:   "Key2",
					Value: Value{"value 2", "value 3"},
				},
			},
		},
	}
	if diff := cmp.Diff(expectedSections, s.Unknown(), cmp.AllowUnexported(Section{})); diff != "" {
		t.Errorf("File.Unknown() mismatch (-want +got):\n%s", diff)
	}
	expectedEntries := []*Entry{
		{
			Key:   "Key1",
			Value: Value{"value 1"},
		},
		{
			Key:   "Key2",
			Value: Value{"value 2", "value 3"},
		},
	}
	if diff := cmp.Diff(expectedEntries, s.Unknown()[0].Unknown(), nil); diff != "" {
		t.Errorf("Section.Unknown() mismatch (-want +got):\n%s", diff)
	}
	expectedSections = []*Section{
		{
			name: "X-Section3",
			entries: []*Entry{
				{
					Key:   "X-Key",
					Value: Value{"value 2", "value 3"},
				},
				{
					Key:   "Key",
					Value: Value{"value 1"},
				},
			},
		},
	}
	if diff := cmp.Diff(expectedSections, s.Extra(), cmp.AllowUnexported(Section{})); diff != "" {
		t.Errorf("File.Extra() mismatch (-want +got):\n%s", diff)
	}
	expectedEntries = []*Entry{
		{
			Key:   "X-Key",
			Value: Value{"value 2", "value 3"},
		},
	}
	if diff := cmp.Diff(expectedEntries, s.Extra()[0].Extra(), nil); diff != "" {
		t.Errorf("Section.Extra() mismatch (-want +got):\n%s", diff)
	}
}

func TestUnmarshalShouldFail(t *testing.T) {
	file := []byte(`[Section]
Key=value
`)
	tests := []struct {
		Name string
		V    interface{}
	}{
		{
			Name: "nil",
			V:    nil,
		},
		{
			Name: "non-struct file",
			V:    true,
		},
		{
			Name: "non-struct section",
			V: struct {
				Section1 struct {
					Entry *string
				}
				Section2 bool
			}{},
		},
		{
			Name: "duplicated entry names",
			V: struct {
				Section1 struct {
					Entry       string
					Duplicated1 string `systemd:"Entry"`
					Duplicated2 string `systemd:"Entry"`
				}
			}{},
		},
		{
			Name: "unsupported entry type",
			V: struct {
				Section1 struct {
					Entry string
				}
				Section2 struct {
					Entry int
				}
			}{},
		},
		{
			Name: "unsupported entry slice type",
			V: struct {
				Section1 struct {
					Entry string
				}
				Section2 struct {
					Entry []int
				}
			}{
				Section2: struct {
					Entry []int
				}{
					Entry: []int{0},
				},
			},
		},
	}
	for _, td := range tests {
		t.Run("Name="+td.Name+";as value", func(t *testing.T) {
			err := Unmarshal(file, td.V)
			if err == nil {
				t.Error("Unmarshal() = nil; want non-nil")
			}
		})
		t.Run("Name="+td.Name+";as pointer", func(t *testing.T) {
			err := Unmarshal(file, &td.V)
			if err == nil {
				t.Error("Unmarshal() = nil; want non-nil")
			}
		})
	}
}

type marshalType string

func (m marshalType) MarshalSystemd() (Value, error) {
	if m == "fail" {
		return Value{string("fail: " + m)}, errors.New("error")
	}
	return Value{string("success: " + m)}, nil
}

func (m *marshalType) UnmarshalSystemd(value Value) error {
	if value.String() == "fail" {
		return errors.New("error")
	}
	*m = marshalType(value.String())
	return nil
}

type marshalTest struct {
	Section struct {
		Key marshalType
	}
}

func TestUnmarshalWithMarshaler(t *testing.T) {
	file := []byte(`[Section]
Key=ok
`)
	expected := marshalTest{
		Section: struct {
			Key marshalType
		}{
			Key: "ok",
		},
	}
	var s marshalTest
	err := Unmarshal(file, &s)
	if err != nil {
		t.Fatalf("Unmarshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(expected, s, nil); diff != "" {
		t.Errorf("Unmarshal() mismatch (-want +got):\n%s", diff)
	}
}

func TestUnmarshalShouldFailOnUnmarshalerError(t *testing.T) {
	file := []byte(`[Section]
Key=fail
`)
	var s marshalTest
	err := Unmarshal(file, &s)
	if err == nil {
		t.Errorf("Unmarshal() = nil; want non-nil")
	}
}
