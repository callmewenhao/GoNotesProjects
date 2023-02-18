# Go 中没有类

Go 和其它经典语言不通，它没有 class，没有对象，也没有继承

但是 Go 提供了struct 和方法。

## 将方法关联到 struct

方法可以被关联到你声明的类型上

```go
type coordinate struct {
    d, m, s float64
    h       rune
}

func (c coordinate) decimal() float64 {
    sign := 1.0
    switch c.h {
    case 'S', 'W', 's', 'w':
        sign = -1
    }
    return sign * (c.d + c.m/60 + c.s/3600)
}
```

## 构造函数

可以使用 struct 复合字面值来初始化你所要的数据。

但如果 struct 初始化的时候还要做很多事情，那就可以考虑写一个构造用的函数。

Go 语言没有专用的构造函数，但以 `new` 或者 `New` 开头的函数，
通常是用来构造数据的。例如 `newPerson()`，`NewPerson()`

```go
type coordinate struct {
    d, m, s float64
    h       rune
}

func (c coordinate) decimal() float64 {
    sign := 1.0
    switch c.h {
    case 'S', 'W', 's', 'w':
        sign = -1
    }
    return sign * (c.d + c.m/60 + c.s/3600)
}

type location struct {
    lat, long float64
}

func newLocation(lat, long coordinate) location {
    return location{lat.decimal(), long.decimal()}
}

func main() {
    lat := coordinate{4, 35, 22.2, 'S'}
    long := coordinate{137, 26, 30.12, 'E'}
    
    fmt.Println(lat.decimal(), long.decimal())
    fmt.Println(newLocation(lat, long))
}
```

## New 函数

有一些用于构造的函数的名称就是 `New`
（例如 `errors` 包里面的 `New` 函数）

这是因为函数调用时使用 `包名.函数名` 的形式。
如果该函数叫 `NewError`，
那么调用的时候就是 `errors.NewError()` 
这就不如 `errors.New()` 简洁

## class 的替代方案

Go 语言没有 class，
但使用 struct 并配备几个方法也可以达到同样的效果





