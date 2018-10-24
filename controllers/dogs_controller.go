package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mojicaj/dog_rescue/db"
	"github.com/mojicaj/dog_rescue/model"
)

// CreateDogHandler handles requests to the create dog endpoint
func CreateDogHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var dog model.Dog

	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		log.Println("failed to decode JSON: ", err)
		http.Error(w, "properly formatted JSON is required", http.StatusUnsupportedMediaType)
		return
	}

	// check that a dog name was provided
	if dog.Name == "" {
		http.Error(w, "a dog's name is required", http.StatusBadRequest)
		return
	}

	// check if dog already exists
	if _, err := db.GetDog(dog.Name); err == nil {
		http.Error(w, "dog name already used", http.StatusBadRequest)
		return
	}

	// doesn't exist, create the new dog
	if err := db.CreateDog(&dog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// convert to JSON to send as the response
	dogBytes, err := json.Marshal(dog)
	if err != nil {
		log.Println("failed to marshal dog into JSON: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", dogBytes)
}

// GetDogHandler handles requests to the get dog endpoint
func GetDogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the name passed as a url parameter
	name := ps.ByName("name")

	// retrieve the specified dog or all dogs if none was specified
	if name == "" {
		// get all stored dogs
		dogs, err := db.GetAllDogs()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// marshal returned dogs into JSON structure
		dogsBytes, err := json.Marshal(dogs)
		if err != nil {
			log.Println("failed to marshal dogs into JSON: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", dogsBytes)
	} else {
		// get the specified dog
		dog, err := db.GetDog(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// marshal returned dog into JSON structure
		dogBytes, err := json.Marshal(dog)
		if err != nil {
			log.Println("failed to marshal dog into JSON: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", dogBytes)
	}
}

// UpdateDogHandler updates the specified dog
func UpdateDogHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// decode and save dog
	dog := model.Dog{}
	if err := json.NewDecoder(r.Body).Decode(&dog); err != nil {
		log.Println("Failed to decode JSON: ", err)
		http.Error(w, "properly formatted JSON is required", http.StatusUnsupportedMediaType)
		return
	}

	// check that a dog name was provided
	if dog.Name == "" {
		http.Error(w, "a dog's name is required", http.StatusBadRequest)
		return
	}

	// check that dog exists in the database
	dogDB, err := db.GetDog(dog.Name)
	if err != nil {
		http.Error(w, "could not find a dog by that name", http.StatusNotFound)
		return
	}

	// only update provided data
	if dog.Age == 0 {
		dog.Age = dogDB.Age
	}
	if dog.Breed == "" {
		dog.Breed = dogDB.Breed
	}
	if dog.Condition == "" {
		dog.Condition = dogDB.Condition
	}
	if dog.Description == "" {
		dog.Description = dogDB.Description
	}
	if dog.ImageURL == "" {
		dog.ImageURL = dogDB.ImageURL
	}
	if dog.Location == "" {
		dog.Location = dogDB.Location
	}
	if dog.Status == "" {
		dog.Status = dogDB.Status
	}
	if dog.Weight == "" {
		dog.Weight = dogDB.Weight
	}

	if err := db.UpdateDog(&dog); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// marshal updated dog into JSON
	dogBytes, err := json.Marshal(dog)
	if err != nil {
		log.Println("Failed to marshal dog into JSON: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", dogBytes)
}

// RemoveDogHandler deletes the specified dog
func RemoveDogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the name passed as a url parameter
	name := ps.ByName("name")

	if err := db.RemoveDog(name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set statuscode and payload
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s deleted", name)
}
