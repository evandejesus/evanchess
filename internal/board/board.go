package board

import (
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/evandejesus/evanchess/internal/piece"
	chess "github.com/notnil/chess"
	image "github.com/notnil/chess/image"
)

type Board struct {
	squares [64]int
}

func CreateGraphicalBoard(fenStr string) {
	file, err := os.Create("_output/output.svg")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer file.Close()
	pos := &chess.Position{}
	pos.UnmarshalText([]byte(fenStr))
	image.SVG(file, pos.Board())
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
				file += int(char)
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
