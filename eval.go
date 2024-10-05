package main

import (
	"github.com/notnil/chess"
)

func evalPos(position *chess.Position, game *chess.Game) float64 {
	whitePieces := []chess.Piece{}
	whiteTargeted := []chess.Square{}
	whitePos := []chess.Square{}
	BlackPieces := []chess.Piece{}
	blackTargeted := []chess.Square{}
	blackPos := []chess.Square{}

	//material eval

	pieceValue := map[chess.PieceType]float64{
		chess.Pawn:   1.0,
		chess.Knight: 3.0,
		chess.Bishop: 3.1,
		chess.Rook:   5,
		chess.Queen:  9,
		chess.King:   0,
	}

	pieceTot := 0.0
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			value := pieceValue[piece.Type()]

			if piece.Color() == chess.White {
				whitePieces = append(whitePieces, piece)
				whitePos = append(whitePos, sq)
				pieceTot += .3 * value
			} else {
				BlackPieces = append(BlackPieces, piece)
				blackPos = append(blackPos, sq)
				pieceTot -= .3 * value
			}
		}

	}

	//fraction of total pieces
	wLeft := len(whitePieces) / 16
	bLeft := len(BlackPieces) / -16

	//attacking pieces
	for _, move := range position.ValidMoves() {
		targetsq := move.S2()
		for _, ws := range blackPos {
			if targetsq == ws {
				blackTargeted = append(blackTargeted, ws)
			}
		}
	}
	wAttacks := len(blackTargeted)

	for _, move := range position.ValidMoves() {
		target := move.S2()
		for _, bs := range whitePos {
			if target == bs {
				whiteTargeted = append(whiteTargeted, bs)
			}
		}
	}
	bAttacks := len(whiteTargeted)
	attPot := .3 * float64(wAttacks+bAttacks)

	eval := (.55 * float64(pieceTot)) + (.3 * float64(attPot)) + (.05 * float64(wLeft)) - (.05 * float64(bLeft))
	return eval
}
