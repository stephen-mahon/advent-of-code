package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	rucksacks := tabulate(input)
	var sum int
	for i := range rucksacks {
		bag1 := rucksacks[i][0]
		bag2 := rucksacks[i][1]

		for j := range bag1 {
			letter := string(bag1[j])
			if strings.Contains(bag2, letter) {
				if IsLower(letter) {
					sum += int(bag1[j]) - 96
				} else if IsUpper(letter) {
					sum += int(bag1[j]) - 38
				}
				break
			}
		}
	}

	fmt.Println("Part 1:", sum)

	sum = 0
	g := make([]string, 3)
	for i := range input {
		g[i%3] = input[i]
		if i != 0 && (i+1)%3 == 0 {
			for _, letter := range g[0] {
				if strings.Contains(g[1], string(letter)) && strings.Contains(g[2], string(letter)) {
					if IsLower(string(letter)) {
						sum += int(letter) - 96
					} else if IsUpper(string(letter)) {
						sum += int(letter) - 38
					}
					break
				}
			}
		}
	}

	fmt.Println("Part 2:", sum)

}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
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

func tabulate(input []string) [][]string {
	var output [][]string

	for i := range input {
		n := len(input[i]) / 2
		output = append(output, []string{input[i][:n], input[i][n:]})
	}

	return output
}
