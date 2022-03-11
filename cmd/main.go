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
	fmt.Println("evanchess c. 2020")
	// fenStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

	board.DrawBoard(board.Emerald, 60)
}
