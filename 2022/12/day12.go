package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type cell struct {
	i      int
	j      int
	height int
	status bool
}

var cols int
var rows int
var cells [][]cell
var stack []cell

func main() {
	fileName := flag.String("f", "test.txt", "input file name")

	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	cols = len(input[0])
	rows = len(input)

	for j := 0; j < rows; j++ {
		line := []cell{}
		for i := 0; i < cols; i++ {
			height := elevation(input[j][i])
			line = append(line, cell{i, j, height, false})
		}
		cells = append(cells, line)
	}
	// Recursive Backtracker
	// 1. Given a current cell as a parameter
	current := cells[0][0]

	for i := 0; i < 30; i++ {
		// 2. Mark the current cell as visited
		cells[current.j][current.i].visited()

		// 3. While the current cell has any unvisited neighbour cells
		// 3.1 Choose one of the unvisited neighbours
		next, check := chooseNeighbor(cells[current.j][current.i].checkNeighbor())

		if check == nil {
			// 3.3 Invoke the routine recursively for the chosen cell
			cells[next.j][next.i].visited()
			stack = append(stack, cells[current.j][current.i])
			current = cells[next.j][next.i]
		} else if len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
}

func (c *cell) visited() {
	c.status = true
}

func (c cell) checkNeighbor() []cell {
	// Clockwise from 12:	top[0],			right[1],		bottom[2], 		left[3]
	// neighbors := []int{	grid[j-1][i],	grid[j][i+1],	grid[j+1][i],	grid[j][i-1]}

	var neighbors []cell

	top := c.j == 0
	right := c.i == cols-1
	bottom := c.j == rows-1
	left := c.i == 0

	// corners
	if left && top {
		if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
			neighbors = append(neighbors, cells[c.j][c.i+1])
		}
		if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
			neighbors = append(neighbors, cells[c.j+1][c.i])
		}
		return neighbors
	}
	if right && top {
		if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
			neighbors = append(neighbors, cells[c.j][c.i-1])
		}
		if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
			neighbors = append(neighbors, cells[c.j+1][c.i])
		}
		return neighbors
	}
	if right && bottom {
		if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
			neighbors = append(neighbors, cells[c.j][c.i-1])
		}
		if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
			neighbors = append(neighbors, cells[c.j-1][c.i])
		}
		return neighbors
	}
	if left && bottom {
		if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
			neighbors = append(neighbors, cells[c.j][c.i+1])
		}
		if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
			neighbors = append(neighbors, cells[c.j-1][c.i])
		}
		return neighbors
	}

	topEdge := (c.i > 0 || c.i < cols-1) && top
	rightEdge := right && (c.j > 0 || c.j < rows-1)
	bottomEdge := (c.i > 0 || c.i < cols-1) && bottom
	leftEdge := left && (c.j > 0 || c.j < rows-1)

	//sides
	if topEdge {
		if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
			neighbors = append(neighbors, cells[c.j][c.i+1])
		}
		if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
			neighbors = append(neighbors, cells[c.j+1][c.i])
		}
		if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
			neighbors = append(neighbors, cells[c.j][c.i-1])
		}
		return neighbors
	}
	if rightEdge {
		if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
			neighbors = append(neighbors, cells[c.j-1][c.i])
		}
		if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
			neighbors = append(neighbors, cells[c.j][c.i-1])
		}
		if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
			neighbors = append(neighbors, cells[c.j+1][c.i])
		}
		return neighbors
	}
	if bottomEdge {
		if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
			neighbors = append(neighbors, cells[c.j-1][c.i])
		}
		if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
			neighbors = append(neighbors, cells[c.j][c.i+1])
		}
		if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
			neighbors = append(neighbors, cells[c.j][c.i-1])
		}
		return neighbors
	}
	if leftEdge {
		if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
			neighbors = append(neighbors, cells[c.j-1][c.i])
		}
		if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
			neighbors = append(neighbors, cells[c.j][c.i+1])
		}
		if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
			neighbors = append(neighbors, cells[c.j+1][c.i])
		}
		return neighbors
	}

	if !cells[c.j-1][c.i].status && condition(c.height, cells[c.j-1][c.i].height) {
		neighbors = append(neighbors, cells[c.j-1][c.i])
	}
	if !cells[c.j][c.i+1].status && condition(c.height, cells[c.j][c.i+1].height) {
		neighbors = append(neighbors, cells[c.j][c.i+1])
	}
	if !cells[c.j+1][c.i].status && condition(c.height, cells[c.j+1][c.i].height) {
		neighbors = append(neighbors, cells[c.j+1][c.i])
	}
	if !cells[c.j][c.i-1].status && condition(c.height, cells[c.j][c.i-1].height) {
		neighbors = append(neighbors, cells[c.j][c.i-1])
	}

	return neighbors
}

func chooseNeighbor(c []cell) (cell, error) {
	rand.Seed(time.Now().UnixNano())
	if len(c) > 0 {
		r := rand.Intn(len(c))
		return c[r], nil
	}
	return cell{}, errors.New("")
}

func condition(x, y int) bool {
	return y-1 == 0 || y-x == 1
}

func elevation(letter byte) int {
	if letter == 'S' {
		return -1
	} else if letter == 'E' {
		return 26
	}

	return int(letter) - 97
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
