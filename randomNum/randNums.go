package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println("random num is", num)

	num = rand.Intn(10) + 1
	fmt.Println("random num is", num)
}
