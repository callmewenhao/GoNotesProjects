# 数

## 浮点型变量——实数

只要数字含有小数部分，那么它的类型就是 float64，相当于默认的浮点数类型

```go
d1 := 365.2425
var d2 = 365.2425
var d3 float64 = 365.2425
```

如果你使用一个整数来初始化某个变量，那么你必须指定它的类型为 float64，否则它就是一个整数类型

```go
var answer float64 = 42
```

### Go 语言里的两种浮点数类型

1. 默认是 float64
    - 64 位的浮点类型
    - 占用 8 字节内存
    - 某些编程语言把这种类型叫做 double（双精度)

2. float32
    - 占用 4 字节内存
    - 精度比 float64 低
    - 有时叫做单精度类型

想要使用单精度类型，你必须在声明变量的时候指定该类型

### 零值

Go 里面每个类型都有一个**默认值**，它称作**零值**

当你声明变量却**不对它进行初始化的时候**，它的值**就是零值**

### 显示浮点类型

使用 Print 或 Println 打印浮点类型的时候，默认的行为是尽可能的多显示几位小数

如果你不想这样，那么你应该使用 Printf 函数，结合 %f 格式化动词来指定显示小数的位数

```go
third := 1.0 / 3
fmt.Println(third)
fmt.Printf("%v\n", third)
fmt.Printf("%f\n", third)
fmt.Printf("%.3f\n", third)
fmt.Printf("%9.2f\n", third) // 不够9位就左填充空格

// output
0.3333333333333333
0.3333333333333333
0.333333
0.333
0.33
```

### %f 格式化动词

它由两部分组成：

宽度：会显示出的最少字符个数（包含小数点和小数）

- 如果宽度大于数字的个数，那么左边会填充空格
- 如果没指定宽度，那么就按实际的位数进行显示

精度：小数点后边显示的位数

使用 0 代替空格进行填充 `fmt.Printf("%05.2f\n", third)`

### 浮点类型的精度

浮点类型不适合用于金融类计算

为了尽量最小化舍入错误，建议先做乘法，再做除法

问题：如何避免舍入错误？可以暂时使用 `math.Abs` 比较精度解决这个问题

```go
func main() {
piggyBank := 0.1
piggyBank += 0.2
fmt.Println(piggyBank == 0.3) // 输出 false

fmt.Println(math.Abs(piggyBank-0.3) < 0.00001) // 输出 true
}
```

## 整数类型

Go 提供了 10 种整数类型：

- 不可以存小数部分

- 范围有限

- 通常根据数值范围来选取整数类型

5 种整数类型是有符号的

- 能表示正数、0、负数

5 种整数类型是无符号的

- 能表示正数、0

最常用的整数类型是 `int` `var year int = 2018`

无符号整数类型是 `uint` `var year uint = 2018`

`int` 和 `uint` 是针对目标设备优化的类型：

- 在树莓派2、比较老的移动设备上，int 和 uint 都是 32 位的。

- 在比较新的计算机上，int，和 uint，都是 64 位的。

Tip:

- 如果你在较老的 32 位设备上，使用了超过 20 亿的整数，并且代码还能运行

- 那么最好使用 int64 和 uint64 来代替 int 和 uint

剩余 8 种数据类型如下：

| Type   | Range                                                   | storage              |
|--------|---------------------------------------------------------|----------------------|
| int8   | -128 to 127                                             | 8-bit (1 byte)       |
| uint8  | 0 to 255                                                ||
| int16  | -32,768 to 32,767                                       | 16-bit (2 bytes)     |
| uint16 | 0 to 65535                                              ||
| int32  | -2,147,483,648 to 2,147,483,647                         | 32-bit (4 bytes)     |
| uint32 | 0 to 4,294,967,295                                      ||
| int64  | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 | 64-bit (eight bytes) |
| uint64 | 0 to 18,446,744,073,709,551,615                         |                      |

