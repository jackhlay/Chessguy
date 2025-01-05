package main

import (
	"math"
	"sort"
	"strconv"
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
	move chess.Move
	eval float64
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

func getPiecePos(position chess.Position, piece chess.Piece) chess.Square {
	var retSq chess.Square
	for sq := chess.A1; sq <= chess.H8; sq++ {
		if position.Board().Piece(sq) == piece {
			retSq = sq
		}
	}
	return retSq
}

func euclideanDist(sq1 chess.Square, sq2 chess.Square) float64 {
	xDiff := sq1.Rank() - sq2.Rank()
	yDiff := sq1.File() - sq2.File()
	//euclidean distance
	xint, _ := strconv.Atoi(xDiff.String())
	yint, _ := strconv.Atoi(yDiff.String())
	dist := math.Sqrt(float64(xint*xint + yint*yint))

	return dist
}

func findKing(pos chess.Position, color chess.Color) chess.Square {
	var king chess.Square
	for sq := chess.A1; sq <= chess.H8; sq++ {
		if pos.Board().Piece(sq).Type() == chess.King && pos.Board().Piece(sq).Color() == color {
			king = sq
		}
	}
	return king
}

//Check for exposed king

func kingCheck(pos chess.Position, white []chess.Square, black []chess.Square) float64 {
	//Get King Square
	var whiteKing, blackKing chess.Square
	enemySum, friendlySum, enemyCt, friendlyCt := 0.0, 0.0, 0.0, 0.0

	enemyDistances := []float64{}
	friendlyDistances := []float64{}

	whiteKing = findKing(pos, chess.White)
	blackKing = findKing(pos, chess.Black)

	//calculate distance from king to all enemy pieces and friendly pieces
	//subtract distance from enemy pieces from distance from friendly pieces

	for _, sq := range white {
		if pos.Turn() == chess.White {
			enemyDistances = append(enemyDistances, euclideanDist(whiteKing, sq))
		} else {
			friendlyDistances = append(friendlyDistances, euclideanDist(whiteKing, sq))
		}
	}
	for _, sq := range black {
		if pos.Turn() == chess.White {
			friendlyDistances = append(friendlyDistances, euclideanDist(blackKing, sq))
		} else {
			enemyDistances = append(enemyDistances, euclideanDist(blackKing, sq))
		}
	}

	//sum the distances
	//average the distances
	for _, dist := range enemyDistances {
		enemySum += dist
		enemyCt++
	}
	for _, dist := range friendlyDistances {
		friendlySum += dist
		friendlyCt++
	}

	return (friendlySum/friendlyCt - enemySum/enemyCt)

}

//TODO: Implement recursive search for live brute force eval.
//TODO: Keep search as a dfs, and implement alpha beta pruning to help with optimization
//Additionally, before running the ABP, run a "Plastic-Bag check" to see if any moves are obviously bad and shouldn't be considered

func bagTest(position chess.Position, game chess.Game) []moveRating {
	//Checks for any egreious positions to not even consider
	mvs := game.ValidMoves()
	initEval := evalPos(position)
	var turn chess.Color
	goodMoves := []moveRating{}
	for _, move := range mvs {
		pos := position.Update(move)
		evalPos := evalPos(*pos)
		if turn == chess.White {
			if evalPos >= initEval {
				mvr := moveRating{*move, evalPos}
				goodMoves = append(goodMoves, mvr)
			}
		} else {
			if evalPos <= initEval {
				mvr := moveRating{*move, evalPos}
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

func alphaBetaPrune(position chess.Position, game chess.Game, depth int, alpha, beta float64, goodMoves []moveRating) float64 {
	if depth == 0 {
		return evalPos(position)
	}
	var mvs []*chess.Move
	if len(goodMoves) > 0 {
		for _, mv := range goodMoves {
			mvs = append(mvs, &mv.move)
		}
	} else {
		mvs = game.ValidMoves()
	}

	bestEval := -math.Inf(1)

	for _, move := range mvs {
		checkPos := game.Position().Update(move)
		eval := -alphaBetaPrune(*checkPos, game, depth-1, -beta, -alpha, goodMoves)

		bestEval = math.Max(bestEval, eval)
		alpha = math.Max(alpha, bestEval)

		if beta <= alpha {
			break
		}
	}

	return bestEval
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

func resultsInMaterialLoss(before chess.Position, after chess.Position) bool {
	//checks if a move results in a material loss
	//if it does, don't consider it
	beforeMaterial := calcMaterial(before)
	afterMaterial := calcMaterial(after)
	return afterMaterial < beforeMaterial
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
		chess.King:   0.0,
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
	eval := (0.6 * pieceTot) + (0.2 * attPot) + (0.1 * (wLeft - bLeft)) + (0.05 * tempo) + (0.05 * mobilityValue)

	// Return the final evaluation score
	return float64(math.Round(eval*100) / 100)
}
