package day5

import (
	"strings"

	"github.com/lonnblad/advent-of-code-2024/util"
)

func B(input string) (_ int, err error) {
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

	checkRules := func(pages []string) bool {
		for idx := 0; idx < len(pages)-1; idx++ {
			page := pages[idx]
			nextPage := pages[idx+1]

			if _, exist := rulesMap[page][nextPage]; !exist {
				return false
			}
		}

		return true
	}

	for _, update := range updates {
		pages := strings.Split(update, ",")

		if checkRules(pages) {
			continue
		}

		for idx := 0; idx < len(pages)-1; idx++ {
			for jdx := idx + 1; jdx < len(pages); jdx++ {
				page := pages[idx]
				nextPage := pages[jdx]

				if _, exist := rulesMap[page][nextPage]; !exist {
					pages[idx], pages[jdx] = pages[jdx], pages[idx]
				}
			}
		}

		midIdx := len(pages) / 2
		result += util.MustParseInt(pages[midIdx])
	}

	return result, nil
}
