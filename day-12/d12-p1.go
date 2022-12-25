package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const Alphabet = "abcdefghijklmnopqrstuvwxyz"

type Point struct {
	x int
	y int
	prev *Point
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

var matrix = [][]int{}
var start, antenna = Point{}, Point{}
var pathways = []int{}
var m_height, m_width = 0, 0

func havePreviouslyVisited(node Point, neighbor Point) bool {
	for node.prev != nil {
		if pointsEqual(*node.prev, neighbor) {
			return true
		}
		node = *node.prev
	}
	return false
}

func parseMatrix(lines []string) {
	for i, line := range lines {
		row := make([]int, len(line))
		for j, c := range line {
			if c == 'S' {
				start = Point{x: j, y: i}
				row[j] = 0
			} else if c == 'E' {
				antenna = Point{x: j, y: i}
				row[j] = 25
			} else {
				row[j] = strings.Index(Alphabet, string(c))
			}
		}
		matrix = append(matrix, row)
	}
	m_height = len(lines)
	m_width  = len(matrix[0])
}

func onGrid(p Point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < m_width && p.y < m_height
}

func pathLength(p Point) int {
	i := 0
	for p.prev != nil {
		i++
		p = *p.prev
	}
	return i
}

func pointElevation(p Point) int {
	return matrix[p.y][p.x]
}

func pointsEqual(p1 Point, p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func withinOneElevation(current Point, destination Point) bool {
	diff := pointElevation(destination) - pointElevation(current)
	return diff <= 1
}

func stepOptions(p Point) []Point {
	options := []Point{}
	neighbors := []Point{
		Point{x: p.x, y: p.y+1},
		Point{x: p.x, y: p.y-1},
		Point{x: p.x+1, y: p.y},
		Point{x: p.x-1, y: p.y},
	}

	for _, n := range neighbors {
		if onGrid(n) && withinOneElevation(p, n) && !havePreviouslyVisited(p, n) {
			options = append(options, n)
		}
	}

	return options
}

func findPath(p Point) {
	if pointsEqual(p, antenna) {
		pathways = append(pathways, pathLength(p))
	} else {
		for _, n := range stepOptions(p) {
			n.prev = &p
			findPath(n)
		}
	}
}

func main() {
	fmt.Println("Parsing Lines...")
	lines := parseLines()

	fmt.Println("Parsing Matrix...")
	parseMatrix(lines)

	fmt.Println("Starting to explore...")
	findPath(start)

	sort.Ints(pathways)
	fmt.Println(pathways[0])
}
