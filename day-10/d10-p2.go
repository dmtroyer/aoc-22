package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const CrtLineLength = 40
const CrtNumLines = 6

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

var cycle int = 1
var reg_x int = 1
var crt []bool = make([]bool, CrtNumLines*CrtLineLength)

func incrementCycle(i int) {
	for j := 0; j < i; j++ {
		current_column := (cycle - 1) % CrtLineLength
		if current_column >= reg_x-1 && current_column <= reg_x+1 {
			crt[cycle-1] = true
		}
		cycle++
	}
}

func initCrt() {
	for i := 0; i < len(crt); i++ {
		crt[i] = false
	}
}

func printCrt() {
	for i := 0; i < len(crt); i++ {
		if i != 0 && i%CrtLineLength == 0 {
			fmt.Print("\n")
		}
		if crt[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Print("\n")
}

func main() {
	initCrt()
	lines := parseLines()

	for _, line := range lines {
		if strings.Contains(line, "noop") {
			incrementCycle(1)
		} else {
			incrementCycle(2)
			var v int
			if _, err := fmt.Sscanf(line, "addx %d", &v); err != nil {
				log.Fatal(err)
			}
			reg_x += v
		}
	}

	printCrt()
}
