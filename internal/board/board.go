package board

import (
	"strings"
	"unicode"

	piece "github.com/evandejesus/evanchess/internal/piece"
)

type Board struct {
	squares [64]int
}

func LoadPositionFromFen(fen string) (board Board) {
	pieceTypeFromSymbol := map[rune]int{
		'k': piece.King,
		'r': piece.Rook,
		'b': piece.Bishop,
		'n': piece.Knight,
		'q': piece.Queen,
		'p': piece.Pawn,
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
				board.squares[rank*8+file] = pieceType | pieceColor
				file += 1
			}
		}
	}
	return board
}

func CreateBoard() {
	var board Board
	board.squares[0] = piece.White | piece.King
}

func IntToPiece(i int) (isWhite bool, pieceType int) {
	isWhite = i>>3 == 1
	pieceType = i & 7

	return isWhite, pieceType
}
