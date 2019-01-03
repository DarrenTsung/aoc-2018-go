package main

import (
	"aoc/puzzles/day1"
	"fmt"
	"io/ioutil"
)

func main() {
	solveAny("day1", day1.Solve)
}

func solveAny(dayName string, solutionFunc interface{}) {
	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("input/%s.txt", dayName))
	if err != nil {
		fmt.Println(dayName, "- error reading input file:", err)
		return
	}

	input := string(inputBytes)

	var res interface{}
	switch solutionFunc := solutionFunc.(type) {
	case func(string) (int, error):
		res, err = solutionFunc(input)
	default:
		fmt.Println(dayName, "- invalid solutionFunc passed:", solutionFunc)
		return
	}

	if err != nil {
		fmt.Println(dayName, "- err:", err)
	} else {
		fmt.Println(dayName, "- solution:", res)
	}
}
