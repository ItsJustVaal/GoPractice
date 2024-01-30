package main

// import (
// 	"flag"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// func main() {
// 	getURL := flag.String("getURL", "https://www.footballnews.net", "URL To get pages from")
// 	flag.Parse()

// 	links := get(*getURL)
// 	for _, l := range links {
// 		fmt.Println(l)
// 	}
// }

// func get(urlString string) []string {
// 	resp, err := http.Get(urlString)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	reqURL := resp.Request.URL
// 	baseURL := &url.URL{
// 		Scheme: reqURL.Scheme,
// 		Host: reqURL.Host,
// 	}

// 	base := baseURL.String()
// 	return filter(hrefs(resp.Body, base), base)
// }

// func hrefs(r io.Reader, base string) []string {
// 	links := Parse(r)
// 	var ret []string
// 	for _, l := range links {
// 		switch {
// 		case strings.HasPrefix(l.Href, "/"):
// 			ret = append(ret, base + l.Href)
// 		case strings.HasPrefix(l.Href, "http"):
// 			ret = append(ret, l.Href)
// 		}
// 	}
// 	return ret
// }

// func filter(links []string, base string) []string {
// 	var ret []string
// 	for _, l := range links {
// 		if !strings.HasPrefix(l, base){
// 			ret = append(ret, l)
// 		}
// 	}
// 	return ret
// }