package board

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func Square(x, y, length int) *image.Rectangle {
	square := image.Rect(x, y, x+length, y+length)
	return &square
}

func DrawBoard(theme Theme, size int) {
	board := image.NewRGBA(image.Rect(0, 0, 8*size, 8*size))

	// rows
	for i := 0; i < 8; i++ {
		// columns
		for j := 0; j < 8; j++ {
			square := Square(size*j, size*i, size)
			var bg color.Color
			if (i+j)%2 == 0 {
				bg = theme.light
			} else {
				bg = theme.dark
			}
			draw.Draw(board, square.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
		}
	}
	f, err := os.Create("_output/board.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, board)
}
