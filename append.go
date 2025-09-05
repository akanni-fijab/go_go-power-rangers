package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Make a simple program to collect use input store it in a slice and print it out later
	//Only integers now or a single type

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How may variable do you wanna store")
	scanner.Scan()

	stop_increment, _ := strconv.ParseInt(scanner.Text(), 10, 64) // --> int64
	user_store := make([]int64, stop_increment)                   //Storage Slice

	var i int64 //Match case

	for i = range stop_increment {
		scan_input := bufio.NewScanner(os.Stdin) //Loop over to get input at each step
		fmt.Print("Enter Them ")
		scan_input.Scan()
		actual_input, _ := strconv.ParseInt(scan_input.Text(), 10, 64)
		// user_store = append(user_store, actual_input)
		user_store[i] = actual_input

	}
	fmt.Print(user_store)

}

