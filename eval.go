package main

import (
	"math"
	"sort"

	"github.com/corentings/chess"
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
	control := 0.0

	for _, sq := range centerSquares {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			if piece.Color() == chess.White {
				control += 0.25
			} else {
				control -= 0.25
			}
		}
	}

	return control
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

func material(position chess.Position) float64 {
	// material eval: Start with raw material balance
	pieceValue := map[chess.PieceType]float64{
		chess.Pawn:   1.0,
		chess.Knight: 3.0,
		chess.Bishop: 3.1,
		chess.Rook:   5.0,
		chess.Queen:  9.0,
		chess.King:   0.0,
	}

	pieceTot := 0.0
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			value := pieceValue[piece.Type()]
			if piece.Color() == chess.White {
				pieceTot += value // Positional value included here
			} else {
				pieceTot -= value // Positional value included here
			}
		}
	}
	return pieceTot

}

func positional(position chess.Position) float64 {
	positional := 0.0
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			positional += getPosModifier(piece, sq)
		}
	}
	return positional
}

func mobility(position chess.Position) float64 {
	var whiteMoves []*chess.Move
	var blackMoves []*chess.Move

	if position.Turn() == chess.White {
		whiteMoves = position.ValidMoves()
		pos := position.ChangeTurn()
		blackMoves = pos.ValidMoves()
	} else {
		blackMoves = position.ValidMoves()
		pos := position.ChangeTurn()
		whiteMoves = pos.ValidMoves()
	}

	whiteMobility := 0.0
	for _, move := range whiteMoves {
		if isCentral(move.S2()) {
			whiteMobility += .2
		} else {
			whiteMobility += .1
		}

	}
	blackMobility := 0.0
	for _, move := range blackMoves {
		if isCentral(move.S2()) {
			blackMobility += .2
		} else {
			blackMobility += .1
		}

	}

	return float64(whiteMobility - blackMobility)

}

func isCentral(sq chess.Square) bool {

	centralSquares := []chess.Square{chess.D4, chess.D5, chess.E4, chess.E5}
	for _, centralSq := range centralSquares {
		if sq == centralSq {
			return true
		}
	}
	return false

}

func evalPos(position chess.Position) float64 {
	material := material(position)
	positional := positional(position)
	mobility := mobility(position)
	pawnStructure := 1.5 // for now
	kingSafety := 20.0   // for now
	centerControl := centerControl(position)

	eval := 0.6*material + 0.2*positional + 0.1*mobility + 0.05*pawnStructure + 0.05*kingSafety + 0.1*centerControl
	return eval
}
