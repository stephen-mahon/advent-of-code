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

type valve struct {
	name     string
	flowrate int
	tunnels  []string
	status   bool
	visited  bool
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var valves []valve
	paths := make(map[string]valve)

	for i := range input {
		v := parseInfo(input[i])
		valves = append(valves, v)
		paths[v.name] = v
	}

	var stack []valve

	var current = valves[0]

	current.visit(true)
	next := current.next(paths)
	fmt.Printf("valve = %v, next = %v\n", current.name, next)
	if next.visited {
		next.visit(true)
		stack = append(stack, current)
		fmt.Printf("valve = %v, next = %v\n", current.name, next)
		current = next
	} else if len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

}

func (v *valve) visit(b bool) {
	v.visited = b
}

func (v *valve) next(paths map[string]valve) valve {
	for i := range v.tunnels {
		next := paths[v.tunnels[i]]
		if next.flowrate != 0 {
			return paths[v.tunnels[i]]
		}
	}
	return valve{}
}

func parseInfo(input string) valve {
	str := strings.Split(input, " ")
	name := str[1]
	flowrate, _ := strconv.Atoi(str[4][5 : len(str[4])-1])
	tunnels := str[9:]

	for i := range tunnels {
		tunnels[i] = tunnels[i][:2]
	}

	return valve{name, flowrate, tunnels, false, false}
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
