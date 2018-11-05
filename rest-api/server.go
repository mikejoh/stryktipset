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

// Bet type
type Bet struct {
	Sek  int `json:"sek"`
	Full int `json:"full"`
	Half int `json:"half"`
}

// Convert will convert a given amount of SEK (Swedish Crowns) to the amount of full and half covers you can bet
// Returns an JSON encoded Bet type
func Convert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var b Bet

	b.Sek, _ = strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(b.Sek)
	b.Full = full
	b.Half = half

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(b)
}

// RandomizedBet returns a JSON encoded array of 13 randomized bets
func RandomizedBet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sek, _ := strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(sek)

	bets := stryktipset.RandomizeBet(full, half)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(bets)
}

// Initializes the router and handler functions
func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/convert/{sek}", Convert).Methods("GET")
	router.HandleFunc("/api/bet/{sek}", RandomizedBet).Methods("GET")

	return router
}

func main() {
	router := Init()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
