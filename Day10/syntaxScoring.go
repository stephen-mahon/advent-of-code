package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var title = "--- Day 10: Syntax Scoring ---"
var bracketOpen = "([{<"

func main() {
	fmt.Println(title)
	fileName := "input.txt"
	data, err := readFile(fileName)
	if err != nil {
		log.Fatalf("could not read file %q: %v", fileName, err)
	}
	var score int
	for i := range data {
		stack := []string{}
		for j := 0; j < len(data[i]); j++ {
			if strings.ContainsAny(data[i][j], bracketOpen) {
				stack = append(stack, data[i][j])
			} else if data[i][j] == chunkSwitch(string(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
			} else {
				log.Printf("Ln %v, Col %v - Expected: `%v`, but found `%v` instead\n", i, j, chunkSwitch(stack[len(stack)-1]), data[i][j])
				score += chunkScore(data[i][j])
				break
			}
		}
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

func chunkSwitch(v string) string {
	switch v {
	case ")":
		return "("
	case "(":
		return ")"
	case "]":
		return "["
	case "[":
		return "]"
	case "}":
		return "{"
	case "{":
		return "}"
	case ">":
		return "<"
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
