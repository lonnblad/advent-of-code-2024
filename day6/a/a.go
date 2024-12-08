package a

import (
	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	var result int

	matrix := util.StringMatrix(input)

	for idx, row := range matrix {
		for jdx, col := range row {
			if col == "^" {
				walk(matrix, idx, jdx, stepUp, "#")
			}
		}
	}

	result = len(uniquePositions)

	return result, nil
}

var uniquePositions = make(map[[2]int]bool)

func walk(matrix [][]string, startIdx, startJdx int, step stepFunc, until string) {
	idx, jdx := startIdx, startJdx

	for {
		uniquePositions[[2]int{idx, jdx}] = true

		newIdx, newJdx, turn := step(idx, jdx)

		if newIdx < 0 || newJdx < 0 || newIdx >= len(matrix) || newJdx >= len(matrix[0]) {
			return
		}

		if matrix[newIdx][newJdx] == until {
			walk(matrix, idx, jdx, turn, until)
			return
		}

		idx, jdx = newIdx, newJdx
	}
}

type stepFunc func(int, int) (int, int, stepFunc)

func stepUp(y, x int) (int, int, stepFunc) {
	return y - 1, x, stepRight
}

func stepRight(y, x int) (int, int, stepFunc) {
	return y, x + 1, stepDown
}

func stepDown(y, x int) (int, int, stepFunc) {
	return y + 1, x, stepLeft
}

func stepLeft(y, x int) (int, int, stepFunc) {
	return y, x - 1, stepUp
}
