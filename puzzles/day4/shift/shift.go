package shift

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Shift represents a record for a single shift that a guard had
type Shift struct {
	GuardID int
	Asleep  []MinuteSpan
}

// MinuteSpan represents an range from [start, end)
type MinuteSpan struct {
	Start, End int
}

var timestampRegex = regexp.MustCompile(`\[([^\]]+)\] (.*)`)

// ParseGuardShifts takes in an input string and returns a list
// of shifts in chronological order
func ParseGuardShifts(input string) (shifts []Shift, err error) {
	const timeLayout = "2006-01-02 15:04"
	type TimestampedEntry struct {
		raw   string
		time  time.Time
		entry string
	}

	timestampedEntries := []TimestampedEntry{}
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		values := timestampRegex.FindStringSubmatch(line)
		if values == nil {
			return nil, fmt.Errorf("line does not match timestamped entry: %s", line)
		}

		time, err := time.Parse(timeLayout, values[1])
		if err != nil {
			return nil, fmt.Errorf("failed to parse time from: %s, err: %s", values[1], err)
		}

		timestampedEntries = append(
			timestampedEntries,
			TimestampedEntry{raw: strings.TrimSpace(line), time: time, entry: values[2]},
		)
	}

	sort.Slice(timestampedEntries, func(i, j int) bool {
		return timestampedEntries[i].time.Before(timestampedEntries[j].time)
	})

	var currentShift *Shift
	var sleepingStart *int
	for _, timestampedEntry := range timestampedEntries {
		entry := strings.TrimSpace(timestampedEntry.entry)

		var guardID int
		_, err := fmt.Sscanf(entry, "Guard #%d begins shift", &guardID)
		if err == nil {
			if currentShift != nil {
				shifts = append(shifts, *currentShift)
			}

			currentShift = &Shift{GuardID: guardID, Asleep: []MinuteSpan{}}
		} else {
			if currentShift == nil {
				return nil, fmt.Errorf("failed to find starting shift, first entry is: %s", timestampedEntry.raw)
			}

			if entry == "falls asleep" {
				if sleepingStart != nil {
					return nil, fmt.Errorf("can't fall asleep when already asleep: %s", timestampedEntry.raw)
				}

				minute := timestampedEntry.time.Minute()
				sleepingStart = &minute
			} else if entry == "wakes up" {
				if sleepingStart == nil {
					return nil, fmt.Errorf("can't wake up if not already asleep: %s", timestampedEntry.raw)
				}

				endMinute := timestampedEntry.time.Minute()
				currentShift.Asleep = append(currentShift.Asleep, MinuteSpan{Start: *sleepingStart, End: endMinute})
				sleepingStart = nil
			} else {
				return nil, fmt.Errorf("failed to parse entry: %s", entry)
			}
		}
	}

	if currentShift != nil {
		shifts = append(shifts, *currentShift)
	}

	return shifts, nil
}
