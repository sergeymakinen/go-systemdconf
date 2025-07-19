package main

import (
	"fmt"
	"go/format"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sergeymakinen/go-systemdconf/v3/tools/cmd/internal/common"
	"gopkg.in/yaml.v3"
)

var (
	reNewLine       = regexp.MustCompile(`\n{2,}`)
	reLeadingSpace  = regexp.MustCompile(`\n +`)
	reTrailingSpace = regexp.MustCompile(` +\n`)
)

type config struct {
	Files configFiles
}

type configFiles []configFile

func (f *configFiles) UnmarshalYAML(value *yaml.Node) error {
	return mapYAML(value, func(key string, value *yaml.Node) error {
		var cf configFile
		if err := value.Decode(&cf); err != nil {
			return fmt.Errorf("failed to decode configFile: %w", err)
		}
		cf.Path = key
		if cf.Name == "" {
			cf.Name = filepath.Base(cf.Path)
			cf.Name = strings.Title(strings.TrimSuffix(cf.Name, filepath.Ext(cf.Name)))
		}
		var (
			structs []configStruct
			fs      *configStruct
		)
		for _, s := range cf.Structs {
			if !strings.HasPrefix(s.Name, ".") {
				structs = append(structs, s)
				continue
			}
			if fs == nil {
				name := cf.Name + "File"
				fs = &configStruct{
					Name:     name,
					Comment:  cf.Comment,
					Embedded: configStrings{"systemdconf.File"},
					Fields:   nil,
				}
			}
			s.Name = s.Name[1:]
			name := cf.Name + s.Name + "Section"
			structs = append(structs, configStruct{
				Name:     name,
				Comment:  s.Comment,
				Embedded: configStrings{"systemdconf.Section"},
				Fields:   s.Fields,
			})
			short := strings.SplitN(s.Comment, "\n", 2)[0]
			fs.Fields = append(fs.Fields, configField{
				Name:        s.Name,
				SystemdName: s.SystemdName,
				Type:        name,
				Comment:     short,
				Short:       true,
				Multiple:    s.Multiple,
			})
		}
		if fs != nil {
			structs = append([]configStruct{*fs}, structs...)
		}
		cf.Structs = structs
		*f = append(*f, cf)
		return nil
	})
}

type configFile struct {
	Path, Name, Comment string
	Structs             configStructs
}

func (f configFile) String() string {
	var buf strings.Builder
	buf.WriteString("// Code generated from " + systemdName + " by generatesdconf. DO NOT EDIT.\n\n")
	buf.WriteString("package " + path.Dir(f.Path) + "\n\n")
	buf.WriteString("import \"github.com/sergeymakinen/go-systemdconf/v3\"\n\n")
	for _, s := range f.Structs {
		buf.WriteString(s.String() + "\n\n")
	}
	b, err := format.Source([]byte(buf.String()))
	if err != nil {
		log.Fatalf("failed to format source: %v\n%s\n", err, buf.String())
	}
	return string(b)
}

type configStructs []configStruct

func (s *configStructs) UnmarshalYAML(value *yaml.Node) error {
	return mapYAML(value, func(key string, value *yaml.Node) error {
		var cs configStruct
		if err := value.Decode(&cs); err != nil {
			return fmt.Errorf("failed to decode configStruct: %w", err)
		}
		if cs.Name != "" {
			cs.SystemdName = cs.Name
		}
		cs.Name = key
		if cs.Comment == "" && strings.HasPrefix(cs.Name, ".") {
			name := cs.Name[1:]
			if cs.SystemdName != "" {
				name = cs.SystemdName
			}
			cs.Comment = fmt.Sprintf("[%s] section", name)
		}
		var fields configStructFields
		value.Decode(&fields)
		if len(fields.Fields) == 1 {
			cs.Comment += fmt.Sprintf("\n(see %s for details)", fields.Fields[0].URL)
		} else if len(fields.Fields) > 1 {
			u := common.ParseURL(fields.Fields[0].URL)
			u.Fragment = ""
			cs.Comment += fmt.Sprintf("\n(see %s for details)", u)
		}
		*s = append(*s, cs)
		return nil
	})
}

type configStruct struct {
	Name, SystemdName, Comment string
	Multiple                   bool
	Embedded                   configStrings
	Fields                     configFields
}

func (s configStruct) String() string {
	var buf strings.Builder
	buf.WriteString(comment(s.Name+" represents "+s.Comment) + ".\n")
	buf.WriteString("type " + s.Name + " struct {\n")
	if len(s.Embedded) > 0 {
		for _, v := range s.Embedded {
			buf.WriteString(v + "\n")
		}
		buf.WriteString("\n")
	}
	for _, f := range s.Fields {
		buf.WriteString(f.String() + "\n")
		if !f.Short {
			buf.WriteString("\n")
		}
	}
	buf.WriteString("}")
	return buf.String()
}

