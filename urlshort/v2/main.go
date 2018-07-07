package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	yamlFilePath string
	jsonFilePath string
)

func init() {
	flag.StringVar(&yamlFilePath, "Yaml", "resources/redirects.yaml", "Provide path to YAML file")
	flag.StringVar(&jsonFilePath, "Json", "resources/redirects.json", "Provide path to JSON file")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/bolt-doc":   "https://godoc.org/github.com/boltdb/bolt",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ReadFile(yamlFilePath)
	checkPanic(err)
	yamlHandler, err := YAMLHandler(yaml, mapHandler)
	checkPanic(err)

	//Build the JSONHandler using the mapHandler as the
	//fallback
	json, err := ReadFile(jsonFilePath)
	checkPanic(err)
	jsonHandler, err := JSONHandler(json, yamlHandler)
	checkPanic(err)

	fmt.Println("Starting the server on :8080")
	if err = http.ListenAndServe(":8080", jsonHandler); err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}
