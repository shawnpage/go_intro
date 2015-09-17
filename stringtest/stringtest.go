package main

import (
	"fmt"
)

func main() {
	atoz := "The quick brown fox jumped over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough
		default:
			consonants += 1
		}
	}
	fmt.Printf("Vowels: %d; Consonants: %d; zeds: %d\n", vowels, consonants, zeds)
	//		fmt.Printf("%c\n", r)

}
