package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text []string
}

// This is if you want to read from a file instead of a get request
// func HtmlLinkParsing() []Link {
// 	file, err := os.Open("")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	return parse(file)
// }

func Parse(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err.Error())
	}

	links := getLinkNodes(doc)
	var linkItems []Link
	for _, item := range links {
		linkItems = append(linkItems, buildLinks(item))
	}
	return linkItems
}

func getLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, getLinkNodes(c)...)
	}
	return nodes
}

func buildLinks(n *html.Node) Link {
	var newLink Link
	for _, att := range n.Attr {
		if att.Key == "href" {
			newLink.Href = att.Val
			break
		}
	}
	newLink.Text = getText(n)
	return newLink
}

func getText(n *html.Node) []string {
	if n.Type == html.TextNode {
		finalText := "-" + strings.Join(strings.Fields(n.Data), " ")
		if finalText == "-" {
			return []string{}
		}
		return []string{finalText}
	}

	var htmlText []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlText = append(htmlText, getText(c)...)
	}
	return htmlText
}
