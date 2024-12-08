package day2

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func A(input string) (_ int, err error) {
	result := 0

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		levels := strings.Split(row, " ")

		var prevLevel *int
		var increase *bool
		safe := true

		for _, l := range levels {
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

		if safe {
			result++
		}
	}

	return result, nil
}
