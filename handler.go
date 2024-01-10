package main

// This goes with the handlerproblem file

import (
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v3"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := pathsToUrls[r.URL.Path]

		if url != "" {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type DATA struct {
	Path string
	Url string
}

func yamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var data []DATA
	if err := yaml.Unmarshal(yml, &data); err != nil {
		log.Fatal(err.Error())
	}

	urlMap := make(map[string]string)
	for _, item := range data {
		urlMap[item.Path] = item.Url
	}

	return mapHandler(urlMap, fallback), nil
}
