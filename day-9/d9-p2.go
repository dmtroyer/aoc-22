package main

import (
	"bufio"
	"errors"
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

func withinOneSpace(p1 Point, p2 Point) bool {
	spaces := append(spacesOneFromPoint(p2), p2)
	return containsPoint(spaces, p1)
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

func containsPoint(points []Point, point Point) bool {
	for _, p := range points {
		if p == point {
			return true
		}
	}
	return false
}

func spacesDiagonalFromPoint(p Point) []Point {
	return []Point{Point{x: p.x - 1, y: p.y - 1},
		Point{x: p.x - 1, y: p.y + 1},
		Point{x: p.x + 1, y: p.y - 1},
		Point{x: p.x + 1, y: p.y + 1},
	}
}

func spacesStraightFromPoint(p Point) []Point {
	return []Point{Point{x: p.x - 1, y: p.y},
		Point{x: p.x + 1, y: p.y},
		Point{x: p.x, y: p.y - 1},
		Point{x: p.x, y: p.y + 1},
	}
}

func spacesOneFromPoint(p Point) []Point {
	return append(spacesDiagonalFromPoint(p), spacesStraightFromPoint(p)...)
}

func commonNeighbor(lead Point, follower Point) (Point, error) {
	follower_options := spacesOneFromPoint(follower)
	for _, p := range follower_options {
		if containsPoint(spacesStraightFromPoint(lead), p) {
			return p, nil
		}
	}
	for _, p := range follower_options {
		if containsPoint(spacesDiagonalFromPoint(lead), p) {
			return p, nil
		}
	}
	// should never get here
	return follower, errors.New("knots are more than two spaces away from each other")
}

func main() {
	lines := parseLines()
	motions := parseMotions(lines)

	starting_point := Point{x: 11, y: 5}
	knots := make([]Point, 10)
	for i := range knots {
		knots[i] = starting_point
	}

	head := &knots[0]
	tail := &knots[9]

	visited := map[Point]bool{*tail: true}

	for _, motion := range motions {
		for i := 0; i < motion.distance; i++ {
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

			for j := 1; j < len(knots); j++ {
				if withinOneSpace(knots[j], knots[j-1]) {
					break
				} else {
					// move to the space one away from both that is in the same row or column as knots[j-1]
					var err error
					knots[j], err = commonNeighbor(knots[j-1], knots[j])
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			visited[*tail] = true
		}
	}

	fmt.Printf("Spaces visited: %d\n", len(visited))
}
