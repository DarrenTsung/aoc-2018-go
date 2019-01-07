package day5

import "testing"

func TestSolvePartOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "aa",
			want:  2,
		},
		{
			input: "aA",
			want:  0,
		},
		{
			input: "abBA",
			want:  0,
		},
		{
			input: "ABba",
			want:  0,
		},
		{
			input: "abAB",
			want:  4,
		},
		{
			input: "aabAAB",
			want:  6,
		},
		{
			input: "aabAAbB",
			want:  5,
		},
		{
			input: "dabAcCaCBAcCcaDA",
			want:  10,
		},
		{
			input: "XXABCDEedcbaXX",
			want:  4,
		},
		{
			input: "XXABCDEedEFGgfecbaBbBXX\n",
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := SolvePartOne(tt.input); got != tt.want {
				t.Errorf("SolvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
