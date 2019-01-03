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

// SolvePartTwo solves the day1 part2 puzzle from
// the input: https://adventofcode.com/2018/day/2
func SolvePartTwo(input string) string {
	lines := strings.FieldsFunc(input, func(r rune) bool { return unicode.IsSpace(r) })

	for i, line := range lines {
		for j := i; j < len(lines); j++ {
			otherLine := []rune(lines[j])

			differences := 0
			differenceIndex := -1
			for rIndex, r := range line {
				otherR := otherLine[rIndex]
				if r != otherR {
					differences++
					differenceIndex = rIndex
					if differences > 1 {
						continue
					}
				}
			}

			if differences == 1 {
				return line[:differenceIndex] + line[differenceIndex+1:]
			}
		}
	}

	return ""
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
