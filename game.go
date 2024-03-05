package main

import "fmt"

type gameState struct {
	board [8][8]Space

	pieceColorBitboards [2]uint64
	pieceTypeBitboards  [6]uint64
	kings               [2]int
	epSquare            int

	castlingVariables uint8 // 0bKQkq

	turn PieceColor

	halfmoves int
	numMoves  int
}

func newGame() gameState {
	// Create a new board
	state := gameState{
		board: [8][8]Space{},

		pieceColorBitboards: [2]uint64{},
		pieceTypeBitboards:  [6]uint64{}, // reference PieceType in pieces.go, but generally 0: pawns, 1: knights, 2: bishops, 3: rooks, 4: queens, 5: kings
		kings:               [2]int{},
		epSquare:            -1,

		castlingVariables: 0,

		turn: White,

		halfmoves: 0,
		numMoves:  0,
	}
	state.pieceColorBitboards[White], state.pieceColorBitboards[Black] = 0, 0
	state.pieceTypeBitboards[0], state.pieceTypeBitboards[1], state.pieceTypeBitboards[2], state.pieceTypeBitboards[3], state.pieceTypeBitboards[4], state.pieceTypeBitboards[5] = 0, 0, 0, 0, 0, 0

	return state
}

func (g *gameState) allBB() uint64 {
	return g.pieceColorBitboards[White] | g.pieceColorBitboards[Black]
}

func (g *gameState) getBitBoards() {
	var wBB, bBB uint64
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if g.board[i][j].Occupied {
				Ptype := g.board[i][j].Piece.Type
				bit := uint64(1 << (i*8 + j))
				fmt.Printf("bit value: %#v", bit)
				if g.board[i][j].Piece.Color == White {

					wBB = wBB | bit
				} else {
					bBB = bBB | bit
				}
				g.pieceTypeBitboards[Ptype] = g.pieceTypeBitboards[Ptype] | bit
			}
		}
	}
	g.pieceColorBitboards[White] = wBB
	g.pieceColorBitboards[Black] = bBB
}
