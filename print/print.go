package main

import "fmt"

func main() {
	// 打印 string 使用逗号分隔 不带换行
	fmt.Print("print string without '\\n'! \n")
	// 打印 string 使用逗号分隔 带换行
	fmt.Println("print string with '\\n'!")
	// 格式化输出
	fmt.Printf("123 + 456 = %v\n", 123456)
	fmt.Printf("123 + 456 = %v\n", "123456")
	// 对齐操作
	// 在格式化动词里指定宽度,就可以对齐文本。
	// 例如, %4v ,就是向左填充到足够 4 个宽度
	// 正数, 向左填充空格
	// 负数, 向右填充空格
	fmt.Printf("%-15v $ %4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $ %4v\n", "Vrigin Galactic", 100)

}
