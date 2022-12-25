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
	x         int
	y         int
	parent    *Point
	elevation int
	explored  bool
}

var matrix = [][]*Point{}
var antenna *Point
var m_height, m_width = 0, 0

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

func parseMatrix(lines []string) {
	for i, line := range lines {
		row := make([]*Point, len(line))

		for j, c := range line {
			if c == 'S' {
				row[j] = &Point{x: j, y: i, elevation: 0, explored: false}
			} else if c == 'E' {
				antenna = &Point{x: j, y: i, elevation: 25, explored: false}
				row[j] = antenna
			} else {
				e := strings.Index(Alphabet, string(c))
				row[j] = &Point{x: j, y: i, elevation: e, explored: false}
			}
		}
		matrix = append(matrix, row)
	}

	m_height = len(lines)
	m_width = len(matrix[0])
}

func onGrid(x int, y int) bool {
	return x >= 0 && y >= 0 && x < m_width && y < m_height
}

func pathLength(p *Point) int {
	var i int
	for i = 0; p.parent != nil; i++ {
		p = p.parent
	}
	return i
}

func printMatrix() {
	for _, row := range matrix {
		for _, p := range row {
			if p.elevation < 10 {
				fmt.Print(" ")
			}
			fmt.Print(p.elevation, " ")
		}
		fmt.Print("\n")
	}
}

func neighborIsReachable(current *Point, destination *Point) bool {
	return destination.elevation-current.elevation <= 1
}

func reachableNeighbors(p *Point) []*Point {
	reachable_neighbors := []*Point{}
	relative_neighbors := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, n := range relative_neighbors {
		if onGrid(p.x+n[0], p.y+n[1]) {
			neighbor := matrix[p.y+n[1]][p.x+n[0]]
			if neighborIsReachable(p, neighbor) {
				reachable_neighbors = append(reachable_neighbors, neighbor)
			}
		}
	}

	return reachable_neighbors
}

func resetMatrix() {
	for _, row := range matrix {
		for _, p := range row {
			p.parent = nil
			p.explored = false
		}
	}
}

func search(start *Point) *Point {
	start.explored = true
	queue := []*Point{start}

	for len(queue) > 0 {
		p := queue[0]
		if p == antenna {
			return p
		}
		queue = queue[1:]

		for _, neighbor := range reachableNeighbors(p) {
			if !neighbor.explored {
				neighbor.explored = true
				neighbor.parent = p
				queue = append(queue, neighbor)
			}
		}
	}

	return nil
}

func main() {
	lines := parseLines()
	parseMatrix(lines)

	path_lengths := []int{}
	for _, row := range matrix {
		for _, p := range row {
			if p.elevation == 0 {
				if end := search(p); end != nil {
					path_lengths = append(path_lengths, pathLength(end))
				}
				resetMatrix()
			}
		}
	}
	sort.Ints(path_lengths)

	fmt.Printf("Shortest Path length: %d\n", path_lengths[0])
}
