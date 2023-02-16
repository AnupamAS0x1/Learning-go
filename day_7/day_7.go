package main

import "fmt"

func main() {

	fmt.Println(" if else in golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {

		result = "exactly 10 login count"
	}

	fmt.Println(result)

	// creating the varial in the running process

	if 9%2 == 0 {
		fmt.Println("Number is even")

	} else {
		fmt.Println("the number is odd")
	}

	if num := 10; num < 10 {
		fmt.Println("num is less than 10")
	} else if num > 10 {
		fmt.Println("num is more than 10")
	} else {
		fmt.Println("num is equal to 10 ")
	}

}
