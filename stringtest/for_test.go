package main

import ( "fmt")

func main() {
	for i, j := 0, 1; i < 10; i=j++ {
		fmt.Printf("hello")
	}
}
