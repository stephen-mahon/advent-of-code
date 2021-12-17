package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var title = "--- Day 9: Smoke Basin ---"

func main() {
	fmt.Println(title)
	fileName := "input.txt"
	data, err := readFile(fileName)
	if err != nil {
		log.Fatalf("could not read file %q: %v", fileName, err)
	}

	truths := lowVals(data)
	fmt.Println(sumHightMap(data, truths))

}

func readFile(path string) (data [][]int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var line []int
		for _, v := range s.Text() {
			val, err := strconv.Atoi(string(v))
			if err != nil {
				log.Printf("discarding bad data point %q: %v", v, err)
				continue
			}
			line = append(line, val)
		}
		data = append(data, line)
	}
	return data, nil
}

func truthArray(data [][]int) [][]bool {
	var truths [][]bool
	for i := range data {
		var line []bool
		for j := range data[i] {
			_ = j
			line = append(line, false)
		}
		truths = append(truths, line)
	}
	return truths
}

func lowVals(data [][]int) [][]bool {
	truths := truthArray(data)
	for i := range data {
		for j := range data[i] {
			// Top Left Corner
			if (i == 0) && (j == 0) {
				if (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j+1]) {
					truths[i][j] = true
				}
			}
			// Top Right Corner
			if (i == 0) && (j == len(data[i])-1) {
				if (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j-1]) {
					truths[i][j] = true
				}
			}
			// Top Right Corner
			if (i == 0) && !(j == 0 || j == len(data[i])-1) {
				if (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j-1]) && (data[i][j] < data[i][j+1]) {
					truths[i][j] = true
				}
			}
			// Middle
			if !(i == 0 || i == len(data)-1) && !(j == 0 || j == len(data[i])-1) {
				if midCond(i, j, data) {
					truths[i][j] = true
				}
			}
			// Left Column
			if !(i == 0 || i == len(data)-1) && j == 0 {
				if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j+1]) {
					truths[i][j] = true
				}
			}
			// Right Column
			if !(i == 0 || i == len(data)-1) && (j == len(data[i])-1) {
				if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j-1]) {
					truths[i][j] = true
				}
			}
			// Bottom Left Corner
			if (i == len(data)-1) && (j == 0) {
				if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i][j+1]) {
					truths[i][j] = true
				}
			}
			// Bottom Row
			if (i == len(data)-1) && !(j == 0 || j == len(data[i])-1) {
				if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i][j-1]) && (data[i][j] < data[i][j+1]) {
					truths[i][j] = true
				}
			}
			// Bottom Right Corner
			if (i == len(data)-1) && (j == len(data[i])-1) {
				if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i][j-1]) {
					truths[i][j] = true
				}
			}
		}
	}
	return truths
}

func midCond(i, j int, data [][]int) bool {
	if (data[i][j] < data[i-1][j]) && (data[i][j] < data[i+1][j]) && (data[i][j] < data[i][j-1]) && (data[i][j] < data[i][j+1]) {
		return true
	}
	return false
}

func sumHightMap(data [][]int, truths [][]bool) int {
	var count int
	for i := range truths {
		for j := range truths[i] {
			if truths[i][j] {
				count += data[i][j] + 1
			}
		}
	}
	return count
}
