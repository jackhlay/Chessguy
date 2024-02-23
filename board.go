package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

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
	board, gameState := startIt()
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
		gameState.turn = "white"
	} else if turn == "b" {
		gameState.turn = "black"
	}

	//Castling
	if castling != "-" {
		for _, char := range castling {
			switch char {
			case 'K':
				gameState.WKCastle = true
			case 'Q':
				gameState.WQCastle = true
			case 'k':
				gameState.BKCastle = true
			case 'q':
				gameState.BQCastle = true
			}
		}
	}

	//En Passant Square Handling
	if epSquare == "-" {
		gameState.epSquare = ""
	} else {
		gameState.epSquare = epSquare
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
