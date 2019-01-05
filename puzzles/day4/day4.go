package day4

import (
	"aoc/puzzles/day4/shift"
	"reflect"
)

// SolvePartOne solves day4 part1.
func SolvePartOne(input string) (int, error) {
	shifts, err := shift.ParseGuardShifts(input)
	if err != nil {
		return 0, err
	}

	type GuardInfo struct {
		totalAsleep   int
		asleepMinutes [60]int
	}

	guardInfos := map[int]*GuardInfo{}

	for _, shift := range shifts {
		if guardInfos[shift.GuardID] == nil {
			guardInfos[shift.GuardID] = &GuardInfo{}
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

	var sleepiestGuardInfo *GuardInfo
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
