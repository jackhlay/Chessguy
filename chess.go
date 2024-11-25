package main

import (
	"github.com/notnil/chess"
)

var games = 0
var maxGames = 3

func main() {
	// This will be fairly boring, since it's mostly just an api for the engine
	if games < maxGames {
		// do a seek
		// wait
	}
}

func game() {
	games++
	game := chess.NewGame()
	moves := game.ValidMoves()
	println(game.Position().Board().Draw())
	for move := range moves {
		println(moves[move].String())

	}
	games--
}
