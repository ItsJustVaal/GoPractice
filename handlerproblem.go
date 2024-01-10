package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func handlerProblem() {
	file := flag.String("yaml", "paths.yaml", "redirects")
	flag.Parse()

	yaml, err := os.ReadFile(*file)
	if err != nil {
		panic(err.Error())
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := mapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
// 	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
	yamlHandler, err := yamlHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
