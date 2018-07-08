package main

import (
	"errors"
	"io/ioutil"
)

var ErrorInvalidPath = errors.New("Invalid path to file")

func ReadFile(path string) ([]byte, error) {
	if path == "" {
		return nil, ErrorInvalidPath
	}
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func WriteToFile(path string) ([]byte, error) {
	if path == "" {
		return nil, ErrorInvalidPath
	}

	return []byte{}, nil
}
