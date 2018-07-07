package main

import (
	"encoding/json"
	"log"
)

func Parse(path string) (map[string]Story, error) {
	var stories map[string]Story
	data, err := ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read the file", path)
		return stories, err
	}
	err = json.Unmarshal(data, &stories)
	if err != nil {
		log.Fatal("Couldn't unmarshall stories:", err)
		return stories, err
	}

	return stories, nil
}
