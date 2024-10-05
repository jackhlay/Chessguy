package main

import (
	"github.com/notnil/chess"
)

var games = 0

func main() {
	games := 0
	game()
	if games == 0 {

		// play game against self
	}

}

func game() {
	if games >= 3 {
		return
	}
	games++
	game := chess.NewGame() //new game for now, later will see if we can load from PGN
	moves := game.ValidMoves()
	for move := range moves {
		println(moves[move].String())
	}

	games--
}
