package a

import (
	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	matrix := util.StringMatrix(input)

	freqs := make(map[string]map[[2]int]bool)
	nodes := make(map[[2]int]bool)

	for idx, row := range matrix {
		for jdx, cell := range row {
			if cell == "." {
				continue
			}

			cellValue := matrix[idx][jdx]
			if _, ok := freqs[cellValue]; !ok {
				freqs[cellValue] = make(map[[2]int]bool)
			}

			freqs[cellValue][[2]int{idx, jdx}] = true
		}
	}

	for key := range freqs {
		for firstCoord := range freqs[key] {
			for secondCoord := range freqs[key] {
				if firstCoord == secondCoord {
					continue
				}

				vector := [2]int{secondCoord[0] - firstCoord[0], secondCoord[1] - firstCoord[1]}

				newFirstCoord := [2]int{firstCoord[0] - vector[0], firstCoord[1] - vector[1]}
				newSecondCoord := [2]int{secondCoord[0] + vector[0], secondCoord[1] + vector[1]}

				if newFirstCoord[0] >= 0 && newFirstCoord[0] < len(matrix) && newFirstCoord[1] >= 0 && newFirstCoord[1] < len(matrix[0]) {
					nodes[newFirstCoord] = true
				}

				if newSecondCoord[0] >= 0 && newSecondCoord[0] < len(matrix) && newSecondCoord[1] >= 0 && newSecondCoord[1] < len(matrix[0]) {
					nodes[newSecondCoord] = true
				}
			}
		}
	}

	return len(nodes), nil
}
