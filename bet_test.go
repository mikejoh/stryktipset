package stryktipset

import (
	"strings"
	"testing"
)

func TestGetRandomBet(t *testing.T) {
	randomBet := GetRandomBet()
	if !strings.ContainsAny(randomBet, "1X2") {
		t.Errorf("GetRandomBet() returned %s expected %s or %s or %s", randomBet, "1", "X", "2")
	}
}

func TestCreateBet(t *testing.T) {
	createBetTests := map[string]int{
		"full":   3,
		"half":   2,
		"single": 1,
	}

	for bet, expectedNum := range createBetTests {
		var b Bet
		createdBet := b.Create(bet)

		numOfMarks := len(createdBet.String())

		switch bet {
		case "full":
			if numOfMarks != expectedNum {
				t.Errorf("CreateBet(\"%s\") returned number of possible marks %d expected %d", bet, numOfMarks, expectedNum)
			}
		case "half":
			if numOfMarks != expectedNum {
				t.Errorf("CreateBet(\"%s\") returned number of possible marks %d expected %d", bet, numOfMarks, expectedNum)
			}
		case "single":
			if numOfMarks != expectedNum {
				t.Errorf("CreateBet(\"%s\") returned number of possible marks %d expected %d", bet, numOfMarks, expectedNum)
			}
		}
	}
}

func TestStringBet(t *testing.T) {
	var b Bet
	b.Home = "1"
	b.Draw = "X"

	if b.String() != "1X" {
		t.Errorf("String() returned %s expexted %s", b.String(), "1X")
	}
}
