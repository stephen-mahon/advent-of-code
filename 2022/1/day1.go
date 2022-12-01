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

// var enable = "enable1.txt"

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	dat, err := readCalories(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	top := topThree(dat)
	fmt.Println("Part 1:", top[0])
	fmt.Println("Part 2:", total(top))

}

func readCalories(path string) (vals [][]float64, err error) {

	dat, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dat.Close()

	s := bufio.NewScanner(dat)
	val := []float64{}

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			vals = append(vals, val)
			val = []float64{}

		} else {
			cal, _ := strconv.ParseFloat(line, 64)
			val = append(val, cal)
		}
	}
	vals = append(vals, val)

	return vals, nil
}

func topThree(vals [][]float64) []float64 {

	var total float64
	top := make([]float64, 3)

	for _, v := range vals {
		total = 0
		for i := range v {
			total += v[i]
		}

		if total > top[0] {
			top[2] = top[1]
			top[1] = top[0]
			top[0] = total
		} else if total > top[1] {
			top[2] = top[1]
			top[1] = total
		} else if total > top[2] {
			top[2] = total
		}
	}

	return top
}

func total(vals []float64) float64 {
	var total float64
	for i := range vals {
		total += vals[i]
	}
	return total
}
