package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World")

	var x int = 42
	var y float64 = 40

	fmt.Println(x + int(y))

	fmt.Println(rand.Intn(5))

}
