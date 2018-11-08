package stryktipset

import (
	"math/rand"
	"time"
)

// Coupon type
type Coupon struct {
	Bets map[int]Bet `json:"bets"`
}

// NewCoupon is a constructor for creating a new Coupon
func NewCoupon() *Coupon {
	var c Coupon
	c.Bets = make(map[int]Bet)
	return &c
}

// Create fills the coupon with random bets
func (c *Coupon) Create(full, half int) {
	numMatches := 13

	rand.Seed(time.Now().UnixNano())
	randomizeMatches := rand.Perm(numMatches)

	for _, match := range randomizeMatches {
		var bet Bet

		if full > 0 {
			c.Bets[match] = bet.Create("full")
			full--
		} else if half > 0 {
			c.Bets[match] = bet.Create("half")
			half--
		} else {
			c.Bets[match] = bet.Create("single")
		}
	}
}
