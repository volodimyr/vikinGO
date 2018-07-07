package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Story struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var (
	path string
	tpl  *template.Template
)

func init() {
	flag.StringVar(&path, "Path",
		"vikinGo/choose-your-adventure/resources/gopher.json",
		"Set up path to json file")
	flag.Parse()

	tpl = template.Must(template.ParseGlob("vikinGo/choose-your-adventure/resources/*.gohtml"))
}

func main() {
	mux := defaultMux()
	stories, err := Parse(path)
	if err != nil {
		panic(err)
	}
	handl := ChapterHandler(stories, mux)
	if err = http.ListenAndServe(":8080", handl); err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)
	return mux
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not found Chapter")
}

func ChapterHandler(stories map[string]Story, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rPath := r.URL.Path
		if rPath == "/" {
			tpl.ExecuteTemplate(w, "chapter.gohtml", stories["intro"])
			return
		}
		rPath = strings.Replace(rPath, "/", "", -1)
		s, exist := stories[rPath]
		if exist {
			tpl.ExecuteTemplate(w, "chapter.gohtml", s)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
