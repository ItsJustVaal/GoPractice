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
					<h1>{{.Title}}</h1>
					{{range .Paragraphs}}
					<p>{{.}}</p>
					{{end}}
					{{range .Option}}
					<ul>
						<li><a href="/{{.Arc}}"></a>{{.Text}}</li>
					</ul>
					{{end}}
				</body>
			</html>`

func init() {
	tpl = template.Must(template.New("").Parse(page))
}

var tpl *template.Template

func makeHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.Story["intro"])
	if err != nil {
		fmt.Println(err.Error())
	}
}

func parseJson(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)

	var file Story
	if err := d.Decode(&file); err != nil {
		return file, err
	}
	return file, nil
}

func main() {
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
