package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
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

	var grid [6][40]bool

	var cycle = 0
	var i = 0
	var x = 1
	var signalStrength = 0
	var part1 int

	for i < len(input) {
		if input[i] == "noop" {
			cycle++
			signalStrength, grid = checkTick(cycle, x, grid)
			part1 += signalStrength

		} else {
			_, num := parseInput(input[i])
			cycle++
			signalStrength, grid = checkTick(cycle, x, grid)
			part1 += signalStrength
			cycle++
			signalStrength, grid = checkTick(cycle, x, grid)
			part1 += signalStrength
			x += num
		}
		i++
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:")
	for _, line := range grid {
		for _, i := range line {
			if i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func checkTick(cycle, x int, grid [6][40]bool) (int, [6][40]bool) {
	vals := []int{20, 60, 100, 140, 180, 220}

	tick := cycle - 1
	if math.Abs(float64(x-tick%40)) <= 1 {
		grid[tick/40][tick%40] = true
	} else {
		grid[tick/40][tick%40] = false
	}

	for i := range vals {
		if cycle == vals[i] {
			return vals[i] * x, grid
		}
	}

	return 0, grid
}

func parseInput(input string) (string, int) {
	vals := strings.Split(input, " ")
	num, _ := strconv.Atoi(vals[1])
	return vals[0], num
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
