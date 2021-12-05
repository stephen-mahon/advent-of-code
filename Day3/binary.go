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

func main() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	vals, err := readData(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	gamma, err := bits(vals)
	if err != nil {
		log.Fatalf("could not convert vals to binary: %v", err)
	}

	epilson, err := binaryInvert(gamma)
	if err != nil {
		log.Fatalf("could not convert vals to binary: %v", err)
	}

	fmt.Printf("Part 1\ngamma: %v\nepilson: %v\nPower consumption: %v\n", gamma, epilson, gamma*epilson)
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

func bits(vals []string) (int, error) {
	var sigB string
	for i := 0; i < len(vals[0]); i++ {
		bin := ""
		for _, val := range vals {
			bin += string(val[i])
		}
		if strings.Count(bin, "0") > strings.Count(bin, "1") {
			sigB += "0"
		} else {
			sigB += "1"
		}
	}

	output, err := strconv.ParseInt(sigB, 2, 64)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return int(output), nil
}

func binaryInvert(input int) (int, error) {
	val := strconv.FormatInt(int64(input), 2)
	var invBin string
	for _, v := range val {
		if string(v) == "0" {
			invBin += "1"
		} else {
			invBin += "0"
		}
	}
	output, err := strconv.ParseInt(invBin, 2, 64)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return int(output), nil
}
