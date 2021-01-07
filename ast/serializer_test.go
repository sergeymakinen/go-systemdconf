package ast

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSerializeComplexFile(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/complex.conf")
	if err != nil {
		panic("Error reading testdata/complex.conf: " + err.Error())
	}
	file, err := NewParser(bytes.NewReader(b)).Parse()
	if err != nil {
		t.Fatalf("Parser.Parse() = _, %v; want nil", err)
	}
	s := &Serializer{}
	var buf bytes.Buffer
	err = s.Serialize(file, &buf)
	if err != nil {
		t.Fatalf("Serializer.Serialize() = %v; want nil", err)
	}
	expected := `[Section A]
KeyOne=value 1
KeyTwo=value 2
# a comment
.include /usr/lib/systemd/system/httpd.service

[Section B]
Setting="something" "some thing" "…"
KeyTwo=value 2  value 2 continued

[Section C]
KeyThree=value 2 value 2 continued continued 3
# this line is ignored\
; this line is ignored too
;      test
`
	if diff := cmp.Diff(expected, buf.String(), nil); diff != "" {
		t.Errorf("Serializer.Serialize() mismatch (-want +got):\n%s", diff)
	}
}

func TestSerializeWithOptions(t *testing.T) {
	simpleFile := &File{
		Children: Children{
			&Section{
				Name: "Section",
				Children: Children{
					&Entry{
						Key:   "Key",
						Value: "Value 123456789",
					},
				},
			},
		},
	}
	type opts struct {
		Newline           string
		KeyValueDelimiter string
		ValueWrapLength   int
	}
	tests := []struct {
		Name     string
		File     *File
		Opts     opts
		Contents string
		Err      error
	}{
		{
			Name:     "nil",
			Contents: "",
		},
		{
			Name: "Newline",
			File: simpleFile,
			Opts: opts{
				Newline: "\r\n",
			},
			Contents: "[Section]\r\nKey=Value 123456789\r\n",
		},
		{
			Name: "KeyValueDelimiter",
			File: simpleFile,
			Opts: opts{
				KeyValueDelimiter: " = ",
			},
			Contents: `[Section]
Key = Value 123456789
`,
		},
		{
			Name: "ValueWrapLength",
			File: simpleFile,
			Opts: opts{
				ValueWrapLength: 4,
			},
			Contents: `[Section]
Key=Value \
    123456789
`,
		},
	}
	for _, td := range tests {
		t.Run("Name="+td.Name, func(t *testing.T) {
			s := &Serializer{}
			s.Newline = td.Opts.Newline
			s.KeyValueDelimiter = td.Opts.KeyValueDelimiter
			s.ValueWrapLength = td.Opts.ValueWrapLength
			var buf bytes.Buffer
			err := s.Serialize(td.File, &buf)
			if err != td.Err {
				t.Fatalf("Serializer.Serialize() = %v; want %v", err, td.Err)
			}
			if diff := cmp.Diff(td.Contents, buf.String(), nil); diff != "" {
				t.Errorf("Serializer.Serialize() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
