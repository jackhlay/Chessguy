package main

type gameState struct {
	turn      string
	board     [8][8]Space
	pieces    [16]Piece
	epSquare  string
	halfmoves int
	numMoves  int

	//Castling variables
	BKCastle bool
	BQCastle bool
	WKCastle bool
	WQCastle bool
}

func startIt() (Board, gameState) {
	// Create a new board
	board := Board{
		Spaces: [8][8]Space{},
		Pieces: [16]Piece{},
	}
	state := gameState{
		turn:      "white",
		board:     board.Spaces,
		pieces:    board.Pieces,
		epSquare:  "",
		halfmoves: 0,
		numMoves:  0,

		BKCastle: false,
		BQCastle: false,
		WKCastle: false,
		WQCastle: false,
	}

	return board, state
}
