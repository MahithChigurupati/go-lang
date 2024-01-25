package router

import (
	contoller "github.com/MahithChigurupati/24-mongo/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api//movies", contoller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", contoller.GetMovie).Methods("GET")
	router.HandleFunc("/api/movies", contoller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", contoller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", contoller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", contoller.DeleteAllMovies).Methods("DELETE")

	return router
}
