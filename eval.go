package main

import (
	"math"
	"sort"
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

type moveRating struct {
	move string
	eval float64
}

func getPosModifier(piece chess.Piece, sq chess.Square) float64 {
	value := 0.0
	index := int(sq)

	if piece.Color() == chess.Black {
		index = 63 - index
	}

	switch piece.Type() {
	case chess.Pawn:
		value += pawnTable[index]

	case chess.Knight:
		value += knightTable[index]

	case chess.Bishop:
		value += bishopTable[index]

	case chess.Rook:
		value += rookTable[index]

	case chess.Queen:
		value += queenTable[index]
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
		fen = strings.Replace(fen, "w", "b", 1)
	} else {
		fen = strings.Replace(fen, "b", "w", 1)
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

//TODO: Implement recursive search for live brute force eval.
//TODO: Keep search as a dfs, and implement alpha beta pruning to help with optimization

func bagTest(position chess.Position, game chess.Game) []moveRating {
	//Checks for any egreious positions to not even consider
	mvs := game.ValidMoves()
	initEval := evalPos(position)
	var turn chess.Color
	goodMoves := []moveRating{}
	for _, move := range mvs {
		pos := position.Update(move)
		if potentialMaterialLoss(position, *pos) {
			continue
		}
		if potentialMaterialGain(position, *pos) {
			goodMoves = append(goodMoves, moveRating{move.String(), evalPos(*pos)})
			continue
		}
		evalPos := evalPos(*pos)
		if turn == chess.White {
			if canTakeQueen(position) {
				mvr := moveRating{move.String(), 999}
				goodMoves = append(goodMoves, mvr)
				continue
			}
			if evalPos >= initEval {
				mvr := moveRating{move.String(), evalPos}
				goodMoves = append(goodMoves, mvr)
			}
		} else {
			if canTakeQueen(position) {
				mvr := moveRating{move.String(), 999}
				goodMoves = append(goodMoves, mvr)
				continue
			}
			if evalPos <= initEval {
				mvr := moveRating{move.String(), evalPos}
				goodMoves = append(goodMoves, mvr)
			}
		}

	}
	if turn == chess.White {
		sort.Slice(goodMoves, func(i, j int) bool {
			return goodMoves[i].eval > goodMoves[j].eval
		})
	} else {
		sort.Slice(goodMoves, func(i, j int) bool {
			return goodMoves[i].eval < goodMoves[j].eval
		})

	}
	return goodMoves
}

func centerControl(position chess.Position) float64 {
	centerSquares := []chess.Square{chess.D4, chess.D5, chess.E4, chess.E5}
	centerControl := 0.0
	for _, sq := range position.ValidMoves() {
		if sq.S2() == centerSquares[0] || sq.S2() == centerSquares[1] || sq.S2() == centerSquares[2] || sq.S2() == centerSquares[3] {
			centerControl += 0.1
		}
		if sq.S1() == centerSquares[0] || sq.S1() == centerSquares[1] || sq.S1() == centerSquares[2] || sq.S1() == centerSquares[3] {
			centerControl += 0.1
		}
	}
	if position.Turn() == chess.Black {
		centerControl = -1 * centerControl
	}
	return centerControl
}

func calcMaterial(position chess.Position) float64 {
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
				pieceTot += value
			} else {
				pieceTot -= value
			}
		}
	}
	return pieceTot
}

func potentialMaterialLoss(before chess.Position, after chess.Position) bool {
	//checks if a move results in a material loss
	//if it does, don't consider it
	oppMoves := after.ValidMoves()
	for _, move := range oppMoves {
		piece := after.Board().Piece(move.S2())
		if piece != chess.NoPiece {
			if piece.Color() != before.Turn() {
				return true
			}
		}
	}
	return false
}

