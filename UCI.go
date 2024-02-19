package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nullmove := "0000"
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read input from the UCI protocol
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Process the input
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := parts[0]
		switch command {
		case "uci":
			// Send UCI identification
			fmt.Println("id name Boby")
			fmt.Println("id author Jack")
			fmt.Println("uciok")
		case "debug":
			// Enable or disable debug mode
			fmt.Println("info string debug not supported")
		case "isready":
			// Send confirmation that engine is ready
			fmt.Println(nullmove)
			startIt()
			fmt.Println("readyok")
		case "setoption":
			fmt.Println("info string setoption not supported")
		case "register":
			// Register the engine
			fmt.Println("info string register not supported")
		case "ucinewgame":
			// Initialize a new game
			// Additional setup can be done here if needed
			fmt.Println("readyok")
		case "position":
			fs := ""
			if parts[1] == "startpos" {
				fs = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
			} else {
				fs = parts[1]
			}
			fmt.Println(fs)
			// Set up the position on the board
			// Example: position startpos moves e2e4 e7e5
			// Additional logic to parse and apply the moves can be added here
			// fmt.Println("fen/startpos:", parts[1])
			// fmt.Println("moves", parts[2])
		case "go":
			// Start searching for the best move
			// Additional logic for searching and sending the best move can be added here
			// For simplicity, let's just output a random move
			fmt.Println("bestmove e2e4")
		case "stop":
			// Stop searching for the best move
			// Additional logic for stopping the search can be added here
		case "ponderhit":
			// The opponent has played the expected move
			// Additional logic for handling the opponent's move can be added here
		case "quit":
			// Exit the program
			os.Exit(0)
		default:
			// Ignore unknown commands

		}

	}
}
