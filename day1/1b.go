package day1

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func Day1B(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	leftNos := make([]int, len(rows))
	rightNos := make(map[int]int, len(rows))

	re := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for idx, row := range rows {
		matches := re.FindAllStringSubmatch(row, -1)

		if len(matches[0]) != 3 {
			err = fmt.Errorf("unexpected number of matches: %d", len(matches))
			return
		}

		leftNos[idx] = util.MustParseInt(matches[0][1])
		rightNo := util.MustParseInt(matches[0][2])
		rightNos[rightNo] = rightNos[rightNo] + 1
	}

	diffs := 0
	for _, leftNo := range leftNos {
		rightNo := rightNos[leftNo]
		diffs += leftNo * rightNo
	}

	return diffs, nil
}
