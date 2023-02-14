package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10
	for {
		if count < 0 {
			break
		}
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	/*for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}*/
	fmt.Println("Liftoff!")
}
