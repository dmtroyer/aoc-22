package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func parseLinesToMatrix(lines []string) [][]uint8 {
	matrix := make([][]uint8, len(lines))

	for i, line := range lines {
		matrix[i] = make([]uint8, len(line))

		for j, c := range line {
			if val, err := strconv.Atoi(string(c)); err == nil {
				matrix[i][j] = uint8(val)
			} else {
				panic(err)
			}
		}
	}
	return matrix
}

func scenicScoreToNorth(matrix *[][]uint8, i int, j int) int {
	score := 0
	height := (*matrix)[i][j]

	for i--; i >= 0; i-- {
		score++
		if height <= (*matrix)[i][j] {
			break
		}
	}
	return score
}

func scenicScoreToEast(matrix *[][]uint8, i int, j int) int {
	score := 0
	height := (*matrix)[i][j]

	for j++; j < len((*matrix)[i]); j++ {
		score++
		if height <= (*matrix)[i][j] {
			break
		}
	}
	return score
}

func scenicScoreToSouth(matrix *[][]uint8, i int, j int) int {
	score := 0
	height := (*matrix)[i][j]

	for i++; i < len((*matrix)); i++ {
		score++
		if height <= (*matrix)[i][j] {
			break
		}
	}
	return score
}

func scenicScoreToWest(matrix *[][]uint8, i int, j int) int {
	score := 0
	height := (*matrix)[i][j]

	for j--; j >= 0; j-- {
		score++
		if height <= (*matrix)[i][j] {
			break
		}
	}
	return score
}

func highestScenicScore(matrix [][]uint8) int {
	high_score := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			score := scenicScoreToNorth(&matrix, i, j)
			if score == 0 {
				break
			}
			score *= scenicScoreToEast(&matrix, i, j)
			if score == 0 {
				break
			}
			score *= scenicScoreToSouth(&matrix, i, j)
			if score == 0 {
				break
			}
			score *= scenicScoreToWest(&matrix, i, j)
			if score > high_score {
				high_score = score
			}
		}
	}

	return high_score
}

func main() {
	lines := parseLines()
	matrix := parseLinesToMatrix(lines)
	fmt.Printf("The highest scenic score is %d\n", highestScenicScore(matrix))
}
