package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Image struct {
	Id string `json:"id"`
}

var images []Image

// Function that will handle the POST requests made on /endpoint
// It adds an image with input ID in the DataBase
func postFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Cast input request to Image
	var image Image
	_ = json.NewDecoder(r.Body).Decode(&image)
	log.Println("Posting Image with ID = ", image.Id, " in my fake database")

	// Add it in the Database
	images = append(images, image)
}

// Function that will handle the GET requests made on /endpoint
// It looks up the input ID in the database, and returns it if it exists
func getFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// Cast input request to Image
	var image Image
	_ = json.NewDecoder(r.Body).Decode(&image)

	// Go through images database, and return (encode in JSON) image if present
	found := false
	for _, item := range images {
		if item.Id == image.Id {
			log.Println("Image with ID = ", image.Id, " found in database")
			json.NewEncoder(w).Encode(item)
			found = true
			break
		}
	}

	// Return (encode) empty image if not present in the database
	if !found {
		log.Println("Image with ID = ", image.Id, " not found in database")
		json.NewEncoder(w).Encode(&Image{})
	}
}

func main() {
	// Initialize router
	router := mux.NewRouter().StrictSlash(true)

	// Set up router endpoints
	router.Methods("POST").Path("/endpoint").HandlerFunc(postFunction)
	router.Methods("GET").Path("/endpoint").HandlerFunc(getFunction)

	// Listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}
