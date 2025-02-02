package main

import (
	"time"

	"github.com/corentings/chess"
)

//Piece Square Tables:

var transposeTable map[string]float32

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
	move chess.Move
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

func calcMaterial(position chess.Position) float32 {
	whitePawns, blackPawns := float32(0), float32(0)
	whiteKnights, blackKnights := float32(0), float32(0)
	whiteBishops, blackBishops := float32(0), float32(0)
	whiteRooks, blackRooks := float32(0), float32(0)
	whiteQueens, blackQueens := float32(0), float32(0)
	whiteKings, blackKings := float32(0), float32(0)
	var material float32
	for sq := chess.A1; sq <= chess.H8; sq++ {
		piece := position.Board().Piece(sq)
		if piece != chess.NoPiece {
			switch piece {
			case chess.Piece(chess.Pawn):
				if piece.Color() == chess.White {
					val := getPosModifier(piece, sq)
					whitePawns += val
				} else {
					val := getPosModifier(piece, sq)
					blackPawns += val * -1
				}
			case chess.Piece(chess.Knight):
				if piece.Color() == chess.White {
					val := getPosModifier(piece, sq)
					whiteKnights += val
				} else {
					val := getPosModifier(piece, sq)
					blackKnights += val * -1
				}
			case chess.Piece(chess.Bishop):
				if piece.Color() == chess.White {
					val := getPosModifier(piece, sq)
					whiteBishops += val
				} else {
					val := getPosModifier(piece, sq)
					blackBishops += val * -1
				}
			case chess.Piece(chess.Rook):
				if piece.Color() == chess.White {
					val := getPosModifier(piece, sq)
					whiteRooks += val
				} else {
					val := getPosModifier(piece, sq)
					blackRooks += val * -1
				}
			case chess.Piece(chess.Queen):
				if piece.Color() == chess.White {
					val := getPosModifier(piece, sq)
					whiteQueens += val
				} else {
					val := getPosModifier(piece, sq)
					blackQueens += val * -1
				}
			case chess.Piece(chess.King):
				if piece.Color() == chess.White {
					whiteKings += float32(200)
				} else {
					blackKings += float32(200) * -1
				}
			}
		}

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

		whiteDoubled, blackDoubled := doubledPawns(position)
		//blocked pawns
		//isolatedpawns

		material = 9*float32(whiteQueens-blackQueens) +
			5*float32(whiteRooks-blackRooks) +
			3.2*float32(whiteKnights-blackKnights) +
			3.3*float32(whiteBishops-blackBishops) +
			float32(whitePawns-blackPawns) -
			0.5*float32(whiteDoubled-blackDoubled) +
			.1*float32(whiteMobility-blackMobility)

		if position.Turn() == chess.Black {
			material *= -1
		}
	}
	return material
}

func doubledPawns(position chess.Position) (whiteDoubled, blackDoubled float32) {
	for file := chess.FileA; file <= chess.FileH; file++ {
		whitePawns := 0
		blackPawns := 0
		for rank := chess.Rank1; rank <= chess.Rank8; rank++ {
			sq := chess.NewSquare(file, rank)
			piece := position.Board().Piece(sq)
			if piece.Type() == chess.Pawn {
				if piece.Color() == chess.White {
					whitePawns++
				} else {
					blackPawns++
				}
			}
		}
		if whitePawns > 1 {
			whiteDoubled += 1
		}
		if blackPawns > 1 {
			blackDoubled += 1
		}
	}
	return whiteDoubled, blackDoubled

}

func deepen(position chess.Position, timeLeft int) chess.Move {
	startTime := time.Now()
	moveChan := make(chan chess.Move)

	go func() {
		for depth := 1; ; depth++ {
			negaMaxRoot(position, depth, moveChan)
			if time.Since(startTime) > (time.Duration(timeLeft)/20 + 3*time.Second/2) {
				break
			}
		}
	}()
	move := <-moveChan
	return move

}

func negaMax(remaining_depth int, position chess.Position) float32 {
	if remaining_depth == 0 {
		return evalPos(position)
	}
	maxeval := float32(-9999999999)
	for _, move := range position.ValidMoves() {
		pos := position.Update(move)
		eval := -negaMax(remaining_depth-1, *pos)
		if eval >= maxeval {
			maxeval = eval
		}
	}
	return maxeval
}

func negaMaxRoot(position chess.Position, depth int, movechan chan chess.Move) {
	moveCache := make(map[float32]chess.Move)
	maxeval := float32(-9999999999)
	for _, move := range position.ValidMoves() {
		pos := position.Update(move)
		movechan <- *move
		eval := -negaMax(3, *pos)
		moveCache[eval] = *move
		if eval >= maxeval {
			maxeval = eval
		}

	}
	movechan <- moveCache[maxeval]
}

func evalPos(position chess.Position) float32 {
	// time allowed for search: remaining time/20 + increment/2
	material := calcMaterial(position)
	// positional := positional(position)
	eval := material
	return eval
}
