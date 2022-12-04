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

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var sum1 int
	var sum2 int

	for i := range input {
		pairs := strings.Split(input[i], ",")

		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		l1, _ := strconv.Atoi(p1[0])
		u1, _ := strconv.Atoi(p1[1])
		l2, _ := strconv.Atoi(p2[0])
		u2, _ := strconv.Atoi(p2[1])

		if (l1 <= l2 && u1 >= u2) || (l2 <= l1 && u2 >= u1) {
			sum1 += 1
		}
		if u1 >= l2 && u2 >= l1 {
			sum2 += 1
		}
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
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
