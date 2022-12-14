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

	max_x := 0
	max_y := 0
	min_x := 10000

	for j := range input {
		rockLine := parse(input[j])
		for i := range rockLine {
			//fmt.Println(rockLine[i])
			max_x = max(max_x, rockLine[i][0])
			min_x = min(min_x, rockLine[i][0])
			max_y = max(max_y, rockLine[i][1])
		}
	}

	cols := 1 + max_x - min_x
	rows := 2 + max_y
	rocks := make([]bool, (cols)*(rows))

	for j := range input {
		corners := parse(input[j])
		for i := 0; i < len(corners)-1; i++ {
			x1 := corners[i][0] - min_x
			y1 := corners[i][1]
			x2 := corners[i+1][0] - min_x
			y2 := corners[i+1][1]

			if y1 == y2 {
				for k := min(x1, x2); k <= max(x1, x2); k++ {
					rocks[k+y1*cols] = true
				}
			} else {
				for k := min(y1, y2); k <= max(y1, y2); k++ {
					rocks[x1+k*cols] = true
				}
			}

		}
	}

	part1 := 0
	sand := rocks
	fin := true
	for fin {
		fin = hourglass(sand, min_x, max_y, cols)
		if !fin {
			break
		}
		part1++
	}

	fmt.Println("Part 1:", part1)

	// part 2

	cols = max_x * 2
	rows = 2 + max_y
	rocks = make([]bool, (cols)*(rows))

	for j := range input {
		corners := parse(input[j])
		for i := 0; i < len(corners)-1; i++ {
			x1 := corners[i][0] - min_x
			y1 := corners[i][1]
			x2 := corners[i+1][0] - min_x
			y2 := corners[i+1][1]

			if y1 == y2 {
				for k := min(x1, x2); k <= max(x1, x2); k++ {
					rocks[k+y1*cols] = true
				}
			} else {
				for k := min(y1, y2); k <= max(y1, y2); k++ {
					rocks[x1+k*cols] = true
				}
			}
		}
	}

	part2 := 0
	sand = rocks
	fin = true

	for {
		coor := hourglass2(sand, max_y, cols)
		x := coor[0]
		y := coor[1]

		sand[x+y*cols] = true
		part2++

		if x == 500 && y == 0 {
			break
		}

	}

	fmt.Println("Part 2:", part2)

}

func hourglass(rocks []bool, min_x, max_y, cols int) bool {
	x := 500 - min_x
	y := 0

	for y <= max_y {
		if !(rocks[x+(y+1)*cols]) {
			y++
			continue
		}
		if !(rocks[(x-1)+(y+1)*cols]) {
			x--
			y++
			continue
		}
		if !(rocks[(x+1)+(y+1)*cols]) {
			x++
			y++
			continue
		}
		rocks[x+y*cols] = true
		return true
	}

	return false
}

func hourglass2(rocks []bool, max_y, cols int) []int {
	x := 500
	y := 0

	if rocks[x+y*cols] == true {
		return []int{x, y}
	}

	for y <= max_y {
		if !(rocks[x+(y+1)*cols]) {
			y++
			continue
		}
		if !(rocks[(x-1)+(y+1)*cols]) {
			x--
			y++
			continue
		}
		if !(rocks[(x+1)+(y+1)*cols]) {
			x++
			y++
			continue
		}

		break
	}

	return []int{x, y}
}

func draw(grid []bool, cols, rows int, fill string) {

	for j := 0; j < rows; j++ {

		for i := 0; i < cols; i++ {
			if grid[i+j*cols] {
				fmt.Print(fill)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
