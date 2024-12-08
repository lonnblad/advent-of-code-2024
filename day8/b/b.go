package b

import (
	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	matrix := util.NewMatrix(input)

	freqs := make(map[string]map[[2]int]bool)
	nodes := make(map[[2]int]bool)

	for idx, row := range matrix.Values {
		for jdx, cell := range row {
			if cell == "." {
				continue
			}

			cellValue := matrix.Values[idx][jdx]
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

				newFirstCoord := firstCoord
				for matrix.IsWithinBounds(newFirstCoord) {
					nodes[newFirstCoord] = true

					newFirstCoord[0] -= vector[0]
					newFirstCoord[1] -= vector[1]
				}

				newSecondCoord := secondCoord
				for matrix.IsWithinBounds(newSecondCoord) {
					nodes[newSecondCoord] = true

					newSecondCoord[0] += vector[0]
					newSecondCoord[1] += vector[1]
				}
			}
		}
	}

	return len(nodes), nil
}
