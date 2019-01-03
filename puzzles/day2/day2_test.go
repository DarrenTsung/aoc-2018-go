package day2

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example given",
			args: args{"abcdef\n bababc\n abbcde\n abcccd\n aabcdd\n abcdee\n ababab\n"},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePartOne(tt.args.input); got != tt.want {
				t.Errorf("SolvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDoublesAndTriples(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name          string
		args          args
		wantHasDouble bool
		wantHasTriple bool
	}{
		{
			name:          "no doubles or triples",
			args:          args{"abcdef"},
			wantHasDouble: false,
			wantHasTriple: false,
		},
		{
			name:          "1 double, 1 triple",
			args:          args{"bababc"},
			wantHasDouble: true,
			wantHasTriple: true,
		},
		{
			name:          "has doubles but no triples",
			args:          args{"abbcde"},
			wantHasDouble: true,
			wantHasTriple: false,
		},
		{
			name:          "has triple but no doubles",
			args:          args{"abcccd"},
			wantHasDouble: false,
			wantHasTriple: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHasDouble, gotHasTriple := checkDoublesAndTriples(tt.args.input)
			if gotHasDouble != tt.wantHasDouble {
				t.Errorf("checkDoublesAndTriples() gotHasDouble = %v, want %v", gotHasDouble, tt.wantHasDouble)
			}
			if gotHasTriple != tt.wantHasTriple {
				t.Errorf("checkDoublesAndTriples() gotHasTriple = %v, want %v", gotHasTriple, tt.wantHasTriple)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple example",
			args: args{"abc\nabd"},
			want: "ab",
		},
		{
			name: "more complicated",
			args: args{"abc\nabc\naii\nmii"},
			want: "ii",
		},
		{
			name: "no solution",
			// no words that are different by 1 character
			args: args{"abc\nabc"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePartTwo(tt.args.input); got != tt.want {
				t.Errorf("SolvePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
