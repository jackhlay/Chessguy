package main

import (
	"strings"

	"github.com/notnil/chess"
)

//Piece Square Tables:

var pawnTable = [64]float64{
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.5, 1.0, 1.0, -2.0, -2.0, 1.0, 1.0, 0.5,
	0.5, 0.5, 0.5, 1.5, 1.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, 0.5,
	0.5, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 1.5, 1.5, 0.5, 0.5, 0.5,
	0.5, 1.0, 1.0, -2.0, -2.0, 1.0, 1.0, 0.5,
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
}

var knightTable = [64]float64{
	-5.0, -4.0, -3.0, -3.0, -3.0, -3.0, -4.0, -5.0,
	-4.0, -2.0, 0.0, 0.5, 0.5, 0.0, -2.0, -4.0,
	-3.0, 0.5, 1.0, 1.5, 1.5, 1.0, 0.5, -3.0,
	-3.0, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, -3.0,
	-3.0, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, -3.0,
	-3.0, 0.5, 1.0, 1.5, 1.5, 1.0, 0.5, -3.0,
	-4.0, -2.0, 0.0, 0.5, 0.5, 0.0, -2.0, -4.0,
	-5.0, -4.0, -3.0, -3.0, -3.0, -3.0, -4.0, -5.0,
}

var bishopTable = [64]float64{
	-2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -2.0,
	-1.0, 0.5, 0.0, 0.0, 0.0, 0.0, 0.5, -1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0,
	-1.0, 1.0, 1.5, 2.0, 2.0, 1.5, 1.0, -1.0,
	-1.0, 1.0, 1.5, 2.0, 2.0, 1.5, 1.0, -1.0,
	-1.0, 1.0, 1.0, 1.5, 1.5, 1.0, 1.0, -1.0,
	-1.0, 0.5, 0.0, 0.0, 0.0, 0.0, 0.5, -1.0,
	-2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -2.0,
}

var rookTable = [64]float64{
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	0.0, 0.0, 0.0, 0.5, 0.5, 0.0, 0.0, 0.0,
}

var queenTable = [64]float64{
	-2.0, -1.0, -1.0, -0.5, -0.5, -1.0, -1.0, -2.0,
	-1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -1.0,
	-1.0, 0.0, 0.5, 0.5, 0.5, 0.5, 0.0, -1.0,
	-0.5, 0.0, 0.5, 0.5, 0.5, 0.5, 0.0, -0.5,
	-0.5, 0.0, 0.5, 0.5, 0.5, 0.5, 0.0, -0.5,
	-1.0, 0.0, 0.5, 0.5, 0.5, 0.5, 0.0, -1.0,
	-1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -1.0,
	-2.0, -1.0, -1.0, -0.5, -0.5, -1.0, -1.0, -2.0,
}

func getPosModifier(piece chess.Piece, sq chess.Square) float64 {
	value := 0.0

	switch piece.Type() {
	case chess.Pawn:
		value += pawnTable[int(sq)]

	case chess.Knight:
		value += knightTable[int(sq)]

	case chess.Bishop:
		value += bishopTable[int(sq)]

	case chess.Rook:
		value += rookTable[int(sq)]

	case chess.Queen:
		value += queenTable[int(sq)]
	}
	if piece.Color() == chess.Black {
		return -1 * value
	}
	return value
}

func mobility(position chess.Position) float64 {
	wFirstMoves := map[chess.Square]bool{}
	bFirstMoves := map[chess.Square]bool{}
	moves := position.ValidMoves()
	for _, move := range moves {
		if !wFirstMoves[move.S1()] {
			wFirstMoves[move.S1()] = true
		}
	}
	whiteMob := float64(len(wFirstMoves) + len(moves))

	fen := position.String()

	if strings.Contains(fen, "w") {
		strings.Replace(fen, "w", "b", 1)
	} else {
		strings.Replace(fen, "b", "w", 1)
	}
	fenVal, _ := chess.FEN(fen)

	otherside := chess.NewGame(fenVal)
	otherMoves := otherside.Position().ValidMoves()
	for _, move := range otherMoves {
		if !bFirstMoves[move.S1()] {
			bFirstMoves[move.S1()] = true
		}

	}
	blackMob := float64(len(bFirstMoves) + len(otherMoves))

	return whiteMob - blackMob
}

//TODO: Write check for exposed king

//TODO: Implement recursive search for live brute force eval.
//TODO: Keep search as a dfs, and implement alpha beta pruning to help with optimization
//Additionally, before running the ABP, run a "Plastic-Bag check" to see if any moves are obviously bad and shouldn't be considered

func evalPos(position chess.Position, game chess.Game) float64 {
	whitePieces := []chess.Piece{}
	whiteTargeted := []chess.Square{}
	whitePos := []chess.Square{}
	BlackPieces := []chess.Piece{}
	blackTargeted := []chess.Square{}
	blackPos := []chess.Square{}

	tempo := 0.0

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
				pieceTot += .3 * (value + getPosModifier(piece, sq))
			} else {
				BlackPieces = append(BlackPieces, piece)
				blackPos = append(blackPos, sq)
				pieceTot -= .3 * (value + getPosModifier(piece, sq))
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

	if position.Turn() == chess.Black {
		tempo = -0.1
	} else {
		tempo = 0.1
	}

	mobility := mobility(position)

	eval := (.55 * float64(pieceTot)) + (.3 * float64(attPot)) + (.05 * float64(wLeft)) - (.05*float64(bLeft) + tempo + (.6 * mobility))
	return eval
}
