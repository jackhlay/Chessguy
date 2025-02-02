package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/corentings/chess"
)

var api_key = ""

func sendToFrontend(pos chess.Position) {
	url := "http://127.0.0.1:5000/api/update"
	payload := map[string]interface{}{
		"fen":         pos.String(),
		"evalSource1": evalPos(pos),
		"evalSource2": "N/A",
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
	resp.Body.Close()

	// Print the response status
	fmt.Println("Response status:", resp.Status)
}

func liChessEval(fen string, apiTok string) string {
	//format fen string for url use
	fen = strings.ReplaceAll(fen, " ", "%20")
	url := "http://127.0.0.1:7000/eval?fenStr=" + fen + "&api_tok=" + apiTok
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "-9999"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "-9999"
	}
	bodyStr := string(body)

	if strings.Contains(bodyStr, "404") {
		return "-404"
	}

	return bodyStr

}

func exitProc() {
	//save the transposition table to a file
	if len(transposeTable) > 0 {
		file, err := os.Create("transpositionTable.txt")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		for k, v := range transposeTable {
			file.WriteString(k + ":" + strconv.FormatFloat(float64(v), 'f', -1, 32) + "\n")
		}
	}
}

func loadTranspositionTable() map[string]float32 {
	file, err := os.Open("transpositionTable.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	transposeTable := make(map[string]float32)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}
		key := parts[0]
		value, err := strconv.ParseFloat(parts[1], 32)
		if err != nil {
			fmt.Println("Error parsing value:", err)
			continue
		}
		transposeTable[key] = float32(value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return transposeTable
}

func main() {
	transposeTable = loadTranspositionTable()
	if transposeTable == nil {
		transposeTable = make(map[string]float32)
	}
	defer exitProc()
	//main func for real time streaming the games to the frontend
	for {
		game := chess.NewGame()
		for game.Outcome() == chess.NoOutcome {
			pos := game.Position()
			moves := pos.ValidMoves()

			startFen := game.Position().String()

			splitFen := strings.Split(startFen, " ")

			trn, _ := strconv.Atoi(splitFen[len(splitFen)-1])
			if len(moves) == 0 || trn > 113 {
				break
			}
			var move *chess.Move
			if pos.Turn() == chess.White {

				// if pos.Turn() == chess.White {
				// 	move, _ = chess.UCINotation{}.Decode(game.Position(), testRes[len(testRes)-1].move)

				// } else {
				// 	move, _ = chess.UCINotation{}.Decode(game.Position(), testRes[0].move)

				// }

			}
			// } else {
			// 	//wait for user input, convert string to move ex Nf3
			// 	fmt.Println("Enter move:")
			// 	var moveStr string
			// 	fmt.Scanln(&moveStr)
			// 	move, _ = chess.UCINotation{}.Decode(game.Position(), moveStr)
			// }
			game.Move(move)
			fmt.Printf("Move: %v\n", move)
			sendToFrontend(*game.Position())
		}
	}
}

// func main() {
// 	//main func for real time streaming the games to the frontend
// 	for {
// 		game := chess.NewGame()
// 		for game.Outcome() == chess.NoOutcome {
// 			pos := game.Position()
// 			moves := pos.ValidMoves()

// 			startFen := game.Position().String()

// 			splitFen := strings.Split(startFen, " ")

// 			trn, _ := strconv.Atoi(splitFen[len(splitFen)-1])
// 			if len(moves) == 0 || trn > 113 {
// 				break
// 			}
// 			var move *chess.Move
// 			if trn < 11 {
// 				move = moves[rand.Intn(len(moves))]
// 			} else {

// 				testRes := bagTest(*pos, *game)
// 				fmt.Printf("Bag Test Result: %v\n", testRes)
// 				if len(testRes) == 0 {
// 					move = moves[rand.Intn(len(moves))]
// 				} else {
// 					if pos.Turn() == chess.White {
// 						move, _ = chess.UCINotation{}.Decode(game.Position(), testRes[len(testRes)-1].move)

// 					} else {
// 						move, _ = chess.UCINotation{}.Decode(game.Position(), testRes[0].move)

// 					}

// 				}
// 			}
// 			game.Move(move)
// 			sendToFrontend(*pos)
// 			sleepTime := time.Duration(rand.Intn(873)) * time.Millisecond

// 			time.Sleep(sleepTime)

// 		}
// 	}
// }

// startRating := evalPos(*pos)
// startRating, _ := strconv.ParseFloat(liChessEval(startFen, api_key), 64)
// if startRating == -404 || startRating == -9999 || startRating == 9999 {
// 	startRating = evalPos(*pos)
// }
// endFen := game.Position().String()
// endRating := evalPos(*pos)
// endRating, _ := strconv.ParseFloat(liChessEval(endFen, api_key), 64)
// fmt.Printf("Lichess Eval: %v\n", endRating)
// if endRating == -404 || endRating == -9999 || endRating == 9999 {
// 	endRating = evalPos(*pos)
// }
// data := PosData{
// 	StartFen:    startFen,
// 	StartRating: startRating,
// 	Action:      move.String(),
// 	EndFen:      endFen,
// 	EndRating:   endRating,
// }
// fmt.Printf("Data: %v\n", data)
// sendJSON(data)
