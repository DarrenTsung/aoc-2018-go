package day5

import (
	"strings"
	"unicode"
)

// SolvePartOne solves day5 part1
func SolvePartOne(input string) int {
	runes := []rune(strings.TrimSpace(input))
	madeChange := true

	for madeChange {
		madeChange = false
		newRunes := []rune{}

		len := len(runes)
		for i := 0; i < len; i++ {
			r := runes[i]

			// if end, automatically append
			if len-1 == i {
				newRunes = append(newRunes, r)
				continue
			}

			nextR := runes[i+1]

			isLower := unicode.IsLower(r)
			nextIsLower := unicode.IsLower(nextR)

			// the two characters react and destroy each other
			if isLower != nextIsLower && unicode.ToLower(r) == unicode.ToLower(nextR) {
				madeChange = true
				i++
				continue
			}

			newRunes = append(newRunes, r)
		}

		runes = newRunes
	}

	return len(runes)
}
