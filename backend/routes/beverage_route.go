package routes

import (
	"github.com/emilsto/alma-task/handlers"

	"github.com/gorilla/mux"
)

func BeverageRoute(r *mux.Router) {
	// Get all beverages
	r.HandleFunc("/beverages", handlers.GetBeverages).Methods("GET")
	// Create a new beverage
	r.HandleFunc("/beverages", handlers.CreateBeverage).Methods("POST")
	// Search beverage by name
	r.HandleFunc("/beverages/{query}", handlers.SearchBeverage).Methods("GET")
}
