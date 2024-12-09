package b

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	nos := strings.Split(input, "")
	diskMap := make([][]int, len(nos))
	spacesToFill := make([]int, len(nos))
	spacesFilled := make([]int, len(nos))

	for jdx, disk := range nos {
		val := util.MustParseInt(disk)

		insertValue := -1

		if jdx%2 == 0 {
			insertValue = jdx / 2
		} else {
			spacesToFill[jdx] = val
		}

		for i := 0; i < val; i++ {
			diskMap[jdx] = append(diskMap[jdx], insertValue)
		}
	}

	// fmt.Println(diskMap)
	// fmt.Println(spacesToFill)

	for i := len(diskMap) - 1; i > 0; i-- {
		if len(diskMap[i]) == 0 || diskMap[i][0] == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if spacesToFill[j] < len(diskMap[i]) {
				continue
			}

			for idx := range diskMap[i] {
				diskMap[j][idx+spacesFilled[j]], diskMap[i][idx] = diskMap[i][idx], -1
			}

			spacesToFill[j] -= len(diskMap[i])
			spacesFilled[j] += len(diskMap[i])

			// fmt.Println("-----------------")
			// fmt.Println(diskMap[j])
			// fmt.Println(diskMap)
			// fmt.Println(spacesToFill)

			break
		}
	}

	// fmt.Println("-----------------")
	// fmt.Println(diskMap)

	var result int

	n := 0

	for _, block := range diskMap {
		for _, disk := range block {
			if disk == -1 {
				n++
				continue
			}

			result += n * disk
			n++
		}
	}

	return result, nil
}
