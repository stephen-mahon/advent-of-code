package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var part2 = 0

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var grid [][]int
	var line []int
	for i := range input {
		line = []int{}
		for j := range input[i] {
			v, _ := strconv.Atoi(string(input[i][j]))
			line = append(line, v)
		}
		grid = append(grid, line)
	}

	part1 := 4*len(input[0]) - 4
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if check(i, j, grid) {
				part1++
			}
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func check(x, y int, trees [][]int) bool {
	//var visable int
	dir := make(map[string][]int)
	var n, e, s, w []int

	for i := range trees {
		for j := range trees[i] {
			if x == i && y == j {
			} else if x == i || y == j {
				if i < x {
					n = append(n, trees[i][j])
				} else if i > x {
					s = append(s, trees[i][j])
				} else if j < y {
					w = append(w, trees[i][j])
				} else if j > y {
					e = append(e, trees[i][j])
				}
			}
		}
	}

	dir["N"] = n
	dir["E"] = e
	dir["S"] = s
	dir["W"] = w

	var sum int
	var visable bool
	var treeCount int
	var scenicScore = 1

	//fmt.Printf("\ntrees[%v][%v] = %v\n", x, y, trees[x][y])
	for i, v := range dir {
		visable = true
		treeCount = 0
		//fmt.Printf("%v, %v ", i, v)

		if i == "S" || i == "E" {
			for j := range v {
				if visable {
					treeCount++
				}
				if v[j] >= trees[x][y] {
					visable = false
				}
			}
		} else {
			for j := len(v) - 1; j >= 0; j-- {
				if visable {
					treeCount++
				}
				if v[j] >= trees[x][y] {
					visable = false
				}
			}

		}
		//fmt.Printf("visable: %v, treeCount: %v\n", visable, treeCount)
		scenicScore *= treeCount
		if visable {
			sum++
		}
	}

	//fmt.Println("scenicScore", scenicScore)
	part2 = max(part2, scenicScore)
	return sum > 0
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
