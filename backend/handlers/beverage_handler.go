package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/emilsto/alma-task/config"

	"github.com/emilsto/alma-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = config.GetCollection(config.Client, "beverages")

// Get all beverages
func GetBeverages(w http.ResponseWriter, r *http.Request) {
	var beverages []models.Beverage

	// Find all beverages
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		respondWithError(w, "Error finding beverages", http.StatusInternalServerError)
		return
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var beverage models.Beverage
		// Decode the beverage
		if err := cur.Decode(&beverage); err != nil {
			log.Fatal(err)
		}
		// Add the beverage to the slice
		beverages = append(beverages, beverage)
	}

	// Close the cursor
	cur.Close(context.TODO())

	// Return the beverages as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(beverages); err != nil {
		log.Println(err)
		respondWithError(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

}
