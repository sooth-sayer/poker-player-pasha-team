package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ferki/poker-player-go/player"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 4711
	}

	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Printf("Error parsing form data: %s", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	action := request.FormValue("action")
	log.Printf("Request method=%s url=%s action=%s from client=%s\n", request.Method, request.URL, action, request.RemoteAddr)
	switch action {
	case "check":
		fmt.Fprint(w, "")
	case "bet_request":
		gameState, err := parseGameState(request.FormValue("game_state"))
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		result := player.BetRequest(gameState)
		fmt.Fprintf(w, "%d", result)
	case "showdown":
		gameState, err := parseGameState(request.FormValue("game_state"))
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		player.Showdown(gameState)
		fmt.Fprint(w, "")
	case "version":
		fmt.Fprint(w, player.Version())
	default:
		http.Error(w, "Invalid action", 400)
	}
}

func parseGameState(stateStr string) (gameState *player.GameState, err error) {
	if err = json.Unmarshal([]byte(stateStr), gameState); err != nil {
		log.Printf("Error parsing game state: %s", err)
		return nil, err
	}

	return gameState, nil
}
