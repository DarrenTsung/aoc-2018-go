package rect

import (
	"reflect"
	"testing"
)

func TestRect_IterPoints(t *testing.T) {
	type fields struct {
		X      int
		Y      int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   []Point
	}{
		{
			name: "simple",
			fields: fields{
				X:      0,
				Y:      0,
				Width:  1,
				Height: 2,
			},
			want: []Point{Point{X: 0, Y: 0}, Point{X: 0, Y: 1}},
		},
		{
			name: "empty",
			fields: fields{
				X:      0,
				Y:      0,
				Width:  0,
				Height: 0,
			},
			want: []Point{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := Rect{
				X:      tt.fields.X,
				Y:      tt.fields.Y,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			got := []Point{}
			rect.IterPoints(func(point Point) { got = append(got, point) })

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rect.IterPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_ContainsPoint(t *testing.T) {
	type fields struct {
		X      int
		Y      int
		Width  int
		Height int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "simple1",
			fields: fields{
				X:      0,
				Y:      0,
				Width:  1,
				Height: 2,
			},
			args: args{x: 0, y: 0},
			want: true,
		},
		{
			name: "simple2",
			fields: fields{
				X:      0,
				Y:      0,
				Width:  1,
				Height: 2,
			},
			args: args{x: 0, y: 1},
			want: true,
		},
		{
			name: "simple3",
			fields: fields{
				X:      0,
				Y:      0,
				Width:  1,
				Height: 2,
			},
			args: args{x: 0, y: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rect := Rect{
				X:      tt.fields.X,
				Y:      tt.fields.Y,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := rect.ContainsPoint(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Rect.ContainsPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
