package router

import (
	"net/http"

	"github.com/gabey14/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// Home route
	router.HandleFunc("/", serveHome).Methods("GET")

	// Getting all movies
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")

	// Creating a movie
	router.HandleFunc("/api/movie", controller.CreatMovie).Methods("POST")

	// Mark/Update a movie as watched
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	// Delete a movie
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	// Delete all movies
	router.HandleFunc("/api/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")

	return router

}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Welcome to Mongo DB API</h1>"))
}
