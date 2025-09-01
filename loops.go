package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	for i := range 100 {
		fmt.Println("Happy Birthday Ayo", i)

	}
	// OR

	for j := range 9001 {
		if j <= 8999 {
			fmt.Printf("It's Just %d\n", j)
		} else {
			fmt.Printf("It's over %d!!!\n", j)
		}
	}

	column_scaner := bufio.NewScanner(os.Stdin)
	row_scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the no of columns : ")
	column_scaner.Scan() //Read and store a buffer

	fmt.Print("Enter the rows : ")
	row_scanner.Scan()

	columns, _ := strconv.ParseInt(column_scaner.Text(), 10, 64)
	rows, _ := strconv.ParseInt(row_scanner.Text(), 10, 64)
	var row_increment int64 //Match type with strconv.ParseInt()
	var column_increment int64

	for row_increment = 0; row_increment < rows; row_increment++ {
		for column_increment = 0; column_increment < columns; column_increment++ { // nestted loops
			fmt.Print("#") //It doesn't do all this before moving, it finishes one iteration
		}
		fmt.Print("\n")

	}

}

