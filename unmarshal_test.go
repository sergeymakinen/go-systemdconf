package systemdconf

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type fileTest struct {
	File

	Section1 *struct {
		*Section

		Key1, Key2 Value
		Key3       **Value
	}
	Section3 struct {
		Key Value
	}
	Section4 []struct {
		Key1, Key2 Value
	}
	Section5 [3]struct {
		Key Value
	}
	Section6 [2]struct {
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
		Keys1       []bool
		Keys2       [10]**bool
		Key3        **bool
	}
	Duration struct {
		Key1  time.Duration
		Keys2 []time.Duration
		Keys3 [10]**time.Duration
		Key4  **time.Duration
	}
	String struct {
		Key1  string
		Keys2 []string
		Keys3 [10]**string
		Key4  **string
	}
}

func TestUnmarshalFile(t *testing.T) {
	b, err := os.ReadFile("testdata/unmarshal.conf")
	if err != nil {
		t.Fatal(err)
	}
	ptrValue := &Value{"value 1", "value 2"}
	boolTrue := true
	ptrBoolTrue := &boolTrue
	boolFalse := false
	ptrBoolFalse := &boolFalse
	duration3Min := 3 * time.Minute
	ptrDuration3Min := &duration3Min
	duration3Sec := 3 * time.Second
	ptrDuration3Sec := &duration3Sec
	duration3Hours := 3 * time.Hour
	ptrDuration3Hours := &duration3Hours
	stringValue1 := "value 1"
	ptrStringValue1 := &stringValue1
	stringValue2 := "value 2"
	ptrStringValue2 := &stringValue2
	stringValue := "value"
	ptrStringValue := &stringValue
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
			Key3       **Value
		}{
			Section: &Section{
				name: "Section1",
				entries: []*Entry{
					{
						Key:   "X-Key1",
						Value: Value{"value 3", "value 6"},
					},
					{
						Key:   "X-Key2",
						Value: Value{"value 4"},
					},
				},
			},
			Key1: Value{"value 1", "value 5"},
			Key2: Value{"value 2"},
			Key3: &ptrValue,
		},
		Section3: struct {
			Key Value
		}{Key: Value{"value 1"}},
		Section4: []struct {
			Key1, Key2 Value
		}{
			{
				Key1: Value{"value 1"},
				Key2: Value{"value 2"},
			},
			{
				Key1: Value{"value 3"},
				Key2: Value{"value 4"},
			},
		},
		Section5: [3]struct {
			Key Value
		}{
			{Key: Value{"value 1"}},
			{Key: Value{"value 2"}},
			{},
		},
		Section6: [2]struct {
			Key Value
		}{
			{Key: Value{"value 1"}},
			{},
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
			Keys1       []bool
			Keys2       [10]**bool
			Key3        **bool
		}{
			True:  true,
			False: false,
			Keys1: []bool{
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
			Keys2: [10]**bool{&ptrBoolTrue, &ptrBoolFalse},
			Key3:  &ptrBoolTrue,
		},
		Duration: struct {
			Key1  time.Duration
			Keys2 []time.Duration
			Keys3 [10]**time.Duration
			Key4  **time.Duration
		}{
			Key1: 3 * time.Second,
			Keys2: []time.Duration{
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
			Keys3: [10]**time.Duration{&ptrDuration3Min, &ptrDuration3Sec},
			Key4:  &ptrDuration3Hours,
		},
		String: struct {
			Key1  string
			Keys2 []string
			Keys3 [10]**string
			Key4  **string
		}{
			Key1:  "value",
			Keys2: []string{"value 1", "value 2", "value 3"},
			Keys3: [10]**string{&ptrStringValue1, &ptrStringValue2},
			Key4:  &ptrStringValue,
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
Key=value 1
Key=value 2
`)
	tests := []struct {
		name     string
		expected string
		f        func() error
	}{
		{
			name:     "invalid",
			expected: "",
			f: func() error {
				return Unmarshal([]byte("foo"), &struct{}{})
			},
		},
		{
			name:     "not pointer",
			expected: "expected non-nil pointer, got struct",
			f: func() error {
				return Unmarshal(file, struct{}{})
			},
		},
		{
			name:     "nil",
			expected: "expected non-nil pointer, got nil",
			f: func() error {
				return Unmarshal(file, nil)
			},
		},
		{
			name:     "nil pointer",
			expected: "expected non-nil pointer, got nil",
			f: func() error {
				var s *struct{}
				return Unmarshal(file, s)
			},
		},
		{
			name:     "non-struct file",
			expected: "expected non-nil pointer, got bool",
			f: func() error {
				return Unmarshal(file, true)
			},
		},
		{
			name:     "non-struct/slice/array section",
			expected: "expected struct, slice, or array, got int",
			f: func() error {
				return Unmarshal(file, &struct {
					Section int
				}{})
			},
		},
		{
			name:     "non-struct array section",
			expected: "expected struct, got int",
			f: func() error {
				return Unmarshal(file, &struct {
					Section []int
				}{})
			},
		},
		{
			name:     "duplicated entry names",
			expected: "conflicts with field struct",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key1 string `systemd:"Key"`
						Key2 string `systemd:"Key"`
					}
				}{})
			},
		},
		{
			name:     "unsupported entry type",
			expected: "unsupported type",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key int
					}
				}{})
			},
		},
		{
			name:     "unsupported entry slice type",
			expected: "unsupported type",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key []int
					}
				}{})
			},
		},
		{
			name:     "invalid duration",
			expected: "invalid systemd duration",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key time.Duration
					}
				}{})
			},
		},
		{
			name:     "invalid duration slice",
			expected: "invalid systemd duration",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key []time.Duration
					}
				}{})
			},
		},
		{
			name:     "invalid bool",
			expected: "invalid systemd boolean",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key bool
					}
				}{})
			},
		},
		{
			name:     "invalid bool slice",
			expected: "invalid systemd boolean",
			f: func() error {
				return Unmarshal(file, &struct {
					Section struct {
						Key []bool
					}
				}{})
			},
		},
	}
	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			if err := td.f(); err == nil || !strings.Contains(err.Error(), td.expected) {
				t.Errorf("Unmarshal() = %v; does not contain %s", err, td.expected)
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

type unmarshalByValueType string

func (m unmarshalByValueType) UnmarshalSystemd(value Value) error {
	if value.String() == "fail" {
		return errors.New("error")
	}
	return nil
}

func TestUnmarshalShouldFailOnUnmarshalerError(t *testing.T) {
	file := []byte(`[Section]
Key=fail
`)

	var s1 marshalTest
	err := Unmarshal(file, &s1)
	if err == nil {
		t.Errorf("Unmarshal() = nil; want non-nil")
	}

	var s2 struct {
		Section struct {
			Key unmarshalByValueType
		}
	}
	err = Unmarshal(file, &s2)
	if err == nil {
		t.Errorf("Unmarshal() = nil; want non-nil")
	}
}
