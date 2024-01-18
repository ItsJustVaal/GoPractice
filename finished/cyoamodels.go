package main

type handler struct {
	Story Story
}

type Story map[string]Section

type Section struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Option     []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
