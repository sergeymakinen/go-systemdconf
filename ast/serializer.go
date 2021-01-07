package ast

import (
	"io"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

const (
	newline           = "\n"
	keyValueDelimiter = "="
	valueWrapLength   = 200
)

// Serializer is a systemd config AST node converter
type Serializer struct {
	Newline           string // Newline character(s) (default is "\n")
	KeyValueDelimiter string // Delimiter is written between entry key and value (default is "=")
	ValueWrapLength   int    // Limit of characters between whitespace characters without a line break (default limit is 200)
}

// Serialize converts an AST node to its textual form and writes it to w
func (s *Serializer) Serialize(node Node, w io.Writer) error {
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
			if err := s.Serialize(n, w); err != nil {
				return err
			}
			if section, ok := n.(*Section); ok && section != lastSection {
				if _, err := w.Write([]byte(stringOrDefault(s.Newline, newline))); err != nil {
					return err
				}
			}
		}
	case *Comment:
		if _, err := w.Write([]byte(strings.Trim(v.Text, whitespaceChars) + stringOrDefault(s.Newline, newline))); err != nil {
			return err
		}
	case *Include:
		if _, err := w.Write([]byte(".include " + strings.Trim(v.Path, whitespaceChars) + stringOrDefault(s.Newline, newline))); err != nil {
			return err
		}
	case *Section:
		if _, err := w.Write([]byte("[" + strings.Trim(v.Name, whitespaceChars) + "]" + stringOrDefault(s.Newline, newline))); err != nil {
			return err
		}
		for _, n := range v.Children {
			if err := s.Serialize(n, w); err != nil {
				return err
			}
		}
	case *Entry:
		var buf strings.Builder
		buf.WriteString(strings.Trim(v.Key, whitespaceChars) + stringOrDefault(s.KeyValueDelimiter, keyValueDelimiter))
		wrap := s.ValueWrapLength
		if wrap == 0 {
			wrap = valueWrapLength
		}
		if wrap > 0 {
			indentLen := buf.Len()
			current := 0
			for _, r := range strings.Trim(v.Value, whitespaceChars) {
				buf.WriteRune(r)
				current++
				if strings.ContainsRune(" \t", r) && current >= wrap {
					buf.WriteString("\\" + stringOrDefault(s.Newline, newline))
					buf.WriteString(strings.Repeat(" ", indentLen))
					current = 0
				}
			}
		} else {
			buf.WriteString(strings.Trim(v.Value, whitespaceChars))
		}
		buf.WriteString(stringOrDefault(s.Newline, newline))
		if _, err := w.Write([]byte(buf.String())); err != nil {
			return err
		}
	default:
		return errors.Errorf("unknown node type %s", reflect.TypeOf(node))
	}
	return nil
}

func stringOrDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
