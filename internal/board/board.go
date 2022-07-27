// Package board implements the board representation for the engine.
package board

import (
	"errors"
	"strings"
	"unicode"

	piece "github.com/evandejesus/evanchess/internal/piece"
)

// a Board represents the location of all the pieces on the chess board using an array of length 64.
// It also stores the possible moves as well as the color to move.
type Board struct {
	Squares     [64]int
	ColorToMove int
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

	// Determine color to move from FEN
	if color := strings.Split(fen, " ")[1]; color == "w" {
		board.ColorToMove = piece.White
	} else if color == "b" {
		board.ColorToMove = piece.Black
	}

	// Populate Board object with pieces
	fenBoard := strings.Split(fen, " ")[0]
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
