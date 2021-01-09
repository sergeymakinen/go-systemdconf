package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/sergeymakinen/go-systemdconf/cmd/internal"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage: generatesdconf config.yml outdir\n")
	fmt.Fprintf(w, "generates systemd config file bindings\n")
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("generatesdconf: ")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}
	cfgpath := flag.Arg(0)
	if !internal.Exist(cfgpath) {
		log.Fatalf("config %q not found", cfgpath)
	}
	b, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to read %q", cfgpath))
	}
	var conf config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse config"))
	}
	outdir := flag.Arg(1)
	if !internal.Exist(outdir) {
		log.Fatalf("not found outdir %q", outdir)
	}
	man, err := goquery.NewDocumentFromReader(internal.Read("https://www.freedesktop.org/software/systemd/man/index.html"))
	if err != nil {
		log.Fatal(err)
	}
	systemdName = man.Find("body > span").Text()
	if systemdName == "" || !strings.HasPrefix(systemdName, "systemd") {
		log.Fatal("no systemd version")
	}
	for _, file := range conf.Files {
		filename := filepath.Join(outdir, file.Path)
		if err := ioutil.WriteFile(filename, []byte(file.String()), 0644); err != nil {
			log.Fatal(errors.Wrapf(err, "failed to write %q", filename))
		}
	}
}

var (
	spaceRE         = regexp.MustCompile(`\s+`)
	newlineRE       = regexp.MustCompile(`[\n]{2,}`)
	leadingSpaceRE  = regexp.MustCompile(`\n[ ]+`)
	trailingSpaceRE = regexp.MustCompile(`[ ]+\n`)
	manSectionRE    = regexp.MustCompile(`\(\d\)`)
	systemdName     string
)

func generate(url string) []configField {
	var fields []configField
	for _, section := range internal.FindDirectives(url).Sections {
		for _, directive := range section.Directives {
			desc := nodeToString(directive.Description.Get(0))
			desc = leadingSpaceRE.ReplaceAllString(desc, "\n")
			desc = trailingSpaceRE.ReplaceAllString(desc, "\n")
			desc = strings.TrimSpace(newlineRE.ReplaceAllString(desc, "\n\n"))
			desc = wordwrap(desc, 100)
			fields = append(fields, configField{
				Name:    directive.Name,
				Comment: desc,
			})
		}
	}
	return fields
}

func nodeToString(node *html.Node) string {
	switch node.Type {
	case html.TextNode:
		return manSectionRE.ReplaceAllString(spaceRE.ReplaceAllString(node.Data, " "), "")
	case html.ElementNode:
		if node.Data == "table" {
			var data [][]string
			nodeToTable(node, &data)
			buf := strings.Builder{}
			table := tablewriter.NewWriter(&buf)
			if len(data[0]) <= 2 {
				table.SetColWidth(80)
			}
			table.SetHeader(data[0])
			for _, row := range data[1:] {
				table.Append(row)
			}
			table.Render()
			return indent(buf.String())
		}
		buf := strings.Builder{}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			buf.WriteString(nodeToString(child))
		}
		switch node.Data {
		case "dd":
			return strings.TrimSpace(buf.String()) + "\n"
		case "div", "span", "code", "strong", "a", "b", "em", "ul", "li", "th", "td":
			return buf.String()
		case "dt":
			return strings.TrimRight(buf.String(), "¶") + ": "
		case "br":
			return buf.String() + "\n"
		case "p":
			if v, ok := lookupAttr(node, "class"); !ok || v != "title" {
				return buf.String() + "\n\n"
			}
		case "pre", "dl":
			return indent(buf.String()) + "\n\n"
		default:
			log.Panicf("unknown tag %q", node.Data)
		}
	default:
		log.Panicf("unknown node type %q", node.Type)
	}
	return ""
}

func nodeToTable(node *html.Node, data *[][]string) {
	switch node.Type {
	case html.TextNode:
	case html.ElementNode:
		switch node.Data {
		case "table", "thead", "tbody", "col", "colgroup":
		case "tr":
			*data = append(*data, nil)
		case "th", "td":
			(*data)[len(*data)-1] = append((*data)[len(*data)-1], nodeToString(node))
			return
		default:
			log.Panicf("unknown table tag %q", node.Data)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			nodeToTable(child, data)
		}
	default:
		log.Panicf("unknown table node type %q", node.Type)
	}
}

func lookupAttr(node *html.Node, name string) (string, bool) {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val, true
		}
	}
	return "", false
}

func indent(s string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		if strings.TrimSpace(lines[i]) != "" {
			lines[i] = "\t" + lines[i]
		} else {
			lines[i] = ""
		}
	}
	return strings.Join(lines, "\n")
}

func comment(s string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = "// " + lines[i]
	}
	return strings.Join(lines, "\n")
}

func wordwrap(s string, limit int) string {
	var (
		buf                     strings.Builder
		current                 int
		indentedLine, tableLine bool
	)
	for _, r := range s {
		if r == '\n' {
			buf.WriteRune(r)
			current = 0
			indentedLine = false
			tableLine = false
		} else if unicode.IsSpace(r) && r != '\t' {
			if !tableLine && current >= limit {
				buf.WriteString("\n")
				if indentedLine {
					buf.WriteString("\t")
				}
				current = 0
			} else {
				buf.WriteRune(r)
			}
		} else {
			if current == 0 && r == '\t' {
				indentedLine = true
			}
			if indentedLine && current == 1 && (r == '|' || r == '+') {
				tableLine = true
			}
			buf.WriteRune(r)
			current++
		}
	}
	return buf.String()
}
