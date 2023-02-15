package main

import (
	"fmt"
	"strconv"
)

func main() {
	//countdown := "Launch in T minus " + "10 seconds."
	//countdown := "Launch in T minus " + string(rune(10)) + " seconds." // 报错
	//age := 41
	//marsAge := float64(age)

	//var bh float64 = 32768
	//var h = int16(bh)
	//fmt.Println(h) // -32768
	//countdown := 9
	//str := fmt.Sprintf("Lanuch in T minus %v seconds.", countdown)
	//fmt.Println(str)

	countdown, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("convert error!")
	}
	fmt.Println(countdown)

}
