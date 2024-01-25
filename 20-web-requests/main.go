package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Hello, playground")

	// performGetRequest()
	// performPostRequest()
	performPostFormDataRequest()
}

func performGetRequest() {
	const url = "https://www.google.com:443/search?q=golang"

	result, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer result.Body.Close()
	fmt.Println(result.StatusCode)
	fmt.Println(result.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(result.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func performPostRequest() {

	const url = "https://httpbin.org/post"

	payload := strings.NewReader(`
		{
			"hello": "world"
		}
	`)

	result, err := http.Post(url, "application/json", payload)

	if err != nil {
		panic(err)
	}

	defer result.Body.Close()

	content, _ := ioutil.ReadAll(result.Body)

	fmt.Println(string(content))

}

func performPostFormDataRequest() {

	const myurl = "https://httpbin.org/post"

	data := url.Values{}
	data.Add("name", "John Doe")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
