package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	// path keeps track of our current directory in a heap
	var path []string
	// files maps the directory path to the total file size in that dir including directories in that path
	files := make(map[string]int)

	for _, v := range input {
		line := strings.Split(v, " ")

		if line[1] == "cd" {
			if line[2] == ".." { // if change directory, pop back one step
				_, path = pop(path)
			} else { // else add the new dir to the path
				path = append(path, line[2])
			}
		} else { // ls lists files and dir in current dir
			// "fileSize fileName"
			s, err := strconv.Atoi(line[0]) // string convert to get the file size
			if err == nil {                 // some of the lines won't contain a number so the err will catch that
				// add this files size to the map *tricky*
				for i := 1; i < len(path)+1; i++ {
					addr := strings.Join(path[:i], "/")
					// add file size to current dir size and the size of all parents
					files[addr] += s
				}
			}
		}
	}

	maxUsed := 40000000     // need 40 mb of storage
	totalUsed := files["/"] // maps are awesome
	neededSpace := totalUsed - maxUsed

	var part1 int
	part2 := math.MaxInt64 // creates a large int
	for i, v := range files {
		fmt.Println(i, v)
		if v <= 100000 {
			part1 += v
		}
		if v >= neededSpace {
			// What is the size of the smallest directory smaller than the need space?
			part2 = int(math.Min(float64(part2), float64(v)))
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func read(path string) ([]string, error) {
	var input []string
	dat, err := os.Open("data/" + path)
	if err != nil {
		return nil, err
	}
	defer dat.Close()

	s := bufio.NewScanner(dat)

	for s.Scan() {
		input = append(input, s.Text())
	}

	return input, nil
}

func pop(s []string) (string, []string) {
	return s[len(s)-1], s[:len(s)-1]
}
