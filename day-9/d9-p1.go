package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Direction byte

const (
	Up    Direction = 'U'
	Down  Direction = 'D'
	Left  Direction = 'L'
	Right Direction = 'R'
)

type RopeMotion struct {
	direction Direction
	distance  int
}

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

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt((float64(p2.x-p1.x) * float64(p2.x-p1.x)) + (float64(p2.y-p1.y) * float64(p2.y-p1.y)))
}

func withinOneSpace(p1 Point, p2 Point) bool {
	return (math.Abs(float64(p2.x-p1.x)) <= 1 && math.Abs(float64(p2.y-p1.y)) <= 1)
}

func parseMotions(lines []string) []RopeMotion {
	motions := make([]RopeMotion, len(lines))

	for i, line := range lines {
		var motion RopeMotion
		_, err := fmt.Sscanf(line, "%c %d", &motion.direction, &motion.distance)
		if err == nil {
			motions[i] = motion
		}
	}

	return motions
}

func main() {
	lines := parseLines()
	motions := parseMotions(lines)

	head := Point{x: 0, y: 0}
	tail := Point{x: 0, y: 0}
	visited := map[Point]bool{tail: true}

	for _, motion := range motions {
		//  I think, if the head and tail are one space apart (diagonally or horizontally) after the move, the head's space before the move will not be counted.
		// However, if the head and tail are more than one space apart after the move, the tail will follow to where the head prevously was.
		for i := 0; i < motion.distance; i++ {
			prev := head
			switch motion.direction {
			case Up:
				head.y++
			case Down:
				head.y--
			case Left:
				head.x--
			case Right:
				head.x++
			}
			if !withinOneSpace(tail, head) {
				tail = prev
				visited[tail] = true
			}
		}
	}

	fmt.Println(len(visited))
}
