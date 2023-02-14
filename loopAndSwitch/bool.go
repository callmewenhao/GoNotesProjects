package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	var c1 = "go east"
	if c1 == "go east" {
		fmt.Println("== go east!")
	} else if c1 == "go inside" {
		fmt.Println("== go inside!")
	} else {
		fmt.Println("others!")
	}
}
