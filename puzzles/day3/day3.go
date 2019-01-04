package day3

import (
	"aoc/puzzles/day3/rect"
	"regexp"
	"strconv"
	"strings"
)

// SolvePartOne solves Day3 Part1
func SolvePartOne(input string) (int, error) {
	claims, err := parseClaims(input)
	if err != nil {
		return 0, err
	}

	occupiedCount := map[rect.Point]int{}
	for _, claim := range claims {
		claim.rect.IterPoints(func(point rect.Point) {
			occupiedCount[point]++
		})
	}

	overlapCount := 0
	for _, count := range occupiedCount {
		if count > 1 {
			overlapCount++
		}
	}

	return overlapCount, nil
}

// SolvePartTwo solves Day3 Part2
func SolvePartTwo(input string) (int, error) {
	claims, err := parseClaims(input)
	if err != nil {
		return 0, err
	}

	occupiedCount := map[rect.Point]int{}
	for _, claim := range claims {
		claim.rect.IterPoints(func(point rect.Point) {
			occupiedCount[point]++
		})
	}

	for _, claim := range claims {
		overlaps := false
		claim.rect.IterPoints(func(point rect.Point) {
			if occupiedCount[point] > 1 {
				overlaps = true
			}
		})

		if !overlaps {
			return claim.id, nil
		}
	}

	return -1, nil
}

type claim struct {
	rect rect.Rect
	id   int
}

var claimRegex = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func parseClaims(input string) (claims []claim, err error) {
	for _, line := range strings.Split(input, "\n") {
		values := claimRegex.FindStringSubmatch(line)
		if values != nil {
			id, err := strconv.Atoi(values[1])
			if err != nil {
				return nil, err
			}

			x, err := strconv.Atoi(values[2])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(values[3])
			if err != nil {
				return nil, err
			}

			width, err := strconv.Atoi(values[4])
			if err != nil {
				return nil, err
			}

			height, err := strconv.Atoi(values[5])
			if err != nil {
				return nil, err
			}

			claims = append(claims, claim{rect.Rect{X: x, Y: y, Width: width, Height: height}, id})
		}
	}

	return
}