type configStrings []string

func (s *configStrings) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		*s = configStrings{value.Value}
		return nil
	}
	var list []string
	if err := value.Decode(&list); err != nil {
		return fmt.Errorf("failed to decode list: %w", err)
	}
	*s = list
	return nil
}

type configField struct {
	Name, Type, SystemdName, Comment string
	Short, Multiple                  bool
}

func (f configField) String() string {
	if f.Type == "" {
		f.Type = "systemdconf.Value"
	}
	if f.Multiple {
		f.Type = "[]" + f.Type
	}
	tag := ""
	if f.SystemdName != "" {
		tag = fmt.Sprintf("`systemd:\"%s\"`", f.SystemdName)
	}
	if f.Short {
		return fmt.Sprintf("%s %s %s // %s", f.Name, f.Type, tag, f.Comment)
	}
	return fmt.Sprintf("%s\n%s %s %s", comment(f.Comment), f.Name, f.Type, tag)
}

type configFields []configField

func (f *configFields) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*f = generate(value.Value, "", "")
		return nil
	case yaml.SequenceNode:
		for _, node := range value.Content {
			var cfs []configField
			switch node.Kind {
			case yaml.MappingNode:
				var addr configFieldAddr
				if err := node.Decode(&addr); err != nil {
					return fmt.Errorf("failed to decode configFields: %w", err)
				}
				cfs = generate(addr.URL, addr.Refsect, addr.Header)
			case yaml.ScalarNode:
				var s string
				if err := node.Decode(&s); err != nil {
					return fmt.Errorf("failed to decode configFields: %w", err)
				}
				cfs = generate(s, "", "")
			default:
				return fmt.Errorf("node %q is not mapping or scalar", node.Tag)
			}
			for _, cf := range cfs {
				*f = append(*f, cf)
			}
		}
		return nil
	default:
		return mapYAML(value, func(key string, value *yaml.Node) error {
			var cf configField
			if err := value.Decode(&cf); err != nil {
				return fmt.Errorf("failed to decode configField: %w", err)
			}
			cf.Name = key
			*f = append(*f, cf)
			return nil
		})
	}
}

type configStructFields struct {
	Fields configFieldAddrs
}

type configFieldAddrs []configFieldAddr

func (a *configFieldAddrs) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*a = append(*a, configFieldAddr{URL: value.Value})
		return nil
	case yaml.SequenceNode:
		for _, node := range value.Content {
			switch node.Kind {
			case yaml.MappingNode:
				var addr configFieldAddr
				if err := node.Decode(&addr); err != nil {
					return fmt.Errorf("failed to decode configFields: %w", err)
				}
				*a = append(*a, addr)
			case yaml.ScalarNode:
				var s string
				if err := node.Decode(&s); err != nil {
					return fmt.Errorf("failed to decode configFields: %w", err)
				}
				*a = append(*a, configFieldAddr{URL: s})
			default:
				return fmt.Errorf("node %q is not mapping or scalar", node.Tag)
			}
		}
	default:
		return fmt.Errorf("node %q is not scalar or sequence", value.Tag)
	}
	return nil
}

type configFieldAddr struct {
	URL     string
	Refsect string
	Header  string
}

func mapYAML(node *yaml.Node, f func(key string, value *yaml.Node) error) error {
	if node.Kind != yaml.MappingNode {
		return fmt.Errorf("node %q is not mapping", node.Tag)
	}
	var (
		valueNode bool
		key       string
	)
	for _, n := range node.Content {
		if !valueNode {
			key = n.Value
		} else {
			if err := f(key, n); err != nil {
				return err
			}
		}
		valueNode = !valueNode
	}
	return nil
}

func generate(url, refsect, header string) []configField {
	sections := common.FindDirectives(url, &common.FindDirectivesOptions{
		Refsect: refsect,
		Header:  header,
	}).Sections
	var fields []configField
	for _, section := range sections {
		for _, directive := range section.Directives {
			desc := nodeToString(directive.Description.Get(0))
			desc = reLeadingSpace.ReplaceAllString(desc, "\n")
			desc = reTrailingSpace.ReplaceAllString(desc, "\n")
			desc = strings.TrimSpace(reNewLine.ReplaceAllString(desc, "\n\n"))
			desc = wordwrap(desc, 100)
			field := configField{
				Name:    directive.Name,
				Comment: desc,
			}
			if strings.Contains(field.Name, "-") {
				field.SystemdName = field.Name
				field.Name = strings.ReplaceAll(field.Name, "-", "")
			}
			fields = append(fields, field)
		}
	}
	return fields
}
