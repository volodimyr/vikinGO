package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	path = "problems.csv"
)

var (
	limit   int
	checked int
)

func init() {
	flag.IntVar(&limit, "Limit", 10, "How long it would take to answer whole questions?")
	flag.Parse()
}

func main() {
	game, err := CSVparse()
	if err != nil {
		log.Fatalln(err)
		return
	}

	end := make(chan bool)
	go startGame(game, end)
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("\nGame Over")
			return
		case <-end:
			fmt.Printf("Done [%d/%d]\n", len(game), checked)
			return
		}
	}
}

func CSVparse() (map[string]int, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()

	game := make(map[string]int)
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) <= 1 {
			return nil, fmt.Errorf("CSV should have [question-answer] form")
		}
		ans, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("Answer should be integer")
		}
		game[record[0]] = ans
	}
	return game, nil
}

func startGame(game map[string]int, end chan<- bool) {
	var input int
	for k, v := range game {
		fmt.Printf("Question: %v is ", k)
		_, err := fmt.Scanf("%d", &input)
		checkErrPanic(err, "")
		if input == v {
			checked++
		}
	}
	end <- true
}

func checkErrPanic(err error, context string) {
	if err != nil {
		fmt.Errorf("%s", context)
		panic(err)
	}
}
