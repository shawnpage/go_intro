package main

import (
	"fmt"
	// "os"
)

//func printer(msg, msg2 string, repeat int) {
//func printer(msg string) (string, error) {
//func printer(msg string) (returnedMessage string, e error) {
func printer(format string, msgs ...string) {

	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
	// _, e = fmt.Printf("%s\n", msg)
	// returnedMessage = "I did it"
	// return

	// f, err := os.Create("helloworld.txt")
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	// f.Write(msg)
	// return err

	// defer fmt.Printf("\n\n")
	// defer fmt.Printf("More\n\n")

	// _, err := fmt.Printf("%s", msg)
	// return err

	// _, err := fmt.Printf("%s", msg)
	// return msg, err

	// for repeat > 0 {
	// 	fmt.Printf("%s", msg)
	// 	fmt.Printf("%s\n", msg2)
	// 	repeat--
	// }
}

func main() {
	//	msg, err := printer("Hello, World")
	printer("% x\n", "Hello, World", "How are you?")

	// if err == nil {
	// 	fmt.Printf("%s\n", msg)
	// }

	// printer("Hello, ", "World", 5)
	// appendedMessage, err := printer("Hello, World")

	// if err == nil {
	// 	fmt.Printf("%q\n", appendedMessage)
	// 	fmt.Printf("%x\n", appendedMessage)
	// 	fmt.Printf("% x\n", appendedMessage)
	// }
}
