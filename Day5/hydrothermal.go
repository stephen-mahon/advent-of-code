package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var title = "Day 5: Hydrthermal Venture"

type coor struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	fmt.Println(title)
	data, _ := readFile("test.txt")

	for _, v := range data {
		fmt.Printf("%v,%v -> %v,%v ", v.x1, v.y1, v.x2, v.y2)
		if v.x1 == v.x2 {
			fmt.Printf("move y %v", v.y2-v.y1)
		} else if v.y1 == v.y2 {
			fmt.Printf("move x %v", v.x2-v.x1)
		}
		fmt.Println()
	}

}

func readFile(path string) (data []coor, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(s.Text(), "%v,%v -> %v,%v", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}

		data = append(data, coor{x1, y1, x2, y2})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return data, nil
}
