package day1

import (
	"testing"
)

func TestSolve(t *testing.T) {
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
			got, err := Solve(tt.args.input)
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
