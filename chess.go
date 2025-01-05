package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/notnil/chess"
)

func sendToFrontend(pos chess.Position) {
	url := "http://localhost:5000/api/update"
	payload := map[string]interface{}{
		"fen":         pos.String(),
		"evalSource1": evalPos(pos),
		"evalSource2": "0",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response status:", resp.Status)
}

func main() {
	//main func for real time streaming the games to the frontend
	for {
		game := chess.NewGame()
		for game.Outcome() == chess.NoOutcome {
			pos := game.Position()
			moves := pos.ValidMoves()
			if len(moves) == 0 {
				break
			}
			bagTest(*pos, *game)

			move := moves[rand.Intn(len(moves))]
			game.Move(move)
			sendToFrontend(*pos)
			sleepTime := time.Duration(rand.Intn(700)) * time.Millisecond
			time.Sleep(sleepTime)
		}
	}
}

// func main() {
// 	//main func for training / forawrd pass of the model
// 			move := moves[rand.Intn(len(moves))]
// 			usePos := pos.Update(move)
// 			startFen := usePos.String()
// 			startRating := evalPos(*usePos)

// 			game.Move(move)
// 			fmt.Printf("Move: %s\n", move.String())
// 			endFen := game.Position().String()
// 			endRating := evalPos(*pos)

// 			data := PosData{
// 				StartFen:    startFen,
// 				StartRating: startRating,
// 				Action:      move.String(),
// 				EndFen:      endFen,
// 				EndRating:   endRating,
// 			}
// 			fmt.Printf("Data: %v\n", data)
// 			// sendJSON(data)
// 			sendToFrontend(*usePos)
// 			sleepTime := 1*time.Second + time.Duration(rand.Intn(700))*time.Millisecond
// 			time.Sleep(sleepTime)
// 		}
// 	}
// }
