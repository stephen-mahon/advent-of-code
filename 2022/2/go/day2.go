package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	pairs, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var part1, part2 int
	for i := range pairs {
		part1 += paperRockSissors(pairs[i])
		part2 += paperRockSissorsDetermined(pairs[i])
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func read(path string) ([][]string, error) {

	var pairs [][]string
	dat, err := os.Open("../data/" + path)
	if err != nil {
		return nil, err
	}
	defer dat.Close()

	s := bufio.NewScanner(dat)

	for s.Scan() {
		pairs = append(pairs, strings.Split(s.Text(), " "))
	}

	return pairs, nil
}

func paperRockSissors(pair []string) int {
	moves := map[string]string{"A": "Rock", "X": "Rock", "B": "Paper", "Y": "Paper", "C": "Scissors", "Z": "Scissors"}
	score := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3, "LOSS": 0, "TIE": 3, "WIN": 6}

	opponent := moves[pair[0]]
	response := moves[pair[1]]
	result := outcome(opponent, response)

	return score[response] + score[result]
}

func paperRockSissorsDetermined(pair []string) int {
	moves := map[string]string{"A": "Rock", "B": "Paper", "C": "Scissors", "X": "LOSE", "Y": "TIE", "Z": "WIN"}
	score := map[string]int{"Rock": 1, "Paper": 2, "Scissors": 3, "LOSS": 0, "TIE": 3, "WIN": 6}

	opponent := moves[pair[0]]
	result := moves[pair[1]]
	response := outcomeDetermined(opponent, result)

	return score[response] + score[result]
}

func outcome(opponent, response string) string {
	if opponent == response {
		return "TIE"
	} else if opponent == "Rock" && response == "Paper" {
		return "WIN"
	} else if opponent == "Rock" && response == "Scissors" {
		return "LOSE"
	} else if opponent == "Paper" && response == "Scissors" {
		return "WIN"
	} else if opponent == "Paper" && response == "Rock" {
		return "LOSE"
	} else if opponent == "Scissors" && response == "Rock" {
		return "WIN"
	} else if opponent == "Scissors" && response == "Paper" {
		return "LOSE"
	} else {
		return ""
	}
}

func outcomeDetermined(opponent, result string) string {
	if result == "TIE" {
		return opponent
	} else if result == "LOSE" && opponent == "Rock" {
		return "Scissors"
	} else if result == "LOSE" && opponent == "Paper" {
		return "Rock"
	} else if result == "LOSE" && opponent == "Scissors" {
		return "Paper"
	} else if result == "WIN" && opponent == "Rock" {
		return "Paper"
	} else if result == "WIN" && opponent == "Paper" {
		return "Scissors"
	} else if result == "WIN" && opponent == "Scissors" {
		return "Rock"
	} else {
		return ""
	}
}
