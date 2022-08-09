// Package board implements the board representation for the engine.
package board

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	piece "github.com/evandejesus/evanchess/internal/piece"
)

// a Board represents the location of all the pieces on the chess board using an array of length 64.
// It also stores the possible moves as well as the color to move.
type Board struct {
	Squares     [64]int
	ColorToMove int
	HalfMoves   int
	FullMoves   int
	MoveHistory []Move
}

var pieceTypeFromSymbol = map[rune]int{
	'k': piece.King,
	'r': piece.Rook,
	'b': piece.Bishop,
	'n': piece.Knight,
	'q': piece.Queen,
	'p': piece.Pawn,
}

// LoadPositionFromFen creates a Board object from FEN notation.
func LoadPositionFromFen(fen string) (board Board, err error) {
	fenArray := strings.Split(fen, " ")

	// Determine color to move from FEN
	if color := fenArray[1]; color == "w" {
		board.ColorToMove = piece.White
	} else if color == "b" {
		board.ColorToMove = piece.Black
	}

	board.HalfMoves, err = strconv.Atoi(fenArray[4])
	if err != nil {
		return Board{}, err
	}
	board.FullMoves, err = strconv.Atoi(fenArray[5])
	if err != nil {
		return Board{}, err
	}

	// Populate Board object with pieces
	fenBoard := fenArray[0]
	file := 0
	rank := 7
	for _, char := range fenBoard {
		// Iterate rank by rank
		if char == '/' {
			// Go to beginning of next rank
			file = 0
			rank -= 1
		} else {
			if unicode.IsDigit(char) {
				// Number of empty spaces encountered
				file += int(char - '0')
			} else {
				var pieceColor int
				if unicode.IsUpper(char) {
					pieceColor = piece.White
				} else {
					pieceColor = piece.Black
				}

				pieceType := pieceTypeFromSymbol[unicode.ToLower(char)]
				if pieceType == 0 {
					return board, errors.New("invalid fen character")
				}

				// Generate integer from color and type
				board.Squares[rank*8+file] = pieceType | pieceColor
				file += 1
			}
		}
	}

	return board, nil
}
