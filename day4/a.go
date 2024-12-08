package day4

import (
	"fmt"
	"strings"
)

func A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	// re := regexp.MustCompile(`X.*M.*A.*S`)

	matrix := make([][]string, len(rows))

	for idx, row := range rows {
		matrix[idx] = strings.Split(row, "")
	}

	var result int

	for idx, row := range matrix {
		for jdx, val := range row {
			if val == "X" {
				// result += findWord(matrix, idx, jdx, up, "XMAS")
				result += findWord(matrix, idx, jdx, upRight, "XMAS")
				result += findWord(matrix, idx, jdx, right, "XMAS")
				result += findWord(matrix, idx, jdx, downRight, "XMAS")
				result += findWord(matrix, idx, jdx, down, "XMAS")
			}

			if val == "S" {
				// result += findWord(matrix, idx, jdx, up, "SAMX")
				result += findWord(matrix, idx, jdx, upRight, "SAMX")
				result += findWord(matrix, idx, jdx, right, "SAMX")
				result += findWord(matrix, idx, jdx, downRight, "SAMX")
				result += findWord(matrix, idx, jdx, down, "SAMX")
			}
		}
	}

	// out := rows
	// fmt.Println(out)
	// for _, row := range out {
	// 	matches := re.FindAllString(row, -1)
	// 	fmt.Println(matches)
	// 	result += len(matches)
	// }
	// fmt.Println("----")

	// out = util.Rotate45Degrees(out)
	// fmt.Println(out)
	// for _, row := range out {
	// 	matches := re.FindAllString(row, -1)
	// 	fmt.Println(matches)
	// 	result += len(matches)
	// }
	// fmt.Println("----")

	// out = util.Rotate45Degrees(out)
	// fmt.Println(out)
	// for _, row := range out {
	// 	matches := re.FindAllString(row, -1)
	// 	fmt.Println(matches)
	// 	result += len(matches)
	// }
	// fmt.Println("----")

	// out = util.Rotate45Degrees(out)
	// fmt.Println(out)
	// for _, row := range out {
	// 	matches := re.FindAllString(row, -1)
	// 	fmt.Println(matches)
	// 	result += len(matches)
	// }

	return result, nil
}

// var checked = map[[2]int]map[string]bool{}

func findWord(matrix [][]string, idx, jdx int, step stepFunc, word string) int {
	if idx < 0 || jdx < 0 || idx >= len(matrix) || jdx >= len(matrix[idx]) {
		return 0
	}

	fmt.Println(idx, jdx, matrix[idx][jdx], word)

	if matrix[idx][jdx] != string(word[0]) {
		return 0
	}

	newWord := word

	if matrix[idx][jdx] == string(word[0]) {
		newWord = word[1:]
	}

	if len(newWord) == 0 {
		fmt.Println("found", idx, jdx, word)
		return 1
	}

	newIdx, newJdx := step(idx, jdx)
	result := findWord(matrix, newIdx, newJdx, step, newWord)

	if result > 0 {
		fmt.Println("found", idx, jdx, word)
	}

	return result

	// var result int

	// for _, pos := range potentialPositions {
	// 	if val, exists := checked[[2]int{pos[0], pos[1]}]; exists {
	// 		if _, exists := val[word]; exists {
	// 			continue
	// 		}
	// 	}

	// 	result += findWords(matrix, pos[0], pos[1], word)
	// }

	// return result
	// }

	// if len(word) == 1 {
	// 	return 1
	// }

	// return findWords(matrix, idx+1, jdx, word[1:]) + findWords(matrix, idx, jdx+1, word[1:]) +
	// 	findWords(matrix, idx-1, jdx, word[1:]) + findWords(matrix, idx, jdx-1, word[1:])

	// result := 0
	// matrix[idx][jdx] = "."

	// for i := -1; i <= 1; i++ {
	// 	for j := -1; j <= 1; j++ {
	// 		if i == 0 && j == 0 {
	// 			continue
	// 		}

	// 		result += findWords(matrix, idx+i, jdx+j, word[1:])
	// 	}
	// }

	// matrix[idx][jdx] = string(word[0])

	// return result
}

type stepFunc func(idx, jdx int) (int, int)

func up(idx, jdx int) (int, int) {
	return idx - 1, jdx
}

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
