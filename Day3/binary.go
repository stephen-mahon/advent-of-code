package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "test.txt"

	vals, err := readData(fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", fileName, err)
	}
	bits := bitSeg(vals)
	gamma, _ := readGamma(bits)
	epsilon, _ := readEpsilon(bits)
	fmt.Println("Part 1:", gamma*epsilon)
	bitCriteria(bits)

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

func bitSeg(vals []string) [][]string {
	bits := make([][]string, len(vals[0]))
	for i := range vals {
		for j := range vals[i] {
			bits[j] = append(bits[j], string(vals[i][j]))
		}
	}
	return bits
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

func bitCriteria(bits [][]string) {
	for i := range bits {
		bit := ""
		for j := range bits[i] {
			bit += bits[i][j]
		}
		fmt.Println(bit, "0s:", strings.Count(bit, "0"), "1s:", strings.Count(bit, "1"))
	}
}
