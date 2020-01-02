// Package ast contains a parser and a serializer for systemd configuration files through AST trees
package ast

import "fmt"

// Position describes an arbitrary source position
// including line and column location
type Position struct {
	Offset int // Byte offset, starting at 0
	Line   int // Line number, starting at 1
	Column int // Column number, starting at 1 (character count per line)
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

// Node describes a basic node
type Node interface {
	Begin() Position // Position of first character belonging to the node
	End() Position   // Position of last character belonging to the node
}

// Children is a collection of child nodes
type Children []Node

func (c Children) Begin() Position { return c[0].Begin() }
func (c Children) End() Position   { return c[len(c)-1].End() }

// File represents a systemd configuration file
type File struct {
	Children Children // Comment, Include, and Section nodes

	begin, end Position
}

func (f *File) Begin() Position { return f.begin }
func (f *File) End() Position   { return f.end }

// Comment represents a single comment
type Comment struct {
	Text string // Trimmed comment text without # or ;

	begin, end Position
}

func (c *Comment) Begin() Position { return c.begin }
func (c *Comment) End() Position   { return c.end }

// Include represents a .include directive (obsolete)
type Include struct {
	Path string // Include path as specified in the directive

	begin, end Position
}

func (i *Include) Begin() Position { return i.begin }
func (i *Include) End() Position   { return i.end }

// Section represents a configuration section
type Section struct {
	Name     string   // Section name
	Children Children // Comment, Include, and Entry nodes

	begin, end Position
}

func (s *Section) Begin() Position { return s.begin }
func (s *Section) End() Position   { return s.end }

// Entry represents a single key-value configuration entry
type Entry struct {
	Key   string // Key name
	Value string // Entry value

	begin, end Position
}

func (e *Entry) Begin() Position { return e.begin }
func (e *Entry) End() Position   { return e.end }

// ParseError is returned for parsing errors
type ParseError struct {
	Err      error    // The actual error
	Position Position // Position where the error occurred
}

func (p *ParseError) Error() string {
	return fmt.Sprintf("%v at %s", p.Err, p.Position)
}
