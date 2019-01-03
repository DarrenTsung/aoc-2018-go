package main

import (
	"aoc/part"
	"aoc/puzzles/day1"
	"fmt"
	"io/ioutil"
)

func main() {
	solveAny("Day1", part.One, day1.SolvePartOne)
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
