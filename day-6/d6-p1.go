package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func containsDups(s string) bool {
	for _, c := range s {
		if strings.Count(s, string(c)) > 1 {
			return true
		}
	}
	return false
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

func main() {
	lines := parseLines()

	for _, line := range lines {
		for i := 4; i <= len(line); i++ {
			if !containsDups(line[i-4 : i]) {
				fmt.Println(i)
				break
			}
		}
	}

}
