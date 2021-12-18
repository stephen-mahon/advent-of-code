package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var title = "--- Day 10: Syntax Scoring ---"

func main() {
	fmt.Println(title)
	fileName := "test.txt"
	data, err := readFile(fileName)
	if err != nil {
		log.Fatalf("could not read file %q: %v", fileName, err)
	}
	var score int
	for i := range data {
		fmt.Println(len(data[i]), data[i])
	}

	fmt.Println("Score:", score)
}

func readFile(path string) (data [][]string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		var line []string
		for _, v := range s.Text() {
			line = append(line, string(v))
		}
		data = append(data, line)
	}
	return data, nil
}

func chunkEnd(v string) string {
	switch v {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	default:
		return ""
	}
}

func chunkScore(v string) int {
	switch v {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}
