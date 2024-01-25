package main

import "fmt"

func main() {

	fmt.Println("Structs")
	// no inheritance in Go, no super, no extends

	mahith := Person{
		FirstName: "Mahith",
		LastName:  "C",
		Age:       25,
		Status:   true,
	}

	fmt.Println(mahith)
	fmt.Printf("%+v \n", mahith)
	fmt.Printf("name is %v %v \n", mahith.FirstName, mahith.LastName)

}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Status   bool
}

