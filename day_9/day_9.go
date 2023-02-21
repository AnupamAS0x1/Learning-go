package main

import "fmt"

func main() {
	fmt.Println("welcome to function in golang")

	greeter()

}

func greeter() {
	fmt.Println("Namstey from golang")
}

func adder(valOne int, valtwo int) int {
	return valOne + valtwo

}
