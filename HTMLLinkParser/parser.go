package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `"json":"href"`
	Text string `"json":"text"`
}

type Document struct {
	Links []Link
}

const (
	input  = "resources/index.html"
	output = "resources/out.json"
)

func main() {
	data, err := ReadFile(input)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
