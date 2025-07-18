// Package ast contains a parser and a printer for systemd configuration files through AST trees.
package ast

import "fmt"

// Position describes an arbitrary source position
// including line and column location.
type Position struct {
	Offset int // byte offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (character count per line)
}

func (p Position) String() string {
	s := "-"
	if p.Line != 0 {
		s = fmt.Sprintf("%d", p.Line)
		if p.Column != 0 {
			s += fmt.Sprintf(":%d", p.Column)
		}
	}
	return s
}

// Node describes a basic node.
type Node interface {
	Pos() Position // position of first character belonging to the node
	End() Position // position of last character belonging to the node
}

// Children is a collection of child nodes.
type Children []Node

func (c Children) Pos() Position { return c[0].Pos() }
func (c Children) End() Position { return c[len(c)-1].End() }

// File represents a systemd configuration file.
type File struct {
	Children Children // Comment, Include, and Section nodes

	pos, end Position
}

func (f *File) Pos() Position { return f.pos }
func (f *File) End() Position { return f.end }

// Comment represents a single comment.
type Comment struct {
	Text string // comment text starting with # or ;

	pos, end Position
}

func (c *Comment) Pos() Position { return c.pos }
func (c *Comment) End() Position { return c.end }

// Include represents a .include directive (obsolete).
type Include struct {
	Path string // include path as specified in the directive

	pos, end Position
}

func (i *Include) Pos() Position { return i.pos }
func (i *Include) End() Position { return i.end }

// Section represents a configuration section.
type Section struct {
	Name     string   // section name
	Children Children // Comment, Include, and Entry nodes

	pos, end Position
}

func (s *Section) Pos() Position { return s.pos }
func (s *Section) End() Position { return s.end }

// Entry represents a single key-value configuration entry.
type Entry struct {
	Key   string // key name
	Value string // entry value

	pos, end Position
}

func (e *Entry) Pos() Position { return e.pos }
func (e *Entry) End() Position { return e.end }

// ParseError is returned for parsing errors.
type ParseError struct {
	Err error    // the actual error
	Pos Position // position where the error occurred
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%v at %s", p.Err, p.Pos)
}
