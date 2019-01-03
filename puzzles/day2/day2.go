package day2

import (
	"strings"
	"unicode"
)

// SolvePartOne solves the day1 part1 puzzle from
// the input: https://adventofcode.com/2018/day/2
func SolvePartOne(input string) int {
	doubleCount := 0
	tripleCount := 0

	for _, line := range strings.FieldsFunc(input, func(r rune) bool { return unicode.IsSpace(r) }) {
		hasDouble, hasTriple := checkDoublesAndTriples(line)
		if hasDouble {
			doubleCount++
		}
		if hasTriple {
			tripleCount++
		}
	}

	return doubleCount * tripleCount
}

func checkDoublesAndTriples(input string) (hasDouble bool, hasTriple bool) {
	occurences := map[rune]int{}
	for _, c := range input {
		occurences[c]++
	}

	for _, count := range occurences {
		if count == 2 {
			hasDouble = true
		} else if count == 3 {
			hasTriple = true
		}
	}

	return
}
