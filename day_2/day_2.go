package main

import "fmt"

func main() {
	//var is used for declaring a variable

	var name string = "anupam"
	fmt.Println(name)
	fmt.Printf("Variable of the type: %T \n", name)
	// %T for for the type

	var smallVal uint8 = 255
	fmt.Println(name)
	fmt.Printf("Variable of the type: %T \n", smallVal)
	// Unit is a untyppe integer constant

	var tryingnew int
	fmt.Println(tryingnew)
	fmt.Printf("Variable of the type: %T \n", tryingnew)

	//implicit type
	var impli = "hello_world"
	fmt.Println(impli)

	//novarstyle

	numberofuser := 30000
	fmt.Println(numberofuser)

	//:= walrus operator in go cant be use in global

	//In go the Capital letter from the function define that it is publically for use.
	// I learned about lexer and types on go, a lexer means to check the syntax of the laguage and type means int,string etc
	//fp for print line short cut
}
