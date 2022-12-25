package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const NumRounds = 20

type Test struct {
	divisible_by int
	true_monkey  int
	false_monkey int
}

type Monkey struct {
	items       []int
	operations  []string
	test        Test
	inspections int
}

func parseLines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return lines
}

func processOperation(item int, operations []string) int {
	var y int
	var z int

	if len(operations) != 3 {
		log.Fatal("Expecting three operations.")
	}

	if operations[0] == "old" {
		y = item
	} else {
		y, _ = strconv.Atoi(operations[0])
	}

	if operations[2] == "old" {
		z = item
	} else {
		z, _ = strconv.Atoi(operations[2])
	}

	if operations[1] == "*" {
		return y * z
	} else {
		return y + z
	}
}

func parseMonkeys(lines []string) []Monkey {
	monkeys := []Monkey{}
	var m Monkey

	for _, line := range lines {
		if match, _ := regexp.MatchString(`^Monkey \d:$`, line); match {
			m = Monkey{inspections: 0, items: []int{}, test: Test{}}

		} else if match, _ := regexp.MatchString(`Starting items: `, line); match {
			item_string := strings.TrimPrefix(line, "  Starting items: ")
			items := strings.Split(item_string, ", ")
			for _, item := range items {
				if item_int, err := strconv.Atoi(item); err == nil {
					m.items = append(m.items, item_int)
				} else {
					log.Fatal(err)
				}
			}

		} else if match, _ := regexp.MatchString(`Operation: `, line); match {
			op_str := strings.TrimPrefix(line, "  Operation: new = ")
			m.operations = strings.Split(op_str, " ")

		} else if match, _ := regexp.MatchString(`Test: `, line); match {
			if _, err := fmt.Sscanf(line, "  Test: divisible by %d", &m.test.divisible_by); err != nil {
				log.Fatal(err)
			}

		} else if match, _ := regexp.MatchString(`If true: `, line); match {
			if _, err := fmt.Sscanf(line, "    If true: throw to monkey %d", &m.test.true_monkey); err != nil {
				log.Fatal(err)
			}

		} else if match, _ := regexp.MatchString(`If false: `, line); match {
			if _, err := fmt.Sscanf(line, "    If false: throw to monkey %d", &m.test.false_monkey); err != nil {
				log.Fatal(err)
			}

		} else {
			monkeys = append(monkeys, m)
		}
	}
	monkeys = append(monkeys, m)

	return monkeys
}

func main() {
	lines := parseLines()
	monkeys := parseMonkeys(lines)

	for r := 0; r < NumRounds; r++ {
		for i := 0; i < len(monkeys); i++ {
			m := &monkeys[i]

			for j := 0; j < len(m.items); j++ {
				item := &m.items[j]
				*item = processOperation(*item, m.operations) / 3

				var next_monkey *Monkey
				if *item%m.test.divisible_by == 0 {
					next_monkey = &monkeys[m.test.true_monkey]
				} else {
					next_monkey = &monkeys[m.test.false_monkey]
				}
				next_monkey.items = append(next_monkey.items, *item)

				m.inspections++
			}

			m.items = []int{}
		}
	}

	// Print the monkeys' final items
	for i, m := range monkeys {
		fmt.Printf("Monkey %d: ", i)
		for j, item := range m.items {
			fmt.Printf("%d", item)
			if j != len(m.items)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// Tabulate final number of inspections
	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, m.inspections)
		inspections[i] = m.inspections
	}
	fmt.Println()

	// Calculate total monkey business
	sort.Ints(inspections)
	total_monkey_business := inspections[len(monkeys)-1] * inspections[len(monkeys)-2]
	fmt.Printf("Total monkey business: %d\n", total_monkey_business)
}
