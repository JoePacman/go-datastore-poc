package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-datastore-poc/authentication"
	"go-datastore-poc/datastore-film"
	"log"
	"net/http"
	"os"
)

func main() {

	// Using 'gorilla' library to register Endpoints; It provides more functionality compared to Go's 'net/http'
	router := mux.NewRouter().StrictSlash(true)

	// This line registers a handler (GetFilm function) for the '/film' endpoint
	// It also wraps the 'GetFilm' into the 'MustAuth' function
	router.Handle("/film/{title}", authentication.MustAuth(http.HandlerFunc(datastore_film.GetFilm))).
		Methods("GET")
	router.Handle("/film", authentication.MustAuth(http.HandlerFunc(datastore_film.PostFilm))).
		Methods("POST")

	// When deployed into GCP it will get the port from the Environment variable (it may not be 8080 in GCP)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Define headers, origins and methods for CORS
	headersOk := handlers.AllowedHeaders([]string{"authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"HEAD", "GET", "POST", "PUT", "OPTIONS"})
	log.Printf("Started Go Server...")

	// Start to listen for all defined Endpoints
	// Use Gorilla's 'CORS' handler function to wrap app endpoints in it (so CORS is applied to all endpoints)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}
