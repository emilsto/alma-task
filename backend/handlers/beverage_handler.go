package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/emilsto/alma-task/config"
	"github.com/gorilla/mux"

	"github.com/emilsto/alma-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection = config.GetCollection(config.Client, config.GetConfig().MongoCollection)

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

// Create a new beverage
func CreateBeverage(w http.ResponseWriter, r *http.Request) {
	var beverage models.Beverage

	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&beverage); err != nil {
		respondWithError(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(beverage)

	// Insert the beverage into the collection
	if _, err := collection.InsertOne(context.TODO(), beverage); err != nil {
		fmt.Println(err.Error())
		respondWithError(w, "Error inserting beverage", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(beverage); err != nil {
		log.Println(err)
		respondWithError(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

// Search for a beverage
func SearchBeverage(w http.ResponseWriter, r *http.Request) {
	var beverages []models.Beverage

	//set and check the params
	params := mux.Vars(r)
	query := params["query"]

	if query == "" {
		//send all beverages
		println("No query, sending all beverages")
		GetBeverages(w, r)
		return
	}

	cur, err := collection.Find(context.TODO(), bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: query, Options: "i"}}})
	if err != nil {
		respondWithError(w, "No beverages found!", http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var beverage models.Beverage
		err := cur.Decode(&beverage)
		if err != nil {
			respondWithError(w, "Invalid req payload", http.StatusBadRequest)
		}
		beverages = append(beverages, beverage)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(beverages); err != nil {
		log.Println(err)
		respondWithError(w, "Error endocing JSON", http.StatusInternalServerError)
	}

}
