package main

import (
	"fmt"
)

func main() {
	var rand_int int = 69

	switch rand_int /*--> 69*/ {
	// case rand_int % 2 ==0: can not compare 69 to true or false
	case 3069:
		fmt.Println("Dario")
	case 69, 96: /*69 or 96*/
		fmt.Println("Go GO GO")
	default:
		fmt.Println("You no go reach me")

	}
	//I could have done What I wanted to do if i dont put anything after switch so kinda like pretty if else
}
