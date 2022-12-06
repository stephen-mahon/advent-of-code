package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	for i := range input {
		for j := range input[i] {
			if j+4 <= len(input[i]) && check(input[i][j:j+4]) {
				fmt.Println("Part 1:", j+4)
				break
			}
		}
	}

	fmt.Println()

	for i := range input {
		for j := range input[i] {
			if j+14 <= len(input[i]) && check(input[i][j:j+14]) {
				fmt.Println("Part 2:", j+14)
				break
			}
		}
	}

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

func check(input string) bool {
	m := map[byte]int{}
	for i := range input {
		if _, exists := m[input[i]]; exists {
			return false
		} else {
			m[input[i]] = 0
		}
	}
	return true
}
