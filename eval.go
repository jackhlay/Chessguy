package main

import (
	"math"
	"sort"
	"strings"

	"github.com/corentings/chess"
)

//Piece Square Tables:

var pawnTable = [64]float32{
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.5, 1.0, 1.0, -2.0, -2.0, 1.0, 1.0, 0.5,
	0.5, 0.5, 0.5, 1.5, 1.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, 0.5,
	0.5, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 1.5, 1.5, 0.5, 0.5, 0.5,
	0.5, 1.0, 1.0, -2.0, -2.0, 1.0, 1.0, 0.5,
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
}

var knightTable = [64]float32{
	-5.0, -4.0, -3.0, -3.0, -3.0, -3.0, -4.0, -5.0,
	-4.0, -2.0, 0.0, 0.5, 0.5, 0.0, -2.0, -4.0,
	-3.0, 0.5, 1.0, 1.5, 1.5, 1.0, 0.5, -3.0,
	-3.0, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, -3.0,
	-3.0, 0.5, 1.5, 2.0, 2.0, 1.5, 0.5, -3.0,
	-3.0, 0.5, 1.0, 1.5, 1.5, 1.0, 0.5, -3.0,
	-4.0, -2.0, 0.0, 0.5, 0.5, 0.0, -2.0, -4.0,
	-5.0, -4.0, -3.0, -3.0, -3.0, -3.0, -4.0, -5.0,
}

var bishopTable = [64]float32{
	-2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -2.0,
	-1.0, 0.5, 0.0, 0.0, 0.0, 0.0, 0.5, -1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0,
	-1.0, 1.0, 1.5, 2.0, 2.0, 1.5, 1.0, -1.0,
	-1.0, 1.0, 1.5, 2.0, 2.0, 1.5, 1.0, -1.0,
	-1.0, 1.0, 1.0, 1.5, 1.5, 1.0, 1.0, -1.0,
	-1.0, 0.5, 0.0, 0.0, 0.0, 0.0, 0.5, -1.0,
	-2.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -2.0,
}

var rookTable = [64]float32{
	0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0,
	0.5, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	-0.5, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, -0.5,
	0.0, 0.0, 0.0, 0.5, 0.5, 0.0, 0.0, 0.0,
}

var queenTable = [64]float32{
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
	eval float32
}

func getPosModifier(piece chess.Piece, sq chess.Square) float32 {
	value := float32(0)
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

func bagTest(position chess.Position, game chess.Game) []moveRating {
	//Checks for any egreious positions to not even consider0
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

func calcMaterial(position chess.Position) float32 {

	fen := position.String()
	whitePawns := strings.Count("P", fen)
	blackPawns := strings.Count("p", fen)

	whiteKnights := strings.Count("N", fen)
	blackKnights := strings.Count("n", fen)

	whiteBishops := strings.Count("B", fen)
	blackBishops := strings.Count("b", fen)

	whiteRooks := strings.Count("R", fen)
	blackRooks := strings.Count("r", fen)

	whiteQueens := strings.Count("Q", fen)
	blackQueens := strings.Count("q", fen)

	whiteMobility := 0
	blackMobility := 0
	if position.Turn() == chess.White {
		whiteMobility = len(position.ValidMoves())
		position.ChangeTurn()
		blackMobility = len(position.ValidMoves())
	} else {
		blackMobility = len(position.ValidMoves())
		position.ChangeTurn()
		whiteMobility = len(position.ValidMoves())
	}

	//doubled pawns
	//blocked pawns
	//isolatedpawns

	material := float32(whitePawns-blackPawns) + 3.2*float32(whiteKnights-blackKnights) + 3.3*float32(whiteBishops-blackBishops) + 5*float32(whiteRooks-blackRooks) + 9*float32(whiteQueens-blackQueens) + .1*float32(whiteMobility-blackMobility)

	return material
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

func canTakeQueen(position chess.Position) bool {
	for _, move := range position.ValidMoves() {
		if position.Board().Piece(move.S2()).Type() == chess.Queen && position.Board().Piece(move.S2()).Color() != position.Turn() {
			return true
		}
	}
	return false

}

func deepen(position chess.Position, depth int, alpha, beta float32) float32 {
	if depth == 0 {
		return evalPos(position)
	}
	moves := position.ValidMoves()
	if len(moves) == 0 {
		return evalPos(position)
	}
	if position.Turn() == chess.White {
		maxEval := float32(math.Inf(-1))
		for _, move := range moves {
			pos := position.Update(move)
			eval := deepen(*pos, depth-1, alpha, beta)
			if eval > float32(maxEval) {
				maxEval = eval
			}

			if eval > alpha {
				alpha = eval
			}

			if beta <= alpha {
				break
			}
		}
		return maxEval
	} else {
		minEval := float32(math.Inf(1))
		for _, move := range moves {
			pos := position.Update(move)
			eval := deepen(*pos, depth-1, alpha, beta)
			minEval = float32(math.Min(float64(minEval), float64(eval)))
			beta = float32(math.Min(float64(beta), float64(eval)))
			if beta <= alpha {
				break
			}
		}
		return minEval
	}
}
func evalPos(position chess.Position) float32 {
	material := calcMaterial(position)
	// positional := positional(position)
	eval := material
	return eval
}
