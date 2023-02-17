package main

import (
	"fmt"
	"sort"
)

func main() {
	//temperature := map[string]int{
	//	"Earth": 15,
	//	"Mars":  -65,
	//}
	//temp := temperature["Earth"]
	//fmt.Printf("on average the Earth is %v° C.\n", temp) // on average the Earth is 15° C.
	//
	//temperature["Earth"] = 16
	//temperature["Venus"] = 464
	//fmt.Println(temperature) // map[Earth:16 Mars:-65 Venus:464]
	//
	//moon := temperature["Moon"]
	//fmt.Println(moon)        // 0
	//fmt.Println(temperature) // map[Earth:16 Mars:-65 Venus:464] 没有因为 moon 而增加元素
	//
	//if moon, ok := temperature["Moon"]; ok {
	//	fmt.Printf("on average the Moon is %v° C.\n", moon)
	//} else {
	//	fmt.Printf("where is the moon?\n")
	//}

	//planets := map[string]string{
	//	"Earth": "sector zZ9",
	//	"Mars":  "sector zz9",
	//}
	//
	//planetsMarkII := planets
	//planets["Earth"] = "whoops"
	//
	//fmt.Println(planets) // 两者指向同一块内存地址
	//fmt.Println(planetsMarkII)
	//
	//delete(planets, "Earth") // 把对应的 key 和 value 删掉
	//fmt.Println(planetsMarkII)

	//temperature := make(map[float64]int, 8)
	//fmt.Println(len(temperature))
	//fmt.Println(temperature)
	// output
	// 0
	// map[]

	//temperature := []float64{
	//	-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	//}
	//
	//frequency := make(map[float64]int)
	//
	//for _, t := range temperature {
	//	frequency[t]++
	//}
	//
	//for t, num := range frequency { // 便利的顺序无法保证
	//	fmt.Printf("%+.2f occurs %d times\n", t, num)
	//}

	//temperature := []float64{
	//	-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	//}
	//groups := make(map[float64][]float64) // map 中存 float64 切片
	//
	//for _, t := range temperature {
	//	g := math.Trunc(t/10) * 10
	//	groups[g] = append(groups[g], t)
	//}
	//
	//for g, temperature := range groups {
	//	fmt.Printf("%v: %v\n", g, temperature)
	//}

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
	sort.Float64s(unique) // 由小到大排序
	fmt.Println(unique)
}
