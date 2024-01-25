package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello Pointers!")

	var ptr0 *int
	fmt.Println("Value of the pointer is ", ptr0)

	var ptr *int = new(int)
	fmt.Println("Value of the pointer is ", ptr)

	// referencing
	var num = 10
	var ptr1 = &num
	fmt.Println("Value of the pointer is ", ptr1)
	fmt.Println("Value of the pointer is ", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("Value of the pointer is ", num)

}
