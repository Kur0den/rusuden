package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Message struct {
	ID         int    `json:"id"`
	CallerName string `json:"callerName"`
	CallerID   string `json:"callerId"`
	Created    string `json:"created"`
	URL        string `json:"url"`
}

func handleMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	testDate := time.Date(2026, 11, 23, 12, 00, 00, 00, time.UTC)
	message := []Message{
		{ID: 1, CallerName: "test01", CallerID: "1234", Created: testDate.String(), URL: "/dummy1.wav"},
		{ID: 2, CallerName: "test01", CallerID: "1234", Created: testDate.Add(time.Hour).String(), URL: "/dummy1.wav"},
		{ID: 3, CallerName: "test01", CallerID: "1234", Created: testDate.Add(time.Second).String(), URL: "/dummy1.wav"},
	}
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		println("えらーだよ: ", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/messages", handleMessages)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
