package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/sergeymakinen/go-systemdconf/v3/tools/cmd/internal/common"
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
	if !common.Exist(cfgpath) {
		log.Fatalf("config %q not found", cfgpath)
	}
	b, err := os.ReadFile(cfgpath)
	if err != nil {
		log.Fatalf("failed to read %q: %v", cfgpath, err)
	}
	var conf config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	outdir := flag.Arg(1)
	if !common.Exist(outdir) {
		log.Fatalf("not found outdir %q", outdir)
	}
	man, err := goquery.NewDocumentFromReader(common.Read("https://www.freedesktop.org/software/systemd/man/index.html"))
	if err != nil {
		log.Fatal(err)
	}
	systemdName = strings.SplitN(man.Find("body > span").Text(), ".", 2)[0]
	if systemdName == "" || !strings.HasPrefix(systemdName, "systemd") {
		log.Fatal("no systemd version")
	}
	for _, file := range conf.Files {
		filename := filepath.Join(outdir, file.Path)
		if err := os.WriteFile(filename, []byte(file.String()), 0644); err != nil {
			log.Fatalf("failed to write %q: %v", filename, err)
		}
	}
}

var (
	reSpace      = regexp.MustCompile(`\s+`)
	reManSection = regexp.MustCompile(`\(\d\)`)
)

var systemdName string

func nodeToString(node *html.Node) string {
	switch node.Type {
	case html.TextNode:
		return reManSection.ReplaceAllString(reSpace.ReplaceAllString(node.Data, " "), "")
	case html.ElementNode:
		if node.Data == "table" {
			var data [][]string
			nodeToTable(&data, node)
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
		var buf strings.Builder
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			buf.WriteString(nodeToString(child))
		}
		switch node.Data {
		case "dd":
			return strings.TrimSpace(buf.String()) + "\n"
		case "div", "h3":
			if attr(node, "class") == "warning" {
				return indent(buf.String()) + "\n\n"
			}
			return buf.String() + "\n\n"
		case "span", "code", "strong", "a", "b", "em", "ul", "li", "th", "td":
			return buf.String()
		case "dt":
			return strings.TrimRight(buf.String(), "¶") + ": "
		case "br":
			return buf.String() + "\n"
		case "p":
			if attr(node, "class") != "title" {
				return buf.String() + "\n\n"
			}
		case "pre", "dl":
			return indent(buf.String()) + "\n\n"
		default:
			a := buf.String()
			_ = a
			log.Panicf("unknown tag %q", node.Data)
		}
	default:
		log.Panicf("unknown node type %q", node.Type)
	}
	return ""
}

func nodeToTable(data *[][]string, node *html.Node) {
	switch node.Type {
	case html.TextNode:
	case html.ElementNode:
		switch node.Data {
		case "table", "thead", "tbody", "col", "colgroup":
		case "tr":
			*data = append(*data, nil)
		case "th", "td":
			(*data)[len(*data)-1] = append((*data)[len(*data)-1], strings.TrimSpace(nodeToString(node)))
			return
		default:
			log.Panicf("unknown table tag %q", node.Data)
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			nodeToTable(data, child)
		}
	default:
		log.Panicf("unknown table node type %q", node.Type)
	}
}

func attr(node *html.Node, name string) string {
	for _, attr := range node.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
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
