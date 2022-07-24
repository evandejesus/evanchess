package piece

// Package piece implements the system for storing piece color and type as a single integer.
// Pieces are stored using the most significant two bits for the piece color and
// the least significant two bits for the piece type.

// 00 000
//    ^ 	0 - 6 indicates piece type
// ^		1 for white, 2 for black

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

// GetPieceType returns whether piece is knight, bishop, etc.
func GetPieceType(square int) int {
	return square & 7
}

// IsPieceType returns whether the piece at `square` is the piece type being compared to
func IsPieceType(square int, pieceType int) bool {
	return square&7 == pieceType
}

// IsSlidingPiece returns whether the piece is a queen, rook, or bishop
func IsSlidingPiece(square int) bool {
	p := GetPieceType(square)
	return p == Queen || p == Rook || p == Bishop
}
