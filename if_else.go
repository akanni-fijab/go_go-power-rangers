package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Create Scanner
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name ")
	scanner.Scan()

	name := scanner.Text()

	if name == "Jose" || name == "Fijabi" {
		fmt.Printf("Welcome %s", name)
	} else if name == "Amina" {
		fmt.Printf("Welcome %s", name)
	} else {
		fmt.Printf("Go away %s, you fuck ass \n", name)
	}

}
