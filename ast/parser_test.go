package ast

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseComplexFile(t *testing.T) {
	b, err := os.ReadFile("testdata/complex.conf")
	if err != nil {
		t.Fatal(err)
	}
	file, err := Parse(bytes.NewReader(b))
	if err != nil {
		t.Fatalf("Parse() = _, %v; want nil", err)
	}
	expectedFile := &File{
		Children: Children{
			&Section{
				Name: "Section A",
				Children: Children{
					&Entry{
						Key:   "KeyOne",
						Value: "value 1",
						pos:   Position{15, 3, 1},
						end:   Position{36, 3, 22},
					},
					&Entry{
						Key:   "KeyTwo",
						Value: "value 2",
						pos:   Position{38, 4, 1},
						end:   Position{51, 4, 14},
					},
					&Comment{
						Text: "# a comment",
						pos:  Position{54, 6, 1},
						end:  Position{64, 6, 11},
					},
					&Include{
						Path: "/usr/lib/systemd/system/httpd.service",
						pos:  Position{67, 8, 1},
						end:  Position{112, 8, 46},
					},
				},
				pos: Position{3, 2, 3},
				end: Position{13, 2, 13},
			},
			&Section{
				Name: "Section B",
				Children: Children{
					&Entry{
						Key:   "Setting",
						Value: `"something" "some thing" "…"`,
						pos:   Position{126, 10, 1},
						end:   Position{163, 10, 36},
					},
					&Entry{
						Key:   "KeyTwo",
						Value: "value 2  value 2 continued",
						pos:   Position{165, 11, 1},
						end:   Position{205, 12, 24},
					},
				},
				pos: Position{114, 9, 1},
				end: Position{124, 9, 11},
			},
			&Section{
				Name: "Section C",
				Children: Children{
					&Entry{
						Key:   "KeyThree",
						Value: "value 2 value 2 continued continued 3",
						pos:   Position{220, 15, 1},
						end:   Position{336, 19, 18},
					},
					&Comment{
						Text: `# this line is ignored\`,
						pos:  Position{242, 16, 5},
						end:  Position{264, 16, 27},
					},
					&Comment{
						Text: "; this line is ignored too",
						pos:  Position{266, 17, 1},
						end:  Position{291, 17, 26},
					},
					&Comment{
						Text: ";      test",
						pos:  Position{338, 20, 1},
						end:  Position{348, 20, 11},
					},
				},
				pos: Position{208, 14, 1},
				end: Position{218, 14, 11},
			},
		},
		pos: Position{0, 1, 1},
		end: Position{353, 24, 1},
	}
	if diff := cmp.Diff(expectedFile, file, cmp.AllowUnexported(File{}, Section{}, Entry{}, Include{}, Comment{})); diff != "" {
		t.Errorf("Parser.Parse() mismatch (-want +got):\n%s", diff)
	}
}

func TestParseEdgeCases(t *testing.T) {
	maxLengthStr := strings.Repeat("x", maxLineLength)
	tests := []struct {
		name, contents string
		file           *File
		err            error
	}{
		{
			name:     "no trailing newline",
			contents: "[Section]\nKey=Value",
			file: &File{
				Children: Children{
					&Section{
						Name: "Section",
						Children: Children{
							&Entry{
								Key:   "Key",
								Value: "Value",
								pos:   Position{10, 2, 1},
								end:   Position{18, 2, 9},
							},
						},
						pos: Position{0, 1, 1},
						end: Position{8, 1, 9},
					},
				},
				pos: Position{0, 1, 1},
				end: Position{19, 2, 10},
			},
		},
		{
			name:     "trailing backslash",
			contents: "[Section]\nKey=Value 1\\\n2\\\n3\\\n",
			file: &File{
				Children: Children{
					&Section{
						Name: "Section",
						Children: Children{
							&Entry{
								Key:   "Key",
								Value: "Value 1 2 3",
								pos:   Position{10, 2, 1},
								end:   Position{26, 4, 1},
							},
						},
						pos: Position{0, 1, 1},
						end: Position{8, 1, 9},
					},
				},
				pos: Position{0, 1, 1},
				end: Position{29, 5, 1},
			},
		},
		{
			name:     "too long section",
			contents: "[Section " + maxLengthStr + "]\nKey=Value\n",
			err:      ErrLineTooLong,
		},
		{
			name:     "too long value",
			contents: "[Section]\nKey=Value + " + maxLengthStr + "\n",
			err:      ErrLineTooLong,
		},
		{
			name:     "too long continuation",
			contents: "[Section]\nKey=Value + " + maxLengthStr[:maxLineLength-100] + "\\\n" + maxLengthStr[:100] + "\n",
			err:      ErrContinuationLineTooLong,
		},
		{
			name:     "missing section",
			contents: "Key=Value\n",
			err:      ErrMissingSection,
		},
		{
			name:     "missing =",
			contents: "[Section]\nKey\n",
			err:      ErrMissingEqual,
		},
		{
			name:     "missing key name",
			contents: "[Section]\n=Value\n",
			err:      ErrMissingKey,
		},
	}
	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			file, err := Parse(strings.NewReader(td.contents))
			var err2 *ParseError
			if errors.As(err, &err2) {
				err = err2.Err
			}
			if err != td.err {
				t.Fatalf("Parse() = _, %v; want %v", err, td.err)
			}
			if td.file != nil {
				if diff := cmp.Diff(td.file, file, cmp.AllowUnexported(File{}, Section{}, Entry{}, Include{}, Comment{})); diff != "" {
					t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
