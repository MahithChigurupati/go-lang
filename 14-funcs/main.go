package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	greeter()

	result := adder(2, 3)

	fmt.Println("Result is", result)

	proResult, msg := proAdder(2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("Pro result is", proResult)
	fmt.Println("Message is", msg)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int,string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello"
}

func greeter() {
	fmt.Println("Hello")

}