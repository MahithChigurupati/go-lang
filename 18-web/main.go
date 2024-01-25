package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Starting the application...")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response status: %T\n", response)

	defer response.Body.Close() // close the response body when we're done

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println("Response body:", content)
}
