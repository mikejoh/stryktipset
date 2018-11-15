package stryktipset

import "testing"

var coupontests = []struct {
	wantFull    int
	wantHalf    int
	wantSingles int
}{
	{0, 6, 7},
	{1, 5, 6},
	{1, 6, 5},
}

func TestCreateCoupon(t *testing.T) {
	for _, tt := range coupontests {
		c := NewCoupon(tt.wantFull, tt.wantHalf)

		gotFull := 0
		gotHalf := 0
		gotSingles := 0

		for _, b := range c.Bets {
			if len(b.String()) == 3 {
				gotFull++
			} else if len(b.String()) == 2 {
				gotHalf++
			} else if len(b.String()) == 1 {
				gotSingles++
			}
		}

		if gotFull != tt.wantFull && gotHalf != tt.wantHalf && gotSingles != tt.wantSingles {
			t.Errorf("CreateCoupon() got = %d full, %d half and %d singles. want %d, %d and %d", gotFull, gotHalf, gotSingles, tt.wantFull, tt.wantHalf, tt.wantSingles)
		}
	}
}
