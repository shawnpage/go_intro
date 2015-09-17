package main

import (
	"fmt"
)

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s\n", word)
	}
}
func main() {

	words := make([]string, 4)

	words[0] = "The"
	words[1] = "Quick"
	words[2] = "Brown"
	words[3] = "Fox"

	printer(words)
	// words := []string{"The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	// printer(words[2:5])
	//fmt.Printf("%d\n", (words))
}
