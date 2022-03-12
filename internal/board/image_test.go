package board

import (
	"image"
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {
	type args struct {
		x      int
		y      int
		length int
	}
	tests := []struct {
		name string
		args args
		want *image.Rectangle
	}{
		{
			name: "base",
			args: args{
				x:      1,
				y:      2,
				length: 3,
			},
			want: &image.Rectangle{
				Min: image.Point{
					X: 1,
					Y: 2,
				},
				Max: image.Point{
					X: 4,
					Y: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Square(tt.args.x, tt.args.y, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDrawBoard(t *testing.T) {
	type args struct {
		board Board
		theme Theme
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DrawBoard(tt.args.board, tt.args.theme)
		})
	}
}
