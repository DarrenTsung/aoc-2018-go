package day3

import (
	"aoc/puzzles/day3/rect"
	"reflect"
	"testing"
)

func Test_parseClaims(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantClaims []claim
		wantErr    bool
	}{
		{
			name:       "single",
			args:       args{"#1 @ 1,3: 4x4"},
			wantClaims: []claim{claim{rect.Rect{X: 1, Y: 3, Width: 4, Height: 4}, 1}},
			wantErr:    false,
		},
		{
			name: "multiple",
			args: args{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"},
			wantClaims: []claim{
				claim{rect.Rect{X: 1, Y: 3, Width: 4, Height: 4}, 1},
				claim{rect.Rect{X: 3, Y: 1, Width: 4, Height: 4}, 2},
				claim{rect.Rect{X: 5, Y: 5, Width: 2, Height: 2}, 3},
			},
			wantErr: false,
		},
		{
			name: "invalid single",
			// regex will ignore invalid claims
			args:       args{"#1 @ 1,3ii: 4x4"},
			wantClaims: nil,
			wantErr:    false,
		},
		{
			name: "invalid too large 1",
			// regex will not ignore too large integers
			args:       args{"#1 @ 9999999999999999999999999999,1: 4x4"},
			wantClaims: nil,
			wantErr:    true,
		},
		{
			name: "invalid too large 2",
			// regex will not ignore too large integers
			args:       args{"#1 @ 1,9999999999999999999999999999: 4x4"},
			wantClaims: nil,
			wantErr:    true,
		},
		{
			name: "invalid too large 3",
			// regex will not ignore too large integers
			args:       args{"#1 @ 1,2: 9999999999999999999999999999x4"},
			wantClaims: nil,
			wantErr:    true,
		},
		{
			name: "invalid too large 4",
			// regex will not ignore too large integers
			args:       args{"#1 @ 1,2: 4x9999999999999999999999999999"},
			wantClaims: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClaims, err := parseClaims(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseClaims() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClaims, tt.wantClaims) {
				t.Errorf("parseClaims() = %v, want %v", gotClaims, tt.wantClaims)
			}
		})
	}
}

func TestSolvePartOne(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "given",
			args:    args{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"},
			want:    4,
			wantErr: false,
		},
		{
			name:    "parse error",
			args:    args{"#1 @ 1,3: 99999999999999999994x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartOne(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePartOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolvePartTwo(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "given",
			args: args{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"},
			// claim #3 does not overlap with any others
			want:    3,
			wantErr: false,
		},
		{
			name:    "all overlap",
			args:    args{"#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 3,1: 2x2"},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "parse error",
			args:    args{"#1 @ 1,3: 99999999999999999994x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartTwo(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePartTwo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
