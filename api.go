package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Lichess API URLs
const lichessAPIURL = "https://lichess.org/api"

// Structs for Lichess Game and Move data
type LichessGame struct {
	ID       string `json:"id"`
	Variant  string `json:"variant"`
	Status   string `json:"status"`
	Started  bool   `json:"started"`
	FullMove int    `json:"fullmove"`
}

type LichessMove struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Piece string `json:"p"`
}

// Endpoint to start a game seeking request
func seekGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Prepare your seeking request payload
	req := map[string]interface{}{
		"rated":   false,
		"variant": "standard",
		"timeControl": map[string]interface{}{
			"minutes":          5,
			"incrementSeconds": 5,
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Error marshalling request data", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post(lichessAPIURL+"/board/seek", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		http.Error(w, "Error seeking game", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var game LichessGame
	if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Respond with the game details
	json.NewEncoder(w).Encode(game)
}

// Function to send a move to Lichess
func sendMove(gameID, move string) error {
	url := fmt.Sprintf("%s/game/%s/move/%s", lichessAPIURL, gameID, move)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send move, status: %s", resp.Status)
	}

	return nil
}

func ongoingGames(w http.ResponseWriter, r *http.Request) {
	url := "https://lichess.org/api/account/playing"
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
	}
	println(resp.Body)
}

// Endpoint for making a move
func handleMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		GameID string `json:"gameID"`
		Move   string `json:"move"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := sendMove(data.GameID, data.Move)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to make move: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Move sent successfully"))
}

func listen() {
	http.HandleFunc("/seek", seekGame)
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/games", ongoingGames)

	log.Println("Starting server on :5000...")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}

func APImain() {
	listen()
}
