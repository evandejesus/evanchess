package board

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/evandejesus/evanchess/internal/piece"
	"github.com/evandejesus/evanchess/internal/projectpath"
)

var size int = 60

func Square(x, y, length int) *image.Rectangle {
	square := image.Rect(x, y, x+length, y+length)
	return &square
}

func DrawBoard(board Board, theme Theme) {
	boardPng := image.NewRGBA(image.Rect(0, 0, 8*size, 8*size))

	// ranks
	for i := 0; i < 8; i++ {
		// files
		for j := 0; j < 8; j++ {
			square := Square(size*j, size*i, size)
			var bg color.Color
			if (i+j)%2 == 0 {
				bg = theme.light
			} else {
				bg = theme.dark
			}
			draw.Draw(boardPng, square.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)

			// render board in reverse order
			pieceVal := board.squares[63-(8*i+(7-j))]
			if pieceVal == 0 {
				continue
			}
			pieceFilepath := projectpath.Root + "/assets/" + getFilepathFromInt(pieceVal)
			pieceFile, err := os.Open(pieceFilepath)
			if err != nil {
				panic(err)
			}
			defer pieceFile.Close()

			piece, _, err := image.Decode(pieceFile)
			if err != nil {
				panic(err)
			}
			draw.Draw(boardPng, square.Bounds(), piece, image.Point{}, draw.Over)
		}
	}
	f, err := os.Create(OutputFilepath())
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, boardPng)
}

func getFilepathFromInt(val int) (pieceFilepath string) {
	isWhite := val>>3 == 1
	pieceType := val & 7

	if pieceType == 0 {
		pieceFilepath = ""
	} else if pieceType == piece.King && isWhite {
		pieceFilepath = "Chess_klt60.png"
	} else if pieceType == piece.King && !isWhite {
		pieceFilepath = "Chess_kdt60.png"
	} else if pieceType == piece.Knight && isWhite {
		pieceFilepath = "Chess_nlt60.png"
	} else if pieceType == piece.Knight && !isWhite {
		pieceFilepath = "Chess_ndt60.png"
	} else if pieceType == piece.Queen && isWhite {
		pieceFilepath = "Chess_qlt60.png"
	} else if pieceType == piece.Queen && !isWhite {
		pieceFilepath = "Chess_qdt60.png"
	} else if pieceType == piece.Rook && isWhite {
		pieceFilepath = "Chess_rlt60.png"
	} else if pieceType == piece.Rook && !isWhite {
		pieceFilepath = "Chess_rdt60.png"
	} else if pieceType == piece.Pawn && isWhite {
		pieceFilepath = "Chess_plt60.png"
	} else if pieceType == piece.Pawn && !isWhite {
		pieceFilepath = "Chess_pdt60.png"
	} else if pieceType == piece.Bishop && isWhite {
		pieceFilepath = "Chess_blt60.png"
	} else if pieceType == piece.Bishop && !isWhite {
		pieceFilepath = "Chess_bdt60.png"
	}

	return pieceFilepath
}

func OutputFilepath() string {
	return projectpath.Root + "/_output/board.png"
}
