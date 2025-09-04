package main

import "fmt"

func main() {
	fmt.Printf("We doing slices Now\n")
	//Simply portions of arrays

	var my_arr [5]int = [5]int{69, 21, 2055, 9000, 2077}

	//5 ints and here they are

	var slice_type_shii []int = my_arr[0:3]
	//Slice operator like in python
	fmt.Println(slice_type_shii)
	fmt.Println(cap(slice_type_shii))

	// fmt.Println(slice_type_shii[:cap(my_arr)])

	var giga []int = []int{1, 23, 4, 8, 21, 2, 1, 2, 64, 4, 6, 32, 3, 1, 3}
	giga = append(giga, 2055, 2077)

	//Makin a slice :-)
	made := make([]int, 0, 10) //Make an array of size 10 with a slice of length 0 but capacity 10, so a way of definig an array without populating or appending
	fmt.Println(made)
}
