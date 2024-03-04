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

func (p *Piece) generateMoves(place Place) []string {
	moves := []string{}
	fmt.Println(moves)
	row := place.Rank - 1
	// convert col to number
	col := int(place.File - 'a')
	fmt.Println(row, col)

	switch p.Type {
	case Pawn:
		// moves = p.pawnMoves()
	case Knight:
		// moves = p.knightMoves()
	case Bishop:
		// moves = p.bishopMoves()
	case Rook:
		// moves = p.rookMoves()
	case Queen:
		// moves = p.queenMoves()
	case King:
		// moves = p.kingMoves()
	}

	return []string{}
}
