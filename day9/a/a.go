package a

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func Run(input string) (_ int, err error) {
	diskMap := []int{}

	for jdx, disk := range strings.Split(input, "") {
		val := util.MustParseInt(disk)

		insertedValue := -1
		if jdx%2 == 0 {
			insertedValue = jdx / 2
		}

		for i := 0; i < val; i++ {
			diskMap = append(diskMap, insertedValue)
		}
	}

	// fmt.Println(diskMap)

	for i := len(diskMap) - 1; i > 0; i-- {
		if diskMap[i] == -1 {
			continue
		}

		for j := 0; j < i; j++ {
			if diskMap[j] == -1 {
				diskMap[j], diskMap[i] = diskMap[i], diskMap[j]
			}
		}
	}

	var result int

	for idx, disk := range diskMap {
		if disk == -1 {
			continue
		}

		result += idx * disk
	}

	// fmt.Println(diskMap)

	return result, nil
}
