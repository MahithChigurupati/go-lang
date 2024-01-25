package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello from Go!"
	
	file, err := os.Create("myFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)
	
	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)

	fmt.Println("length is", length)
	defer file.Close()

	readFile("myFile.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	checkError(err)

	fmt.Println(string(databyte))
	

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}