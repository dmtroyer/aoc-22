package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

type Point struct {
	x int
	y int
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

func parseMatrix(lines []string) ([][]int, Point, Point) {
	matrix := [][]int{}
	var start, antenna Point

	for i, line := range lines {
		row := make([]int, len(line))
		for j, c := range []rune(line) {
			if c == 'S' {
				start = Point{x: i, y: j}
				row[j] = 0
			} else if c == 'E' {
				antenna = Point{x: i, y: j}
				row[j] = 25
			} else {
				row[j] = strings.Index(Alphabet, string(c))
			}
			matrix = append(matrix, row)
		}
	}
	return matrix, start, antenna
}

func main() {
	lines := parseLines()
	matrix, start, antenna := parseMatrix(lines)
	fmt.Println(matrix)
	fmt.Println(start)
	fmt.Println(antenna)
}
