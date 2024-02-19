package main

import "fmt"

type gameState struct {
	turn   string
	board  [8][8]Space
	pieces [16]Piece
}

func startIt() {
	// Create a new board
	board := Board{
		Spaces: [8][8]Space{},
		Pieces: [16]Piece{},
	}
	state := gameState{
		turn:   "white",
		board:  board.Spaces,
		pieces: board.Pieces,
	}

	fmt.Print(state)
}
