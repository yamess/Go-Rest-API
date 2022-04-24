package services

import (
	"encoding/json"
	"github.com/api_app/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Set the ID and creation date and append to the db
	user.ID = uuid.New()
	user.CreatedAt = time.Now().UTC()

	result := user.CreateRecord()

	if result.RowsAffected > 0 {
		log.Println("New User created")

		err = json.NewEncoder(w).Encode(&user)
		if err != nil {
			panic(err)
		}
	} else {
		err = json.NewEncoder(w).Encode("Enable to insert record")
		if err != nil {
			panic(err)
		}
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users models.Users

	// Get db object
	users.ReadRecords()

	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		panic("Enable to handle request")
	}
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	vars := mux.Vars(r)
	id := vars["id"]

	result := user.ReadRecord("id", id)

	if result.Error != nil {
		log.Println("Enable to get data from database")
		panic(result.Error)
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Enable to encode response")
		panic(err.Error())
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("Enable to decode the request body")
		log.Println(err.Error())
		panic(err.Error())
	}

	result := user.UpdateRecord()
	if result != nil {
		log.Println("Enable to update data in the database")
		panic(result.Error)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("Enable to encode response data")
		panic(err.Error())
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User deleted")
}
