package systemdconf

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalFile(t *testing.T) {
	unmarshal, err := os.ReadFile("testdata/unmarshal.conf")
	if err != nil {
		t.Fatal(err)
	}
	marshal, err := os.ReadFile("testdata/marshal.conf")
	if err != nil {
		t.Fatal(err)
	}
	var s fileTest
	err = Unmarshal(unmarshal, &s)
	if err != nil {
		t.Fatalf("Unmarshal() = %v; want nil", err)
	}
	b, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(marshal, b, nil); diff != "" {
		t.Errorf("Marshal() mismatch (-want +got):\n%s", diff)
	}
}

func TestMarshalShouldFail(t *testing.T) {
	tests := []struct {
		name, expected string
		v              any
	}{
		{
			name:     "nil",
			expected: "expected value, got nil",
			v:        nil,
		},
		{
			name:     "non-struct file",
			expected: "expected struct, got bool",
			v:        true,
		},
		{
			name:     "non-struct/slice/array section",
			expected: "expected struct, slice, or array, got int",
			v: struct {
				Section1 struct {
					Entry *string
				}
				Section2 int
			}{},
		},
		{
			name:     "non-struct array section",
			expected: "expected struct, got int",
			v: struct {
				Section []int
			}{Section: []int{1, 2}},
		},
		{
			name:     "duplicated entry names",
			expected: "conflicts with field struct",
			v: struct {
				Section1 struct {
					Entry       string
					Duplicated1 string `systemd:"Entry"`
					Duplicated2 string `systemd:"Entry"`
				}
			}{
				Section1: struct {
					Entry       string
					Duplicated1 string `systemd:"Entry"`
					Duplicated2 string `systemd:"Entry"`
				}{},
			},
		},
		{
			name:     "unsupported entry type",
			expected: "unsupported type",
			v: struct {
				Section1 struct {
					Entry string
				}
				Section2 struct {
					Entry int
				}
			}{},
		},
		{
			name:     "unsupported entry slice type",
			expected: "unsupported type",
			v: struct {
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
		t.Run(td.name, func(t *testing.T) {
			_, err := Marshal(td.v)
			if err == nil || !strings.Contains(err.Error(), td.expected) {
				t.Errorf("Marshal() = _, %v; does not contain %s", err, td.expected)
			}
		})
	}
}

func TestMarshalWithMarshaler(t *testing.T) {
	expected := []byte(`[Section]
Key=success: ok
`)
	s := marshalTest{
		Section: struct {
			Key marshalType
		}{
			Key: "ok",
		},
	}
	b, err := Marshal(s)
	if err != nil {
		t.Fatalf("Marshal() = %v; want nil", err)
	}
	if diff := cmp.Diff(expected, b, nil); diff != "" {
		t.Errorf("Marshal() mismatch (-want +got):\n%s", diff)
	}
}

func TestMarshalShouldFailOnMarshalerError(t *testing.T) {
	s := marshalTest{
		Section: struct {
			Key marshalType
		}{
			Key: "fail",
		},
	}
	_, err := Marshal(s)
	if err == nil {
		t.Errorf("Marshal() = nil; want non-nil")
	}
}
