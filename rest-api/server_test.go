package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
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

type TestBet struct {
	sek          int
	expectedFull int
	expectedHalf int
}

var testBets = []TestBet{
	{2, 0, 1}, {4, 0, 2}, {8, 0, 3}, {16, 0, 4}, {48, 1, 4}, {96, 1, 5}, {144, 2, 4}, {432, 3, 4},
}

func TestGetConvertEndpoint(t *testing.T) {
	for _, tt := range testBets {
		sek := strconv.Itoa(tt.sek)
		expectedFullStr := strconv.Itoa(tt.expectedFull)
		expectedHalfStr := strconv.Itoa(tt.expectedHalf)

		req, _ := http.NewRequest("GET", "/api/convert/"+sek, nil)

		response := executeRequest(req)

		checkResponseCode(t, http.StatusOK, response.Code)

		var b Bet

		json.Unmarshal(response.Body.Bytes(), &b)

		if b.Sek != tt.sek {
			t.Errorf("Expected amount of SEK to be '"+sek+"'. Got '%v'", b.Sek)
		}

		if b.Full != tt.expectedFull {
			t.Errorf("Expected full covers expected to be '"+expectedFullStr+"'. Got '%v'", b.Full)
		}

		if b.Half != tt.expectedHalf {
			t.Errorf("Expected half covers expected to be '"+expectedHalfStr+"'. Got '%v'", b.Half)
		}
	}
}
