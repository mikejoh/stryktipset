package stryktipset

import "testing"

var bettests = []struct {
	sek             int
	expectedFull    int
	expectedHalf    int
	expectedSingles int
}{
	{2, 0, 1, 12},
	{4, 0, 2, 11},
	{8, 0, 3, 10},
	{16, 0, 4, 9},
	{48, 1, 4, 8},
	{96, 1, 5, 7},
	{192, 1, 6, 6},
	{144, 2, 4, 7},
	{432, 3, 4, 6},
}

func TestConvertSekToBet(t *testing.T) {
	for _, tt := range bettests {
		gotFull, gotHalf := ConvertSekToBet(tt.sek)
		if gotFull != tt.expectedFull && gotHalf != tt.expectedHalf {
			t.Errorf("ConvertToCovers() got = %d %d, want = %d %d", tt.expectedFull, tt.expectedHalf, gotFull, gotHalf)
		}
	}
}

func TestConvertSekToBetNumberOfSingles(t *testing.T) {
	for _, tt := range bettests {
		gotFull, gotHalf := ConvertSekToBet(tt.sek)
		gotSingles := 13 - (gotFull + gotHalf)
		if gotSingles != tt.expectedSingles {
			t.Errorf("ConvertToCovers(%d) got = %d singles, want = %d singles", tt.sek, tt.expectedSingles, gotSingles)
		}
	}
}
