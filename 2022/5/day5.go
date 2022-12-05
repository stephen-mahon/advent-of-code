package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	num   int
	start int
	end   int
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	stacks1 := parseStacks(input)
	stacks2 := parseStacks(input)

	var moves []move

	for i := range input {
		if strings.Contains(input[i], "move") {
			m, err := parseMoves(input[i])
			if err != nil {
				log.Fatal(err)
			}
			moves = append(moves, m)
		}

	}

	for i := range moves {
		stacks1 = arrange1(stacks1, moves[i])
		stacks2 = arrange2(stacks2, moves[i])
	}

	var part1, part2 string
	for i := range stacks1 {
		part1 += stacks1[i][len(stacks1[i])-1]
		part2 += stacks2[i][len(stacks2[i])-1]
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func arrange1(stacks [][]string, move move) [][]string {
	for i := 0; i < move.num; i++ {
		val, newStack := pop(stacks[move.start])
		stacks[move.start] = newStack
		stacks[move.end] = append(stacks[move.end], val)

	}

	return stacks
}

func arrange2(stacks [][]string, move move) [][]string {

	newStack, vals := popArray(move.num, stacks[move.start])
	stacks[move.start] = newStack
	for i := range vals {
		stacks[move.end] = append(stacks[move.end], vals[i])
	}

	return stacks
}

func pop(s []string) (string, []string) {
	return s[len(s)-1], s[:len(s)-1]
}

func popArray(num int, s []string) ([]string, []string) {
	return s[:len(s)-num], s[len(s)-num:]
}

func read(path string) ([]string, error) {
	var input []string
	dat, err := os.Open("data/" + path)
	if err != nil {
		return nil, err
	}
	defer dat.Close()

	s := bufio.NewScanner(dat)

	for s.Scan() {
		input = append(input, s.Text())
	}

	return input, nil
}

func parseMoves(input string) (move, error) {

	values := strings.Split(input, " ")

	num, err := strconv.Atoi(values[1])
	if err != nil {
		return move{}, err
	}

	start, err := strconv.Atoi(values[3])
	if err != nil {
		return move{}, err
	}

	end, err := strconv.Atoi(values[5])
	if err != nil {
		return move{}, err
	}

	return move{num, start - 1, end - 1}, nil
}

func parseStacks(input []string) [][]string {

	var vals []string

	for i := range input {
		str := string(input[i])
		if str != "" {
			vals = append(vals, str)
		} else {
			break
		}
	}
	vals = vals[:len(vals)-1]

	var containers []string

	for i := len(vals) - 1; i >= 0; i-- {
		str := vals[i]
		for j := range str {
			if (j-1)%4 == 0 {
				containers = append(containers, string(str[j]))
			}
		}
	}
	n := (len(input[0]) + 1) / 4
	stacks := make([][]string, n)

	for i := range containers {
		if strings.ContainsAny(containers[i], "QWERTYUIOPASDFGHJKLZXCVBNM") {
			stacks[(i+n)%n] = append(stacks[(i+n)%n], containers[i])
		}
	}

	return stacks
}
