package main

import (
	"fmt"
	"math/rand"
	"time"
)

// learning swithc case here

func main() {
	fmt.Println("learning the switch case here")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("dice value is 2 and you can open")
	case 3:
		fmt.Println("dice value is 3 and you can open")
	case 4:
		fmt.Println("dice value is 4 and you can open")
	case 5:
		fmt.Println("dice value is 5 and you can open")
	case 6:
		fmt.Println("dice value is 6 and you can open")

	}

}
