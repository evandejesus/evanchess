package board

import (
	"github.com/evandejesus/evanchess/internal/piece"
)

var directionOffsets = []int{8, -8, -1, 1, 7, -7, 9, -9}
var numSquaresToEdge [64][]int

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
				min(numSouth, numEast),
				min(numNorth, numEast),
				min(numSouth, numWest),
			}
		}
	}
}

type Move struct {
	startSquare  int
	targetSquare int
}

func GenerateMoves(board *Board) {
	PrecomputedMoveData()
	for i := 0; i < 64; i++ {
		p := board.Squares[i]
		if piece.IsColor(p, board.ColorToMove) {
			if piece.IsSlidingPiece(p) {
				generateSlidingMoves(board, i, p)
			} else if piece.IsPieceType(p, piece.Pawn) {

			}
		}
	}
}

func generateSlidingMoves(board *Board, startSquare, p int) {
	var slidingMoves []Move

	var startDirIndex, endDirIndex int
	if piece.IsPieceType(p, piece.Bishop) {
		startDirIndex = 4
	} else {
		startDirIndex = 0
	}

	if piece.IsPieceType(p, piece.Rook) {
		endDirIndex = 4
	} else {
		endDirIndex = 8
	}

	// for each possible direction
	for directionIndex := startDirIndex; directionIndex < endDirIndex; directionIndex++ {
		// for each square in this direction
		for n := 0; n < numSquaresToEdge[startSquare][directionIndex]; n++ {
			targetSquare := startSquare + directionOffsets[directionIndex]*(n+1)
			pieceOnTargetSquare := board.Squares[targetSquare]
			// stop looking if friendly piece is in the way
			if piece.IsColor(pieceOnTargetSquare, board.ColorToMove) {
				break
			}
			// moves = append(moves, Move{startSquare: startSquare, targetSquare: targetSquare})
			slidingMoves = append(slidingMoves, Move{startSquare: startSquare, targetSquare: targetSquare})

			// stop looking if enemy piece is in the way
			if piece.IsOpponentColor(pieceOnTargetSquare, board.ColorToMove) {
				break
			}
		}
	}
	board.Moves = append(board.Moves, slidingMoves...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
