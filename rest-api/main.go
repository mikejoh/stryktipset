package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mikejoh/stryktipset"
)

type Bet struct {
	Sek  int `json:"sek"`
	Full int `json:"full"`
	Half int `json:"half"`
}

func Convert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var b Bet

	b.Sek, _ = strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(b.Sek)
	b.Full = full
	b.Half = half

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(b)
}

func RandomizedBet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sek, _ := strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(sek)

	bets := stryktipset.RandomizeBet(full, half)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(bets)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/convert/{sek}", Convert).Methods("GET")
	router.HandleFunc("/api/bet/{sek}", RandomizedBet).Methods("GET")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
