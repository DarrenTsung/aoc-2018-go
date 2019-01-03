package part

import "log"

// Part is either One or Two
type Part int

const (
	// One is the first part
	One Part = iota
	// Two is the second part
	Two
)

func (part Part) String() string {
	switch part {
	case One:
		return "Part One"
	case Two:
		return "Part Two"
	default:
		log.Panic("Impossible!")
		return ""
	}
}
