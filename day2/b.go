package day2

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func B(input string) (_ int, err error) {
	result := 0

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		levels := strings.Split(row, " ")

		for i := -1; i < len(levels); i++ {
			if checkLevels(levels, i) {
				result++
				break
			}
		}
	}

	return result, nil
}

func checkLevels(levels []string, skipLevel int) bool {
	var prevLevel *int
	var increase *bool
	safe := true

	for idx, l := range levels {
		if idx == skipLevel {
			continue
		}

		level := util.MustParseInt(l)

		if prevLevel == nil {
			prevLevel = &level
			continue
		}

		if level > *prevLevel {
			if increase == nil {
				increase = util.Ptr(true)
			} else if !*increase {
				safe = false
			}

			if level-*prevLevel > 3 {
				safe = false
			}
		} else if level < *prevLevel {
			if increase == nil {
				increase = util.Ptr(false)
			} else if *increase {
				safe = false
			}

			if *prevLevel-level > 3 {
				safe = false
			}
		} else {
			safe = false
		}

		prevLevel = &level
	}

	return safe
}
