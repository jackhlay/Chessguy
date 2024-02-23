package main

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
