package main

type gameState struct {
	board  [8][8]Space
	pieces [32]Piece

	pieceColorBitboards [2]uint64
	pieceTypeBitboards  [6]uint64
	kings               [2]int
	epSquare            int

	castlingVariables uint8 // 0bKQkq

	turn PieceColor

	halfmoves int
	numMoves  int
}

func newGame() (Board, gameState) {
	// Create a new board
	Board := Board{
		Spaces: [8][8]Space{},
		Pieces: [32]Piece{},
	}
	state := gameState{
		board:  Board.Spaces,
		pieces: Board.Pieces,

		pieceColorBitboards: [2]uint64{},
		pieceTypeBitboards:  [6]uint64{},
		kings:               [2]int{},
		epSquare:            -1,

		castlingVariables: 0,

		turn: White,

		halfmoves: 0,
		numMoves:  0,
	}
	state.pieceColorBitboards[White], state.pieceColorBitboards[Black] = 0, 0
	state.pieceTypeBitboards[0], state.pieceTypeBitboards[1], state.pieceTypeBitboards[2], state.pieceTypeBitboards[3], state.pieceTypeBitboards[4], state.pieceTypeBitboards[5] = 0, 0, 0, 0, 0, 0

	return Board, state
}

func (g *gameState) allBB() uint64 {
	return g.pieceColorBitboards[White] | g.pieceColorBitboards[Black]
}

func (g *gameState) getBitBoards() {
	wBB, bBB := uint64(0), uint64(0)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g.board[i][j].Occupied {
				if g.board[i][j].Piece.Color == White {
					wBB |= 1 << (i*8 + j)
				} else {
					bBB |= 1 << (i*8 + j)
				}
			}
		}
	}
	g.pieceColorBitboards[White] = wBB
	g.pieceColorBitboards[Black] = bBB

}
