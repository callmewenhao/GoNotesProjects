# map

map 是 Go 提供的另外一种集合：

- 它可以将 key 映射到 value

- 可快速通过 key 找到对应的 value

- 它的 key 几乎可以是任何类型

## 声明 map

声明 map，必须指定 key 和 value 的类型：

```go
temperature := map[string]int{
    "Earth": 15,
    "Mars":  -65,
}
temp := temperature["Earth"]
fmt.Printf("on average the Earth is %v° C.\n", temp) // on average the Earth is 15° C.

temperature["Earth"] = 16
temperature["Venus"] = 464
fmt.Println(temperature) // map[Earth:16 Mars:-65 Venus:464]

moon := temperature["Moon"]
fmt.Println(moon)        // 0 返回了 int 的零值
fmt.Println(temperature) // map[Earth:16 Mars:-65 Venus:464] 没有因为 moon 而增加元素
```

## 逗号与 ok 写法

```go
if moon, ok := temperature["Moon"]; ok {
    fmt.Printf("on average the Moon is %v° C.\n", moon)
} else {
    fmt.Printf("where is the moon?\n")
}
```

## map 不会被复制

- 数组、int、float64 等类型在赋值给新变量或传递至函数/方法时候会创建相应的副本
 
- 但 map不会

```go
planets := map[string]string{
    "Earth": "sector zZ9",
    "Mars":  "sector zz9",
}

planetsMarkII := planets
planets["Earth"] = "whoops"

fmt.Println(planets) // 两者指向同一块内存地址，在传递给函数时也是一样的
fmt.Println(planetsMarkII)

delete(planets, "Earth") // 内置函数 把对应的 key 和 value 删掉
fmt.Println(planetsMarkII)
```

## 使用 make 函数对 map 进行预分配

除非你使用复合字面值来初始化 map，否则必须使用内置的 make 函数来为 map 分配空间。

创建 map 时，make 函数可接受一个或两个参数

第二个参数用于为指定数量的 key 预先分配空间

使用 make 函数创建的 map 的初始长度为 0

```go
temperature := make(map[float64]int, 8)
fmt.Println(len(temperature))
fmt.Println(temperature)
// output
0
map[]
```

## 使用 map 作计数器

```go
temperature := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -33.0,
}

frequency := make(map[float64]int)

for _, t := range temperature {
    frequency[t]++
}

for t, num := range frequency { // 便利的顺序无法保证
    fmt.Printf("%+.2f occurs %d times\n", t, num)
}
```

## 使用 map 和 slice 实现数据分组

```go
temperature := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
}

groups := make(map[float64][]float64) // map 中存 float64 切片

for _, t := range temperature {
    g := math.Trunc(t/10) * 10
    groups[g] = append(groups[g], t)
}

for g, temperature := range groups {
    fmt.Printf("%v: %v\n", g, temperature)
}

// output
-20: [-28 -29 -23 -29 -28]
30: [32]
-30: [-31 -33]
```

# 将 map 用作 set

Set 这种集合与数组类似，但元素不会重复 

Go 语言里没有提供 set 集合

```go
temperature := []float64{
    -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
}

set := make(map[float64]bool)
for _, t := range temperature {
    set[t] = true
}
if set[-28.0] {
    fmt.Println("set member!")
}
fmt.Println(set)

unique := make([]float64, 0, len(set))
for t := range set {
    unique = append(unique, t)
}
sort.Float64s(unique)  // 由小到大排序
fmt.Println(unique)

// output
set member!
map[-33:true -31:true -29:true -28:true -23:true 32:true]
[-33 -31 -29 -28 -23 32]
```













