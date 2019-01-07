package day5

import (
	"math"
	"strings"
	"unicode"
)

// SolvePartOne solves day5 part1
func SolvePartOne(input string) int {
	runes := []rune(strings.TrimSpace(input))
	madeChange := true

	if len(runes) <= 0 {
		return 0
	}

	for madeChange {
		madeChange = false
		len := len(runes)
		if len <= 0 {
			break
		}

		newRunes := make([]rune, 0, len)
		var cachedLowerR *rune

		for i := 0; i < len; i++ {
			r := runes[i]

			// if end, automatically append
			if len-1 == i {
				newRunes = append(newRunes, r)
				continue
			}

			nextR := runes[i+1]

			// the two characters react and destroy each other
			if r != nextR {
				if cachedLowerR == nil {
					lowerR := unicode.ToLower(r)
					cachedLowerR = &lowerR
				}

				nextLowerR := unicode.ToLower(nextR)
				if *cachedLowerR == nextLowerR {
					cachedLowerR = nil
					madeChange = true
					i++
					continue
				} else {
					cachedLowerR = &nextLowerR
				}
			}

			newRunes = append(newRunes, r)
		}

		runes = newRunes
	}

	return len(runes)
}

// SolvePartTwo solves day5 part2
func SolvePartTwo(input string) int {
	input = strings.TrimSpace(input)

	chars := map[rune]bool{}
	for _, r := range input {
		chars[unicode.ToLower(r)] = true
	}

	min := math.MaxInt32
	for r := range chars {
		removedInput := strings.Replace(input, string(r), "", -1)
		removedInput = strings.Replace(removedInput, string(unicode.ToUpper(r)), "", -1)

		s := SolvePartOne(removedInput)
		if min > s {
			min = s
		}
	}

	return min
}
