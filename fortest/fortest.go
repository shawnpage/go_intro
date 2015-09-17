package main

import (
	"fmt"
)

func main() {
	var stop bool

	i := 0

	for !stop {
		fmt.Printf("Hello, World!\n")

		i += 1
		if i == 10 {
			stop = true
		}
	}
	// for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
	// 	fmt.Printf("%d Hello World!\n", j)
	// }

}
