package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var title = "Bingo!"

func main() {
	fmt.Println(title)
	file, _ := readFile("test.txt")

	draw := strings.Split(file[0], ",")
	readout, err := strToInt(draw)
	if err != nil {
		log.Fatalf("could not read first line of input file: %v\n", err)
	}

	file = file[1:]
	var bingoData []string
	for i := range file {
		if file[i] != "" {
			bingoData = append(bingoData, file[i])
		}
	}

	tables, err := bingoBoard(bingoData)
	if err != nil {
		log.Fatalf("could not convert bingo tables: %v\n", err)
	}

	booltables := boolBingo(tables)

	fmt.Printf("%v\n\n", readout)

	for i := range readout {
		booltables = boolCheck(readout[i], tables, booltables)
		fmt.Printf("%v, ", readout[i])
		check, val := bingoCheck(booltables)
		if check {
			fmt.Println()
			fmt.Println(tables[val])
			break
		}
	}

}

func strToInt(vals []string) ([]int, error) {
	var draw []int
	for i := range vals {
		val, err := strconv.Atoi(vals[i])
		if err != nil {
			return nil, err
		}
		draw = append(draw, val)
	}
	return draw, nil
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

func bingoBoard(input []string) (bingo [][][]int, err error) {
	var matrix [][]int
	for i := range input {
		var vals []int
		line := strings.Split(input[i], " ")
		for j := range line {
			if line[j] != "" {
				val, err := strconv.Atoi(line[j])
				if err != nil {
					return nil, err
				}
				vals = append(vals, val)
			}
		}
		matrix = append(matrix, vals)
		if (i+1)%5 == 0 {
			bingo = append(bingo, matrix)
			matrix = [][]int{}
		}
	}
	return bingo, nil
}

func boolLine(input []int) []bool {
	return make([]bool, len(input))
}

func boolBingo(input [][][]int) (bingo [][][]bool) {
	var matrix [][]bool
	for i := range input {
		for j := range input[i] {
			matrix = append(matrix, boolLine(input[i][j]))
		}
		bingo = append(bingo, matrix)
		matrix = [][]bool{}
	}

	return bingo
}

func printTables(ints [][][]int, bools [][][]bool) {
	for i := range ints {
		for j := range ints[i] {
			fmt.Println(ints[i][j])
		}
		for j := range bools[i] {
			fmt.Println(bools[i][j])
		}
		fmt.Println()
	}
}

func boolCheck(check int, ints [][][]int, bools [][][]bool) [][][]bool {
	for i := range ints {
		for j := range ints[i] {
			for k := range ints[i][j] {
				if ints[i][j][k] == check {
					bools[i][j][k] = true
				}
			}
		}
	}
	return bools
}

func bingoCheck(bools [][][]bool) (bingo bool, i int) {

	for i := range bools {
		for j := range bools[i] {
			bingo = true
			for k := range bools[i][j] {
				bingo = bingo && bools[i][j][k]
			}
			if bingo == true {
				return bingo, i
			}
		}
	}
	return false, -1
}
