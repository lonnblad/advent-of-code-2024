package day4

import "strings"

func B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	// re := regexp.MustCompile(`X.*M.*A.*S`)

	matrix := make([][]string, len(rows))

	for idx, row := range rows {
		matrix[idx] = strings.Split(row, "")
	}

	var result int

	for idx, row := range matrix {
		for jdx, val := range row {
			if val == "M" {
				dr := findWord(matrix, idx, jdx, downRight, "MAS")
				ur1 := findWord(matrix, idx+2, jdx, upRight, "MAS")
				ur2 := findWord(matrix, idx+2, jdx, upRight, "SAM")

				if dr == 1 && (ur1 == 1 || ur2 == 1) {
					result++
				}
			}

			if val == "S" {
				dr := findWord(matrix, idx, jdx, downRight, "SAM")
				ur1 := findWord(matrix, idx+2, jdx, upRight, "SAM")
				ur2 := findWord(matrix, idx+2, jdx, upRight, "MAS")

				if dr == 1 && (ur1 == 1 || ur2 == 1) {
					result++
				}
			}
		}
	}

	return result, nil
}
