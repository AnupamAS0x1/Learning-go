package main

import "fmt"

// learning math in go lang

func main() {

	//fmt.Println("welcome to Maths in go ")

	//	var mynumbeOne int = 2
	//	var mynumbetwo float64 = 4

	// fmt.Println("the sum is: ", mynumbeOne+int(mynumbetwo)) bad syntax

	//RANDOM NUMBER
	// rand is a crpyto graphy library for generating random numbers here.

	//myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	//fmt.Println(myRandomNum)

	// POINTERS

	// Why Pointers are needed , as sometimes the variable fo the memory is  passed on which caused some irregularities in the program .So Some time it pass the copied values
	//Pointer help us to pass the memory address of the program .

	//var ptr *int // * is used  to declare the pointer variable

	//fmt.Println("Value of the pointer is ", ptr)

	mynumber := 23

	var ptr = &mynumber

	//& sign is used for referenceing

	fmt.Println("value of actual pointer is ", ptr)  // it pass the memory address
	fmt.Println("value of actual pointer is ", *ptr) // it pass the value inside it

	*ptr = *ptr + 2

	fmt.Println("the actual value now is of mynumber is  ", mynumber)

	//NOTE: The operation will; be perform on the actuale  value

	// NOw SLICES

}
