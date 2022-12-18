package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const TREE_HEIGHT_MASK = uint8(0xF)
const VISIBLE_FROM_WEST_FLAG = uint8(1 << 7)
const VISIBLE_FROM_NORTH_FLAG = uint8(1 << 6)
const VISIBLE_FROM_EAST_FLAG = uint8(1 << 5)
const VISIBLE_FROM_SOUTH_FLAG = uint8(1 << 4)

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

func calcRowVisibility(matrix [][]uint8) {
	for i := 0; i < len(matrix); i++ {
		var high uint8
		var j int

		for j = 0; j < len(matrix[i]); j++ {
			if j == 0 || (matrix[i][j]&TREE_HEIGHT_MASK) > high {
				high = matrix[i][j] & TREE_HEIGHT_MASK
				matrix[i][j] |= VISIBLE_FROM_WEST_FLAG
			}
		}

		for j--; j >= 0; j-- {
			if j == len(matrix[i])-1 || (matrix[i][j]&TREE_HEIGHT_MASK) > high {
				high = matrix[i][j] & TREE_HEIGHT_MASK
				matrix[i][j] |= VISIBLE_FROM_EAST_FLAG
			}
		}
	}
}

func calcColVisibility(matrix [][]uint8) {
	for i := 0; i < len(matrix[0]); i++ {
		var high uint8
		var j int

		for j = 0; j < len(matrix); j++ {
			if j == 0 || (matrix[j][i]&TREE_HEIGHT_MASK) > high {
				high = matrix[j][i] & TREE_HEIGHT_MASK
				matrix[j][i] |= VISIBLE_FROM_NORTH_FLAG
			}
		}

		for j--; j >= 0; j-- {
			if j == len(matrix[i])-1 || (matrix[j][i]&TREE_HEIGHT_MASK) > high {
				high = matrix[j][i] & TREE_HEIGHT_MASK
				matrix[j][i] |= VISIBLE_FROM_SOUTH_FLAG
			}
		}
	}
}

func numVisibleTrees(matrix [][]uint8) int {
	n := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > TREE_HEIGHT_MASK {
				n++
			}
		}
	}
	return n
}

func main() {
	lines := parseLines()
	matrix := parseLinesToMatrix(lines)
	calcRowVisibility(matrix)
	calcColVisibility(matrix)
	num_vis_trees := numVisibleTrees(matrix)
	fmt.Println(num_vis_trees)
}
