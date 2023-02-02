package main

import (
	"log"
	"net/http"

	"github.com/emilsto/alma-task/config"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Set up the router
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)

	// Connect to the database
	config.ConnectDB()


	// Start the server, listen on port 5000 and use gorilla/mux as the router
	log.Fatal(http.ListenAndServe(":5000", handler))
}
