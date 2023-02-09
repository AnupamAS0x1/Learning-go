package main

import (
	"fmt"
)

// We will learn about maps go lang

func main() {

	fmt.Println("maps in go lang ")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all the languages:", languages)
	fmt.Println("JS Short for:", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all the language:", languages)

	// loops are interesting in golang

	for key, value := range languages {

		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