虽然在某些设备上 int 可以看作是 int32，在某些设备上 int 可以作是 int64，但它们其实是 3 种不同的类型。 
int 并不是其它类型的别名。

### 16 进制表示法

Go 语言里，在数前面加上 Ox 前缀，就可以用十六进制的形式来表示数值。

```go
var r, g, b uint8 = 0, 141, 255
var r, g, b uint8 = 0x00, 0x8d, 0xff
```

打印 16 进制：

打印十六进制的数，使用 `%x` 格式化动词：`fmt.Printf("%x %x %x", r, g, b)`

也可以指定最小宽度和填充：

```go
var r, g, b uint8 = 0x00, 0x8d, 0xff
fmt.Printf("%02x %02x %02x", r, g, b)
```

### 整数环绕

> 类似 python 的溢出处理，似乎更安全

所有的整数类型都有一个取值范围，超出这个范围，就会发生“环绕”

### 打印 bit 位

使用 `%b` 格式化动词

```go
var g uint8 = 3
fmt.Printf("%08b\n", g) // 00000011
```

### 整数最大、小值

math 包里，为与架构无关的整数类型，定义了最大、最小值常量：

- `math.MaxInt16`

- `math.MinInt64`

而 int 和 uint，可能是 32 或 64 位的。

## 很大的数

浮点类型可以存储非常大的数值，但是精度不高

整型很精确，但是取值范围有限。

如果你需要很大的数，而且要求很精确，那么怎么办

- int64 可以容纳很大的数，如果还不行，那么

- uint64可以容纳更大的正数，如果还不行，那么

- 也可以凑合用浮点类型，但是还有另外一种方法

- 使用 **big** 包。

```go
var dis uint64 = 24e18
fmt.Println(dis)
// # command-line-arguments
// .\nums.go:9:19: cannot use 24e18 (untyped float constant 2.4e+19) as uint64 value in variable declaration (truncated)
```

### big 包

对于较大的整数（超过10^18）：`big.Int`

对于任意精度的浮点类型，`big.Float`

对于分数，`big.Rat`

```go
lightSpeed := big.NewInt(299792)
secondsPerDay := big.NewInt(86400)

distance = new(big.Int)
distance.SetString("24000000000000000000", 10) // 10 进制
fmt.Println("Andromeda Galaxy is", distance, "km away.")

seconds := new( big.Int)
seconds.Div(distance, lightSpeed)

days := new(big.Int)
days.Div(seconds, secondsPerDay)

fmt.Println("That is", days,"days of travel at light speed.")
```
- 一旦使用了 `big.Int`，那么等式里其它的部分也必须使用 `big.Int`

- `NewInt()` 函数可以把 `int64` 转化为 `big.Int` 类型，但是不支持初始化大于 `int64` 的数

- 如何把 `24x1018` 转化为 `big.Int` 类型?

  - 首先 `new` 一个 `big.Int`

  - 再通过 `SetString` 函数把数值的字符串形式，和几进制传递进行即可

缺点：用起来繁琐，且速度慢

## 较大数值的常量

在 Go 里面，可以为常量指明类型（这句话会报错）

`const distance uint64 = 24000000000008000000  // 数超过了 uint64 的范围`

也可以不指明常量的类型

对于变量，Go会使用类型推断

而在 Go 里面，常量是可以无类型的（untyped），这句话就不会报错：

`const distance = 24000000000000000000`

常量使用 const 关键字来声明，程序里的每个字面值都是常量。

这意味着：比较大的数值可以直接使用（作为字面值）

针对字面值和常量的计算是在编译阶段完成的。

Go 的编译器是用 Go 编写的，这种无类型的数值字面值就是由 big 包所支持的。
这使你可以操作很大的数（超过 18 的 10^18）
只要能够容纳得下,那么常量就可以赋值给变量。

尽管 Go 编译器使用 big 包来处理无类型的数值常量，但是常量和 big.Int 的值是不能互换的。



