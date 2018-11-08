package stryktipset

import (
	"fmt"
	"math/rand"
	"time"
)

// Bet type
type Bet struct {
	Home string `json:"home"`
	Draw string `json:"draw"`
	Away string `json:"away"`
}

func (b Bet) String() string {
	return fmt.Sprintf("%s%s%s", b.Home, b.Draw, b.Away)
}

// GetRandomBet returns a random bet (home, draw or away)
func GetRandomBet() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3)

	var b string
	switch r {
	case 0:
		b = "1"
	case 1:
		b = "X"
	case 2:
		b = "2"
	}

	return b
}

// Create returns a Bet struct with randomized bets depending on the type (full, half or single bet)
func (b Bet) Create(t string) Bet {
	switch t {
	case "full":
		b.Home = "1"
		b.Draw = "X"
		b.Away = "2"
	case "half":
		var randomBets [2]string
		randomBets[0] = GetRandomBet()
		randomBets[1] = GetRandomBet()

		for randomBets[0] == randomBets[1] {
			randomBets[1] = GetRandomBet()
		}

		for _, randomBet := range randomBets {
			if randomBet == "1" {
				b.Home = randomBet
			} else if randomBet == "X" {
				b.Draw = randomBet
			} else if randomBet == "2" {
				b.Away = randomBet
			}
		}
	case "single":
		randomBet := GetRandomBet()
		if randomBet == "1" {
			b.Home = randomBet
		} else if randomBet == "X" {
			b.Draw = randomBet
		} else if randomBet == "2" {
			b.Away = randomBet
		}
	}

	return b
}
