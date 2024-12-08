package b

import (
	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	var result int

	matrix := util.StringMatrix(input)

	startPos := findStartPos(matrix)
	walk(matrix, startPos[0], startPos[1], stepUp, "#")

	for pos := range uniquePositions {
		idx, jdx := pos[0], pos[1]

		if idx == startPos[0] && jdx == startPos[1] || matrix[idx][jdx] == "#" {
			continue
		}

		matrix[idx][jdx] = "#"

		walkedPositions := make(map[[2]int]map[int]bool)

		if isInALoop(matrix, startPos[0], startPos[1], stepUp, walkedPositions) {
			result++
		}

		matrix[idx][jdx] = "."
	}

	return result, nil
}

func findStartPos(matrix [][]string) [2]int {
	for idx, row := range matrix {
		for jdx, col := range row {
			if col == "^" {
				return [2]int{idx, jdx}
			}
		}
	}

	return [2]int{}
}

var uniquePositions = make(map[[2]int]bool)

func walk(matrix [][]string, startIdx, startJdx int, step stepFunc, until string) {
	idx, jdx := startIdx, startJdx

	for {
		uniquePositions[[2]int{idx, jdx}] = true

		newIdx, newJdx, _, turn := step(idx, jdx)

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

func isInALoop(matrix [][]string, idx, jdx int, step stepFunc, walkedPositions map[[2]int]map[int]bool) bool {
	maxY, maxX := len(matrix), len(matrix[0])

	for {
		newIdx, newJdx, dir, turn := step(idx, jdx)

		if walkedPositions[[2]int{newIdx, newJdx}][dir] {
			return true
		}

		if _, ok := walkedPositions[[2]int{idx, jdx}]; !ok {
			walkedPositions[[2]int{idx, jdx}] = make(map[int]bool)
		}

		walkedPositions[[2]int{idx, jdx}][dir] = true

		if newIdx < 0 || newJdx < 0 || newIdx >= maxY || newJdx >= maxX {
			return false
		}

		if matrix[newIdx][newJdx] == "#" {
			return isInALoop(matrix, idx, jdx, turn, walkedPositions)
		}

		idx, jdx = newIdx, newJdx
	}
}

type stepFunc func(int, int) (int, int, int, stepFunc)

func stepUp(y, x int) (int, int, int, stepFunc) {
	return y - 1, x, 0, stepRight
}

func stepRight(y, x int) (int, int, int, stepFunc) {
	return y, x + 1, 1, stepDown
}

func stepDown(y, x int) (int, int, int, stepFunc) {
	return y + 1, x, 2, stepLeft
}

func stepLeft(y, x int) (int, int, int, stepFunc) {
	return y, x - 1, 3, stepUp
}
