package main

import (
	"hieu/pkgs/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create mux router
	router := mux.NewRouter().StrictSlash(true)

	// register handler to router
	router.Methods(http.MethodGet).Path("/api/cinema").HandlerFunc(handler.Cinema)
	router.Methods(http.MethodGet).Path("/api/movie").HandlerFunc(handler.MoviesList)
	router.Methods(http.MethodGet).Path("/api/movieid").HandlerFunc(handler.MovieID)
	router.Methods(http.MethodGet).Path("/api/movieManager").HandlerFunc(handler.MoviesManager)
	router.Methods(http.MethodGet).Path("/api/movieSelected").HandlerFunc(handler.MovieSelected)
	router.Methods(http.MethodPost).Path("/api/createShow").HandlerFunc(handler.CreateShow)

	// serve router on port
	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		panic(err)
	}
}
