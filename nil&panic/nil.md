# nil

`nil` 是一个名词，表示 “无” 或者 “零”。

在 Go 里，`nil` 是一个零值。

如果一个指针没有明确的指向，那么它的值就是  `nil`

除了指针，`nil` 还是 `slice`、`map` 和接口的零值。

Go 语言的 `nil`，比以往语言中的 `null` 更为友好，
并且用的没那么频繁，但是仍需谨慎使用。

## nil 会导致 panic

如果指针没有明确的指向，
那么程序将无法对其实施的解引用。

尝试解引用一个 nil 指针将导致程序崩溃。

## 保护你的方法

避免 `nil` 引发 panic

```go
type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	var nobody *person
	fmt.Println(nobody)

	nobody.birthday() // 虽然时 nil 但是仍然可以调用，本质上相当于传参
}
```

因为值为 `nil` 的接收者和值为 `nil` 的参数在行为上并没有区别，
所以 Go 语言即使在接收者为 nil 的情况下，也会继续调用方法。

## nil 函数值

当变量被声明为函数类型时，它的默认值是 `nil`

```go
var fn func(a, b int) int
fmt.Println(fn == nil)  // true
```

检查函数值是否为 `nil`，并在有需要时提供默认行

```go
func sortStrings(s []string, less func(i, j int) bool) {
    if less == nil {
        less = func(i, j int) bool { return s[i] < s[j] }
    }
    sort.Slice(s, less)
}

func main() {
    food := []string{"onion", "carrot", "celery"}
    sortStrings(food, nil)
    fmt.Println(food)
    // [carrot celery onion]
}
```

## nil slice

如果 slice 在声明之后没有使用复合字面值或内置的 `make` 函数进行初始化，
那么它的值就是 `nil`。

幸运的是，`range`、`len`、`append` 等内置函数都可以正常处理值为 `nil` 的 slice。

虽然空 slice 和值为 `nil` 的 slice 并不相等，但它们通常可以替换使用。

## nil map

和 slice一样，
如果 `map` 在声明后没有使用复合字面值或内置的 make 函数进行初始化，
那么它的值将会是默认的 `nil`

```go
var soup map[string]int
fmt.Println(soup == nil)
measurement, ok := soup["onion"]
if ok {
    fmt.Println(measurement)
}
for ingredient, measurement := range soup {
    fmt.Println(ingredient, measurement)
}
```

## nil 接口

声明为接口类型的变量在未被赋值时，它的零值是 `nil`

对于一个未被赋值的接口变量来说，
它的接口类型和值都是 nil，并且变量本身也等于 nil。

当接口类型的变量被赋值后，接口就会在内部指向该变量的类型和值

在 Go 中，接口类型的变量只有在类型和值都为 nil 时才等于 nil
- 即使接口变量的值仍为 nil，但只要它的类型不是 nil，那么该变量就不等于 nil






