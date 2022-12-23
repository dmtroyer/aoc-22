package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
var signal_strength_sum int = 0

func incrementCycle(i int) {
	for j := 0; j < i; j++ {
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signal_strength_sum += cycle * reg_x
		}
		cycle++
	}
	return
}

func main() {
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

	fmt.Printf("Signal strength sum: %d\n", signal_strength_sum)
}
