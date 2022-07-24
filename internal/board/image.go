package board

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/evandejesus/evanchess/internal/piece"
	"github.com/evandejesus/evanchess/internal/projectpath"
)

var squareSize int = 60
var boardSize float32 = 480

func Square(x, y, length int) *image.Rectangle {
	square := image.Rect(x, y, x+length, y+length)
	return &square
}

func DrawBoard(board Board, theme Theme) (fyne.Window, error) {
	boardPng := image.NewRGBA(image.Rect(0, 0, 8*squareSize, 8*squareSize))

	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			square := Square(squareSize*rank, squareSize*file, squareSize)
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
				return nil, err
			}
			defer pieceFile.Close()

			piece, _, err := image.Decode(pieceFile)
			if err != nil {
				return nil, err
			}
			draw.Draw(boardPng, square.Bounds(), piece, image.Point{}, draw.Over)
		}
	}
	a := app.New()
	w := a.NewWindow("Images")
	img := canvas.NewImageFromImage(boardPng)
	w.SetContent(img)
	w.Resize(fyne.NewSize(boardSize, boardSize))

	return w, nil
}

func pieceFilepathFromSquare(square int) (pieceFilepath string) {
	// return the location of the piece image file from numerical value

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

// OutputFilepath returns the absolute file location of the board image
func OutputFilepath() string {
	return fmt.Sprintf("%s/_output/board.png", projectpath.Root)
}
