package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MahithChigurupati/24-mongo/router"
)

func main() {
	router := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("Server started on the port 8080...")
}
