package day1

import (
	"strconv"
	"strings"
	"unicode"
)

// SolvePartOne solves the day1 part1 puzzle from
// the input: https://adventofcode.com/2018/day/1
func SolvePartOne(input string) (int, error) {
	output, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, v := range output {
		sum += v
	}

	return sum, nil
}

// SolvePartTwo solves the day1 part2 puzzle from
// the input: https://adventofcode.com/2018/day/1
func SolvePartTwo(input string) (int, error) {
	output, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	valuesSeen := map[int]bool{}
	valuesSeen[0] = true
	sum := 0
	for {
		for _, v := range output {
			sum += v
			if valuesSeen[sum] {
				return sum, nil
			}
			valuesSeen[sum] = true
		}
	}
}

func parseInput(input string) ([]int, error) {
	isCommaOrWhitespace := func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	}

	output := []int{}
	for _, s := range strings.FieldsFunc(input, isCommaOrWhitespace) {
		v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
		if err != nil {
			return output, err
		}

		output = append(output, int(v))
	}

	return output, nil
}
