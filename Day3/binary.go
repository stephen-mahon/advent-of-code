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

	vals, err := readData(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	bits(vals)
}

func readData(path string) (vals []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var val string
		_, err := fmt.Sscanf(s.Text(), "%s", &val)
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

func bits(vals []string) {

}
