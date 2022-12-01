package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var title = "Day 8: Seven Segment Search"

var digit = " aaaa \n" +
	"b    c\n" +
	"b    c\n" +
	" dddd \n" +
	"e    f\n" +
	"d    f\n" +
	" gggg "

func main() {
	fmt.Println(title)
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	var count int
	for _, v := range data {
		if len(v) == 2 {
			count++
		} else if len(v) == 4 {
			count++
		} else if len(v) == 3 {
			count++
		} else if len(v) == 7 {
			count++
		}
	}
	fmt.Println(count)

}

func readFile(path string) (data []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		splitStr := strings.SplitAfter(s.Text(), " | ")
		combos := strings.Split(splitStr[1], " ")
		for _, v := range combos {
			data = append(data, v)
		}
	}

	return data, nil
}

func strBin(num int) string {
	return strconv.FormatInt(int64(num), 2)
}

func binLU(bin string) (string, error) {
	switch bin {
	case strBin(0):
		return "abcefg", nil
	case strBin(1):
		return "cf", nil
	case strBin(2):
		return "acdeg", nil
	case strBin(3):
		return "acdfg", nil
	case strBin(4):
		return "bcdf", nil
	case strBin(5):
		return "abdfg", nil
	case strBin(6):
		return "abdefg", nil
	case strBin(7):
		return "acf", nil
	case strBin(8):
		return "abcdefg", nil
	case strBin(9):
		return "abcdfg", nil
	default:
		return "", fmt.Errorf("not a valid binary number %v", bin)

	}
}

func segLU(seg string) (int, error) {
	switch seg {
	case "abcefg":
		return 0, nil
	case "cf":
		return 1, nil
	case "acdeg":
		return 2, nil
	case "acdfg":
		return 3, nil
	case "bcdf":
		return 4, nil
	case "abdfg":
		return 5, nil
	case "abdefg":
		return 6, nil
	case "acf":
		return 7, nil
	case "abcdefg":
		return 8, nil
	case "abcdfg":
		return 9, nil
	default:
		return -1, fmt.Errorf("not a valid seven segment sequence: %v", seg)

	}
}

func display(seg string) (string, error) {
	_, err := segLU(seg)
	if err != nil {
		return "", err
	}
	var num string
	for _, v := range digit {
		if strings.ContainsAny(string(v), seg) || strings.ContainsAny(string(v), " \n") {
			num += fmt.Sprint(string(v))
		} else {
			num += fmt.Sprint(".")
		}
	}
	return num, nil
}
