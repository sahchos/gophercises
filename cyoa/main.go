package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var storyData Story

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", storyHandler)

	storyFile, err := ioutil.ReadFile("cyoa/gopher.json")
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(storyFile, &storyData)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}

func executeTemplate(w http.ResponseWriter, data *Chapter) {
	tmpl := template.Must(template.ParseFiles("cyoa/chapter.html"))
	tmpl.Execute(w, data)
}

func storyHandler(w http.ResponseWriter, req *http.Request) {
	chapter := strings.Split(req.URL.Path, "/")[1]

	if data, ok := storyData[chapter]; ok {
		executeTemplate(w, &data)
	} else if chapter == "" {
		data := storyData["intro"]
		executeTemplate(w, &data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Chapter not found")
	}
}
