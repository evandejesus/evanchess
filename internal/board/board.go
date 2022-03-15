package board

import (
	"errors"
	"strings"
	"unicode"

	piece "github.com/evandejesus/evanchess/internal/piece"
)

type Board struct {
	Squares     [64]int
	moves       []Move
	ColorToMove int
}

func LoadPositionFromFen(fen string) (board Board, err error) {
	pieceTypeFromSymbol := map[rune]int{
		'k': piece.King,
		'r': piece.Rook,
		'b': piece.Bishop,
		'n': piece.Knight,
		'q': piece.Queen,
		'p': piece.Pawn,
	}

	if color := strings.Split(fen, " ")[1]; color == "w" {
		board.ColorToMove = piece.White
	} else if color == "b" {
		board.ColorToMove = piece.Black
	}

	fenBoard := strings.Split(fen, " ")[0]
	file := 0
	rank := 7
	for _, char := range fenBoard {
		if char == '/' {
			file = 0
			rank -= 1
		} else {
			if unicode.IsDigit(char) {
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
				board.Squares[rank*8+file] = pieceType | pieceColor
				file += 1
			}
		}
	}
	GenerateMoves(&board)

	return board, nil
}
