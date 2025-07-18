package ast

import (
	"cmp"
	"errors"
	"io"
	"reflect"
	"strconv"
	"strings"
)

const (
	newline           = "\n"
	keyValueDelimiter = "="
	valueWrapLength   = 200
)

// Fprint "pretty-prints" an AST node to output.
// It calls [Printer.Fprint] with default settings.
func Fprint(output io.Writer, node Node) error {
	return (&Printer{}).Fprint(output, node)
}

// A Printer controls the output of Fprint.
type Printer struct {
	Newline           string // default: "\n"
	KeyValueDelimiter string // default: "="
	ValueWrapLength   int    // default: 200
}

// Fprint "pretty-prints" an AST node to output for a given configuration.
func (p *Printer) Fprint(output io.Writer, node Node) error {
	if v := reflect.ValueOf(node); !v.IsValid() || (v.Kind() == reflect.Ptr && v.IsNil()) {
		return nil
	}
	switch v := node.(type) {
	case *File:
		var lastSection *Section
		for _, n := range v.Children {
			if section, ok := n.(*Section); ok {
				lastSection = section
			}
		}
		for _, n := range v.Children {
			if err := p.Fprint(output, n); err != nil {
				return err
			}
			if section, ok := n.(*Section); ok && section != lastSection {
				if _, err := output.Write([]byte(cmp.Or(p.Newline, newline))); err != nil {
					return err
				}
			}
		}
	case *Comment:
		if _, err := output.Write([]byte(strings.Trim(v.Text, whitespaceChars) + cmp.Or(p.Newline, newline))); err != nil {
			return err
		}
	case *Include:
		if _, err := output.Write([]byte(".include " + strings.Trim(v.Path, whitespaceChars) + cmp.Or(p.Newline, newline))); err != nil {
			return err
		}
	case *Section:
		if _, err := output.Write([]byte("[" + strings.Trim(v.Name, whitespaceChars) + "]" + cmp.Or(p.Newline, newline))); err != nil {
			return err
		}
		for _, n := range v.Children {
			if err := p.Fprint(output, n); err != nil {
				return err
			}
		}
	case *Entry:
		var buf strings.Builder
		buf.WriteString(strings.Trim(v.Key, whitespaceChars) + cmp.Or(p.KeyValueDelimiter, keyValueDelimiter))
		if cmp.Or(p.ValueWrapLength, valueWrapLength) > 0 {
			indentLen := buf.Len()
			current := 0
			for _, r := range strings.Trim(v.Value, whitespaceChars) {
				buf.WriteRune(r)
				current++
				if strings.ContainsRune(" \t", r) && current >= cmp.Or(p.ValueWrapLength, valueWrapLength) {
					buf.WriteString("\\" + cmp.Or(p.Newline, newline))
					buf.WriteString(strings.Repeat(" ", indentLen))
					current = 0
				}
			}
		} else {
			buf.WriteString(strings.Trim(v.Value, whitespaceChars))
		}
		buf.WriteString(cmp.Or(p.Newline, newline))
		if _, err := output.Write([]byte(buf.String())); err != nil {
			return err
		}
	default:
		return errors.New("unknown node type " + strconv.Quote(reflect.TypeOf(node).String()))
	}
	return nil
}
