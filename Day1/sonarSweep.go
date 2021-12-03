package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	depths, err := readData(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}
	fmt.Println(incCount(depths))
}

func readData(path string) (vals []float64, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var val float64
		_, err := fmt.Sscanf(s.Text(), "%f", &val)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		vals = append(vals, val)
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return vals, nil
}

func incCount(vals []float64) int {
	var count int
	for i := 1; i < len(vals); i++ {
		if vals[i] > vals[i-1] {
			count++
		}
	}
	return count
}
