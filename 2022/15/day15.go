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

type sensor struct {
	sx int
	sy int
	bx int
	by int
	d  int
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	checkY := flag.Int("y", 10, "row check of grid")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var positions []sensor
	minX := 10000000
	maxX := 0
	minY := 10000000
	maxY := 0
	maxD := 0

	for i := range input {
		var s sensor
		nodes := strings.Split(input[i], ": ")
		coor := strings.Split(nodes[0][10:], ", ")

		s.sx, _ = strconv.Atoi(coor[0][2:])
		s.sy, _ = strconv.Atoi(coor[1][2:])

		beacon := strings.Split(nodes[1][21:], ", ")

		s.bx, _ = strconv.Atoi(beacon[0][2:])
		s.by, _ = strconv.Atoi(beacon[1][2:])

		s.d = distance(s.sx, s.sy, s.bx, s.by)

		maxD = max(maxD, s.d)

		positions = append(positions, s)

		minX = min(minX, s.sx)
		minX = min(minX, s.bx)
		minY = min(minY, s.sy)
		minY = min(minY, s.by)
		maxX = max(maxX, s.sx)
		maxX = max(maxX, s.bx)
		maxY = max(maxY, s.sy)
		maxY = max(maxY, s.by)

	}

	noCoverage := make([]bool, maxX-minX+(2*maxD))
	x := minX - maxD
	part1 := 0
	for j := range noCoverage {
		for i := range positions {
			if positions[i].bx == x && positions[i].by == *checkY {
				break
			}
			if distance(positions[i].sx, positions[i].sy, x, *checkY) <= positions[i].d {
				noCoverage[j] = true
				part1++
				break

			}
		}
		x++
	}

	fmt.Println("y =", *checkY)
	fmt.Printf("Part 1: %v\n", part1)

	len := 20
	for y := 0; y < len; y++ {
		line := make([]bool, len)
		for x := 0; x < len; x++ {
			for i := range positions {
				if positions[i].bx == x && positions[i].by == y {
					break
				}
				if distance(positions[i].sx, positions[i].sy, x, y) <= positions[i].d {
					line[x] = true
					break
				}
			}

			//fmt.Printf("(%v, %v), %v\n", x, y, tuningFreq(x, y))

		}
		for x := range line {
			if !line[x] {
				fmt.Printf("(%v, %v): %v\n", x, y, tuningFreq(x, y))
			}
		}
		//printLine(line)
	}

}

func tuningFreq(x, y int) int {
	return x*4000000 + y
}

func printLine(arr []bool) {
	for i := range arr {
		if arr[i] {
			fmt.Print(".")
		} else {
			fmt.Print("F")
		}
	}
	fmt.Printf("\n")
}

func distance(x1, y1, x2, y2 int) int {
	d := math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2))
	return int(d)
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

func index(i, j, c int) int {
	return i + j*c
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
