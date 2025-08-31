package main

import (
	"fmt"
)

func main() {

	//Chained conditionals
	val := 5 < 7 || 3 > 9
	fmt.Printf("%t\n", val)

	var andy bool = 3 < 5 && 5 > 33333
	fmt.Printf("%t", andy)

}
