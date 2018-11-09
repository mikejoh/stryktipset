package stryktipset

import "testing"

func TestCreateCoupon(t *testing.T) {
	c := NewCoupon()

	wantFull := 1
	wantHalf := 6
	wantSingles := 5

	gotFull := 0
	gotHalf := 0
	gotSingles := 0

	c.Create(wantFull, wantHalf)

	for _, b := range c.Bets {
		if len(b.String()) == 3 {
			gotFull++
		} else if len(b.String()) == 2 {
			gotHalf++
		} else if len(b.String()) == 1 {
			gotSingles++
		}
	}

	if gotFull != wantFull && gotHalf != wantHalf && gotSingles != wantSingles {
		t.Errorf("CreateCoupon() got = %d full, %d half and %d singles. want %d, %d and %d", wantFull, wantHalf, gotFull, gotHalf, gotSingles, wantFull, wantHalf, wantSingles)
	}
}
