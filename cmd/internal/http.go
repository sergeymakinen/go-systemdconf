package internal

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	urlpkg "net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func Read(url string) io.Reader {
	filename := filepath.Join(os.TempDir(), fmt.Sprintf("systemdconf-%x", md5.Sum([]byte(url))))
	if Exist(filename) {
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

func ParseURL(url string) *urlpkg.URL {
	u, err := urlpkg.Parse(url)
	if err != nil {
		log.Fatal(errors.Wrapf(err, "failed to parse URL %q", url))
	}
	return u
}

type page struct {
	Description string
	Sections    []section
}

type section struct {
	Selection  *goquery.Selection
	Title, URL string
	Directives []directive
}

type directive struct {
	Name        string
	Description *goquery.Selection
}

var (
	spaceRE         = regexp.MustCompile(`\s+`)
	directiveNameRE = regexp.MustCompile(`([A-Z][\w-]*)=([^,]+)?`)
)

func FindDirectives(url string) page {
	u := ParseURL(url)
	id := u.Fragment
	u.Fragment = ""
	doc, err := goquery.NewDocumentFromReader(Read(u.String()))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create document"))
	}
	p := page{
		Description: doc.Find(".titlepage ~ .refnamediv p").Text(),
	}
	doc.Find(".refsect1:has(dl.variablelist)").Each(func(i int, s *goquery.Selection) {
		h2 := s.ChildrenFiltered("h2")
		if id != "" && s.ChildrenFiltered("h2").AttrOr("id", "") != id {
			return
		}
		title := strings.TrimRight(h2.Text(), "¶")
		title = spaceRE.ReplaceAllString(strings.ReplaceAll(title, "\n", ""), " ")
		section := section{
			Selection: s,
			Title:     title,
			URL:       u.ResolveReference(ParseURL(h2.ChildrenFiltered("a").AttrOr("href", ""))).String(),
		}
		s.Find("dl.variablelist").ChildrenFiltered("dt").Each(func(i int, s *goquery.Selection) {
			var dirs []string
			for _, match := range directiveNameRE.FindAllStringSubmatch(s.Text(), -1) {
				dirs = append(dirs, match[1])
			}
			if len(dirs) == 0 {
				return
			}
			for _, dir := range dirs {
				section.Directives = append(section.Directives, directive{
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
