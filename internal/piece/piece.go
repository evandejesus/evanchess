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
