package day5

import (
	"math"
	"strings"
	"unicode"
)

// SolvePartOne solves day5 part1
func SolvePartOne(input string) int {
	runes := []rune(strings.TrimSpace(input))
	if len(runes) <= 0 {
		return 0
	}

	newRunes := []rune{runes[0]}
	for _, r := range runes[1:] {
		if len(newRunes) > 0 {
			lastR := newRunes[len(newRunes)-1]
			if lastR != r && unicode.ToLower(lastR) == unicode.ToLower(r) {
				newRunes = newRunes[:len(newRunes)-1]
				continue
			}
		}

		newRunes = append(newRunes, r)
	}

	return len(newRunes)
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
