package common

import (
	"bytes"
	"cmp"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	urlpkg "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Read(url string) io.Reader {
	filename := filepath.Join(os.TempDir(), fmt.Sprintf("systemdconf-%x", md5.Sum([]byte(url))))
	if Exist(filename) {
		b, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalf("failed to read %q: %v", filename, err)
		}
		return bytes.NewBuffer(b)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to open %q: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error for %q: %d %s", url, resp.StatusCode, resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read %q: %v", url, err)
	}
	if err := os.WriteFile(filename, b, 0644); err != nil {
		log.Fatalf("failed to write %q: %v", filename, err)
	}
	return bytes.NewBuffer(b)
}

func ParseURL(url string) *urlpkg.URL {
	u, err := urlpkg.Parse(url)
	if err != nil {
		log.Fatalf("failed to parse URL %q: %v", url, err)
	}
	return u
}

type Page struct {
	Description string
	Sections    []Section
}

type Section struct {
	Selection  *goquery.Selection
	Title, URL string
	Directives []Directive
}

type Directive struct {
	Name        string
	Description *goquery.Selection
}

var (
	reSpace         = regexp.MustCompile(`\s+`)
	reDirectiveName = regexp.MustCompile(`([A-Z][\w-]*)=([^,]+)?`)
)

type FindDirectivesOptions struct {
	Refsect string
	Header  string
}

func FindDirectives(url string, o *FindDirectivesOptions) Page {
	var opts FindDirectivesOptions
	if o != nil {
		opts = *o
	}
	u := ParseURL(url)
	id := u.Fragment
	u.Fragment = ""
	doc, err := goquery.NewDocumentFromReader(Read(u.String()))
	if err != nil {
		log.Fatalf("failed to create document: %v", err)
	}
	p := Page{
		Description: doc.Find(".titlepage ~ .refnamediv p").Text(),
	}
	doc.Find("." + cmp.Or(opts.Refsect, "refsect1") + ":has(dl.variablelist)").Each(func(i int, s *goquery.Selection) {
		header := s.ChildrenFiltered(cmp.Or(opts.Header, "h2"))
		if id != "" && header.AttrOr("id", "") != id {
			return
		}
		title := strings.TrimRight(header.Text(), "¶")
		title = reSpace.ReplaceAllString(strings.ReplaceAll(title, "\n", ""), " ")
		section := Section{
			Selection: s,
			Title:     title,
			URL:       u.ResolveReference(ParseURL(header.ChildrenFiltered("a").AttrOr("href", ""))).String(),
		}
		s.Find("dl.variablelist").ChildrenFiltered("dt").Each(func(i int, s *goquery.Selection) {
			var dirs []string
			for _, match := range reDirectiveName.FindAllStringSubmatch(s.Text(), -1) {
				dirs = append(dirs, match[1])
			}
			if len(dirs) == 0 {
				return
			}
			for _, dir := range dirs {
				section.Directives = append(section.Directives, Directive{
					Name:        dir,
					Description: s.NextFiltered("dd"),
				})
			}
		})
		if id != "" && len(section.Directives) == 0 {
			log.Fatalf("empty section %q", section.Title)
		}
		if len(section.Directives) > 0 {
			p.Sections = append(p.Sections, section)
		}
	})
	return p
}
