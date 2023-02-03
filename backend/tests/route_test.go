package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emilsto/alma-task/config"
	"github.com/emilsto/alma-task/models"
	"github.com/emilsto/alma-task/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

// Struct for testing the routes
var routeTests = []struct {
	name           string
	url            string
	method         string
	expectedStatus int
}{
	{"Get all beverages", "/beverages", "GET", 200},
	{"Search beverage by name", "/beverages/kahvi", "GET", 200},
}

// Clear the database for testing
func ClearDB() {
	// Get the collection
	collection := config.GetCollection(config.Client, config.GetConfig().MongoCollection)
	// Delete all documents
	collection.DeleteMany(context.TODO(), bson.D{})

}

// Add a beverage to the database
func AddBeverage() {
	// Get the collection
	collection := config.GetCollection(config.Client, config.GetConfig().MongoCollection)
	// Create a beverage
	beverage := models.Beverage{
		Name:        "Ethiopian Harrar",
		Weight:      500,
		Price:       10,
		RoastDegree: 3,
		Type:        "Coffee",
	}
	// Insert the beverage
	collection.InsertOne(context.TODO(), beverage)
}

func TestDBConnection(t *testing.T) {
	// Get the Collection name, log it
	fmt.Println("Collection name: ", config.GetConfig().MongoCollection)
	// Clear the database, add a beverage
	ClearDB()
	AddBeverage()
}

// TestRoutes tests the routes
func TestRoutes(t *testing.T) {
	for _, tt := range routeTests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to the handler
			req, err := http.NewRequest(tt.method, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Create a router and pass the request and ResponseRecorder to the handler
			router := mux.NewRouter()
			routes.BeverageRoute(router)
			router.ServeHTTP(rr, req)

			// Check the status code of the response
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}

// TestGetBeverages tests the GetBeverages route, check that the response is not empty, is json and is formatted correctly
func TestGetBeverages(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/beverages", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass the request and ResponseRecorder to the handler
	router := mux.NewRouter()
	routes.BeverageRoute(router)
	router.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that response is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}

	// Check that response is json
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

	// Format the response body to a model Beverage
	var beverages []models.Beverage
	err = json.Unmarshal(rr.Body.Bytes(), &beverages)
	if err != nil {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "application/json")
	}

}

// TestSearchBeverage tests the SearchBeverage route
func TestSearchBeverage(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/beverages/Ethiopian", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass the request and ResponseRecorder to the handler
	router := mux.NewRouter()
	routes.BeverageRoute(router)
	router.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that response is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}

	// Check that response is json
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

	// Format the response body to a model Beverage
	var beverages []models.Beverage
	err = json.Unmarshal(rr.Body.Bytes(), &beverages)
	if err != nil {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "application/json")
	}

	// Check that the response body contains the searched beverage
	if beverages[0].Name != "Ethiopian Harrar" {
		t.Errorf("handler returned wrong body: got %v want %v",
			beverages[0].Name, "Ethiopian")
	}

}

// TestCreateBeverage tests the CreateBeverage route
func TestCreateBeverage(t *testing.T) {
	// Create a request to pass to the handler
	req, err := http.NewRequest("POST", "/beverages", strings.NewReader(`{"name":"Pirkka Costa Rica","weight":500,"price":5.50,"roastDegree":1,"type":"Coffee"}`))
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a router and pass the request and ResponseRecorder to the handler
	router := mux.NewRouter()
	routes.BeverageRoute(router)
	router.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check that response is not empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}

	// Check that response is json
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

	// Format the response body to a model Beverage
	var beverage models.Beverage
	err = json.Unmarshal(rr.Body.Bytes(), &beverage)
	if err != nil {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), "application/json")
	}

	// Check that the response body contains the created beverage
	if beverage.Name != "Pirkka Costa Rica" {
		t.Errorf("handler returned wrong body: got %v want %v",
			beverage.Name, "Pirkka Costa Rica")
	}

}
