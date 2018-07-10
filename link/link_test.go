package main

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestParseNode(t *testing.T) {
	cases := []struct {
		fileName string
		links    []Link
	}{
		{"ex1.html", []Link{
			{Href: "/other-page", Text: "A link to another page"},
		}},
		{"ex2.html", []Link{
			{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
			{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
		}},
	}

	for caseInd, tableCase := range cases {
		htmlFile, err := ioutil.ReadFile(tableCase.fileName)
		if err != nil {
			panic(err.Error())
		}

		doc, err := html.Parse(strings.NewReader(string(htmlFile)))
		if err != nil {
			log.Fatal(err)
		}

		links := parseNode(doc)
		if !cmp.Equal(links, tableCase.links) {
			t.Errorf("case number %d not passed", caseInd)
		}
	}

}
