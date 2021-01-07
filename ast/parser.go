package ast

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var (
	ErrLineTooLong             = errors.New("line too long")
	ErrContinuationLineTooLong = errors.New("continuation line too long")
	ErrMissingSection          = errors.New("entry outside of section")
	ErrMissingEqual            = errors.New("missing =")
	ErrMissingKey              = errors.New("missing key name before =")
)

const (
	maxLineLength   = 1024 * 1024
	whitespaceChars = " \t\n\r"
	commentChars    = ";#"
)

// Parser is a systemd config file parser which conforms systemd.syntax
// and follows conventions defined by a default C parser
// (see https://github.com/systemd/systemd/blob/master/src/shared/conf-parser.c for details)
type Parser struct {
	b       *bufio.Reader
	buf     bytes.Buffer
	file    *File
	section *Section
	pos     Position
}

// NewParser returns a new parser that reads from r
func NewParser(r io.Reader) *Parser {
	return &Parser{
		b: bufio.NewReader(r),
	}
}

// Parse parses a systemd configuration file and returns its File node
func (p *Parser) Parse() (*File, error) {
	p.pos = Position{
		Line:   1,
		Column: 1,
	}
	p.file = &File{
		begin: p.pos,
	}
	p.section = nil
	var (
		eof                                bool
		continuation                       bytes.Buffer
		continuationBegin, continuationEnd Position
		continuationComments               []*Comment
	)
	for !eof {
		s, err, begin, end := p.readLine()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			eof = true
		}
		if len(s) == 0 {
			continue
		}
		if len(s) > maxLineLength {
			return nil, &ParseError{
				Err:      ErrLineTooLong,
				Position: end,
			}
		}
		s = p.trimLeft(s, &begin)
		if strings.IndexAny(s, commentChars) == 0 {
			s = p.trimRight(s, &end)
			n := &Comment{
				Text:  s,
				begin: begin,
				end:   end,
			}
			if p.section != nil {
				if continuation.Len() > 0 {
					continuationComments = append(continuationComments, n)
				} else {
					p.section.Children = append(p.section.Children, n)
				}
			} else {
				p.file.Children = append(p.file.Children, n)
			}
			continue
		}
		if continuation.Len()+len(s) > maxLineLength {
			return nil, &ParseError{
				Err:      ErrContinuationLineTooLong,
				Position: end,
			}
		}
		if strings.HasSuffix(s, "\\") {
			if continuation.Len() == 0 {
				continuationBegin = begin
			}
			continuationEnd = end
			continuation.WriteString(s[:len(s)-1])
			continuation.WriteString(" ")
			continue
		} else if continuation.Len() > 0 {
			continuation.WriteString(s)
			s = continuation.String()
			continuation.Reset()
			begin = continuationBegin
		}
		if err := p.parseLine(s, begin, end); err != nil {
			return nil, err
		}
		if len(continuationComments) > 0 {
			for _, c := range continuationComments {
				p.section.Children = append(p.section.Children, c)
			}
			continuationComments = nil
		}
	}
	if continuation.Len() > 0 {
		if err := p.parseLine(continuation.String(), continuationBegin, continuationEnd); err != nil {
			return nil, err
		}
	}
	p.file.end = p.pos
	return p.file, nil
}

func (p *Parser) parseLine(s string, begin, end Position) error {
	if len(s) == 0 {
		return nil
	}
	if strings.HasPrefix(s, ".include") {
		path := strings.TrimLeft(strings.TrimPrefix(s, ".include"), whitespaceChars)
		if path != "" {
			path = p.trimRight(path, &end)
			n := &Include{
				Path:  path,
				begin: begin,
				end:   end,
			}
			if p.section != nil {
				p.section.Children = append(p.section.Children, n)
			} else {
				p.file.Children = append(p.file.Children, n)
			}
			return nil
		}
	}
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		p.section = &Section{
			Name:  s[1 : len(s)-1],
			begin: begin,
			end:   end,
		}
		p.file.Children = append(p.file.Children, p.section)
		return nil
	}
	if p.section == nil {
		return &ParseError{
			Err:      ErrMissingSection,
			Position: end,
		}
	}
	s = p.trimRight(p.trimLeft(s, &begin), &end)
	parts := strings.SplitN(s, "=", 2)
	if len(parts) == 1 {
		return &ParseError{
			Err:      ErrMissingEqual,
			Position: end,
		}
	}
	if len(parts[0]) == 0 {
		return &ParseError{
			Err:      ErrMissingKey,
			Position: end,
		}
	}
	p.section.Children = append(p.section.Children, &Entry{
		Key:   strings.TrimRight(parts[0], whitespaceChars),
		Value: strings.TrimLeft(parts[1], whitespaceChars),
		begin: begin,
		end:   end,
	})
	return nil
}

// readLine reads a next line until it meets \n\r\0 in every combination (\0 must be last).
// It follows rules from the original C code:
//	Previous char was a NUL? This is not an EOL, but the previous char was? This type of
//	EOL marker has been seen right before?  In either of these three cases we are
//	done.
func (p *Parser) readLine() (s string, err error, begin, end Position) {
	var (
		done    bool
		r       rune
		size    int
		prevEnd *Position
	)
	metEOL := map[rune]bool{}
	unread := func() {
		p.b.UnreadRune()
		p.pos.Offset -= size
		p.pos.Column--
		if prevEnd != nil {
			end = *prevEnd
		}
		p.pos.Line++
		p.pos.Column = 1
		err = nil
	}
	begin, end = p.pos, p.pos
	p.buf.Reset()
	for !done {
		r, size, err = p.b.ReadRune()
		if err != nil {
			if err != io.EOF {
				err = &ParseError{
					Err:      err,
					Position: p.pos,
				}
				return
			}
			if len(metEOL) > 0 {
				unread()
			}
			break
		}
		prevEnd = nil
		switch r {
		case 0:
			done = true
		case '\n', '\r':
			if _, ok := metEOL[r]; ok {
				unread()
				done = true
			} else {
				metEOL[r] = true
			}
		default:
			if len(metEOL) > 0 {
				unread()
				done = true
			} else {
				prevEnd = &end
				end = p.pos
				p.pos.Column++
				p.buf.WriteRune(r)
			}
		}
		p.pos.Offset += size
	}
	s = p.buf.String()
	return
}

func (p *Parser) trimLeft(s string, begin *Position) string {
	i := 0
	for i < len(s) {
		r, w := utf8.DecodeRuneInString(s[i:])
		if !strings.ContainsRune(whitespaceChars, r) {
			break
		}
		begin.Offset += w
		begin.Column += 1
		i += w
	}
	return s[i:]
}

func (p *Parser) trimRight(s string, end *Position) string {
	i := len(s)
	for i > 0 {
		r, w := utf8.DecodeLastRuneInString(s[:i])
		if !strings.ContainsRune(whitespaceChars, r) {
			break
		}
		end.Offset -= w
		end.Column -= 1
		i -= w
	}
	return s[:i]
}
