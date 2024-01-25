package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"Apple", "Orange", "Banana", "Grape"}

	fmt.Printf("type of fruits is %T \n", fruits)
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	fruits = append(fruits, "Mango", "Pineapple")
	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println(fruits)

	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("After sorting")
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove values from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby", "php"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
