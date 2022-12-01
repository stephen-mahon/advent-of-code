package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var title = "Bingo!"

func main() {
	fmt.Println(title)
	file, _ := readFile("test.txt")

	readout := strings.Split(file[0], ",")

	for i := range readout {
		fmt.Printf("%s ", readout[i])
	}
	fmt.Println()

	boards := bingoBoard(file[2:])

	fmt.Println(boards)

	var boolBoards [][][]bool
	var line [][]bool
	var val []bool
	for i := range boards {
		for j := range boards[i] {
			for k := range boards[i][j] {
				_ = k
				val = append(val, false)
			}
			line = append(line, val)
			val = []bool{}
		}
		boolBoards = append(boolBoards, line)
		line = [][]bool{}
	}
}

func readFile(path string) (vals []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		vals = append(vals, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return vals, nil
}

func bingoBoard(boards []string) [][][]string {

	var bingoTables [][][]string
	var line [][]string
	var num []string

	for i := range boards {
		ln := strings.Split(boards[i], " ")

		if len(ln) > 1 {
			num = []string{}
			for j := range ln {
				if ln[j] != "" {
					num = append(num, ln[j])
				}
			}
			line = append(line, num)
		} else {
			bingoTables = append(bingoTables, line)
			line = [][]string{}
		}
	}
	bingoTables = append(bingoTables, line)

	return bingoTables
}
