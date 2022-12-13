package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type cell struct {
	i       int
	j       int
	height  int
	visited bool
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")

	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}
	rows := len(input)
	cols := len(input[0])

	var grid []cell

	var current cell
	for j := range input {
		for i := range input[j] {
			height := elevation(string(input[j][i]))
			grid = append(grid, cell{i, j, height, false})
			if height == -1 {

				current = cell{j, i, height, true}
			}
		}
	}

	//var steps int
	fmt.Println(current)

}

func index(i, j, cols int) int {
	return i + j*cols
}

func (c *cell) checkNeighbor(i, j int) []cell {
	var neighbors []cell

	return neighbors
}

/*
func checkNeighbor(j, i int, grid [][]int) []bool {
	// up, right, down, left. Clockwise from 12
	//neighbors := []int{grid[j-1][i], grid[j][i+1], grid[j+1][i], grid[j][i-1]}

	top := j == 0
	right := i == len(grid[j])-1
	bottom := j == len(grid)-1
	left := i == 0
	//Corners
	if left && top {
		return []bool{false, condition(grid[j+1][i], grid[j][i]), condition(grid[j][i+1], grid[j][i]), false}
	}
	if right && top {
		return []bool{false, false, condition(grid[j+1][i], grid[j][i]), condition(grid[j][i-1], grid[j][i])}
	}
	if right && bottom {
		return []bool{condition(grid[j-1][i], grid[j][i]), false, false, condition(grid[j][i-1], grid[j][i])}
	}
	if left && bottom {
		return []bool{condition(grid[j-1][i], grid[j][i]), condition(grid[j][i+1], grid[j][i]), false, false}
	}

	topEdge := i > 0 || (i < len(grid[j])-1) && top
	rightEdge := right && ((j > 0) || (j < len(grid)-1))
	bottomEdge := (i > 0 || (i < len(grid[j])-1)) && bottom
	leftEdge := left && (j > 0 || j < len(grid)-1)
	//sides
	if topEdge {
		return []bool{false, condition(grid[j][i+1], grid[j][i]), condition(grid[j+1][i], grid[j][i]), condition(grid[j][i-1], grid[j][i])}
	}
	if rightEdge {
		return []bool{condition(grid[j-1][i], grid[j][i]), false, condition(grid[j+1][i], grid[j][i]), condition(grid[j][i-1], grid[j][i])}
	}
	if bottomEdge {
		return []bool{condition(grid[j-1][i], grid[j][i]), condition(grid[j][i+1], grid[j][i]), false, condition(grid[j][i-1], grid[j][i])}
	}
	if leftEdge {
		return []bool{condition(grid[j-1][i], grid[j][i]), condition(grid[j][i+1], grid[j][i]), condition(grid[j+1][i], grid[j][i]), false}
	}

	return []bool{condition(grid[j-1][i], grid[j][i]), condition(grid[j][i+1], grid[j][i]), condition(grid[j+1][i], grid[j][i]), condition(grid[j][i-1], grid[j][i])}
}
*/
func condition(x, y int) bool {
	return x-y == 0 || x-y == 1
}

func elevation(line string) int {
	var height int
	for i := range line {
		if line[i] == 'S' {
			return -1
		} else if line[i] == 'E' {
			return 26
		} else {
			return int(line[i]) - 97
		}
	}
	return height
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
