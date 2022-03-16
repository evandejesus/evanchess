package board

import (
	"reflect"
	"testing"

	"github.com/evandejesus/evanchess/internal/piece"
)

func TestLoadPositionFromFen(t *testing.T) {
	type args struct {
		fen string
	}
	tests := []struct {
		name      string
		args      args
		wantBoard Board
		wantErr   bool
	}{
		{
			name: "random",
			args: args{
				fen: "8/2p1N1R1/1Pn1p3/1b6/8/2p1B2K/pqP1k3/1Qb5 b - - 1 1",
			},
			wantBoard: Board{
				Squares: [64]int{
					0, 14, 20, 0, 0, 0, 0, 0, 18, 22, 10, 0, 17, 0, 0, 0, 0, 0, 18, 0, 12, 0, 0, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 10, 19, 0, 18, 0, 0, 0, 0, 0, 18, 0, 11, 0, 13, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				},
				Moves: []Move{
					{startSquare: 2, targetSquare: 11}, {startSquare: 2, targetSquare: 20}, {startSquare: 9, targetSquare: 17}, {startSquare: 9, targetSquare: 25}, {startSquare: 9, targetSquare: 1}, {startSquare: 9, targetSquare: 10}, {startSquare: 9, targetSquare: 16}, {startSquare: 9, targetSquare: 0}, {startSquare: 33, targetSquare: 40}, {startSquare: 33, targetSquare: 26}, {startSquare: 33, targetSquare: 19}, {startSquare: 33, targetSquare: 24},
				},
				ColorToMove: piece.Black,
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				fen: "8/2x1N1R1/1Pn1p3/1b6/8/2p1B2K/pqP1k3/1Qb5 b - - 1 1",
			},
			wantBoard: Board{ColorToMove: piece.Black},
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoard, err := LoadPositionFromFen(tt.args.fen)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPositionFromFen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoard, tt.wantBoard) {
				t.Errorf("LoadPositionFromFen() = %v, want %v", gotBoard, tt.wantBoard)
			}
		})
	}
}
