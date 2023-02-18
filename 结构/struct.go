package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//var curiosity struct {
	//	lat  float64
	//	long float64
	//}
	//
	//curiosity.lat = -4.5678
	//curiosity.long = -137.4417
	//
	//fmt.Println(curiosity)

	//type location struct {
	//	lat  float64
	//	long float64
	//}
	//
	//var spirit location
	//spirit.lat = -14.5684
	//spirit.long = 175.472636
	//
	//var opportunity location
	//opportunity.lat = -1.9462
	//opportunity.long = 345.4734
	//
	//fmt.Println(spirit, opportunity)

	//type location struct {
	//	lat, long float64
	//}

	//opportunity := location{lat: -1.9462, long: 345.4734}
	//fmt.Println(opportunity)
	//
	//insight := location{lat: 4.5, long: 135.9}
	//fmt.Println(insight)
	//
	//spirit := location{173.09, -152.0}
	//fmt.Println(spirit)
	//
	//fmt.Printf("%v\n", insight)
	//fmt.Printf("%+v\n", insight)

	//spirit := location{173.09, -152.0}
	//new_spirit := spirit
	//new_spirit.long += 100
	//fmt.Println(spirit, new_spirit)

	//locations := []location{
	//	{lat: 1.0, long: 2.0},
	//	{lat: 2.0, long: 3.5},
	//	{lat: 4.0, long: -1.0},
	//}
	//fmt.Println(locations)

	type location struct {
		Lat  float64 `json:"latitude"` // 也可以是 xml 格式
		Long float64 `json:"longitude"`
	}

	curiosity := location{Lat: 134.09, Long: -152.1}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
