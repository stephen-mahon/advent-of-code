package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var title = "Day 5: Hydrthermal Venture"

type coor struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	fmt.Println(title)
	data, _ := readFile("input.txt")

	board := [999][999]int{}

	board = drawLine(data, board)
	var count int
	for i := range board {
		for j := range board[i] {
			if board[i][j] > 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func readFile(path string) (data []coor, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(s.Text(), "%v,%v -> %v,%v", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}

		data = append(data, coor{x1, y1, x2, y2})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return data, nil
}

func drawLine(data []coor, board [999][999]int) [999][999]int {
	for _, v := range data {
		if v.x1 == v.x2 {
			for y := min(v.y1, v.y2); y <= max(v.y1, v.y2); y++ {
				board[v.x1][y] += 1
			}
		} else if v.y1 == v.y2 {
			for x := min(v.x1, v.x2); x <= max(v.x1, v.x2); x++ {
				board[x][v.y1] += 1
			}
		}
	}
	return board
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
