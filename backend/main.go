package main

import (
	"log"
	"net/http"

	"github.com/emilsto/alma-task/config"
	"github.com/emilsto/alma-task/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Set up the router
	r := mux.NewRouter()
	handler := cors.Default().Handler(r)

	api := r.PathPrefix("/api/v1").Subrouter()

	// Connect to the database
	config.ConnectDB()

	// Set up the routes
	routes.BeverageRoute(api)

	// Start the server, listen on port 5000 and use gorilla/mux as the router
	log.Fatal(http.ListenAndServe(":5000", handler))
}
