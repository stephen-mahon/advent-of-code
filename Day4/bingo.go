package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var title = "Bingo!"

func main() {
	fmt.Println(title)
	fmt.Println(readBingo("test.txt"))
}

func readBingo(path string) (draw []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var lines []string
	//var arrays [][]string
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	line := strings.Split(lines[0], ",")
	lines = lines[1:]

	for _, v := range line {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		draw = append(draw, val)
	}
	var boards [][][]int
	for i := 1; i < len(lines); i++ {
		board := [][]int{}

		if lines[i] != "" {
			vals := strings.Split(lines[i], " ")
			for _, v := range vals {
				b := []int{}
				v = strings.Trim(v, " ")
				if v != "" {
					val, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}
					fmt.Println(val)
				}
			}
		}
	}
	fmt.Println(boards)

	return draw, nil
}
