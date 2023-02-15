package main

import "fmt"

func main() {
	// 声明字符串
	//peace := "peace"
	//var peace = "peace"
	//var peace string = "peace"

	//c1 := '🙂'
	//var c2 = '🤣'
	//var c3 rune = '🤗'
	//
	//fmt.Printf("%v \n", c1)
	//fmt.Printf("%v \n", c2)
	//fmt.Printf("%v \n", c3)

	//message := "shalom"
	//c := message[5]
	//fmt.Printf("%c\n", c)
	//
	//message[5] = 'd' // 无法赋值

	//c := 'c'
	//a := c - 2
	//fmt.Printf("%v \n", c)
	//fmt.Printf("%v \n", a)

	question := "Do you ❤️🤗?"
	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}

}
