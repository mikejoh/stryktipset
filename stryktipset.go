package stryktipset

import (
	"math/rand"
	"strings"
	"time"
)

type Bet struct {
	Rows [13]string `json:"bets,omitempty"`
}

// Converts given amount of money you want to bet to the amount of full and half covers you can use
func ConvertSekToBet(sek int) (int, int) {
	full := 0
	half := 0

	for {
		if sek == 1 {
			break
		}

		if sek%2 == 0 {
			half++
			sek = sek / 2
		} else {
			full++
			sek = sek / 3
		}
	}

	return full, half
}

// Randomize bets of 13 matches
func RandomizeBet(full, half int) Bet {

	var bet Bet

	rand.Seed(time.Now().UnixNano())

	possibleBets := []string{"1", "X", "2"}

	randomizeMatches := rand.Perm(13)

	for index, match := range randomizeMatches {
		if full > 0 {
			bet.Rows[match] = strings.Join(possibleBets[:], "")
			full--
		} else if half > 0 {
			randomizeHalves := rand.Perm(len(possibleBets))
			bet.Rows[match] = possibleBets[randomizeHalves[0]] + possibleBets[randomizeHalves[1]]
			half--
		} else {
			bet.Rows[match] = possibleBets[rand.Intn(len(possibleBets))]
		}

		if index == 12 {
			break
		}
	}

	return bet
}
