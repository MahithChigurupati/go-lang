package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.google.com:443/search?q=golang"

func main() {

	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	// fmt.Printf("Scheme: %s\n", result.Scheme)
	// fmt.Printf("Host: %s\n", result.Host)
	// fmt.Printf("Path: %s\n", result.Path)
	// fmt.Printf("Port: %s\n", result.Port())
	fmt.Printf("RawQuery: %s\n", result.RawQuery)

	qParams := result.Query()

	fmt.Printf("qParams: %s\n", qParams)
	fmt.Printf("qParams: %T\n", qParams)

	for k, v := range qParams {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "www.google.com",
		Path: "/search",
		RawQuery: "q=golang",
	}

	anoterUrl := partsOfUrl.String()
	fmt.Println(anoterUrl)
	
}