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
	fileName := "test.txt"
	data, err := readFile(fileName)
	if err != nil {
		log.Fatalf("could not read file %q: %v", fileName, err)
	}
	stacks := part2(data)
	for i := range stacks {
		//for i := 9; i < 10; i++ {
		score := 0
		for j := len(stacks[i]) - 1; j >= 0; j-- {
			closer := chunkSwitch(stacks[i][j])
			score = (5 * score) + chunkScore2(closer)
		}
		fmt.Println(score)
	}
}

// Some problem here with breaking from data.
// Adding to the stack when there is corrupted line.
// Probably with the break - only breaking from the inner loop (j) and not the outter (i)
func part2(data [][]string) [][]string {
	var lines [][]string
	var stacks [][]string
	for i := range data {
		stack := []string{}
		for j := 0; j < len(data[i]); j++ {
			if strings.ContainsAny(data[i][j], bracketOpen) {
				stack = append(stack, data[i][j])
			} else if data[i][j] == chunkSwitch(string(stack[len(stack)-1])) {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		lines = append(lines, data[i])
		stacks = append(stacks, stack)
		fmt.Printf("%v: %v - \"%v\"\n", i, data[i], stack)
	}
	return stacks
}

func part1(data [][]string) string {
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
	return fmt.Sprintf("Score: %v", score)
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

func chunkScore2(v string) int {
	switch v {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	default:
		return 0
	}
}
