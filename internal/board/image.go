package board

import (
	"fmt"
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

func DrawBoard(board Board, theme Theme) error {
	boardPng := image.NewRGBA(image.Rect(0, 0, 8*size, 8*size))

	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			square := Square(size*rank, size*file, size)
			var bg color.Color
			if (file+rank)%2 == 0 {
				bg = theme.light
			} else {
				bg = theme.dark
			}
			draw.Draw(boardPng, square.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)

			// render board in reverse order
			pieceVal := board.Squares[63-(8*file+(7-rank))]
			if pieceVal == 0 {
				continue
			}
			pieceFilepath := pieceFilepathFromSquare(pieceVal)
			pieceFile, err := os.Open(pieceFilepath)
			if err != nil {
				return err
			}
			defer pieceFile.Close()

			piece, _, err := image.Decode(pieceFile)
			if err != nil {
				return err
			}
			draw.Draw(boardPng, square.Bounds(), piece, image.Point{}, draw.Over)
		}
	}
	f, err := os.Create(OutputFilepath())
	if err != nil {
		return err
	}
	defer f.Close()
	png.Encode(f, boardPng)

	return nil
}

func pieceFilepathFromSquare(square int) (pieceFilepath string) {
	isWhite := piece.IsColor(square, piece.White)
	pieceType := piece.GetPieceType(square)

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

	return fmt.Sprintf("%s/assets/%s", projectpath.Root, pieceFilepath)
}

func OutputFilepath() string {
	return fmt.Sprintf("%s/_output/board.png", projectpath.Root)
}
