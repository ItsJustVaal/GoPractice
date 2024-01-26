package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct{
	Urls []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}


func main() {
	getURL := flag.String("getURL", "https://www.footballnews.net", "URL To get pages from")
	maxDepth := flag.Int("maxDepth", 2, "BFS Search depth")
	flag.Parse()

	links := bfs(*getURL, *maxDepth)
	toXML := urlset{
		Xmlns: xmlns,
	}

	for _, link := range links {
		toXML.Urls = append(toXML.Urls, loc{link})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
	fmt.Println()
}

func bfs(urlString string, maxDepth int) []string {
	var queue map[string]struct{}
	seen := make(map[string]struct{})
	newQueue := map[string]struct{}{
		urlString: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		queue, newQueue = newQueue, make(map[string]struct{})
		if len(queue) == 0 {
			break
		}
		for url, _ := range queue {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					newQueue[link] = struct{}{}
				}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

func get(urlString string) []string {
	resp, err := http.Get(urlString)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()
	return filter(hrefs(resp.Body, base), base)
}

func hrefs(r io.Reader, base string) []string {
	links := Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(links []string, base string) []string {
	var ret []string
	for _, l := range links {
		if strings.HasPrefix(l, base) {
			ret = append(ret, l)
		}
	}
	return ret
}
