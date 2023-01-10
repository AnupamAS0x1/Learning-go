package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "this simple user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter intput ")

	// comman ok syntax  , its like a try catch

	input, _ := reader.ReadString('\n')

	// reader is a imported from bufio package to read the user input
	// The  ReadString here is read until a new line is there.

	fmt.Println("Thanks for reading .", input)

	// Now we will convert the string in to a number and 1 ;

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("added 1 to you input ", numrating+1)
	}

}
