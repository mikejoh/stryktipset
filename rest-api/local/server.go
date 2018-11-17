package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mikejoh/stryktipset"
)

// GetConvert will convert a given amount of SEK (Swedish Crowns) to the amount of full and half covers you can bet
// Returns an JSON encoded Bet type
func GetConvert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var c stryktipset.Convert

	c.Sek, _ = strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(c.Sek)
	c.Full = full
	c.Half = half

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	json.NewEncoder(w).Encode(c)
}

// GetCoupon returns a JSON encoded array of 13 randomized bets
func GetCoupon(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	sek, _ := strconv.Atoi(params["sek"])
	full, half := stryktipset.ConvertSekToBet(sek)

	c := stryktipset.NewCoupon(full, half)
	bets := c.Bets

	outputType := r.URL.Query().Get("output")

	switch outputType {
	case "html":
		// Respond with a HTML page
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		for _, b := range bets {
			fmt.Fprint(w, b.String()+"<br>")
		}
	default:
		// Always respond with the bets JSON encoded
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(bets)
	}
}

// Initializes the router and handler functions
func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/convert/{sek}", GetConvert).Methods("GET")
	router.HandleFunc("/api/coupon/{sek}", GetCoupon).Methods("GET")

	return router
}

func main() {
	router := Init()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
