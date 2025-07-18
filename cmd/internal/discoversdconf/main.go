package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sergeymakinen/go-systemdconf/v3/cmd/internal/common"
)

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage: discoversdconf [-bare] [-refsect refsect] [-header header] [url|config.yml]\n")
	fmt.Fprintf(w, "fetches systemd directive pages and aggregates their content\n")
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("discoversdconf: ")
	flag.Usage = usage
	bare := flag.Bool("bare", false, "")
	refsect := flag.String("refsect", "", "")
	header := flag.String("header", "", "")
	flag.Parse()
	switch flag.NArg() {
	case 0:
		overview("")
	case 1:
		pathOrURL := flag.Arg(0)
		if strings.HasPrefix(pathOrURL, "http://") || strings.HasPrefix(pathOrURL, "https://") {
			config(pathOrURL, *refsect, *header, *bare)
		} else {
			overview(pathOrURL)
		}
	default:
		usage()
	}
}

type directiveGroup struct {
	Title, URL string
	Directives map[string][]string // map[url][]directive
}

func overview(cfgpath string) {
	var cfg []byte
	if cfgpath != "" {
		if !common.Exist(cfgpath) {
			log.Fatalf("config %q not found", cfgpath)
		}
		if b, err := os.ReadFile(cfgpath); err == nil {
			cfg = b
		} else {
			log.Fatalf("failed to read %q: %v", cfgpath, err)
		}
	}
	baseURL := common.ParseURL("https://www.freedesktop.org/software/systemd/man/systemd.directives.html")
	var dgs []directiveGroup
	for _, section := range common.FindDirectives(baseURL.String(), nil).Sections {
		dg := directiveGroup{
			Title:      section.Title,
			URL:        section.URL,
			Directives: map[string][]string{},
		}
		for _, directive := range section.Directives {
			directive.Description.Find("a").Each(func(i int, s *goquery.Selection) {
				dirURL, err := url.Parse(s.AttrOr("href", ""))
				if err != nil {
					log.Printf("ignored directive %s (%s): %s\n", directive.Name, section.Title, err)
					return
				}
				dirURL.Fragment = ""
				dirURL = baseURL.ResolveReference(dirURL)
				dg.Directives[dirURL.String()] = append(dg.Directives[dirURL.String()], directive.Name)
			})
		}
		dgs = append(dgs, dg)
	}
	for _, dg := range dgs {
		fmt.Printf("\n%s\n", dg.Title)
		for u, directives := range dg.Directives {
			mark := ""
			if cfg != nil {
				if bytes.Contains(cfg, []byte(u)) {
					mark = "\033[32m✔\033[0m "
				} else {
					mark = "\033[31m✗\033[0m "
				}
			}
			if len(directives) > 5 {
				directives = directives[0:6]
				directives[5] = "…"
			}
			fmt.Printf("\t%s%s (%s)\n", mark, u, strings.Join(directives, ", "))
		}
	}
}

var reSectionName = regexp.MustCompile(`\[([A-Z][\w-]*)]`)

func config(url, refsect, header string, bare bool) {
	page := common.FindDirectives(url, &common.FindDirectivesOptions{
		Refsect: refsect,
		Header:  header,
	})
	comment := fmt.Sprintf("TODO: %s", url)
	if page.Description != "" {
		comment = page.Description
	}
	code := "file.go:\n"
	if bare {
		code += fmt.Sprintf("  structs:\n    (TODO):\n      comment: %s\n      fields:", comment)
		if len(page.Sections) == 0 {
			code += "\n"
		} else if len(page.Sections) == 1 {
			code += " " + page.Sections[0].URL + "\n"
		} else {
			code += "\n"
			for _, section := range page.Sections {
				code += "        - " + section.URL + "\n"
			}
		}
	} else {
		code += fmt.Sprintf("  comment: %s\n  structs:\n", comment)
		for _, section := range page.Sections {
			sectionName := "(TODO)"
			if reSectionName.MatchString(section.Title) {
				sectionName = reSectionName.FindStringSubmatch(section.Title)[1]
			}
			code += fmt.Sprintf("    .%s:\n      comment: TODO\n      fields: %s\n", sectionName, section.URL)
		}
	}
	fmt.Print(code)
}
