package day5

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func A(input string) (_ int, err error) {
	var result int

	parts := strings.Split(input, "\n\n")
	rules, updates := strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")

	rulesMap := make(map[string]map[string]bool)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		if _, ok := rulesMap[parts[0]]; !ok {
			rulesMap[parts[0]] = make(map[string]bool)
		}

		rulesMap[parts[0]][parts[1]] = true
	}

	for _, update := range updates {
		ok := true

		pages := strings.Split(update, ",")

		for idx := 0; idx < len(pages)-1; idx++ {
			page := pages[idx]
			nextPage := pages[idx+1]

			if _, exist := rulesMap[page][nextPage]; !exist {
				ok = false
				break
			}
		}

		if ok {
			midIdx := len(pages) / 2
			result += util.MustParseInt(pages[midIdx])
		}
	}

	return result, nil
}
