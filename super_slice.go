package main

import (
	"fmt"
)

func main() {
	//We are slicing some more
	// my_slice := make([]int, 7)
	my_slice := []int{12, 32, 12343, 65, 12, 7686}

	// for i := range len(my_slice) { //Range incremets itself so no    i++ shii
	// 	fmt.Println(i)
	// }
	for iterator, slice_element := range my_slice { // _ --> throw am
		for jeterator, slice_element_v2 := range my_slice {
			if slice_element == slice_element_v2 && jeterator > iterator { // j before i cuz print it where is smallest i.e first repetition
				fmt.Println(slice_element)

			}

		}

		// fmt.Println(slice_element)

	}
}
