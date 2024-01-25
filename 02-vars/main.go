package main

import "fmt"

// LoginToken exported
// public type implicitly
const LoginToken string = "asdas"

func main()  {
	var message string = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("The message is of type: %T \n", message)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("The message is of type: %T \n", isTrue)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("The message is of type: %T \n", smallValue)

	var smallFloat float32 = 255.4554545454545
	fmt.Println(smallFloat)
	fmt.Printf("The message is of type: %T \n", smallFloat)

	var smallFloat64 float64 = 255.4554545454545
	fmt.Println(smallFloat64)
	fmt.Printf("The message is of type: %T \n", smallFloat)

	// Default values
	var anotherNumber int
	fmt.Println(anotherNumber)
	fmt.Printf("The message is of type: %T \n", anotherNumber)

	var anotherMessage int
	fmt.Println(anotherMessage)
	fmt.Printf("The message is of type: %T \n", anotherMessage)

	var anotherBool bool
	fmt.Println(anotherBool)
	fmt.Printf("The message is of type: %T \n", anotherBool)

	// implicit type
	var message2 = "Hello, World!"
	fmt.Println(message2)
	fmt.Printf("The message is of type: %T \n", message2)

	// no var style
	message3 := 300000.0
	fmt.Println(message3)
	fmt.Printf("The message is of type: %T \n", message3)

	// public
	fmt.Println(LoginToken)
	fmt.Printf("The message is of type: %T \n", LoginToken)

}
