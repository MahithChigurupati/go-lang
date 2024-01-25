package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"red", "yellow", "green"}

var wg sync.WaitGroup // pointer to a wait group
var mut sync.Mutex    // pointer to a mutex

func main() {

	// go greeter("Hello")
	// greeter("World")

	website := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	for _, site := range website {
		go getStatusCode(site)

		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)

}

// func greeter(s string) {

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	resp, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error:", err)
	}

	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Println(resp.StatusCode)

	// return resp.StatusCode
}
