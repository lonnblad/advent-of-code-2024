package day1

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func Day1A(input string) (_ int, err error) {
	rows := strings.Split(input, "\n")

	leftNos := make([]int, len(rows))
	rightNos := make([]int, len(rows))

	re := regexp.MustCompile(`(\d+)\s+(\d+)`)

	for idx, row := range rows {
		matches := re.FindAllStringSubmatch(row, -1)

		if len(matches[0]) != 3 {
			err = fmt.Errorf("unexpected number of matches: %d", len(matches))
			return
		}

		leftNos[idx] = util.MustParseInt(matches[0][1])
		rightNos[idx] = util.MustParseInt(matches[0][2])
	}

	sort.Ints(leftNos)
	sort.Ints(rightNos)

	diffs := 0
	for idx, leftNo := range leftNos {
		rightNo := rightNos[idx]
		diff := rightNo - leftNo

		if diff < 0 {
			diff = -diff
		}

		diffs += diff
	}

	return diffs, nil
}
