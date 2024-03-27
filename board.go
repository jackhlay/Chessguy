package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var placeToInd = map[string]int{
	"a1": 63, "b1": 62, "c1": 61, "d1": 60, "e1": 59, "f1": 58, "g1": 57, "h1": 56,
	"a2": 55, "b2": 54, "c2": 53, "d2": 52, "e2": 51, "f2": 50, "g2": 49, "h2": 48,
	"a3": 47, "b3": 46, "c3": 45, "d3": 44, "e3": 43, "f3": 42, "g3": 41, "h3": 40,
	"a4": 39, "b4": 38, "c4": 37, "d4": 36, "e4": 35, "f4": 34, "g4": 33, "h4": 32,
	"a5": 31, "b5": 30, "c5": 29, "d5": 28, "e5": 27, "f5": 26, "g5": 25, "h5": 24,
	"a6": 23, "b6": 22, "c6": 21, "d6": 20, "e6": 19, "f6": 18, "g6": 17, "h6": 16,
	"a7": 15, "b7": 14, "c7": 13, "d7": 12, "e7": 11, "f7": 10, "g7": 9, "h7": 8,
	"a8": 7, "b8": 6, "c8": 5, "d8": 4, "e8": 3, "f8": 2, "g8": 1, "h8": 0,
}

type Board struct {
	Board [8][8]Space
}

type Place struct {
	File rune
	Rank int
}

type Space struct {
	Occupied bool
	Piece    *Piece
	Ind      int
	Place    Place
}

// Confirmed working accurately with random string 3Q4/4Pb2/p4p1q/rk6/3P4/pB1PP3/p5K1/3R4 w - - 0 1
func fenParsing(fen string) gameState {
	gameState := newGame()
	// Parts of a fen string
	// 1. Piece Placement
	// 2. Turn COlor (w or b)
	// 3. Castling Availability (a combination of KQkq or -)
	// 4. En Passant Target Square (a square in algebraic notation or -)
	// 5. Halfmove Clock (a nonnegative integer, representing the number of halfmoves since the last capture or pawn advance, used for the fifty-move rule.)
	// 6. Fullmove Number (a positive integer, representing the number of the full move. It starts at 1, and is incremented after)
	// Opening FEN String: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1

	parts := strings.Split(fen, " ")
	pos := parts[0]
	turn := parts[1]
	castling := parts[2]
	epSquare := parts[3]
	halfmoves := parts[4]
	moves := parts[5]

	fmt.Println(pos, turn, castling, epSquare, halfmoves, moves)

	// create a map for the piece symbols
	pieceMap := map[rune]Piece{
		'p': {Pawn, 'p', Black, -1, false},
		'P': {Pawn, 'P', White, 1, false},
		'n': {Knight, 'n', Black, -3, false},
		'N': {Knight, 'N', White, 3, false},
		'b': {Bishop, 'b', Black, -3, false},
		'B': {Bishop, 'B', White, 3, false},
		'r': {Rook, 'r', Black, -5, false},
		'R': {Rook, 'R', White, 5, false},
		'q': {Queen, 'q', Black, -9, false},
		'Q': {Queen, 'Q', White, 9, false},
		'k': {King, 'k', Black, -10, false},
		'K': {King, 'K', White, 10, false},
	}

	//Positional Parsing block
	row, col := 0, 0
	for _, char := range pos {
		// fmt.Printf("ROW : %d, COL : %d\n", row, col)
		if char == '/' {
			row++
			col = 0
		}

		if unicode.IsDigit(char) {
			intVal := int(char - '0')
			for j := 0; j < intVal; j++ {
				gameState.board[row][col+j] = Space{
					Occupied: false,
					Piece:    nil,
					Ind:      row*8 + col,
					Place:    Place{File: rune('a' + col), Rank: 8 - row},
				}
			}
			col += intVal

		} else {
			newPiece, ok := pieceMap[char]
			if ok {
				gameState.board[row][col] = Space{
					Occupied: true,
					Piece:    &newPiece,
					Ind:      row*8 + col,
					Place:    Place{File: rune('a' + col), Rank: 8 - row},
				}
				col++
			}
		}
	}
	//print board for debugging

	//Turn Handling
	if turn == "w" {
		gameState.turn = White
	} else {
		gameState.turn = Black
	}

	//Castling
	if castling != "-" {
		for _, char := range castling {
			switch char {
			case 'K':
				gameState.castlingVariables = gameState.castlingVariables | 1
			case 'Q':
				gameState.castlingVariables = gameState.castlingVariables | 2
			case 'k':
				gameState.castlingVariables = gameState.castlingVariables | 4
			case 'q':
				gameState.castlingVariables = gameState.castlingVariables | 8
			}
		}
	} else {
		gameState.castlingVariables = 0
	}

	//En Passant Square Handling
	if epSquare == "-" {
		gameState.epSquare = -1
	} else {
		num, err := strconv.Atoi(epSquare)
		if err != nil {
			gameState.epSquare = num
		}
	}

	//Halfmove count (moves since pawn move or capture)
	pawnMoves, err := strconv.Atoi(halfmoves)
	if err != nil {
		fmt.Errorf("Error parsing halfmove count: %v", err)
	}
	gameState.halfmoves = pawnMoves

	//Move Count
	numMoves, err := strconv.Atoi(moves)
	if err != nil {
		fmt.Errorf("Error parsing move count: %v", err)
	}
	gameState.numMoves = numMoves

	gameState.getBitBoards()
	// fmt.Printf("White Bitboard: %#v \n", gameState.pieceColorBitboards[White])
	// fmt.Printf("Black Bitboard: %#v", gameState.pieceColorBitboards[Black])
	return gameState
}

func printBoard(bb uint64) {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			square := uint64(1) << uint(8*row+col)
			if (bb & square) != 0 {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}
