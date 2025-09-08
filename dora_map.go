package main

import "fmt"

func main() {

	//I'm the map, I'm the map.

	//We can store key value pairs

	// var map_name map[key type]value type

	var my_map map[string]int = map[string]int{
		"jojo":   5,
		"koichi": 4,
		"whamu":  10,
	}

	fmt.Println(my_map) //Maps don't keep order cuz dora can't remember shit
	my_map["koichi"] = 2055
	fmt.Println(my_map)
	my_map["waste_my_time"] = 16
	fmt.Println(my_map)

	delete(my_map, "waste_my_time")
	fmt.Println(my_map)
	// for i, element := range my_map {
	// 	fmt.Printf("My key is %s and my value is %d \n", i, element)
	// }

	test, err := my_map["karss"]

	fmt.Println(test, err) // err --> bool for some reason
}
