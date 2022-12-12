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

type monkey struct {
	num        int
	items      []int
	operation  string
	test       int
	condition1 int
	condition2 int
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var monkeys []monkey
	var items []int

	for i := range input {
		line := strings.Split(input[i], " ")

		if line[0] == "Monkey" {
			items = []int{}
			num, _ := strconv.Atoi(string(line[1][0]))
			si := strings.Split(input[i+1], ", ")
			for i := range si {
				if i == 0 {
					item, _ := strconv.Atoi(string(si[i][18:]))
					items = append(items, item)
				} else {
					item, _ := strconv.Atoi(string(si[i]))
					items = append(items, item)
				}
			}

			operation := input[i+2][19:]
			test, _ := strconv.Atoi(input[i+3][21:])
			lineC1 := strings.Split(input[i+4], " ")
			condition1, _ := strconv.Atoi(string(lineC1[len(lineC1)-1]))
			lineC2 := strings.Split(input[i+5], " ")
			condition2, _ := strconv.Atoi(string(lineC2[len(lineC2)-1]))
			monkeys = append(monkeys, monkey{num, items, operation, test, condition1, condition2})
		}
	}

	for i := range monkeys {
		fmt.Println(monkeys[i].num, monkeys[i].items)
	}
	fmt.Println()

	for _, v := range monkeys {
		for len(v.items) != 0 {
			var headItem int
			headItem, v.items = v.head()
			item := operation(v.operation, headItem) / 3
			fmt.Println(v.num, v.items)
			if item%v.test == 0 {
				fmt.Printf("    %v (%v) -> monkey #%v\n", headItem, item, monkeys[v.condition1].num)
				monkeys[v.condition1].addItem(item)
			} else {
				fmt.Printf("    %v (%v) -> monkey #%v\n", headItem, item, monkeys[v.condition2].num)
				monkeys[v.condition2].addItem(item)
			}
		}
		fmt.Println()
	}

	for i := range monkeys {
		fmt.Println(monkeys[i].num, monkeys[i].items)
	}
	fmt.Println()
}

func (m *monkey) head() (int, []int) {
	return m.items[0], m.items[1:]
}

func (m *monkey) addItem(item int) []int {
	m.items = append(m.items, item)
	return m.items
}

func operation(line string, n int) int {
	lineSplit := strings.Split(line, " ")
	op := lineSplit[1]
	switch op {
	case "+":
		val, err := strconv.Atoi(lineSplit[2])
		if err == nil {
			return n + val
		}
		return n + n
	case "-":
		val, err := strconv.Atoi(lineSplit[2])
		if err == nil {
			return n - val
		}
		return n - n
	case "*":
		val, err := strconv.Atoi(lineSplit[2])
		if err == nil {
			return n * val
		}
		return n * n
	}
	return n
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
