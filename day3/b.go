package day3

import (
	"regexp"
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func B(input string) (_ int, err error) {
	dos := strings.Split(input, "do()")

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0

	for _, do := range dos {
		parts := strings.Split(do, "don't()")

		matches := re.FindAllStringSubmatch(parts[0], -1)
		for _, match := range matches {
			a := util.MustParseInt(match[1])
			b := util.MustParseInt(match[2])
			sum += a * b
		}
	}

	return sum, nil
}
