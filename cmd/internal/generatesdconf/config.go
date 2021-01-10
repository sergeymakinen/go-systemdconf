package main

import (
	"fmt"
	"go/format"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sergeymakinen/go-systemdconf/v2/cmd/internal"
	"gopkg.in/yaml.v3"
)

type configField struct {
	Name, Type, SystemdName, Comment string
	Short                            bool
}

func (f configField) String() string {
	if f.Type == "" {
		f.Type = "systemdconf.Value"
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
		*f = generate(value.Value)
		return nil
	case yaml.SequenceNode:
		var list []string
		if err := value.Decode(&list); err != nil {
			return errors.Wrap(err, "failed to decode list")
		}
		for _, v := range list {
			for _, cf := range generate(v) {
				*f = append(*f, cf)
			}
		}
		return nil
	default:
		return mapYAML(value, func(key string, value *yaml.Node) error {
			var cf configField
			if err := value.Decode(&cf); err != nil {
				return errors.Wrap(err, "failed to decode configField")
			}
			cf.Name = key
			*f = append(*f, cf)
			return nil
		})
	}
}

type rawConfigStruct struct {
	Fields configStrings
}

type configStruct struct {
	Name, SystemdName, Comment string
	Embedded                   configStrings
	Fields                     configFields
}

func (s configStruct) String() string {
	buf := strings.Builder{}
	buf.WriteString(comment(s.Name+" represents "+s.Comment) + "\n")
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

type configStructs []configStruct

func (s *configStructs) UnmarshalYAML(value *yaml.Node) error {
	return mapYAML(value, func(key string, value *yaml.Node) error {
		var cs configStruct
		if err := value.Decode(&cs); err != nil {
			return errors.Wrap(err, "failed to decode configStruct")
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
		var rcs rawConfigStruct
		value.Decode(&rcs)
		if len(rcs.Fields) == 1 {
			cs.Comment += fmt.Sprintf("\n(see %s for details)", rcs.Fields[0])
		} else if len(rcs.Fields) > 1 {
			u := internal.ParseURL(rcs.Fields[0])
			u.Fragment = ""
			cs.Comment += fmt.Sprintf("\n(see %s for details)", u)
		}
		*s = append(*s, cs)
		return nil
	})
}

type configFile struct {
	Path, Name, Comment string
	Structs             configStructs
}

func (f configFile) String() string {
	buf := strings.Builder{}
	buf.WriteString("// DO NOT EDIT. This file is generated from " + systemdName + " by generatesdconf\n\n")
	buf.WriteString("package " + path.Dir(f.Path) + "\n\n")
	buf.WriteString("import \"github.com/sergeymakinen/go-systemdconf/v2\"\n\n")
	for _, s := range f.Structs {
		buf.WriteString(s.String() + "\n\n")
	}
	b, err := format.Source([]byte(buf.String()))
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to format source:\n%s\n", buf.String()))
	}
	return string(b)
}

type configFiles []configFile

func (f *configFiles) UnmarshalYAML(value *yaml.Node) error {
	return mapYAML(value, func(key string, value *yaml.Node) error {
		var cf configFile
		if err := value.Decode(&cf); err != nil {
			return errors.Wrap(err, "failed to decode configFile")
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
				Comment:     strings.ToUpper(short[0:1]) + short[1:],
				Short:       true,
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

type config struct {
	Files configFiles
}

type configStrings []string

func (s *configStrings) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		*s = configStrings{value.Value}
		return nil
	}
	var list []string
	if err := value.Decode(&list); err != nil {
		return errors.Wrap(err, "failed to decode list")
	}
	*s = list
	return nil
}

func mapYAML(node *yaml.Node, f func(key string, value *yaml.Node) error) error {
	if node.Kind != yaml.MappingNode {
		return errors.Errorf("node %q is not mapping", node.LongTag())
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
