package main

import "fmt"

//type person struct {
//	name, superpower string
//	age              int
//}

//func birthday(p *person) {
//	p.age++
//}

//func (p *person) birthday() {
//	p.age++
//}

//type stats struct {
//	level             int
//	endurance, health int
//}
//
//func levelUp(s *stats) {
//	s.level++
//	s.endurance = 42 + (14 * s.level)
//	s.health = 5 * s.endurance
//}
//
//type character struct {
//	name  string
//	stats stats
//}

//func reset(board *[8][8]rune) {
//	board[0][0] = '😊'
//}

func reclassify(planets *[]string) {
	*planets = (*planets)[1:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets)

	//var board [8][8]rune
	//reset(&board)
	//fmt.Printf("%c\n", board[0][0])

	//player := character{
	//	name: "Matthias",
	//}
	//levelUp(&player.stats) // 直接使用 . 标记法 返回的就是指针
	//fmt.Printf("%+v\n", player.stats)

	//rebecca := person{
	//	name:       "Rebecca",
	//	superpower: "imagination",
	//	age:        14,
	//}
	//rebecca.birthday()
	//fmt.Printf("%+v\n", rebecca)
	//// {name:Rebecca superpower:imagination age:15}
	//
	//nathan := person{
	//	name: "Nathan",
	//	age:  17,
	//}
	//nathan.birthday() // 也可以调用 birthday()
	//fmt.Printf("%+v\n", nathan)
	// {name:Nathan superpower: age:18}

	//a := 100
	//fmt.Println(&a)

	//answer := 42
	//fmt.Println(&answer)

	//address := &answer
	//fmt.Println(*address)

	//answer := 42
	//address := &answer                       // 类型 *int
	//fmt.Printf("address is a %T\n", address) // address is a *int

	//canada := "Canada"
	//
	//var home *string
	//fmt.Printf("home is a %T\n", home)
	//
	//home = &canada
	//fmt.Println(home)

	//type person struct {
	//	name, superpower string
	//	age              int
	//}
	//
	//timmy := &person{
	//	name: "Timothy",
	//	age:  10,
	//}
	//
	//(*timmy).superpower = "flying" // 解引用加不加都行
	//timmy.superpower = "flying"
	//
	//fmt.Printf("%+v\n", timmy)
	// output
	//&{name:Timothy superpower:flying age:10}

	//superpowers := [3]string{"flight", "invisibility", "super strength"}
	//fmt.Println(superpowers[0])
	//fmt.Println(superpowers[1:2])
	// output
	// flight
	// [invisibility]

}
