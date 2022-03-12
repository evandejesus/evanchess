package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/evandejesus/evanchess/internal/board"
)

const path = "_output"

func init() {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	fmt.Println("evanchess c. 2022")
	fenStr := "5rk1/4R1pp/3q1p2/p1p2P2/P3Q2P/5p2/2P2PPK/8 w - - 0 34"

	pos := board.LoadPositionFromFen(fenStr)
	// pos.PrintBoard()
	board.DrawBoard(pos, board.Sandcastle)
}
