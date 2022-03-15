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
	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("evanchess c. 2022")

	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	// fen := "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2"
	fen := "8/2p1N1R1/1Pn1p3/1b6/8/2p1B2K/pqP1k3/1Qb5 b - - 1 1"

	pos, err := board.LoadPositionFromFen(fen)
	if err != nil {
		panic(err)
	}
	if err = board.DrawBoard(pos, board.Sandcastle); err != nil {
		panic(err)
	} else {
		// board.GenerateMoves(&pos)
		fmt.Println("board generated")
	}

}
