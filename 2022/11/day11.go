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
	operation  string
	test       int
	condition1 int
	condition2 int
}

func main() {
	fileName := flag.String("f", "test.txt", "input file name")
	worryLevel := flag.Int("w", 3, "worry level")
	rounds := flag.Int("r", 20, "number of rounds")

	flag.Parse()

	input, err := read(*fileName)
	if err != nil {
		log.Fatalf("could not read %s: %v", *fileName, err)
	}

	var monkeys []monkey
	var items [][]int

	for i := range input {
		line := strings.Split(input[i], " ")

		if line[0] == "Monkey" {
			itemList := []int{}
			num, _ := strconv.Atoi(string(line[1][0]))
			si := strings.Split(input[i+1], ", ")
			for i := range si {
				if i == 0 {
					item, _ := strconv.Atoi(string(si[i][18:]))
					itemList = append(itemList, item)
				} else {
					item, _ := strconv.Atoi(string(si[i]))
					itemList = append(itemList, item)
				}
			}

			items = append(items, itemList)

			operation := input[i+2][19:]
			test, _ := strconv.Atoi(input[i+3][21:])
			lineC1 := strings.Split(input[i+4], " ")
			condition1, _ := strconv.Atoi(string(lineC1[len(lineC1)-1]))
			lineC2 := strings.Split(input[i+5], " ")
			condition2, _ := strconv.Atoi(string(lineC2[len(lineC2)-1]))
			monkeys = append(monkeys, monkey{num, operation, test, condition1, condition2})
		}
	}

	inspections := make([]int, len(monkeys))
	for i := 0; i < *rounds; i++ {
		inspections = monkeyBusiness(monkeys, inspections, items, *worryLevel)
	}

	fmt.Printf("== After round %v ==\n", *rounds)
	for i := range inspections {
		fmt.Printf("Monkey %v inspected items %v times\n", i, inspections[i])
	}

	bubbleSort(inspections)
	fmt.Println(inspections[0] * inspections[1])

}

func monkeyBusiness(monkeys []monkey, inspections []int, items [][]int, worryLevel int) []int {
	for i, v := range monkeys {
		for len(items[i]) != 0 {
			head := items[i][0]
			items[i] = items[i][1:]
			item := operation(v.operation, head) / worryLevel
			inspections[i]++

			if item%v.test == 0 {
				items[v.condition1] = append(items[v.condition1], item)
			} else {
				items[v.condition2] = append(items[v.condition2], item)
			}
		}
	}

	return inspections
}

func swap(xp, yp *int) {
	temp := *xp
	*xp = *yp
	*yp = temp
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
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
		return 0
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
