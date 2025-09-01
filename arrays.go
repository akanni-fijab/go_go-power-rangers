package main

import "fmt"

func main() {

	//We doi'n arrays
	//An array is an ordered collection of elements with fixed length

	var my_arr [7]string

	my_arr[0] = "lola"
	my_arr[1] = "is"
	my_arr[2] = "a"
	my_arr[3] = "girl's"
	my_arr[4] = "name"

	for i := range len(my_arr) { //Start at zero and stop at 6 in this case, i.e 7-1 ... my_arr[7] does not exist
		fmt.Println(my_arr[i], i)
		i++

	}

}
