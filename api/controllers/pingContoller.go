package controllers

import (
	"encoding/json"
	"net/http"
)

type pongMessage struct {
	Message string `json:"message"`
}

func InitPingController(){
	
	http.Handle("GET /ping", http.HandlerFunc(Ping))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	response := pongMessage{
		Message: "pong",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
