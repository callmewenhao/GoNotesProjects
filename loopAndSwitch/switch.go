package main

import "fmt"

func main() {
	var command = "go inside"
	switch command {
	case "go east":
		fmt.Println("go east!")
	case "enter", "go inside":
		fmt.Println("enter")
		fmt.Println("or go inside")
		fallthrough
	case "read sign":
		fmt.Println("read sign")
	default:
		fmt.Println("others")
	}
}
