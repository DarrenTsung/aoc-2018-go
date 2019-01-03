package day1

import (
	"reflect"
	"testing"
)

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
			name:    "all positive",
			args:    args{"+1, +1, +1"},
			want:    3,
			wantErr: false,
		},
		{
			name:    "positive & negative",
			args:    args{"+1, +1, -2"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "different formats",
			args:    args{"+1,1,-2"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "all negatives",
			args:    args{"-1, -2, -3"},
			want:    -6,
			wantErr: false,
		},
		{
			name:    "some whitespace",
			args:    args{"    -1  ,-2  , -3  "},
			want:    -6,
			wantErr: false,
		},
		{
			name:    "bad numbers",
			args:    args{"_1, 23.45"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "newlines",
			args:    args{"1\n 2  \n3"},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePartOne(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
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
			name:    "simple",
			args:    args{"+1, -1"},
			want:    0,
			wantErr: false,
		},
		{
			name:    "loops",
			args:    args{"+3, +3, +4, -2, -4"},
			want:    10,
			wantErr: false,
		},
		{
			name:    "medium1",
			args:    args{"-6, +3, +8, +5, -6"},
			want:    5,
			wantErr: false,
		},
		{
			name:    "medium2",
			args:    args{"+7, +7, -2, -7, -4"},
			want:    14,
			wantErr: false,
		},
		{
			name:    "bad parsing",
			args:    args{"+7, ++7, _2, -7.14, -4"},
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

func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "some whitespace",
			args:    args{"    -1  ,-2  , -3  "},
			want:    []int{-1, -2, -3},
			wantErr: false,
		},
		{
			name:    "bad numbers",
			args:    args{"_1, 23.45"},
			want:    []int{},
			wantErr: true,
		},
		{
			name:    "newlines",
			args:    args{"1\n 2  \n3"},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
