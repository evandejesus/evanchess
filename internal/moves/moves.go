package moves

import (
	"fmt"

	"github.com/evandejesus/evanchess/internal/board"
	"github.com/evandejesus/evanchess/internal/piece"
)

// var directionOffsets = []int{8, -8, -1, 1, 7, -7, 9, -9}
var numSquaresToEdge [][]int

func PrecomputedMoveData() {
	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			numNorth := 7 - rank
			numSouth := rank
			numEast := 7 - file
			numWest := file
			squareIndex := rank*8 + file

			numSquaresToEdge[squareIndex] = []int{
				numNorth,
				numSouth,
				numWest,
				numEast,
				min(numNorth, numWest),
				min(numNorth, numWest),
				min(numNorth, numWest),
				min(numNorth, numWest),
			}
		}
	}
}

type Move struct {
	startSquare  int
	targetSquare int
}

var moves []Move

func GenerateMoves(board board.Board) {
	for i := 0; i < 64; i++ {
		p := board.Squares[i]
		if piece.IsColor(p, board.ColorToMove) {
			fmt.Println(board.ColorToMove == piece.White)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
