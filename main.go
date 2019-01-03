package main

import (
	"aoc/part"
	"aoc/puzzles/day1"
	"aoc/puzzles/day2"
	"fmt"
	"io/ioutil"
)

func main() {
	solveAny("Day1", part.One, day1.SolvePartOne)
	solveAny("Day1", part.Two, day1.SolvePartTwo)

	solveAny("Day2", part.One, day2.SolvePartOne)
	solveAny("Day2", part.Two, day2.SolvePartTwo)
}

func solveAny(dayName string, part part.Part, solutionFunc interface{}) {
	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("input/%s.txt", dayName))
	if err != nil {
		fmt.Println(dayName, part, "- error reading input file:", err)
		return
	}

	input := string(inputBytes)

	var res interface{}
	switch solutionFunc := solutionFunc.(type) {
	case func(string) (int, error):
		res, err = solutionFunc(input)
	case func(string) int:
		res = solutionFunc(input)
	case func(string) string:
		res = solutionFunc(input)
	default:
		fmt.Println(dayName, part, "- invalid solutionFunc passed:", solutionFunc)
		return
	}

	if err != nil {
		fmt.Println(dayName, part, "- err:", err)
	} else {
		fmt.Println(dayName, part, "- solution:", res)
	}
}
