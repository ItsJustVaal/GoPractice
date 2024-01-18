package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var page = `<!DOCTYPE html>
				<html lang="en">
				<head>
					<meta charset="UTF-8" />
					<meta name="viewport" content="width=device-width, initial-scale=1.0" />
					<title>Choose Your Own Adventure</title>
				</head>
				<body>
					<div class="page">
						<h1>{{.Title}}</h1>
						{{range .Paragraphs}}
							<p>{{.}}</p>
						{{end}}
						<ul>
						{{range .Option}}
							<li><a href="/{{.Arc}}">{{.Text}}</a></li>
						{{end}}
						</ul>
					</div>
				</body>
				<style>
				h1 {
					text-align: center;
					position:relative;
				}
				.page {
					width:80%;
					max-width:500px;
					margin:auto;
					margin-bottom:40px;
					padding: 80px;
					padding-top: 5px;
					padding-bottom: 5px;
					border: 1px solid black;
					background: #FFFCF6;
					box-shadow: 0 10px 6px -6px #777;
				}

				li {
					padding: 5px;
				}
				
				</style>
			</html>`

func init() {
	tpl = template.Must(template.New("").Parse(page))
}

var tpl *template.Template

func makeHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" || path == "/intro" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.Story[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter Not Found", http.StatusNotFound)
}

func parseJson(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var file Story
	if err := d.Decode(&file); err != nil {
		return file, err
	}
	return file, nil
}

func cyoa() {
	// load file
	data, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	story, err := parseJson(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	h := makeHandler(story)

	log.Fatal(http.ListenAndServe(":8000", h))
}
