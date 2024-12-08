package a

import (
	"regexp"
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	equations := strings.Split(input, "\n")

	re := regexp.MustCompile(`(\d+)`)

	var result int

	for _, equation := range equations {
		matches := re.FindAllStringSubmatch(equation, -1)

		var testValue int
		nos := make([]int, len(matches)-1)

		for idx, match := range matches {
			if idx == 0 {
				testValue = util.MustParseInt(match[0])
				continue
			}

			nos[idx-1] = util.MustParseInt(match[0])
		}

		if solveable(testValue, 0, nos) {
			result += testValue
		}
	}

	return result, nil
}

func solveable(testValue, ackumulator int, nos []int) bool {
	if len(nos) == 0 {
		return false
	}

	if testValue == ackumulator*nos[0] || testValue == ackumulator+nos[0] {
		return true
	}

	return solveable(testValue, ackumulator*nos[0], nos[1:]) || solveable(testValue, ackumulator+nos[0], nos[1:])
}
