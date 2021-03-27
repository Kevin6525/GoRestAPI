package main

import (
	"encoding/json"
	"net/http"
)

type Player struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height int     `json:"height"`
	Salary float32 `json:"salary"`
}

type playerHandlers struct {
	store map[string]Player
}

func (h *playerHandlers) get(w http.ResponseWriter, r *http.Request) {
	players := make([]Player, len(h.store))

	i := 0
	for _, player := range h.store {
		players[i] = player
		i++
	}
	jsonBytes, err := json.Marshal(players)
	if err != nil {
		// TODO
	}
	w.Write(jsonBytes)
}

func newPlayerHandlers() *playerHandlers {
	return &playerHandlers{
		store: map[string]Player{
			"player1": Player{
				Name:   "Aaron Rodgers",
				Age:    37,
				Height: 72,
				Salary: 33500000.0,
			},
		},
	}
}

func handleRequests() {
	playerHandlers := newPlayerHandlers()
	http.HandleFunc("/players", playerHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	handleRequests()
}
