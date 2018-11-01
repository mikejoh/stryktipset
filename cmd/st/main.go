package main

import (
	"flag"
	"fmt"

	st "github.com/mikejoh/stryktipset"
)

func main() {
	sekPtr := flag.Int("sek", 2, "Amount of SEK to bet.")
	flag.Parse()

	full, half := st.ConvertSekToBet(*sekPtr)
	fmt.Printf("Full: %d\nHalf: %d\n", full, half)
}
