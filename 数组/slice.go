package main

import "fmt"

func main() {
	//neptune := "Neptune"
	//tune := neptune[3:]
	//
	//fmt.Println(tune) // tune
	//neptune = "Poseidon"
	//fmt.Println(tune) // tune

	//planets := []string{"Mercury", "Venus", "Earth", "Mars",
	//	"Jupiter", "Saturn", "Uranus", "Neptune",
	//}
	//// 先将 planets 转成 StringSlice 类型，在进行排序
	//sort.StringSlice(planets).Sort()
	//
	//fmt.Println(planets)

	//dwarfs := []string{"ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//dwarfs = append(dwarfs, "Orus")
	//dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	//fmt.Println(dwarfs)

	//dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	//dwarfs2 := append(dwarfs1, "Orcus")
	//dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	//
	//dump("dwarfs1", dwarfs1)
	//dump("dwarfs2", dwarfs2)
	//dump("dwarfs3", dwarfs3)
	//
	//dwarfs3[1] = "Pluto!"
	//dump("dwarfs1", dwarfs1)
	//dump("dwarfs2", dwarfs2)
	//dump("dwarfs3", dwarfs3)

	//planets := []string{
	//	"Mercury", "Venus", "Earth", "Mars",
	//	" Jupiter", "Saturn", "Uranus", "Neptune",
	//}
	//terrestrial := planets[0:4:4]
	//worlds := append(terrestrial, "Ceres")
	//dump("planets", planets)
	//dump("terrestrial", terrestrial)
	//dump("worlds", worlds)

	//dwarfs1 := make([]string, 0, 10) // 长度是 0 容量是 10
	//dwarfs2 := make([]string, 10)    // 长度是 10 容量是 10
	//dump(" dwarfs1", dwarfs1)
	//dump(" dwarfs2", dwarfs2)
	//dwarfs1 = append(dwarfs1, "ceres", "Pluto", "Haumea", "Makemake", "Eris") // 在原来的基础上 append
	//dump(" dwarfs1", dwarfs1)

	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)

}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice),
		cap(slice), slice)
}
