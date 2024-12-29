package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PosData struct {
	StartFen    string  `json:"start_fen"`
	StartRating float64 `json:"start_rating"`
	Action      string  `json:"action"`
	EndFen      string  `json:"end_fen"`
	EndRating   float64 `json:"end_rating"`
}

func sendJSON(data PosData) {
	//send to dqn
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("http://10.0.0.51:8000", "application/json", bytes.NewBuffer(jsonData))
	// resp, err := http.Post("chessdqn.default.svc.cluster.local:8000", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error Making POST req: ", err)
	}

	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("200OK")
		fmt.Printf("%s", resp.Body)
	} else if resp.StatusCode == http.StatusResetContent {
		fmt.Printf("205 - End of the line! %s\n", resp.Body)
	} else {
		fmt.Printf("Failed with status code: %d\n", resp.StatusCode)
	}
}
