package main

import (
	"fmt"
	"math/rand"
	"os"

	"fyne.io/fyne/v2"
	"github.com/evandejesus/evanchess/internal/board"
)

var keys []*fyne.KeyEvent

func setup() {

	// create output directory for moves log
	os.Mkdir("_output", os.ModePerm)
}

func main() {
	setup()

	fmt.Println("evanchess c. 2022")
	fens := []string{
		/* starting position */
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		/* e4 e5 */
		"rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
		/* 16/32 */
		"8/2p1N1R1/1Pn1p3/1b6/8/2p1B2K/pqP1k3/1Qb5 b - - 1 1",
		/* 27/32 */
		"3r4/PPk4p/r4p1P/bP6/3pRK2/np1pp3/p1qn1PBP/Nb1NR3 w - - 0 1",
		/* study */
		"8/8/7p/3KNN1k/2p4p/8/3P2p1/8 w - - 0 1",
		/* scandinavian */
		"rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR w KQkq d6 0 2",
		/* random pawns */
		"rnbqkbnr/p1p1pppp/8/1p1p4/P3P3/8/1PPP1PPP/RNBQKBNR w KQkq b6 0 3",
	}
	fen := fens[3]

	fmt.Println("loading position from FEN...")
	pos, err := board.LoadPositionFromFen(fen)
	if err != nil {
		panic(err)
	}

	// Generate first ply moves
	fmt.Println("generating moves...")
	moves := board.GenerateMoves(&pos)
	board.PrintMoves(moves)
	fmt.Println("moves generated")

	// make random move
	board.MakeMove(moves[rand.Int()%len(moves)], &pos)

	// Draw Board
	if w, err := board.DrawBoard(pos, board.Dusk); err != nil {
		panic(err)
	} else {
		w.Canvas().SetOnTypedKey(keyHandler)
		w.ShowAndRun()
	}
}

// keyHandler attaches functionality to the fyne board window
func keyHandler(k *fyne.KeyEvent) {
	keys = append(keys, k)
	if k.Name == "Return" {
		for _, key := range keys[:len(keys)-1] {
			fmt.Print(key.Name)
		}
		fmt.Println("")

		keys = nil
	}
}
