package main

import "fmt"

func main() {

	fmt.Println("If Else")

	loginCount := 20
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "something else"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	// if with initialization
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	// if err != nil {
	// 	fmt.Println(err)
	// }

}