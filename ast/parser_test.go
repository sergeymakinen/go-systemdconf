package ast

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseComplexFile(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/complex.conf")
	if err != nil {
		panic("failed to read testdata/complex.conf: " + err.Error())
	}
	file, err := NewParser(bytes.NewReader(b)).Parse()
	if err != nil {
		t.Fatalf("Parser.Parse() = _, %v; want nil", err)
	}
	expectedFile := &File{
		Children: Children{
			&Section{
				Name: "Section A",
				Children: Children{
					&Entry{
						Key:   "KeyOne",
						Value: "value 1",
						begin: Position{15, 3, 1},
						end:   Position{36, 3, 22},
					},
					&Entry{
						Key:   "KeyTwo",
						Value: "value 2",
						begin: Position{38, 4, 1},
						end:   Position{51, 4, 14},
					},
					&Comment{
						Text:  "# a comment",
						begin: Position{54, 6, 1},
						end:   Position{64, 6, 11},
					},
					&Include{
						Path:  "/usr/lib/systemd/system/httpd.service",
						begin: Position{67, 8, 1},
						end:   Position{112, 8, 46},
					},
				},
				begin: Position{3, 2, 3},
				end:   Position{13, 2, 13},
			},
			&Section{
				Name: "Section B",
				Children: Children{
					&Entry{
						Key:   "Setting",
						Value: `"something" "some thing" "…"`,
						begin: Position{126, 10, 1},
						end:   Position{163, 10, 36},
					},
					&Entry{
						Key:   "KeyTwo",
						Value: "value 2  value 2 continued",
						begin: Position{165, 11, 1},
						end:   Position{205, 12, 24},
					},
				},
				begin: Position{114, 9, 1},
				end:   Position{124, 9, 11},
			},
			&Section{
				Name: "Section C",
				Children: Children{
					&Entry{
						Key:   "KeyThree",
						Value: "value 2 value 2 continued continued 3",
						begin: Position{220, 15, 1},
						end:   Position{336, 19, 18},
					},
					&Comment{
						Text:  `# this line is ignored\`,
						begin: Position{242, 16, 5},
						end:   Position{264, 16, 27},
					},
					&Comment{
						Text:  "; this line is ignored too",
						begin: Position{266, 17, 1},
						end:   Position{291, 17, 26},
					},
					&Comment{
						Text:  ";      test",
						begin: Position{338, 20, 1},
						end:   Position{348, 20, 11},
					},
				},
				begin: Position{208, 14, 1},
				end:   Position{218, 14, 11},
			},
		},
		begin: Position{0, 1, 1},
		end:   Position{353, 24, 1},
	}
	if diff := cmp.Diff(expectedFile, file, cmp.AllowUnexported(File{}, Section{}, Entry{}, Include{}, Comment{})); diff != "" {
		t.Errorf("Parser.Parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestParseEdgeCases(t *testing.T) {
	maxLengthStr := strings.Repeat("x", maxLineLength)
	tests := []struct {
		Name, Contents string
		File           *File
		Err            error
	}{
		{
			Name:     "no trailing newline",
			Contents: "[Section]\nKey=Value",
			File: &File{
				Children: Children{
					&Section{
						Name: "Section",
						Children: Children{
							&Entry{
								Key:   "Key",
								Value: "Value",
								begin: Position{10, 2, 1},
								end:   Position{18, 2, 9},
							},
						},
						begin: Position{0, 1, 1},
						end:   Position{8, 1, 9},
					},
				},
				begin: Position{0, 1, 1},
				end:   Position{19, 2, 10},
			},
		},
		{
			Name:     "trailing backslash",
			Contents: "[Section]\nKey=Value 1\\\n2\\\n3\\\n",
			File: &File{
				Children: Children{
					&Section{
						Name: "Section",
						Children: Children{
							&Entry{
								Key:   "Key",
								Value: "Value 1 2 3",
								begin: Position{10, 2, 1},
								end:   Position{26, 4, 1},
							},
						},
						begin: Position{0, 1, 1},
						end:   Position{8, 1, 9},
					},
				},
				begin: Position{0, 1, 1},
				end:   Position{29, 5, 1},
			},
		},
		{
			Name:     "too long section",
			Contents: "[Section " + maxLengthStr + "]\nKey=Value\n",
			Err:      ErrLineTooLong,
		},
		{
			Name:     "too long value",
			Contents: "[Section]\nKey=Value + " + maxLengthStr + "\n",
			Err:      ErrLineTooLong,
		},
		{
			Name:     "too long continuation",
			Contents: "[Section]\nKey=Value + " + maxLengthStr[:maxLineLength-100] + "\\\n" + maxLengthStr[:100] + "\n",
			Err:      ErrContinuationLineTooLong,
		},
		{
			Name:     "missing section",
			Contents: "Key=Value\n",
			Err:      ErrMissingSection,
		},
		{
			Name:     "missing =",
			Contents: "[Section]\nKey\n",
			Err:      ErrMissingEqual,
		},
		{
			Name:     "missing key name",
			Contents: "[Section]\n=Value\n",
			Err:      ErrMissingKey,
		},
	}
	for _, td := range tests {
		t.Run(td.Name, func(t *testing.T) {
			file, err := NewParser(strings.NewReader(td.Contents)).Parse()
			if err2, ok := err.(*ParseError); ok {
				err = err2.Err
			}
			if err != td.Err {
				t.Fatalf("Parser.Parse() = _, %v; want %v", err, td.Err)
			}
			if td.File != nil {
				if diff := cmp.Diff(td.File, file, cmp.AllowUnexported(File{}, Section{}, Entry{}, Include{}, Comment{})); diff != "" {
					t.Errorf("Parser.Parse() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
