package main

import "fmt"

type celsius float64

type report struct {
	// 在这里嵌入被指类型也可以 比如 int
	// int
	sol          int
	temperatures // 后面在调用变量时直接使用类型名就行
	locations
}

type temperatures struct {
	high, low celsius
}

func (t temperatures) average() celsius {
	return (t.low + t.high) / 2
}

func (r report) average() celsius {
	return r.temperatures.average() // 这里是对函数的转发
}

type locations struct {
	lat, long float64
}

func main() {
	bradbury := locations{-4.5895, 137.4417}
	t := temperatures{high: -1.0, low: -78.0}
	fmt.Println(t.average())

	report := report{
		sol:          15,
		temperatures: t,
		locations:    bradbury,
	}
	fmt.Println(report.average()) // 方法被转发 可以直接调用

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.high)              // 字段被转发 可以直接调用
	fmt.Printf("a balmy %v° C\n", report.temperatures.high) // 字段的访问
}
