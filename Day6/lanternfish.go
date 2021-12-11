package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var title = "Day 6: Lanternfish"

func main() {
	fmt.Println(title)
	data, _ := readFile("input.txt")
	fmt.Println("Initial state:", data)

	var day int
	fmt.Println()
	for day < 256 {
		day++
		data = subSlice(data)
		fmt.Printf("After %v day: %v\n", day, len(data))
	}

}

func readFile(path string) (data []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		dataStr := strings.Split(s.Text(), ",")
		for _, v := range dataStr {
			val, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("could not scan: %v", err)
			}
			data = append(data, val)
		}
	}

	return data, nil
}

func subSlice(s []int) []int {
	for i := range s {
		s[i] -= 1
		if s[i] < 0 {
			s[i] = 6
			s = append(s, 8)
		}
	}
	return s
}
