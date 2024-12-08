package day4

import (
	"strings"
)

func A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	matrix := make([][]string, len(rows))

	for idx, row := range rows {
		matrix[idx] = strings.Split(row, "")
	}

	var result int

	for idx, row := range matrix {
		for jdx, val := range row {
			if val == "X" {
				result += findWord(matrix, idx, jdx, upRight, "XMAS")
				result += findWord(matrix, idx, jdx, right, "XMAS")
				result += findWord(matrix, idx, jdx, downRight, "XMAS")
				result += findWord(matrix, idx, jdx, down, "XMAS")
			}

			if val == "S" {
				result += findWord(matrix, idx, jdx, upRight, "SAMX")
				result += findWord(matrix, idx, jdx, right, "SAMX")
				result += findWord(matrix, idx, jdx, downRight, "SAMX")
				result += findWord(matrix, idx, jdx, down, "SAMX")
			}
		}
	}

	return result, nil
}

func findWord(matrix [][]string, idx, jdx int, step stepFunc, word string) int {
	if idx < 0 || jdx < 0 || idx >= len(matrix) || jdx >= len(matrix[idx]) {
		return 0
	}

	if matrix[idx][jdx] != string(word[0]) {
		return 0
	}

	newWord := word

	if matrix[idx][jdx] == string(word[0]) {
		newWord = word[1:]
	}

	if len(newWord) == 0 {
		return 1
	}

	newIdx, newJdx := step(idx, jdx)
	result := findWord(matrix, newIdx, newJdx, step, newWord)

	if result > 0 {
	}

	return result
}

type stepFunc func(idx, jdx int) (int, int)

func upRight(idx, jdx int) (int, int) {
	return idx - 1, jdx + 1
}

func right(idx, jdx int) (int, int) {
	return idx, jdx + 1
}

func downRight(idx, jdx int) (int, int) {
	return idx + 1, jdx + 1
}

func down(idx, jdx int) (int, int) {
	return idx + 1, jdx
}
