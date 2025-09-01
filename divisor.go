package main

import (
	"fmt"
)

func main() {

	for quotient := range 1000 { // quotient is declared as an int unlike the loops one where i had to specify for concurrency
		quotient += 1 // Tell it to start from 1
		if quotient%3 == 0 && quotient%7 == 0 && quotient%9 == 0 {
			fmt.Printf("%d is Divisible by all ------ \n", quotient)
		} else {
			fmt.Printf("%d is not divisible by all \n", quotient)
		}

		// quotient++
	}
}
