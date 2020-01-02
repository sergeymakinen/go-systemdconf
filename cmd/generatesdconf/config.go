package main

import (
	"fmt"
	"go/format"
	"log"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type configField struct {
	Name, Type, Comment string
	Short               bool
}

func (f configField) String() string {
	if f.Type == "" {
		f.Type = "systemdconf.Value"
	}
	if f.Short {
		return fmt.Sprintf("%s %s // %s", f.Name, f.Type, f.Comment)
	}
	return fmt.Sprintf("%s\n%s %s", comment(f.Comment), f.Name, f.Type)
}

type configFields []configField

func (f *configFields) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		*f = generate(value.Value)
	case yaml.SequenceNode:
		var list []string
		if err := value.Decode(&list); err != nil {
			log.Panic(err)
		}
		for _, v := range list {
			for _, cf := range generate(v) {
				*f = append(*f, cf)
			}
		}
	default:
		mapYAML(value, func(key string, value *yaml.Node) {
			var cf configField
			if err := value.Decode(&cf); err != nil {
				log.Panic(err)
			}
			cf.Name = key
			*f = append(*f, cf)
		})
	}
	return nil
}

type configStruct struct {
	Name, Comment string
	Embedded      configStrings
	Fields        configFields
}

func (s configStruct) String() string {
	buf := strings.Builder{}
	buf.WriteString(comment(s.Comment) + "\n")
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
	mapYAML(value, func(key string, value *yaml.Node) {
		var cs configStruct
		if err := value.Decode(&cs); err != nil {
			log.Panic(err)
		}
		cs.Name = key
		*s = append(*s, cs)
	})
	return nil
}

type configFile struct {
	Path, Name, Comment string
	Structs             configStructs
}

func (f configFile) String() string {
	buf := strings.Builder{}
	buf.WriteString("// DO NOT EDIT. This file is generated from " + systemdName + " by generatesdconf\n\n")
	buf.WriteString("package " + path.Dir(f.Path) + "\n\n")
	buf.WriteString("import \"github.com/sergeymakinen/go-systemdconf\"\n\n")
	for _, s := range f.Structs {
		buf.WriteString(s.String() + "\n\n")
	}
	b, err := format.Source([]byte(buf.String()))
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

type configFiles []configFile

func (f *configFiles) UnmarshalYAML(value *yaml.Node) error {
	mapYAML(value, func(key string, value *yaml.Node) {
		var cf configFile
		if err := value.Decode(&cf); err != nil {
			log.Panic(err)
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
					Comment:  name + " represents " + cf.Comment,
					Embedded: configStrings{"systemdconf.File"},
					Fields:   nil,
				}
			}
			s.Name = s.Name[1:]
			name := cf.Name + s.Name + "Section"
			structs = append(structs, configStruct{
				Name:     name,
				Comment:  name + " represents " + s.Comment,
				Embedded: configStrings{"systemdconf.Section"},
				Fields:   s.Fields,
			})
			fs.Fields = append(fs.Fields, configField{
				Name:    s.Name,
				Type:    name,
				Comment: strings.ToUpper(s.Comment[0:1]) + s.Comment[1:],
				Short:   true,
			})
		}
		if fs != nil {
			structs = append(structs, *fs)
		}
		cf.Structs = structs
		*f = append(*f, cf)
	})
	return nil
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
		log.Panic(err)
	}
	*s = list
	return nil
}

func mapYAML(node *yaml.Node, f func(key string, value *yaml.Node)) {
	if node.Kind != yaml.MappingNode {
		log.Panicf("node %q is not mapping", node.LongTag())
	}
	var (
		valueNode bool
		key       string
	)
	for _, n := range node.Content {
		if !valueNode {
			key = n.Value
		} else {
			f(key, n)
		}
		valueNode = !valueNode
	}
}
