package shift

import (
	"reflect"
	"testing"
)

func TestParseGuardShifts(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantShifts []Shift
		wantErr    bool
	}{
		{
			name: "single ordered",
			args: args{
				`[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:05] falls asleep
				[1518-11-01 00:25] wakes up
				[1518-11-01 00:30] falls asleep
				[1518-11-01 00:55] wakes up`,
			},
			wantShifts: []Shift{
				Shift{
					GuardID: 10,
					Asleep: []MinuteSpan{
						MinuteSpan{Start: 5, End: 25},
						MinuteSpan{Start: 30, End: 55},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple unordered",
			args: args{
				`
				[1518-11-01 00:05] falls asleep
				[1518-11-03 00:05] Guard #10 begins shift
				[1518-11-03 00:24] falls asleep
				[1518-11-01 00:30] falls asleep
				[1518-11-02 00:50] wakes up
				[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:55] wakes up
				[1518-11-01 23:58] Guard #99 begins shift
				[1518-11-02 00:40] falls asleep
				[1518-11-01 00:25] wakes up
				[1518-11-03 00:29] wakes up
				`,
			},
			wantShifts: []Shift{
				Shift{
					GuardID: 10,
					Asleep: []MinuteSpan{
						MinuteSpan{Start: 5, End: 25},
						MinuteSpan{Start: 30, End: 55},
					},
				},
				Shift{
					GuardID: 99,
					Asleep: []MinuteSpan{
						MinuteSpan{Start: 40, End: 50},
					},
				},
				Shift{
					GuardID: 10,
					Asleep: []MinuteSpan{
						MinuteSpan{Start: 24, End: 29},
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "very simple",
			args:       args{"[1518-11-01 00:00] Guard #10 begins shift"},
			wantShifts: []Shift{Shift{GuardID: 10, Asleep: []MinuteSpan{}}},
			wantErr:    false,
		},
		{
			name:    "not timestamped entry",
			args:    args{"just random garbage"},
			wantErr: true,
		},
		{
			name:    "invalid timestamp",
			args:    args{"[1518-11-01 x00:00] Guard #10 begins shift"},
			wantErr: true,
		},
		{
			name:    "invalid begins shift",
			args:    args{"[1518-11-01 00:00] Guardxx #10 begins shift"},
			wantErr: true,
		},
		{
			name: "invalid entry after beginning shift",
			args: args{
				`[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:55] dances a dance`,
			},
			wantErr: true,
		},
		{
			name: "falls asleep twice",
			args: args{
				`[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:05] falls asleep
				[1518-11-01 00:30] falls asleep
				[1518-11-01 00:55] wakes up`,
			},
			wantErr: true,
		},
		{
			name: "wakes up twice",
			args: args{
				`[1518-11-01 00:00] Guard #10 begins shift
				[1518-11-01 00:05] falls asleep
				[1518-11-01 00:30] wakes up
				[1518-11-01 00:55] wakes up`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShifts, err := ParseGuardShifts(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGuardShifts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotShifts, tt.wantShifts) {
				t.Errorf("ParseGuardShifts() = %v, want %v", gotShifts, tt.wantShifts)
			}
		})
	}
}
