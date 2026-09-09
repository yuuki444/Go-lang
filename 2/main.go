package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /movies", getMovies)
	http.HandleFunc("GET /movies/{id}", getMovieByID)
	http.HandleFunc("POST /movies", createMovie)
	http.HandleFunc("PUT /movies/{id}", updateMovie)
	http.HandleFunc("DELETE /movies/{id}", deleteMovie)

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
