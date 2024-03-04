package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const StartingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	debugg := false
	nullmove := "0000"
	reader := bufio.NewReader(os.Stdin)
	for {
		// Read input from the UCI protocol\

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Process the input
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := strings.ToLower(parts[0])
		//Maybe use a map in future for optimization, but for now, this is fine
		switch command {
		case "uci":
			// Send UCI identification
			fmt.Println("id name Boby")
			fmt.Println("id author Jack")
			fmt.Println("uciok")
		case "debug":
			// Enable or disable debug mode
			debugg = !debugg
			fmt.Println("debug mode:", debugg)
		case "isready":
			// Send confirmation that engine is ready
			fmt.Println(nullmove)
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
			if len(parts) > 1 {
				if parts[1] == "startpos" {
					fenParsing(StartingFen)

				} else {
					fenParsing(strings.Join(parts[1:], " "))
				}
			}
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
