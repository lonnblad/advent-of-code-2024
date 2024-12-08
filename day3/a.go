package day3

import (
	"regexp"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func A(input string) (_ int, err error) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0

	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a := util.MustParseInt(match[1])
		b := util.MustParseInt(match[2])
		sum += a * b
	}

	return sum, nil
}
