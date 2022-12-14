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

const cols = 10
const rows = 10

type cell struct {
	i    int
	j    int
	rock bool
	sand bool
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	grid := make([]cell, rows*cols)
	for j := range input {
		rockLine := parse(input[j])
		for i := 0; i < len(rockLine)-1; i++ {
			x1 := rockLine[i][0] - 494
			y1 := rockLine[i][1]
			x2 := rockLine[i+1][0] - 494
			y2 := rockLine[i+1][1]

			if y1 == y2 {
				for k := min(x1, x2); k <= max(x1, x2); k++ {
					grid[k+y1*cols].rock = true
				}
			} else {
				for k := min(y1, y2); k <= max(y1, y2); k++ {
					grid[x1+k*cols].rock = true
				}
			}

		}
	}

	x := 6
	y := 0

	n := 0
	for n < 10 {
		current := x + (y+n)*cols
		grid[current].sand = true
		draw(grid)
		grid[current].sand = false
		n++
	}

}

func draw(grid []cell) {
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			if grid[i+j*cols].rock {
				fmt.Print("#")
			} else if grid[i+j*cols].sand {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
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

func parse(line string) [][]int {
	var output [][]int
	pairs := strings.Split(line, " -> ")
	for i := range pairs {
		vals := strings.Split(pairs[i], ",")
		x, _ := strconv.Atoi(vals[0])
		y, _ := strconv.Atoi(vals[1])
		output = append(output, []int{x, y})
	}
	return output
}
