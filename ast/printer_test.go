package ast

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFprintComplexFile(t *testing.T) {
	b, err := os.ReadFile("testdata/complex.conf")
	if err != nil {
		t.Fatal(err)
	}
	file, err := Parse(bytes.NewReader(b))
	if err != nil {
		t.Fatalf("Parse() = _, %v; want nil", err)
	}
	var buf bytes.Buffer
	err = Fprint(&buf, file)
	if err != nil {
		t.Fatalf("Fprint() = %v; want nil", err)
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
		t.Errorf("Fprint() mismatch (-want +got):\n%s", diff)
	}
}

func TestFprintWithOptions(t *testing.T) {
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
		name     string
		file     *File
		opts     opts
		contents string
		err      error
	}{
		{
			name:     "nil",
			contents: "",
		},
		{
			name: "Newline",
			file: simpleFile,
			opts: opts{
				Newline: "\r\n",
			},
			contents: "[Section]\r\nKey=Value 123456789\r\n",
		},
		{
			name: "KeyValueDelimiter",
			file: simpleFile,
			opts: opts{
				KeyValueDelimiter: " = ",
			},
			contents: `[Section]
Key = Value 123456789
`,
		},
		{
			name: "ValueWrapLength",
			file: simpleFile,
			opts: opts{
				ValueWrapLength: 4,
			},
			contents: `[Section]
Key=Value \
    123456789
`,
		},
	}
	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			p := &Printer{
				Newline:           td.opts.Newline,
				KeyValueDelimiter: td.opts.KeyValueDelimiter,
				ValueWrapLength:   td.opts.ValueWrapLength,
			}
			var buf bytes.Buffer
			err := p.Fprint(&buf, td.file)
			if err != td.err {
				t.Fatalf("Printer.Fprint() = %v; want %v", err, td.err)
			}
			if diff := cmp.Diff(td.contents, buf.String(), nil); diff != "" {
				t.Errorf("Printer.Fprint() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
