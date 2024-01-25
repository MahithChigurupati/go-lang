package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["es"] = "Spanish"
	languages["en"] = "English"
	languages["fr"] = "French"

	fmt.Println(languages)

	fmt.Println(languages["es"])

	delete(languages, "fr")

	fmt.Println(languages)

	// Iterating over a map
	for key, value := range languages {
		fmt.Printf("for key : %v and for value : %v \n", key, value)
	}
}