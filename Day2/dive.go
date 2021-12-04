package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type cmds struct {
	v string  // vector (forward, up and down)
	s float64 // scalar
}

type pos struct {
	h float64 // horizontal
	d float64 // horizontal * aim
	a float64 // aim = depth
}

func main() {
	fileName := flag.String("f", "input.txt", "input file name")
	flag.Parse()

	vals, err := readData(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}
	subPos := calcPos1(vals)
	fmt.Printf("Part 1\nHorizontal: %.0f, Depth: %.0f\nans = %.0f\n\n", subPos.h, subPos.d, subPos.h*subPos.d)
	subPos = calcPos2(vals)
	fmt.Printf("Part 2\nHorizontal: %.0f, Depth: %.0f\nans = %.0f\n", subPos.h, subPos.d, subPos.h*subPos.d)
}

func readData(path string) (vals []cmds, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var vec string
		var sca float64
		_, err := fmt.Sscanf(s.Text(), "%s %f", &vec, &sca)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		vals = append(vals, cmds{vec, sca})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return vals, nil
}

func calcPos1(vals []cmds) pos {
	var fPos pos
	for i := range vals {
		if vals[i].v == "forward" {
			fPos.h += vals[i].s
		} else {
			if vals[i].v == "up" {
				fPos.d -= vals[i].s
			} else {
				fPos.d += vals[i].s
			}
		}
	}
	return fPos
}

func calcPos2(vals []cmds) pos {
	var fPos pos
	for i := range vals {
		if vals[i].v == "forward" {
			fPos.h += vals[i].s
			fPos.d += vals[i].s * fPos.a
		} else {
			if vals[i].v == "up" {
				fPos.a -= vals[i].s
			} else {
				fPos.a += vals[i].s
			}
		}
	}
	return fPos
}
