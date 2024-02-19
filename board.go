package main

type Board struct {
	Spaces [8][8]Space
	Pieces [16]Piece
}

type Place struct {
	File rune
	Rank int
}

type Space struct {
	Occupied bool
	Place    Place
}
