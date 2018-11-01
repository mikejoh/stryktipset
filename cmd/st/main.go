package main

import "flag"
import st "github.com/mikejoh/stryktipset"

func main() {
	sekPtr := flag.Int("sek", 0, "Amount of SEK to bet.")
	flag.Parse()

	st.ConvertSekToBet()
}