func potentialMaterialGain(before chess.Position, after chess.Position) bool {
	//checks if a move results in a material loss
	//if it does, don't consider it
	oppMoves := after.ValidMoves()
	for _, move := range oppMoves {
		piece := after.Board().Piece(move.S2())
		if piece != chess.NoPiece {
			if piece.Color() != before.Turn() {
				return true
			}
		}
	}
	return false
}

func pieceCoordination(before chess.Position) float64 {
	coordination := 0.0
	//evaluates how well pieces are coordinated
	for _, move := range before.ValidMoves() {
		if before.Board().Piece(move.S1()) != chess.NoPiece && before.Board().Piece(move.S2()).Color() == before.Turn() {
			coordination += .01
		}
	}
	if before.Turn() == chess.Black {
		coordination = -1 * coordination
	}
	return coordination
}

func canTakeQueen(position chess.Position) bool {
	for _, move := range position.ValidMoves() {
		if position.Board().Piece(move.S2()).Type() == chess.Queen && position.Board().Piece(move.S2()).Color() != position.Turn() {
			return true
		}
	}
	return false

}

func deepen(position chess.Position, depth int, alpha, beta float64) float64 {
	if depth == 0 {
		return evalPos(position)
	}
	moves := position.ValidMoves()
	if len(moves) == 0 {
		return evalPos(position)
	}
	if position.Turn() == chess.White {
		maxEval := -math.MaxFloat64
		for _, move := range moves {
			pos := position.Update(move)
			eval := deepen(*pos, depth-1, alpha, beta)
			maxEval = math.Max(maxEval, eval)
			alpha = math.Max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := math.MaxFloat64
		for _, move := range moves {
			pos := position.Update(move)
			eval := deepen(*pos, depth-1, alpha, beta)
			minEval = math.Min(minEval, eval)
			beta = math.Min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}

func evalPos(position chess.Position) float64 {
	whitePieces := []chess.Piece{}
	whiteTargeted := []chess.Square{}
	whitePos := []chess.Square{}
	blackPieces := []chess.Piece{}
	blackTargeted := []chess.Square{}
	blackPos := []chess.Square{}

	tempo := 0.0

	// material eval: Start with raw material balance
	pieceValue := map[chess.PieceType]float64{
		chess.Pawn:   1.0,
		chess.Knight: 3.0,
		chess.Bishop: 3.1,
		chess.Rook:   5.0,
		chess.Queen:  9.0,
		chess.King:   4.0,
	}

	pieceTot := 0.0
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			value := pieceValue[piece.Type()]
			if piece.Color() == chess.White {
				whitePieces = append(whitePieces, piece)
				whitePos = append(whitePos, sq)
				pieceTot += value + getPosModifier(piece, sq) // Positional value included here
			} else {
				blackPieces = append(blackPieces, piece)
				blackPos = append(blackPos, sq)
				pieceTot -= value + getPosModifier(piece, sq) // Positional value included here
			}
		}
	}

	// Fraction of total pieces
	wLeft := float64(len(whitePieces)) / 16.0
	bLeft := float64(len(blackPieces)) / 16.0

	// Attacking pieces: Count attacks by each side
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
	attPot := float64(wAttacks + bAttacks)

	// Adjust tempo based on whose turn it is
	if position.Turn() == chess.Black {
		tempo = -0.1
	} else {
		tempo = 0.1
	}

	// Calculate mobility (how many legal moves each side has)
	mobilityValue := mobility(position)

	// Now calculate the final evaluation:
	// - Material weight (60%)
	// - Attacking potential (20%)
	// - Leftover material ratio (10%)
	// - Tempo (5%)
	// - Mobility (5%)
	eval := (0.5 * pieceTot) + (0.2 * attPot) + (0.15 * (wLeft - bLeft)) + (0.1 * mobilityValue) + (0.05 * tempo) + pieceCoordination(position) + (0.1 * centerControl(position))

	// Return the final evaluation score
	return float64(math.Round(eval*100) / 100)
}
