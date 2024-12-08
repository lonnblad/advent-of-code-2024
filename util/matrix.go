package util

import "strings"

func StringMatrix(a string) [][]string {
	rows := strings.Split(a, "\n")
	result := make([][]string, len(rows))

	for i, row := range rows {
		result[i] = strings.Split(row, "")
	}

	return result
}

type Matrix struct {
	Values  [][]string
	noRows  int
	noCells int
}

func NewMatrix(a string) Matrix {
	rows := strings.Split(a, "\n")
	result := make([][]string, len(rows))

	for i, row := range rows {
		result[i] = strings.Split(row, "")
	}

	return Matrix{
		Values:  result,
		noRows:  len(result),
		noCells: len(result[0]),
	}
}

func (m Matrix) IsWithinBounds(coord [2]int) bool {
	return coord[0] >= 0 && coord[0] < m.noRows && coord[1] >= 0 && coord[1] < m.noCells
}
