package main

import (
	"bufio"
	"fmt"
	"log"
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
	distance  uint8
}

type Point struct {
	x float32
	y float32
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

	head := Point{x: 0.0, y: 0.0}
	// tail := Knot{x: 0, y: 0}
	visited := make(map[Point]bool)

	for _, motion := range motions {
		var i uint8
		//  I think, if the head and tail are one space apart (diagonally or horizontally) after the move, the head's space before the move will not be counted.
    // However, if the head and tail are more than one space apart after the move, the tail will follow to where the head prevously was.
		for i = 0; i < motion.distance; i++ {
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
		}
		visited[head] = true
	}

	fmt.Println(head)
	fmt.Println(visited)
}
