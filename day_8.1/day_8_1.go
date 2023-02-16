package main

import (
	"fmt"
)

func main() {

	fmt.Println("leaning about loops in go")

	days := []string{"sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)
	// for i := 0; i < len(days); i++ {

	// 	fmt.Println(days[i])
	// }

	for i := range days {
		fmt.Println(days[i])
		// here it return the index not he value
	}
	for _, day := range days {

		fmt.Printf("index is  and value is %v\n", day)

	}
	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {

			goto lco

		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println(" value is:", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println(" jumping at anupamas0x1.github.com")

}
