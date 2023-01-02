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

	head := [2]int{0, 0}
	tail := head
	//start := head
	moves := 0

	checkPos := make(map[[2]int]int)

	for i := range input {
		instr := strings.Split(input[i], " ")
		dir := instr[0]
		mag, _ := strconv.Atoi(instr[1])
		for mag != 0 {
			head = move(head, dir)
			tail, moves = follow(head, tail, moves)
			checkPos[tail]++
			mag--
		}
	}

	fmt.Println("Part 1:", len(checkPos))

}

func distance(v1, v2 [2]int) float64 {
	return math.Sqrt(math.Pow(float64(v2[0])-float64(v1[0]), 2) + math.Pow(float64(v2[1])-float64(v1[1]), 2))

}

func follow(pos, f [2]int, count int) ([2]int, int) {
	if magnitude(pos, f) <= 1 {
		return f, count
	}

	d := distance(pos, f)
	var next [2]int

	if distance(pos, [2]int{f[0] + 1, f[1]}) < d {
		next = [2]int{f[0] + 1, f[1]}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0] - 1, f[1]}) < d {
		next = [2]int{f[0] - 1, f[1]}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0], f[1] + 1}) < d {
		next = [2]int{f[0], f[1] + 1}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0], f[1] - 1}) < d {
		next = [2]int{f[0], f[1] - 1}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0] + 1, f[1] + 1}) < d {
		next = [2]int{f[0] + 1, f[1] + 1}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0] + 1, f[1] - 1}) < d {
		next = [2]int{f[0] + 1, f[1] - 1}
		d = distance(pos, next)
	}
	if distance(pos, [2]int{f[0] - 1, f[1] + 1}) < d {
		next = [2]int{f[0] - 1, f[1] + 1}
		d = distance(pos, next)
	}

	if distance(pos, [2]int{f[0] - 1, f[1] - 1}) < d {
		next = [2]int{f[0] - 1, f[1] - 1}
	}

	return next, count + 1
}

func move(pos [2]int, dir string) [2]int {
	direction := map[string][]int{"U": {0, 1}, "R": {1, 0}, "D": {0, -1}, "L": {-1, 0}}
	next, check := direction[dir]
	if !check {
		return pos
	}
	return [2]int{pos[0] + next[0], pos[1] + next[1]}
}

func magnitude(pos1, pos2 [2]int) int {
	x1 := pos1[0]
	y1 := pos1[1]
	x2 := pos2[0]
	y2 := pos2[1]
	dx := math.Pow(float64(x2)-float64(x1), 2)
	dy := math.Pow(float64(y2)-float64(y1), 2)

	return int(math.Sqrt(dx + dy))
}

func read(path string) ([]string, error) {
	var input []string
	dat, err := os.Open("../data/" + path)
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
