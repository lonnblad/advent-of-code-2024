package util

import (
	"strconv"
)

func Ptr[T any](v T) *T {
	return &v
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Transpose(a []string) []string {
	newArr := make([][]byte, len(a[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	newStringArr := make([]string, len(newArr))
	for idx, row := range newArr {
		newStringArr[idx] = string(row)
	}

	return newStringArr
}

func Rotate45Degrees(a []string) []string {
	n := len(a)
	if n == 0 {
		return nil
	}

	m := len(a[0])
	if m == 0 {
		return nil
	}

	// The size of the resulting array will be n + m - 1
	result := make([]string, n+m-1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result[i+j] += string(a[i][j])
		}
	}

	return result
}
