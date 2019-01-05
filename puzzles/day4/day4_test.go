package day4

import (
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
			name: "example given",
			args: args{
				`
				[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:05] falls asleep
				[1518-11-01 00:25] wakes up
				[1518-11-01 00:30] falls asleep
				[1518-11-01 00:55] wakes up
				[1518-11-01 23:58] Guard #99 begins shift
				[1518-11-02 00:40] falls asleep
				[1518-11-02 00:50] wakes up
				[1518-11-03 00:05] Guard #10 begins shift
				[1518-11-03 00:24] falls asleep
				[1518-11-03 00:29] wakes up
				[1518-11-04 00:02] Guard #99 begins shift
				[1518-11-04 00:36] falls asleep
				[1518-11-04 00:46] wakes up
				[1518-11-05 00:03] Guard #99 begins shift
				[1518-11-05 00:45] falls asleep
				[1518-11-05 00:55] wakes up
				`,
			},
			want: 240,
		},
		{
			name:    "fail parse",
			args:    args{"[1518-11-01 00"},
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

func Test_max(t *testing.T) {
	type args struct {
		slice interface{}
		less  func(i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find max",
			args: args{[]int{3, 0, 21, -4}, func(i, j int) bool {
				arr := []int{3, 0, 21, -4}
				return arr[i] < arr[j]
			}},
			// index
			want: 2,
		},
		{
			name: "zero elements",
			args: args{[]int{}, func(i, j int) bool { return true }},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.slice, tt.args.less); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
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
			name: "example given",
			args: args{
				`
					[1518-11-01 00:00] Guard #10 begins shift
					[1518-11-01 00:05] falls asleep
					[1518-11-01 00:25] wakes up
					[1518-11-01 00:30] falls asleep
					[1518-11-01 00:55] wakes up
					[1518-11-01 23:58] Guard #99 begins shift
					[1518-11-02 00:40] falls asleep
					[1518-11-02 00:50] wakes up
					[1518-11-03 00:05] Guard #10 begins shift
					[1518-11-03 00:24] falls asleep
					[1518-11-03 00:29] wakes up
					[1518-11-04 00:02] Guard #99 begins shift
					[1518-11-04 00:36] falls asleep
					[1518-11-04 00:46] wakes up
					[1518-11-05 00:03] Guard #99 begins shift
					[1518-11-05 00:45] falls asleep
					[1518-11-05 00:55] wakes up
					`,
			},
			want: 4455,
		},
		{
			name:    "fail parse",
			args:    args{"[1518-11-01 00"},
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
