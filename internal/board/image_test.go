package board

import (
	"errors"
	"image"
	"os"
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

func setup(t testing.TB) func(t testing.TB) {
	if _, err := os.Stat("/_output/board.png"); err == nil {
		if err := os.Remove("_output/board.png"); err != nil {
			t.Error(err.Error())
		}
	}

	return func(t testing.TB) {
		// teardown
	}
}

func TestDrawBoard(t *testing.T) {
	type args struct {
		board Board
		theme Theme
	}
	pos, _ := LoadPositionFromFen("5rk1/4R1pp/3q1p2/p1p2P2/P3Q2P/5p2/2P2PPK/8 w - - 0 34")
	tests := []struct {
		name string
		args args
	}{
		{
			name: "random fen",
			args: args{
				board: pos,
				theme: Sandcastle,
			},
		},
	}

	teardown := setup(t)
	defer teardown(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DrawBoard(tt.args.board, tt.args.theme)
		})
	}

	// board image successfully created
	t.Run("created file", func(t *testing.T) {
		if _, err := os.Stat(OutputFilepath()); errors.Is(err, os.ErrNotExist) {
			t.Error(err.Error())
			t.Errorf("file \"" + OutputFilepath() + "\" not created")
		}
	})
}

func Test_getFilepathFromInt(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name              string
		args              args
		wantPieceFilepath string
	}{
		{
			name: "none",
			args: args{
				val: 0,
			},
			wantPieceFilepath: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPieceFilepath := getFilepathFromInt(tt.args.val); gotPieceFilepath != tt.wantPieceFilepath {
				t.Errorf("getFilepathFromInt() = %v, want %v", gotPieceFilepath, tt.wantPieceFilepath)
			}
		})
	}

}
