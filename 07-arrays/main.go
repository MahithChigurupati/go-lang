package main

import "fmt"

func main()  {

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	// fruits[2] = "Banana"
	fruits[3] = "Grape"

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var vegitables = [3]string{"potato", "tomato", "onion"}

	fmt.Println(vegitables)
	
}