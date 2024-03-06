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

func getRowCol(s string) (int, int) {
	ind := placeToInd[s]
	col := ind % 8
	row := ind / 8
	return row, col

}

func (g *gameState) pieceAt(s string) bool {
	return g.board[8-int(s[1]-'0')][int(s[0]-'a')].Piece != nil

}

func (g *gameState) clearOldSq(s string) {
	row, col := getRowCol(s)
	g.board[row][col].Piece = nil
	g.board[row][col].Occupied = false

	fmt.Println()
	for _, row := range g.board {
		for _, col := range row {
			if col.Occupied {
				fmt.Printf("%c ", col.Piece.Symbol)
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}

func (g *gameState) getPieceAt(s string) *Piece {
	row, col := getRowCol(s)
	return g.board[row][col].Piece
}

func (g *gameState) setPieceAt(s string, piece *Piece) {
	row, col := getRowCol(s)
	g.board[row][col].Piece = piece
	g.board[row][col].Occupied = true
}

func (g *gameState) makeMove(move string) {
	// move := "e2e4"
	from := move[:2]
	to := move[2:]
	piece := g.getPieceAt(from)
	g.clearOldSq(from)
	g.setPieceAt(to, piece)
}
