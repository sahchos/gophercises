package main

import (
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

type Link struct {
	Href, Text string
}

func getNodeText(n * html.Node) string {
	text := ""
	for inner := n.FirstChild; inner != nil; inner = inner.NextSibling {
		switch inner.Type {
		case html.TextNode:
			{
				text += inner.Data
			}
		case html.ElementNode:
			{
				text += getNodeText(inner)
			}
		}
	}

	return strings.TrimSuffix(strings.TrimSpace(text), "\n")
}

func parseNode(n *html.Node) []Link {
	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := getNodeText(n)
				link := &Link{Href: a.Val, Text: text}
				links = append(links, *link)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseNode(c)...)
	}

	return links
}

func main() {
	htmlFile, err := ioutil.ReadFile("link/ex2.html")
	if err != nil {
		panic(err.Error())
	}

	doc, err := html.Parse(strings.NewReader(string(htmlFile)))
	if err != nil {
		log.Fatal(err)
	}

	links := parseNode(doc)

	for _, link := range links {
		fmt.Println("Link: ", link.Href)
		fmt.Println("Text: ", link.Text)
	}

}
