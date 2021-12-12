package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var title = "Day 7: The Treachery of Whales"

func main() {
	fmt.Println(title)
	data, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	med := median(data)
	fmt.Println(data, med)
	var fuelTotal float64
	for _, v := range data {
		fuelTotal += math.Abs(float64(v - med))
	}
	fmt.Println("Fuel total:", fuelTotal)
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

func median(data []int) int {
	sort.Ints(data)

	var median int
	if len(data)%2 == 0 {

		median = (data[(len(data)-1)/2] + data[(len(data)+1)/2]) / 2
	} else {
		median = (data[(len(data)-1)/2])
	}
	return median
}
