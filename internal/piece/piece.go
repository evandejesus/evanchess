package piece

/*
00 000
   ^ 	0 - 6 indicates piece type
^		1 for white, 2 for black
*/

const (
	None int = iota
	King
	Pawn
	Knight
	Bishop
	Rook
	Queen

	White int = 8
	Black int = 16
)

// IsColor returns whether the square is white or black based on position in board array
func IsColor(square int, color int) bool {
	return square&color == color
}
func IsOpponentColor(square int, color int) bool {
	return square != 0 && square&color == 0
}

func GetPieceType(square int) int {
	return square & 7
}

func IsPieceType(square int, pieceType int) bool {
	return square&7 == pieceType
}

func IsSlidingPiece(square int) bool {
	p := GetPieceType(square)
	return p == Queen || p == Rook || p == Bishop
}
