package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random number is", rand.Intn(6)+1)

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number 1")
	case 2:
		fmt.Println("Dice number 2")
	case 3:
		fmt.Println("Dice number 3")
	case 4:
		fmt.Println("Dice number 4")
		fallthrough
	case 5:
		fmt.Println("Dice number 5")
	default:
		fmt.Println("Dice number 6")
	}
}
