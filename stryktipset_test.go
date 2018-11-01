package stryktipset

import (
	"testing"
)

type TestBet struct {
	sek          int
	expectedFull int
	expectedHalf int
}

var testBets = []TestBet{
	{2, 0, 1}, {4, 0, 2}, {8, 0, 3}, {16, 0, 4}, {48, 1, 4}, {96, 1, 5}, {144, 2, 4}, {432, 3, 4},
}

func TestConvertSekToBet(t *testing.T) {
	for _, tt := range testBets {
		full, half := ConvertSekToBet(tt.sek)
		if full != tt.expectedFull && half != tt.expectedHalf {
			t.Errorf("ConvertToCovers(%d) should return %d %d but we got %d %d", tt.sek, tt.expectedFull, tt.expectedHalf, full, half)
		}
	}
}

func TestRandomizeBets(t *testing.T) {
	for _, tt := range testBets {
		numOfHalf := 0
		numOfFull := 0
		bet := RandomizeBet(tt.expectedFull, tt.expectedHalf)
		for _, row := range bet.Rows {
			if len(row) == 2 {
				numOfHalf++
			} else if len(row) == 3 {
				numOfFull++
			}
		}
		if numOfHalf != tt.expectedHalf && numOfFull != tt.expectedFull {
			t.Errorf("Number of half and full covers was incorrect in randomized bet, got: %d %d want %d %d", numOfHalf, tt.expectedHalf, numOfFull, tt.expectedFull)
		}
	}
}
