package day1

import (
	"strconv"
	"strings"
	"unicode"
)

// Solve the day1 puzzle from the input: https://adventofcode.com/2018/day/1
func Solve(input string) (int, error) {
	isCommaOrWhitespace := func(c rune) bool {
		return c == ',' || unicode.IsSpace(c)
	}

	sum := 0
	for _, s := range strings.FieldsFunc(input, isCommaOrWhitespace) {
		v, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
		if err != nil {
			return 0, err
		}

		sum += int(v)
	}

	return sum, nil
}
