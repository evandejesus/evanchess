package board

import (
	"bufio"
	"fmt"
	"os"

	"github.com/evandejesus/evanchess/internal/piece"
	"github.com/evandejesus/evanchess/internal/projectpath"
)

var directionOffsets = []int{8, -8, -1, 1, 7, -7, 9, -9}
var numSquaresToEdge [64][]int

// Generate matrix of available squares in each direction from each starting square
func precomputedMoveData() {

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

// Move is a representation of a chess move, including start, target, type of piece, and whether the move is a capture.
type Move struct {
	startSquare  int
	targetSquare int
	pieceType    int
	// isCapture    bool
}

// GenerateMoves finds each piece of the color to play and adds their possible moves to a list.
// Depends on function to compute all reachable squares from each square.
func GenerateMoves(board *Board) []Move {
	var moves []Move

	precomputedMoveData()

	//
	for i := 0; i < 64; i++ {
		p := board.Squares[i]
		if piece.IsColor(p, board.ColorToMove) {
			if piece.IsSlidingPiece(p) {
				moves = append(moves, generateSlidingMoves(board, i, p)...)
			} else if piece.IsPieceType(p, piece.Pawn) {
				moves = append(moves, generatePawnMoves(board, i)...)
			} else if piece.IsPieceType(p, piece.King) {
				moves = append(moves, generateKingMoves(board, i)...)
			} else if piece.IsPieceType(p, piece.Knight) {
				moves = append(moves, generateKnightMoves(board, i)...)
			}
		}
	}
	return moves
}

func generatePawnMoves(board *Board, startSquare int) []Move {
	var pawnMoves []Move

	// forward white move
	if board.ColorToMove == piece.White && getRank(startSquare) < 7 {
		if getFile(startSquare) > 0 && piece.IsOpponentColor(board.Squares[startSquare+7], board.ColorToMove) {
			// left capture
			pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare + 7, pieceType: piece.Pawn})
		}
		if getFile(startSquare) < 7 && piece.IsOpponentColor(board.Squares[startSquare+9], board.ColorToMove) {
			// right capture
			pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare + 9, pieceType: piece.Pawn})
		}
		if board.Squares[startSquare+8] == 0 {
			if getRank(startSquare) == 6 {
				// promotion
				pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare + 8, pieceType: piece.Queen})

			} else {
				// forward one
				pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare + 8, pieceType: piece.Pawn})
			}
		}
	}

	// forward black move
	if board.ColorToMove == piece.Black && getRank(startSquare) > 0 {
		if getFile(startSquare) > 0 && piece.IsOpponentColor(board.Squares[startSquare-9], board.ColorToMove) {
			// left capture
			pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 9, pieceType: piece.Pawn})
		}
		if getFile(startSquare) < 7 && piece.IsOpponentColor(board.Squares[startSquare-7], board.ColorToMove) {
			// right capture
			pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 7, pieceType: piece.Pawn})
		}
		if board.Squares[startSquare-8] == 0 {

			if getRank(startSquare) == 1 {
				// promotion
				pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 8, pieceType: piece.Queen})
				// pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 8, pieceType: piece.Rook})
				// pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 8, pieceType: piece.Bishop})
				// pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 8, pieceType: piece.Knight})

			} else {
				// forward one
				pawnMoves = append(pawnMoves, Move{startSquare: startSquare, targetSquare: startSquare - 8, pieceType: piece.Pawn})
			}

		}
	}

	return pawnMoves
}

func generateSlidingMoves(board *Board, startSquare, p int) []Move {
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
			slidingMoves = append(slidingMoves, Move{startSquare: startSquare, targetSquare: targetSquare, pieceType: piece.GetPieceType(p)})

			// stop looking if enemy piece is in the way
			if piece.IsOpponentColor(pieceOnTargetSquare, board.ColorToMove) {
				break
			}
		}
	}
	return slidingMoves
}

func generateKingMoves(board *Board, startSquare int) []Move {
	var kingMoves []Move
	for directionIndex := 0; directionIndex < 8; directionIndex++ {
		if numSquaresToEdge[startSquare][directionIndex] > 0 {
			targetSquare := startSquare + directionOffsets[directionIndex]
			pieceOnTargetSquare := board.Squares[targetSquare]
			if !piece.IsColor(pieceOnTargetSquare, board.ColorToMove) {
				kingMoves = append(kingMoves, Move{startSquare: startSquare, targetSquare: targetSquare, pieceType: piece.King})
			}
		}
	}
	return kingMoves
}

func generateKnightMoves(board *Board, startSquare int) []Move {
	var knightMoves []Move

	var targetSquares []int
	if getFile(startSquare) >= 1 {
		targetSquares = append(targetSquares, startSquare+15)
		targetSquares = append(targetSquares, startSquare-17)
	}
	if getFile(startSquare) >= 2 {
		targetSquares = append(targetSquares, startSquare+6)
		targetSquares = append(targetSquares, startSquare-10)
	}
	if getFile(startSquare) <= 5 {
		targetSquares = append(targetSquares, startSquare+10)
		targetSquares = append(targetSquares, startSquare-6)
	}
	if getFile(startSquare) <= 6 {
		targetSquares = append(targetSquares, startSquare+17)
		targetSquares = append(targetSquares, startSquare-15)
	}

	for _, t := range targetSquares {
		if t > 0 && t < 64 && !piece.IsColor(board.Squares[t], board.ColorToMove) {
			knightMoves = append(knightMoves, Move{startSquare: startSquare, targetSquare: t, pieceType: piece.Knight})

		}
	}
	return knightMoves
}

// MakeMove updates the board position with the provided move and sets the colorToMove to the opposite color.
func MakeMove(move Move, b *Board) {

	// move piece to new square, replace existing piece
	b.Squares[move.startSquare] = 0
	b.Squares[move.targetSquare] = move.pieceType | b.ColorToMove

	// switch sides
	if b.ColorToMove == piece.White {
		b.ColorToMove = piece.Black
	} else {
		b.ColorToMove = piece.White
		b.FullMoves += 1
	}
	b.HalfMoves += 1
}

// PrintMoves prints each move in moves to a file with the format <target>-<destination>.
// Expects an existing directory _output
func PrintMoves(moves []Move) {
	f, err := os.Create(fmt.Sprintf("%s/_output/moves.log", projectpath.Root))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, move := range moves {
		str := printMove(move)
		w.WriteString(str)
	}
	w.Flush()
}

func printMove(move Move) string {
	return fmt.Sprintf("%c%d-%c%d\n", 97+move.startSquare%8, 1+move.startSquare/8,
		97+move.targetSquare%8, 1+move.targetSquare/8)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getFile(square int) int {
	return square % 8
}

func getRank(square int) int {
	return square / 8
}
