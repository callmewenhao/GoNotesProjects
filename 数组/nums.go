package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury✔️"
	planets[1] = "Venus❤️"
	planets[2] = "Earth🤗"

	earth := planets[2]
	fmt.Println(len(planets)) // 8
	fmt.Println(earth)

	dwarfs := [5]string{"ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}
