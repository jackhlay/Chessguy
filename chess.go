package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

func sendToFrontend(fen string) {
	fmt.Printf("wip: %s\n", fen)
}

func main() {
	for {
		game := chess.NewGame()
		for game.Outcome() == chess.NoOutcome {
			pos := game.Position()
			moves := pos.ValidMoves()
			if len(moves) == 0 {
				break
			}
			move := moves[rand.Intn(len(moves))]
			startFen := game.Position().String()
			startRating := evalPos(*pos)

			game.Move(move)
			fmt.Printf("Move: %s\n", move.String())
			endFen := game.Position().String()
			endRating := evalPos(*pos)

			data := PosData{
				StartFen:    startFen,
				StartRating: startRating,
				Action:      move.String(),
				EndFen:      endFen,
				EndRating:   endRating,
			}
			fmt.Printf("Data: %v\n", data)
			sendJSON(data)
			// sendToFrontend(endFen)
			time.Sleep(150 * time.Millisecond)
		}
	}
}
