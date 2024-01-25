package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Hello, playground")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func greeter() {
	fmt.Println("Hello, playground")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, playground"))
}
