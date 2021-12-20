package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// add bits func back into oxygenscrubber for new version of gamma based on remaining vals.
// rework gamma func to return consistent types for new functions. Maybe just a string of binary numbers and do the conversion in main()

func main() {
	fileName := "input.txt"

	vals, err := readData(fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", fileName, err)
	}
	bits := make([][]string, len(vals[0]))
	for i := range vals {
		for j := range vals[i] {
			bits[j] = append(bits[j], string(vals[i][j]))
		}
	}

	gamma, _ := readGamma(bits)
	epsilon, _ := readEpsilon(bits)
	fmt.Println(gamma * epsilon)

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

func cannonize(vals []string) map[string]int {
	dict := make(map[string]int)
	for _, v := range vals {
		dict[v] = dict[v] + 1
	}
	return dict
}

func readGamma(bits [][]string) (int64, error) {
	var bit string
	for i := range bits {
		dict := cannonize(bits[i])
		if dict["0"] > dict["1"] {
			bit += "0"
		} else {
			bit += "1"
		}
	}
	return strconv.ParseInt(bit, 2, 64)
}

func readEpsilon(bits [][]string) (int64, error) {
	var bit string
	for i := range bits {
		dict := cannonize(bits[i])
		if dict["1"] > dict["0"] {
			bit += "0"
		} else {
			bit += "1"
		}
	}
	return strconv.ParseInt(bit, 2, 64)
}
