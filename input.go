package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// fmt.Println("We gettig User Input baby")

	// var scanner = bufio.NewScanner(os.Stdin) // Create a scanner that reads from standard in while inferrig the type
	// fmt.Printf("Enter some text : ")

	// scanner.Scan()             //When to start read(scan) from stdin
	// var input = scanner.Text() //Store what is scanned into a variable as a text
	// scanner.Text() converts what was stored in scanner into text
	// //The equivalent of v = input(lalalala) in python in a way

	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your date of birth : ")

	scanner.Scan()
	var dob, _ = strconv.ParseInt(scanner.Text(), 10, 64) // parse scanner.Text() as an integer
	// err tells go to store any error if it is unable to covert into err, _ >> throw away
	fmt.Printf("You are %d", 2025-dob)

	// fmt.Printf("%q", err)

}
