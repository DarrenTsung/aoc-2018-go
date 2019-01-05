package day4

import (
	"aoc/puzzles/day4/shift"
	"reflect"
)

// SolvePartOne solves day4 part1.
func SolvePartOne(input string) (int, error) {
	guardInfos, err := getGuardInfos(input)
	if err != nil {
		return 0, err
	}

	var sleepiestGuardInfo *guardInfo
	var sleepiestGuardID int
	for guardID, info := range guardInfos {
		if sleepiestGuardInfo == nil || sleepiestGuardInfo.totalAsleep < info.totalAsleep {
			sleepiestGuardID = guardID
			sleepiestGuardInfo = info
		}
	}

	sleepiestMinute := max(sleepiestGuardInfo.asleepMinutes, func(i, j int) bool {
		return sleepiestGuardInfo.asleepMinutes[i] < sleepiestGuardInfo.asleepMinutes[j]
	})

	return sleepiestGuardID * sleepiestMinute, nil
}

// SolvePartTwo solves day4 part2.
func SolvePartTwo(input string) (int, error) {
	guardInfos, err := getGuardInfos(input)
	if err != nil {
		return 0, err
	}

	var sleepiestMinute *int
	var sleepiestMinuteAmount int
	var sleepiestGuardID int
	for guardID, info := range guardInfos {
		for minute := 0; minute < 60; minute++ {
			if sleepiestMinute == nil || sleepiestMinuteAmount < info.asleepMinutes[minute] {
				// Because we're using pointers we have to copy a new int to a different address
				var storedMinute = minute
				sleepiestMinute = &storedMinute
				sleepiestMinuteAmount = info.asleepMinutes[minute]
				sleepiestGuardID = guardID
			}
		}
	}

	return sleepiestGuardID * *sleepiestMinute, nil
}

type guardInfo struct {
	totalAsleep   int
	asleepMinutes [60]int
}

func getGuardInfos(input string) (map[int]*guardInfo, error) {
	shifts, err := shift.ParseGuardShifts(input)
	if err != nil {
		return nil, err
	}

	guardInfos := map[int]*guardInfo{}

	for _, shift := range shifts {
		if guardInfos[shift.GuardID] == nil {
			guardInfos[shift.GuardID] = &guardInfo{}
		}

		for _, asleepSpan := range shift.Asleep {
			guardInfos[shift.GuardID].totalAsleep += asleepSpan.End - 1 - asleepSpan.Start
			for minute := asleepSpan.Start; minute < asleepSpan.End; minute++ {
				asleepMinutes := guardInfos[shift.GuardID].asleepMinutes
				asleepMinutes[minute]++
				guardInfos[shift.GuardID].asleepMinutes = asleepMinutes
			}
		}
	}

	return guardInfos, nil
}

func max(slice interface{}, less func(i, j int) bool) int {
	s := reflect.ValueOf(slice)
	len := s.Len()
	if len <= 0 {
		return -1
	}

	currentIndex := 0
	for i := 1; i < len; i++ {
		if less(currentIndex, i) {
			currentIndex = i
		}
	}

	return currentIndex
}
