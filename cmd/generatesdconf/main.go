package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	gourl "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"golang.org/x/net/html"
	"gopkg.in/yaml.v3"
)

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage: generatesdconf config.yaml outdir\n")
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
	if !exist(cfgpath) {
		log.Fatalf("not found config %q", cfgpath)
	}
	b, err := ioutil.ReadFile(cfgpath)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to read %q", cfgpath))
	}
	var conf config
	if err := yaml.Unmarshal(b, &conf); err != nil {
		log.Fatal(err)
	}
	outdir := flag.Arg(1)
	if !exist(outdir) {
		log.Fatalf("not found outdir %q", outdir)
	}
	man, err := goquery.NewDocumentFromReader(read("https://www.freedesktop.org/software/systemd/man/index.html"))
	if err != nil {
		log.Fatal(err)
	}
	systemdName = man.Find("body > span").Text()
	if systemdName == "" {
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
	keyNameRE       = regexp.MustCompile(`(\w+)=([^,]+)?`)
	manSectionRE    = regexp.MustCompile(`\(\d\)`)
	systemdName     string
)

func read(url string) io.Reader {
	filename := filepath.Join(os.TempDir(), fmt.Sprintf("generatesdconf-%x", md5.Sum([]byte(url))))
	if exist(filename) {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(errors.Wrapf(err, "failed to read %q", filename))
		}
		return bytes.NewBuffer(b)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to open %q", url))
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error for %q: %d %s", url, resp.StatusCode, resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to read %q", url))
	}
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Fatal(errors.Wrapf(err, "failed to write %q", filename))
	}
	return bytes.NewBuffer(b)
}

func generate(url string) []configField {
	u, err := gourl.Parse(url)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to parse %q", url))
	}
	header := u.Fragment
	u.Fragment = ""
	doc, err := goquery.NewDocumentFromReader(read(u.String()))
	if err != nil {
		log.Fatal(err)
	}
	var fields []configField
	doc.Find(".refsect1:has(dl.variablelist)").Each(func(i int, s *goquery.Selection) {
		if v, ok := s.ChildrenFiltered("h2").Attr("id"); !ok || v != header {
			return
		}
		s.Find("dl.variablelist").ChildrenFiltered("dt").Each(func(i int, s *goquery.Selection) {
			var keys []string
			for _, match := range keyNameRE.FindAllStringSubmatch(s.Text(), -1) {
				keys = append(keys, match[1])
			}
			if len(keys) == 0 {
				return
			}
			dd := nodeToString(s.NextFiltered("dd").Get(0))
			dd = leadingSpaceRE.ReplaceAllString(dd, "\n")
			dd = trailingSpaceRE.ReplaceAllString(dd, "\n")
			dd = strings.TrimSpace(newlineRE.ReplaceAllString(dd, "\n\n"))
			for _, v := range keys {
				fields = append(fields, configField{
					Name:    v,
					Comment: wordwrap(dd, 100),
				})
			}
			if len(fields) == 0 {
				log.Panicf("empty header %q", header)
			}
		})
	})
	if len(fields) == 0 {
		log.Panicf("not found header %q", header)
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
	for i, _ := range lines {
		lines[i] = "// " + lines[i]
	}
	return strings.Join(lines, "\n")
}

func exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.Panic(err)
		return false
	}
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
