package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := Init()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetConvertEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/convert/192", nil)

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var b Bet

	json.Unmarshal(response.Body.Bytes(), &b)

	if b.Sek != 192 {
		t.Errorf("Expected amount of SEK to be '192'. Got '%v'", b.Sek)
	}

	if b.Full != 1 {
		t.Errorf("Expected full covers expected to be '1'. Got '%v'", b.Full)
	}

	if b.Half != 6 {
		t.Errorf("Expected half covers expected to be '6'. Got '%v'", b.Half)
	}
}
