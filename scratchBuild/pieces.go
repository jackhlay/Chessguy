package main

import "fmt"

type PieceType int
type PieceColor int

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	White PieceColor = iota
	Black
)

type Piece struct {
	Type   PieceType
	Symbol rune
	Color  PieceColor
	Value  int
	Moved  bool
}

func (p *Piece) generateMoves(place string, g gameState) uint64 {
	moves := uint64(0)
	row := place[1] - 1
	col := int(place[0] - 'a')
	fmt.Println(row, col)

	switch p.Type {
	case Pawn:
		// moves = p.pawnMoves()
	case Knight:
		// moves = p.knightMoves()
	case Bishop:
		// moves = p.bishopMoves()
	case Rook:
		moves = p.rookMoves(g)
	case Queen:
		// moves = p.queenMoves()
	case King:
		// moves = p.kingMoves()
	}
	return moves
}

func (p *Piece) rookMoves(g gameState) uint64 {
	board := g.board
	g.getBitBoards()
	occupancyBB := (g.pieceColorBitboards[White] | g.pieceColorBitboards[Black])
	fmt.Println("Occupancy Board:")
	printBoard(occupancyBB)
	row, col := p.getPieceLocation(board)
	ind := row*8 + col
	moves := p.rookGen(ind, occupancyBB, g)
	fmt.Println(moves)
	printBoard(moves)
	return moves

}

func (p *Piece) getPieceLocation(board [8][8]Space) (int, int) {
	for i, row := range board {
		for j, col := range row {
			if p == col.Piece {
				return i, j
			}
		}
	}
	return -1, -1
}

func (p *Piece) rookGen(square int, occupancyBB uint64, g gameState) uint64 {
	attacksBB := uint64(0)
	Rcol, Rrow := p.getPieceLocation(g.board)
	fmt.Println(Rrow, " ", Rcol)

	for col := Rcol + 1; col < 8; col++ {
		fmt.Println("DOWN")

		inda := 8*col + Rrow
		fmt.Println(inda)
		if g.board[Rrow][col].Occupied {
			attacksBB |= uint64(1) << inda
			break
		}
		attacksBB |= uint64(1) << inda
	}
	for coll := Rcol - 1; coll >= 0; coll-- {
		fmt.Println("UP1")
		indb := 8*coll + Rrow
		fmt.Println(indb)
		if g.board[Rrow][coll].Occupied {
			attacksBB |= uint64(1) << indb
			break
		}
		attacksBB |= uint64(1) << indb
	}
	for row := Rrow + 1; row < 8; row++ {
		fmt.Println("LEFT")
		indc := 8*Rcol + row
		fmt.Println(indc)
		if squareOccupied(occupancyBB, indc) {
			attacksBB |= uint64(1) << indc
			break
		}
		attacksBB |= uint64(1) << indc
	}
	for roww := Rrow - 1; roww >= 0; roww-- {
		fmt.Println("RIGHT")
		indd := 8*Rcol + roww
		fmt.Println(indd)
		if squareOccupied(occupancyBB, indd) {
			fmt.Println("OCCUPIED")
			attacksBB |= uint64(1) << indd
			fmt.Println("Breaks here")
			break
		}
		attacksBB |= uint64(1) << indd
	}

	return attacksBB
}

func squareOccupied(occupancyBB uint64, square int) bool {
	sq := uint64(1) << uint(square)
	return (occupancyBB & sq) != 0
}
