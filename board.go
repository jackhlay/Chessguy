package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var placeToInd = map[string]int{
	"a1": 0, "b1": 1, "c1": 2, "d1": 3, "e1": 4, "f1": 5, "g1": 6, "h1": 7,
	"a2": 8, "b2": 9, "c2": 10, "d2": 11, "e2": 12, "f2": 13, "g2": 14, "h2": 15,
	"a3": 16, "b3": 17, "c3": 18, "d3": 19, "e3": 20, "f3": 21, "g3": 22, "h3": 23,
	"a4": 24, "b4": 25, "c4": 26, "d4": 27, "e4": 28, "f4": 29, "g4": 30, "h4": 31,
	"a5": 32, "b5": 33, "c5": 34, "d5": 35, "e5": 36, "f5": 37, "g5": 38, "h5": 39,
	"a6": 40, "b6": 41, "c6": 42, "d6": 43, "e6": 44, "f6": 45, "g6": 46, "h6": 47,
	"a7": 48, "b7": 49, "c7": 50, "d7": 51, "e7": 52, "f7": 53, "g7": 54, "h7": 55,
	"a8": 56, "b8": 57, "c8": 58, "d8": 59, "e8": 60, "f8": 61, "g8": 62, "h8": 63,
}

type Board struct {
	Spaces [8][8]Space
	Pieces [32]Piece
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

func (b Board) String() string {
	var str string
	for _, row := range b.Spaces {
		for _, space := range row {
			if space.Occupied {
				str += fmt.Sprintf("%v ", string(rune(space.Piece.Symbol))) // assuming Piece has a Symbol field
			} else {
				str += ". "
			}
		}
		str += "\n"
	}
	return str
}

// Confirmed working accurately with random string 3Q4/4Pb2/p4p1q/rk6/3P4/pB1PP3/p5K1/3R4 w - - 0 1
func fenParsing(fen string) {
	board, gameState := newGame()
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
				board.Spaces[row][col+j] = Space{
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
				board.Spaces[row][col] = Space{
					Occupied: true,
					Piece:    &newPiece,
					Ind:      row*8 + col,
					Place:    Place{File: rune('a' + col), Rank: 8 - row},
				}
				col++
			}
		}
	}

	fmt.Println(board)

	//Turn Handling
	if turn == "w" {
		gameState.turn = White
	} else if turn == "b" {
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
		return
	}
	gameState.halfmoves = pawnMoves

	//Move Count
	numMoves, err := strconv.Atoi(moves)
	if err != nil {
		return
	}
	gameState.numMoves = numMoves
}

func (b *Board) getPieceAt(place Place) (*Piece, int) {
	return b.Spaces[place.Rank][place.File-'a'].Piece, placeToInd[string(place.Rank)+string(place.File)]
}

func (b *Board) setPieceAt(place Place, piece *Piece) {
	b.Spaces[place.Rank][place.File-'a'].Piece = piece
}
