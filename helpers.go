package stryktipset

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
