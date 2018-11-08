package stryktipset

import "testing"

func TestCreateCoupon(t *testing.T) {
	c := NewCoupon()

	expectedFull := 1
	expectedHalf := 6
	expectedSingles := 5

	returnedFull := 0
	returnedHalf := 0
	returnedSingles := 0

	c.Create(expectedFull, expectedHalf)

	for _, b := range c.Bets {
		if len(b.String()) == 3 {
			returnedFull++
		} else if len(b.String()) == 2 {
			returnedHalf++
		} else if len(b.String()) == 1 {
			returnedSingles++
		}
	}

	if returnedFull != expectedFull && returnedHalf != expectedHalf && returnedSingles != expectedSingles {
		t.Errorf("CreateCoupon(\"%d, %d\") returned %d full, %d half and %d singles. Expected %d, %d and %d", expectedFull, expectedHalf, returnedFull, returnedHalf, returnedSingles, expectedFull, expectedHalf, expectedSingles)
	}
}
